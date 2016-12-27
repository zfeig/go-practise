package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "test.txt"
	outputFile := "out.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("read file failed!")
	}
	fmt.Println("out:", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}
}
