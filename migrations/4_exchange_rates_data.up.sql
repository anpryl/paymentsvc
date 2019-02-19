CREATE TABLE exchange_rates (
       id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
       currency_numeric_code_from smallint NOT NULL,
       currency_numeric_code_to smallint NOT NULL,
       rate numeric NOT NULL,

       CONSTRAINT fk_currency_code_from
       FOREIGN KEY (currency_numeric_code_from)
       REFERENCES currencies(numeric_code)
       ON DELETE RESTRICT
       ON UPDATE CASCADE,

       CONSTRAINT fk_currency_code_to
       FOREIGN KEY (currency_numeric_code_to)
       REFERENCES currencies(numeric_code)
       ON DELETE RESTRICT
       ON UPDATE CASCADE,

       CONSTRAINT exchange_rates_nonnegative_rate CHECK (rate >= 0)
);

CREATE INDEX exchange_rates_currency_code_to_from_idx 
    ON exchange_rates(currency_numeric_code_to, currency_numeric_code_from);

CREATE UNIQUE INDEX exchange_currencies 
    ON exchange_rates (currency_numeric_code_from, currency_numeric_code_to);
 
INSERT INTO exchange_rates (currency_numeric_code_from, currency_numeric_code_to, rate)
VALUES (840, 978, 0.8846),
       (840, 980, 27.173),
       (840, 643, 66.2112),
       (840, 933, 2.1567);
