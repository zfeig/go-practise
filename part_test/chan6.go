package main 
import(
  "fmt"
  "time"
)

func main() {
	a := make(chan int)
	
	go func(){
		  fmt.Println("wait v",<-a)
		}()
    
    go func() {
    	a<-10
	    fmt.Printf("Type of a is : %T\n", a)
    }()

    time.Sleep(1 * time.Second)
}