package htmlFormatter

import (
	"log-formatter/domFormatter"
	"github.com/yosssi/gohtml"
)

type Formatter struct {
	domFormatter.Formatter
}

func (f Formatter) Format(content []byte) []byte {
	return gohtml.FormatBytes(content)
}