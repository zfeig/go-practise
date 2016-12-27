package main
import(
  "fmt"
)

func incr() func() int {
  i := 0
  return func() int {
   i++;
   return i
  }

}


func main() {

  index := 0

  for j:=0;j<5;j++ {
    index += incr()()
  }

  fmt.Println("r:",index)

}
