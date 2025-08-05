CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100) NOT NULL UNIQUE,
                       password_hash VARCHAR(255) NOT NULL,
                       email VARCHAR(150) UNIQUE,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);