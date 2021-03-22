package main

import (
	"fmt"
	"io"

	"github.com/FiftyLinesOfCode/wordide/wordide"
)

func main() {
	ctx, err := wordide.OpenContext("test.odt")
	if err != nil {
		fmt.Println("Could not open context")
		return
	}
	defer ctx.Close()

	file, err := ctx.GetFile("content.xml")
	if err != nil {
		fmt.Println("Can't find content.xml")
		return
	}

	reader, _ := file.Open()
	fileContent, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Failed to read from content.xml")
		return
	}
	fmt.Printf("%s\n", fileContent)

	content, err := wordide.Parse(fileContent)
	if err != nil {
		fmt.Printf("Cannot parse the file, reported error: %v\n", err)
	}

	fmt.Println(content.String())
}
