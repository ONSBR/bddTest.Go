package actions

import (
	"testing"

	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/stretchr/testify/assert"
)

//
func TestSendTextToElement(t *testing.T) {
	// mock
	page := pageObject.NewPageObject("pageName", "http://172.17.0.1:8080")
	element := pageObject.NewPageElement(page, "id", "txtId", "textInput", "teste")
	action := &lexer.Expect_action{Param: "valor"}

	page.Open()
	defer page.Driver.Quit()

	// act
	sendText := SendTextAction{}
	sendText.Execute(element, action)

	// assert
	el, _ := page.Driver.FindElement("id", "txtId")
	value, _ := el.GetAttribute("value")
	assert.Equal(t, value, action.Param)
}
