package pageObject

import "github.com/ONSBR/bddTest.Go/compiler/lexer"

type (
	action func(p *PageElement, a *lexer.Expect_action) error
)

var (
	actions = map[string]action{
		"click": clickAction,
		"set":   sendText,
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

// sendText sends text to a web element.
func sendText(p *PageElement, a *lexer.Expect_action) error {
	var err error

	if err = p.Find(); err != nil {
		return err
	}

	if err = p.element.Clear(); err != nil {
		return err
	}

	return p.element.SendKeys(a.Param)
}
