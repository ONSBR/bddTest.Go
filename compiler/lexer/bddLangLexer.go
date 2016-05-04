//line bddLangLexer.y:2

//TODO Put your favorite license here

// yacc source generated by ebnf2y[1]
// at 2016-05-03 14:19:12.727122628 -0300 BRT
//
//  $ ebnf2y -pkg lexer -start Feature bddLang.ebnf
//
// CAUTION: If this file is a Go source file (*.go), it was generated
// automatically by '$ go tool yacc' from a *.y file - DO NOT EDIT in that case!
//
//   [1]: http://github.com/cznic/ebnf2y

package lexer

import __yyfmt__ "fmt"

//line bddLangLexer.y:15
import (
	"bytes"
	//"fmt"
	"github.com/ONSBR/bddTest.Go/util"
	"strconv"
)

var log = util.GetLogger("lexer.lexer")

//line bddLangLexer.y:28
type FeatureSymType struct {
	yys  int
	item interface{} //TODO insert real field(s)
}

const ACTION = 57346
const FEATURE_LABEL = 57347
const LABEL = 57348
const NEW_LINE = 57349
const NUMBER = 57350
const OBJECT_TYPE = 57351
const SCENARIO_LABEL = 57352
const SIGN = 57353
const TEXT = 57354
const TEXT_PARAM = 57355

var FeatureToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ACTION",
	"FEATURE_LABEL",
	"LABEL",
	"NEW_LINE",
	"NUMBER",
	"OBJECT_TYPE",
	"SCENARIO_LABEL",
	"SIGN",
	"TEXT",
	"TEXT_PARAM",
	"'.'",
}
var FeatureStatenames = [...]string{}

const FeatureEofCode = 1
const FeatureErrCode = 2
const FeatureInitialStackSize = 16

//line bddLangLexer.y:422

type (
	Assertion struct {
		LineNum    int
		FullText   string
		Label      string
		Action     string
		ObjectType string
		ObjectId   string
		Param      string
	}
	Assertion1 string
	Assertion2 string
	Assertion3 string
	Assertion4 string
	Feature    struct {
		LineNum   int
		FullText  string
		Label     string
		Name      string
		Scenarios []Scenario
	}
	Feature1      []Scenario
	Feature_block struct {
		LineNum  int
		FullText string
		Label    string
		Name     string
	}
	Feature_block1 string
	Number_param   string
	Number_param1  string
	Number_param2  string
	Object         struct {
		ObjectType string
		ObjectId   string
	}
	Param    string
	Scenario struct {
		LineNum    int
		FullText   string
		Label      string
		Name       string
		Assertions []Assertion
	}
	Scenario1     []Assertion
	Scenario_line struct {
		FullText string
		Label    string
		Name     string
		LineNum  int
	}
	Start       interface{}
	Test_block  []Assertion
	Test_block1 []Assertion
	Test_line   Assertion
	Test_line1  interface{}
	Text_line   string
	Text_line1  string
)

var parsedCode Feature

func GetParsedCode() interface{} {
	return parsedCode
}

//line yacctab:1
var FeatureExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 6,
	7, 36,
	12, 36,
	-2, 13,
	-1, 31,
	8, 17,
	-2, 8,
}

const FeatureNprod = 38
const FeaturePrivate = 57344

var FeatureTokenNames []string
var FeatureStates []string

const FeatureLast = 43

var FeatureAct = [...]int{

	14, 39, 34, 36, 30, 42, 25, 29, 18, 32,
	9, 43, 10, 19, 26, 40, 22, 16, 4, 11,
	21, 24, 17, 20, 13, 1, 8, 12, 7, 35,
	28, 41, 38, 37, 6, 3, 5, 2, 33, 31,
	27, 23, 15,
}
var FeaturePact = [...]int{

	13, -1000, -1000, -1000, -1000, 0, -1000, -1000, 11, -1000,
	-1000, 1, -1000, -1000, -1000, 9, -1000, -1000, -1000, -1000,
	11, -1000, -1000, 2, -1000, -1000, -1000, -5, -1000, -1000,
	-3, -10, -1000, -1000, -1000, -1000, -1000, -1000, 7, -1000,
	-9, -1000, 3, -1000,
}
var FeaturePgo = [...]int{

	0, 42, 41, 40, 39, 38, 37, 36, 35, 34,
	33, 32, 31, 30, 29, 28, 27, 26, 25, 24,
	23, 0, 20, 12, 19,
}
var FeatureR1 = [...]int{

	0, 1, 2, 2, 3, 3, 4, 4, 5, 5,
	6, 7, 7, 8, 9, 9, 10, 11, 11, 12,
	12, 13, 14, 14, 15, 16, 16, 17, 18, 19,
	20, 20, 21, 22, 22, 23, 24, 24,
}
var FeatureR2 = [...]int{

	0, 7, 0, 2, 0, 2, 0, 2, 0, 1,
	2, 0, 2, 2, 0, 2, 3, 0, 1, 0,
	2, 2, 1, 1, 2, 0, 1, 2, 1, 2,
	0, 2, 2, 0, 1, 2, 0, 2,
}
var FeatureChk = [...]int{

	-1000, -18, -6, -8, 5, -7, -9, -15, -17, 10,
	-23, -24, -16, -19, -21, -1, 6, -23, 7, 12,
	-20, -22, 7, -2, -21, 4, 12, -3, -13, 12,
	9, -4, 12, -5, 12, -14, 13, -10, -11, 11,
	8, -12, 14, 8,
}
var FeatureDef = [...]int{

	0, -2, 28, 11, 14, 10, -2, 12, 25, 36,
	15, 0, 24, 26, 30, 33, 2, 27, 35, 37,
	29, 32, 34, 0, 31, 4, 3, 0, 6, 5,
	0, -2, 21, 1, 7, 9, 22, 23, 0, 18,
	19, 16, 0, 20,
}
var FeatureTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 14,
}
var FeatureTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13,
}
var FeatureTok3 = [...]int{
	0,
}

var FeatureErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	FeatureDebug        = 0
	FeatureErrorVerbose = true
)

type FeatureLexer interface {
	Lex(lval *FeatureSymType) int
	Error(s string)
}

type FeatureParser interface {
	Parse(FeatureLexer) int
	Lookahead() int
}

type FeatureParserImpl struct {
	lval  FeatureSymType
	stack [FeatureInitialStackSize]FeatureSymType
	char  int
}

func (p *FeatureParserImpl) Lookahead() int {
	return p.char
}

func FeatureNewParser() FeatureParser {
	return &FeatureParserImpl{}
}

const FeatureFlag = -1000

func FeatureTokname(c int) string {
	if c >= 1 && c-1 < len(FeatureToknames) {
		if FeatureToknames[c-1] != "" {
			return FeatureToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func FeatureStatname(s int) string {
	if s >= 0 && s < len(FeatureStatenames) {
		if FeatureStatenames[s] != "" {
			return FeatureStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func FeatureErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !FeatureErrorVerbose {
		return "syntax error"
	}

	for _, e := range FeatureErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + FeatureTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := FeaturePact[state]
	for tok := TOKSTART; tok-1 < len(FeatureToknames); tok++ {
		if n := base + tok; n >= 0 && n < FeatureLast && FeatureChk[FeatureAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if FeatureDef[state] == -2 {
		i := 0
		for FeatureExca[i] != -1 || FeatureExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; FeatureExca[i] >= 0; i += 2 {
			tok := FeatureExca[i]
			if tok < TOKSTART || FeatureExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if FeatureExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += FeatureTokname(tok)
	}
	return res
}

func Featurelex1(lex FeatureLexer, lval *FeatureSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = FeatureTok1[0]
		goto out
	}
	if char < len(FeatureTok1) {
		token = FeatureTok1[char]
		goto out
	}
	if char >= FeaturePrivate {
		if char < FeaturePrivate+len(FeatureTok2) {
			token = FeatureTok2[char-FeaturePrivate]
			goto out
		}
	}
	for i := 0; i < len(FeatureTok3); i += 2 {
		token = FeatureTok3[i+0]
		if token == char {
			token = FeatureTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = FeatureTok2[1] /* unknown char */
	}
	if FeatureDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", FeatureTokname(token), uint(char))
	}
	return char, token
}

func FeatureParse(Featurelex FeatureLexer) int {
	return FeatureNewParser().Parse(Featurelex)
}

func (Featurercvr *FeatureParserImpl) Parse(Featurelex FeatureLexer) int {
	var Featuren int
	var FeatureVAL FeatureSymType
	var FeatureDollar []FeatureSymType
	_ = FeatureDollar // silence set and not used
	FeatureS := Featurercvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Featurestate := 0
	Featurercvr.char = -1
	Featuretoken := -1 // Featurercvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Featurestate = -1
		Featurercvr.char = -1
		Featuretoken = -1
	}()
	Featurep := -1
	goto Featurestack

ret0:
	return 0

ret1:
	return 1

Featurestack:
	/* put a state and value onto the stack */
	if FeatureDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", FeatureTokname(Featuretoken), FeatureStatname(Featurestate))
	}

	Featurep++
	if Featurep >= len(FeatureS) {
		nyys := make([]FeatureSymType, len(FeatureS)*2)
		copy(nyys, FeatureS)
		FeatureS = nyys
	}
	FeatureS[Featurep] = FeatureVAL
	FeatureS[Featurep].yys = Featurestate

Featurenewstate:
	Featuren = FeaturePact[Featurestate]
	if Featuren <= FeatureFlag {
		goto Featuredefault /* simple state */
	}
	if Featurercvr.char < 0 {
		Featurercvr.char, Featuretoken = Featurelex1(Featurelex, &Featurercvr.lval)
	}
	Featuren += Featuretoken
	if Featuren < 0 || Featuren >= FeatureLast {
		goto Featuredefault
	}
	Featuren = FeatureAct[Featuren]
	if FeatureChk[Featuren] == Featuretoken { /* valid shift */
		Featurercvr.char = -1
		Featuretoken = -1
		FeatureVAL = Featurercvr.lval
		Featurestate = Featuren
		if Errflag > 0 {
			Errflag--
		}
		goto Featurestack
	}

Featuredefault:
	/* default state action */
	Featuren = FeatureDef[Featurestate]
	if Featuren == -2 {
		if Featurercvr.char < 0 {
			Featurercvr.char, Featuretoken = Featurelex1(Featurelex, &Featurercvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if FeatureExca[xi+0] == -1 && FeatureExca[xi+1] == Featurestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Featuren = FeatureExca[xi+0]
			if Featuren < 0 || Featuren == Featuretoken {
				break
			}
		}
		Featuren = FeatureExca[xi+1]
		if Featuren < 0 {
			goto ret0
		}
	}
	if Featuren == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Featurelex.Error(FeatureErrorMessage(Featurestate, Featuretoken))
			Nerrs++
			if FeatureDebug >= 1 {
				__yyfmt__.Printf("%s", FeatureStatname(Featurestate))
				__yyfmt__.Printf(" saw %s\n", FeatureTokname(Featuretoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Featurep >= 0 {
				Featuren = FeaturePact[FeatureS[Featurep].yys] + FeatureErrCode
				if Featuren >= 0 && Featuren < FeatureLast {
					Featurestate = FeatureAct[Featuren] /* simulate a shift of "error" */
					if FeatureChk[Featurestate] == FeatureErrCode {
						goto Featurestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if FeatureDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", FeatureS[Featurep].yys)
				}
				Featurep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if FeatureDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", FeatureTokname(Featuretoken))
			}
			if Featuretoken == FeatureEofCode {
				goto ret1
			}
			Featurercvr.char = -1
			Featuretoken = -1
			goto Featurenewstate /* try again in the same state */
		}
	}

	/* reduction by production Featuren */
	if FeatureDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Featuren, FeatureStatname(Featurestate))
	}

	Featurent := Featuren
	Featurept := Featurep
	_ = Featurept // guard against "declared and not used"

	Featurep -= FeatureR2[Featuren]
	// Featurep is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Featurep+1 >= len(FeatureS) {
		nyys := make([]FeatureSymType, len(FeatureS)*2)
		copy(nyys, FeatureS)
		FeatureS = nyys
	}
	FeatureVAL = FeatureS[Featurep+1]

	/* consult goto table to find next state */
	Featuren = FeatureR1[Featuren]
	Featureg := FeaturePgo[Featuren]
	Featurej := Featureg + FeatureS[Featurep].yys + 1

	if Featurej >= FeatureLast {
		Featurestate = FeatureAct[Featureg]
	} else {
		Featurestate = FeatureAct[Featurej]
		if FeatureChk[Featurestate] != -Featuren {
			Featurestate = FeatureAct[Featureg]
		}
	}
	// dummy call; replaced with literal code
	switch Featurent {

	case 1:
		FeatureDollar = FeatureS[Featurept-7 : Featurept+1]
		//line bddLangLexer.y:88
		{
			log.Infof("Assertion found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[2].item.(string))
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[3].item.(string))
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[4].item.(string))
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[5].item.(Object).ObjectType)
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[5].item.(Object).ObjectId)
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[6].item.(string))
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[7].item.(string))
			FeatureVAL.item = Assertion{FullText: buffer.String(), Label: FeatureDollar[1].item.(string), Action: FeatureDollar[3].item.(string), ObjectType: FeatureDollar[5].item.(Object).ObjectType, ObjectId: FeatureDollar[5].item.(Object).ObjectId, Param: FeatureDollar[7].item.(string)}
		}
	case 2:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:111
		{
			log.Infof("Assertion1 empty found!")
			FeatureVAL.item = ""
		}
	case 3:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:116
		{
			log.Infof("Assertion1 found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 4:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:129
		{
			log.Infof("Assertion2 empty found!")
			FeatureVAL.item = ""
		}
	case 5:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:134
		{
			log.Infof("Assertion2 found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 6:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:147
		{
			log.Infof("Assertion3 empty found!")
			FeatureVAL.item = ""
		}
	case 7:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:152
		{
			log.Infof("Assertion3 found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 8:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:165
		{
			log.Infof("Assertion4 empty found!")
			FeatureVAL.item = ""
		}
	case 9:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:170
		{
			log.Infof("Assertion4 found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 10:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:177
		{
			log.Infof("Feature found!")
			featureBlock := FeatureDollar[1].item.(Feature_block)
			feature := Feature{LineNum: featureBlock.LineNum, FullText: featureBlock.FullText, Name: featureBlock.Name, Label: featureBlock.Label, Scenarios: FeatureDollar[2].item.([]Scenario)}
			FeatureVAL.item = feature
		}
	case 11:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:186
		{
			FeatureVAL.item = []Scenario(nil)
		}
	case 12:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:190
		{
			log.Infof("Feature1 found!")
			if nil != FeatureDollar[2].item {
				FeatureVAL.item = append(FeatureDollar[1].item.([]Scenario), FeatureDollar[2].item.(Scenario))
			} else {
				FeatureVAL.item = FeatureDollar[1].item.([]Scenario)
			}
		}
	case 13:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:201
		{
			log.Infof("Feature_block found!")

			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[2].item.(string))
			featureBlock := Feature_block{LineNum: 1, FullText: buffer.String(), Label: FeatureDollar[1].item.(string), Name: FeatureDollar[2].item.(string)}
			FeatureVAL.item = featureBlock
		}
	case 14:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:214
		{
			FeatureVAL.item = ""
		}
	case 15:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:218
		{
			log.Infof("Feature_block1 found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString("\n")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 16:
		FeatureDollar = FeatureS[Featurept-3 : Featurept+1]
		//line bddLangLexer.y:231
		{
			log.Infof("Number_param found!")
			var buffer bytes.Buffer
			if FeatureDollar[1].item != nil {
				buffer.WriteString(FeatureDollar[1].item.(string))
			}

			buffer.WriteString(strconv.Itoa(FeatureDollar[2].item.(int)))

			if FeatureDollar[3].item != nil {
				buffer.WriteString(FeatureDollar[3].item.(string))
			}

			FeatureVAL.item = buffer.String()
		}
	case 17:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:249
		{
			log.Infof("Number_param1 empty found!")
			FeatureVAL.item = ""
		}
	case 18:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:254
		{
			log.Infof("Number_param1 found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 19:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:261
		{
			log.Infof("Number_param2 empty found!")
			FeatureVAL.item = ""
		}
	case 20:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:266
		{
			log.Infof("Number_param2 found!")
			var buffer bytes.Buffer
			buffer.WriteString(".")
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 21:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:276
		{
			log.Infof("Object found!")
			FeatureVAL.item = Object{ObjectType: FeatureDollar[1].item.(string), ObjectId: FeatureDollar[2].item.(string)}
		}
	case 22:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:283
		{
			log.Infof("Text param found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 23:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:288
		{
			log.Infof("Number Param found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 24:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:295
		{
			log.Infof("Scenario found!")
			line := FeatureDollar[1].item.(Scenario_line)
			scenario := Scenario{LineNum: line.LineNum, FullText: line.FullText, Label: line.Label, Name: line.Name, Assertions: FeatureDollar[2].item.([]Assertion)}
			FeatureVAL.item = scenario
		}
	case 25:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:304
		{
			FeatureVAL.item = []Assertion(nil)
		}
	case 26:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:308
		{
			log.Infof("Scenario1 found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 27:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:315
		{
			log.Infof("Scenario_line found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			buffer.WriteString(" ")
			buffer.WriteString(FeatureDollar[2].item.(string))
			lineNum := 0
			if v, ok := Featurelex.(*BddTestLex); ok {
				lineNum = v.lineNum
			}
			line := Scenario_line{LineNum: lineNum, FullText: buffer.String(), Label: FeatureDollar[1].item.(string), Name: FeatureDollar[2].item.(string)}
			FeatureVAL.item = line
		}
	case 28:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:331
		{
			log.Infof("Start")
			//if v,ok := Featurelex.(*BddTestLex); ok {
			//	v.OnBddTest($1)
			//}
			parsedCode = FeatureDollar[1].item.(Feature)
		}
	case 29:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:341
		{
			log.Infof("Test_block")
			lineNum := 0
			if v, ok := Featurelex.(*BddTestLex); ok {
				lineNum = v.lineNum
			}
			assertion := FeatureDollar[1].item.(Assertion)
			assertion.LineNum = lineNum
			log.Infof("Line Num %d", assertion.LineNum)
			FeatureVAL.item = append([]Assertion{assertion}, FeatureDollar[2].item.([]Assertion)...)
		}
	case 30:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:355
		{
			FeatureVAL.item = []Assertion(nil)
		}
	case 31:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:359
		{
			log.Infof("Test_block1")
			if nil != FeatureDollar[2].item {
				lineNum := 0
				if v, ok := Featurelex.(*BddTestLex); ok {
					lineNum = v.lineNum
				}
				assertion := FeatureDollar[2].item.(Assertion)
				assertion.LineNum = lineNum
				log.Infof("Line Num %d", assertion.LineNum)
				FeatureVAL.item = append(FeatureDollar[1].item.([]Assertion), assertion)
			} else {
				FeatureVAL.item = FeatureDollar[1].item.([]Assertion)
			}
		}
	case 32:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:377
		{
			log.Infof("Test_line")
			FeatureVAL.item = FeatureDollar[1].item.(Assertion)
		}
	case 33:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:384
		{
			FeatureVAL.item = nil
		}
	case 34:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:388
		{
			log.Infof("New line")
			FeatureVAL.item = nil
		}
	case 35:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:395
		{
			log.Infof("Text_line found")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			//buffer.WriteString("\n")
			FeatureVAL.item = buffer.String()
		}
	case 36:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:405
		{
			FeatureVAL.item = ""
		}
	case 37:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:409
		{
			log.Infof("Text_line1 found")

			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))

			FeatureVAL.item = buffer.String()
		}
	}
	goto Featurestack /* stack new state and value */
}
