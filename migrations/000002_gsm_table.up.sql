BEGIN;

CREATE TABLE IF NOT EXISTS gsm_table (
                                        id SERIAL PRIMARY KEY,
                                        dt_receiving date,
                                        dt_crch date,
                                        income_kg real,
                                        been_changed boolean,
                                        db_data_creation timestamp,
                                        site_id integer REFERENCES sites ON DELETE CASCADE,
                                        operator_id integer REFERENCES operators ON DELETE CASCADE,
                                        provider_id integer REFERENCES providers ON DELETE CASCADE,
                                        contractor_id integer REFERENCES contractors ON DELETE CASCADE,
                                        license_plate_id integer REFERENCES license_plates ON DELETE CASCADE,
                                        status_id integer REFERENCES  statuses ON DELETE CASCADE,           
                                        guid text);    
CREATE INDEX IF NOT EXISTS gsm_table_dt_receiving ON gsm_table (dt_receiving);
CREATE INDEX IF NOT EXISTS gsm_table_guid ON gsm_table (guid);

COMMIT;