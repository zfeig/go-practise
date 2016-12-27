package main

import(
  "fmt"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

func main(){
  db, err := sql.Open("mysql", "root:root@/weather")
  if err != nil {
    panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
      panic(err.Error()) // proper error handling instead of panic in your app
  }
  
  fmt.Printf("db connect ok!\n")
  fmt.Printf("-----------------------------------\n")

  //query data and print

  rows,err:=db.Query("SELECT * FROM weather_city WHERE id='CN101010300'")
  if err !=nil {
    fmt.Printf("parse sql error!")
  }

   columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
   }


   values := make([]sql.RawBytes, len(columns))

   scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

   
     for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        // Now do something with the data.
        // Here we just print each column as a string.
        var value string
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }
            fmt.Println(columns[i], ": ", value)
        }
        fmt.Println("-----------------------------------")
    }


    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }  

}
