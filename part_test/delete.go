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
 
  fmt.Printf("===========delete data from db ==========\n")

    //get stmt from db
  stmt, err := db.Prepare("DELETE FROM test_news WHERE id=?")
  errMsg(err)	
  defer stmt.Close()

    //get row from stmt
  id := 1
  res,err := stmt.Exec(id)
  errMsg(err)
    
   //get affevted rows
  rowCnt, err := res.RowsAffected()
  errMsg(err)
   
   //out res
  fmt.Printf("delete %d,RowsAffected=>%d\n",id,rowCnt)

}

