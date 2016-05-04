package pageObject_test

import (
	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/util"
	. "github.com/ONSBR/bddTest.Go/pageObject"
	

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var logCC = util.GetLogger("test.compiler")

var _ = Describe("PageObjectDefParser", func() {
	var (
		parser = &PageObjectDefParser{}
		
		goodTree1 compiler.ExecutionTestTree
		goodTree2 compiler.ExecutionTestTree
		goodTree3 compiler.ExecutionTestTree
		goodTree4 compiler.ExecutionTestTree
		definition1 string
		definition2 string
		
	)
	BeforeEach(func(){
			goodTree1 = compiler.ExecutionTestTree{
					NumScenarios:1,
					Feature: lexer.Feature{
						LineNum:0,
						FullText:"Aspecto: Primeiro aspecto",
						Label:"Aspecto",
						Name: "Primeiro aspecto",
						PageName:"Home",
						Scenarios:[]lexer.Scenario{
							lexer.Scenario{
								LineNum:0,
								FullText:"Cenario: Cenario1",
								Label:"Cenario",
								Name:"Cenario1",
								Username:"clovis.chedid",
								Actions:[]lexer.Expect_action{
									lexer.Expect_action{
										LineNum: 0,
										FullText: "Quando eu clico no botao teste com o valor \"clovis1\"",
										Label:"Quando",
										Action:"clico",
										ObjectType:"botao",
										ObjectId:"teste",
										Param:"clovis1",
									},
									lexer.Expect_action{
										LineNum: 0,
										FullText: "Quando eu preencho no campo teste com o valor \"clovis1\"",
										Label:"Quando",
										Action:"preencho",
										ObjectType:"campo",
										ObjectId:"teste",
										Param:"clovis1",
									},
								},
								Expectations:[]lexer.Expect_expression{
									lexer.Expect_expression{
										LineNum: 0,
										FullText: "Entao eu espero no botao teste com o valor \"clovis1\"",
										Label:"Entao",
										Action:"espero",
										ObjectType:"botao",
										ObjectId:"teste",
										Param:"clovis1",
									},
								},
							},
						},
					},
				}
			
			goodTree2 = compiler.ExecutionTestTree{
					NumScenarios:0,
					Feature: lexer.Feature{
						LineNum:0,
						FullText:"Aspecto: Primeiro aspecto",
						Label:"Aspecto",
						Name: "Primeiro aspecto",
						PageName:"Home",
						Scenarios:[]lexer.Scenario{},
					},
				}
			goodTree3 = compiler.ExecutionTestTree{
					NumScenarios:0,
					Feature: lexer.Feature{
						LineNum:0,
						FullText:"Aspecto: Primeiro aspecto",
						Label:"Aspecto",
						Name: "Primeiro aspecto",
						PageName:"Home",
						Scenarios:[]lexer.Scenario{
							lexer.Scenario{
								LineNum:0,
								FullText:"Cenario: Cenario1",
								Label:"Cenario",
								Name:"Cenario1",
								Username:"clovis.chedid",
								Actions:[]lexer.Expect_action{},
								Expectations:[]lexer.Expect_expression{},
							},
						},
					},
				}
			goodTree4 = compiler.ExecutionTestTree{
					NumScenarios:1,
					Feature: lexer.Feature{
						LineNum:0,
						FullText:"Aspecto: Primeiro aspecto",
						Label:"Aspecto",
						Name: "Primeiro aspecto",
						PageName:"Home",
						Scenarios:[]lexer.Scenario{
							lexer.Scenario{
								LineNum:0,
								FullText:"Cenario: Cenario1",
								Label:"Cenario",
								Name:"Cenario1",
								Username:"clovis.chedid",
								Actions:[]lexer.Expect_action{
									lexer.Expect_action{
										LineNum: 0,
										FullText: "Quando eu clico no botao teste com o valor \"clovis1\"",
										Label:"Quando",
										Action:"click",
										ObjectType:"button",
										ObjectId:"teste",
										Param:"clovis1",
									},
								},
								Expectations:[]lexer.Expect_expression{},
							},
						},
					},
				}
			definition1 = "PAGE \"Home\"\n\tURI \"<ENTER PAGE URI>\"\n"
			definition2 = "PAGE \"Home\"\n\tURI \"<ENTER PAGE URI>\"\n\tELEMENT \"teste\"\n\t\tLOCATOR \"<ENTER ELEMENT LOCATOR>\"\n\t\tTYPE \"button\"\n"
		})
	Describe("Generating Page Object Definitions", func(){
		Context("When valid incomplete trees", func(){
			It("Should return string with Page without objects for only Feature",func(done Done){
				definition := parser.GetDefinitionFromTree(goodTree2)
				Expect(definition).To(Equal(definition1))
				close(done)
			})
			It("Should return string with Page without objects for Scenario without actions",func(done Done){
				definition := parser.GetDefinitionFromTree(goodTree3)
				Expect(definition).To(Equal(definition1))
				close(done)
			})
		})
		Context("When valid complete trees", func(){
			It("Should return string with Page and one Element from action",func(done Done){
				definition := parser.GetDefinitionFromTree(goodTree4)
				Expect(definition).To(Equal(definition2))
				close(done)
			})
		})
	})
})
