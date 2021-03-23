package wordide

import (
	"encoding/xml"
	"strconv"
)

func GetIntAttr(attrs *[]xml.Attr, name string) int {
	for _, attr := range *attrs {
		if attr.Name.Local == name {
			i, _ := strconv.Atoi(attr.Value)
			return i
		}
	}

	return -1 // FIXME: We need to return an error!
}

type DocumentText struct {
	Data string
}

func (dt *DocumentText) ParseElement(d *xml.Decoder) xml.Token {
	token, _ := d.Token()
	switch tt := token.(type) {
	case xml.StartElement:
		switch tt.Name.Local {
		case "p":
			for {
				switch tt := dt.ParseElement(d).(type) {
				case xml.EndElement:
					if tt.Name.Local == "p" {
						dt.Data += "\n"
						return tt
					}
				}
			}

		case "s":
			for i := 0; i < GetIntAttr(&tt.Attr, "c"); i += 1 {
				dt.Data += " "
			}

		case "tab":
			dt.Data += "\t"

		case "span":
			for {
				switch tt := dt.ParseElement(d).(type) {
				case xml.EndElement:
					if tt.Name.Local == "span" {
						return tt
					}
				}
			}

		default:
			for {
				switch x := dt.ParseElement(d).(type) {
				case xml.EndElement:
					if tt.Name.Local == x.Name.Local {
						return x
					}
				}
			}
		}

	case xml.EndElement:
		return tt

	case xml.CharData:
		dt.Data += string(tt)
	}

	return token
}

func (dt *DocumentText) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		switch x := dt.ParseElement(d).(type) {
		case xml.EndElement:
			if x.Name.Local == "text" {
				return nil // Yes we need to return errors some time
			}
		}
	}
}

type DocumentBody struct {
	XMLName xml.Name     `xml:"body"`
	Text    DocumentText `xml:"text"`
}

type DocumentContent struct {
	XMLName xml.Name `xml:"document-content"`
	Body    DocumentBody
}

func Parse(file []byte) (*DocumentContent, error) {
	content := new(DocumentContent)

	err := xml.Unmarshal([]byte(file), content)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (content *DocumentContent) String() string {
	return content.Body.Text.Data
}
