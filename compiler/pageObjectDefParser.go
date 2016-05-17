package compiler

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/ONSBR/bddTest.Go/util"

	yaml "gopkg.in/yaml.v2"

	//	"bytes"
	"fmt"
	"strings"
)

var logPDef = util.GetLogger("pageObject.defParser")

var elements map[string]string

type (
	ParseError struct {
		Errors []string
	}

	YamlElement struct {
		Element string
		Locator string
		Expression string
		Type    string
	}

	YamlPage struct {
		Page     string
		Uri      string
		Elements []YamlElement
	}

	iPageObjectDef interface {
		GetDefinitionFromTree(tree ExecutionTestTree) string
		GetYamlPage(definition string) (page YamlPage, err error)
		GetPageObject(definition string, baseUri string) (page *pageObject.PageObject, err error)
	}
	PageObjectDefParser struct{}
)

func (e ParseError) Error() string {
	return strings.Join(e.Errors, "\n")
}

func (this *PageObjectDefParser) GetDefinitionFromTree(tree ExecutionTestTree) (string, error) {
	elements = map[string]string{}
	page := YamlPage{}
	err := definePage(&page, tree.Feature)
	if err != nil {
		return "", err
	}
	err = defineActionElement(&page, tree.Feature)
	if err != nil {
		return "", err
	}
	err = defineExpectationElement(&page, tree.Feature)
	if err != nil {
		return "", err
	}
	return serializePage(&page)
}

func serializePage(page *YamlPage) (string, error) {
	serialized := ""
	byteArr, err := yaml.Marshal(&page)
	if err == nil {
		serialized = string(byteArr[:len(byteArr)])
		logPDef.Infof("%s", serialized)
		//		fmt.Printf("|%s|",serialized)
	}
	return serialized, err
}

func definePage(page *YamlPage, feature lexer.Feature) error {
	page.Page = feature.PageName
	page.Uri = "<ENTER PAGE URI>"
	return nil
}

func defineActionElement(page *YamlPage, feature lexer.Feature) error {
	for _, scn := range feature.Scenarios {
		for _, action := range scn.Actions {
			if t, ok := elements[action.ObjectId]; ok && t != action.ObjectType {
				return ParseError{Errors: []string{fmt.Sprintf("Duplicated element %s with different types %s || %s", action.ObjectId, action.ObjectType, t)}}
			}
			if _, ok := elements[action.ObjectId]; !ok {
				element := YamlElement{
					Element: action.ObjectId,
					Locator: "<ENTER ELEMENT LOCATOR>",
					Expression: "<ENTER ELEMENT LOCATOR EXPRESSION>",
					Type:    action.ObjectType,
				}
				elements[action.ObjectId] = action.ObjectType
				page.Elements = append(page.Elements, element)
			}
		}
	}
	return nil
}

func defineExpectationElement(page *YamlPage, feature lexer.Feature) error {
	for _, scn := range feature.Scenarios {
		for _, expect := range scn.Expectations {
			if t, ok := elements[expect.ObjectId]; ok && t != expect.ObjectType {
				return ParseError{Errors: []string{fmt.Sprintf("Duplicated element %s with different types %s || %s", expect.ObjectId, expect.ObjectType, t)}}
			}
			if _, ok := elements[expect.ObjectId]; !ok {
				element := YamlElement{
					Element: expect.ObjectId,
					Locator: "<ENTER ELEMENT LOCATOR>",
					Expression: "<ENTER ELEMENT LOCATOR EXPRESSION>",
					Type:    expect.ObjectType,
				}
				elements[expect.ObjectId] = expect.ObjectType
				page.Elements = append(page.Elements, element)
			}
		}
	}
	return nil
}

func (this *PageObjectDefParser) GetYamlPage(definition string) (page YamlPage, err error) {
	elements = map[string]string{}
	page = YamlPage{}
	errYaml := yaml.Unmarshal([]byte(definition), &page)
	if errYaml != nil {
		err = ParseError{Errors: []string{errYaml.Error()}}
		return
	}
	if page.Page == "" {
		err = ParseError{Errors: []string{fmt.Sprintf("Missing page name")}}
		return
	}
	if page.Uri == "<ENTER PAGE URI>" || page.Uri == "" {
		err = ParseError{Errors: []string{fmt.Sprintf("Missing page uri")}}
		return
	}
	for _, element := range page.Elements {
		if t, ok := elements[element.Element]; ok && t != element.Type {
			err = ParseError{Errors: []string{fmt.Sprintf("Duplicated element %s with different types %s || %s", element.Element, element.Type, t)}}
			return
		}
		if element.Element == "" {
			err = ParseError{Errors: []string{fmt.Sprintf("Missing element name for locator %s and type %s", element.Locator, element.Type)}}
			return
		}
		if element.Locator == "<ENTER ELEMENT LOCATOR>" || element.Locator == "" {
			err = ParseError{Errors: []string{fmt.Sprintf("Missing element locator for name %s and type %s", element.Element, element.Type)}}
			return
		}
		if element.Expression == "<ENTER ELEMENT LOCATOR EXPRESSION>" || element.Expression == "" {
			err = ParseError{Errors: []string{fmt.Sprintf("Missing element locator expression for name %s and type %s", element.Element, element.Type)}}
			return
		}
		if element.Type == "" {
			err = ParseError{Errors: []string{fmt.Sprintf("Missing element type for locator %s and name %s", element.Locator, element.Element)}}
			return
		}
		elements[element.Element] = element.Type
	}
	return
}

func (this *PageObjectDefParser) GetPageObject(definition string, baseUri string) (page *pageObject.PageObject, err error) {
	yamlPage, yamlErr := this.GetYamlPage(definition)
	if yamlErr != nil {
		err = yamlErr
		return
	}

	uri := baseUri

	if baseUri[len(baseUri)-1:] != "/" {
		uri += "/"
	}
	if yamlPage.Uri[0:1] == "/" {
		uri += yamlPage.Uri[1:]
	} else {
		uri += yamlPage.Uri
	}

	page = pageObject.NewPageObject(yamlPage.Page, uri)
	for _, element := range yamlPage.Elements {
		_ = pageObject.NewPageElement(page, element.Locator, element.Expression, element.Type, element.Element)
	}
	return
}
