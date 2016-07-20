package actions

import (
	"testing"

	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/stretchr/testify/assert"
)

func TestClickOnElement(t *testing.T) {
	// mock
	page := pageObject.NewPageObject("pageName", "http://172.17.0.1:8080")
	element := pageObject.NewPageElement(page, "id", "buttonId", "button", "teste")

	page.Open()
	defer page.Driver.Quit()

	// act
	click := ClickAction{}
	err := click.Execute(element, nil)

	// assert
	assert.Nil(t, err)
}
