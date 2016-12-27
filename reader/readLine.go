package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	input, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open failed!")
	}
	defer input.Close()

	reader := bufio.NewReader(input)
	var count int
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		count++
		fmt.Println("line:", count, ":", line)
	}

}
