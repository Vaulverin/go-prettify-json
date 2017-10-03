package domFormatter

import (
	"bytes"
	"encoding/json"
	"regexp"
	"go/doc"
	"encoding/xml"
)

type Formatter struct {
	firstDOMElement []byte
	lastDOMElement []byte
}

/*

If no begin index found, function returns -1.
*/
func (f Formatter) FindBeginIndex(content []byte) int {
	beginIndex := -1
	pattern, err := regexp.Compile("<([a-zA-Z]+)")
	if err == nil {
		loc := pattern.FindIndex(content)
		if loc != nil {
			// Set name of the first element
			f.firstDOMElement = content[loc[0]:loc[1]]
			f.lastDOMElement = append([]byte {'<', '/'}, content[loc[0] + 1:loc[1]]...)
			return loc[0]
		}
	}
	return beginIndex
}

/*
If no end index found, function returns -1.
 */
func (f Formatter) FindEndIndex(content []byte) int {
	pattern, err := regexp.Compile("<([a-zA-Z\\/]+)")
	if err == nil {
		matches := pattern.FindAllIndex(content, 0)
		if matches != nil {
			tagCounter := 0
			for _, match := range matches {
				switch content[match[0]:match[1]] {
				case f.firstDOMElement:
					tagCounter++
				case f.lastDOMElement:
					tagCounter--
				}
				if tagCounter == 0 {
					endIndex := bytes.IndexByte(content[match[0]:], byte('>'))
					if endIndex != -1 {
						endIndex += match[0]
					}
					return endIndex
				} else if tagCounter < 0 {
					break
				}
			}
		}
	}
	return -1
}

/*
Simple JSON formatting.
If JSON string is corrupted, function returns empty array.
*/
func (f Formatter) Format(content []byte) []byte {
	var data bytes.Buffer

	json.Indent(&data, content, "", "  ")
	return data.Bytes()
}