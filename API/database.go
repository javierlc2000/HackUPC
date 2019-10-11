package main

import (
	"fmt"
	_ "mysql-master"
	"database/sql"
)


func main(){
	fmt.Println("starting")

	db, err := sql.Open("mysql",
		"user:password@tcp(10.192.70.129:3306)/hello")

	fmt.Println("connecting")


	if err != nil {
		fmt.Println("rejected")
		panic(err)
	}
	defer db.Close()

	fmt.Println("here we are")

	name := "app_user"


	_,err = db.Exec("CREATE DATABASE "+name)
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("USE "+name)
   if err != nil {
       panic(err)
   }

   _,err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
   if err != nil {
       panic(err)
   }


   print(db)
}
