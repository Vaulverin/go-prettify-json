package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"errors"
)

var (
	jsonBegin = []byte{'{', '['}
	jsonEnd   = []byte{'}', ']'}
)

func main() {
	fileContent, err := getInputFileContent()
	if err == nil {
		stopOnError, formatTypesToPrettify := getFlags()
		prettyContent := prettifyContent(fileContent, stopOnError, formatTypesToPrettify)
		fmt.Println(string(prettyContent))
	}
}
func prettifyContent(content []byte, stopOnError bool, formatTypesToPrettify string) []byte {
	if strings.Contains(formatTypesToPrettify, "xml") {

	}
	if strings.Contains(formatTypesToPrettify, "json") {
		content = prettifyJson(content)
	}
	return content
}
func getFlags() (bool, string) {
	return false, "json,xml"
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
	return nil, errors.New("No input file")
}

func prettifyJson(content []byte) []byte {
	for i := 0; i < len(content); i++ {
		beginIndex := bytes.IndexByte(content[i:], jsonBegin[0])
		tempIndex := bytes.IndexByte(content[i:], jsonBegin[1])
		if beginIndex > tempIndex && tempIndex != -1 {
			beginIndex = tempIndex
		}
		if beginIndex != -1 {
			endIndex := findJsonEnd( content[ beginIndex :])
			if endIndex != -1 {
				length := beginIndex + endIndex + 1
				jsonData := parseJson( content[ beginIndex : length ] )
				if len(jsonData) != 0 {
					newLine := []byte{'\r', '\n'}
					jsonData = append(newLine, jsonData...)
					jsonData = append(jsonData, newLine...)
					newContent := append(content[: beginIndex ], jsonData...)
					tailBeginIndex := beginIndex + endIndex + 1
					if tailBeginIndex >= len(content) {
						content = newContent
					} else {
						content = append(newContent, content[ tailBeginIndex : ]...)
					}
					i = beginIndex + len(jsonData) - 1
				}
			}
		}
	}
	return content
}

func parseJson(content []byte) []byte {
	var data bytes.Buffer
	json.Indent(&data, content, "", "  ")
	return data.Bytes()
}

func findJsonEnd(content []byte) int {
	signIndex := bytes.IndexByte(jsonBegin, content[0])
	if signIndex != -1 {
		beginsCount := 0
		for i, symbol := range content {
			switch symbol {
			case jsonBegin[ signIndex ]:
				beginsCount++
			case jsonEnd[ signIndex ]:
				beginsCount--
			}
			if beginsCount == 0 {
				return i
			} else if beginsCount < 0 {
				break
			}
		}
	}
	return -1
}

func parseXml(content []byte) []byte {
	return content
}