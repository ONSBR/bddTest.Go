package testProcessor

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
)

type ActionContext struct {
	actions map[string]Action
}

type Action interface {
	GetName() string
	Execute(assertion lexer.Assertion) bool
}

func (a *ActionContext) GetAction(n string) Action {
	return a.actions[n]
}

func NewActionContext() ActionContext {
	a := ActionContext{}
	a.actions = make(map[string]Action)
	a.actions["click"] = NewClickAction()
	return a
}
