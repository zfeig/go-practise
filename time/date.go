package main
import(
  "fmt"
  "time"
)

func main(){
  t:= time.Date(2016,time.December,13,0,10,10,0,time.UTC)
  fmt.Printf("go launched at: %s\n",t.Local())
  tm := time.Now()
  fmt.Println("now is:",tm)
}
