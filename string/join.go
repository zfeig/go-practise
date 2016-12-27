package main
import(
  "fmt"
  "strings"
)

func main(){
  s:=[]string{"hello","world","wonderfull","day"}
  fmt.Println(strings.Join(s," "))
}
