package getall

import (
	"fmt"

	. "test/db/conn"
)

func Getall() []User {

	data, err := DbConnection() //create db instance
	ErrorCheck(err)

	//variables used to store data from the query
	var (
		id    string
		name  string
		email string
		Users []User //used to store all users
	)
	i := 0 //used to get how many scans

	//get from database
	rows, err := data.Query("select * from user")
	ErrorCheck(err)

	//cycle through the rows to collect all the data
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		ErrorCheck(err)

		i++
		fmt.Println("scan ", i)

		//store into memory
		u := User{ID: id, Name: name, Email: email}
		Users = append(Users, u)

	}
	//close everything
	defer rows.Close()
	defer data.Close()
	return Users

}

type User struct {
	ID    string `json:"id" form:"id"`
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required,email"`
}
