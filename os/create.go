package main
import(
  "fmt"
  "os"
)

func main(){
  flname:="gowell.txt"
  fl,err:= os.Create(flname)
  if err!=nil{
   fmt.Println("create faild!")
  }
  defer fl.Close()
  fmt.Println("create ok!")
}

