package main

import(
	"fmt"
	"database/sql"
	 _ "github.com/go-sql-driver/mysql"
)
const (
  DB_TYPE = "mysql"
  DB_ADDR = "root:@tcp(127.0.0.1:3306)/test"
)

var (
		id int
		title string
)

func errMsg(e error) {
	if e !=nil {
		panic(e.Error)
	}
}

func main() {
	//connect db
	db, err := sql.Open(DB_TYPE, DB_ADDR)
	errMsg(err)	
	defer db.Close()

    //get data from db
	rows, err := db.Query("select id, title from title where id = ?", 2)
	errMsg(err)	
	defer rows.Close()

    //get rows from res
	for rows.Next() {
		err := rows.Scan(&id, &title)
		errMsg(err)	
		fmt.Printf("id =>%d,title=>%s",id, title)
	}
   
   //process error handling
	err = rows.Err()
	errMsg(err)	
}