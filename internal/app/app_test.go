package app

import (
	"os"
	"testing"
)

func TestParseOSID(t *testing.T) {
	testCases := []struct {
		desc   string
		input  string
		output [3]string
		error  error
	}{
		{
			desc:   "Positive test case",
			input:  "Gen.1.12",
			output: [3]string{"Gen", "1", "12"},
			error:  nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			book, chapter, verse, err := ParseOSID(tC.input)
			if tC.output[0] != book || tC.output[1] != chapter || tC.output[2] != verse || tC.error != err {
				t.Logf("expected the book to be %s, chapter to be %s and verse to be %s but found the book to be %s, chapter to be %s and verse to be %s", tC.output[0], tC.output[1], tC.output[2], book, chapter, verse)
				t.Fail()
			}
		})
	}
}

func TestParseOSISFile(t *testing.T) {
	f, err := os.Open("../../data/en/kjv.xml")
	if err != nil {
		t.Logf("unable to find the osis xml file %v", err)
		t.FailNow()
	}
	if err != nil {
		t.Logf("unable to parse the osis xml file %v", err)
		t.FailNow()
	}
	bible, err := ParseOSISFile(f)
	if err != nil {
		t.Logf("unable to parse the osis xml file %v", err)
		t.FailNow()
	}
	if len(bible.Books) != 66 {
		t.Logf("number of books expected in the bible is %d but found %d", 66, len(bible.Books))
	}

}
