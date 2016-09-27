package readers

import (
	"github.com/ONSBR/bddTest.Go/pageObject"
	"sourcegraph.com/sourcegraph/go-selenium"
)

var (
	readers = map[string]ValueReader{
		"textbox": &TextInputReader{},
	}
)

// ValueReader ...
type ValueReader interface {
	Read(element selenium.WebElement) (string, error)
}

// Read reads element value
func Read(p *pageObject.PageElement) (string, error) {
	reader := readers[p.ElementType]

	if p.Element == nil {
		if err := p.Find(); err != nil {
			return "", err
		}
	}

	return reader.Read(p.Element)
}
