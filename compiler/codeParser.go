package compiler

// clovis
import (
	. "github.com/ONSBR/bddTest.Go/compiler/langMaps"
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/util"
)

var (
	log = util.GetLogger("compiler.parser")
)

type TranslatorError struct {
	Message string
	Token string
}

var ParseCode = func(token string) (feature lexer.Feature, retCode int, err lexer.ParserError) {
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

var TranslateTokens = func(feat lexer.Feature) (feature lexer.Feature, retCode int, err TranslatorError) {
	feature = feat
	fLab,ok := FeatureTokens[feature.Label]
	if !ok {
		retCode = 1
		err = TranslatorError{Message:"Invalid feature token translation",Token:feature.Label}
		return
	}
	feature.Label = fLab
	for idx,scenario := range feature.Scenarios {
		_ = idx
		sLab,ok := ScenarioTokens[scenario.Label]
		if !ok {
			retCode = 1
			err = TranslatorError{Message:"Invalid scenario token translation",Token:scenario.Label}
			return
		}
		scenario.Label = sLab
		for idx1, action := range scenario.Actions {
			_ = idx1
			aLab, ok := ExpectActionTokens[action.Label]
			if !ok {
				retCode = 1
				err = TranslatorError{Message:"Invalid action label token translation",Token:action.Label}
				return
			}
			action.Label = aLab
			aAct, ok := ActionTokens[action.Action]
			if !ok {
				retCode = 1
				err = TranslatorError{Message:"Invalid action token translation",Token:action.Action}
				return
			}
			action.Action = aAct
			aObjT, ok := ObjectTokens[action.ObjectType]
			if !ok {
				retCode = 1
				err = TranslatorError{Message:"Invalid action object type token translation",Token:action.ObjectType}
				return
			}
			action.ObjectType = aObjT
		}
		
		for idx2, expect := range scenario.Expectations {
			_ = idx2
			aLab, ok := ExpectExpressionTokens[expect.Label]
			if !ok {
				retCode = 1
				err = TranslatorError{Message:"Invalid expectation label token translation",Token:expect.Label}
				return
			}
			expect.Label = aLab
			aAct, ok := ActionTokens[expect.Action]
			if !ok {
				retCode = 1
				err = TranslatorError{Message:"Invalid expectation token translation",Token:expect.Action}
				return
			}
			expect.Action = aAct
			aObjT, ok := ObjectTokens[expect.ObjectType]
			if !ok {
				retCode = 1
				err = TranslatorError{Message:"Invalid expectation object type token translation",Token:expect.ObjectType}
				return
			}
			expect.ObjectType = aObjT
		}
	}
	retCode = 0
	return
}



