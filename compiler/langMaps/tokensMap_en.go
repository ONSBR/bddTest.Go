package langMaps

import (
    . "github.com/ONSBR/bddTest.Go/compiler/lexer"
)

//TokenMapEn strategy implementation of lex pt_BR
type TokenMapEn struct {}

var tokenMapsEn = map[string]string{
    "button":"button",
    "textbox":"textbox",
    "selectbox":"selectbox",
    "session":"session",
    "When":"When",
    "And":"And",
    "Then":"Then",
    "Given":"Given",
    "User":"User",
    "Page":"Page",
    "Scenario":"Scenario",
    "Feature":"Feature",
    "click":"click",
    "select":"select",
    "set":"set",
    "exists":"exists",
    "expect":"expect",
    "equal":"eq",
    "greater then":"gt",
    "less then":"lt",
    "greater then or equal":"gte",
    "less then or equal":"lte",
    "contains":"ct",
}

func init() {
    RegisterTokenMap("en",(*TokenMapEn)(nil))
}

//GetTranslation finds a translation token fo supplied token
func (*TokenMapEn) GetTranslation(token string) (translatedToken string, err bool) {
    translatedToken, err = tokenMapsEn[token]
    return
}