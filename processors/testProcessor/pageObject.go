package testProcessor

import (
	"fmt"
	"sourcegraph.com/sourcegraph/go-selenium"
)

/*

 */
type PageObject struct {
	driver   selenium.WebDriver
	Uri      string
	Elements []PageElement
}

/*
Creates a new instance of PageObject
*/
func NewPageObject(uri string) *PageObject {
	p := &PageObject{Uri: uri}
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "phantomjs"})

	var err error

	if p.driver, err = selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
	}

	return p
}

/*
Navigate to the web page object.
*/
func (p *PageObject) Open() {
	var err error

	if err = p.driver.Get(p.Uri); err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
	}
}

/*

 */
type PageElement struct {
	Page        PageObject
	Locator     string
	ElementType string
	ElementId   string
}

/*
Creates a new instance of PageElement
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
