package main

import (
    "github.com/gin-gonic/gin"
    "github.com/yourusername/swift-api/database"
    "github.com/yourusername/swift-api/handlers"
)

func main() {
    database.ConnectDB()
    
    r := gin.Default()

    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)
    r.DELETE("/v1/swift-codes/:swiftCode", handlers.DeleteSwiftCode)

    r.Run(":8080")
}