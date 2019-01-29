-- this is sample database.

-- create database
create database testdb encoding 'UTF-8';
\c testdb;

-- test table
CREATE TABLE test_user(
    id serial NOT NULL UNIQUE,
    name    TEXT       NOT NULL,
    age     INTEGER    DEFAULT 0,
    PRIMARY KEY (id)
) WITHOUT OIDS;

INSERT INTO test_user (name, age) VALUES ('test', 19);
