// CAUTION: Generated file - DO NOT EDIT.

//
package lexer

import (
	//"fmt"
	"github.com/ONSBR/bddTest.Go/util"
)

var logP = util.GetLogger("lexer.parser")

func (this *BddTestLex) Lex(lval *FeatureSymType) (ret int) { // begin main()

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
		goto yystate3 // c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c == '!' || c >= '#' && c <= ',' || c == '.' || c == '/' || c >= ':' && c <= '@' || c == 'B' || c >= 'F' && c <= 'P' || c >= 'R' && c <= 'a' || c == 'd' || c >= 'f' && c <= 'k' || c >= 'm' && c <= 'o' || c == 'q' || c == 'r' || c >= 't' && c <= 'ÿ'
	case c == '"':
		goto yystate6
	case c == '-':
		goto yystate11
	case c == 'A':
		goto yystate13
	case c == 'C':
		goto yystate21
	case c == 'D':
		goto yystate29
	case c == 'E':
		goto yystate33
	case c == 'Q':
		goto yystate36
	case c == '\n':
		goto yystate5
	case c == '\t' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == 'b':
		goto yystate39
	case c == 'c':
		goto yystate44
	case c == 'e':
		goto yystate51
	case c == 'l':
		goto yystate56
	case c == 'p':
		goto yystate60
	case c == 's':
		goto yystate66
	case c >= '0' && c <= '9':
		goto yystate12
	}

yystate2:
	c = this.next()
	goto yyrule10

yystate3:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate4:
	c = this.next()
	switch {
	default:
		goto yyrule12
	case c == '\t' || c == ' ':
		goto yystate4
	}

yystate5:
	c = this.next()
	switch {
	default:
		goto yyrule1
	case c == '\n':
		goto yystate5
	}

yystate6:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == '\t' || c == ' ':
		goto yystate8
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate7
	}

yystate7:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == '"':
		goto yystate10
	case c == '\t' || c == ' ':
		goto yystate8
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c == '!' || c >= '#' && c <= 'ÿ':
		goto yystate7
	}

yystate8:
	c = this.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate9
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate8
	}

yystate9:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == '"':
		goto yystate9
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate8
	}

yystate10:
	c = this.next()
	switch {
	default:
		goto yyrule9
	case c == '"':
		goto yystate10
	case c == '\t' || c == ' ':
		goto yystate8
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c == '!' || c >= '#' && c <= 'ÿ':
		goto yystate7
	}

yystate11:
	c = this.next()
	switch {
	default:
		goto yyrule8
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate12:
	c = this.next()
	switch {
	default:
		goto yyrule2
	case c >= '0' && c <= '9':
		goto yystate12
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '/' || c >= ':' && c <= 'ÿ':
		goto yystate3
	}

yystate13:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 's':
		goto yystate14
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate14:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'p':
		goto yystate15
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'o' || c >= 'q' && c <= 'ÿ':
		goto yystate3
	}

yystate15:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'e':
		goto yystate16
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate16:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'c':
		goto yystate17
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'b' || c >= 'd' && c <= 'ÿ':
		goto yystate3
	}

yystate17:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 't':
		goto yystate18
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate18:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'o':
		goto yystate19
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate19:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == ':':
		goto yystate20
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '9' || c >= ';' && c <= 'ÿ':
		goto yystate3
	}

yystate20:
	c = this.next()
	switch {
	default:
		goto yyrule4
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate21:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'e':
		goto yystate22
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate22:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'n':
		goto yystate23
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate23:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'a':
		goto yystate24
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate24:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'r':
		goto yystate25
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'q' || c >= 's' && c <= 'ÿ':
		goto yystate3
	}

yystate25:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'i':
		goto yystate26
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate26:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'o':
		goto yystate27
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate27:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == ':':
		goto yystate28
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '9' || c >= ';' && c <= 'ÿ':
		goto yystate3
	}

yystate28:
	c = this.next()
	switch {
	default:
		goto yyrule3
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate29:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'a':
		goto yystate30
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate30:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'd':
		goto yystate31
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'c' || c >= 'e' && c <= 'ÿ':
		goto yystate3
	}

yystate31:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'o':
		goto yystate32
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate32:
	c = this.next()
	switch {
	default:
		goto yyrule5
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate33:
	c = this.next()
	switch {
	default:
		goto yyrule5
	case c == 'n':
		goto yystate34
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate34:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 't':
		goto yystate35
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate35:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'a':
		goto yystate31
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate36:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'u':
		goto yystate37
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 't' || c >= 'v' && c <= 'ÿ':
		goto yystate3
	}

yystate37:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'a':
		goto yystate38
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate38:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'n':
		goto yystate30
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate39:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'o':
		goto yystate40
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate40:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 't':
		goto yystate41
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate41:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'a':
		goto yystate42
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate42:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'o':
		goto yystate43
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate43:
	c = this.next()
	switch {
	default:
		goto yyrule7
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate44:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'a':
		goto yystate45
	case c == 'l':
		goto yystate47
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'k' || c >= 'm' && c <= 'ÿ':
		goto yystate3
	}

yystate45:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'm':
		goto yystate46
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'l' || c >= 'n' && c <= 'ÿ':
		goto yystate3
	}

yystate46:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'p':
		goto yystate42
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'o' || c >= 'q' && c <= 'ÿ':
		goto yystate3
	}

yystate47:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'i':
		goto yystate48
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate48:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'c':
		goto yystate49
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'b' || c >= 'd' && c <= 'ÿ':
		goto yystate3
	}

yystate49:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'o':
		goto yystate50
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate50:
	c = this.next()
	switch {
	default:
		goto yyrule6
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate51:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'x':
		goto yystate52
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'w' || c >= 'y' && c <= 'ÿ':
		goto yystate3
	}

yystate52:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'i':
		goto yystate53
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate53:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 's':
		goto yystate54
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate54:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 't':
		goto yystate55
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate55:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'e':
		goto yystate50
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate56:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'i':
		goto yystate57
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate57:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 's':
		goto yystate58
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate58:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 't':
		goto yystate59
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate59:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'a':
		goto yystate43
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate60:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'r':
		goto yystate61
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'q' || c >= 's' && c <= 'ÿ':
		goto yystate3
	}

yystate61:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'e':
		goto yystate62
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate62:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'e':
		goto yystate63
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate63:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'n':
		goto yystate64
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate64:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'c':
		goto yystate65
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'b' || c >= 'd' && c <= 'ÿ':
		goto yystate3
	}

yystate65:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'h':
		goto yystate49
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'g' || c >= 'i' && c <= 'ÿ':
		goto yystate3
	}

yystate66:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'e':
		goto yystate67
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate67:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'l':
		goto yystate68
	case c == 's':
		goto yystate73
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'k' || c >= 'm' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate68:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'e':
		goto yystate69
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate69:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'c':
		goto yystate70
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'b' || c >= 'd' && c <= 'ÿ':
		goto yystate3
	}

yystate70:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'i':
		goto yystate71
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate71:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'o':
		goto yystate72
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate72:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 'n':
		goto yystate49
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate73:
	c = this.next()
	switch {
	default:
		goto yyrule11
	case c == 's':
		goto yystate41
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yyrule1: // [\n]+
	{
		logP.Infof("returning newline")
		this.newLine()
		return NEW_LINE
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
yyrule3: // {SCNLAB}
	{
		lval.item = this.getLabel()
		logP.Infof("Scenario Label found")
		return SCENARIO_LABEL
		goto yystate0
	}
yyrule4: // {FEATLAB}
	{
		lval.item = this.getLabel()
		logP.Infof("Feature Label found")
		return FEATURE_LABEL
		goto yystate0
	}
yyrule5: // {LAB}
	{
		lval.item = this.getLabel()
		logP.Infof("Label found")
		return LABEL
		goto yystate0
	}
yyrule6: // {ACT}
	{
		lval.item = this.getStr()
		logP.Infof("Action found")
		return ACTION
		goto yystate0
	}
yyrule7: // {OBJTY}
	{
		lval.item = this.getStr()
		logP.Infof("Object Type found")
		return OBJECT_TYPE
		goto yystate0
	}
yyrule8: // {SIGN}
	{
		lval.item = this.getStr()
		return SIGN
		goto yystate0
	}
yyrule9: // {TXTPA}
	{
		lval.item = this.getParam()
		return TEXT_PARAM
		goto yystate0
	}
yyrule10: // \0
	{
		return -1
	}
yyrule11: // {TXT}
	{
		lval.item = this.getStr()
		logP.Infof("Text found: %s", lval.item)
		return TEXT
		goto yystate0
	}
yyrule12: // [ \t]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized

	return -1
}
