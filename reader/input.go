package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Input something...")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("The input was:%s\n", err)
	}

	switch strings.Trim(input, "\r\n") {
	case "M":
		fmt.Println("今天是周一")
	case "T":
		fmt.Println("今天是周二")
	case "W":
		fmt.Println("今天是周三")
	case "TH":
		fmt.Println("今天是周四")
	case "Fri":
		fmt.Println("今天是周五")
	case "S":
		fmt.Println("今天是周末")
	default:
		fmt.Println("今天日期未知！")

	}
}
