

CREATE TABLE cities (
    city_id SERIAL PRIMARY KEY,
    city_name VARCHAR(32) NOT NULL,
    city_state VARCHAR(64) NULL,
    country VARCHAR(64) UNIQUE NOT NULL,
    city_location POINT NOT NULL
);
