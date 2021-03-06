package util

import (
	"fmt"
	"os"
	"time"

	loggo "github.com/juju/loggo"
)

type (
	CustomFormatter struct{}
)

func InitLog(path string) {
	if path == "" {
		path = "./util/conf.json"
	}
	configuration, err := GetConfig(path)
	if err != nil {
		fmt.Println("Error reading config")
		return
	}
	config := SerializeConfig(configuration.Logging)
	loggo.ConfigureLoggers(config)
}

func GetLogger(module string) loggo.Logger {
	log := loggo.GetLogger(module)
	return log
}

func (formatter *CustomFormatter) Format(level loggo.Level, module, filename string, line int, timestamp time.Time, message string) string {
	return fmt.Sprintf("%s", message)
}

func ReplaceLoggerDefaultFormatter() {
	writer := loggo.NewSimpleWriter(os.Stderr, &CustomFormatter{})
	loggo.ReplaceDefaultWriter(writer)
}
