package database

import (
	"database/sql"
	"fmt"
	"pr1idk/pkg/config"
	"pr1idk/pkg/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New(c config.DB) *sql.DB {
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
		helper.ErrFatal(err)
	}

	if err = db.Ping(); err != nil {
		helper.ErrFatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}
