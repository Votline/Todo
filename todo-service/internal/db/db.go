package db

import (
	"os"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func InitDB() *sqlx.DB {
	var err error
	var db *sqlx.DB
	connStr := os.Getenv("POSTGRES_URL")
	
	for i := 0; i < 10; i++ {
		db, err = sqlx.Connect("postgres", connStr)
		if err != nil {
			log.Println("Failed to connect to DB", err)
		} else {
			return db
		}
		log.Println("Waiting for DB. \n")
		time.Sleep(2*time.Second)
	}
	return db
}
