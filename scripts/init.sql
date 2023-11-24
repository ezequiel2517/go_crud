\c postgres;

CREATE TABLE IF NOT EXISTS drugs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    approved BOOLEAN,
    min_dose INTEGER CHECK (min_dose >= 0),
    max_dose INTEGER CHECK (max_dose >= 0),
    available_at DATE
);

CREATE TABLE IF NOT EXISTS usuario (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
	email VARCHAR(50),
	password VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS vaccination (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
	drug_id INTEGER,
	dose INTEGER,
	fecha DATE,
	CONSTRAINT fk_drug_id
		FOREIGN KEY(drug_id) 
		REFERENCES drugs(id)
);

