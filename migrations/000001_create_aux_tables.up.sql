BEGIN;

CREATE TABLE IF NOT EXISTS sites (
                            id SERIAL PRIMARY KEY,
                            name text
                    );


CREATE TABLE IF NOT EXISTS operators (
                            id SERIAL PRIMARY KEY,
                            name text
                    );

CREATE TABLE IF NOT EXISTS providers (
                            id SERIAL PRIMARY KEY,
                            name text
                    );

 CREATE TABLE IF NOT EXISTS contractors (
                            id SERIAL PRIMARY KEY,
                            name text
                    );

CREATE TABLE IF NOT EXISTS license_plates (
                            id SERIAL PRIMARY KEY,
                            name text
                    );

CREATE TABLE IF NOT EXISTS statuses (
                            id SERIAL PRIMARY KEY,
                            name text
                    );


COMMIT;