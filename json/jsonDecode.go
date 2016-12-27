package main

import(
  "fmt"
  "encoding/json"
)

type Server struct{
  ServerName string
  ServerIp string
}

type ServerList struct{
  Servers []Server
}


func main(){
  var s ServerList
  str:=`{"servers":[{"serverName":"www.baidu.com","serverIp":"232.67.123.121"},{"serverName":"www.qq.com","serverIp":"212.34.12.58"}]}`
  json.Unmarshal([]byte(str),&s)
  fmt.Println(s)
}
