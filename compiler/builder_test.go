package compiler_test

import (
	. "github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/pageObject"
	"github.com/ONSBR/bddTest.Go/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
	"os"
)

var (
	tree1       ExecutionTestTree
	tree2       ExecutionTestTree
	definition1 string
	definition2 string
	page1       *pageObject.PageObject
	page2       *pageObject.PageObject
)

var _ = Describe("Builder", func() {
	Describe("building specs", func() {
		It("should build a single spec file", func(done Done) {
			builder := NewBuilder()
			parsedTree := builder.BuildFile("../test/specs/teste1.spec")
			Expect(parsedTree.HasError).To(BeFalse())
			Expect(parsedTree.NumScenarios).To(Equal(1))
			close(done)
		})
		It("should build a multiple spec files", func(done Done) {
			builder := NewBuilder()
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
			builder := NewBuilder()
			yaml, err := builder.GenerateYamlPageObject(tree1)
			Expect(err).To(BeNil())
			Expect(yaml).To(Equal(definition1))
			close(done)
		})
		It("should generate multiple YAML strings", func(done Done) {
			builder := NewBuilder()
			yamls, err := builder.GenerateYamlPageObjects([]ExecutionTestTree{tree1, tree2})
			Expect(err).To(BeNil())
			Expect(len(yamls)).To(Equal(2))
			Expect(yamls[0]).To(Equal(definition1))
			Expect(yamls[1]).To(Equal(definition2))
			close(done)
		})
	})
	Describe("generating Page Objects", func() {
		BeforeEach(func() {
			page1 = pageObject.NewPageObject("Home", "http://localhost:3000/test/teste1.html")
			_ = pageObject.NewPageElement(page1, "By.Id(\"teste\")", "button", "teste")
			_ = pageObject.NewPageElement(page1, "By.Id(\"teste1\")", "textbox", "teste1")
			_ = pageObject.NewPageElement(page1, "By.Id(\"teste2\")", "selectbox", "teste2")

			page2 = pageObject.NewPageObject("Home", "http://localhost:3000/test/teste2.html")
			_ = pageObject.NewPageElement(page2, "By.Id(\"teste4\")", "textbox", "teste4")
			_ = pageObject.NewPageElement(page2, "By.Id(\"teste5\")", "textbox", "teste5")
			_ = pageObject.NewPageElement(page2, "By.Id(\"salvar\")", "button", "salvar")
			_ = pageObject.NewPageElement(page2, "By.Id(\"teste6\")", "textbox", "teste6")
		})
		It("should generate a single Page Object based on page file", func(done Done) {
			builder := NewBuilder()
			pageObject, err := builder.GeneratePageObject("../test/specs/teste1.spec.page", "http://localhost:3000")
			Expect(err).To(BeNil())
			Expect(len(pageObject.Elements)).To(Equal(3))
			Expect(pageObject).To(Equal(*page1))
			close(done)
		})
	})
	Describe("generating execution tree", func() {
		BeforeEach(func() {
			tree1 = ExecutionTestTree{
				HasError: false,
				Feature: lexer.Feature{
					LineNum:  1,
					FullText: "Aspecto Este é um aspecto",
					Label:    "Feature",
					Name:     "Este é um aspecto",
					PageName: "Cadastro de Clientes",
					Scenarios: []lexer.Scenario{
						lexer.Scenario{
							LineNum:  3,
							FullText: "Cenario primeiro cenário",
							Label:    "Scenario",
							Name:     "primeiro cenário",
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
					FullText: "Aspecto Este é um aspecto",
					Label:    "Feature",
					Name:     "Este é um aspecto",
					PageName: "Cadastro de Clientes",
					Scenarios: []lexer.Scenario{
						lexer.Scenario{
							LineNum:  3,
							FullText: "Cenario primeiro cenário",
							Label:    "Scenario",
							Name:     "primeiro cenário",
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
						lexer.Scenario{
							LineNum:  8,
							FullText: "Cenario segundo cenário",
							Label:    "Scenario",
							Name:     "segundo cenário",
							Username: "clovis.chedid",
							Actions: []lexer.Expect_action{
								lexer.Expect_action{
									LineNum:    10,
									FullText:   "Quando eu clico no botao teste com o valor clovis1",
									Label:      "When",
									Action:     "click",
									ObjectType: "button",
									ObjectId:   "teste",
									Param:      "clovis1",
								},
								lexer.Expect_action{
									LineNum:    11,
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
									LineNum:    12,
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
			page1 = pageObject.NewPageObject("Home", "http://localhost:3000/test/teste1.html")
			_ = pageObject.NewPageElement(page1, "By.Id(\"teste\")", "button", "teste")
			_ = pageObject.NewPageElement(page1, "By.Id(\"teste1\")", "textbox", "teste1")
			_ = pageObject.NewPageElement(page1, "By.Id(\"teste2\")", "selectbox", "teste2")

			page2 = pageObject.NewPageObject("Home", "http://localhost:3000/test/teste2.html")
			_ = pageObject.NewPageElement(page2, "By.Id(\"teste4\")", "textbox", "teste4")
			_ = pageObject.NewPageElement(page2, "By.Id(\"teste5\")", "textbox", "teste5")
			_ = pageObject.NewPageElement(page2, "By.Id(\"salvar\")", "button", "salvar")
			_ = pageObject.NewPageElement(page2, "By.Id(\"teste6\")", "textbox", "teste6")
		})
		It("should generate a execution object based on a single spec file", func(done Done) {
			builder := NewBuilder()
			filename := "../test/specs/teste1.spec"
			baseUri := "http://localhost:3000"
			execution := builder.BuildExecution(filename, baseUri)
			Expect(execution.HasError).To(BeFalse())
			Expect(execution.Filename).To(Equal(filename))
			Expect(execution.ExecutionTree.HasError).To(BeFalse())
			Expect(execution.ExecutionTree.Feature).To(Equal(tree1.Feature))
			Expect(execution.PageObject).To(Equal(*page1))
			close(done)
		})
		It("should generate a collection of execution objects based on a spec folder pattern", func(done Done) {
			builder := NewBuilder()
			folderPattern := "../test/**/"
			baseUri := "http://localhost:3000"
			filename1 := "../test/specs/teste1.spec"
			filename2 := "../test/specs/teste2.spec"
			executions, err := builder.BuildExecutions(folderPattern, baseUri)
			Expect(err).To(BeNil())
			Expect(len(executions)).To(Equal(5))

			for _, execution := range executions {
				if execution.Filename == filename1 {
					Expect(execution.HasError).To(BeFalse())
					Expect(execution.Filename).To(Equal(filename1))
					Expect(execution.ExecutionTree.HasError).To(BeFalse())
					Expect(execution.ExecutionTree.Feature).To(Equal(tree1.Feature))
					Expect(execution.PageObject).To(Equal(*page1))
				} else if execution.Filename == filename2 {
					Expect(execution.HasError).To(BeFalse())
					Expect(execution.Filename).To(Equal(filename2))
					Expect(execution.ExecutionTree.HasError).To(BeFalse())
					Expect(execution.ExecutionTree.Feature).To(Equal(tree2.Feature))
					Expect(execution.PageObject).To(Equal(*page2))
				} else {
					Expect(execution.HasError).To(BeTrue())
					Expect(execution.Error).To(Equal(fmt.Sprintf("Page Object file missing: %s.page", execution.Filename)))
				}
			}

			close(done)
		})
	})
	Describe("generating page files", func() {
		BeforeEach(func() {
			os.Remove("../test/specs/teste3.spec.page")
			os.Remove("../test/specifications/1.spec.page")
			os.Remove("../test/specifications/2.spec.page")
			os.Rename("../test/specs/teste1.spec.page.bkp", "../test/specs/teste1.spec.page")
			os.Rename("../test/specs/teste2.spec.page.bkp", "../test/specs/teste2.spec.page")
		})
		It("should write a spec.page file based on a filename", func(done Done) {
			builder := NewBuilder()
			fileHandler := util.NewFileHandler()

			filename1 := "../test/specs/teste3.spec"

			err := builder.BuildYamlPageObjectFile(filename1,false)
			Expect(err).To(BeNil())
			Expect(fileHandler.DoesFileExists(filename1 + ".page")).To(BeTrue())
			close(done)
		})
		
		It("should write a spec.page file based on a filename and backup existing one", func(done Done) {
			builder := NewBuilder()
			fileHandler := util.NewFileHandler()

			filename1 := "../test/specs/teste2.spec"
			pageContent,_ := fileHandler.ReadFile(filename1 + ".page")

			err := builder.BuildYamlPageObjectFile(filename1,true)
			Expect(err).To(BeNil())
			Expect(fileHandler.DoesFileExists(filename1 + ".page")).To(BeTrue())
			Expect(fileHandler.DoesFileExists(filename1 + ".page.bkp")).To(BeTrue())
			
			_ = fileHandler.WriteFile(filename1 + ".page", pageContent)
			_ = fileHandler.RemoveFile(filename1 + ".page.bkp")
			
			close(done)
		})
		
		It("should write spec.page files based on a folder pattern", func(done Done) {
			builder := NewBuilder()
			fileHandler := util.NewFileHandler()

			folderPattern := "../test/**/"

			err := builder.BuildYamlPageObjectFiles(folderPattern,true)
			Expect(len(err.(*BuilderError).Errors)).To(Equal(0))

			files, _ := fileHandler.FindFiles(folderPattern + "*.spec.page")
			filesBkp, _ := fileHandler.FindFiles(folderPattern + "*.spec.page.bkp")

			Expect(len(files)).To(Equal(5))
			Expect(len(filesBkp)).To(Equal(2))
			close(done)
		})
		AfterEach(func() {
			os.Remove("../test/specs/teste3.spec.page")
			os.Remove("../test/specifications/1.spec.page")
			os.Remove("../test/specifications/2.spec.page")
			os.Rename("../test/specs/teste1.spec.page.bkp", "../test/specs/teste1.spec.page")
			os.Rename("../test/specs/teste2.spec.page.bkp", "../test/specs/teste2.spec.page")
		})
	})
})
