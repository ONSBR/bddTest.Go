package compiler

// clovis
import (
	//	"fmt"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/util"
)

var log = util.GetLogger("compiler.parser")

type ParsedTest struct {
	NumScenarios int
	Feature lexer.Feature
	Error    lexer.ParserError
	HasError bool
}

type OnBddTestParserCb func(ParsedTest)

func ParseBddTest(token string) ParsedTest {
	log.Infof("Parsing token: %s", token)
	bddTestLex := lexer.BddTestLex{S: token, OnBddTestParse: nil}
	var retCode = lexer.FeatureParse(&bddTestLex)
	if retCode == 0 {
		res := lexer.GetParsedCode().(lexer.Feature)
		numScenarios := len(res.Scenarios)
		return ParsedTest{NumScenarios: numScenarios, Feature: res}
	} else {
		return ParsedTest{Error:bddTestLex.ParseError, HasError:true}
	}
//	log.Infof("Return code: %d", retCode)
////	
//	log.Infof("Err %s",bddTestLex.ParseError)
//	return retCode
////	return lexer.FeatureParse(&lexer.BddTestLex{S: token, OnBddTestParse: func(res *lexer.BddTestParseRes) {
////		if res.HasError == false {
////			numScenarios := len(res.Feature.Scenarios)
////			cb(ParsedTest{NumScenarios: numScenarios, Feature:res.Feature})
////		} else {
////			cb(ParsedTest{Error: res.Error, HasError:true})
////		}
////	}})
}
