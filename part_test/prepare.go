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
 
    fmt.Printf("===========get row ==========\n")

    //get stmt from db
	stmt, err := db.Prepare("select id, title from test_news where id = ?")
	errMsg(err)	
	defer stmt.Close()

    //get row from stmt
	err = stmt.QueryRow(2).Scan(&id,&title)
    errMsg(err)
   
   //out res
   fmt.Printf("id=>%d,title=>%s\n",id,title)

  //======================================================
  fmt.Printf("===========get more ==========\n")

  //get stmt obj  
  sdb, err := db.Prepare("select id, title from test_news where id > ?")
  errMsg(err)	
  defer sdb.Close()

  //get query obj
  rows,err := sdb.Query(0)
  errMsg(err)
  defer rows.Close()

  //itera data and output
  for rows.Next() {
  	err := rows.Scan(&id, &title)
	errMsg(err)
	fmt.Printf("row:id=>%d,title=>%s\n",id,title)
  }
  
  //handing error
  err = rows.Err()
  errMsg(err)

}

