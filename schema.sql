CREATE TABLE swift_codes (
    id SERIAL PRIMARY KEY,
    swift_code VARCHAR(255) NOT NULL,
    bank_name VARCHAR(255),
    address TEXT,
    country_iso2 VARCHAR(2),
    country_name VARCHAR(255),
    is_headquarter BOOLEAN
);

CREATE INDEX idx_swift_code ON swift_codes (swift_code);
CREATE INDEX idx_country_iso2 ON swift_codes (country_iso2);