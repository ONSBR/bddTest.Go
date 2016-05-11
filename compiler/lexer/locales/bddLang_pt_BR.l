%{
//
package locales

import (
    //"fmt"
    . "github.com/ONSBR/bddTest.Go/compiler/lexer"
)

//LexerLocalized_pt_BR strategy implementation of lex pt_BR
type LexerLocalized_pt_BR struct {}

func init() {
    RegisterLexLocalization("pt_BR",(*LexerLocalized_pt_BR)(nil))
}

func (this *LexerLocalized_pt_BR) LexLocalized(next NextFunc, back BackFunc, resetBuffer ResetBuffFunc) (ret int) { // begin main()

    var c byte = ' '
    
    defer func(){
		back()
	}()

%}

D           [0-9]
L           [A-Za-z]
LAB         "Quando"|"E"
INITLAB     "Dado"
PAGELAB     "Pagina:"
USERLAB     "usuario"
EXPTLAB     "Entao"
EXPTACT     "espero"
ACT         "clico"|"seleciono"|"preencho"|"existe"
OBJTY       "botao"|"campo"|"lista"|"sessao"
SIGN        "-"
TXTPA       \".+\"
TXT     [^ \t\n]+
SCNLAB      "Cenario:"
FEATLAB     "Aspecto:"

%yyc c
%yyn c = next()

%%

         resetBuffer()
        
        
[\n]+               logP.Infof("returning newline");return NEW_LINE  

{D}+                logP.Infof("Number found");return NUMBER
        
{INITLAB}           logP.Infof("Init assertion Label found"); return INIT_SCENARIO_LABEL

{PAGELAB}           logP.Infof("Page Label found"); return PAGE_LABEL

{USERLAB}           logP.Infof("User Label found"); return USER_SCENARIO_LABEL

{EXPTLAB}           logP.Infof("Expect assertion Label found"); return EXPECT_ACTION_LABEL

{EXPTACT}           logP.Infof("Expect Action found"); return EXPECT_ACTION_ACTION
            
{SCNLAB}            logP.Infof("Scenario Label found"); return SCENARIO_LABEL

{FEATLAB}           logP.Infof("Feature Label found"); return FEATURE_LABEL

{LAB}               logP.Infof("Label found"); return LABEL

{ACT}               logP.Infof("Action found"); return ACTION

{OBJTY}             logP.Infof("Object Type found"); return OBJECT_TYPE

{SIGN}              logP.Infof("Sign found");return SIGN

{TXTPA}             logP.Infof("Text param found");return TEXT_PARAM

\0                  return -1

{TXT}               logP.Infof("Text found"); return TEXT

[ \t]+          

%%

return -1
}

