package compiler_test

import (
	. "github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	tree1       ExecutionTestTree
	tree2       ExecutionTestTree
	definition1 string
	definition2 string
)

var _ = Describe("Builder", func() {
	Describe("building specs", func() {
		It("should build a single spec file", func(done Done) {
			builder := &Builder{}
			parsedTree := builder.BuildFile("../test/specs/teste1.spec")
			Expect(parsedTree.HasError).To(BeFalse())
			Expect(parsedTree.NumScenarios).To(Equal(1))
			close(done)
		})
		It("should build a multiple spec files", func(done Done) {
			builder := &Builder{}
			parsedTrees, err := builder.BuildFiles("../test/**/")
			Expect(err).To(BeNil())
			Expect(len(parsedTrees)).To(Equal(5))
			close(done)
		})
	})
	Describe("generating YAML Page Objects", func() {
		BeforeEach(func() {
			tree1 = ExecutionTestTree{
				HasError: false,
				Feature: lexer.Feature{
					LineNum:  1,
					FullText: "Aspecto Primeiro aspecto",
					Label:    "Feature",
					Name:     "Primeiro aspecto",
					PageName: "Home",
					Scenarios: []lexer.Scenario{
						lexer.Scenario{
							LineNum:  3,
							FullText: "Cenario Cenario1",
							Label:    "Scenario",
							Name:     "Cenario1",
							Username: "clovis.chedid",
							Actions: []lexer.Expect_action{
								lexer.Expect_action{
									LineNum:    5,
									FullText:   "Quando eu clico no botao teste com o valor clovis1",
									Label:      "When",
									Action:     "click",
									ObjectType: "button",
									ObjectId:   "teste",
									Param:      "clovis1",
								},
								lexer.Expect_action{
									LineNum:    6,
									FullText:   "E eu preencho o campo teste1 com o valor clovis2",
									Label:      "And",
									Action:     "set",
									ObjectType: "textbox",
									ObjectId:   "teste1",
									Param:      "clovis2",
								},
							},
							Expectations: []lexer.Expect_expression{
								lexer.Expect_expression{
									LineNum:    7,
									FullText:   "Entao eu espero a lista teste2 com a opcao clovis3",
									Label:      "Then",
									Action:     "expect",
									ObjectType: "selectbox",
									ObjectId:   "teste2",
									Param:      "clovis3",
								},
							},
						},
					},
				},
			}

			tree2 = ExecutionTestTree{
				HasError: false,
				Feature: lexer.Feature{
					LineNum:  1,
					FullText: "Aspecto Primeiro aspecto",
					Label:    "Feature",
					Name:     "Primeiro aspecto",
					PageName: "Home",
					Scenarios: []lexer.Scenario{
						lexer.Scenario{
							LineNum:  3,
							FullText: "Cenario Cenario1",
							Label:    "Scenario",
							Name:     "Cenario1",
							Username: "clovis.chedid",
							Actions: []lexer.Expect_action{
								lexer.Expect_action{
									LineNum:    5,
									FullText:   "Quando eu clico no botao teste com o valor clovis1",
									Label:      "When",
									Action:     "set",
									ObjectType: "textbox",
									ObjectId:   "teste4",
									Param:      "clovis1",
								},
								lexer.Expect_action{
									LineNum:    6,
									FullText:   "E eu preencho o campo teste1 com o valor clovis2",
									Label:      "And",
									Action:     "set",
									ObjectType: "textbox",
									ObjectId:   "teste5",
									Param:      "clovis2",
								},
								lexer.Expect_action{
									LineNum:    7,
									FullText:   "E eu preencho o campo teste1 com o valor clovis2",
									Label:      "And",
									Action:     "click",
									ObjectType: "button",
									ObjectId:   "salvar",
									Param:      "clovis2",
								},
							},
							Expectations: []lexer.Expect_expression{
								lexer.Expect_expression{
									LineNum:    7,
									FullText:   "Entao eu espero a lista teste2 com a opcao clovis3",
									Label:      "Then",
									Action:     "expect",
									ObjectType: "textbox",
									ObjectId:   "teste6",
									Param:      "clovis3",
								},
							},
						},
					},
				},
			}

			definition1 = "page: Home\nuri: <ENTER PAGE URI>\nelements:\n- element: teste\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste1\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n- element: teste2\n  locator: <ENTER ELEMENT LOCATOR>\n  type: selectbox\n"
			definition2 = "page: Home\nuri: <ENTER PAGE URI>\nelements:\n- element: teste4\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n- element: teste5\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n- element: salvar\n  locator: <ENTER ELEMENT LOCATOR>\n  type: button\n- element: teste6\n  locator: <ENTER ELEMENT LOCATOR>\n  type: textbox\n"
		})
		It("should generate a single YAML string", func(done Done) {
			builder := &Builder{}
			yaml, err := builder.GenerateYamlPageObject(tree1)
			Expect(err).To(BeNil())
			Expect(yaml).To(Equal(definition1))
			close(done)
		})
		It("should generate multiple YAML strings", func(done Done) {
			builder := &Builder{}
			yamls, err := builder.GenerateYamlPageObjects([]ExecutionTestTree{tree1, tree2})
			Expect(err).To(BeNil())
			Expect(len(yamls)).To(Equal(2))
			Expect(yamls[0]).To(Equal(definition1))
			Expect(yamls[1]).To(Equal(definition2))
			close(done)
		})
	})
})
