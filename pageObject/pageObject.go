package pageObject

import (
	"errors"
	"fmt"
	"sourcegraph.com/sourcegraph/go-selenium"
)

/*

 */
type PageObject struct {
	Page     string
	driver   selenium.WebDriver
	Uri      string
	Elements []PageElement
}

/*
NewPageObject creates a new instance of PageObject
*/
func NewPageObject(page string, uri string) *PageObject {
	p := &PageObject{Page: page, Uri: uri}
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "phantomjs"})

	var err error

	if p.driver, err = selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
	}

	return p
}

/*
Open navigates to the web page object.
*/
func (p *PageObject) Open() {
	var err error

	if err = p.driver.Get(p.Uri); err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
	}
}

/*
FindPageElement finds an element by name.
*/
func (p *PageObject) FindPageElement(objectId string) (*PageElement, error) {
	for _, el := range p.Elements {
		if el.ElementId == objectId {
			return &el, nil
		}
	}

	return nil, errors.New("Element could not be found.")
}
