package main
import(
  "fmt"
)

func main(){

   obj := make(map[string]string)
   obj["Italy"] = "Rome"
   obj["Japan"] = "Tokyo"
   obj["India"] = "New Delhi"
   for key,v := range obj {
      fmt.Println("Capital of",key,"is",v)
   }
}
