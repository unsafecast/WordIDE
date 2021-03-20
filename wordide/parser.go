package wordide

import (
	"encoding/xml"
	"fmt"
)

type Space struct {
	Number int `xml:"c,attr,omitempty"`
}

type Paragraph struct {
	Space Space       `xml:"s"`
	Data  string      `xml:",chardata"`
	Spans []Paragraph `xml:"span"`
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
		//for _, s := range p.Spaces {
		for i := 0; i <= p.Space.Number; i += 1 {
			str += " "
		}
		if p.Space.Number == 0 {
			str += " "
		}
		//}
		str += p.Data

		for _, s := range p.Spans {
			str += s.Data
		}
		str += "\n"
	}

	return str
}
