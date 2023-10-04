package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitializePostgres() (*sql.DB, error) {
	fmt.Println("Init database here")

	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%v", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
		db.Close()

		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
		db.Close()

		return nil, err
	}

	fmt.Println("Connected to the PostgreSQL database")

	return db, nil
}
