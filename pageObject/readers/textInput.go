package readers

import "sourcegraph.com/sourcegraph/go-selenium"

//TextInputReader ...
type TextInputReader struct {
	ValueReader
}

// Read ...
func (t *TextInputReader) Read(element selenium.WebElement) (string, error) {
	return element.GetAttribute("value")
}
