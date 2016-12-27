package main
import(
  "fmt"
  "strings"
)

func main(){
  str:="hello,world good well,what a hello day,no hello problemokyesbye"
  fmt.Println(strings.Replace(str,"hello","****",2))
  fmt.Println(strings.Replace(str,"hello","****",-1))
}
