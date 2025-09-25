-- create_user.sql

CREATE USER igorrm19 WITH PASSWORD 'chunda123';

CREATE DATABASE database;

GRANT ALL PRIVILEGES ON DATABASE database TO igorrm19;

ALTER ROLE igorrm19 LOGIN;
