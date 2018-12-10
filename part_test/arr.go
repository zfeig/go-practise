package main
import(
	"fmt"
)

func main() {
   arr := [...]string{"a","b","c","d"}
   for i:= range arr {
      fmt.Println("arr item:",i," is ",arr[i])
   }	
}