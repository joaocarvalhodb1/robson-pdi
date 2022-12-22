package main

import (
	"fmt"
	"log"
	"net/http"
	"poc-pessoa/controller"
	"poc-pessoa/db"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn = "root:root@tcp(localhost:3306)/db_account"
)

func main() {
	conn, err := db.MysqlConnecttion(dsn)
	if err != nil {
		log.Println(err)
	}
	// migration
	err = db.CreateTables(conn)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	http.HandleFunc("/persons", controller.PersonHandler)
	fmt.Println("API server listening in port", 8888)
	log.Panic(http.ListenAndServe(":8888", nil))
}
