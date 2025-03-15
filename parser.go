package parser

import (
    "encoding/csv"
    "os"
    "strings"

    "github.com/marcinoPlayGames/swift-api/models"
)

func ParseCSV(filename string) ([]models.SwiftCode, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var swiftCodes []models.SwiftCode
    for _, record := range records[1:] { // Pomijamy nagłówki
        isHeadquarter := strings.HasSuffix(record[0], "XXX")
        swiftCodes = append(swiftCodes, models.SwiftCode{
            SwiftCode:    record[0],
            BankName:     record[1],
            Address:      record[2],
            CountryISO2:  strings.ToUpper(record[3]),
            CountryName:  strings.ToUpper(record[4]),
            IsHeadquarter: isHeadquarter,
        })
    }
    return swiftCodes, nil
}