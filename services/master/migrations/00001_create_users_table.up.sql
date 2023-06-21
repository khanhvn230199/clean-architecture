CREATE TABLE if not exists schema_migration5 (
                                                id serial NOT NULL,
                                                service varchar not null,
                                                version bigint NOT null,
                                                tstamp timestamp NULL default now(),
    PRIMARY KEY(id),
    UNIQUE (service, version)
    );