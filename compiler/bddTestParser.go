package compiler

// clovis
import (
	//	"fmt"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/util"
)

var (
	log = util.GetLogger("compiler.parser")
)

type ParsedTest struct {
	NumScenarios int
	Feature lexer.Feature
	Error    lexer.ParserError
	HasError bool
}

type OnBddTestParserCb func(ParsedTest)

func ParseCode(token string) (feature lexer.Feature, retCode int, err lexer.ParserError) {
	log.Infof("Parsing token: %s", token)
	bddTestLex := lexer.BddTestLex{S: token, OnBddTestParse: nil}
	retCode = lexer.FeatureParse(&bddTestLex)
	if retCode == 0 {
		feature = lexer.GetParsedCode().(lexer.Feature)
		err = lexer.ParserError{}
	} else {
		feature = lexer.Feature{}
		err = bddTestLex.ParseError
	}
	return
}

//func ParseCode(token string) ParsedTest {
//	log.Infof("Parsing token: %s", token)
//	bddTestLex := lexer.BddTestLex{S: token, OnBddTestParse: nil}
//	var retCode = lexer.FeatureParse(&bddTestLex)
//	if retCode == 0 {
//		res := lexer.GetParsedCode().(lexer.Feature)
//		numScenarios := len(res.Scenarios)
//		return ParsedTest{NumScenarios: numScenarios, Feature: res}
//	} else {
//		return ParsedTest{Error:bddTestLex.ParseError, HasError:true}
//	}
//}

func TranslateTokens(feature lexer.Feature) {
	
}

