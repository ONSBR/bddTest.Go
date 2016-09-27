package assertions

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
)

var (
	assertions = map[string]Assertion{
		"eq": &EqualsAssertion{},
	}
)

// Assertion ...
type Assertion interface {
	Assert(p *pageObject.PageElement, e string) error
}

// Assert evaluates an expectation
func Assert(p *pageObject.PageElement, e *lexer.Expect_expression) error {
	assertion := assertions[e.Matcher]
	return assertion.Assert(p, e.Param)
}
