package testProcessor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldInstanceNewActionContext(t *testing.T) {
	// mock
	context := NewActionContext()

	// assert
	assert.NotNil(t, context.actions)
	assert.NotEmpty(t, context.actions)
}

func TestActionContextShouldReturnCorrectActionRunner(t *testing.T) {
	// mock
	context := NewActionContext()
	action := context.GetAction("click")

	// assert
	assert.NotNil(t, context)
	assert.Equal(t, action.GetName(), "click")
}
