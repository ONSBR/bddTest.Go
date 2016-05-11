package lexer

type ParserError struct {
	Message string
	LineNum int
	LinePos int
	Token string
}

type BddTestParseRes struct {
	Error ParserError
	HasError bool
	Feature Feature
}


type OnBddTestParseCb func(*BddTestParseRes)

type BddTestLex struct {
	S string
	pos int
	lineNum int
	linePos int
	OnBddTestParse OnBddTestParseCb
	Itemchan chan *BddTestParseRes
	buf []byte
	ParseError ParserError
	lexStrategy ILexerLocalized
} 
