package main

import (
	"bufio"
	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/util"
	_ "github.com/ONSBR/bddTest.Go/compiler/lexer/locales"
	//	"fmt"
	//	"os"
	//	"flag"
	"strings"
)

//func usage() {
//	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
//	flag.PrintDefaults()
//	os.Exit(2)
//}
//
//func init() {
//	flag.Usage = usage
//	// NOTE: This next line is key you have to call flag.Parse() for the command line
//	// options or "flags" that are defined in the glog module to be picked up.
//	flag.Parse()
//}

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}

func main() {

	util.InitLog("")
	log := util.GetLogger("test.compiler")

	//	fi := bufio.NewReader(os.Stdin)

	var token string
	//	var ok bool

	tokenArr := []string{"Aspecto: Este é um aspecto",
		"Com uma segunda linha",
		"Pagina: Cadastro de Usuários",
		"Cenario: setimo 7 cenário",
		"Dado que estou usando o usuario clovis.chedid",
		"Quando eu clico no botao teste com o valor \"clovis1\"",
		"E eu preencho o campo teste1 com o valor \"clovis2\"",
		"Entao eu espero a lista teste2 com a opcao \"clovis3\"",
		"Cenario: sexto cenário",
		"Dado que estou usando o usuario clovis.chedid2",
		"Quando eu clico no botao teste.2 com o valor \"eu1\"",
		"E eu preencho o campo teste1.2 com o valor \"eu2\"",
		"Entao eu espero a lista teste2.2 com a opcao \"eu3\""}

	token = strings.Join(tokenArr, "\n")

	token = `Aspecto: Este é um aspecto 9
Pagina: Cadastro de Clientes
Cenario: primeiro 1 cenário
Dado que estou usando o usuario clovis.chedid
Quando eu clico no botao teste com o valor "clovis1"
E eu preencho o campo teste1 com o valor "clovis2"
Entao eu espero a lista teste2 com a opcao maior que "clovis3"`
	//	token = "aspecto: Este é um aspecto\ncenario: primeiro cenário\nquando eu clico no botao teste com o valor \"clovis1\"\nquando eu preencho o campo teste1 com o valor \"clovis2\"\nquando eu seleciono a lista teste2 com a opcao \"clovis3\""

	log.Infof("Lines %s", token)
	codeParser := &compiler.CodeParser{}
	Feature, retCode, err := codeParser.ParseCode(token)
	if retCode > 0 {
		log.Errorf("Error: %s\n", err.Message)
		log.Errorf("Line: %d, Pos: %d, Token: %s", err.LineNum, err.LinePos, err.Token)
	} else {
		log.Infof("Page Name: %s", Feature.PageName)
		log.Infof("Line: %d |%s|: %s", Feature.LineNum, Feature.Label,Feature.Name)
		for idx, scn := range Feature.Scenarios {
			_ = idx
			log.Infof("Line: %d Scenario: %s", scn.LineNum, scn.Name)
			log.Infof("Scenario user: %s", scn.Username)
			for idx2, line := range scn.Actions {
				_ = idx2
				log.Infof("Assert: Line %d %s Object: %s<%s>.%s(%s)\n", line.LineNum, line.Label, line.ObjectId, line.ObjectType, line.Action, line.Param)
			}
			for idx2, line := range scn.Expectations {
				_ = idx2
				log.Infof("Expect: Line %d %s Object: %s<%s>.%s(%s(%s))\n", line.LineNum, line.Label, line.ObjectId, line.ObjectType, line.Action, line.Matcher, line.Param)
			}
		}
	}
	//	for {
	//		fmt.Printf(": ")
	//		if token, ok = readline(fi); ok {
	//
	//		} else {
	//			break
	//		}
	//	}
}
