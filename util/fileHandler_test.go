package util_test

import (
	. "github.com/ONSBR/bddTest.Go/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"
)

var logFH = GetLogger("test")

var _ = Describe("FileHandler", func() {
	Describe("handling directory patterns", func() {
		It("should list files of a path", func(done Done) {
			fileHandler := &FileHandler{}
			files, err := fileHandler.FindFiles("../test/specs/")
			Expect(err).To(BeNil())
			Expect(len(files)).To(Equal(3))
			Expect(files[0]).To(Equal("../test/specs/teste1.spec"))
			Expect(files[1]).To(Equal("../test/specs/teste2.spec"))
			Expect(files[2]).To(Equal("../test/specs/teste3.spec"))
			close(done)
		})
		It("should list files of a path pattern", func(done Done) {
			fileHandler := &FileHandler{}
			files, err := fileHandler.FindFiles("../test/**/")
			Expect(err).To(BeNil())
			Expect(len(files)).To(Equal(5))
			Expect(files[0]).To(Equal("../test/specifications/1.spec"))
			Expect(files[1]).To(Equal("../test/specifications/2.spec"))
			Expect(files[2]).To(Equal("../test/specs/teste1.spec"))
			Expect(files[3]).To(Equal("../test/specs/teste2.spec"))
			Expect(files[4]).To(Equal("../test/specs/teste3.spec"))
			close(done)
		})
	})
	Describe("reading files", func() {
		It("should return file exists", func(done Done) {
			fileExist := "../test/specs/teste1.spec"
			fileNotExist := "../test/specs/teste1.spec1"

			fileHandler := &FileHandler{}
			exists := fileHandler.DoesFileExists(fileExist)
			Expect(exists).To(BeTrue())
			exists = fileHandler.DoesFileExists(fileNotExist)
			Expect(exists).To(BeFalse())
			close(done)
		})
		It("should read a single file", func(done Done) {
			ret := `Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"`
			fileHandler := &FileHandler{}
			content, err := fileHandler.ReadFile("../test/specs/teste1.spec")
			Expect(err).To(BeNil())
			Expect(content).To(Equal(ret))
			close(done)
		})
		It("should read all files of a pattern path and return a map", func(done Done) {
			ret := `Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"`
			fileHandler := &FileHandler{}
			files, err := fileHandler.ReadFiles("../test/**/")
			Expect(err).To(BeNil())
			Expect(len(files)).To(Equal(5))
			file := files["../test/specs/teste1.spec"]
			Expect(file.Filename).To(Equal("teste1.spec"))
			Expect(file.Path).To(Equal("../test/specs"))
			Expect(file.Content).To(Equal(ret))
			close(done)
		})
	})
	Describe("writing files", func() {
		It("should write content to path/file", func(done Done) {
			content := `Aspecto: Este é um aspecto
Pagina: Cadastro de Clientes
Cenario: primeiro cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao "clovis3"`
			fileHandler := &FileHandler{}
			err := fileHandler.WriteFile("../test/specs/teste4.spec", content)
			Expect(err).To(BeNil())

			bytes, errR := ioutil.ReadFile("../test/specs/teste4.spec")
			Expect(errR).To(BeNil())
			str := string(bytes)
			Expect(str).To(Equal(content))

			os.Remove("../test/specs/teste4.spec")

			close(done)
		})
	})
})
