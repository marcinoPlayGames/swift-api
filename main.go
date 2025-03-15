package main

import (
    "github.com/gin-gonic/gin"
    "github.com/marcinoPlayGames/swift-api/database"
    "github.com/marcinoPlayGames/swift-api/handlers"  // Dodaj ten import
    "github.com/marcinoPlayGames/swift-api/parser"
    "github.com/marcinoPlayGames/swift-api/models"
    "log"
)

var swiftCodes []models.SwiftCode

func main() {
    // Połączenie z bazą danych
    database.ConnectDB()

    // Parsowanie CSV i przypisanie wyników do zmiennej swiftCodes
    var err error
    swiftCodes, err = parser.ParseCSV("assets/Interns_2025_SWIFT_CODES.csv")
    if err != nil {
        panic(err)
    }

    // Wstawianie danych do bazy
    for _, code := range swiftCodes {
        err := database.InsertSwiftCode(code)  // Funkcja wstawiająca dane do bazy
        if err != nil {
            log.Fatal("Błąd przy wstawianiu danych: ", err)
        }
    }

    r := gin.Default()

    // Definiowanie tras API
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)
    r.DELETE("/v1/swift-codes/:swiftCode", handlers.DeleteSwiftCode)

    r.Run(":8080")  // Uruchomienie serwera na porcie 8080
}