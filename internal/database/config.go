package database

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func Init() error {
	var err error

	DB, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetPostgres() *sql.DB {
	return DB
}
