package actions

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
)

// SendTextAction is a send text action.
type SendTextAction struct {
	Action
}

// Execute sends text to a web element.
func (s *SendTextAction) Execute(p *pageObject.PageElement, a *lexer.Expect_action) error {
	var err error

	if err = p.Find(); err != nil {
		return err
	}

	if err = p.Element.Clear(); err != nil {
		return err
	}

	return p.Element.SendKeys(a.Param)
}
