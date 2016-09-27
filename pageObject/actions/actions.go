package actions

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
)

var (
	actions = map[string]Action{
		"click": &ClickAction{},
		"set":   &SendTextAction{},
	}
)

// Action is an action which can be executed against selenium webdriver
type Action interface {
	Execute(p *pageObject.PageElement, a *lexer.Expect_action) error
}

// Execute fires an action.
func Execute(p *pageObject.PageElement, a *lexer.Expect_action) error {
	return actions[a.Action].Execute(p, a)
}
