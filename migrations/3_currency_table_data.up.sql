CREATE TABLE currencies (
       numeric_code smallint PRIMARY KEY,
       alpha_code text UNIQUE NOT NULL,
       minor smallint NOT NULL
);

ALTER TABLE accounts 
  ADD CONSTRAINT fk_currency_code 
      FOREIGN KEY (currency_numeric_code) 
      REFERENCES currencies(numeric_code)
      ON DELETE RESTRICT
      ON UPDATE CASCADE;

INSERT INTO currencies (alpha_code, numeric_code, minor)
VALUES ('USD', 840, 2),
       ('EUR', 978, 2),
       ('UAH', 980, 2),
       ('RUB', 643, 2),
       ('BYN', 933, 2);
