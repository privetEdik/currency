DROP TABLE IF EXISTS currencies;
CREATE TABLE IF NOT EXISTS currencies
(
    id   SERIAL PRIMARY KEY,
    name TEXT              NOT NULL,
    code VARCHAR(3) UNIQUE NOT NULL,
    sign TEXT              NOT NULL
);