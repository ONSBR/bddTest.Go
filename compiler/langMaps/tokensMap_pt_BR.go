package langMaps

import (
    . "github.com/ONSBR/bddTest.Go/compiler/lexer"
)

//TokenMapPtBR strategy implementation of lex pt_BR
type TokenMapPtBR struct {}

var tokenMapsPtBR = map[string]string{
    "botao":"button",
    "campo":"textbox",
    "lista":"selectbox",
    "sessao":"session",
    "Quando":"When",
    "E":"And",
    "Entao":"Then",
    "Dado":"Given",
    "usuario":"User",
    "Pagina":"Page",
    "Cenario":"Scenario",
    "Aspecto":"Feature",
    "clico":"click",
    "seleciono":"select",
    "preencho":"set",
    "existe":"exists",
    "espero":"expect",
    "igual a":"eq",
    "maior que":"gt",
    "menor que":"lt",
    "maior ou igual a":"gte",
    "menor ou igual a":"lte",
    "contem":"ct",
}

func init() {
    RegisterTokenMap("pt_BR",(*TokenMapPtBR)(nil))
}

//GetTranslation finds a translation token fo supplied token
func (*TokenMapPtBR) GetTranslation(token string) (translatedToken string, err bool) {
    translatedToken, err = tokenMapsPtBR[token]
    return
}