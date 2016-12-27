package main
import(
  "fmt"
  "time"
)

func main(){
  c:= time.Tick(10*time.Second)
  for now := range c {
   fmt.Printf("hello,world",now)
  }
}
