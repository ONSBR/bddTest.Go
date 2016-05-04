package compiler

import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
)



type ExecutionTestTree struct {
	NumScenarios int
	Feature lexer.Feature
	Error    lexer.ParserError
	HasError bool
}

var BuildExecutionTestTree = func(token string) ExecutionTestTree {
	log.Infof("Parsing token: %s", token)
	feature, retCode, err := ParseCode(token)
	
	if retCode == 0 {
		numScenarios := len(feature.Scenarios)
		transFeat, transRetCode, transErr := TranslateTokens(feature)
		if transRetCode == 0 {
			return ExecutionTestTree{NumScenarios: numScenarios, Feature: transFeat}
		} else {
			
			return ExecutionTestTree{NumScenarios:0,Error:lexer.ParserError{Message:transErr.Message,Token:transErr.Token},HasError:true}
		}
	} else {
		return ExecutionTestTree{NumScenarios:0,Error:err,HasError:true}
	}
}

