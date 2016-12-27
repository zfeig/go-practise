package main

import(
  "fmt"
  "os"
)


func FileStatus( fileName string) (int64,bool,bool) {
  f,e := os.Stat(fileName)
  if e != nil {
   return 0,false,false
  }
  return f.Size(),true,!f.IsDir()
}


func main() {
  //fileName := "info.txt"
  fileName := "/home/lisi/demo/goproject/os/info.txt"
  size,exist,isFile :=  FileStatus(fileName)
  fmt.Println("file is exist?",exist,"  file size is:",size," is File?",isFile)
}
