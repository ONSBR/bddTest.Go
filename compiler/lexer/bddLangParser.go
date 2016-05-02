// CAUTION: Generated file - DO NOT EDIT.

//
package lexer

import (
	//"fmt"
	"github.com/ONSBR/bddTest.Go/util"
)

var logP = util.GetLogger("lexer.parser")

func (this *BddTestLex) Lex(lval *Test_ScenarioSymType) (ret int) { // begin main()

	var err error
	var c byte = ' '

	defer func() {
		this.back()
	}()

yystate0:

	if nil != this.buf {
		this.buf = this.buf[len(this.buf)-1:]
	}

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = this.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate5
	case c == '-':
		goto yystate8
	case c == 'D' || c == 'd':
		goto yystate11
	case c == 'E' || c == 'e':
		goto yystate15
	case c == 'Q' || c == 'q':
		goto yystate18
	case c == '\n':
		goto yystate4
	case c == '\t' || c == ' ':
		goto yystate3
	case c == '\x00':
		goto yystate2
	case c == 'b':
		goto yystate21
	case c == 'c':
		goto yystate26
	case c == 'l':
		goto yystate33
	case c == 'p':
		goto yystate37
	case c == 's':
		goto yystate43
	case c >= '0' && c <= '9':
		goto yystate9
	case c >= 'A' && c <= 'C' || c >= 'F' && c <= 'P' || c >= 'R' && c <= 'Z' || c == 'a' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c == 'r' || c >= 't' && c <= 'z':
		goto yystate10
	}

yystate2:
	c = this.next()
	goto yyrule8

yystate3:
	c = this.next()
	switch {
	default:
		goto yyrule10
	case c == '\t' || c == ' ':
		goto yystate3
	}

yystate4:
	c = this.next()
	switch {
	default:
		goto yyrule1
	case c == '\n':
		goto yystate4
	}

yystate5:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate6
	}

yystate6:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate7
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate6
	}

yystate7:
	c = this.next()
	switch {
	default:
		goto yyrule7
	case c == '"':
		goto yystate7
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate6
	}

yystate8:
	c = this.next()
	goto yyrule6

yystate9:
	c = this.next()
	switch {
	default:
		goto yyrule2
	case c >= '0' && c <= '9':
		goto yystate9
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate10:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate11:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'a':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate12:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'd':
		goto yystate13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate10
	}

yystate13:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'o':
		goto yystate14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate10
	}

yystate14:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate15:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c == 'n':
		goto yystate16
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate16:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 't':
		goto yystate17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate17:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'a':
		goto yystate13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate18:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'u':
		goto yystate19
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate10
	}

yystate19:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'a':
		goto yystate20
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate20:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'n':
		goto yystate12
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate21:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'o':
		goto yystate22
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate10
	}

yystate22:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 't':
		goto yystate23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate23:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'a':
		goto yystate24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate24:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'o':
		goto yystate25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate10
	}

yystate25:
	c = this.next()
	switch {
	default:
		goto yyrule5
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate26:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'a':
		goto yystate27
	case c == 'l':
		goto yystate29
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate10
	}

yystate27:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'm':
		goto yystate28
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate10
	}

yystate28:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'p':
		goto yystate24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate10
	}

yystate29:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'i':
		goto yystate30
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate30:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'c':
		goto yystate31
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate10
	}

yystate31:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'o':
		goto yystate32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate10
	}

yystate32:
	c = this.next()
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate10
	}

yystate33:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'i':
		goto yystate34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate34:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 's':
		goto yystate35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate10
	}

yystate35:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 't':
		goto yystate36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate10
	}

yystate36:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'a':
		goto yystate25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'b' && c <= 'z':
		goto yystate10
	}

yystate37:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'r':
		goto yystate38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate10
	}

yystate38:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'e':
		goto yystate39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate39:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'e':
		goto yystate40
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate40:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'n':
		goto yystate41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yystate41:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'c':
		goto yystate42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate10
	}

yystate42:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'h':
		goto yystate31
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate10
	}

yystate43:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'e':
		goto yystate44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate44:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'l':
		goto yystate45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate10
	}

yystate45:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'e':
		goto yystate46
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate10
	}

yystate46:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'c':
		goto yystate47
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate10
	}

yystate47:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'i':
		goto yystate48
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate10
	}

yystate48:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'o':
		goto yystate49
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate10
	}

yystate49:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == 'n':
		goto yystate31
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate10
	}

yyrule1: // [\n]+
	{
		logP.Infof("returning newline")
		this.newLine() //; return NEW_LINE
		goto yystate0
	}
yyrule2: // {D}+
	{
		if lval.item, err = this.getInt(); nil != err {
			return -1
		} else {
			return NUMBER
		}
		goto yystate0
	}
yyrule3: // {LAB}
	{
		lval.item = this.getStr()
		logP.Infof("Label found")
		return LABEL
		goto yystate0
	}
yyrule4: // {ACT}
	{
		lval.item = this.getStr()
		logP.Infof("Action found")
		return ACTION
		goto yystate0
	}
yyrule5: // {OBJTY}
	{
		lval.item = this.getStr()
		logP.Infof("Object Type found")
		return OBJECT_TYPE
		goto yystate0
	}
yyrule6: // {SIGN}
	{
		lval.item = this.getStr()
		return SIGN
		goto yystate0
	}
yyrule7: // {TXTPA}
	{
		lval.item = this.getParam()
		return TEXT_PARAM
		goto yystate0
	}
yyrule8: // \0
	{
		return -1
	}
yyrule9: // {TXT}
	{
		lval.item = this.getStr()
		logP.Infof("Text found")
		return TEXT
		goto yystate0
	}
yyrule10: // [ \t]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized

	return -1
}
