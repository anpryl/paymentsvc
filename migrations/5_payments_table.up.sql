CREATE TABLE payments (
       id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
       from_account uuid NOT NULL,
       to_account uuid NOT NULL,
       currency_numeric_code smallint NOT NULL,
       amount numeric NOT NULL,
       created_at time NOT NULL,

       CONSTRAINT fk_from_account
       FOREIGN KEY (from_account)
       REFERENCES accounts(id)
       ON DELETE RESTRICT
       ON UPDATE CASCADE,

       CONSTRAINT fk_to_account
       FOREIGN KEY (to_account)
       REFERENCES accounts(id)
       ON DELETE RESTRICT
       ON UPDATE CASCADE
);

CREATE INDEX payments_from_account_idx 
    ON payments(from_account);
