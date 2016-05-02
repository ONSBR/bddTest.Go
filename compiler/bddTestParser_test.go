package compiler_test

import (
	"github.com/ONSBR/bddTest.Go/compiler"
	"strings"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BddTestParser", func() {

	var (
		goodLines []string
		badLines []string
		lines []string
	)
	
	BeforeEach(func() {
		goodLines = []string{
			"quando eu clico no botao teste com o valor \"clovis\"", 
			"quando eu preencho o campo teste1 com o valor \"clovis\"",
			"quando eu seleciono na lista teste2 a opcao \"clovis\"",
			"dado eu clico no botao teste com o valor \"clovis\"", 
			"entao eu clico no botao teste com o valor \"clovis\"", 
			"e eu clico no botao teste com o valor \"clovis\"", 
			"Quando eu clico no botao teste com o valor \"clovis\"",
			"Dado eu clico no botao teste com o valor \"clovis\"", 
			"Entao eu clico no botao teste com o valor \"clovis\"", 
			"E eu clico no botao teste com o valor \"clovis\"", }		
		badLines = []string{
			"quand eu clico no botao teste com o valor \"clovis\"", 
			"quando eu preench o campo teste1 com o valor \"clovis\"",
			"quando eu seleciono na list teste2 a opcao \"clovis\""}
	})
	
	Describe("parsing string line", func(){
		Context("when string line is correct",func(){
			BeforeEach(func(){
				lines = goodLines
			})
			It("should parse and return button click command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[0]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("quando"))
					Expect(line.ObjectId).To(Equal("teste"))
					Expect(line.ObjectType).To(Equal("botao"))
					Expect(line.Action).To(Equal("clico"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return selection select command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[2]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("quando"))
					Expect(line.ObjectId).To(Equal("teste2"))
					Expect(line.ObjectType).To(Equal("lista"))
					Expect(line.Action).To(Equal("seleciono"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return field complete command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[1]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("quando"))
					Expect(line.ObjectId).To(Equal("teste1"))
					Expect(line.ObjectType).To(Equal("campo"))
					Expect(line.Action).To(Equal("preencho"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return button click command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[3]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("dado"))
					Expect(line.ObjectId).To(Equal("teste"))
					Expect(line.ObjectType).To(Equal("botao"))
					Expect(line.Action).To(Equal("clico"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return button click command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[4]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("entao"))
					Expect(line.ObjectId).To(Equal("teste"))
					Expect(line.ObjectType).To(Equal("botao"))
					Expect(line.Action).To(Equal("clico"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return button click command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[5]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("e"))
					Expect(line.ObjectId).To(Equal("teste"))
					Expect(line.ObjectType).To(Equal("botao"))
					Expect(line.Action).To(Equal("clico"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return button click command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[6]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("Quando"))
					Expect(line.ObjectId).To(Equal("teste"))
					Expect(line.ObjectType).To(Equal("botao"))
					Expect(line.Action).To(Equal("clico"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return button click command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[7]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("Dado"))
					Expect(line.ObjectId).To(Equal("teste"))
					Expect(line.ObjectType).To(Equal("botao"))
					Expect(line.Action).To(Equal("clico"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return button click command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[8]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("Entao"))
					Expect(line.ObjectId).To(Equal("teste"))
					Expect(line.ObjectType).To(Equal("botao"))
					Expect(line.Action).To(Equal("clico"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
			It("should parse and return button click command",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[9]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(false))
					line := res.Lines[0]
					Expect(line.Label).To(Equal("E"))
					Expect(line.ObjectId).To(Equal("teste"))
					Expect(line.ObjectType).To(Equal("botao"))
					Expect(line.Action).To(Equal("clico"))
					Expect(line.Param).To(Equal("clovis"))
					close(done)
				})
			})
		})
		Context("when string line is not correct", func(){
			BeforeEach(func(){
				lines = badLines
			})
			It("should not parse with wrong label",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[0]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(true))
					Expect(res.Error.Message).To(Equal("syntax error: unexpected TEXT, expecting LABEL or NEW_LINE"))
					Expect(res.Error.LineNum).To(Equal(1))
					Expect(res.Error.LinePos).To(Equal(6))
					Expect(res.Error.Token).To(Equal("quand "))
					close(done)
				})
			})
			It("should not parse with wrong action",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[1]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(true))
					Expect(res.Error.Message).To(Equal("syntax error: unexpected OBJECT_TYPE, expecting ACTION or TEXT"))
					Expect(res.Error.LineNum).To(Equal(1))
					Expect(res.Error.LinePos).To(Equal(26))
					Expect(res.Error.Token).To(Equal("campo "))
					close(done)
				})
			})
			It("should not parse with wrong object type",func(done Done){
				compiler.ParseBddTest(strings.TrimSpace(lines[2]),func(res compiler.ParsedTest) {
					Expect(res.HasError).To(Equal(true))
					Expect(res.Error.Message).To(Equal("syntax error: unexpected TEXT_PARAM, expecting OBJECT_TYPE or TEXT"))
					Expect(res.Error.LineNum).To(Equal(1))
					Expect(res.Error.LinePos).To(Equal(52))
					Expect(res.Error.Token).To(Equal("\"clovis\""))
					close(done)
				})
			})
		})
	})
})
