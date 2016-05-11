package pageObject

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/*

 */
func TestShouldCreateNewPageObject(t *testing.T) {
	// mock
	uri := "http://172.17.0.1:8080"
	pageObject := NewPageObject("pageName", uri)

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
	initTestServer()

	pageObject := NewPageObject("pageName", "http://172.17.0.1:8080")

	// act
	pageObject.Open()
	defer pageObject.driver.Quit()
	pageSource, _ := pageObject.driver.PageSource()

	// assert
	assert.True(t, strings.Contains(pageSource, "HTML Content"))
}
