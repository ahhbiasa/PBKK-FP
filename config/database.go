package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:Xadenth04*@tcp(localhost:3306)/go_crud?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
	DB = db
}
