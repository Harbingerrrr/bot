package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect(connectionString string) {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening PostgreSQL database: %q", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging PostgreSQL database: %q", err)
	}

	fmt.Println("Connected to PostgreSQL database")
}

func Close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing PostgreSQL database: %q", err)
		}
		fmt.Println("Closed database PostgreSQL connection")
	}
}
