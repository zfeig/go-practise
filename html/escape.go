package main

import (
	"fmt"
	"html"
)

func main() {
	str := "<script src='http://www.ssl.com/jquery.min.js'></script>"
	res := html.EscapeString(str)
	fmt.Println("escaped:", res)
	fmt.Println("unescaped:", html.UnescapeString(res))

}
