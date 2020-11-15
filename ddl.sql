CREATE DATABASE julo;

CREATE SEQUENCE wallets_pk_seq
	INCREMENT BY 1
	MINVALUE 0
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE wallets(
	wallet_id bigint NOT NULL DEFAULT nextval('wallets_pk_seq'::regclass),
	wallet_uuid VARCHAR(225) NOT NULL,
	customer_uuid VARCHAR(225) NOT NULL,
	status varchar(20) NOT NULL,
	balance float8 NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NULL,
    enable_at timestamp without time zone NULL,
    disabled_at timestamp without time zone NULL,
	CONSTRAINT pk_wallet_id PRIMARY KEY (wallet_id)
);

CREATE INDEX idx_wallet_uuid ON comments (wallet_uuid);
CREATE INDEX idx_customer_uuid ON comments (customer_uuid);


CREATE SEQUENCE wallet_transactions_pk_seq
	INCREMENT BY 1
	MINVALUE 0
	MAXVALUE 2147483647
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE wallet_transactions(
	wallet_transaction_id bigint NOT NULL DEFAULT nextval('wallet_transactions_pk_seq'::regclass),
	wallet_transaction_uuid VARCHAR(225) NOT NULL,
	wallet_id int NOT NULL,
	status varchar(20) NOT NULL,
    amount float8 NOT NULL,
    reference_id VARCHAR(255) NOT NULL,
	types int NOT NULL,
    created_at timestamp without time zone NOT NULL,
	CONSTRAINT pk_wallet_transaction_id PRIMARY KEY (wallet_transaction_id)
);

CREATE INDEX idx_wallet_transaction_uuid ON wallet_transactions (wallet_transaction_uuid);
CREATE INDEX idx_customer_id ON wallet_transactions (customer_id);
CREATE INDEX idx_reference_id ON wallet_transactions (reference_id);