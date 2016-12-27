package main
import(
 "html/template"
 "os"
)

type Person struct{
  UserName string
}

func main(){
  t := template.New("tpl")
  t,_ = t.Parse("hello {{.UserName}}!\n")
  p:= Person{UserName:"zfeig"}
  t.Execute(os.Stdout,p)
}
