#!/bin/bash

set -e
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER bexsy;
	CREATE DATABASE daily;
	ALTER USER bexsy WITH ENCRYPTED PASSWORD 'mich';
	GRANT ALL PRIVILEGES ON DATABASE daily TO bexsy;
EOSQL

psql -v ON_ERROR_STOP=1 --username postgres --dbname daily <<-EOSQL
	CREATE TABLE equations (id SERIAL PRIMARY KEY, equation VARCHAR(6));
	CREATE TABLE day_solution (id SERIAL PRIMARY KEY, solution VARCHAR(6), dt DATE);
	INSERT INTO equations (equation) VALUES ('1+1+40'), ('2+2*20'), ('-1*-42');
	INSERT INTO day_solution (solution, dt) VALUES ('1+1+40', NOW());
EOSQL
