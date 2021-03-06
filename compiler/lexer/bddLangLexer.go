//line bddLangLexer.y:2

//TODO Put your favorite license here

// yacc source generated by ebnf2y[1]
// at 2016-05-04 11:47:01.771562225 -0300 BRT
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
const EXPECT_ACTION_ACTION = 57347
const EXPECT_ACTION_LABEL = 57348
const MATCHER = 57349
const FEATURE_LABEL = 57350
const INIT_SCENARIO_LABEL = 57351
const LABEL = 57352
const NEW_LINE = 57353
const NUMBER = 57354
const OBJECT_TYPE = 57355
const PAGE_LABEL = 57356
const SCENARIO_LABEL = 57357
const SIGN = 57358
const TEXT = 57359
const TEXT_PARAM = 57360
const USER_SCENARIO_LABEL = 57361

var FeatureToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ACTION",
	"EXPECT_ACTION_ACTION",
	"EXPECT_ACTION_LABEL",
	"MATCHER",
	"FEATURE_LABEL",
	"INIT_SCENARIO_LABEL",
	"LABEL",
	"NEW_LINE",
	"NUMBER",
	"OBJECT_TYPE",
	"PAGE_LABEL",
	"SCENARIO_LABEL",
	"SIGN",
	"TEXT",
	"TEXT_PARAM",
	"USER_SCENARIO_LABEL",
	"'.'",
}
var FeatureStatenames = [...]string{}

const FeatureEofCode = 1
const FeatureErrCode = 2
const FeatureInitialStackSize = 16

//line bddLangLexer.y:639

type (
	Expect_action struct {
		LineNum    int
		FullText   string
		Label      string
		Action     string
		ObjectType string
		ObjectId   string
		Param      string
	}
	Expect_action1    string
	Expect_action2    string
	Expect_action3    string
	Expect_action4    string
	Expect_expression struct {
		LineNum    int
		FullText   string
		Label      string
		Action     string
		ObjectType string
		ObjectId   string
		Matcher    string
		Param      string
	}
	Expect_expression1 string
	Expect_expression2 string
	Expect_expression3 string
	Expect_expression4 string
	Feature            struct {
		LineNum   int
		FullText  string
		Label     string
		Name      string
		PageName  string
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
	Init_scenario  string
	Init_scenario1 string
	Number_param   string
	Number_param1  string
	Number_param2  string
	Object         struct {
		ObjectType string
		ObjectId   string
	}
	Page_line string
	Param     string
	Scenario  struct {
		LineNum      int
		FullText     string
		Label        string
		Name         string
		Username     string
		Actions      []Expect_action
		Expectations []Expect_expression
	}
	Scenario1     Test_block
	Scenario_line struct {
		FullText string
		Label    string
		Name     string
		LineNum  int
	}
	Start      interface{}
	Test_block struct {
		Username     string
		Actions      []Expect_action
		Expectations []Expect_expression
	}
	Test_block1             []Expect_action
	Test_block2             []Expect_expression
	Test_expect_line        Expect_expression
	Test_expect_line1       int
	Test_init_scenario_line string
	Test_line               Expect_action
	Text_line               string
	Text_line1              string
	Text_line11             string
)

var parsedCode Feature

func GetParsedCode() interface{} {
	log.Infof("GetParsedCode |%s|", parsedCode.Label)
	return parsedCode
}

//line yacctab:1
var FeatureExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 7,
	14, 22,
	-2, 53,
	-1, 54,
	12, 29,
	-2, 8,
	-1, 67,
	12, 29,
	-2, 17,
}

const FeatureNprod = 57
const FeaturePrivate = 57344

var FeatureTokenNames []string
var FeatureStates []string

const FeatureLast = 78

var FeatureAct = [...]int{

	60, 50, 35, 25, 70, 64, 59, 61, 34, 64,
	33, 61, 15, 18, 67, 52, 52, 48, 17, 57,
	51, 55, 39, 41, 68, 9, 14, 6, 73, 49,
	27, 66, 44, 11, 36, 40, 31, 38, 28, 23,
	24, 27, 4, 38, 16, 47, 10, 21, 43, 42,
	30, 20, 1, 13, 19, 56, 12, 5, 69, 63,
	62, 29, 22, 7, 3, 8, 2, 71, 72, 65,
	53, 45, 37, 58, 54, 46, 32, 26,
}
var FeaturePact = [...]int{

	34, -1000, -1000, 13, -1000, -1000, -1000, -1000, 11, -1000,
	1, -1000, -1000, 30, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 20, 27, -1000, -1000, -1000, 25, -1000, -1000, -9,
	31, -1000, 18, 6, -1000, -1000, -1000, 21, -1000, -1000,
	-1000, -1000, 37, -1000, -1000, 12, 3, -1000, -1000, -1000,
	-1000, -1000, 4, 2, -11, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 19, -1000, 7, -16, -7, -1000, -1000,
	16, -1000, -1000, -1000,
}
var FeaturePgo = [...]int{

	0, 77, 76, 75, 74, 73, 72, 71, 70, 69,
	67, 66, 65, 64, 63, 62, 61, 60, 59, 58,
	1, 57, 0, 56, 54, 53, 52, 51, 50, 49,
	2, 48, 47, 3, 25, 46, 44,
}
var FeatureR1 = [...]int{

	0, 1, 2, 2, 3, 3, 4, 4, 5, 5,
	6, 7, 7, 8, 8, 9, 9, 10, 10, 11,
	12, 12, 13, 14, 14, 15, 16, 16, 17, 18,
	18, 19, 19, 20, 21, 22, 22, 23, 24, 24,
	25, 26, 27, 28, 28, 29, 29, 30, 31, 31,
	32, 33, 34, 35, 35, 36, 36,
}
var FeatureR2 = [...]int{

	0, 7, 0, 2, 0, 2, 0, 2, 0, 1,
	8, 0, 2, 0, 2, 0, 2, 0, 1, 3,
	0, 2, 2, 0, 2, 4, 0, 2, 3, 0,
	1, 0, 2, 2, 2, 1, 1, 2, 0, 1,
	2, 1, 5, 0, 2, 0, 2, 2, 0, 1,
	2, 2, 2, 0, 2, 1, 1,
}
var FeatureChk = [...]int{

	-1000, -26, -11, -13, 8, -21, 14, -14, -12, -34,
	-35, -34, -23, -25, 15, 11, -36, 17, 12, -24,
	-27, -32, -15, 9, -34, -33, -1, 10, 11, -16,
	-28, 11, -2, 19, 17, -30, -33, -6, 6, 4,
	17, 17, -29, -31, 11, -7, -3, -30, 5, 17,
	-20, 17, 13, -8, -4, 17, -20, 17, -5, 17,
	-22, 18, -17, -18, 16, -9, 12, 7, 17, -19,
	20, -10, -22, 12,
}
var FeatureDef = [...]int{

	0, -2, 41, 0, 23, 20, 53, -2, 19, 34,
	0, 24, 21, 38, 53, 52, 54, 55, 56, 37,
	39, 0, 0, 26, 40, 43, 0, 2, 50, 0,
	0, 51, 0, 0, 27, 45, 44, 48, 11, 4,
	3, 25, 42, 47, 49, 0, 0, 46, 13, 12,
	6, 5, 0, 0, -2, 33, 15, 14, 1, 7,
	9, 35, 36, 0, 30, 0, 31, -2, 16, 28,
	0, 10, 18, 32,
}
var FeatureTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 20,
}
var FeatureTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19,
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
		//line bddLangLexer.y:111
		{
			log.Infof("Expect_action found! |%s|", FeatureDollar[1].item.(string))
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
			buffer.WriteString(" \"")
			buffer.WriteString(FeatureDollar[7].item.(string))
			buffer.WriteString("\"")
			FeatureVAL.item = Expect_action{FullText: buffer.String(), Label: FeatureDollar[1].item.(string), Action: FeatureDollar[3].item.(string), ObjectType: FeatureDollar[5].item.(Object).ObjectType, ObjectId: FeatureDollar[5].item.(Object).ObjectId, Param: FeatureDollar[7].item.(string)}
		}
	case 2:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:135
		{
			log.Infof("Expect_action1 empty found!")
			FeatureVAL.item = ""
		}
	case 3:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:140
		{
			log.Infof("Expect_action1 found!")
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
		//line bddLangLexer.y:153
		{
			log.Infof("Expect_action2 empty found!")
			FeatureVAL.item = ""
		}
	case 5:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:158
		{
			log.Infof("Expect_action2 found!")
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
		//line bddLangLexer.y:171
		{
			log.Infof("Expect_action3 empty found!")
			FeatureVAL.item = ""
		}
	case 7:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:176
		{
			log.Infof("Expect_action3 found!")
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
		//line bddLangLexer.y:189
		{
			log.Infof("Expect_action4 empty found!")
			FeatureVAL.item = ""
		}
	case 9:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:194
		{
			log.Infof("Expect_action4 found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 10:
		FeatureDollar = FeatureS[Featurept-8 : Featurept+1]
		//line bddLangLexer.y:201
		{
			log.Infof("Expect_expression found!")
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
			buffer.WriteString(" \"")
			buffer.WriteString(FeatureDollar[8].item.(string))
			buffer.WriteString("\"")
			FeatureVAL.item = Expect_expression{FullText: buffer.String(), Label: FeatureDollar[1].item.(string), Action: FeatureDollar[3].item.(string), ObjectType: FeatureDollar[5].item.(Object).ObjectType, ObjectId: FeatureDollar[5].item.(Object).ObjectId, Matcher: FeatureDollar[7].item.(string), Param: FeatureDollar[8].item.(string)}
		}
	case 11:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:227
		{
			log.Infof("Expect_expression1 empty found!")
			FeatureVAL.item = ""
		}
	case 12:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:232
		{
			log.Infof("Expect_expression1 found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 13:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:245
		{
			log.Infof("Expect_expression2 empty found!")
			FeatureVAL.item = ""
		}
	case 14:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:250
		{
			log.Infof("Expect_expression2 found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 15:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:263
		{
			log.Infof("Expect_expression3 empty found!")
			FeatureVAL.item = ""
		}
	case 16:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:268
		{
			log.Infof("Expect_expression3 found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 17:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:281
		{
			log.Infof("Expect_expression4 empty found!")
			FeatureVAL.item = ""
		}
	case 18:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:286
		{
			log.Infof("Expect_expression4 found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 19:
		FeatureDollar = FeatureS[Featurept-3 : Featurept+1]
		//line bddLangLexer.y:293
		{
			log.Infof("Feature found! |%s|", FeatureDollar[1].item.(Feature_block).Label)
			featureBlock := FeatureDollar[1].item.(Feature_block)
			page := FeatureDollar[2].item.(string)
			feature := Feature{LineNum: featureBlock.LineNum, FullText: featureBlock.FullText, Name: featureBlock.Name, Label: featureBlock.Label, Scenarios: FeatureDollar[3].item.([]Scenario), PageName: page}
			FeatureVAL.item = feature
		}
	case 20:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:303
		{
			FeatureVAL.item = []Scenario(nil)
		}
	case 21:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:307
		{
			log.Infof("Feature1 found!")
			if nil != FeatureDollar[2].item {
				FeatureVAL.item = append(FeatureDollar[1].item.([]Scenario), FeatureDollar[2].item.(Scenario))
			} else {
				FeatureVAL.item = FeatureDollar[1].item.([]Scenario)
			}
		}
	case 22:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:318
		{
			log.Infof("Feature_block found! %s", FeatureDollar[1].item.(string))

			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			buffer.WriteString(": ")
			buffer.WriteString(FeatureDollar[2].item.(string))
			featureBlock := Feature_block{LineNum: 1, FullText: buffer.String(), Label: FeatureDollar[1].item.(string), Name: FeatureDollar[2].item.(string)}
			FeatureVAL.item = featureBlock
		}
	case 23:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:331
		{
			FeatureVAL.item = ""
		}
	case 24:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:335
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
	case 25:
		FeatureDollar = FeatureS[Featurept-4 : Featurept+1]
		//line bddLangLexer.y:348
		{
			log.Infof("Init_scenario found!")
			FeatureVAL.item = FeatureDollar[4].item.(string)
		}
	case 26:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:355
		{
			log.Infof("Init_scenario1 empty found!")
			FeatureVAL.item = ""
		}
	case 27:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:360
		{
			log.Infof("Init_scenario1 found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			if len(FeatureDollar[1].item.(string)) > 0 {
				buffer.WriteString(" ")
			}
			buffer.WriteString(FeatureDollar[2].item.(string))
			FeatureVAL.item = buffer.String()
		}
	case 28:
		FeatureDollar = FeatureS[Featurept-3 : Featurept+1]
		//line bddLangLexer.y:373
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
	case 29:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:391
		{
			log.Infof("Number_param1 empty found!")
			FeatureVAL.item = ""
		}
	case 30:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:396
		{
			log.Infof("Number_param1 found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 31:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:403
		{
			log.Infof("Number_param2 empty found!")
			FeatureVAL.item = ""
		}
	case 32:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:408
		{
			log.Infof("Number_param2 found!")
			var buffer bytes.Buffer
			buffer.WriteString(".")
			buffer.WriteString(strconv.Itoa(FeatureDollar[2].item.(int)))
			FeatureVAL.item = buffer.String()
		}
	case 33:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:418
		{
			log.Infof("Object found!")
			FeatureVAL.item = Object{ObjectType: FeatureDollar[1].item.(string), ObjectId: FeatureDollar[2].item.(string)}
		}
	case 34:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:425
		{
			log.Infof("Page_line found!")

			FeatureVAL.item = FeatureDollar[2].item.(string)
		}
	case 35:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:433
		{
			log.Infof("Text param found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 36:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:438
		{
			log.Infof("Number Param found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 37:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:445
		{
			log.Infof("Scenario found!")
			line := FeatureDollar[1].item.(Scenario_line)
			testBlock := FeatureDollar[2].item.(Test_block)
			scenario := Scenario{LineNum: line.LineNum,
				FullText:     line.FullText,
				Label:        line.Label,
				Name:         line.Name,
				Actions:      testBlock.Actions,
				Expectations: testBlock.Expectations,
				Username:     testBlock.Username}
			FeatureVAL.item = scenario
		}
	case 38:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:461
		{
			FeatureVAL.item = Test_block{}
		}
	case 39:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:465
		{
			log.Infof("Scenario1 found!")
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 40:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:472
		{
			log.Infof("Scenario_line found!")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			buffer.WriteString(": ")
			buffer.WriteString(FeatureDollar[2].item.(string))
			lineNum := 0
			if v, ok := Featurelex.(*BddTestLex); ok {
				lineNum = v.lineNum
			}
			line := Scenario_line{LineNum: lineNum, FullText: buffer.String(), Label: FeatureDollar[1].item.(string), Name: FeatureDollar[2].item.(string)}
			FeatureVAL.item = line
		}
	case 41:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:488
		{
			log.Infof("Start,|%s|", FeatureDollar[1].item.(Feature).Label)
			//if v,ok := Featurelex.(*BddTestLex); ok {
			//	v.OnBddTest($1)
			//}
			parsedCode = FeatureDollar[1].item.(Feature)
		}
	case 42:
		FeatureDollar = FeatureS[Featurept-5 : Featurept+1]
		//line bddLangLexer.y:498
		{
			log.Infof("Test_block")

			action := FeatureDollar[2].item.(Expect_action)
			actions := append([]Expect_action{action}, FeatureDollar[3].item.([]Expect_action)...)

			expectation := FeatureDollar[4].item.(Expect_expression)
			expectations := append([]Expect_expression{expectation}, FeatureDollar[5].item.([]Expect_expression)...)

			FeatureVAL.item = Test_block{Username: FeatureDollar[1].item.(string), Actions: actions, Expectations: expectations}
		}
	case 43:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:512
		{
			FeatureVAL.item = []Expect_action(nil)
		}
	case 44:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:516
		{
			log.Infof("Test_block1")
			if nil != FeatureDollar[2].item {
				lineNum := 0
				if v, ok := Featurelex.(*BddTestLex); ok {
					lineNum = v.lineNum
				}
				action := FeatureDollar[2].item.(Expect_action)
				action.LineNum = lineNum
				log.Infof("Line Num %d", action.LineNum)
				FeatureVAL.item = append(FeatureDollar[1].item.([]Expect_action), action)
			} else {
				FeatureVAL.item = FeatureDollar[1].item.([]Expect_action)
			}
		}
	case 45:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:534
		{
			FeatureVAL.item = []Expect_expression(nil)
		}
	case 46:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:538
		{
			log.Infof("Test_block2")
			if nil != FeatureDollar[2].item {
				lineNum := 0
				if v, ok := Featurelex.(*BddTestLex); ok {
					lineNum = v.lineNum
				}
				expectation := FeatureDollar[2].item.(Expect_expression)
				expectation.LineNum = lineNum
				FeatureVAL.item = append(FeatureDollar[1].item.([]Expect_expression), expectation)
			} else {
				FeatureVAL.item = FeatureDollar[1].item.([]Expect_expression)
			}
		}
	case 47:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:555
		{
			log.Infof("Test_expect_line")

			expectation := FeatureDollar[1].item.(Expect_expression)

			lineNum := 0
			if v, ok := Featurelex.(*BddTestLex); ok {
				if FeatureDollar[2].item.(int) == -1 {
					v.newLine()
				}
				lineNum = v.lineNum
			}
			expectation.LineNum = lineNum
			FeatureVAL.item = expectation
		}
	case 48:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:571
		{
			FeatureVAL.item = -1
		}
	case 49:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:575
		{
			log.Infof("New line")
			FeatureVAL.item = 0
		}
	case 50:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:582
		{
			log.Infof("Test_init_scenario_line found!")
			FeatureVAL.item = FeatureDollar[1].item.(string)
		}
	case 51:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:589
		{
			log.Infof("Test_line")
			action := FeatureDollar[1].item.(Expect_action)
			lineNum := 0
			if v, ok := Featurelex.(*BddTestLex); ok {
				lineNum = v.lineNum
			}
			action.LineNum = lineNum
			FeatureVAL.item = action
		}
	case 52:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:602
		{
			log.Infof("Text_line found")
			var buffer bytes.Buffer
			buffer.WriteString(FeatureDollar[1].item.(string))
			//buffer.WriteString("\n")
			FeatureVAL.item = buffer.String()
		}
	case 53:
		FeatureDollar = FeatureS[Featurept-0 : Featurept+1]
		//line bddLangLexer.y:612
		{
			FeatureVAL.item = ""
		}
	case 54:
		FeatureDollar = FeatureS[Featurept-2 : Featurept+1]
		//line bddLangLexer.y:616
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
	case 55:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:631
		{
			FeatureVAL.item = FeatureDollar[1].item
		}
	case 56:
		FeatureDollar = FeatureS[Featurept-1 : Featurept+1]
		//line bddLangLexer.y:635
		{
			FeatureVAL.item = strconv.Itoa(FeatureDollar[1].item.(int))
		}
	}
	goto Featurestack /* stack new state and value */
}
