CREATE DATABASE user_db;

DROP TABLE IF EXISTS Users CASCADE;
CREATE TABLE Users (
    id          SERIAL NOT NULL PRIMARY KEY,
    email       VARCHAR(255) NOT NULL UNIQUE,
    username    VARCHAR(50) NOT NULL UNIQUE,
    first_name  VARCHAR(50),
    last_name   VARCHAR(50),
    biography   TEXT,
    birthday    TIMESTAMP,
    password    VARCHAR(255),
    profile_pic VARCHAR(255),
    points      INTEGER,
    date_created TIMESTAMP
)