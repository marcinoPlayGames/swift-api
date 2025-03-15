CREATE TABLE swift_codes (
    id SERIAL PRIMARY KEY,
    swift_code VARCHAR(11) UNIQUE NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    country_iso2 CHAR(2) NOT NULL,
    country_name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    is_headquarter BOOLEAN NOT NULL
);

CREATE INDEX idx_swift_code ON swift_codes (swift_code);
CREATE INDEX idx_country_iso2 ON swift_codes (country_iso2);