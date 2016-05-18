package pageObject

import (
	"testing"

	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/stretchr/testify/assert"
)

//
func TestShouldClickOnElement(t *testing.T) {
	// mock
	pageObject := NewPageObject("pageName", serverUrl)
	pageElement := NewPageElement(pageObject, "id", "buttonId", "button", "teste")
	pageObject.Open()

	// act
	err := clickAction(pageElement, nil)

	// assert
	assert.Nil(t, err)

	// cleanup
	pageObject.driver.Quit()
}

//
func TestShouldSendTextToElement(t *testing.T) {
	// mock
	pageObject := NewPageObject("pageName", serverUrl)
	pageElement := NewPageElement(pageObject, "id", "txtId", "textInput", "teste")
	action := &lexer.Expect_action{Param: "valor"}

	pageObject.Open()

	// act
	err := sendText(pageElement, action)

	// assert
	assert.Nil(t, err)

	// cleanup
	pageObject.driver.Quit()
}
