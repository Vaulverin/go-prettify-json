package jsonParser

import (
	"bytes"
	"encoding/json"
)

var (
	beginSymbol = []byte{'{', '['}
	endSymbol   = []byte{'}', ']'}
)

type Parser struct {
}

func (p Parser) FindBeginIndex(content []byte) int {
	beginIndex := bytes.IndexByte(content, beginSymbol[0])
	tempIndex := bytes.IndexByte(content, beginSymbol[1])
	if beginIndex > tempIndex && tempIndex != -1 {
		beginIndex = tempIndex
	}
	return beginIndex
}

func (p Parser) FindEndIndex(content []byte) int {
	signIndex := bytes.IndexByte(beginSymbol, content[0])
	if signIndex != -1 {
		beginsCount := 0
		for i, symbol := range content {
			switch symbol {
			case beginSymbol[ signIndex ]:
				beginsCount++
			case endSymbol[ signIndex ]:
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

func (p Parser) Parse(content []byte) []byte {
	var data bytes.Buffer
	json.Indent(&data, content, "", "  ")
	return data.Bytes()
}

