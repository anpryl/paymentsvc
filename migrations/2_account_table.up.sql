CREATE TABLE accounts (
       id uuid DEFAULT uuid_generate_v4() NOT NULL,
       currency_numeric_code integer NOT NULL,
       balance numeric NOT NULL
);
