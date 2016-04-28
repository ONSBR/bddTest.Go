package compiler

import (
//	"fmt"
	"github.com/ONSBR/bddTest.Go/lexer"
	"github.com/ONSBR/bddTest.Go/util"
)

var log = util.GetLogger("compiler.parser")
type ParsedTest struct {
	NumLines int
	Lines []lexer.Assertion
}

type OnBddTestParserCb func(ParsedTest)

func ParseBddTest(token string, cb OnBddTestParserCb) int {
	log.Infof("Parsing token: %s",token)
	return lexer.Test_ScenarioParse(&lexer.BddTestLex{S:token,OnBddTestParse:func(res *lexer.BddTestParseRes){
		lines := len(res.Lines.([]lexer.Assertion))
		cb(ParsedTest{NumLines:lines,Lines:res.Lines.([]lexer.Assertion)})
	}})
}

