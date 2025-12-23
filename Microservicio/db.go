package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func conectarDB() {
	var err error
	dsn := "usuariodb:1234@tcp(127.0.0.1:3306)/proyecto_go"
	db, err = sql.Open("mysql", dsn)
	if err != nil || db.Ping() != nil {
		fmt.Println("Error conectando a la base de datos")
	}
}