package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	env "github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	// dns := "knox:knox@tcp(localhost:3306)/db_contacts"
	err := env.Load()
	if err != nil {
		return nil, err
	}

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dns)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión a la db exitosa")
	return db, nil
}
