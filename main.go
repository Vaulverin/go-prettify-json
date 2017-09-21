package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"log-formatter/jsonFormatter"
)

func main() {
	fileContent, err := getInputFileContent()
	if err == nil {
		var prettyContent []byte
		parsers := getFlags()
		for _, parser := range parsers {
			prettyContent = prettifyContent(fileContent, parser)
		}
		fmt.Println(string(prettyContent))
	}
}

type iFormatter interface {
	FindBeginIndex(content []byte) int
	FindEndIndex(content []byte) int
	Parse(content []byte) []byte
}

func getFlags() []iFormatter {
	return []iFormatter{jsonFormatter.Formatter{}}
}

// Trying to get input file from STDin
func getInputFileContent() ([]byte, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := ioutil.ReadFile("D:/Projects/go-projects/src/log-formatter/t.txt")
		//data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New("no input file")
}
//
func prettifyContent(content []byte, parser iFormatter) []byte {
	// Loop through file, looking for text to parse
	for i := 0; i < len(content); i++ {
		beginIndex := parser.FindBeginIndex(content[i:])
		if beginIndex != -1 {
			endIndex := parser.FindEndIndex(content[beginIndex:])
			if endIndex != -1 {
				textLength := beginIndex + endIndex + 1
				formattedText := parser.Parse(content[beginIndex:textLength])
				if len(formattedText) != 0 {
					// Append new lines at the begin and at the end of formattedText
					newLine := []byte{'\r', '\n'}
					formattedText = append(newLine, formattedText...)
					formattedText = append(formattedText, newLine...)
					// Replace old text with new formatted one
					newContent := append(content[:beginIndex], formattedText...)
					tailBeginIndex := beginIndex + endIndex + 1
					if tailBeginIndex >= len(content) {
						content = newContent
					} else {
						content = append(newContent, content[tailBeginIndex:]...)
					}
					// Move index to the end of formatted text
					i = beginIndex + len(formattedText) - 1
				}
			}
		}
	}
	return content
}
