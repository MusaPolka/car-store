CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INTEGER NOT NULL,
                        car_id INTEGER NOT NULL,
                        quantity INTEGER DEFAULT 1,
                        total NUMERIC(12, 2) NOT NULL,
                        status VARCHAR(50) DEFAULT 'pending',
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
