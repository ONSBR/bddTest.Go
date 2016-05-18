package processor

import (
	"testing"

	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/stretchr/testify/assert"
)

var mockFeature = lexer.Feature{
	LineNum:  1,
	FullText: "Aspecto Primeiro aspecto",
	Label:    "Aspecto",
	Name:     "Primeiro aspecto",
	PageName: "Home",
	Scenarios: []lexer.Scenario{
		lexer.Scenario{
			LineNum:  3,
			FullText: "Cenario Cenario1",
			Label:    "Cenario",
			Name:     "Cenario1",
			Username: "clovis.chedid",
			Actions: []lexer.Expect_action{
				lexer.Expect_action{
					LineNum:    5,
					FullText:   "Quando eu clico no botao teste",
					Label:      "Quando",
					Action:     "click",
					ObjectType: "botao",
					ObjectId:   "objectId",
					Param:      "clovis1",
				},
			},
			Expectations: []lexer.Expect_expression{
				lexer.Expect_expression{
					LineNum:    7,
					FullText:   "Entao eu espero a lista teste2 com a opcao clovis3",
					Label:      "Entao",
					Action:     "espero",
					ObjectType: "lista",
					ObjectId:   "teste2",
					Param:      "clovis3",
				},
			},
		},
	},
}

var mockPageObject = pageObject.NewPageObject("pageName", "http://172.17.0.1:8080")

var mockPageElement = pageObject.NewPageElement(mockPageObject, "id", "objectId", "button", "teste")

var mockedExecution = compiler.Execution{
	Filename:   "teste.spec",
	PageObject: *mockPageObject,
	Feature:    mockFeature,
	HasError:   false,
	Error:      "",
}

var mockedAction = lexer.Expect_action{
	LineNum:    5,
	FullText:   "Quando eu clico no botao teste",
	Label:      "Quando",
	Action:     "click",
	ObjectType: "button",
	ObjectId:   "objectId",
	Param:      "clovis1",
}

func TestShouldCreateNewTestProcessor(t *testing.T) {
	// act
	processor := NewTestProcessor(&mockedExecution)

	// assert
	assert.NotNil(t, processor)
	assert.NotNil(t, processor.Execution)
}

func TestShouldProcessFeature(t *testing.T) {
	processor := NewTestProcessor(&mockedExecution)

	// act
	err := processor.Process()

	// assert
	assert.NoError(t, err)
}
