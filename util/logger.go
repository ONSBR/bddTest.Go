package util

import (
	"fmt"
	loggo "github.com/juju/loggo"
)

func InitLog() {
	configuration, err := GetConfig("/Users/coichedid/Projetos/Golang/src/github.com/ONSBR/bddTest.Go/util/conf.json")
	if err != nil {
		fmt.Println("Error reading config")
		return
	}
	config := SerializeConfig(configuration.Logging)
	loggo.ConfigureLoggers(config)
}

func GetLogger(module string) loggo.Logger {
	return loggo.GetLogger(module)
}
