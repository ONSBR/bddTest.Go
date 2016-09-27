package pageObject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewPageElement(t *testing.T) {
	// mock
	pageObject := &PageObject{}

	// act
	pageElement := NewPageElement(pageObject, "id", "objectId", "button", "teste")

	// assert
	assert.NotNil(t, pageElement)
	assert.Equal(t, pageElement.Locator, "id")
	assert.Equal(t, pageElement.Expression, "objectId")
	assert.Equal(t, pageElement.ElementType, "button")
	assert.Equal(t, pageElement.ElementId, "teste")
}

func TestShouldAppendElementToPageObject(t *testing.T) {
	// mock
	pageObject := &PageObject{}

	// act
	NewPageElement(pageObject, "id", "objectId", "button", "elementId")

	// assert
	assert.Len(t, pageObject.Elements, 1)
}

func TestShouldFindPageElement(t *testing.T) {
	pageObject := NewPageObject("pageName", serverURL)
	pageElement := NewPageElement(pageObject, "id", "buttonId", "button", "objectId")

	// act
	pageObject.Open()
	pageElement.Find()

	// assert
	assert.NotNil(t, pageElement.Element)

	// cleanup
	pageObject.Driver.Quit()
}

func TestShouldFindPageElementFromObjectId(t *testing.T) {
	// mock
	pageObject := NewPageObject("pageName", serverURL)
	pageElement := NewPageElement(pageObject, "id", "objectId", "button", "objectId")

	// act
	foundElement, err := pageObject.FindPageElement("objectId")

	// assert
	assert.Equal(t, pageElement, foundElement)
	assert.Nil(t, err)

	// cleanup
	pageObject.Driver.Quit()
}

func TestShouldReturnErrorForNonExistingObjectId(t *testing.T) {
	// mock
	pageObject := NewPageObject("pageName", serverURL)
	NewPageElement(pageObject, "id", "objectId", "button", "objectId")

	// act
	foundElement, err := pageObject.FindPageElement("wrong_objectId")

	// assert
	assert.Nil(t, foundElement)
	assert.NotNil(t, err)

	// cleanup
	pageObject.Driver.Quit()
}
