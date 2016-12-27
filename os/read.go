package main
import(
 "fmt"
 "os"
)

func main(){
  flname := "./info.txt"
  fl,err:= os.Open(flname)
  if err!=nil{
   fmt.Println("error!")
   return
  }
  defer fl.Close()
  buf:= make([]byte,1024)
  for{
    n,_:= fl.Read(buf)
    if 0 == n {
      break
    }
    fmt.Println(string(buf[:n]))
    // os.Stdout.Write(buf[:n])
  }
}
