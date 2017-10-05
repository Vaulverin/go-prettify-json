package domFormatter

import (
	"bytes"
	"github.com/go-xmlfmt/xmlfmt"
	"regexp"
	"fmt"
)

type Formatter struct {
	FirstDOMElement []byte
	LastDOMElement  []byte
}

/*

If no begin index found, function returns -1.
*/
func (f Formatter) FindBeginIndex(content []byte) int {
	beginIndex := -1
	pattern, err := regexp.Compile("<[a-zA-Z]+")
	if err == nil {
		loc := pattern.FindIndex(content)
		if loc != nil {
			// Set name of the first element
			f.FirstDOMElement = content[loc[0]:loc[1]]
			f.LastDOMElement = append([]byte{'<', '/'}, content[loc[0]+1:loc[1]]...)
			return loc[0]
		}
	}
	return beginIndex
}

/*
If no end index found, function returns -1.
*/
func (f Formatter) FindEndIndex(content []byte) int {
	pattern, err := regexp.Compile("<[a-zA-Z/]+")
	if err == nil {
		matches := pattern.FindAllIndex(content, -1)
		if matches != nil {
			tagCounter := 0
			for _, match := range matches {
				fmt.Println(string(content[match[0]:match[1]]))
				if bytes.Equal(content[match[0]:match[1]], f.FirstDOMElement) {
					tagCounter++
				} else if bytes.Equal(content[match[0]:match[1]], f.LastDOMElement) {
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
Not implemented.
*/
func (f Formatter) Format(content []byte) []byte {
	return []byte(xmlfmt.FormatXML(string(content), "\t", "  "))
}
