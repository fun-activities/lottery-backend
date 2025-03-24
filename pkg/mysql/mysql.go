package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	MaxOpenConns int
}

func Open(dsn string) (*sqlx.DB, error) {
	con, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Mysql connect fail %w", err)
	}

	con.SetMaxOpenConns(5)
	con.SetMaxIdleConns(5)
	con.SetConnMaxIdleTime(15 * time.Minute)

	if err = con.Ping(); err != nil {
		return nil, fmt.Errorf("Mysql ping fail %w", err)
	}
	return con, nil

}

func Close(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}
