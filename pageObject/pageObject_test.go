package pageObject

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
func TestShouldCreateNewPageObject(t *testing.T) {
	// mock
	pageObject := NewPageObject("pageName", serverURL)

	// assert
	assert.NotNil(t, pageObject)
	assert.NotNil(t, pageObject.Driver)
	assert.Equal(t, pageObject.URI, serverURL)

	// cleanup
	pageObject.Driver.Quit()
}

//
func TestShouldOpenWebPage(t *testing.T) {
	// mock
	fmt.Println(serverURL)
	pageObject := NewPageObject("pageName", serverURL)

	// act
	pageObject.Open()
	defer pageObject.Driver.Quit()
	pageSource, _ := pageObject.Driver.PageSource()

	// assert
	assert.True(t, strings.Contains(pageSource, "HTML Content"))
}
