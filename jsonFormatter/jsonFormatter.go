package jsonFormatter

import (
	"bytes"
	"encoding/json"
)

var (
	beginSymbol = []byte{'{', '['}
	endSymbol   = []byte{'}', ']'}
)

type Formatter struct {
}

/*
There is two symbols, which could be the beginning of JSON string,
we need to check both and return closest.
If no begin index found, function returns -1.
*/
func (p Formatter) FindBeginIndex(content []byte) int {
	beginIndex := bytes.IndexByte(content, beginSymbol[0])
	tempIndex := bytes.IndexByte(content, beginSymbol[1])
	if beginIndex > tempIndex && tempIndex != -1 {
		beginIndex = tempIndex
	}
	return beginIndex
}

/*
To find the end of JSON string we need to count all closing and openings symbols starting from first.
So the first symbol has to be one of the beginning symbols.
 */
func (p Formatter) FindEndIndex(content []byte) int {
	signIndex := bytes.IndexByte(beginSymbol, content[0])
	if signIndex != -1 {
		beginsCount := 0
		for i, symbol := range content {
			switch symbol {
			case beginSymbol[signIndex]:
				beginsCount++
			case endSymbol[signIndex]:
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

/*
Simple JSON formatting.
If JSON string is corrupted, function returns empty array.
*/
func (p Formatter) Format(content []byte) []byte {
	var data bytes.Buffer
	json.Indent(&data, content, "", "  ")
	return data.Bytes()
}
