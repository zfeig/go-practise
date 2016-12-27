package main
import "fmt"
func main() {
 s := []int{1,2,3,4,5,6,7,8}
 for i:=0;i<len(s);i++ {
  fmt.Println("s[%d] == %d",i,s[i])
 }
}
