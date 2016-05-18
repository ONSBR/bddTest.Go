package main

import (
	"io/ioutil"
	"log"

	"github.com/ONSBR/bddTest.Go/cli"

	//Bootstrap strategy handler
	_ "github.com/ONSBR/bddTest.Go/compiler/langMaps"
	_ "github.com/ONSBR/bddTest.Go/compiler/lexer/locales"

	"os"
)

func main() {
	log.SetOutput(ioutil.Discard)

	mainCli := cli.NewCli()
	retCode := mainCli.Run()
	os.Exit(retCode)
}
