package lexer

type BddTestParseRes struct {
	Error string
	Lines interface{}
}

type OnBddTestParseCb func(*BddTestParseRes)

type BddTestLex struct {
	S string
	pos int
	OnBddTestParse OnBddTestParseCb
	Itemchan chan *BddTestParseRes
	buf []byte
}
