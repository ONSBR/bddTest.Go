package lexer

import (
//	"fmt"
	"strconv"
	"strings"
	"github.com/ONSBR/bddTest.Go/util"
)

const (
	C_Feature = iota
	C_Scenario
	C_Assertion
)

var _emptystruct struct{} = struct{}{}

var logUtil = util.GetLogger("lexer.lexer") 

func (this *BddTestLex) peek() (bret byte) {
	bret = this.next()
	this.pos -= 1
	this.linePos -= 1
	if this.linePos < 0 {
		this.linePos = 0
//		this.lineNum -= 1
	}
	return
}

func (this *BddTestLex) next() (bret byte) {

	if this.pos < len(this.S) {
		bret = byte(this.S[this.pos])
		this.buf = append(this.buf, bret)
	} else {
		bret = 0
	}
	this.pos += 1
	this.linePos += 1
	return
}

func (this *BddTestLex) back() {
	if 0 < this.pos {
		this.pos -= 1
		this.linePos -= 1
		if this.linePos < 0 {
			this.linePos = 0
//			this.lineNum -= 1
		}
	}
	return
}

func (this *BddTestLex) newLine() {
	this.lineNum += 1
	this.linePos = 0
}

func (this BddTestLex) data() (bb []byte) {
	if this.pos < len(this.S) {
		bb = this.buf[:len(this.buf)-1]
	} else {
		bb = this.buf
	}
	return
}

func (this *BddTestLex) Error(s string) {
	err := ParserError{Message:s,LineNum:this.lineNum+1,LinePos:this.linePos+1,Token:string(this.buf)}
	this.ParseError = err
//	res := &BddTestParseRes{Error:err, HasError:true}
//	if nil != this.OnBddTestParse {
//		this.OnBddTestParse(res)
//	}
}

func (this BddTestLex) getInt() (n int, err error) {
	s := string(this.data())
	if n, err = strconv.Atoi(strings.TrimSpace(s)); nil != err {
		this.Error(err.Error() + ";s=" + s)
	}
	return
}

func (this BddTestLex) getStr() (s string) {
	s = strings.TrimSpace(string(this.data()))
	return
}

func (this BddTestLex) getLabel() (s string) {
	s = strings.TrimSpace(string(this.data()))
	c := string(s[len(s)-1:])
	if ":" == c {
		s = s[:len(s)-1]
	}
	return
}

func (this BddTestLex) getParam() (s string) {
	str := strings.TrimSpace(string(this.data()))
	s = str[1:len(str)-1]
	return
}

func (this *BddTestLex) SetLexStrategy(strategy ILexerLocalized) {
	this.lexStrategy = strategy
}

//Lex is the wrapper to localized strategy Lex
func (this *BddTestLex) Lex(lval *FeatureSymType) (ret int) {
	var err error
	
	retLex := this.lexStrategy.LexLocalized(this.next, this.back, this.resetBuffer)
	switch retLex {
		case NEW_LINE:
			this.newLine()
			break
		case NUMBER:
			if lval.item, err = this.getInt(); nil!=err {
				retLex = -1
			}
			break
		case INIT_SCENARIO_LABEL:
			lval.item = this.getLabel()
			break
		case PAGE_LABEL:
			lval.item = this.getLabel()
			break
		case USER_SCENARIO_LABEL:
			lval.item = this.getLabel()
			break
		case EXPECT_ACTION_LABEL:
			lval.item = this.getLabel()
			break
		case EXPECT_ACTION_ACTION:
			lval.item = this.getStr()
			break
		case SCENARIO_LABEL:
			lval.item = this.getLabel()
			break
		case FEATURE_LABEL:
			lval.item = this.getLabel()
			break
		case LABEL:
			lval.item = this.getLabel()
			break
		case ACTION:
			lval.item = this.getStr()
			break
		case OBJECT_TYPE:
			lval.item = this.getStr()
			break
		case SIGN:
			lval.item = this.getStr()
			break
		case TEXT_PARAM:
			lval.item = this.getParam()
			break
		case TEXT:
			lval.item = this.getStr()
			break
	}
	ret = retLex
	return
}

func (this *BddTestLex) resetBuffer()  {
	if nil!=this.buf {
		this.buf = this.buf[len(this.buf)-1:]
	}
}

