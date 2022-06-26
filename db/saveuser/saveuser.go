package saveuser

import (
	"context"
	"log"

	. "test/db/conn"
)

func Saveuser(name, email string) {
	//opening database
	data, err := DbConnection() //create db instance
	ErrorCheck(err)
	var exists bool
	stmts := data.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM user WHERE name=?)", name)
	err = stmts.Scan(&exists)
	ErrorCheck(err)

	//prepare the statement to ensure no sql injection
	stmt, err := data.Prepare("INSERT INTO user(name, email) VALUES(?, ?)")
	ErrorCheck(err)

	//actually make the execution of the query
	res, err := stmt.Exec(name, email)
	ErrorCheck(err)

	//get last id to double check
	lastId, err := res.LastInsertId()
	ErrorCheck(err)

	//get rows affected to double check
	rowCnt, err := res.RowsAffected()
	ErrorCheck(err)

	//print out what you actually did
	log.Printf("lastid = %d, affected = %d, name = %s\n", lastId, rowCnt, name)
	defer data.Close()

}
