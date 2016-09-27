package assertions

import (
	"testing"

	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/stretchr/testify/assert"
)

func TestAssertTwoValuesAreEquals(t *testing.T) {
	// mock
	page := pageObject.NewPageObject("pageName", "http://172.17.0.1:8080")
	element := pageObject.NewPageElement(page, "id", "txtId", "textInput", "teste")

	page.Open()
	defer page.Driver.Quit()

	// act
	assertion := EqualsAssertion{}

	err := assertion.Assert(element, "value")

	// assert
	assert.Nil(t, err)
}

func TestAssertTwoValuesAreEqualsCanFail(t *testing.T) {
	// mock
	page := pageObject.NewPageObject("pageName", "http://172.17.0.1:8080")
	element := pageObject.NewPageElement(page, "id", "txtId", "textInput", "teste")

	page.Open()
	defer page.Driver.Quit()

	// act
	assertion := EqualsAssertion{}

	err := assertion.Assert(element, "different value")

	// assert
	assert.NotNil(t, err)

}
