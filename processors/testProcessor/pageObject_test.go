package testProcessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*

 */
func TestShouldCreateNewPageObject(t *testing.T) {
	uri := "http://172.17.0.1:8080"
	pageObject := NewPageObject(uri)

	assert.NotNil(t, pageObject)
	assert.NotNil(t, pageObject.driver)
	assert.Equal(t, pageObject.Uri, uri)

	pageObject.driver.Quit()
}

/*

 */
func TestShouldOpenWebPage(t *testing.T) {
	uri := "http://172.17.0.1:8080"
	pageObject := NewPageObject(uri)

	pageObject.Open()
	pageSource, _ := pageObject.driver.PageSource()

	assert.Contains(t, pageSource, "Hello")
	pageObject.driver.Quit()
}

/*

 */
func TestShouldCreateNewPageElement(t *testing.T) {
	pageObject := &PageObject{}
	pageElement := NewPageElement(pageObject, "css", "button", "elementId")

	assert.NotNil(t, pageElement)
	assert.Equal(t, pageElement.Locator, "css")
	assert.Equal(t, pageElement.ElementType, "button")
	assert.Equal(t, pageElement.ElementId, "elementId")
}

/*

 */
func TestShouldAppendElementToPageObject(t *testing.T) {
	pageObject := &PageObject{}
	NewPageElement(pageObject, "css", "button", "elementId")

	assert.Len(t, pageObject.Elements, 1)
}

// func TestShouldFindAnWebElement(t *testing.T) {
// 	uri := "http://172.17.0.1:8080"
// 	pageObject := NewPageObject(uri)
// 	pageElement := NewPageElement(pageObject, "css", "button", "elementId")

// 	pageObject.Open()

// }
