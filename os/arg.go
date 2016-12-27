package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	sep := "*"
	for k, v := range os.Args[0:] {
		if k > 0 {
			s += sep + v
		}
		fmt.Println("k is:", k, " v is:", v)
	}
	fmt.Println("string is:", s)
	fmt.Println("string is:", strings.Join(os.Args[1:], "**"))
}
