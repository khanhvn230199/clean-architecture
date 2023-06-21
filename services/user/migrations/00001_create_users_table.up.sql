CREATE TABLE if not exists schema_migration4 (
                                                id serial NOT NULL,
                                                service varchar not null,
                                                version_id bigint NOT null UNIQUE,
                                                tstamp timestamp NULL default now(),
    PRIMARY KEY(id)
    );