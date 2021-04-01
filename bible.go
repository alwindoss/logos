package logos

type Bible struct {
	Title       string
	Description string
	Publisher   string
	Identifier  string
	Lang        string
	Rights      string
	Books       []Book
}

type Book struct {
	Name     string
	Chapters []Chapter
}

type Chapter struct {
	ID     string
	Verses []Verse
}

type Verse struct {
	ChapterID string
	ID        string
	Text      string
}
