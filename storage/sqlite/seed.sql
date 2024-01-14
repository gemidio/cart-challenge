CREATE TABLE products_offers (
    id INTEGER PRIMARY KEY,
    product_id TEXT,
    expire_at TEXT
);

INSERT INTO products_offers VALUES (1, "6d5544f4-8dfe-4d11-9dca-4160ff265893", "2024-01-01 12:00:00.000");
INSERT INTO products_offers VALUES (2, "e16fefdd-dd32-424a-a4dd-466dc86613d3", "2024-02-01 23:59:59.000");