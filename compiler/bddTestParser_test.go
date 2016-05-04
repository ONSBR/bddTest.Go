package compiler_test

import (
	"github.com/ONSBR/bddTest.Go/compiler"
	"strings"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ONSBR/bddTest.Go/util"
)

var logT = util.GetLogger("test.compiler") 

var _ = Describe("BddTestParser", func() {

	var (
		goodSpec1 []string
		goodSpec2 []string
		goodSpec3 []string
		goodSpec4 []string
		goodSpec5 []string
		goodSpec6 []string
		goodSpec7 []string
		badSpec1 []string
		badSpec2 []string
		badSpec3 []string
		badSpec4 []string
	)
	
	BeforeEach(func() {
		goodSpec1 = []string{"Aspecto: Este é um aspecto",
			"Cenario: primeiro cenário",
			"Dado que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\""}
		
		goodSpec2 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Cenario: segundo cenário",
			"Dado que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\""}

		goodSpec3 = []string{"Aspecto: Este é um aspecto",
			"Cenario: terceiro cenário",
			"Dado que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\"",
			"Cenario: quarto cenário",
			"Dado que existe a sessao xpto",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu seleciono a lista teste2.2 com a opcao \"eu3\""}
		
		goodSpec4 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Cenario: quinto cenário",
			"Dado que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que existe a sessao xpto",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu seleciono a lista teste2.2 com a opcao \"eu3\""}
		
		goodSpec5 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			""}
		
		goodSpec6 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Cenario: quinto cenário",
			""}
		
		goodSpec7 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Cenario: quinto cenário",
			"Cenario: sexto cenário",
			""}
		
		badSpec1 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Cenario: setimo cenário",
			"Dad que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que existe a sessao xpto",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu seleciono a lista teste2.2 com a opcao \"eu3\""}
		
		badSpec2 = []string{"Aspect: Este é um aspecto",
			"Com uma segunda linha",
			"Cenario: quinto cenário",
			"Dado que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que existe a sessao xpto",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu seleciono a lista teste2.2 com a opcao \"eu3\""}
		
		badSpec3 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Cenari: quinto cenário",
			"Dado que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que existe a sessao xpto",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu seleciono a lista teste2.2 com a opcao \"eu3\""}
		
		badSpec4 = []string{"Dado que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que existe a sessao xpto",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu seleciono a lista teste2.2 com a opcao \"eu3\""}
	})
	
	Describe("parsing string line", func(){
		Context("when string line is correct",func(){
			It("should parse and a feature line with a single scenario",func(done Done){
				token := strings.Join(goodSpec1,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(false))
				Expect(res.Feature.Name).To(Equal("Este é um aspecto"))
				Expect(res.NumScenarios).To(Equal(1))
				Expect(len(res.Feature.Scenarios)).To(Equal(1))
				Expect(res.Feature.Scenarios[0].Name).To(Equal("primeiro cenário"))
				Expect(len(res.Feature.Scenarios[0].Assertions)).To(Equal(4))
				line := res.Feature.Scenarios[0].Assertions[1]
				//Quando eu clico no botao teste com o valor \"clovis1\"
				Expect(line.Label).To(Equal("Quando"))
				Expect(line.ObjectId).To(Equal("teste"))
				Expect(line.ObjectType).To(Equal("botao"))
				Expect(line.Action).To(Equal("clico"))
				Expect(line.Param).To(Equal("clovis1"))
				Expect(res.Feature.LineNum).To(Equal(1))
				Expect(res.Feature.Scenarios[0].LineNum).To(Equal(2))
				Expect(line.LineNum).To(Equal(4))
				close(done)
			})
			It("should parse and a feature block name with a single scenario",func(done Done){
				token := strings.Join(goodSpec2,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(false))
				Expect(res.Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(res.NumScenarios).To(Equal(1))
				Expect(len(res.Feature.Scenarios)).To(Equal(1))
				Expect(res.Feature.Scenarios[0].Name).To(Equal("segundo cenário"))
				Expect(len(res.Feature.Scenarios[0].Assertions)).To(Equal(4))
				line := res.Feature.Scenarios[0].Assertions[1]
				//Quando eu clico no botao teste com o valor \"clovis1\"
				Expect(line.Label).To(Equal("Quando"))
				Expect(line.ObjectId).To(Equal("teste"))
				Expect(line.ObjectType).To(Equal("botao"))
				Expect(line.Action).To(Equal("clico"))
				Expect(line.Param).To(Equal("clovis1"))
				Expect(res.Feature.LineNum).To(Equal(1))
				Expect(res.Feature.Scenarios[0].LineNum).To(Equal(3))
				Expect(line.LineNum).To(Equal(5))
				close(done)
			})
			It("should parse and a feature line with a two scenario",func(done Done){
				token := strings.Join(goodSpec3,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(false))
				Expect(res.Feature.Name).To(Equal("Este é um aspecto"))
				Expect(res.NumScenarios).To(Equal(2))
				Expect(len(res.Feature.Scenarios)).To(Equal(2))
				Expect(res.Feature.Scenarios[0].Name).To(Equal("terceiro cenário"))
				Expect(res.Feature.Scenarios[1].Name).To(Equal("quarto cenário"))
				Expect(len(res.Feature.Scenarios[0].Assertions)).To(Equal(4))
				line := res.Feature.Scenarios[0].Assertions[1]
				//Quando eu clico no botao teste com o valor \"clovis1\"
				Expect(line.Label).To(Equal("Quando"))
				Expect(line.ObjectId).To(Equal("teste"))
				Expect(line.ObjectType).To(Equal("botao"))
				Expect(line.Action).To(Equal("clico"))
				Expect(line.Param).To(Equal("clovis1"))
				Expect(res.Feature.LineNum).To(Equal(1))
				Expect(res.Feature.Scenarios[0].LineNum).To(Equal(2))
				Expect(line.LineNum).To(Equal(4))
				Expect(res.Feature.Scenarios[1].LineNum).To(Equal(7))
				Expect(res.Feature.Scenarios[1].Assertions[1].LineNum).To(Equal(9))
				close(done)
			})
			It("should parse and a feature block with a two scenario",func(done Done){
				token := strings.Join(goodSpec4,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(false))
				Expect(res.Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(res.NumScenarios).To(Equal(2))
				Expect(len(res.Feature.Scenarios)).To(Equal(2))
				Expect(res.Feature.Scenarios[0].Name).To(Equal("quinto cenário"))
				Expect(res.Feature.Scenarios[1].Name).To(Equal("sexto cenário"))
				Expect(len(res.Feature.Scenarios[0].Assertions)).To(Equal(4))
				line := res.Feature.Scenarios[0].Assertions[1]
				//Quando eu clico no botao teste com o valor \"clovis1\"
				Expect(line.Label).To(Equal("Quando"))
				Expect(line.ObjectId).To(Equal("teste"))
				Expect(line.ObjectType).To(Equal("botao"))
				Expect(line.Action).To(Equal("clico"))
				Expect(line.Param).To(Equal("clovis1"))
				Expect(res.Feature.LineNum).To(Equal(1))
				Expect(res.Feature.Scenarios[0].LineNum).To(Equal(3))
				Expect(line.LineNum).To(Equal(5))
				Expect(res.Feature.Scenarios[1].LineNum).To(Equal(8))
				Expect(res.Feature.Scenarios[1].Assertions[1].LineNum).To(Equal(10))
				close(done)
			})
			It("should parse and a feature block without a scenario",func(done Done){
				token := strings.Join(goodSpec5,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(false))
				Expect(res.Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(res.NumScenarios).To(Equal(0))
				Expect(res.Feature.LineNum).To(Equal(1))
				close(done)
			})
			It("should parse and a feature block with a scenario with out assertions",func(done Done){
				token := strings.Join(goodSpec6,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(false))
				Expect(res.Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(res.NumScenarios).To(Equal(1))
				Expect(len(res.Feature.Scenarios)).To(Equal(1))
				Expect(res.Feature.Scenarios[0].Name).To(Equal("quinto cenário"))
				Expect(len(res.Feature.Scenarios[0].Assertions)).To(Equal(0))
				Expect(res.Feature.LineNum).To(Equal(1))
				Expect(res.Feature.Scenarios[0].LineNum).To(Equal(3))
				close(done)
			})
			It("should parse and a feature block with two scenarios with out assertions",func(done Done){
				token := strings.Join(goodSpec7,"\n")
				logT.Errorf("%s",token)
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(false))
				Expect(res.Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(res.NumScenarios).To(Equal(2))
				Expect(len(res.Feature.Scenarios)).To(Equal(2))
				Expect(res.Feature.Scenarios[0].Name).To(Equal("quinto cenário"))
				Expect(res.Feature.Scenarios[1].Name).To(Equal("sexto cenário"))
				Expect(len(res.Feature.Scenarios[0].Assertions)).To(Equal(0))
				Expect(len(res.Feature.Scenarios[1].Assertions)).To(Equal(0))
				Expect(res.Feature.LineNum).To(Equal(1))
				Expect(res.Feature.Scenarios[0].LineNum).To(Equal(3))
				Expect(res.Feature.Scenarios[1].LineNum).To(Equal(4))
				close(done)
			})
		})
		Context("when string line is not correct", func(){
			It("should not parse with wrong Assertion label",func(done Done){
				token := strings.Join(badSpec1,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(true))
				Expect(res.Error.Message).To(Equal("syntax error: unexpected TEXT"))
				Expect(res.Error.LineNum).To(Equal(4))
				Expect(res.Error.LinePos).To(Equal(4))
				Expect(res.Error.Token).To(Equal("Dad "))
				close(done)
			})
			It("should not parse with wrong feature label",func(done Done){
				token := strings.Join(badSpec2,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(true))
				Expect(res.Error.Message).To(Equal("syntax error: unexpected TEXT, expecting FEATURE_LABEL"))
				Expect(res.Error.LineNum).To(Equal(1))
				Expect(res.Error.LinePos).To(Equal(8))
				Expect(res.Error.Token).To(Equal("Aspect: "))
				close(done)
			})
			It("should not parse with wrong scenario label",func(done Done){
				token := strings.Join(badSpec3,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(true))
				Expect(res.Error.Message).To(Equal("syntax error: unexpected LABEL"))
				Expect(res.Error.LineNum).To(Equal(4))
				Expect(res.Error.LinePos).To(Equal(5))
				Expect(res.Error.Token).To(Equal("Dado "))
				close(done)
			})
			It("should not parse with only assertions",func(done Done){
				token := strings.Join(badSpec4,"\n")
				res := compiler.ParseBddTest(token)
				Expect(res.HasError).To(Equal(true))
				Expect(res.Error.Message).To(Equal("syntax error: unexpected LABEL, expecting FEATURE_LABEL"))
				Expect(res.Error.LineNum).To(Equal(1))
				Expect(res.Error.LinePos).To(Equal(5))
				Expect(res.Error.Token).To(Equal("Dado "))
				close(done)
			})
		})
	})
})
