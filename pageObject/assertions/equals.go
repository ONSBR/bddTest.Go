package assertions

import (
	"fmt"

	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/ONSBR/bddTest.Go/pageObject/readers"
)

// EqualsAssertion asserts that two given values are equals
type EqualsAssertion struct {
	Assertion
}

// Assert ..
func (e *EqualsAssertion) Assert(p *pageObject.PageElement, expected string) error {
	var value string
	var err error

	if value, err = readers.Read(p); err != nil {
		return err
	}

	if value == expected {
		return nil
	}

	return fmt.Errorf("Expected '%s', but got: '%s'", expected, value) // errors.New(" ")
}
