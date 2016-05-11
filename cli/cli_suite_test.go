package cli_test

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
	folders = []string{"../test/CliSpecs","../test/CliSpecs/specs1","../test/CliSpecs/specs2","../test/CliSpecs/specs3"}
)

func TestFlagParser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cli Suite")
}

func prepareFiles() {
	fileHandler := util.NewFileHandler()
			
			destinations["error1.spec"] = []string{"../test/CliSpecs/specs1","../test/CliSpecs/specs2"}
			destinations["good1.spec"] = []string{"../test/CliSpecs/specs1","../test/CliSpecs/specs2"}
			destinations["good1.spec.page"] = []string{"../test/CliSpecs/specs1","../test/CliSpecs/specs2"}
			destinations["good2.spec"] = []string{"../test/CliSpecs/specs1","../test/CliSpecs/specs2"}
			destinations["teste1.spec"] = []string{"../test/CliSpecs/specs3"}
			destinations["teste1.spec.page"] = []string{"../test/CliSpecs/specs3"}
			destinations["teste2.spec"] = []string{"../test/CliSpecs/specs3"}
			destinations["teste2.spec.page"] = []string{"../test/CliSpecs/specs3"}
			
			files["error1.spec"] = `#pt_BR
Aspect: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"`

			files["good1.spec"] = `#pt_BR
Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"`
			files["teste1.spec"] = `#en
Feature: Este é um aspecto
Page: Cadastro de Clientes
Scenario: primeiro cenário
Given que estou usando o user clovis.chedid
When eu click no button teste com o valor "clovis1"
And eu set o textbox teste1 com o valor "clovis2"
Then eu expect a selectbox teste2 com a opcao "clovis3"
`
			files["teste2.spec"] = `#pt_BR
Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"

Cenario: segundo cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"
`
			files["good1.spec.page"] = `page: Cadastro de Clientes
uri: <ENTER PAGE URI>
elements:
- element: teste
  locator: <ENTER ELEMENT LOCATOR>
  type: button
- element: teste1
  locator: <ENTER ELEMENT LOCATOR>
  type: textbox
- element: teste2
  locator: <ENTER ELEMENT LOCATOR>
  type: selectbox
`
			files["teste1.spec.page"] = `page: Home
uri: test/teste1.html
elements:
- element: teste
  locator: By.Id("teste")
  type: button
- element: teste1
  locator: By.Id("teste1")
  type: textbox
- element: teste2
  locator: By.Id("teste2")
  type: selectbox
`
			files["teste2.spec.page"] = `page: Home
uri: test/teste2.html
elements:
- element: teste4
  locator: By.Id("teste4")
  type: textbox
- element: teste5
  locator: By.Id("teste5")
  type: textbox
- element: salvar
  locator: By.Id("salvar")
  type: button
- element: teste6
  locator: By.Id("teste6")
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
