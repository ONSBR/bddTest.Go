package main

import (
	"fmt"
	"github.com/ONSBR/bddTest.Go/util"
)

func main() {
	conf := util.GetConfig("/Users/coichedid/Projetos/Golang/src/github.com/ONSBR/bddTest.Go/util/conf.json")
	fmt.Println(conf.Logging["teste1"])
	fmt.Println(util.SerializeConfig(conf.Logging))
}
