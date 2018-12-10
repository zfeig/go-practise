package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    go sendData(ch)
    go getData(ch)

    ch1 := make(chan int)
    go pump(ch1)       // pump hangs
    fmt.Println("get send data:", <-ch1) // prints only 0

    time.Sleep(1e9)
}

func sendData(ch chan string) {
	fmt.Println("try to send data...")
    ch <- "aaaa"
    fmt.Println("send data aaaa")
    ch <- "bbbb"
    fmt.Println("send data bbbb")
    ch <- "cccc"
    fmt.Println("send data cccc")
    ch <- "dddd"
    fmt.Println("send data dddd")
    ch <- "eeee"
    fmt.Println("send data eeee")
}

func getData(ch chan string) {
	fmt.Println("try to get data ...")
    var input string
    // time.Sleep(2e9)
    for {
        input = <-ch
        fmt.Printf("get data:%s\n", input)
    }
}

func pump(ch chan int) {
    for i := 10; ; i++ {
        ch <- i
    }
}