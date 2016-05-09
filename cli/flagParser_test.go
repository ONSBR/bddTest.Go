package cli_test

import (
	. "github.com/ONSBR/bddTest.Go/cli"
	"github.com/ONSBR/bddTest.Go/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/voxelbrain/goptions"
	
	"strings"
	"os"
)

var (
	// usageArr []string 
	// usageStr string
	logFP = util.GetLogger("test")
)

var _ = Describe("FlagParser", func() {
	Describe("base construction", func(){
		BeforeEach(func ()  {
			prepareFiles()
		})
		
		AfterEach(func ()  {
			removeFiles()
		})
		It("should return usage string", func (done Done)  {
			var flagParser *FlagParser = NewFlagParser()
			usage := flagParser.Usage()
			usageSplit := strings.Split(usage, "\n")
			Expect(len(usageSplit)).To(Equal(16))
			close(done)
		})
		It("should parse and return invalid command", func (done Done)  {
			os.Args = []string{"cli.test", "error"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(-1))
			close(done)
		})
		It("should parse and return invalid command", func (done Done)  {
			os.Args = []string{"cli.test", "-f", "../test/specs/teste1.spec"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(-1))
			close(done)
		})
		It("should parse and return invalid command", func (done Done)  {
			os.Args = []string{"cli.test", "error"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(ErrCommand))
			close(done)
		})
		It("should parse and return valid validate command", func (done Done)  {
			os.Args = []string{"cli.test", "validate"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(ValidadeCommand))
			close(done)
		})
		It("should parse and return valid yaml command", func (done Done)  {
			os.Args = []string{"cli.test", "yaml"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(YamlCommand))
			Expect(flagParser.Options.Yaml.Backup).To(BeFalse())
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("yaml")))
			close(done)
		})
		It("should parse and return valid yaml command with backup long", func (done Done)  {
			os.Args = []string{"cli.test", "yaml", "--backup"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(YamlCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("yaml")))
			Expect(flagParser.Options.Yaml.Backup).To(BeTrue())
			close(done)
		})
		It("should parse and return valid yaml command with backup short", func (done Done)  {
			os.Args = []string{"cli.test", "yaml", "-b"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(YamlCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("yaml")))
			Expect(flagParser.Options.Yaml.Backup).To(BeTrue())
			close(done)
		})
		It("should parse and return valid run command", func (done Done)  {
			os.Args = []string{"cli.test", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			close(done)
		})
		It("should parse and return valid run command with file spec flag short", func (done Done)  {
			os.Args = []string{"cli.test", "-f", "../test/specs/teste1.spec", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			Expect(flagParser.Options.SpecFile).To(Equal("../test/specs/teste1.spec"))
			close(done) 
		})
		It("should parse and return valid run command with file spec flag long", func (done Done)  {
			os.Args = []string{"cli.test", "--specfile", "../test/specs/teste1.spec", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			Expect(flagParser.Options.SpecFile).To(Equal("../test/specs/teste1.spec"))
			close(done) 
		})
		It("should parse and return valid run command with file pattern flag short", func (done Done)  {
			os.Args = []string{"cli.test", "-m", "../test/**", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			Expect(flagParser.Options.Multi).To(Equal("../test/**"))
			close(done) 
		})
		It("should parse and return valid run command with file pattern flag long", func (done Done)  {
			os.Args = []string{"cli.test", "--multi", "../test/**", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			Expect(flagParser.Options.Multi).To(Equal("../test/**"))
			close(done) 
		})
		It("should parse and return valid run command with base uri flag short", func (done Done)  {
			os.Args = []string{"cli.test", "-u", "http://10.0.0.1:8000", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			Expect(flagParser.Options.BaseURI).To(Equal("http://10.0.0.1:8000"))
			close(done) 
		})
		It("should parse and return valid run command with base uri flag long", func (done Done)  {
			os.Args = []string{"cli.test", "--uri", "http://10.0.0.1:8000", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			Expect(flagParser.Options.BaseURI).To(Equal("http://10.0.0.1:8000"))
			close(done) 
		})
		It("should parse and return valid run command with selenium server flag short", func (done Done)  {
			os.Args = []string{"cli.test", "-s", "http://10.0.0.1:8000/hub", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			Expect(flagParser.Options.SeleniumServer).To(Equal("http://10.0.0.1:8000/hub"))
			close(done) 
		})
		It("should parse and return valid run command with selenium server flag short", func (done Done)  {
			os.Args = []string{"cli.test", "--seleniumserver", "http://10.0.0.1:8000/hub", "run"}
			flagParser := NewFlagParser()
			commandOk := flagParser.Parse()
			Expect(commandOk).To(Equal(RunCommand))
			Expect(flagParser.Options.Verb).To(Equal(goptions.Verbs("run")))
			Expect(flagParser.Options.SeleniumServer).To(Equal("http://10.0.0.1:8000/hub"))
			close(done) 
		})
	})
})
