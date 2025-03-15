package main

import (
    "github.com/gin-gonic/gin"
    "github.com/marcinoPlayGames/swift-api/database"
    "github.com/marcinoPlayGames/swift-api/handlers"  // Add this import
    "github.com/marcinoPlayGames/swift-api/parser"
    "github.com/marcinoPlayGames/swift-api/models"
    "log"
)

var swiftCodes []models.SwiftCode

func main() {
    // Connecting to the database
    database.ConnectDB()

    // Parsing CSV and assigning results to the swiftCodes variable
    var err error
    swiftCodes, err = parser.ParseCSV("assets/Interns_2025_SWIFT_CODES.csv")
    if err != nil {
        panic(err)
    }

    // Inserting data into the database
    for _, code := range swiftCodes {
        err := database.InsertSwiftCode(code)  // Function for inserting data into the database
        if err != nil {
            log.Fatal("Error inserting data: ", err)
        }
    }

    r := gin.Default()

    // Defining API routes
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)
    r.DELETE("/v1/swift-codes/:swiftCode", handlers.DeleteSwiftCode)

    r.Run(":8080")  // Running the server on port 8080
}