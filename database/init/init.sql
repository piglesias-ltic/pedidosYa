CREATE SCHEMA main;

CREATE TABLE main.version (
    app_version character varying(30) NOT NULL,
    release_date timestamp without time zone
);

INSERT INTO main.version (app_version, release_date)
VALUES ('0.0.1', current_timestamp);
