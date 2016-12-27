package main

import(
  "fmt"
  "reflect"
)

type TagType struct {
  field1 bool "An important answer"
  field2 string "The name of the thing"
  field3 int `how much three are`
}

func main() {
  tt := TagType{true,"hello,world",2}
  reflectType := reflect.TypeOf(tt)
  
  for i:=0;i<3; i++ {
     field :=  reflectType.Field(i)
     fmt.Println("Tag:",i," is:",field.Tag)
  }

}
