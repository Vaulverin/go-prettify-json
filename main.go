package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"encoding/json"
	"bytes"
)
var (
	jsonBegin = byte('{')
	jsonEnd = byte('}')
)

func main() {
	fileContent := getInputFileContent()
	stopOnError, formatTypesToPrettify := getFlags()
	prettyContent := prettifyContent(fileContent, stopOnError, formatTypesToPrettify)
	fmt.Println(string(prettyContent))
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
func getInputFileContent() []byte {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("No input file!")
	}
	return data
}

func prettifyJson(content []byte)[]byte {
	for i:= 0; i < len(content); i++ {
		beginIndex := bytes.IndexByte(content[i:], jsonBegin)
		if beginIndex != -1 {
			endIndex := findJsonEnd(content[beginIndex + 1:])
			if endIndex != -1 {
				jsonData := parseJson(content[ beginIndex : endIndex ])
				if len(jsonData) != 0 {
					newLine := []byte{'\r', '\n'}
					jsonData = append(newLine, jsonData...)
					jsonData = append(jsonData, newLine...)
					content = bytes.Replace(content, content[ beginIndex : endIndex ], jsonData, -1)
					i = beginIndex + len(jsonData)
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
	beginsCount := 1
	for i, symbol := range content {
		switch symbol {
		case jsonBegin:
			beginsCount++
		case jsonEnd:
			beginsCount--
			if beginsCount == 0 {
				return i
			} else if beginsCount < 0 {
				return -1
			}
		}
	}
	return -1
}

func parseXml(content []byte) []byte {
	return content
}
