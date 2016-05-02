package compiler

// clovis
import (
	//	"fmt"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/util"
)

var log = util.GetLogger("compiler.parser")

type ParsedTest struct {
	NumLines int
	Lines    []lexer.Assertion
	Error    lexer.ParserError
	HasError bool
}

type OnBddTestParserCb func(ParsedTest)

func ParseBddTest(token string, cb OnBddTestParserCb) int {
	log.Infof("Parsing token: %s", token)
	return lexer.Test_ScenarioParse(&lexer.BddTestLex{S: token, OnBddTestParse: func(res *lexer.BddTestParseRes) {
		if res.HasError == false {
			lines := len(res.Lines.([]lexer.Assertion))
			log.Infof("First %s",res.Lines.([]lexer.Assertion)[0])
			cb(ParsedTest{NumLines: lines, Lines: res.Lines.([]lexer.Assertion)})
		} else {
			cb(ParsedTest{Error: res.Error, HasError:true})
		}
	}})
}
