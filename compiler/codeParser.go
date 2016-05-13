package compiler

// clovis
import (
	"github.com/ONSBR/bddTest.Go/compiler/lexer"
	"github.com/ONSBR/bddTest.Go/util"
	
	"strings"
)

var (
	log = util.GetLogger("compiler.parser")
	strategyHandler = lexer.NewStrategyHandler()
	locale string
)

type (
	//TranslatorError holds translation error
	TranslatorError struct {
		Message string
		Token   string
	}

	//CodeParser basic struct for parser
	CodeParser struct{}

	iCodeParser interface {
		ParseCode(token string) (feature lexer.Feature, retCode int, err lexer.ParserError)
		TranslateTokens(feat lexer.Feature) (feature lexer.Feature, retCode int, err TranslatorError)
	}
)

//ParseCode receives a token string and try to parse it, respecting localization strategy
func (*CodeParser) ParseCode(token string) (feature lexer.Feature, retCode int, err lexer.ParserError) {
	log.Infof("Parsing token: %s", token)
	
	
	localeFound, preparedToken := getLocale(token)
	locale = localeFound
	strategy := strategyHandler.GetLexer(locale)
	
	bddTestLex := lexer.BddTestLex{S: preparedToken, OnBddTestParse: nil}
	bddTestLex.SetLexStrategy(strategy)
	
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

//TranslateTokens translate all tokens found in Feature labels to common names, respecting localization strategy
func (*CodeParser) TranslateTokens(feat lexer.Feature) (feature lexer.Feature, retCode int, err TranslatorError) {
	strategy := strategyHandler.GetTokenMap(locale)
	
	feature = feat
	fLab, ok := strategy.GetTranslation(feature.Label)
	if !ok {
		retCode = 1
		err = TranslatorError{Message: "Invalid feature token translation", Token: feature.Label}
		return
	}
	feature.Label = fLab
	for idx, scenario := range feature.Scenarios {
		sLab, ok := strategy.GetTranslation(scenario.Label)

		if !ok {
			retCode = 1
			err = TranslatorError{Message: "Invalid scenario token translation", Token: scenario.Label}
			return
		}
		feature.Scenarios[idx].Label = sLab
		for idx1, action := range scenario.Actions {
			aLab, ok := strategy.GetTranslation(action.Label)
			if !ok {
				retCode = 1
				err = TranslatorError{Message: "Invalid action label token translation", Token: action.Label}
				return
			}
			feature.Scenarios[idx].Actions[idx1].Label = aLab
			aAct, ok := strategy.GetTranslation(action.Action)
			if !ok {
				retCode = 1
				err = TranslatorError{Message: "Invalid action token translation", Token: action.Action}
				return
			}
			feature.Scenarios[idx].Actions[idx1].Action = aAct
			aObjT, ok := strategy.GetTranslation(action.ObjectType)
			if !ok {
				retCode = 1
				err = TranslatorError{Message: "Invalid action object type token translation", Token: action.ObjectType}
				return
			}
			feature.Scenarios[idx].Actions[idx1].ObjectType = aObjT
		}

		for idx2, expect := range scenario.Expectations { 
			aLab, ok := strategy.GetTranslation(expect.Label)
			if !ok { 
				retCode = 1
				err = TranslatorError{Message: "Invalid expectation label token translation", Token: expect.Label}
				return
			}
			feature.Scenarios[idx].Expectations[idx2].Label = aLab
			aAct, ok := strategy.GetTranslation(expect.Action)
			if !ok {
				retCode = 1
				err = TranslatorError{Message: "Invalid expectation token translation", Token: expect.Action}
				return
			}
			feature.Scenarios[idx].Expectations[idx2].Action = aAct
			aObjT, ok := strategy.GetTranslation(expect.ObjectType)
			if !ok {
				retCode = 1
				err = TranslatorError{Message: "Invalid expectation object type token translation", Token: expect.ObjectType}
				return
			}
			feature.Scenarios[idx].Expectations[idx2].ObjectType = aObjT
			aMat, ok := strategy.GetTranslation(expect.Matcher)
			if !ok {
				retCode = 1
				err = TranslatorError{Message: "Invalid matcher token translation", Token: expect.Matcher}
				return
			}
			feature.Scenarios[idx].Expectations[idx2].Matcher = aMat
		}
	}
	retCode = 0
	return
}

func getLocale(token string) (locale string, preparedToken string) {
	locale = "pt_BR"
	preparedToken = token
	lines := strings.Split(token, "\n")
	line := lines[0]
	if (string(line[0]) == "#") {
		locale = line[1:]
		lines = lines[1:]
		preparedToken = strings.Join(lines, "\n")
	} 
	return 
}
