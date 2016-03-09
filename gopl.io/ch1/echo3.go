package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], ", "))

	// exe 1.1
	fmt.Println(os.Args[0])

	// exe 1.2
	for i, arg := range os.Args[0:] {
		fmt.Println(i, " = " + arg)
	}
}