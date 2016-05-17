package cli_test

import (
	. "github.com/ONSBR/bddTest.Go/cli"
	"github.com/ONSBR/bddTest.Go/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	"os"
	"strings"
)

var (
	logCliTest = util.GetLogger("test")
)
var _ = Describe("Cli", func() {
	Describe("validating specs", func ()  {
		BeforeEach(func ()  {
			prepareFiles()
		})
		
		AfterEach(func ()  {
			removeFiles()
		})
		
		It("should execute validation and return error", func (done Done)  {
			os.Args = []string{"cli.test", "-f", "../test/CLISpecs/specs1/error1.spec", "validate"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(-1))
			close(done)	
		})
		It("should execute validation and return ok", func (done Done)  {
			os.Args = []string{"cli.test", "-f", "../test/CLISpecs/specs1/good1.spec", "validate"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(0))
			close(done)	
		})
		It("should execute validation of multple files and return ok", func (done Done)  {
			os.Args = []string{"cli.test", "-m", "../test/CLISpecs/**/", "validate"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(-1))
			close(done)	
		})
	})
	Describe("generating yaml pages", func ()  {
		BeforeEach(func ()  {
			prepareFiles()
		})
		
		AfterEach(func ()  {
			removeFiles()
		})
		It("should execute yaml page generation and return ok without backup", func (done Done)  {
			filename := "../test/CLISpecs/specs1/good1.spec"
			os.Args = []string{"cli.test", "-f", filename, "yaml"}
			fileHandler := &util.FileHandler{}
			// pageContent,_ := fileHandler.ReadFile(filename + ".page")
			cli := NewCli()
			retCode := cli.Run() 
	
			Expect(retCode).To(Equal(0))
			Expect(fileHandler.DoesFileExists(filename+".page")).To(BeTrue())
			Expect(fileHandler.DoesFileExists(filename+".page.bkp")).To(BeFalse())
			// fileHandler.WriteFile(filename + ".page",pageContent)
			close(done)
		})
		It("should execute yaml page generation and return ok with backup", func (done Done)  {
			filename := "../test/CLISpecs/specs2/good1.spec"
			os.Args = []string{"cli.test", "-f", filename, "yaml", "-b"}
			fileHandler := &util.FileHandler{}
			// pageContent,_ := fileHandler.ReadFile(filename + ".page")
			cli := NewCli()
			retCode := cli.Run()
	
			Expect(retCode).To(Equal(0))
			Expect(fileHandler.DoesFileExists(filename+".page")).To(BeTrue())
			Expect(fileHandler.DoesFileExists(filename+".page.bkp")).To(BeTrue())
			// fileHandler.WriteFile(filename + ".page",pageContent)
			fileHandler.RemoveFile(filename+".page.bkp")
			close(done)
		}) 
		It("should execute yaml page generation for all found files and return ok without backup", func (done Done)  {
			filename := "../test/CLISpecs/**/"
			
// 						fileContent1 := `page: Home
// uri: test/teste1.html
// elements:
// - element: teste
//   locator: By.Id("teste")
//   type: button
// - element: teste1
//   locator: By.Id("teste1")
//   type: textbox
// - element: teste2
//   locator: By.Id("teste2")
//   type: selectbox
// `
//   			fileContent2 := `page: Home
// uri: test/teste2.html
// elements:
// - element: teste4
//   locator: By.Id("teste4")
//   type: textbox
// - element: teste5
//   locator: By.Id("teste5")
//   type: textbox
// - element: salvar
//   locator: By.Id("salvar")
//   type: button
// - element: teste6
//   locator: By.Id("teste6")
//   type: textbox
// ` 
  
			os.Args = []string{"cli.test", "-m", filename, "yaml"}
			fileHandler := &util.FileHandler{}
			
			cli := NewCli()
			retCode := cli.Run() 
	
			Expect(retCode).To(Equal(-1))
			
			files,_ := fileHandler.FindFiles(filename+"*.spec")
			
			for _,file := range files { 
				if (strings.LastIndex(file,"error1.spec") == -1) {
					Expect(fileHandler.DoesFileExists(file+".page")).To(BeTrue())
					Expect(fileHandler.DoesFileExists(file+".page.bkp")).To(BeFalse())
				} else {
					Expect(fileHandler.DoesFileExists(file+".page")).To(BeFalse())
				}
			}
			
			// fileHandler.WriteFile("../test/CLISpecs/specs3/teste1.spec.page",fileContent1)
			// fileHandler.WriteFile("../test/CLISpecs/specs3/teste2.spec.page",fileContent2)
			close(done)
		})
		It("should execute yaml page generation for all found files and return ok with backup", func (done Done)  {
			filename := "../test/CLISpecs/**/"
			os.Args = []string{"cli.test", "-m", filename, "yaml", "-b"}
			fileHandler := &util.FileHandler{}
			
// 			fileContent1 := `page: Home
// uri: test/teste1.html
// elements:
// - element: teste
//   locator: By.Id("teste")
//   type: button
// - element: teste1
//   locator: By.Id("teste1")
//   type: textbox
// - element: teste2
//   locator: By.Id("teste2")
//   type: selectbox
// `
//   			fileContent2 := `page: Home
// uri: test/teste2.html
// elements:
// - element: teste4
//   locator: By.Id("teste4")
//   type: textbox
// - element: teste5
//   locator: By.Id("teste5")
//   type: textbox
// - element: salvar
//   locator: By.Id("salvar")
//   type: button
// - element: teste6
//   locator: By.Id("teste6")
//   type: textbox
// ` 
  
			cli := NewCli()
			retCode := cli.Run()
	
			Expect(retCode).To(Equal(-1))
			
			files,_ := fileHandler.FindFiles(filename+"*.spec")
			
			for _,file := range files {
				
				if (strings.LastIndex(file,"error1.spec") == -1) {
					Expect(fileHandler.DoesFileExists(file+".page")).To(BeTrue())
					Expect(fileHandler.DoesFileExists(file+".page.bkp")).To(BeTrue())
					fileHandler.RemoveFile(file+".page.bkp")
				} else {
					Expect(fileHandler.DoesFileExists(file+".page")).To(BeFalse())
				}
			}
			
			// fileHandler.WriteFile("../test/CLISpecs/specs3/teste1.spec.page",fileContent1)
			// fileHandler.WriteFile("../test/CLISpecs/specs3/teste2.spec.page",fileContent2)
			close(done)
		}) 
	})
	Describe("generating Execution Tree", func ()  {
		BeforeEach(func ()  {
			prepareFiles()
		})
		
		AfterEach(func ()  {
			removeFiles()
		})
		It("should generate single execution tree and return ok", func (done Done)  {
			os.Args = []string{"cli.test", "-f", "../test/CLISpecs/specs3/teste1.spec", "run"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(0))
			close(done)
		})
		It("should generate a collection of execution tree and return ok", func (done Done)  {
			os.Args = []string{"cli.test", "-m", "../test/CLISpecs/**/", "run"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(0))
			close(done) 
		})
		It("should generate a collection of execution tree based on a guide script file and return ok", func (done Done)  {
			os.Args = []string{"cli.test", "-g", "../test/CLISpecs/guide.script", "run"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(0))
			close(done) 
		})
		It("should generate an error of parsing based on a guide script file and return ok", func (done Done)  {
			os.Args = []string{"cli.test", "-g", "../test/CLISpecs/guide2.script", "run"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(-1))
			close(done) 
		})
	})
})
