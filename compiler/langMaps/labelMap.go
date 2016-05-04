package bddTestLangMaps

import (

)

var (
	ExpectActionTokens = map[string]string{
		"Quando":"When",
		"E":"And",
	}
	ExpectExpressionTokens = map[string]string{
		"Entao":"Then",
	}
	InitScenarioTokens = map[string]string{
		"Dado":"Given",
		"usuario":"User",
	}
	PageTokens = map[string]string{
		"Pagina":"Page",
	}
	ScenarioTokens = map[string]string{
		"Cenario":"Scenario",
	}
	FeatureTokens = map[string]string{
		"Aspecto":"Feature",
	}
)