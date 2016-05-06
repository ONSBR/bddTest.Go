package compiler

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
)

type (
	ExecutionTestTree struct {
		NumScenarios int
		Feature      lexer.Feature
		Error        lexer.ParserError
		HasError     bool
	}

	CodeCompiler struct{}

	CodeCompilerInterface interface {
		BuildExecutionTestTree(token string) ExecutionTestTree
	}
)

func (this *CodeCompiler) BuildExecutionTestTree(token string) ExecutionTestTree {
	codeParser := &CodeParser{}
	log.Infof("Parsing token: %s", token)
	feature, retCode, err := codeParser.ParseCode(token)

	if retCode == 0 {
		numScenarios := len(feature.Scenarios)
		transFeat, transRetCode, transErr := codeParser.TranslateTokens(feature)
		if transRetCode == 0 {
			return ExecutionTestTree{NumScenarios: numScenarios, Feature: transFeat}
		} else {

			return ExecutionTestTree{NumScenarios: 0, Error: lexer.ParserError{Message: transErr.Message, Token: transErr.Token}, HasError: true}
		}
	} else {
		return ExecutionTestTree{NumScenarios: 0, Error: err, HasError: true}
	}
}
