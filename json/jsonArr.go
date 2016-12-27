package main
import(
  "fmt"
  "encoding/json"
)


type Student struct{
  Age int
  Name string
  Address string
}


type Students struct{
  Students []Student
}


func main(){
  strArr:= Students{Students:[]Student{Student{Age:25,Name:"lisi",Address:"guangdong guangzhou"},Student{Age:28,Name:"xiaoq",Address:"hubei wuhan"}}}

  fmt.Println("Students",strArr);

  data,err:=json.Marshal(&strArr)
  if err!=nil{
   fmt.Printf("error!")
  }
  fmt.Println("json:",string(data))

  var s Students
  s.Students = append(s.Students,Student{Age:24,Name:"hehe",Address:"chengdu sichuan"})
  s.Students = append(s.Students,Student{Age:27,Name:"dacui",Address:"hangzhou zhejiang"})
  s.Students = append(s.Students,Student{Age:26,Name:"lilei",Address:"qingdao shangdong"})

  j,err:= json.Marshal(s)
  if err != nil{
    fmt.Printf("parse error!")
  }
  fmt.Println("json:",string(j))

}
