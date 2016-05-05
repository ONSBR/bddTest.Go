package pageObject_test

import (
	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	. "github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/ONSBR/bddTest.Go/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var logCC = util.GetLogger("test.compiler")

var _ = Describe("PageObjectDefParser", func() {
	var (
		parser = &PageObjectDefParser{}

		goodTree1 compiler.ExecutionTestTree //Complete tree two actions two expects
		goodTree2 compiler.ExecutionTestTree //Just Feature tree
		goodTree3 compiler.ExecutionTestTree //Tree with only empty scenario
		goodTree4 compiler.ExecutionTestTree //Tree with one action
		goodTree5 compiler.ExecutionTestTree //Tree with two actions
		goodTree6 compiler.ExecutionTestTree //Tree with one expect
		goodTree7 compiler.ExecutionTestTree //Tree with two expects

		definition1 string //Just Feature tree
		definition2 string //Tree with one action
		definition3 string //Tree with two actions
		//		definition4 string
		//		definition5 string
		definition6 string //Tree with one expect
		definition7 string //Tree with two expects
		definition8 string //Tree with four elements

		definition1Compl string
		definition2Compl string
		definition3Compl string
		definition8Compl string

		badTree1 compiler.ExecutionTestTree //Tree with duplicate element action
		badTree2 compiler.ExecutionTestTree //Tree with duplicate element expectation
		badTree3 compiler.ExecutionTestTree //Tree with duplicate element action and expectation
	)
	BeforeEach(func() {
		goodTree1 = compiler.ExecutionTestTree{
			NumScenarios: 1,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "Home",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  0,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions: []lexer.Expect_action{
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu clico no botao teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "click",
								ObjectType: "button",
								ObjectId:   "teste1",
								Param:      "clovis1",
							},
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu preencho no campo teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "set",
								ObjectType: "textbox",
								ObjectId:   "teste2",
								Param:      "clovis1",
							},
						},
						Expectations: []lexer.Expect_expression{
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "button",
								ObjectId:   "teste3",
								Param:      "clovis1",
							},
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "textbox",
								ObjectId:   "teste4",
								Param:      "clovis1",
							},
						},
					},
				},
			},
		}

		goodTree2 = compiler.ExecutionTestTree{
			NumScenarios: 0,
			Feature: lexer.Feature{
				LineNum:   0,
				FullText:  "Aspecto: Primeiro aspecto",
				Label:     "Aspecto",
				Name:      "Primeiro aspecto",
				PageName:  "Home",
				Scenarios: []lexer.Scenario{},
			},
		}
		goodTree3 = compiler.ExecutionTestTree{
			NumScenarios: 0,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "Home",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:      0,
						FullText:     "Cenario: Cenario1",
						Label:        "Cenario",
						Name:         "Cenario1",
						Username:     "clovis.chedid",
						Actions:      []lexer.Expect_action{},
						Expectations: []lexer.Expect_expression{},
					},
				},
			},
		}
		goodTree4 = compiler.ExecutionTestTree{
			NumScenarios: 1,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "Home",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  0,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions: []lexer.Expect_action{
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu clico no botao teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "click",
								ObjectType: "button",
								ObjectId:   "teste",
								Param:      "clovis1",
							},
						},
						Expectations: []lexer.Expect_expression{},
					},
				},
			},
		}
		goodTree5 = compiler.ExecutionTestTree{
			NumScenarios: 1,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "Home",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  0,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions: []lexer.Expect_action{
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu clico no botao teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "click",
								ObjectType: "button",
								ObjectId:   "teste",
								Param:      "clovis1",
							},
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu preencho no campo teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "set",
								ObjectType: "textbox",
								ObjectId:   "teste1",
								Param:      "clovis1",
							},
						},
						Expectations: []lexer.Expect_expression{},
					},
				},
			},
		}
		goodTree6 = compiler.ExecutionTestTree{
			NumScenarios: 1,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "HomeExpect",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  0,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions:  []lexer.Expect_action{},
						Expectations: []lexer.Expect_expression{
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "button",
								ObjectId:   "teste1",
								Param:      "clovis1",
							},
						},
					},
				},
			},
		}
		goodTree7 = compiler.ExecutionTestTree{
			NumScenarios: 1,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "HomeExpect2",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  0,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions:  []lexer.Expect_action{},
						Expectations: []lexer.Expect_expression{
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "button",
								ObjectId:   "teste1",
								Param:      "clovis1",
							},
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "textbox",
								ObjectId:   "teste2",
								Param:      "clovis1",
							},
						},
					},
				},
			},
		}

		badTree1 = compiler.ExecutionTestTree{
			NumScenarios: 1,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "Home",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  0,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions: []lexer.Expect_action{
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu clico no botao teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "click",
								ObjectType: "button",
								ObjectId:   "teste",
								Param:      "clovis1",
							},
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu preencho no campo teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "set",
								ObjectType: "textbox",
								ObjectId:   "teste",
								Param:      "clovis1",
							},
						},
					},
				},
			},
		}

		badTree2 = compiler.ExecutionTestTree{
			NumScenarios: 1,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "HomeExpect2",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  0,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions:  []lexer.Expect_action{},
						Expectations: []lexer.Expect_expression{
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "button",
								ObjectId:   "teste1",
								Param:      "clovis1",
							},
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "textbox",
								ObjectId:   "teste1",
								Param:      "clovis1",
							},
						},
					},
				},
			},
		}

		badTree3 = compiler.ExecutionTestTree{
			NumScenarios: 1,
			Feature: lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "Home",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  0,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions: []lexer.Expect_action{
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu clico no botao teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "click",
								ObjectType: "button",
								ObjectId:   "teste1",
								Param:      "clovis1",
							},
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu preencho no campo teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "set",
								ObjectType: "textbox",
								ObjectId:   "teste2",
								Param:      "clovis1",
							},
						},
						Expectations: []lexer.Expect_expression{
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "button",
								ObjectId:   "teste3",
								Param:      "clovis1",
							},
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "expect",
								ObjectType: "textbox",
								ObjectId:   "teste1",
								Param:      "clovis1",
							},
						},
					},
				},
			},
		}

		definition1 = "page: Home\nuri: <ENTER PAGE URI>\nelements: []\n"
		definition2 = "page: Home\nuri: <ENTER PAGE URI>\nelements:\n- element: teste\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n"
		definition3 = "page: Home\nuri: <ENTER PAGE URI>\nelements:\n- element: teste\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste1\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n"
		definition6 = "page: HomeExpect\nuri: <ENTER PAGE URI>\nelements:\n- element: teste1\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n"
		definition7 = "page: HomeExpect2\nuri: <ENTER PAGE URI>\nelements:\n- element: teste1\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste2\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n"
		definition8 = "page: Home\nuri: <ENTER PAGE URI>\nelements:\n- element: teste1\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste2\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n- element: teste3\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste4\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n"
		definition1Compl = "page: Home\nuri: teste1/home.html\nelements: []\n"
		definition2Compl = "page: Home\nuri: teste1/home.html\nelements:\n- element: teste\n  locator: By.Id()\n  type: button\n"
		definition3Compl = "page: Home\nuri: teste1/home.html\nelements:\n- element: teste\n  locator: By.Id()\n  type: button\n- element: teste1\n  locator: By.Id()\n  type: textbox\n"
		definition8Compl = "page: Home\nuri: teste1/home.html\nelements:\n- element: teste1\n  locator: By.Id()\n  type: button\n- element: teste2\n  locator: By.Id()\n  type: textbox\n- element: teste3\n  locator: By.Id()\n  type: button\n- element: teste4\n  locator: By.Id()\n  type: textbox\n"
	})
	Describe("Generating Page Object Definitions", func() {
		Context("When valid incomplete trees", func() {
			It("Should return string with Page without objects for only Feature", func(done Done) {
				parser = &PageObjectDefParser{}
				definition, err := parser.GetDefinitionFromTree(goodTree2)
				Expect(err).To(BeNil())
				Expect(definition).To(Equal(definition1))
				close(done)
			})
			It("Should return string with Page without objects for Scenario without actions", func(done Done) {
				parser = &PageObjectDefParser{}
				definition, err := parser.GetDefinitionFromTree(goodTree3)
				Expect(err).To(BeNil())
				Expect(definition).To(Equal(definition1))
				close(done)
			})
		})
		Context("When valid complete trees", func() {
			It("Should return string with Page and one Element from action", func(done Done) {
				parser = &PageObjectDefParser{}
				definition, err := parser.GetDefinitionFromTree(goodTree4)
				Expect(err).To(BeNil())
				Expect(definition).To(Equal(definition2))
				close(done)
			})
			It("Should return string with Page and two Element from action", func(done Done) {
				parser = &PageObjectDefParser{}
				definition, err := parser.GetDefinitionFromTree(goodTree5)
				Expect(err).To(BeNil())
				Expect(definition).To(Equal(definition3))
				close(done)
			})
			It("Should return string with Page and one Element from expectation", func(done Done) {
				parser = &PageObjectDefParser{}
				definition, err := parser.GetDefinitionFromTree(goodTree6)
				Expect(err).To(BeNil())
				Expect(definition).To(Equal(definition6))
				close(done)
			})
			It("Should return string with Page and two Element from expectation", func(done Done) {
				parser = &PageObjectDefParser{}
				definition, err := parser.GetDefinitionFromTree(goodTree7)
				Expect(err).To(BeNil())
				Expect(definition).To(Equal(definition7))
				close(done)
			})
			It("Should return string with Page and four Elements from action and expectation", func(done Done) {
				parser = &PageObjectDefParser{}
				definition, err := parser.GetDefinitionFromTree(goodTree1)
				Expect(err).To(BeNil())
				Expect(definition).To(Equal(definition8))
				close(done)
			})
		})
		Context("When invalid complete trees", func() {
			It("Should return err with duplicate action elements with different types", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetDefinitionFromTree(badTree1)
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Duplicated element teste with different types textbox || button"))
				close(done)
			})
		})
	})
	Describe("Parsing Page Object Definition", func() {
		Context("When valid definition", func() {
			It("Should return a Page without elements", func(done Done) {
				//"page: Home\nuri: http://www\nelements: []\n"
				parser = &PageObjectDefParser{}
				yamlPage, err := parser.GetYamlPage(definition1Compl)
				Expect(err).To(BeNil())
				Expect(yamlPage.Page).To(Equal("Home"))
				Expect(len(yamlPage.Elements)).To(Equal(0))
				close(done)
			})
			It("Should return a Page with one element", func(done Done) {
				//"page: Home\nuri: <ENTER PAGE URI>\nelements:\n- element: teste\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n"
				parser = &PageObjectDefParser{}
				yamlPage, err := parser.GetYamlPage(definition2Compl)
				Expect(err).To(BeNil())
				Expect(yamlPage.Page).To(Equal("Home"))
				Expect(len(yamlPage.Elements)).To(Equal(1))
				Expect(yamlPage.Elements[0].Element).To(Equal("teste"))
				Expect(yamlPage.Elements[0].Locator).To(Equal("By.Id()"))
				Expect(yamlPage.Elements[0].Type).To(Equal("button"))
				close(done)
			})
			It("Should return a Page with two elements", func(done Done) {
				//"page: Home\nuri: <ENTER PAGE URI>\nelements:\n- element: teste\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste1\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n"
				parser = &PageObjectDefParser{}
				yamlPage, err := parser.GetYamlPage(definition3Compl)
				Expect(err).To(BeNil())
				Expect(yamlPage.Page).To(Equal("Home"))
				Expect(len(yamlPage.Elements)).To(Equal(2))
				Expect(yamlPage.Elements[0].Element).To(Equal("teste"))
				Expect(yamlPage.Elements[0].Locator).To(Equal("By.Id()"))
				Expect(yamlPage.Elements[0].Type).To(Equal("button"))
				Expect(yamlPage.Elements[1].Element).To(Equal("teste1"))
				Expect(yamlPage.Elements[1].Locator).To(Equal("By.Id()"))
				Expect(yamlPage.Elements[1].Type).To(Equal("textbox"))
				close(done)
			})
			It("Should return a Page with four elements", func(done Done) {
				//"page: Home\nuri: <ENTER PAGE URI>\nelements:\n- element: teste1\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste2\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n- element: teste3\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste4\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n"
				parser = &PageObjectDefParser{}
				yamlPage, err := parser.GetYamlPage(definition8Compl)
				Expect(err).To(BeNil())
				Expect(yamlPage.Page).To(Equal("Home"))
				Expect(len(yamlPage.Elements)).To(Equal(4))
				Expect(yamlPage.Elements[0].Element).To(Equal("teste1"))
				Expect(yamlPage.Elements[0].Locator).To(Equal("By.Id()"))
				Expect(yamlPage.Elements[0].Type).To(Equal("button"))
				Expect(yamlPage.Elements[1].Element).To(Equal("teste2"))
				Expect(yamlPage.Elements[1].Locator).To(Equal("By.Id()"))
				Expect(yamlPage.Elements[1].Type).To(Equal("textbox"))
				Expect(yamlPage.Elements[2].Element).To(Equal("teste3"))
				Expect(yamlPage.Elements[2].Locator).To(Equal("By.Id()"))
				Expect(yamlPage.Elements[2].Type).To(Equal("button"))
				Expect(yamlPage.Elements[3].Element).To(Equal("teste4"))
				Expect(yamlPage.Elements[3].Locator).To(Equal("By.Id()"))
				Expect(yamlPage.Elements[3].Type).To(Equal("textbox"))
				close(done)
			})
		})
		Context("When invalid definition", func() {
			It("Should return a err with wrong definition", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("pag: home\nuri:<ENTER PAGE URI>\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("yaml: line 2: could not find expected ':'"))
				close(done)
			})
			It("Should return a err with missing page name field", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("pag: home\nuri: <ENTER PAGE URI>\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Missing page name"))
				close(done)
			})
			It("Should return a err with unset page uri field", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("page: home\nuri: <ENTER PAGE URI>\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Missing page uri"))
				close(done)
			})
			It("Should return a err with missing page uri field", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("page: home\nur: <ENTER PAGE URI>\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Missing page uri"))
				close(done)
			})
			It("Should return a err with missing element name field", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("page: Home\nuri: http://www\nelements:\n- elemen: teste\n  locator: By.Id()\n  type: button\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Missing element name for locator By.Id() and type button"))
				close(done)
			})
			It("Should return a err with unset element locator field", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("page: Home\nuri: http://www\nelements:\n- element: teste\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Missing element locator for name teste and type button"))
				close(done)
			})
			It("Should return a err with missing element locator field", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("page: Home\nuri: http://www\nelements:\n- element: teste\n  locatr: <ENTER ELEMENT LOCATOR>\n  type: button\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Missing element locator for name teste and type button"))
				close(done)
			})
			It("Should return a err with missing element type field", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("page: Home\nuri: http://www\nelements:\n- element: teste\n  locator: By.Id()\n  tye: button\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Missing element type for locator By.Id() and name teste"))
				close(done)
			})
			It("Should return a err with duplicated element name", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetYamlPage("page: Home\nuri: http://www\nelements:\n- element: teste\n  locator: By.Id()\n  type: button\n- element: teste1\n  locator: By.Id()\n  type: textbox\n- element: teste\n  locator: By.Id()\n  type: textbox\n")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("Duplicated element teste with different types textbox || button"))
				close(done)
			})
		})
	})
	Describe("Getting PageObjects types", func() {
		Context("When valid definitions", func() {
			It("Should return a valid PageObject without elements", func(done Done) {
				//"page: Home\nuri: http://www\nelements: []\n"
				parser = &PageObjectDefParser{}
				pageObj, err := parser.GetPageObject(definition1Compl, "http://localhost:433/site/test")
				Expect(err).To(BeNil())
				Expect(pageObj.Page).To(Equal("Home"))
				Expect(pageObj.Uri).To(Equal("http://localhost:433/site/test/teste1/home.html"))
				Expect(len(pageObj.Elements)).To(Equal(0))
				close(done)
			})
			It("Should return a valid PageObject one element", func(done Done) {
				//"page: Home\nuri: teste1/home.html\nelements:\n- element: teste\n  locator: By.Id()\n  type: button\n"
				parser = &PageObjectDefParser{}
				pageObj, err := parser.GetPageObject(definition2Compl, "http://localhost:433/site/test")
				Expect(err).To(BeNil())
				Expect(pageObj.Page).To(Equal("Home"))
				Expect(pageObj.Uri).To(Equal("http://localhost:433/site/test/teste1/home.html"))
				Expect(len(pageObj.Elements)).To(Equal(1))
				Expect(pageObj.Elements[0].ElementId).To(Equal("teste"))
				Expect(pageObj.Elements[0].ElementType).To(Equal("button"))
				Expect(pageObj.Elements[0].Locator).To(Equal("By.Id()"))
				close(done)
			})
			It("Should return a valid PageObject four elements", func(done Done) {
				//"page: Home\nuri: teste1/home.html\nelements:\n- element: teste1\n  locator: By.Id()\n  type: button\n- element: teste2\n  locator: By.Id()\n  type: textbox\n- element: teste3\n  locator: By.Id()\n  type: button\n- element: teste4\n  locator: By.Id()\n  type: textbox\n"
				parser = &PageObjectDefParser{}
				pageObj, err := parser.GetPageObject(definition8Compl, "http://localhost:433/site/test")
				Expect(err).To(BeNil())
				Expect(pageObj.Page).To(Equal("Home"))
				Expect(pageObj.Uri).To(Equal("http://localhost:433/site/test/teste1/home.html"))
				Expect(len(pageObj.Elements)).To(Equal(4))
				Expect(pageObj.Elements[0].ElementId).To(Equal("teste1"))
				Expect(pageObj.Elements[0].ElementType).To(Equal("button"))
				Expect(pageObj.Elements[0].Locator).To(Equal("By.Id()"))
				Expect(pageObj.Elements[1].ElementId).To(Equal("teste2"))
				Expect(pageObj.Elements[1].ElementType).To(Equal("textbox"))
				Expect(pageObj.Elements[1].Locator).To(Equal("By.Id()"))
				Expect(pageObj.Elements[2].ElementId).To(Equal("teste3"))
				Expect(pageObj.Elements[2].ElementType).To(Equal("button"))
				Expect(pageObj.Elements[2].Locator).To(Equal("By.Id()"))
				Expect(pageObj.Elements[3].ElementId).To(Equal("teste4"))
				Expect(pageObj.Elements[3].ElementType).To(Equal("textbox"))
				Expect(pageObj.Elements[3].Locator).To(Equal("By.Id()"))
				close(done)
			})
		})
		Context("When invalid definition", func() {
			It("Should return an error with invalid definition", func(done Done) {
				parser = &PageObjectDefParser{}
				_, err := parser.GetPageObject("pag: home\nuri:<ENTER PAGE URI>\n", "http://localhost:433/site/test")
				Expect(err).To(Not(BeNil()))
				Expect(err.Error()).To(Equal("yaml: line 2: could not find expected ':'"))
				close(done)
			})
		})
	})
})
