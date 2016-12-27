package main
import(
  "fmt"
  "os"
)


func main(){
  os.Mkdir("hello",0777)
  os.MkdirAll("world/test/good",0777)
  fmt.Println("mkdir ok!")
}
