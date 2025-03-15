package handlers

import (
	"encoding/json"  // Użycie tej biblioteki do kodowania odpowiedzi
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/marcinoPlayGames/swift-api/models"
)

// Pobiera kod SWIFT po identyfikatorze
func GetSwiftCode(c *gin.Context) {
	swiftCode := c.Param("swiftCode")
	for _, code := range swiftCodes {
		if code.SwiftCode == swiftCode {
			c.Header("Content-Type", "application/json")  // Ustalenie nagłówka odpowiedzi
			json.NewEncoder(c.Writer).Encode(code)       // Użycie JSON do zakodowania odpowiedzi
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Swift code not found"})
}

// Przykładowa struktura przechowująca dane SWIFT
var swiftCodes = []models.SwiftCode{
	{SwiftCode: "BANKPLPWXXX", BankName: "Bank Polska", Address: "Warszawa", CountryISO2: "PL", CountryName: "Polska", IsHeadquarter: true},
	{SwiftCode: "BANKDEFFXXX", BankName: "Bank Niemcy", Address: "Berlin", CountryISO2: "DE", CountryName: "Niemcy", IsHeadquarter: true},
}

// Pobiera wszystkie kody SWIFT dla danego kraju
func GetSwiftCodesByCountry(c *gin.Context) {
	countryISO2 := c.Param("countryISO2code")
	var countryCodes []models.SwiftCode
	for _, code := range swiftCodes {
		if code.CountryISO2 == countryISO2 {
			countryCodes = append(countryCodes, code)
		}
	}

	if len(countryCodes) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No Swift codes found for this country"})
		return
	}

	c.JSON(http.StatusOK, countryCodes)
}

// Dodaje nowy kod SWIFT
func AddSwiftCode(c *gin.Context) {
	var newCode models.SwiftCode
	if err := c.ShouldBindJSON(&newCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	swiftCodes = append(swiftCodes, newCode)
	c.JSON(http.StatusCreated, newCode)
}

// Usuwa kod SWIFT
func DeleteSwiftCode(c *gin.Context) {
	swiftCode := c.Param("swiftCode")
	for i, code := range swiftCodes {
		if code.SwiftCode == swiftCode {
			swiftCodes = append(swiftCodes[:i], swiftCodes[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Swift code deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Swift code not found"})
}