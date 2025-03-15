package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/marcinoPlayGames/swift-api/handlers"
    "github.com/marcinoPlayGames/swift-api/models"
)

/*

Normal unit tests cases

*/

// 1. Test for the GET /v1/swift-codes/:swiftCode endpoint
// Verifies that the endpoint returns the correct result when the provided swiftCode exists in the database
// and properly handles the case when the swiftCode does not exist.

func TestGetSwiftCode_Success(t *testing.T) {
    // Preparation: Set up a test server and test data
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)
    // Add sample data for testing (e.g., mocking the database)

    req, _ := http.NewRequest("GET", "/v1/swift-codes/BANKPLPWXXX", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    // Additional assertions can be added to verify the returned data
}

func TestGetSwiftCode_NotFound(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)

    req, _ := http.NewRequest("GET", "/v1/swift-codes/INVALIDCODE", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusNotFound, resp.Code)
    // Check if the response contains an appropriate error message
}

// 2. Test for the GET /v1/swift-codes/country/:countryISO2code endpoint
// Ensures that all SWIFT codes for a given country are correctly retrieved
// and checks the behavior when there are no codes available for that country.

func TestGetSwiftCodesByCountry_Success(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)

    req, _ := http.NewRequest("GET", "/v1/swift-codes/country/PL", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    // Verify that the response contains the correct SWIFT code data
}

func TestGetSwiftCodesByCountry_NotFound(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)

    req, _ := http.NewRequest("GET", "/v1/swift-codes/country/XY", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusNotFound, resp.Code)
    // Verify that the response contains an appropriate error message
}

// 3. Test for the POST /v1/swift-codes endpoint
// Tests the addition of a new SWIFT code to the system, including validation of input data and response handling.

func TestAddSwiftCode_Success(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    newCode := models.SwiftCode{
        SwiftCode:  "BANKNEWXXX",
        BankName:   "Bank Nowy",
        Address:    "Kraków",
        CountryISO2: "PL",
        CountryName: "Polska",
        IsHeadquarter: true,
    }

    jsonData, _ := json.Marshal(newCode)
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader(jsonData))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusCreated, resp.Code)
    // Additional assertions can be added to verify the response content
}

func TestAddSwiftCode_BadRequest(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    invalidCode := `{ "swiftCode": "BANKINVALIDXXX" }` // Invalid data structure
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader([]byte(invalidCode)))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// 4. Test for the DELETE /v1/swift-codes/:swiftCode endpoint
// Checks if the SWIFT code deletion works correctly and how the application reacts when deleting a non-existent code.

func TestDeleteSwiftCode_Success(t *testing.T) {
    r := gin.Default()
    r.DELETE("/v1/swift-codes/:swiftCode", handlers.DeleteSwiftCode)

    req, _ := http.NewRequest("DELETE", "/v1/swift-codes/BANKPLPWXXX", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteSwiftCode_NotFound(t *testing.T) {
    r := gin.Default()
    r.DELETE("/v1/swift-codes/:swiftCode", handlers.DeleteSwiftCode)

    req, _ := http.NewRequest("DELETE", "/v1/swift-codes/INVALIDCODE", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusNotFound, resp.Code)
}

/*

Edge cases unit tests

*/

// 1. Testing excessively short and long SWIFT codes in input data
// Ensures that the system correctly handles attempts to add a SWIFT code
// that is either too short or too long (not conforming to the required format).

func TestAddSwiftCode_ShortSwiftCode(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    invalidCode := models.SwiftCode{
        SwiftCode:  "BANK", // Too short SWIFT code
        BankName:   "Bank Nowy",
        Address:    "Kraków",
        CountryISO2: "PL",
        CountryName: "Polska",
        IsHeadquarter: true,
    }

    jsonData, _ := json.Marshal(invalidCode)
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader(jsonData))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusBadRequest, resp.Code)
}

func TestAddSwiftCode_LongSwiftCode(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    invalidCode := models.SwiftCode{
        SwiftCode:  "BANKNEWVERYLONGCODE12345", // Too long SWIFT code
        BankName:   "Bank Nowy",
        Address:    "Kraków",
        CountryISO2: "PL",
        CountryName: "Polska",
        IsHeadquarter: true,
    }

    jsonData, _ := json.Marshal(invalidCode)
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader(jsonData))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// 2. Testing multiple additions of the same SWIFT code
// Ensures that the system correctly handles attempts to add the same SWIFT code
// more than once, returning a 409 Conflict status if the code already exists in the system.

func TestAddSwiftCode_Duplicate(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    newCode := models.SwiftCode{
        SwiftCode:  "BANKNEWXXX",
        BankName:   "Bank Nowy",
        Address:    "Kraków",
        CountryISO2: "PL",
        CountryName: "Polska",
        IsHeadquarter: true,
    }

    jsonData, _ := json.Marshal(newCode)
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader(jsonData))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusCreated, resp.Code)

    // Attempting to add the same code again
    resp2 := httptest.NewRecorder()
    r.ServeHTTP(resp2, req)

    assert.Equal(t, http.StatusConflict, resp2.Code) // Returns 409 Conflict if it already exists
}

// 3. Missing data / Non-existent record
// Tests how the system responds when a user tries to access a SWIFT code
// that does not exist in the database.

func TestGetSwiftCode2_NotFound(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)

    // Using a non-existent SWIFT code
    req, _ := http.NewRequest("GET", "/v1/swift-codes/INVALIDCODE", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Checking if a 404 Not Found status is returned
    assert.Equal(t, http.StatusNotFound, resp.Code)
    // The response body can also be checked for an error message
}

// 4. Invalid input data format
// Tests cases where input data is in an incorrect format, such as invalid JSON or missing required fields.

func TestAddSwiftCode_BadRequest_InvalidJSON(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    // Niepoprawny JSON (brakująca część struktury)
    invalidCode := `{ "swiftCode": "BANKINVALID" }` // Invalid JSON (missing structure part)

    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader([]byte(invalidCode)))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Checking if a 400 Bad Request status is returned
    assert.Equal(t, http.StatusBadRequest, resp.Code)
    // The response body content can also be checked
}

func TestAddSwiftCode_BadRequest_InvalidData(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    // Invalid data, e.g., SWIFT code with an incorrect length
    invalidCode := models.SwiftCode{
        SwiftCode:  "BANK", // Too short SWIFT code
        BankName:   "Bank Nowy",
        Address:    "Kraków",
        CountryISO2: "PL",
        CountryName: "Polska",
        IsHeadquarter: true,
    }

    jsonData, _ := json.Marshal(invalidCode)
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader(jsonData))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Checking if a 400 Bad Request status is returned
    assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// 5. Incorrect URL parameters
// Tests cases where a user provides incorrect parameters in the URL,
// such as a non-existent country or an invalid SWIFT code.

func TestGetSwiftCodesByCountry_InvalidCountryISO2(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)

    // Using a non-existent country
    req, _ := http.NewRequest("GET", "/v1/swift-codes/country/XYZ", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Checking if a 404 Not Found status is returned
    assert.Equal(t, http.StatusNotFound, resp.Code)
    // The response body content can also be checked for an error message
}

func TestGetSwiftCode_InvalidSwiftCode(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)

    // Using an invalid SWIFT code
    req, _ := http.NewRequest("GET", "/v1/swift-codes/INVALIDCODE", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Checking if a 404 Not Found status is returned
    assert.Equal(t, http.StatusNotFound, resp.Code)
    // The response body content can also be checked
}

// 6. Invalid input data in requests
// Tests unexpected input data, such as HTML instead of JSON, which could cause errors.

func TestAddSwiftCode_InvalidContentType(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    // Sending HTML instead of JSON
    invalidContent := "<html><body>Error</body></html>"
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader([]byte(invalidContent)))
    req.Header.Set("Content-Type", "text/html") // Incorrect Content-Type header

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Checking if a 415 Unsupported Media Type status is returned
    assert.Equal(t, http.StatusUnsupportedMediaType, resp.Code)
    // The response body content can also be checked
}

func TestGetSwiftCode_InvalidRequestBody(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)

    // Sending an empty request body (even though the GET endpoint should not require one)
    req, _ := http.NewRequest("GET", "/v1/swift-codes/BANKPLPWXXX", nil)
    req.Header.Set("Content-Type", "application/json") // Empty body, but with Content-Type set to JSON

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Checking if the response is valid even without a request body
    assert.Equal(t, http.StatusOK, resp.Code)
    // Also, ensuring that no errors occur in the response
}