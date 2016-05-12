package compiler_test

import (
	. "github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var logCC = util.GetLogger("test.compiler")

var _ = Describe("codeCompiler", func() {

	var (
		goodSpec []string
		badSpec  []string
		goodFeat lexer.Feature
		badFeat  lexer.Feature
	)

	BeforeEach(func() {
		goodSpec = []string{"Aspecto: Primeiro aspecto",
			"Pagina: Home",
			"Cenario: Cenario1",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			""}

		badSpec = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: setimo cenário",
			"Dad que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que existe a sessao xpto",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
	})
	Describe("Building execution test tree", func() {
		BeforeEach(func() {
			goodFeat = lexer.Feature{
				LineNum:  1,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspecto",
				Name:     "Primeiro aspecto",
				PageName: "Home",
				Scenarios: []lexer.Scenario{
					lexer.Scenario{
						LineNum:  3,
						FullText: "Cenario: Cenario1",
						Label:    "Cenario",
						Name:     "Cenario1",
						Username: "clovis.chedid",
						Actions: []lexer.Expect_action{
							lexer.Expect_action{
								LineNum:    5,
								FullText:   "Quando eu clico no botao teste com o valor clovis1",
								Label:      "Quando",
								Action:     "clico",
								ObjectType: "botao",
								ObjectId:   "teste",
								Param:      "clovis1",
							},
							lexer.Expect_action{
								LineNum:    6,
								FullText:   "E eu preencho o campo teste1 com o valor clovis2",
								Label:      "E",
								Action:     "preencho",
								ObjectType: "campo",
								ObjectId:   "teste1",
								Param:      "clovis2",
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

			badFeat = lexer.Feature{
				LineNum:  0,
				FullText: "Aspecto: Primeiro aspecto",
				Label:    "Aspect",
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
								Action:     "clico",
								ObjectType: "botao",
								ObjectId:   "teste",
								Param:      "clovis1",
							},
							lexer.Expect_action{
								LineNum:    0,
								FullText:   "Quando eu preencho no campo teste com o valor \"clovis1\"",
								Label:      "Quando",
								Action:     "preencho",
								ObjectType: "campo",
								ObjectId:   "teste",
								Param:      "clovis1",
							},
						},
						Expectations: []lexer.Expect_expression{
							lexer.Expect_expression{
								LineNum:    0,
								FullText:   "Entao eu espero no botao teste com o valor \"clovis1\"",
								Label:      "Entao",
								Action:     "espero",
								ObjectType: "botao",
								ObjectId:   "teste",
								Param:      "clovis1",
							},
						},
					},
				},
			}
		})
		Context("when string line is correct", func() {
			It("should generate execution tree ", func(done Done) {
				codeCompiler := &CodeCompiler{}
				token := strings.Join(goodSpec, "\n")
				nFeat := goodFeat
				nFeat.Label = "Feature"
				nFeat.Scenarios[0].Label = "Scenario"
				nFeat.Scenarios[0].Actions[0].Label = "When"
				nFeat.Scenarios[0].Actions[0].Action = "click"
				nFeat.Scenarios[0].Actions[0].ObjectType = "button"
				nFeat.Scenarios[0].Actions[1].Label = "And"
				nFeat.Scenarios[0].Actions[1].Action = "set"
				nFeat.Scenarios[0].Actions[1].ObjectType = "textbox"
				nFeat.Scenarios[0].Expectations[0].Label = "Then"
				nFeat.Scenarios[0].Expectations[0].Action = "expect"
				nFeat.Scenarios[0].Expectations[0].ObjectType = "selectbox"

				executionTree := codeCompiler.BuildExecutionTestTree(token)

				Expect(executionTree.HasError).To(Equal(false))
				Expect(executionTree.NumScenarios).To(Equal(1))
				Expect(executionTree.Feature).To(Equal(nFeat))

				close(done)
			})
			// It("should not generate execution tree with unknown token translation", func(done Done) {
			// 	codeCompiler := &CodeCompiler{}

			// 	token := strings.Join(goodSpec, "\n")

			// 	executionTree := codeCompiler.BuildExecutionTestTree(token)

			// 	Expect(executionTree.HasError).To(Equal(true))
			// 	Expect(executionTree.NumScenarios).To(Equal(0))
			// 	Expect(executionTree.Error).To(Equal(lexer.ParserError{Message: "Invalid feature token translation", Token: "Aspect"}))

			// 	close(done)
			// })
		})
		Context("when string line is incorrect", func() {
			It("should not generate execution tree with syntax error", func(done Done) {
				codeCompiler := &CodeCompiler{}
				token := strings.Join(badSpec, "\n")

				executionTree := codeCompiler.BuildExecutionTestTree(token)

				Expect(executionTree.HasError).To(Equal(true))
				Expect(executionTree.NumScenarios).To(Equal(0))
				Expect(executionTree.Error).To(Equal(lexer.ParserError{Message: "syntax error: unexpected TEXT", Token: "Dad ", LineNum: 5, LinePos: 4}))

				close(done)
			})
		})
	})
})
