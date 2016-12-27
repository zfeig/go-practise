package main
import(
  "fmt"
  "strings"
)

func main(){
  str:="hello,world good well,what a fuck day,no problemokyesbye"
  fmt.Println(strings.Index(str,"well"))
  fmt.Println(strings.Index(str,"good"))
  fmt.Println(strings.Index(str,"hello"))
  fmt.Println(strings.Index(str,"test"))
}
