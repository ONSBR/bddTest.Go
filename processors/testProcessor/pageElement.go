package testProcessor

import (
	"fmt"
	"sourcegraph.com/sourcegraph/go-selenium"
)

/*

 */
type PageElement struct {
	Page        PageObject
	Locator     string
	ElementType string
	ElementId   string
	element     selenium.WebElement
}

/*
Find finds the web element in the Page
*/
func (p *PageElement) Find() {
	var err error
	if p.element, err = p.Page.driver.FindElement(selenium.ById, p.ElementId); err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
	}
}

/*
NewPageElement creates a new instance of PageElement
*/
func NewPageElement(page *PageObject, locator string, elementType string, elementId string) *PageElement {
	pageElement := &PageElement{
		Page:        *page,
		Locator:     locator,
		ElementType: elementType,
		ElementId:   elementId,
	}

	page.Elements = append(page.Elements, *pageElement)
	return pageElement
}
