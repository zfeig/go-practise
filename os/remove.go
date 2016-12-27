package main
import(
  "fmt"
  "os"
)


func main(){
 flname:="test.txt"
 err:= os.Remove(flname)
 if err!=nil{
  fmt.Println("error")
  return
 }

}
