# swift-api
This is the repository for the swift-api home exercise.

## Swift API

### Jak uruchomić aplikację:

1. **Lokalnie**:
   1. Zainstaluj Go (wersja 1.21 lub nowsza).
   2. Sklonuj repozytorium:
      ```bash
      git clone https://github.com/yourusername/swift-api.git
      cd swift-api
      ```
   3. Zainstaluj zależności:
      ```bash
      go mod tidy
      ```
   4. Uruchom komendę:
      ```bash
      go run main.go
      ```
   5. Aplikacja będzie dostępna pod `http://localhost:8080`.

2. **W Dockerze**:
   1. Zainstaluj Docker i Docker Compose.
   2. Sklonuj repozytorium:
      ```bash
      git clone https://github.com/yourusername/swift-api.git
      cd swift-api
      ```
   3. Uruchom aplikację i bazę danych w kontenerach:
      ```bash
      docker-compose up --build
      ```
   4. Aplikacja będzie dostępna pod `http://localhost:8080`.

### Jak testować API:

Przykład zapytań cURL:

- **GET**: Pobierz dane kodu SWIFT:
  ```bash
  curl http://localhost:8080/v1/swift-codes/XXXXX

POST: Dodaj nowy kod SWIFT:

curl -X POST -d '{"swiftCode": "XXXXX", "bankName": "Bank", "address": "Address", "countryISO2": "US", "countryName": "USA", "isHeadquarter": true}' http://localhost:8080/v1/swift-codes

DELETE: Usuń kod SWIFT:

curl -X DELETE http://localhost:8080/v1/swift-codes/XXXXX

Przykład odpowiedzi:

GET:

{
  "swiftCode": "XXXXX",
  "bankName": "Bank",
  "address": "Address",
  "countryISO2": "US",
  "countryName": "USA",
  "isHeadquarter": true
}

POST: Odpowiedź w przypadku powodzenia:

{
  "message": "Swift code successfully added"
}

DELETE: Odpowiedź w przypadku powodzenia:

{
  "message": "Swift code successfully deleted"
}

Testowanie jednostkowe
Testy jednostkowe dla logiki biznesowej i endpointów są zawarte w repozytorium. Można je uruchomić za pomocą narzędzia do testowania w Go, np. go test:

go test ./…

Zależności

Go 1.21 lub nowsza.
PostgreSQL (lub Docker z kontenerem PostgreSQL).

Podsumowanie
Aplikacja umożliwia zarządzanie danymi kodów SWIFT poprzez API. Możesz uruchomić ją lokalnie lub w kontenerze Docker, a także testować różne endpointy przy użyciu cURL lub Postmana.

### Co dodano:
1. **Instalacja i konfiguracja Docker** – szczegóły dotyczące uruchamiania aplikacji i bazy danych w kontenerach.
2. **Więcej szczegółów na temat testów jednostkowych** – jak uruchomić testy w Go.
3. **Przykłady odpowiedzi API** – użytkownik wie, czego się spodziewać po wysłaniu zapytania.
4. **Instrukcje dotyczące zależności** – Go 1.21, PostgreSQL oraz Docker.

To sprawia, że README.md będzie kompletne i pełne, umożliwiając łatwe uruchomienie i testowanie aplikacji.