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
	
	token = "quando eu clico no botao teste com o valor \"clovis1\"\nquando eu preencho o campo teste1 com o valor \"clovis2\"\nquando eu seleciono a lista teste2 com a opcao \"clovis3\""
	
	log.Infof("Line %s",token)
	ret := compiler.ParseBddTest(strings.TrimSpace(token),func(res compiler.ParsedTest) {
			log.Infof("Lines: %d", res.NumLines)
			log.Errorf("Error: %s",res.Error)
			for idx, line := range res.Lines {
				_ = idx
				log.Infof("Line: %d", line.LineNum)
				log.Infof("%s Object: %s<%s>.%s(%s)\n",line.Label, line.ObjectId, line.ObjectType, line.Action, line.Param)
			}
		})
	
	log.Infof("ParseBddTest(): %d", ret)
//	for {
//		fmt.Printf(": ")
//		if token, ok = readline(fi); ok {
//			
//		} else {
//			break
//		}
//	}
}

