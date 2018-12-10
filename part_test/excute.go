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
 
  fmt.Printf("===========exec db ==========\n")

    //get stmt from db
	stmt, err := db.Prepare("INSERT INTO test_news(`title`,`desc`,`content`) VALUES(?,?,?)")
	errMsg(err)	
	defer stmt.Close()

    //get row from stmt
	res,err := stmt.Exec("golang insert data","study for use db model to insert data","in this lesson,we will try to use db to insert data into target")
  errMsg(err)
    
    //get last insert id
  lastId,err := res.LastInsertId()
  errMsg(err)
   
   //get affevted rows
  rowCnt, err := res.RowsAffected()
  errMsg(err)
   
   //out res
  fmt.Printf("insertId=>%d,RowsAffected=>%d\n",lastId,rowCnt)

}

