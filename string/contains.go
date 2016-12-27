package main
import(
  "fmt"
  "strings"
)

func main(){
  str:="hello,world good well,what a fuck day,no problemokyesbye"
  fmt.Println(strings.Contains(str,"well"))
  fmt.Println(strings.Contains(str,"good"))
  fmt.Println(strings.Contains(str,"talk"))
  fmt.Println(strings.Contains(str,"what"))
}
