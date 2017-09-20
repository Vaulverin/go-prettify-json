package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"errors"
	"go-prettify-json-xml/jsonParser"
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

type iParser interface {
	FindBeginIndex(content []byte) int
	FindEndIndex(content []byte) int
	Parse(content []byte) []byte
}

func getFlags() []iParser {
	return []iParser{jsonParser.Parser{}}
}

func getInputFileContent() ( []byte, error ){
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := ioutil.ReadFile("C:/go-path/src/go-prettify-json-xml/t.txt")
		//data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, errors.New("No input file.")
}

func prettifyContent(content []byte, parser iParser) []byte {
	for i := 0; i < len(content); i++ {
		beginIndex := parser.FindBeginIndex( content[i:] )
		if beginIndex != -1 {
			endIndex := parser.FindEndIndex( content[ beginIndex :] )
			if endIndex != -1 {
				length := beginIndex + endIndex + 1
				data := parser.Parse( content[ beginIndex : length ] )
				if len(data) != 0 {
					newLine := []byte{'\r', '\n'}
					data = append(newLine, data...)
					data = append(data, newLine...)
					newContent := append(content[: beginIndex ], data...)
					tailBeginIndex := beginIndex + endIndex + 1
					if tailBeginIndex >= len(content) {
						content = newContent
					} else {
						content = append(newContent, content[ tailBeginIndex : ]...)
					}
					i = beginIndex + len(data) - 1
				}
			}
		}
	}
	return content
}