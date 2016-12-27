package main
import(
  "fmt"
  "os"
)

func main(){
 defer fmt.Println("ready to exit!")
 os.Exit(1)
}
