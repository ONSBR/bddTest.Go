package cli

import (
	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/processor"
	// "github.com/ONSBR/bddTest.Go/processor"
	"github.com/ONSBR/bddTest.Go/util"
)

var logCli = util.GetLogger("cli.main")

type (
	//Cli is the structure of Command Line Interface
	Cli struct{}

	//iCli exposed methods
	iCli interface {
		Run() int
	}
)

//Run is the main execution function of bddTest.Go
func (cli *Cli) Run() int {
	retCode := -1
	parseFlag := NewFlagParser()
	parseCode := parseFlag.Parse()
	util.InitLog(parseFlag.Options.Config)
	util.ReplaceLoggerDefaultFormatter()
	switch parseCode {
	case ErrCommand:
		usage := parseFlag.Usage()
		logCli.Infof("%s", usage)
		break
	case ValidadeCommand: //Run validation of files
		builder := compiler.NewBuilder()
		if parseFlag.Options.SpecFile != "" {
			tree := builder.BuildFile(parseFlag.Options.SpecFile)
			result := compiler.ExecutionTestTreeResult{Filename: parseFlag.Options.SpecFile, ExecutionTree: tree}
			trees := append([]compiler.ExecutionTestTreeResult{}, result)
			retCode = handleTreeResults(trees)
		} else if parseFlag.Options.Multi != "" {
			trees, err := builder.BuildFiles(parseFlag.Options.Multi)
			if err != nil {
				logCli.Errorf("%s", err.Error())
				retCode = -1
			} else {
				retCode = handleTreeResults(trees)
			}
		}
		break
	case YamlCommand: //Create .spec.page files
		builder := compiler.NewBuilder()
		if parseFlag.Options.SpecFile != "" {
			err := builder.BuildYamlPageObjectFile(parseFlag.Options.SpecFile, parseFlag.Options.Yaml.Backup)
			if err != nil {
				logCli.Errorf("%s", err.Error())
				retCode = -1
			} else {
				retCode = 0
			}
		} else if parseFlag.Options.Multi != "" {
			err := builder.BuildYamlPageObjectFiles(parseFlag.Options.Multi, parseFlag.Options.Yaml.Backup)
			if err != nil {
				logCli.Errorf("%s", err.Error())
				retCode = -1
			} else {
				retCode = 0
			}
		}
		break
	case RunCommand: //Execute known tests
		retCode = -1
		builder := compiler.NewBuilder()
		executions := []compiler.Execution{}
		if parseFlag.Options.SpecFile != "" {
			execution := builder.BuildExecution(parseFlag.Options.SpecFile, parseFlag.Options.BaseURI)
			executions = append(executions, execution)
			retCode = 0
		} else if parseFlag.Options.Multi != "" {
			executions, _ = builder.BuildExecutions(parseFlag.Options.Multi, parseFlag.Options.BaseURI)
			retCode = 0
		} else if parseFlag.Options.Guide != "" {
			var err error
			executions, err = builder.BuildExecutionsFromGuide(parseFlag.Options.Guide, parseFlag.Options.BaseURI)
			logCli.Errorf("Executions %d %s", len(executions), err)
			if err != nil {
				retCode = -1
			} else {
				retCode = 0
			}
		}

		t := processor.NewTestProcessor(&executions[0])
		t.Process()

		break
	}
	return retCode
}

//NewCli Constructor of cli
func NewCli() *Cli {
	cli := &Cli{}
	return cli
}

func handleTreeResults(results []compiler.ExecutionTestTreeResult) int {
	retCode := 0
	for _, executionTestTreeResult := range results {
		filename := executionTestTreeResult.Filename
		tree := executionTestTreeResult.ExecutionTree
		if tree.HasError {
			logCli.Errorf("%s\n", filename)
			logCli.Errorf("%s Line: %d Pos: %d Token: %s\n", tree.Error.Message, tree.Error.LineNum, tree.Error.LinePos, tree.Error.Token)
			retCode = -1
		} else {
			logCli.Infof("%s is valid!\n", filename)
		}
	}
	return retCode
}
