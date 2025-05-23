CREATE TABLE cars (
                      id SERIAL PRIMARY KEY,
                      name VARCHAR(100) NOT NULL,
                      description TEXT,
                      price NUMERIC(12, 2) NOT NULL,
                      carbrand_id INTEGER REFERENCES carbrands(id) ON DELETE CASCADE,
                      stock INTEGER DEFAULT 0,
                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);