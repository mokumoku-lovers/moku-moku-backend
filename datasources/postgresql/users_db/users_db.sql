CREATE DATABASE user_db;

CONNECT TO user_db;

CREATE SCHEMA user_db;
SET search_path TO user_db;

DROP TABLE IF EXISTS Users CASCADE;
CREATE TABLE Users (
   id              SERIAL NOT NULL PRIMARY KEY,
   email           VARCHAR(255) NOT NULL UNIQUE,
   username        VARCHAR(50) NOT NULL UNIQUE,
   display_name    VARCHAR(255),
   biography       TEXT,
   birthday        DATE,
   password        VARCHAR(255),
   profile_pic     VARCHAR(255),
   points          INTEGER,
   date_created    DATE
);