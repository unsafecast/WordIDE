package wordide

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type Space struct {
	Number int `xml:"c,attr,omitempty"`
}

type Element struct {
	Data string
}

type Paragraph struct {
	Data []Element `xml:",any"`
}

func (e *Element) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "s":
		for _, x := range start.Attr {
			if x.Name.Local == "c" {
				x, _ := strconv.Atoi(x.Value)
				for i := 0; i < x; i += 1 {
					e.Data += " "
				}
			}
		}
		d.Skip()

	case "span":
		d.Skip() // TODO: This is where I left off
		x, _ := d.Token()
		cd := x.(xml.CharData)
		e.Data += string(cd)
		d.Skip()

	default:
		e.Data += start.Name.Local
		d.Skip()
	}

	return nil
}

type DocumentText struct {
	XMLName    xml.Name    `xml:"text"`
	Paragraphs []Paragraph `xml:"p"`
}

type DocumentBody struct {
	XMLName xml.Name `xml:"body"`
	Text    DocumentText
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

	fmt.Printf("\n%#v\n", content)

	return content, nil
}

func (content *DocumentContent) String() string {
	str := ""
	for _, p := range content.Body.Text.Paragraphs {
		for _, x := range p.Data {
			str += x.Data
		}
		str += "\n"
	}

	return str
}
