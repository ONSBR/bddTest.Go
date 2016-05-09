package main

import (
    "github.com/ONSBR/bddTest.Go/cli"
    
    "os"
)

func main()  {
    mainCli := cli.NewCli()
    retCode := mainCli.Run()
    os.Exit(retCode)
}