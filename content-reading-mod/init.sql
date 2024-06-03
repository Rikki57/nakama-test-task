CREATE TABLE IF NOT EXISTS files (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50),
    version VARCHAR(50),
    hash VARCHAR(64),  
    content TEXT
);