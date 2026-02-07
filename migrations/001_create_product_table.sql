CREATE TABLE IF NOT EXISTS product (
    id INTEGER PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    image VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    unit_of_measure VARCHAR(128) NOT NULL,
    total_qty INTEGER NOT NULL,
    description JSONB,
    rating INTEGER,
    rating_count INTEGER,
    tag VARCHAR(128)
);
