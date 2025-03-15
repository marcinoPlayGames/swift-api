package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
    var err error
    connStr := "user=postgres password=admin dbname=swift sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }
    fmt.Println("Connected to database")
}