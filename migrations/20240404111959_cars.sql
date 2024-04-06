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
    removed boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX car_id_idx ON cars(id);
CREATE INDEX car_regNum_idx ON cars(regNum);
CREATE INDEX car_mark_idx ON cars(mark);
CREATE INDEX car_model_idx ON cars(model);
CREATE INDEX car_year_idx ON cars(year);
CREATE INDEX car_name_idx ON cars(name);
CREATE INDEX car_surname_idx ON cars(surname);
CREATE INDEX car_patronymic_idx ON cars(patronymic);
CREATE INDEX car_removed_idx ON cars(removed);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cars;

DROP INDEX IF EXISTS car_id_idx;
DROP INDEX IF EXISTS car_regNum_idx;
DROP INDEX IF EXISTS car_mark_idx;
DROP INDEX IF EXISTS car_model_idx;
DROP INDEX IF EXISTS car_year_idx;
DROP INDEX IF EXISTS car_name_idx;
DROP INDEX IF EXISTS car_surname_idx;
DROP INDEX IF EXISTS car_patronymic_idx;
DROP INDEX IF EXISTS car_removed_idx;

DROP SEQUENCE car_seq;
-- +goose StatementEnd