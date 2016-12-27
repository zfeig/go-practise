package main
import(
  "fmt"
)


func sum(unit ...int) int {
  res := 0
  for k,v:= range unit{
    res += v
    fmt.Printf("%d is %d\n",k,v)  
  }
  return  res
}




func main(){
  
 r:= sum(1,5,8)

 fmt.Println("sum is:",r)

}
