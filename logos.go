package logos

import "encoding/xml"

type BibleXML struct {
	XMLName        xml.Name `xml:"osis"`
	Text           string   `xml:",chardata"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xsi            string   `xml:"xsi,attr"`
	OsisText       struct {
		Text        string `xml:",chardata"`
		OsisIDWork  string `xml:"osisIDWork,attr"`
		Lang        string `xml:"lang,attr"`
		OsisRefWork string `xml:"osisRefWork,attr"`
		Header      struct {
			Text         string `xml:",chardata"`
			RevisionDesc struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				P    string `xml:"p"`
			} `xml:"revisionDesc"`
			Work struct {
				Text        string `xml:",chardata"`
				OsisWork    string `xml:"osisWork,attr"`
				Title       string `xml:"title"`
				Contributor string `xml:"contributor"`
				Creator     []struct {
					Text string `xml:",chardata"`
					Role string `xml:"role,attr"`
				} `xml:"creator"`
				Subject     string `xml:"subject"`
				Date        string `xml:"date"`
				Description string `xml:"description"`
				Publisher   string `xml:"publisher"`
				Type        struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"type"`
				Identifier struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"identifier"`
				Source   string `xml:"source"`
				Language struct {
					Text string `xml:",chardata"`
					Type string `xml:"type,attr"`
				} `xml:"language"`
				Relation  string `xml:"relation"`
				Coverage  string `xml:"coverage"`
				Rights    string `xml:"rights"`
				Scope     string `xml:"scope"`
				RefSystem string `xml:"refSystem"`
			} `xml:"work"`
		} `xml:"header"`
		Div []struct {
			Text    string `xml:",chardata"`
			Type    string `xml:"type,attr"`
			OsisID  string `xml:"osisID,attr"`
			Chapter []struct {
				Text   string `xml:",chardata"`
				OsisID string `xml:"osisID,attr"`
				Verse  []struct {
					Text   string `xml:",chardata"`
					OsisID string `xml:"osisID,attr"`
				} `xml:"verse"`
				Title []string `xml:"title"`
			} `xml:"chapter"`
		} `xml:"div"`
	} `xml:"osisText"`
}
