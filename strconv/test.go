package main
import(
  "fmt"
  "strconv"
)

func main(){
  e:= strconv.Itoa(2345)
  fmt.Println(e+"00")
  f,_:= strconv.Atoi("2345")
  fmt.Println(f+2)
  g,_:= strconv.ParseInt("2345",10,64)
  fmt.Println(g+2)
  h,_:= strconv.ParseFloat("2345.34",64)
  fmt.Println(h+2.45)
}
