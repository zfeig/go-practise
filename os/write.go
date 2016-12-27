package main
import(
  "fmt"
  "os"
)

func main(){
  flname:="test.txt"
  fl,err:=os.Create(flname)
  if err!=nil{
    fmt.Println("err")
    return
  }
  defer fl.Close()
  for i:=0;i<10;i++{
    fl.WriteString("hello,world!\n")
    fl.Write([]byte("write Test!\n"))
  }

}
