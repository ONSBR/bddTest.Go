package main

import (
	"fmt"
	"os"
)

const (
	sep = string(os.PathSeparator)
)

func main() {
	fmt.Println(sep)
}
