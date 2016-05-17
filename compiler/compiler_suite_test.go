package compiler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ONSBR/bddTest.Go/util"
	_ "github.com/ONSBR/bddTest.Go/compiler/lexer/locales"
	_ "github.com/ONSBR/bddTest.Go/compiler/langMaps"

	"testing"
)

var (
	destinations = map[string][]string{}
	files = map[string]string{}
	folders = []string{"../test/BuilderSpecs","../test/BuilderSpecs/specs1","../test/BuilderSpecs/specs2"}
)

func TestCompiler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Compiler Suite")
}

func prepareFiles() {
	fileHandler := util.NewFileHandler()
			
			destinations["1.spec"] = []string{"../test/BuilderSpecs/specs2"}
			destinations["2.spec"] = []string{"../test/BuilderSpecs/specs2"}
			destinations["Readme.txt"] = []string{"../test/BuilderSpecs/specs1"}
			destinations["teste1.spec"] = []string{"../test/BuilderSpecs/specs1"}
			destinations["teste1.spec.page"] = []string{"../test/BuilderSpecs/specs1"}
			destinations["teste2.spec"] = []string{"../test/BuilderSpecs/specs1"}
			destinations["teste2.spec.page"] = []string{"../test/BuilderSpecs/specs1"}
			destinations["teste3.spec"] = []string{"../test/BuilderSpecs/specs1"}
			destinations["guide.script"] = []string{"../test/BuilderSpecs"}
			destinations["guide2.script"] = []string{"../test/BuilderSpecs"}
			
			files["guide.script"] = `../test/BuilderSpecs/specs1/teste1.spec
../test/BuilderSpecs/specs1/teste2.spec`
			files["guide2.script"] = `../test/CliSpecs/specs1/error1.spec
../test/CliSpecs/specs3/teste1.spec
../test/CliSpecs/specs3/teste2.spec`
			
			files["1.spec"] = `#pt_BR
Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao maior que "clovis3"`

			files["2.spec"] = `#pt_BR
Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao menor que "clovis3"
Entao eu espero o campo teste4 com o valor igual a "clovis3"`

			files["Readme.txt"] = `page: Home
uri: test/teste1.html
elements:
- element: teste
  locator: Id
  expression: teste
  type: button
- element: teste1
  locator: Id
  expression: teste1
  type: textbox
- element: teste2
  locator: Id
  expression: teste2
  type: selectbox


page: Home
uri: test/teste2.html
elements:
- element: teste4
  locator: Id
  expression: teste4
  type: textbox
- element: teste5
  locator: Id
  expression: teste5
  type: textbox
- element: salvar
  locator: Id
  expression: salvar
  type: button
- element: teste6
  locator: Id
  expression: teste6
  type: textbox
`
			files["teste1.spec"] = `#pt_BR
Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao menor que "clovis3"
`
			files["teste2.spec"] = `#pt_BR
Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao igual a "clovis3"

Cenario: segundo cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao contem "clovis3"
`
			files["teste3.spec"] = `#en
Feature: Este é um aspecto
Page: Cadastro de Clientes
Scenario: primeiro cenário
Given que estou usando o user clovis.chedid
When eu click no button teste com o valor "clovis1"
And eu select o textbox teste1 com o valor "clovis2"
Then eu expect a selectbox teste2 com a opcao grater then or equal "clovis3"`
			files["teste1.spec.page"] = `page: Home
uri: test/teste1.html
elements:
- element: teste
  locator: Id
  expression: teste
  type: button
- element: teste1
  locator: Id
  expression: teste1
  type: textbox
- element: teste2
  locator: Id
  expression: teste2
  type: selectbox
`
			files["teste2.spec.page"] = `page: Home
uri: test/teste2.html
elements:
- element: teste4
  locator: Id
  expression: teste4
  type: textbox
- element: teste5
  locator: Id
  expression: teste5
  type: textbox
- element: salvar
  locator: Id
  expression: salvar
  type: button
- element: teste6
  locator: Id
  expression: teste6
  type: textbox
`
			for _,f := range folders {
				fileHandler.CreateFolder(f)
			}
			for k,v := range files {
				dest := destinations[k]
				for _,d := range dest {
					_ = fileHandler.WriteFile(d+"/"+k,v)
				}
			}
}

func removeFiles()  {
	fileHandler := util.NewFileHandler()
			for k := range files {
				dest := destinations[k]
				for _,d := range dest {
					_ = fileHandler.RemoveFile(d+"/"+k)
				}
			}
			for i := range folders {
				f := folders[len(folders)-1-i]
				fileHandler.RemoveFolder(f)
			}
}
