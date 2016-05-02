%{

//TODO Put your favorite license here
		
// yacc source generated by ebnf2y[1]
// at 2016-04-28 11:32:14.164479108 -0300 BRT
//
//  $ ebnf2y -pkg lexer -start Test_Scenario bddLang.ebnf
//
// CAUTION: If this file is a Go source file (*.go), it was generated
// automatically by '$ go tool yacc' from a *.y file - DO NOT EDIT in that case!
// 
//   [1]: http://github.com/cznic/ebnf2y

package lexer //TODO real package name

//TODO required only be the demo _dump function
import (
	"bytes"
	//"fmt"
	"strconv"
	"github.com/ONSBR/bddTest.Go/util"
)

var log = util.GetLogger("lexer.lexer") 


%}

%union {
	item interface{} //TODO insert real field(s)
}

%token	ACTION
%token	LABEL
%token	NEW_LINE
%token	NUMBER
%token	OBJECT_TYPE
%token	SIGN
%token	TEXT
%token	TEXT_PARAM

%type	<item> 	/*TODO real type(s), if/where applicable */
	ACTION
	LABEL
	NEW_LINE
	NUMBER
	OBJECT_TYPE
	SIGN
	TEXT
	TEXT_PARAM

%type	<item> 	/*TODO real type(s), if/where applicable */
	Assertion
	Assertion1
	Assertion2
	Assertion3
	Assertion4
	Number_param
	Number_param1
	Number_param2
	Object
	Param
	Start
	Test_Scenario
	Test_Scenario1
	Test_line

/*TODO %left, %right, ... declarations */

%start Start

%%

Assertion:
	LABEL Assertion1 ACTION Assertion2 Object Assertion3 Assertion4
	{
		log.Infof("Assertion found!")
		var buffer bytes.Buffer
		buffer.WriteString($1.(string))
		buffer.WriteString(" ")
		buffer.WriteString($2.(string))
		buffer.WriteString(" ")
		buffer.WriteString($3.(string))
		buffer.WriteString(" ")
		buffer.WriteString($4.(string))
		buffer.WriteString(" ")
		buffer.WriteString($5.(Object).ObjectType)
		buffer.WriteString(" ")
		buffer.WriteString($5.(Object).ObjectId)
		buffer.WriteString(" ")
		buffer.WriteString($6.(string))
		buffer.WriteString(" ")
		buffer.WriteString($7.(string))
		$$ = Assertion{FullText:buffer.String(),Label:$1.(string),Action:$3.(string),ObjectType:$5.(Object).ObjectType,ObjectId:$5.(Object).ObjectId,Param:$7.(string)}
	}

Assertion1:
	/* EMPTY */
	{
		log.Infof("Assertion1 empty found!")
		$$ = ""
	}
|	Assertion1 TEXT
	{
		log.Infof("Assertion1 found!")
		var buffer bytes.Buffer
		buffer.WriteString($1.(string))
		buffer.WriteString($2.(string))
		$$ = buffer.String()
	}

Assertion2:
	/* EMPTY */
	{
		log.Infof("Assertion2 empty found!")
		$$ = ""
	}
|	Assertion2 TEXT
	{
		log.Infof("Assertion2 found!")
		var buffer bytes.Buffer
		buffer.WriteString($1.(string))
		buffer.WriteString($2.(string))
		$$ = buffer.String()
	}

Assertion3:
	/* EMPTY */
	{
		log.Infof("Assertion3 empty found!")
		$$ = ""
	}
|	Assertion3 TEXT
	{
		log.Infof("Assertion3 found!")
		var buffer bytes.Buffer
		buffer.WriteString($1.(string))
		buffer.WriteString($2.(string))
		$$ = buffer.String()
	}

Assertion4:
	/* EMPTY */
	{
		log.Infof("Assertion4 empty found!")
		$$ = ""
	}
|	Param
	{
		log.Infof("Assertion4 found!")
		$$ = $1
	}

Number_param:
	Number_param1 NUMBER Number_param2
	{
		log.Infof("Number_param found!")
		var buffer bytes.Buffer
		if $1 != nil {
			buffer.WriteString($1.(string))
		}
		
		buffer.WriteString(strconv.Itoa($2.(int)))
		
		if $3 != nil {
			buffer.WriteString($3.(string))
		}
		
		$$ = buffer.String()
	}

Number_param1:
	/* EMPTY */
	{
		log.Infof("Number_param1 empty found!")
		$$ = "" 
	}
|	SIGN
	{
		log.Infof("Number_param1 found!")
		$$ = $1
	}

Number_param2:
	/* EMPTY */
	{
		log.Infof("Number_param2 empty found!")
		$$ = ""
	}
|	'.' NUMBER
	{
		log.Infof("Number_param2 found!")
		var buffer bytes.Buffer
		buffer.WriteString(".")
		buffer.WriteString($2.(string))
		$$ = buffer.String()
	}

Object:
	OBJECT_TYPE TEXT
	{
		log.Infof("Object found!")
		$$ = Object{ObjectType:$1.(string), ObjectId:$2.(string)}
	}

Param:
	TEXT_PARAM
	{
		log.Infof("Text param found!")
		$$ = $1
	}
|	Number_param
	{
		log.Infof("Number Param found!")
		$$ = $1
	}

Start:
	Test_Scenario
	{
		log.Infof("Start")
		if v,ok := Test_Scenariolex.(*BddTestLex); ok {
			v.OnBddTest($1)
		}
	}

Test_Scenario:
	Test_line Test_Scenario1
	{
		lineNum := 1
		assertion := $1.(Assertion)
		assertion.LineNum = lineNum
		log.Infof("Line Num %d", assertion.LineNum)
		$$ = append([]Assertion{assertion}, $2.([]Assertion)...)
	}

Test_Scenario1:
	/* EMPTY */
	{
		$$ = []Assertion(nil)
	}
|	Test_Scenario1 Test_line
	{
		lineNum := (len($1.([]Assertion)) + 2)
		assertion := $2.(Assertion)
		assertion.LineNum = lineNum
		log.Infof("Line Num %d", assertion.LineNum)
		$$ = append($1.([]Assertion),assertion)
	}

Test_line:
	NEW_LINE
	{
		//fmt.Println("Test_Line new line found!")
	}
|	Assertion
	{
		$$ = $1.(Assertion)
	}

%%

type (
	Assertion struct {
		LineNum int
		FullText string
		Label string
		Action string
		ObjectType string
		ObjectId string
		Param string
	}
	Assertion1 string
	Assertion2 string
	Assertion3 string
	Assertion4 string
	Number_param string
	Number_param1 string
	Number_param2 string
	Object struct {
		ObjectType string
		ObjectId string
	}
	Param string
	Start interface{}
	Test_Scenario []Assertion
	Test_Scenario1 []Assertion
	Test_line Assertion
)
