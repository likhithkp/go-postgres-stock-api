package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load the env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB")
	return db
}
