// CAUTION: Generated file - DO NOT EDIT.

//
package locales

import (
	//"fmt"
	. "github.com/ONSBR/bddTest.Go/compiler/lexer"
)

//LexerLocalized_pt_BR strategy implementation of lex pt_BR
type LexerLocalized_pt_BR struct{}

func init() {
	RegisterLexLocalization("pt_BR", (*LexerLocalized_pt_BR)(nil))
}

func (this *LexerLocalized_pt_BR) LexLocalized(next NextFunc, back BackFunc, resetBuffer ResetBuffFunc) (ret int) { // begin main()

	var c byte = ' '

	defer func() {
		back()
	}()

yystate0:

	resetBuffer()

	goto yystart1

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = next()
yystart1:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c == '!' || c >= '#' && c <= ',' || c == '.' || c == '/' || c >= ':' && c <= '@' || c == 'B' || c >= 'F' && c <= 'O' || c >= 'R' && c <= 'a' || c == 'd' || c >= 'f' && c <= 'h' || c == 'j' || c == 'k' || c == 'n' || c == 'o' || c == 'q' || c == 'r' || c == 't' || c >= 'v' && c <= 'ÿ'
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
	case c == 'P':
		goto yystate38
	case c == 'Q':
		goto yystate45
	case c == '\n':
		goto yystate5
	case c == '\t' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == 'b':
		goto yystate51
	case c == 'c':
		goto yystate56
	case c == 'e':
		goto yystate68
	case c == 'i':
		goto yystate78
	case c == 'l':
		goto yystate85
	case c == 'm':
		goto yystate89
	case c == 'p':
		goto yystate106
	case c == 's':
		goto yystate112
	case c == 'u':
		goto yystate120
	case c >= '0' && c <= '9':
		goto yystate12
	}

yystate2:
	c = next()
	goto yyrule16

yystate3:
	c = next()
	switch {
	default:
		goto yyrule17
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate4:
	c = next()
	switch {
	default:
		goto yyrule18
	case c == '\t' || c == ' ':
		goto yystate4
	}

yystate5:
	c = next()
	switch {
	default:
		goto yyrule1
	case c == '\n':
		goto yystate5
	}

yystate6:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == '\t' || c == ' ':
		goto yystate8
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate7
	}

yystate7:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == '"':
		goto yystate10
	case c == '\t' || c == ' ':
		goto yystate8
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c == '!' || c >= '#' && c <= 'ÿ':
		goto yystate7
	}

yystate8:
	c = next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate9
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate8
	}

yystate9:
	c = next()
	switch {
	default:
		goto yyrule15
	case c == '"':
		goto yystate9
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate8
	}

yystate10:
	c = next()
	switch {
	default:
		goto yyrule15
	case c == '"':
		goto yystate10
	case c == '\t' || c == ' ':
		goto yystate8
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c == '!' || c >= '#' && c <= 'ÿ':
		goto yystate7
	}

yystate11:
	c = next()
	switch {
	default:
		goto yyrule14
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate12:
	c = next()
	switch {
	default:
		goto yyrule2
	case c >= '0' && c <= '9':
		goto yystate12
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '/' || c >= ':' && c <= 'ÿ':
		goto yystate3
	}

yystate13:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 's':
		goto yystate14
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate14:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'p':
		goto yystate15
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'o' || c >= 'q' && c <= 'ÿ':
		goto yystate3
	}

yystate15:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate16
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate16:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'c':
		goto yystate17
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'b' || c >= 'd' && c <= 'ÿ':
		goto yystate3
	}

yystate17:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 't':
		goto yystate18
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate18:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate19
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate19:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == ':':
		goto yystate20
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '9' || c >= ';' && c <= 'ÿ':
		goto yystate3
	}

yystate20:
	c = next()
	switch {
	default:
		goto yyrule10
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate21:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate22
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate22:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'n':
		goto yystate23
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate23:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate24
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate24:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'r':
		goto yystate25
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'q' || c >= 's' && c <= 'ÿ':
		goto yystate3
	}

yystate25:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'i':
		goto yystate26
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate26:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate27
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate27:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == ':':
		goto yystate28
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '9' || c >= ';' && c <= 'ÿ':
		goto yystate3
	}

yystate28:
	c = next()
	switch {
	default:
		goto yyrule9
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate29:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate30
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate30:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'd':
		goto yystate31
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'c' || c >= 'e' && c <= 'ÿ':
		goto yystate3
	}

yystate31:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate32
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate32:
	c = next()
	switch {
	default:
		goto yyrule3
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate33:
	c = next()
	switch {
	default:
		goto yyrule11
	case c == 'n':
		goto yystate34
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate34:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 't':
		goto yystate35
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate35:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate36
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate36:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate37
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate37:
	c = next()
	switch {
	default:
		goto yyrule6
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate38:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate39
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate39:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'g':
		goto yystate40
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'f' || c >= 'h' && c <= 'ÿ':
		goto yystate3
	}

yystate40:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'i':
		goto yystate41
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate41:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'n':
		goto yystate42
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate42:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate43
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate43:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == ':':
		goto yystate44
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '9' || c >= ';' && c <= 'ÿ':
		goto yystate3
	}

yystate44:
	c = next()
	switch {
	default:
		goto yyrule4
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate45:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'u':
		goto yystate46
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 't' || c >= 'v' && c <= 'ÿ':
		goto yystate3
	}

yystate46:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate47
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate47:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'n':
		goto yystate48
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate48:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'd':
		goto yystate49
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'c' || c >= 'e' && c <= 'ÿ':
		goto yystate3
	}

yystate49:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate50
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate50:
	c = next()
	switch {
	default:
		goto yyrule11
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate51:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate52
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate52:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 't':
		goto yystate53
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate53:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate54
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate54:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate55
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate55:
	c = next()
	switch {
	default:
		goto yyrule13
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate56:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate57
	case c == 'l':
		goto yystate59
	case c == 'o':
		goto yystate63
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'k' || c == 'm' || c == 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate57:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'm':
		goto yystate58
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'l' || c >= 'n' && c <= 'ÿ':
		goto yystate3
	}

yystate58:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'p':
		goto yystate54
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'o' || c >= 'q' && c <= 'ÿ':
		goto yystate3
	}

yystate59:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'i':
		goto yystate60
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate60:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'c':
		goto yystate61
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'b' || c >= 'd' && c <= 'ÿ':
		goto yystate3
	}

yystate61:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate62
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate62:
	c = next()
	switch {
	default:
		goto yyrule12
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate63:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'n':
		goto yystate64
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate64:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 't':
		goto yystate65
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate65:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate66
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate66:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'm':
		goto yystate67
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'l' || c >= 'n' && c <= 'ÿ':
		goto yystate3
	}

yystate67:
	c = next()
	switch {
	default:
		goto yyrule8
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate68:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 's':
		goto yystate69
	case c == 'x':
		goto yystate74
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'w' || c >= 'y' && c <= 'ÿ':
		goto yystate3
	}

yystate69:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'p':
		goto yystate70
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'o' || c >= 'q' && c <= 'ÿ':
		goto yystate3
	}

yystate70:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate71
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate71:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'r':
		goto yystate72
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'q' || c >= 's' && c <= 'ÿ':
		goto yystate3
	}

yystate72:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate73
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate73:
	c = next()
	switch {
	default:
		goto yyrule7
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate74:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'i':
		goto yystate75
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate75:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 's':
		goto yystate76
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate76:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 't':
		goto yystate77
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate77:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate62
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate78:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'g':
		goto yystate79
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'f' || c >= 'h' && c <= 'ÿ':
		goto yystate3
	}

yystate79:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'u':
		goto yystate80
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 't' || c >= 'v' && c <= 'ÿ':
		goto yystate3
	}

yystate80:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate81
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate81:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'l':
		goto yystate82
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'k' || c >= 'm' && c <= 'ÿ':
		goto yystate3
	}

yystate82:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == ' ':
		goto yystate83
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate83:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'a':
		goto yystate84
	}

yystate84:
	c = next()
	goto yyrule8

yystate85:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'i':
		goto yystate86
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate86:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 's':
		goto yystate87
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate87:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 't':
		goto yystate88
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 's' || c >= 'u' && c <= 'ÿ':
		goto yystate3
	}

yystate88:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate55
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate89:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate90
	case c == 'e':
		goto yystate105
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate90:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'i':
		goto yystate91
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate91:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate92
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate92:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'r':
		goto yystate93
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'q' || c >= 's' && c <= 'ÿ':
		goto yystate3
	}

yystate93:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == ' ':
		goto yystate94
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yystate94:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'o':
		goto yystate95
	case c == 'q':
		goto yystate103
	}

yystate95:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'u':
		goto yystate96
	}

yystate96:
	c = next()
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate97
	}

yystate97:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'i':
		goto yystate98
	}

yystate98:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'g':
		goto yystate99
	}

yystate99:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'u':
		goto yystate100
	}

yystate100:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'a':
		goto yystate101
	}

yystate101:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'l':
		goto yystate102
	}

yystate102:
	c = next()
	switch {
	default:
		goto yyabort
	case c == ' ':
		goto yystate83
	}

yystate103:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'u':
		goto yystate104
	}

yystate104:
	c = next()
	switch {
	default:
		goto yyabort
	case c == 'e':
		goto yystate84
	}

yystate105:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'n':
		goto yystate91
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate106:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'r':
		goto yystate107
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'q' || c >= 's' && c <= 'ÿ':
		goto yystate3
	}

yystate107:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate108
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate108:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate109
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate109:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'n':
		goto yystate110
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate110:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'c':
		goto yystate111
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'b' || c >= 'd' && c <= 'ÿ':
		goto yystate3
	}

yystate111:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'h':
		goto yystate61
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'g' || c >= 'i' && c <= 'ÿ':
		goto yystate3
	}

yystate112:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate113
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate113:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'l':
		goto yystate114
	case c == 's':
		goto yystate119
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'k' || c >= 'm' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate114:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'e':
		goto yystate115
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'd' || c >= 'f' && c <= 'ÿ':
		goto yystate3
	}

yystate115:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'c':
		goto yystate116
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'b' || c >= 'd' && c <= 'ÿ':
		goto yystate3
	}

yystate116:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'i':
		goto yystate117
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate117:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate118
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate118:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'n':
		goto yystate61
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'm' || c >= 'o' && c <= 'ÿ':
		goto yystate3
	}

yystate119:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 's':
		goto yystate53
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate120:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 's':
		goto yystate121
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'r' || c >= 't' && c <= 'ÿ':
		goto yystate3
	}

yystate121:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'u':
		goto yystate122
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 't' || c >= 'v' && c <= 'ÿ':
		goto yystate3
	}

yystate122:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate123
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate3
	}

yystate123:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'r':
		goto yystate124
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'q' || c >= 's' && c <= 'ÿ':
		goto yystate3
	}

yystate124:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'i':
		goto yystate125
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'h' || c >= 'j' && c <= 'ÿ':
		goto yystate3
	}

yystate125:
	c = next()
	switch {
	default:
		goto yyrule17
	case c == 'o':
		goto yystate126
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'n' || c >= 'p' && c <= 'ÿ':
		goto yystate3
	}

yystate126:
	c = next()
	switch {
	default:
		goto yyrule5
	case c >= '\x01' && c <= '\b' || c >= '\v' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate3
	}

yyrule1: // [\n]+
	{
		logP.Infof("returning newline")
		return NEW_LINE
		goto yystate0
	}
yyrule2: // {D}+
	{
		logP.Infof("Number found")
		return NUMBER

		goto yystate0
	}
yyrule3: // {INITLAB}
	{
		logP.Infof("Init assertion Label found")
		return INIT_SCENARIO_LABEL
		goto yystate0
	}
yyrule4: // {PAGELAB}
	{
		logP.Infof("Page Label found")
		return PAGE_LABEL
		goto yystate0
	}
yyrule5: // {USERLAB}
	{
		logP.Infof("User Label found")
		return USER_SCENARIO_LABEL
		goto yystate0
	}
yyrule6: // {EXPTLAB}
	{
		logP.Infof("Expect assertion Label found")
		return EXPECT_ACTION_LABEL
		goto yystate0
	}
yyrule7: // {EXPTACT}
	{
		logP.Infof("Expect Action found")
		return EXPECT_ACTION_ACTION
		goto yystate0
	}
yyrule8: // {MATCHER}
	{
		logP.Infof("Matcher found")
		return MATCHER

		goto yystate0
	}
yyrule9: // {SCNLAB}
	{
		logP.Infof("Scenario Label found")
		return SCENARIO_LABEL
		goto yystate0
	}
yyrule10: // {FEATLAB}
	{
		logP.Infof("Feature Label found")
		return FEATURE_LABEL
		goto yystate0
	}
yyrule11: // {LAB}
	{
		logP.Infof("Label found")
		return LABEL
		goto yystate0
	}
yyrule12: // {ACT}
	{
		logP.Infof("Action found")
		return ACTION
		goto yystate0
	}
yyrule13: // {OBJTY}
	{
		logP.Infof("Object Type found")
		return OBJECT_TYPE
		goto yystate0
	}
yyrule14: // {SIGN}
	{
		logP.Infof("Sign found")
		return SIGN
		goto yystate0
	}
yyrule15: // {TXTPA}
	{
		logP.Infof("Text param found")
		return TEXT_PARAM
		goto yystate0
	}
yyrule16: // \0
	{
		logP.Infof("EOF")
		return -1
		goto yystate0
	}
yyrule17: // {TXT}
	{
		logP.Infof("Text found")
		return TEXT
		goto yystate0
	}
yyrule18: // [ \t]+

	goto yystate0
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized

	return -1
}
