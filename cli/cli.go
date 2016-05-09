package cli

import (
    "github.com/ONSBR/bddTest.Go/compiler"
    "github.com/ONSBR/bddTest.Go/util"
)

var logCli = util.GetLogger("cli.main")

type(
    //Cli is the structure of Command Line Interface
    Cli struct {}
    
    //iCli exposed methods
    iCli interface {
        Run() int
    }
)

//Run is the main execution function of bddTest.Go
func (cli *Cli) Run() int  {
    util.ReplaceLoggerDefaultFormatter()
    retCode := -1
    parseFlag := NewFlagParser()
    parseCode := parseFlag.Parse()
    switch parseCode {
        case ErrCommand:
            parseFlag.Usage()
            return -1
        case ValidadeCommand: //Run validation of files
            builder := &compiler.Builder{}
            if parseFlag.Options.SpecFile != "" {
                tree := builder.BuildFile(parseFlag.Options.SpecFile)
                result := compiler.ExecutionTestTreeResult{Filename: parseFlag.Options.SpecFile, ExecutionTree: tree}
		        trees := append([]compiler.ExecutionTestTreeResult{}, result)
                retCode = handleTreeResults(trees)
            } else if parseFlag.Options.Multi != "" {
                trees, err := builder.BuildFiles(parseFlag.Options.Multi)
                if err != nil {
                    logCli.Errorf("%s",err.Error())
                    retCode = -1 
                } else {
                    retCode = handleTreeResults(trees)
                }
            }
            break
        case YamlCommand: //Create .spec.page files
            break
        case RunCommand: //Execute known tests
            break
    }
    return retCode 
}

//NewCli Constructor of cli
func NewCli()  *Cli{
    cli := &Cli{}
    return cli
}

func handleTreeResults(results []compiler.ExecutionTestTreeResult) int {
    retCode := 0
    for _,executionTestTreeResult := range results {
        filename := executionTestTreeResult.Filename
        tree := executionTestTreeResult.ExecutionTree
        if (tree.HasError) {
            logCli.Errorf("%s\n", filename)
            logCli.Errorf("%s Line: %d Pos: %d Token: %s\n", tree.Error.Message,tree.Error.LineNum,tree.Error.LinePos,tree.Error.Token)
            retCode = -1
        } else {
            logCli.Infof("%s is valid!\n", filename)
        }
    }
    return retCode
}