package handlers

import (
	"encoding/json"  // Using this library for encoding responses
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/marcinoPlayGames/swift-api/models"
)

// Retrieves a SWIFT code by its identifier
func GetSwiftCode(c *gin.Context) {
	swiftCode := c.Param("swiftCode")
	for _, code := range swiftCodes {
		if code.SwiftCode == swiftCode {
			c.Header("Content-Type", "application/json")  // Setting the response header
			json.NewEncoder(c.Writer).Encode(code)       // Using JSON to encode the response
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Swift code not found"})
}

// Example structure for storing SWIFT data
var swiftCodes = []models.SwiftCode{
	{SwiftCode: "BANKPLPWXXX", BankName: "Bank Polska", Address: "Warszawa", CountryISO2: "PL", CountryName: "Polska", IsHeadquarter: true},
	{SwiftCode: "BANKDEFFXXX", BankName: "Bank Niemcy", Address: "Berlin", CountryISO2: "DE", CountryName: "Niemcy", IsHeadquarter: true},
}

// Retrieves all SWIFT codes for a given country
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

// Adds a new SWIFT code
func AddSwiftCode(c *gin.Context) {
	var newCode models.SwiftCode
	if err := c.ShouldBindJSON(&newCode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	swiftCodes = append(swiftCodes, newCode)
	c.JSON(http.StatusCreated, newCode)
}

// Deletes a SWIFT code
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