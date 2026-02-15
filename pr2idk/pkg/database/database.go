package database

import (
	"database/sql"
	"fmt"
	"log"
	"pr2idk/pkg/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New(c config.DB) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.Option,
	)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(1 * time.Minute)

	log.Println("👉️ MySQL connected...")
	return db
}
