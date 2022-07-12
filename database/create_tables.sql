CREATE DATABASE FoodDatabase;
CREATE TABLE IF NOT EXISTS Food
(
    Category     character varying(255) NOT NULL,
    Name         character varying(255) NOT NULL,
    UnitPer100   character varying(2) NOT NULL,
    KiloCalories int                    NOT NULL,
    KiloJoule    int
);
COPY Food FROM '/calories.csv' WITH (FORMAT csv);
