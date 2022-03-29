package main

import (
	"go-learn/db"
	"go-learn/router"

	_ "github.com/go-sql-driver/mysql"
)



func main() {
	dbConnect := db.InitDB()
	
	router.Router(dbConnect)
}
