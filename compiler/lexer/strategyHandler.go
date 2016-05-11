package lexer

import (
    "reflect"
)

var (
    lexers = make(map[string]reflect.Type)
    tokenMaps = make(map[string]reflect.Type)
)

type (
    NextFunc func() byte
    BackFunc func()
    ResetBuffFunc func()
    
    //StrategyHandler is base type to strategy chooser
    StrategyHandler struct {}
    
    iStrategyHandler interface {
        GetLexer(string) ILexerLocalized
        GetTokenMap(string) ITokenMap
    }
    
    //ILexerLocalized lexer localized interface
    ILexerLocalized interface {
        LexLocalized(NextFunc, BackFunc, ResetBuffFunc) int
    }
    
    //ITokenMap token localized translation interface
    ITokenMap interface {
        GetTranslation(string) (string, bool)
    }
)

//GetLexer strategy chooser for lex parse
func (*StrategyHandler) GetLexer(lang string) ILexerLocalized {
    if lexerLocalizedType,ok := lexers[lang]; ok {
        return reflect.New(lexerLocalizedType).Elem().Interface().(ILexerLocalized)
    }
    return reflect.New(lexers["pt_BR"]).Elem().Interface().(ILexerLocalized)
}

//GetTokenMap strategy chooser for token translation
func (*StrategyHandler) GetTokenMap(lang string) ITokenMap  {
    if tokenMapType,ok := tokenMaps[lang]; ok {
        return reflect.New(tokenMapType).Elem().Interface().(ITokenMap)
    }
    return reflect.New(lexers["pt_BR"]).Elem().Interface().(ITokenMap)
}

//RegisterLexLocalization register a new strategy based on LexLocated type
func RegisterLexLocalization(locale string, typedNil ILexerLocalized) {
    t := reflect.TypeOf(typedNil)
    lexers[locale] = t
}

func RegisterTokenMap(locale string, typedNil ITokenMap) {
    t := reflect.TypeOf(typedNil)
    tokenMaps[locale] = t
}

//NewStrategyHandler basic constructor
func NewStrategyHandler() *StrategyHandler {
    return &StrategyHandler{}
}