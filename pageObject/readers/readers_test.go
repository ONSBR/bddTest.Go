package readers

import (
	"testing"

	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/stretchr/testify/assert"
)

// TestReadTextInputValue
func TestReadTextInputValue(t *testing.T) {
	// mock
	page := pageObject.NewPageObject("pageName", "http://172.17.0.1:8080")

	page.Open()
	defer page.Driver.Quit()

	element, _ := page.Driver.FindElement("id", "txtId")

	// // act
	reader := TextInputReader{}
	actual, err := reader.Read(element)
	//
	// // assert
	assert.Nil(t, err)
	assert.Equal(t, "value", actual)
}
