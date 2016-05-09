package cli_test

import (
	. "github.com/ONSBR/bddTest.Go/cli"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	
	"os"
)

var (
	
)
var _ = Describe("Cli", func() {
	Describe("validating specs", func ()  {
		It("should execute validation and return error", func (done Done)  {
			os.Args = []string{"cli.test", "-f", "../CLISpecs/specs1/error1.spec", "validate"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(-1))
			close(done)	
		})
		It("should execute validation and return ok", func (done Done)  {
			os.Args = []string{"cli.test", "-f", "../CLISpecs/specs1/good1.spec", "validate"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(0))
			close(done)	
		})
		It("should execute validation of multple files and return ok", func (done Done)  {
			os.Args = []string{"cli.test", "-m", "../CLISpecs/**/", "validate"}
			cli := NewCli()
			retCode := cli.Run()
			
			Expect(retCode).To(Equal(-1))
			close(done)	
		})
	})
})
