package app

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"

	"imkernel.dev/logos"
)

func ParseOSID(id string) (book, chapter, verse string, err error) {
	parsedString := strings.Split(id, ".")
	if len(parsedString) != 3 {
		return "", "", "", errors.New("unable to parse the input")
	}
	return parsedString[0], parsedString[1], parsedString[2], nil
}

func ParseOSISFile(f io.Reader) (*logos.Bible, error) {
	bible := &logos.Bible{}
	bibleXML := logos.BibleXML{}
	if err := xml.NewDecoder(f).Decode(&bibleXML); err != nil {
		return nil, fmt.Errorf("unable to decode the osis xml format %w", err)
	}
	bible.Title = bibleXML.OsisText.Header.Work.Title
	bible.Description = bibleXML.OsisText.Header.Work.Description
	bible.Lang = bibleXML.OsisText.Lang
	bible.Identifier = bibleXML.OsisText.Header.Work.Identifier.Text
	bible.Rights = bibleXML.OsisText.Header.Work.Rights
	bible.Publisher = bibleXML.OsisText.Header.Work.Publisher

	booksXML := bibleXML.OsisText.Div
	for _, bXML := range booksXML {
		book := logos.Book{}
		t := bXML.Type
		if t == "book" {
			book.Name = bXML.OsisID
		}
		// Read the Chapters here and internally read the verses, only after that append the book to the bible
		chaptersXML := bXML.Chapter
		for _, cXML := range chaptersXML {
			chapter := logos.Chapter{}
			chapter.ID = cXML.OsisID
			versesXML := cXML.Verse
			for _, vXML := range versesXML {
				verse := logos.Verse{}
				verse.ID = vXML.OsisID
				verse.ChapterID = cXML.OsisID
				verse.Text = vXML.Text
				chapter.Verses = append(chapter.Verses, verse)
			}
			book.Chapters = append(book.Chapters, chapter)
		}

		bible.Books = append(bible.Books, book)
	}

	return bible, nil
}
