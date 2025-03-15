// 1. Test dla endpointu GET /v1/swift-codes/:swiftCode
// Testowanie, czy zwraca poprawny wynik, kiedy podany swiftCode jest obecny w bazie danych, oraz czy poprawnie obsługuje sytuację, gdy // swiftCode nie istnieje.

func TestGetSwiftCode_Success(t *testing.T) {
    // Przygotowanie: Tworzymy serwer testowy i dane
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)
    // Dodanie przykładowych danych do testu (np. mockowanie bazy danych)

    req, _ := http.NewRequest("GET", "/v1/swift-codes/BANKPLPWXXX", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    // Możesz dodać więcej asercji sprawdzających zwrócone dane
}

func TestGetSwiftCode_NotFound(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)

    req, _ := http.NewRequest("GET", "/v1/swift-codes/INVALIDCODE", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusNotFound, resp.Code)
    // Sprawdzanie, czy w odpowiedzi jest komunikat o błędzie
}

// 2. Test dla endpointu GET /v1/swift-codes/country/:countryISO2code
// Sprawdzenie, czy zwraca wszystkie kody SWIFT dla danego kraju, oraz co się dzieje, gdy nie ma kodów dla tego kraju.

func TestGetSwiftCodesByCountry_Success(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)

    req, _ := http.NewRequest("GET", "/v1/swift-codes/country/PL", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    // Sprawdzanie, czy odpowiedź zawiera poprawne dane o kodach SWIFT
}

func TestGetSwiftCodesByCountry_NotFound(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)

    req, _ := http.NewRequest("GET", "/v1/swift-codes/country/XY", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusNotFound, resp.Code)
    // Sprawdzanie, czy odpowiedź zawiera odpowiedni komunikat o błędzie
}

// 3. Test dla endpointu POST /v1/swift-codes
// Testowanie dodawania nowego kodu SWIFT do systemu, w tym sprawdzenie poprawności danych wejściowych i odpowiedzi.

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
    // Możesz dodać więcej asercji na sprawdzenie zawartości odpowiedzi
}

func TestAddSwiftCode_BadRequest(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    invalidCode := `{ "swiftCode": "BANKINVALIDXXX" }` // Zła struktura
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader([]byte(invalidCode)))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// 4. Test dla endpointu DELETE /v1/swift-codes/:swiftCode
// Sprawdzenie, czy usuwanie kodu SWIFT działa poprawnie i jak aplikacja reaguje na nieistniejący kod.

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

// 5. Testowanie za krótkiego i za długiego kodu SWIFT w danych wejściowych
// Test sprawdza, czy system poprawnie reaguje na próby dodania kodu SWIFT, 
// który jest zbyt krótki lub zbyt długi (niezgodny z wymaganiami formatu).

func TestAddSwiftCode_ShortSwiftCode(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    invalidCode := models.SwiftCode{
        SwiftCode:  "BANK", // Za krótki kod SWIFT
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
        SwiftCode:  "BANKNEWVERYLONGCODE12345", // Za długi kod SWIFT
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

// 6. Testowanie wielokrotnego dodawania tego samego kodu SWIFT
// Test sprawdza, czy system poprawnie obsługuje próbę dodania tego samego kodu SWIFT 
// więcej niż raz, zwracając odpowiedni status 409 Conflict, jeśli kod już istnieje w systemie.

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

    // Próba dodania tego samego kodu ponownie
    resp2 := httptest.NewRecorder()
    r.ServeHTTP(resp2, req)

    assert.Equal(t, http.StatusConflict, resp2.Code) // Zwraca 409 Conflict jeśli już istnieje
}

// 7. Brak danych / Nieistniejący rekord
// Testowanie sytuacji, gdy użytkownik próbuje uzyskać dostęp do kodu SWIFT, który nie istnieje w bazie danych.

func TestGetSwiftCode_NotFound(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)

    // Używamy nieistniejącego kodu SWIFT
    req, _ := http.NewRequest("GET", "/v1/swift-codes/INVALIDCODE", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Sprawdzamy, czy zwrócono status 404 Not Found
    assert.Equal(t, http.StatusNotFound, resp.Code)
    // Możemy również sprawdzić treść odpowiedzi, np. komunikat o błędzie
}

// 8. Nieprawidłowy format danych wejściowych
// Testowanie sytuacji, gdzie dane wejściowe mają niepoprawny format, np. zły JSON lub brak wymaganych pól.

func TestAddSwiftCode_BadRequest_InvalidJSON(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    // Niepoprawny JSON (brakująca część struktury)
    invalidCode := `{ "swiftCode": "BANKINVALID" }` // Brak wymaganych pól, np. BankName

    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader([]byte(invalidCode)))
    req.Header.Set("Content-Type", "application/json")

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Sprawdzamy, czy zwrócono status 400 Bad Request
    assert.Equal(t, http.StatusBadRequest, resp.Code)
    // Można dodać również sprawdzenie treści odpowiedzi
}

func TestAddSwiftCode_BadRequest_InvalidData(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    // Niepoprawne dane, np. kod SWIFT z niewłaściwą długością
    invalidCode := models.SwiftCode{
        SwiftCode:  "BANK", // Za krótki kod SWIFT
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

    // Sprawdzamy, czy zwrócono status 400 Bad Request
    assert.Equal(t, http.StatusBadRequest, resp.Code)
}

// 9. Złe parametry w URL
// Testowanie sytuacji, w której użytkownik poda niepoprawne parametry w URL, np. nieistniejący kraj lub błędny kod SWIFT.

func TestGetSwiftCodesByCountry_InvalidCountryISO2(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/country/:countryISO2code", handlers.GetSwiftCodesByCountry)

    // Używamy nieistniejącego kraju
    req, _ := http.NewRequest("GET", "/v1/swift-codes/country/XYZ", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Sprawdzamy, czy zwrócono status 404 Not Found
    assert.Equal(t, http.StatusNotFound, resp.Code)
    // Można także sprawdzić treść odpowiedzi, np. komunikat o błędzie
}

func TestGetSwiftCode_InvalidSwiftCode(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)

    // Używamy błędnego kodu SWIFT
    req, _ := http.NewRequest("GET", "/v1/swift-codes/INVALIDCODE", nil)
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Sprawdzamy, czy zwrócono status 404 Not Found
    assert.Equal(t, http.StatusNotFound, resp.Code)
    // Można także sprawdzić treść odpowiedzi
}

// 10. Błędne dane wejściowe w zapytaniach
// Testowanie nieoczekiwanych danych wejściowych, np. tekst HTML zamiast JSON, co może prowadzić do błędów.

func TestAddSwiftCode_InvalidContentType(t *testing.T) {
    r := gin.Default()
    r.POST("/v1/swift-codes", handlers.AddSwiftCode)

    // Wysyłamy HTML zamiast JSON
    invalidContent := "<html><body>Error</body></html>"
    req, _ := http.NewRequest("POST", "/v1/swift-codes", bytes.NewReader([]byte(invalidContent)))
    req.Header.Set("Content-Type", "text/html") // Zły nagłówek Content-Type

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Sprawdzamy, czy zwrócono status 415 Unsupported Media Type
    assert.Equal(t, http.StatusUnsupportedMediaType, resp.Code)
    // Można także sprawdzić treść odpowiedzi
}

func TestGetSwiftCode_InvalidRequestBody(t *testing.T) {
    r := gin.Default()
    r.GET("/v1/swift-codes/:swiftCode", handlers.GetSwiftCode)

    // Wysyłamy pusty request body (chociaż endpoint GET nie powinien go wymagać)
    req, _ := http.NewRequest("GET", "/v1/swift-codes/BANKPLPWXXX", nil)
    req.Header.Set("Content-Type", "application/json") // Pusty body, ale z Content-Type jako JSON

    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    // Sprawdzamy, czy odpowiedź jest poprawna, nawet bez ciała
    assert.Equal(t, http.StatusOK, resp.Code)
    // Możemy również sprawdzić, czy nie wystąpił żaden błąd w odpowiedzi
}