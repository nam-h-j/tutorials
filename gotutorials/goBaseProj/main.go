package main

import (
	"database/sql"
	"fmt"
	"goBaseProj/router"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

// @title     GO BASE PROJ
// @version         1.0
// @description     GO BASE PROJ API

// @securityDefinitions.apikey user-token
// @in header
// @name Authorization

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	//dbInit
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PWD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "go_crud",
		AllowNativePasswords: true,
	}

	// 디비 연결
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 디비 연결 확인
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to mysql version: ", version)

	// r := router.Router(db)
	// router.Run(":1234")

	http.ListenAndServe(":1234", router.Router(db))
}
