%{
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
	
	defer func(){
		this.back()
	}()


%}

D			[0-9]
L			[A-Za-z]
LAB			"dado"|"Dado"|"quando"|"Quando"|"entao"|"Entao"|"e"|"E"
ACT			"clico"|"seleciono"|"preencho"
OBJTY		"botao"|"campo"|"lista"
SIGN		"-"
TXTPA		\".+\"
TXT		[A-Za-z0-9]+

%yyc c
%yyn c = this.next()

%%

		if nil!=this.buf {
			this.buf = this.buf[len(this.buf)-1:]
		}
		
[\n]+				logP.Infof("returning newline");this.newLine() //; return NEW_LINE	

{D}+				if lval.item, err = this.getInt(); nil!=err {
						return -1
					} else {
						return NUMBER
					}

{LAB}				lval.item = this.getStr(); logP.Infof("Label found"); return LABEL

{ACT}				lval.item = this.getStr(); logP.Infof("Action found"); return ACTION

{OBJTY}				lval.item = this.getStr(); logP.Infof("Object Type found"); return OBJECT_TYPE

{SIGN}				lval.item = this.getStr(); return SIGN

{TXTPA}				lval.item = this.getParam(); return TEXT_PARAM

\0					return -1

{TXT}				lval.item = this.getStr(); logP.Infof("Text found"); return TEXT

[ \t]+			

%%

return -1
}

