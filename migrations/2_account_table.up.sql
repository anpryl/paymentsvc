CREATE TABLE accounts (
       id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
       currency_numeric_code smallint NOT NULL,
       balance numeric NOT NULL,

       CONSTRAINT accounts_nonnegative_balance CHECK (balance >= 0)
);
