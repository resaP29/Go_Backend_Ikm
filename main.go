package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main(){

	fmt.Println("connet to database")
	db, err := sql.Open("mysql", "root:passwort@tcp(127.0.0.1:3306)/testdb")

	if err!= null {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("berhasil koneksi")
}