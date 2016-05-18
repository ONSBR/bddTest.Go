package pageObject

import (
	"errors"
	"fmt"
	"os"

	"sourcegraph.com/sourcegraph/go-selenium"
)

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
func (p *PageObject) Open() error {
	var err error

	if err = p.driver.Get(p.Uri); err != nil {
		return err
	}

	// if _, err := p.driver.Status(); err != nil {
	// 	return err
	// }

	return nil
}

/*
FindPageElement finds an element by name.
*/
func (p *PageObject) FindPageElement(objectID string) (*PageElement, error) {
	for _, el := range p.Elements {
		if el.ElementId == objectID {
			return &el, nil
		}
	}

	return nil, errors.New("Element could not be found.")
}

/*
SaveScreenShot saves a screenshot.
*/
func (p *PageObject) SaveScreenShot(filename string) {
	screenshot, err := p.driver.Screenshot()

	if err != nil {
		panic(err)
	}

	fo, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := fo.Write(screenshot); err != nil {
		panic(err)
	}
}
