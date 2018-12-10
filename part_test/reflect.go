package main 
import (
"fmt"
"reflect"
)

func main() {
  var  x float64 =3.45
  fmt.Println("type:",reflect.TypeOf(x))
  v := reflect.ValueOf(x)
  fmt.Println("v value:", v)
  fmt.Println("v type:", v.Type())
  fmt.Println("v kind:", v.Kind())
  fmt.Println("v value:", v.Float())
  fmt.Println("v interface:",v.Interface())
  fmt.Printf("value is %5.2e\n", v.Interface())
  y := v.Interface().(float64)
  fmt.Println(y)

  

}