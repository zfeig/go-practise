package main 
import "fmt"

func main() {
	ch :=make(chan int)

	go func() {
		for {
		   v := <-ch
           fmt.Println("get ",v)
		}
	}()

	for i := 0;i<50;i++ {
    	ch <- i
	}
	fmt.Println("get 2")
}