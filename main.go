package main

import (
	"fmt"
	"strings"
)

func main() {
	fileContent := getInputFileContent()
	stopOnError, formatTypesToPrettify := getFlags()
	prettyContent := prettifyContent(fileContent, stopOnError, formatTypesToPrettify)
	fmt.Println(prettyContent)
}
func prettifyContent(content string, stopOnError bool, formatTypesToPrettify string) string {
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
func getInputFileContent() string {
	return ""
}
func parseJson(content string) string {
	return content
}
func parseXml(content string) string {
	return content
}
