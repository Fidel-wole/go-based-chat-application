package db

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	var err error
	DB, err = sql.Open("postgres", "postgresql://postgres:Fidelwole@27@localhost:5432/chat-app?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	log.Println("Successfully connected to database")
}
