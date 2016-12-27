package main

import "fmt"

func main() {
	var a = new([5]int)
	test(a)
	fmt.Println(a, len(a))
}
func test(a *[5]int) {
	a[1] = 5
        fmt.Println(a)
}
