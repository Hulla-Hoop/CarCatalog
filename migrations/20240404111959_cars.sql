-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE  car_seq
START WITH 1
INCREMENT BY 1;

CREATE TABLE cars (
    id int DEFAULT nextval('car_seq'::regclass) NOT NULL,
    regNum text NOT NULL,
    mark text,
    model text,
    year int,
    name text NOT NULL,
    surname text NOT NULL,
    patronymic text,
    PRIMARY KEY (id)
);

CREATE INDEX car_id_idx ON cars(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cars;

DROP INDEX IF EXISTS car_id_idx;

DROP SEQUENCE car_seq;
-- +goose StatementEnd