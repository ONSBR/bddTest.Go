package compiler_test

import (
	. "github.com/ONSBR/bddTest.Go/compiler"
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
		badSpec5 []string
		badSpec6 []string
		badSpec7 []string
		badSpec8 []string
	)
	
	BeforeEach(func() {
		goodSpec1 = []string{"Aspecto: Este é um aspecto",
			"Pagina: Cadastro de Clientes",
			"Cenario: primeiro cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		goodSpec2 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: segundo cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}

		goodSpec3 = []string{"Aspecto: Este é um aspecto",
			"Pagina: Cadastro de Clientes",
			"Cenario: terceiro cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: quarto cenário",
			"Dado que estou usando o usuario clovis.chedid2",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		goodSpec4 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: quinto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		goodSpec5 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			""}
		
		goodSpec6 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: quinto cenário",
			""}
		
		goodSpec7 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: quinto cenário",
			"Cenario: sexto cenário",
			""}
		
		badSpec1 = []string{"Aspecto: Este é um aspecto",
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
		
		badSpec2 = []string{"Aspect: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: quinto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		badSpec3 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenari: quinto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		badSpec4 = []string{"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		badSpec5 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Cenario: quinto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		badSpec6 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: quinto cenário",
			"Dado que estou usando o usuar clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		badSpec7 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: quinto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Ento eu espero a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
		
		badSpec8 = []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Pagina: Cadastro de Clientes",
			"Cenario: quinto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu esper a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que estou usando o usuario clovis.chedid",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu espero a lista teste2 com a opcao \"clovis3\""}
	})
	
	Describe("parsing string line", func(){
		Context("when string line is correct",func(){
			It("should parse and a feature line with a single scenario",func(done Done){
				token := strings.Join(goodSpec1,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = err
				Expect(retCode).To(Equal(0))
				Expect(Feature.Name).To(Equal("Este é um aspecto"))
				Expect(len(Feature.Scenarios)).To(Equal(1))
				Expect(Feature.Scenarios[0].Name).To(Equal("primeiro cenário"))
				Expect(len(Feature.Scenarios[0].Actions)).To(Equal(2))
				Expect(len(Feature.Scenarios[0].Expectations)).To(Equal(1))
				line := Feature.Scenarios[0].Actions[0]
				//Quando eu clico no botao teste com o valor \"clovis1\"
				Expect(line.Label).To(Equal("Quando"))
				Expect(line.ObjectId).To(Equal("teste"))
				Expect(line.ObjectType).To(Equal("botao"))
				Expect(line.Action).To(Equal("clico"))
				Expect(line.Param).To(Equal("clovis1"))
				Expect(Feature.LineNum).To(Equal(1))
				Expect(Feature.Scenarios[0].LineNum).To(Equal(3))
				Expect(line.LineNum).To(Equal(5))
				close(done)
			})
			It("should parse and a feature block name with a single scenario",func(done Done){
				token := strings.Join(goodSpec2,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = err
				Expect(retCode).To(Equal(0))
				Expect(Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(len(Feature.Scenarios)).To(Equal(1))
				Expect(Feature.Scenarios[0].Name).To(Equal("segundo cenário"))
				Expect(len(Feature.Scenarios[0].Actions)).To(Equal(2))
				Expect(len(Feature.Scenarios[0].Expectations)).To(Equal(1))
				line := Feature.Scenarios[0].Actions[0]
				//Quando eu clico no botao teste com o valor \"clovis1\"
				Expect(line.Label).To(Equal("Quando"))
				Expect(line.ObjectId).To(Equal("teste"))
				Expect(line.ObjectType).To(Equal("botao"))
				Expect(line.Action).To(Equal("clico"))
				Expect(line.Param).To(Equal("clovis1"))
				Expect(Feature.LineNum).To(Equal(1))
				Expect(Feature.Scenarios[0].LineNum).To(Equal(4))
				Expect(line.LineNum).To(Equal(6))
				close(done)
			})
			It("should parse and a feature line with a two scenario",func(done Done){
				token := strings.Join(goodSpec3,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = err
				Expect(retCode).To(Equal(0))
				Expect(Feature.Name).To(Equal("Este é um aspecto"))
				Expect(len(Feature.Scenarios)).To(Equal(2))
				Expect(Feature.Scenarios[0].Name).To(Equal("terceiro cenário"))
				Expect(Feature.Scenarios[1].Name).To(Equal("quarto cenário"))
				Expect(len(Feature.Scenarios[0].Actions)).To(Equal(2))
				Expect(len(Feature.Scenarios[0].Expectations)).To(Equal(1))
				Expect(len(Feature.Scenarios[1].Actions)).To(Equal(2))
				Expect(len(Feature.Scenarios[1].Expectations)).To(Equal(1))
				line := Feature.Scenarios[0].Actions[0]
				//Quando eu clico no botao teste com o valor \"clovis1\"
				Expect(line.Label).To(Equal("Quando"))
				Expect(line.ObjectId).To(Equal("teste"))
				Expect(line.ObjectType).To(Equal("botao"))
				Expect(line.Action).To(Equal("clico"))
				Expect(line.Param).To(Equal("clovis1"))
				Expect(Feature.LineNum).To(Equal(1))
				Expect(Feature.Scenarios[0].LineNum).To(Equal(3))
				Expect(line.LineNum).To(Equal(5))
				Expect(Feature.Scenarios[1].LineNum).To(Equal(8))
				Expect(Feature.Scenarios[1].Actions[1].LineNum).To(Equal(11))
				
				eline := Feature.Scenarios[0].Expectations[0]
				Expect(eline.Label).To(Equal("Entao"))
				Expect(eline.ObjectId).To(Equal("teste2"))
				Expect(eline.ObjectType).To(Equal("lista"))
				Expect(eline.Action).To(Equal("espero"))
				Expect(eline.Param).To(Equal("clovis3"))
				Expect(eline.LineNum).To(Equal(7))
				close(done)
			})
			It("should parse and a feature block with a two scenario",func(done Done){
				token := strings.Join(goodSpec4,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = err
				Expect(retCode).To(Equal(0))
				Expect(Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(len(Feature.Scenarios)).To(Equal(2))
				Expect(Feature.Scenarios[0].Name).To(Equal("quinto cenário"))
				Expect(Feature.Scenarios[1].Name).To(Equal("sexto cenário"))
				Expect(len(Feature.Scenarios[0].Actions)).To(Equal(2))
				Expect(len(Feature.Scenarios[0].Expectations)).To(Equal(1))
				line := Feature.Scenarios[0].Actions[0]
				//Quando eu clico no botao teste com o valor \"clovis1\"
				Expect(line.Label).To(Equal("Quando"))
				Expect(line.ObjectId).To(Equal("teste"))
				Expect(line.ObjectType).To(Equal("botao"))
				Expect(line.Action).To(Equal("clico"))
				Expect(line.Param).To(Equal("clovis1"))
				Expect(Feature.LineNum).To(Equal(1))
				Expect(Feature.Scenarios[0].LineNum).To(Equal(4))
				Expect(line.LineNum).To(Equal(6))
				Expect(Feature.Scenarios[1].LineNum).To(Equal(9))
				Expect(Feature.Scenarios[1].Actions[1].LineNum).To(Equal(12))
				close(done)
			})
			It("should parse and a feature block without a scenario",func(done Done){
				token := strings.Join(goodSpec5,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = err
				Expect(retCode).To(Equal(0))
				Expect(Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(Feature.LineNum).To(Equal(1))
				Expect(len(Feature.Scenarios)).To(Equal(0))
				close(done)
			})
			It("should parse and a feature block with a scenario with out assertions",func(done Done){
				token := strings.Join(goodSpec6,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = err
				Expect(retCode).To(Equal(0))
				Expect(Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(len(Feature.Scenarios)).To(Equal(1))
				Expect(Feature.Scenarios[0].Name).To(Equal("quinto cenário"))
				Expect(len(Feature.Scenarios[0].Actions)).To(Equal(0))
				Expect(len(Feature.Scenarios[0].Expectations)).To(Equal(0))
				Expect(Feature.LineNum).To(Equal(1))
				Expect(Feature.Scenarios[0].LineNum).To(Equal(4))
				close(done)
			})
			It("should parse and a feature block with two scenarios with out assertions",func(done Done){
				token := strings.Join(goodSpec7,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = err
				Expect(retCode).To(Equal(0))
				Expect(Feature.Name).To(Equal("Este é um aspecto\nCom uma segunda linha"))
				Expect(len(Feature.Scenarios)).To(Equal(2))
				Expect(Feature.Scenarios[0].Name).To(Equal("quinto cenário"))
				Expect(Feature.Scenarios[1].Name).To(Equal("sexto cenário"))
				Expect(len(Feature.Scenarios[0].Actions)).To(Equal(0))
				Expect(len(Feature.Scenarios[1].Actions)).To(Equal(0))
				Expect(len(Feature.Scenarios[0].Expectations)).To(Equal(0))
				Expect(len(Feature.Scenarios[1].Expectations)).To(Equal(0))
				Expect(Feature.LineNum).To(Equal(1))
				Expect(Feature.Scenarios[0].LineNum).To(Equal(4))
				Expect(Feature.Scenarios[1].LineNum).To(Equal(5))
				close(done)
			})
		})
		Context("when string line is not correct", func(){
			It("should not parse with wrong Assertion label",func(done Done){
				token := strings.Join(badSpec1,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = Feature
				Expect(retCode).To(BeNumerically(">",0))
				Expect(err.Message).To(Equal("syntax error: unexpected TEXT"))
				Expect(err.LineNum).To(Equal(5))
				Expect(err.LinePos).To(Equal(4))
				Expect(err.Token).To(Equal("Dad "))
				close(done)
			})
			It("should not parse with wrong feature label",func(done Done){
				token := strings.Join(badSpec2,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = Feature
				Expect(retCode).To(BeNumerically(">",0))
				Expect(err.Message).To(Equal("syntax error: unexpected TEXT, expecting FEATURE_LABEL"))
				Expect(err.LineNum).To(Equal(1))
				Expect(err.LinePos).To(Equal(8))
				Expect(err.Token).To(Equal("Aspect: "))
				close(done)
			})
			It("should not parse with wrong scenario label",func(done Done){
				token := strings.Join(badSpec3,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = Feature
				Expect(retCode).To(BeNumerically(">",0))
				Expect(err.Message).To(Equal("syntax error: unexpected TEXT"))
				Expect(err.LineNum).To(Equal(4))
				Expect(err.LinePos).To(Equal(8))
				Expect(err.Token).To(Equal("Cenari: "))
				close(done)
			})
			It("should not parse with only assertions",func(done Done){
				token := strings.Join(badSpec4,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = Feature
				Expect(retCode).To(BeNumerically(">",0))
				Expect(err.Message).To(Equal("syntax error: unexpected INIT_SCENARIO_LABEL, expecting FEATURE_LABEL"))
				Expect(err.LineNum).To(Equal(1))
				Expect(err.LinePos).To(Equal(5))
				Expect(err.Token).To(Equal("Dado "))
				close(done)
			})
			It("should not parse without page label",func(done Done){
				token := strings.Join(badSpec5,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = Feature
				Expect(retCode).To(BeNumerically(">",0))
				Expect(err.Message).To(Equal("syntax error: unexpected SCENARIO_LABEL, expecting NEW_LINE or TEXT"))
				Expect(err.LineNum).To(Equal(3))
				Expect(err.LinePos).To(Equal(9))
				Expect(err.Token).To(Equal("Cenario: "))
				close(done)
			})
			It("should not parse without usuario label",func(done Done){
				token := strings.Join(badSpec6,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = Feature
				Expect(retCode).To(BeNumerically(">",0))
				Expect(err.Message).To(Equal("syntax error: unexpected NEW_LINE, expecting TEXT or USER_SCENARIO_LABEL"))
				Expect(err.LineNum).To(Equal(6))
				Expect(err.LinePos).To(Equal(1))
				Expect(err.Token).To(Equal("\nQ"))
				close(done)
			})
			It("should not parse with bad expectation label",func(done Done){
				token := strings.Join(badSpec7,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = Feature
				Expect(retCode).To(BeNumerically(">",0))
				Expect(err.Message).To(Equal("syntax error: unexpected TEXT, expecting EXPECT_ACTION_LABEL or LABEL"))
				Expect(err.LineNum).To(Equal(8))
				Expect(err.LinePos).To(Equal(5))
				Expect(err.Token).To(Equal("Ento "))
				close(done)
			})
			It("should not parse with bad expectation action label",func(done Done){
				token := strings.Join(badSpec8,"\n")
				Feature, retCode, err := ParseCode(token)
				_ = Feature
				Expect(retCode).To(BeNumerically(">",0))
				Expect(err.Message).To(Equal("syntax error: unexpected OBJECT_TYPE, expecting EXPECT_ACTION_ACTION or TEXT"))
				Expect(err.LineNum).To(Equal(8))
				Expect(err.LinePos).To(Equal(23))
				Expect(err.Token).To(Equal("lista "))
				close(done)
			})
		})
	})
})
