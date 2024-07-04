--CREATE gardens DATABASE
CREATE DATABASE gardens;

-- CREATE ENUM types
CREATE TYPE garden_type AS ENUM ('balcony', 'rooftop', 'indoor', 'community', 'backyard');
CREATE TYPE plant_status AS ENUM ('planned', 'planted', 'growing', 'harvesting', 'dormant');

-- CREATE gardens TABLE
CREATE TABLE gardens (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    name VARCHAR(100) NOT NULL,
    type garden_type,
    area_sqm DECIMAL(6, 2),
    created_at VARCHAR,
    updated_at VARCHAR,
    deleted_at VARCHAR
);

-- CREATE plants TABLE
CREATE TABLE plants (
    id uuid PRIMARY KEY,
    garden_id uuid REFERENCES gardens(id),
    species VARCHAR(100) NOT NULL,
    quantity INTEGER,
    planting_date DATE,
    status plant_status DEFAULT 'planned',
    created_at VARCHAR,
    updated_at VARCHAR,
    deleted_at VARCHAR
);

-- CREATE care_logs TABLE
CREATE TABLE care_logs (
    id uuid PRIMARY KEY,
    plant_id uuid REFERENCES plants(id),
    action VARCHAR(50),
    notes TEXT,
    logged_at VARCHAR
);