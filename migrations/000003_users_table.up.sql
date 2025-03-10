BEGIN;

CREATE TABLE IF NOT EXISTS users (
                            id SERIAL PRIMARY KEY,
                            username text,
                            password text,
                            admin boolean
                    );
CREATE INDEX IF NOT EXISTS users_username ON users (username);

COMMIT;