package main

import (
	"github.com/ONSBR/bddTest.Go/compiler"
	"github.com/ONSBR/bddTest.Go/util"
	"bufio"
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

	util.InitLog()
	log := util.GetLogger("test.compiler")
	
//	fi := bufio.NewReader(os.Stdin)

	var token string
//	var ok bool
	
	tokenArr := []string{"Aspecto: Este é um aspecto",
			"Com uma segunda linha",
			"Cenario: setimo cenário",
			"Dad que existe a sessao abc",
			"Quando eu clico no botao teste com o valor \"clovis1\"",
			"E eu preencho o campo teste1 com o valor \"clovis2\"",
			"Entao eu seleciono a lista teste2 com a opcao \"clovis3\"",
			"Cenario: sexto cenário",
			"Dado que existe a sessao xpto",
			"Quando eu clico no botao teste.2 com o valor \"eu1\"",
			"E eu preencho o campo teste1.2 com o valor \"eu2\"",
			"Entao eu seleciono a lista teste2.2 com a opcao \"eu3\""}
	
	token = strings.Join(tokenArr,"\n")
//	token = "aspecto: Este é um aspecto\ncenario: primeiro cenário\nquando eu clico no botao teste com o valor \"clovis1\"\nquando eu preencho o campo teste1 com o valor \"clovis2\"\nquando eu seleciono a lista teste2 com a opcao \"clovis3\""
	
	log.Infof("Lines %s",token)
	res := compiler.ParseBddTest(token)
	if res.HasError {
		log.Errorf("Error: %s\n",res.Error.Message)
		log.Errorf("Line: %d, Pos: %d, Token: %s",res.Error.LineNum, res.Error.LinePos, res.Error.Token)
	} else {
		log.Infof("Scenarios: %d", res.NumScenarios)
		log.Infof("Line: %d Feature: %s",res.Feature.LineNum, res.Feature.Name)
		for idx, scn := range res.Feature.Scenarios {
			_ = idx
			log.Infof("Line: %d Scenario: %s",scn.LineNum, scn.Name)
			for idx2, line := range scn.Assertions {
				_ = idx2
				log.Infof("Line %d %s Object: %s<%s>.%s(%s)\n",line.LineNum, line.Label, line.ObjectId, line.ObjectType, line.Action, line.Param)	
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

