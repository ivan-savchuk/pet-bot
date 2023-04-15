
CREATE TABLE chats (
    chat_id VARCHAR(64) PRIMARY KEY,
    city_id VARCHAR(32) NULL,
    last_activity TIMESTAMP NOT NULL,
    FOREIGN KEY (city_id) REFERENCES cities (city_id)
);

CREATE TABLE chat_functions (
    chat_id VARCHAR(64) PRIMARY KEY,
    weather_daily BOOLEAN NULL,
    aqi_daily BOOLEAN NULL,
    FOREIGN KEY (chat_id) REFERENCES chats (chat_id)
);