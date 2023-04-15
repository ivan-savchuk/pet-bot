

CREATE TABLE weather (
    descr VARCHAR(128) NOT NULL,
    temparature FLOAT NOT NULL,
    feels_like FLOAT NOT NULL,
    pressure INT NOT NULL,
    humidity INT NOT NULL,
    wind_speed FLOAT NOT NULL,
    search_point POINT NOT NULL,
    weather_time TIMESTAMP NOT NULL
);
