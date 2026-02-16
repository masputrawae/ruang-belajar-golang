package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// connect db
func ConnectDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Error: ", err)
	}
	log.Println("✅ Success: MySQL connected")
	return db
}
