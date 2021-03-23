package main

import (
	"fmt"
	"io"
	"os"

	"github.com/FiftyLinesOfCode/wordide/wordide"
)

const (
	PROG_NAME          string = "wide"
	PROG_VERSION_MAJOR int    = 1
	PROG_VERSION_MINOR int    = 0
)

func printUsage() {
	println("USAGE: wide <command> <filename>.odt")
	println("Available commands:")
	println("\tcompile - Outputs plain text in stdout")
	println("\thelp - displays this page")
}

func main() {
	args := os.Args
	if len(args) <= 2 {
		printUsage()
		return
	}

	switch args[1] {
	case "help":
		printUsage()
		return

	case "compile":
		ctx, err := wordide.OpenContext(args[2])
		if err != nil {
			println("Can't open a context. Are you sure the file specified exists?")
			return
		}
		defer ctx.Close()

		file, err := ctx.GetFile("content.xml")
		if err != nil {
			println("The specified file doesn't contain a content.xml. Are you sure it's a valid .odt file?")
			return
		}

		reader, _ := file.Open()
		fileContent, err := io.ReadAll(reader)
		if err != nil {
			println("Failed to read from content.xml")
			return
		}

		generated, err := wordide.Parse(fileContent)
		if err != nil {
			println("Couldn't parse content.xml. Are you sure it's a valid .odt file?")
			return
		}
		fmt.Println(generated)

	default:
		printUsage()
		return
	}
}
