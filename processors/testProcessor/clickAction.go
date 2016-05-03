package testProcessor

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
)

type ClickAction struct {
	Action
}

func (c *ClickAction) GetName() string {
	return "click"
}

func (c *ClickAction) Execute(lexer.Assertion) bool {
	return true
}

func NewClickAction() Action {
	return &ClickAction{}
}
