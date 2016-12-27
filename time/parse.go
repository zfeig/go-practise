package main

import(
  "fmt"
  "time"
)

func main(){
  const tm = "Jan 12,2016 at 6:30pm (UTC)"
  t,_ := time.Parse(tm,"2010-10-10 6:45:20")
  fmt.Println("foramt time is:",t) 

}
