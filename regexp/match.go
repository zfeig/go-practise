package main
import(
  "fmt"
  "regexp"
)

func isIP(ip string) (b bool) {
  m, _ := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}$", ip)
  if m {
    return true
  }
  return false
}


func main(){
 res:= isIP("127.0.0.1")
 fmt.Println("res:",res)
 r:= isIP("127.a.0.f**k")
 fmt.Println("res:",r)

}
