all: bddLangParser.go bddLangLexer.y bddLangLexer.go

bddLangParser.go: bddLang.l
	golex -t -o bddLangParser.go bddLang.l | gofmt > bddLangParser.go
	
bddLangLexer.go: bddLangLexer.y
	param-yacc -p Feature -o bddLangLexer.go -d 0 -e=true bddLangLexer.y
	
bddLangLexer.y: bddLang.ebnf
	ebnf2y -pkg lexer -start Feature bddLang.ebnf > bddLangLexer.y
	
	
param-yacc:
	cd $(GOPATH)/src/github.com/ONSBR/bddTest.Go/param-yacc
	go install	
	
definitions-backup:
	if [ -e old/ ];then rm -rf old/; fi
	mkdir old
	cp bddLang.ebnf old/bddLang.ebnf.old
	cp bddLangLexer.y old/bddLangLexer.y.old
	cp bddLang.l old/bddLang.l.old
clean:
	rm -f bddLangParser.go bddLangLexer.y bddLangLexer.go y.output
	#rm -f lexer.yy.go