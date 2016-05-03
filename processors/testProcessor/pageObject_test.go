package testProcessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*

 */
func TestShouldCreateNewPageObject(t *testing.T) {
	// mock
	uri := "http://172.17.0.1:8080"
	pageObject := NewPageObject(uri)

	// assert
	assert.NotNil(t, pageObject)
	assert.NotNil(t, pageObject.driver)
	assert.Equal(t, pageObject.Uri, uri)

	// cleanup
	pageObject.driver.Quit()
}

/*

 */
func TestShouldOpenWebPage(t *testing.T) {
	// mock
	uri := "http://172.17.0.1:8080"
	pageObject := NewPageObject(uri)

	// act
	pageObject.Open()
	pageSource, _ := pageObject.driver.PageSource()

	// assert
	assert.Contains(t, pageSource, "Hello")

	// cleanup
	pageObject.driver.Quit()
}
