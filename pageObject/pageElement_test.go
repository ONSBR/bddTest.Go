package pageObject

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestShouldCreateNewPageElement(t *testing.T) {
	// mock
	pageObject := &PageObject{}

	// act
	pageElement := NewPageElement(pageObject, "css", "button", "elementId")

	// assert
	assert.NotNil(t, pageElement)
	assert.Equal(t, pageElement.Locator, "css")
	assert.Equal(t, pageElement.ElementType, "button")
	assert.Equal(t, pageElement.ElementId, "elementId")
}

func TestShouldAppendElementToPageObject(t *testing.T) {
	// mock
	pageObject := &PageObject{}

	// act
	NewPageElement(pageObject, "css", "button", "elementId")

	// assert
	assert.Len(t, pageObject.Elements, 1)
}

// func TestShouldFindPageElement(t *testing.T) {
// 	initTestServer()

// 	pageObject := NewPageObject("http://172.17.0.1:8080")
// 	pageElement := NewPageElement(pageObject, "css", "button", "elementId")

// 	// act
// 	pageObject.Open()
// 	pageElement.Find()

// 	// assert
// 	assert.NotNil(t, pageElement.element)

// 	// cleanup
// 	pageObject.driver.Quit()
// }
