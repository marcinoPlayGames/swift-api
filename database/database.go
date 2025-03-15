package database

import (
    "fmt"
    "io/ioutil"
    "github.com/jmoiron/sqlx"
    "github.com/marcinoPlayGames/swift-api/models"
    "log"
    _ "github.com/lib/pq"
)

var db *sqlx.DB

// ConnectDB - funkcja łącząca się z bazą danych
func ConnectDB() {
    var err error
    db, err = sqlx.Connect("postgres", "user=postgres password=admin dbname=swift host=db sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}

// ExecuteSchema - funkcja wykonująca schema.sql
func ExecuteSchema() {
    // Odczytujemy plik schema.sql
    data, err := ioutil.ReadFile("schema.sql")
    if err != nil {
        log.Fatal(err)
    }

    // Wykonujemy zapytanie SQL z pliku
    _, err = db.Exec(string(data))
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Schema applied successfully!")
}

// InsertSwiftCode - funkcja wstawiająca rekord SwiftCode do bazy danych
func InsertSwiftCode(code models.SwiftCode) error {
    query := `INSERT INTO swift_codes (swift_code, bank_name, address, country_iso2, country_name, is_headquarter)
              VALUES (:swift_code, :bank_name, :address, :country_iso2, :country_name, :is_headquarter)`
    
    _, err := db.NamedExec(query, code)
    return err
}