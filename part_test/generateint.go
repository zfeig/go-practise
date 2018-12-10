package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int,100)
    go pump(ch1)       // pump hangs
    go  pick(ch1)
    time.Sleep(1e9)
}



func pump(ch chan int) {
    for i := 1; ; i++ {
        ch <- i
    }
}

func pick(ch chan int) {
    for {
        fmt.Println(<-ch)
    }
}

func test() {
    fmt.Println("this is a nest test")
}
