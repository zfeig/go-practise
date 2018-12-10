package main 
import (
"fmt"
"strconv"
)
type T struct{
   a int
   b float64
   c string
}

func (tt T) String() string {
	return  strconv.Itoa(tt.a) + "/" + strconv.FormatFloat(tt.b, 'f', 3, 32) + "/" + tt.c
}

func main() {
  test := T{7,3.351,"abc\tdef"}
  str := test.String()
  fmt.Println(str)
}