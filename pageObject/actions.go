package pageObject

import "github.com/ONSBR/bddTest.Go/compiler/lexer"

type (
	action func(p *PageElement, a *lexer.Expect_action) error
)

var (
	actions = map[string]action{
		"click": clickAction,
	}
)

// Execute fires an action.
func Execute(p *PageElement, a *lexer.Expect_action) error {
	return actions[a.Action](p, a)
}

// clickAction clicks on a web element.
func clickAction(p *PageElement, a *lexer.Expect_action) error {
	if err := p.Find(); err != nil {
		return err
	}

	return p.element.Click()
}
