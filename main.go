package main

import (
    "github.com/ONSBR/bddTest.Go/cli"
    
    //Bootstrap strategy handler
	_ "github.com/ONSBR/bddTest.Go/compiler/lexer/locales"
	_ "github.com/ONSBR/bddTest.Go/compiler/langMaps"
    
    "os"
)

func main()  {
    mainCli := cli.NewCli()
    retCode := mainCli.Run()
    os.Exit(retCode)
}