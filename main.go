package main

import (
	"fmt"
	"os"

	"github.com/georgfedermann/goecho/echolib"
)

func main() {
	fmt.Println(echolib.ReadArgsIntoString(os.Args[1:]))
}
