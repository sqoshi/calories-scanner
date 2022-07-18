CREATE DATABASE fooddatabase;
\c fooddatabase;
CREATE SCHEMA IF NOT EXISTS foodschema;
CREATE TABLE IF NOT EXISTS foodschema.food
(
    Category     character varying(255) NOT NULL,
    Name         character varying(255) NOT NULL,
    UnitPer100   character varying(2)   NOT NULL,
    KiloCalories int                    NOT NULL,
    KiloJoule    int
);
COPY foodschema.food FROM '/calories.csv' WITH (FORMAT csv);
