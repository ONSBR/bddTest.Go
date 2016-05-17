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
	pageObject := NewPageObject("pageName", serverUrl)

	// assert
	assert.NotNil(t, pageObject)
	assert.NotNil(t, pageObject.driver)
	assert.Equal(t, pageObject.Uri, serverUrl)

	// cleanup
	pageObject.driver.Quit()
}

/*

 */
func TestShouldOpenWebPage(t *testing.T) {
	// mock
	pageObject := NewPageObject("pageName", serverUrl)

	// act
	pageObject.Open()
	defer pageObject.driver.Quit()
	pageSource, _ := pageObject.driver.PageSource()

	// assert
	assert.True(t, strings.Contains(pageSource, "HTML Content"))
}
