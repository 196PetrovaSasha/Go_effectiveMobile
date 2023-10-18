CREATE TABLE fio
(
    id SERIAL,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT NOT NULL,
    nationality TEXT,
    age SMALLINT,
    gender TEXT,
    deleted_on  TIMESTAMP(0) WITHOUT TIME ZONE,

    PRIMARY KEY (id)
);