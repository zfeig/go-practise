package main 
import (
 "fmt"
 "time"
)


func produce(p chan<- int) {
    for i := 0; i < 20; i++ {
        p <- i
        fmt.Println("send:", i)
    }
}


func consumer(c <-chan int) {
    for i := 0; i < 20; i++ {
        v := <-c
        fmt.Println("receive:", v)
    }
}


func main() {
    c := make(chan int, 3)
    var ss chan<- int = c
    var sr <-chan int = c
    ss<-15
    fmt.Println("ss get:",<-sr)

    //product - consumer
    ch := make(chan int,5)
    go produce(ch)
    go consumer(ch)
    time.Sleep(1 * time.Second)
}