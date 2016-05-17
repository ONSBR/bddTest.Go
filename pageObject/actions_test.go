package pageObject

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
func TestShouldClickOnElement(t *testing.T) {
	// mock
	pageObject := NewPageObject("pageName", serverUrl)
	pageElement := NewPageElement(pageObject, "id", "objectId", "button", "objectId")
	pageObject.Open()

	// act
	err := clickAction(pageElement, nil)

	// assert
	assert.Nil(t, err)

	// cleanup
	pageObject.driver.Quit()
}

// func TestShouldExecuteAnAc()  {
//
// }
