package util

import (
	"encoding/json"
	"fmt"
	"os"
	// "path/filepath"
)

type Configuration struct {
	Logging map[string]string
}

func GetConfig(path string) (config Configuration, er error) {
	file, err := os.Open(path)
	if err != nil {
		er = err
		return
	}
	fmt.Println(err)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		er = err
		return
	}
	config = configuration
	er = nil
	return
}

func SerializeConfig(items map[string]string) string {
	ret := ""
	for k, v := range items {
		ret += k + "=" + v + ";"
	}
	return ret
}
