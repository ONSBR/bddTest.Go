package pageObject

import (
	"fmt"

	"sourcegraph.com/sourcegraph/go-selenium"
)

/*
PageElement is a PageObject web element.
*/
type PageElement struct {
	Page        PageObject
	Locator     string
	Expression  string
	ElementType string
	ElementId   string
	element     selenium.WebElement
}

/*
Find finds the web element in the Pagep
*/
func (p *PageElement) Find() error {
	var err error

	if p.element, err = p.Page.driver.FindElement(p.Locator, p.Expression); err != nil {
		return fmt.Errorf("Failed to find element \"%s\" using %s=%s", p.ElementId, p.Locator, p.Expression)
	}

	return nil
}

/*
NewPageElement creates a new instance of PageElement
*/
func NewPageElement(page *PageObject, locator string, expression string, elementType string, elementID string) *PageElement {
	pageElement := &PageElement{
		Page:        *page,
		Locator:     locator,
		Expression:  expression,
		ElementType: elementType,
		ElementId:   elementID,
	}

	page.Elements = append(page.Elements, *pageElement)
	return pageElement
}
