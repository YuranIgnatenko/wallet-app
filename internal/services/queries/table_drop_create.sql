DROP TABLE Wallets;
CREATE TABLE Wallets (id UUID NOT NULL, balance NUMERIC(10,2) NOT NULL);
INSERT INTO Wallets (id, balance) VALUES (gen_random_uuid (),50.215), (gen_random_uuid (),221.214);
