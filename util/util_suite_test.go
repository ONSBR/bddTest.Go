package util_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	"github.com/ONSBR/bddTest.Go/util"

	"testing"
)

var (
	destinations = map[string][]string{}
	files = map[string]string{}
	folders = []string{"../test/FileHandlerSpecs","../test/FileHandlerSpecs/specs1","../test/FileHandlerSpecs/specs2"}
)

func TestUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Util Suite")
}




func prepareFiles() {
	fileHandler := util.NewFileHandler()
			
			destinations["1.spec"] = []string{"../test/FileHandlerSpecs/specs2"}
			destinations["2.spec"] = []string{"../test/FileHandlerSpecs/specs2"}
			destinations["Readme.txt"] = []string{"../test/FileHandlerSpecs/specs1"}
			destinations["teste1.spec"] = []string{"../test/FileHandlerSpecs/specs1"}
			destinations["teste1.spec.page"] = []string{"../test/FileHandlerSpecs/specs1"}
			destinations["teste2.spec"] = []string{"../test/FileHandlerSpecs/specs1"}
			destinations["teste2.spec.page"] = []string{"../test/FileHandlerSpecs/specs1"}
			destinations["teste3.spec"] = []string{"../test/FileHandlerSpecs/specs1"}
			
			files["1.spec"] = `Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"`

			files["2.spec"] = `Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"
Entao eu espero o campo teste4 com o valor "clovis3"`

			files["Readme.txt"] = `page: Home
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


page: Home
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
			files["teste1.spec"] = `Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"
`
			files["teste2.spec"] = `Aspecto: Este é um aspecto
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
			files["teste3.spec"] = `Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu seleciono o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"`
			files["teste1.spec.page"] = `Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu seleciono o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"`
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