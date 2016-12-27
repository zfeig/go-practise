package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

type Article struct {
   Id int
   Title   string
   Content  string
   Author  string
   Cid   int
   Brief  string
   Add_time string
   Key   string
   Weburl  string
}


/**
func (a *Article) TableName() string {
  return "article"
}**/


func init() {
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:root@/go?charset=utf8mb4", 30)
    orm.RegisterModel(new(Article))
}

func main() {
    fmt.Println("==============Raw QUery===================")  
    o := orm.NewOrm()    
    var maps  []orm.Params
    num,err := o.Raw("SELECT * FROM `article` WHERE author= ?","zfeig").Values(&maps)
    if err==nil && num>0 {
      for _,v:= range maps {
         fmt.Println("Id:",v["id"],";title:",v["title"])
      } 
    }else{
         fmt.Println("not find!")
    }

   
   /**
    res,err := o.Raw("UPDATE `article` SET cid =?  WHERE id in (?,?)",3,[]int{12,13}).Exec()
    if err == nil {
      num,_ := res.RowsAffected()
      fmt.Println("update row affected nums:",num)
    }else{
      fmt.Println("query error")
    }
   **/

   
    fmt.Println("============QueryTable===================")
    var datas []orm.Params
    nums, errs := o.QueryTable("article").Filter("id__gt",9).Values(&datas)
    if errs == nil {
        fmt.Printf("Result Nums: %d\n", nums)
        for _, m := range datas {
            fmt.Println(m["Id"], m["Title"],m["Add_time"])
        }
    }

    
}
