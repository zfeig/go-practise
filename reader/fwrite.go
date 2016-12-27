package main

import (
	"os"
)

func main() {
	os.Stdout.WriteString("输出文件内容\n")
	f, _ := os.OpenFile("write.txt", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()
	f.WriteString("hello,world!\n")
}
