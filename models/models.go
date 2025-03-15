package models

type SwiftCode struct {
    SwiftCode    string `json:"swiftCode" db:"swift_code"`
    BankName     string `json:"bankName" db:"bank_name"`
    Address      string `json:"address" db:"address"`
    CountryISO2  string `json:"countryISO2" db:"country_iso2"`
    CountryName  string `json:"countryName" db:"country_name"`
    IsHeadquarter bool   `json:"isHeadquarter" db:"is_headquarter"`
}

type CountrySwiftCodes struct {
    CountryISO2 string      `json:"countryISO2"`
    CountryName string      `json:"countryName"`
    SwiftCodes  []SwiftCode `json:"swiftCodes"`
}