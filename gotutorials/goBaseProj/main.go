package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

	router := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r := router.Router(db)

	// r.Run(":1234")

	// http.ListenAndServe(":1234", router.Router(db))
}
