package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"log-formatter/jsonFormatter"
	"log-formatter/xmlFormatter"
	"log-formatter/htmlFormatter"
	"log-formatter/domFormatter"
)

func main() {
	fileContent, err := getInputFileContent()
	if err == nil {
		formatters := getFormatters()
		for _, formatter := range formatters {
			fileContent = formatInput(fileContent, formatter)
		}
		fmt.Println(string(fileContent))
	}
}
// Interface for all formatters
type iFormatter interface {
	FindBeginIndex(content []byte) int
	FindEndIndex(content []byte) int
	Format(content []byte) []byte
}

// Extract formatters from terminal flags
func getFormatters() []iFormatter {
	// TODO implement flags parsing
	return []iFormatter{
		jsonFormatter.Formatter{},
		htmlFormatter.Formatter{
					Formatter: domFormatter.Formatter{},
				},
		xmlFormatter.Formatter{
			Formatter: domFormatter.Formatter{},
		}}
}

// Trying to get input file from STDin
func getInputFileContent() ([]byte, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := ioutil.ReadFile("c://go-path/src/log-formatter/t.txt")
		//data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	// TODO test this check (no error when no file - this is bad)
	return nil, errors.New("no input file")
}

// Go through file, looking for text to format
func formatInput(fileContent []byte, formatter iFormatter) []byte {
	for i := 0; i < len(fileContent); i++ {
		beginIndex := formatter.FindBeginIndex(fileContent[i:])
		if beginIndex != -1 {
			endIndex := formatter.FindEndIndex(fileContent[beginIndex:])
			if endIndex != -1 {
				textLength := beginIndex + endIndex + 1
				formattedText := formatter.Format(fileContent[beginIndex:textLength])
				if len(formattedText) != 0 {

					// Append new lines at the begin and at the end of formatted text
					newLine := []byte{'\r', '\n'}
					formattedText = append(newLine, formattedText...)
					formattedText = append(formattedText, newLine...)

					// Replace old text with new formatted one
					newContent := append(fileContent[:beginIndex], formattedText...)
					tailBeginIndex := beginIndex + endIndex + 1
					if tailBeginIndex >= len(fileContent) {
						fileContent = newContent
					} else {
						fileContent = append(newContent, fileContent[tailBeginIndex:]...)
					}

					// Move index to the end of formatted text
					i = beginIndex + len(formattedText) - 1
				}
			}
		}
	}
	return fileContent
}
