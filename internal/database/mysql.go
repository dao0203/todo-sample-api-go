package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
	*sql.DB
}

func New() (*MySql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "password", "localhost", "3306", "todo")

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MySql{DB: db}, nil
}

func (ms *MySql) Close() {
	ms.DB.Close()
}
