package pageObject

import (
	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	
	"bytes"
	"fmt"
)

type (
	PageObjectDefInterface interface {
		getDefinitionFromTree(tree compiler.ExecutionTestTree) string
	}
	PageObjectDefParser struct{}
)

func (this *PageObjectDefParser) GetDefinitionFromTree(tree compiler.ExecutionTestTree) string {
	var buffer bytes.Buffer
	page := definePage(tree.Feature)
	buffer.WriteString(page)
	return buffer.String()
}

func definePage(feature lexer.Feature) string {
	var buffer bytes.Buffer
	
	page := fmt.Sprintf("PAGE \"%s\"\n",feature.PageName)
	uri := fmt.Sprintf("\tURI \"%s\"\n","<ENTER PAGE URI>")
	
	buffer.WriteString(page)
	buffer.WriteString(uri)
	
	for _,scn := range feature.Scenarios {
		for _,action := range scn.Actions {
			element := defineActionElement(action)
			buffer.WriteString(element)
		}
	} 
	return buffer.String()
}

func defineActionElement(action lexer.Expect_action) string {
	var buffer bytes.Buffer
	
	element := fmt.Sprintf("\tELEMENT \"%s\"\n",action.ObjectId)
	locator := fmt.Sprintf("\t\tLOCATOR \"%s\"\n","<ENTER ELEMENT LOCATOR>")
	elType := fmt.Sprintf("\t\tTYPE \"%s\"\n",action.ObjectType)	
	
	buffer.WriteString(element)
	buffer.WriteString(locator)
	buffer.WriteString(elType)
	
	return buffer.String()
}
