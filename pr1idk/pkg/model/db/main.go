package db

import (
	"database/sql"
	"fmt"
	"log"
	"pr1idk/pkg/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Repository interface {
}

type API struct {
}

func New(c model.DBConfig) *sql.DB {
	opt := "?parseTime=true&charset=utf8mb4&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		opt,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("\n/pkg/model/db/main.go\nError in line [30]:\n", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("\n/pkg/model/db/main.go\nError in line [34]:\n", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}
