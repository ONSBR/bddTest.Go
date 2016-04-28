all: bddLangParser.go bddLangLexer.y bddLangLexer.go

bddLangParser.go: bddLang.l
	golex -t -o bddLangParser.go bddLang.l | gofmt > bddLangParser.go
	
bddLangLexer.go: bddLangLexer.y
	if [ -e bddLangLexer.go ];then cp -vf bddLangLexer.go bddLangLexer.go.bak; fi
	go tool yacc -p Test_Scenario -o bddLangLexer.go bddLangLexer.y
	
bddLangLexer.y: bddLang.ebnf
	if [ -e bddLangLexer.y ];then cp -vf bddLangLexer.y bddLangLexer.y.bak; fi
	ebnf2y -pkg lexer -start Test_Scenario bddLang.ebnf > bddLangLexer.y
	
clean:
	rm -f bddLangParser.go bddLangLexer.y bddLangLexer.go y.output
	#rm -f lexer.yy.go