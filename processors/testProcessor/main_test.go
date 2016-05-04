package testProcessor

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	// "github.com/stretchr/testify/assert"
	"testing"
)

/*

 */
func TestShouldFindPageObjectFromAssertion(t *testing.T) {
	// mock
	assertions := []lexer.Assertion{{
		LineNum:    0,
		FullText:   "Quando clicar no botão elementId",
		Label:      "",
		Action:     "clicar",
		ObjectType: "Button",
		ObjectId:   "elementId",
		Param:      "",
	}}

	Process(assertions)
}
