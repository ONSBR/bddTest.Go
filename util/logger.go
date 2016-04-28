package util

import (
	loggo "github.com/juju/loggo"
)

func InitLog() {
	config := "<root>=INFO; compiler.parser=ERROR; lexer.lexer=ERROR; lexer.parser=ERROR; test.compiler=INFO"
	loggo.ConfigureLoggers(config)
}

func GetLogger(module string) loggo.Logger {
	return loggo.GetLogger(module)
}