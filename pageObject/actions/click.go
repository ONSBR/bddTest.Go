package actions

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
)

// ClickAction represents a click on an element.
type ClickAction struct {
	Action
}

// Execute clicks on a web element using selenium webdriver.
func (c *ClickAction) Execute(p *pageObject.PageElement, a *lexer.Expect_action) error {
	if err := p.Find(); err != nil {
		return err
	}

	return p.Element.Click()
}
