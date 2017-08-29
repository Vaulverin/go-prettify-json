package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"encoding/json"
)

func main() {
	fileContent := getInputFileContent()
	stopOnError, formatTypesToPrettify := getFlags()
	prettyContent := prettifyContent(fileContent, stopOnError, formatTypesToPrettify)
	fmt.Println(string(prettyContent))
}
func prettifyContent(content []byte, stopOnError bool, formatTypesToPrettify string) []byte {
	if strings.Contains(formatTypesToPrettify, "xml") {
		content = parseXml(content)
	}
	if strings.Contains(formatTypesToPrettify, "json") {
		content = parseJson(content)
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
func parseJson(content []byte) []byte {
	var obj interface{}
	json.Unmarshal(content, &obj)
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println("Bad json!")
	}
	return data
}
func parseXml(content []byte) []byte {
	return content
}
