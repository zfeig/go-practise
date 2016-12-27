package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}

func init() {
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:root@/go?charset=utf8mb4", 30)
    
    // register model
    orm.RegisterModel(new(User))

    // create table
    orm.RunSyncdb("default", false, true)
}

func main() {
    o := orm.NewOrm()

    user := User{Name: "slene"}

    // insert
    id, err := o.Insert(&user)
    fmt.Printf("Insert ID: %d, ERR: %v\n", id, err)

    // update
    user.Name = "zfeig"
    num, err := o.Update(&user)
    fmt.Printf("Update Affeced NUM: %d, ERR: %v\n", num, err)

    // read one
    u := User{Id: user.Id}
    err = o.Read(&u)
    fmt.Printf("Read ERR: %v\n", err)

    // delete
    //num, err = o.Delete(&u)
    //fmt.Printf("Delete Affected NUM: %d, ERR: %v\n", num, err)
}
