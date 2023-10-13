-- Table: public.transactions

-- DROP TABLE IF EXISTS public.transactions;

CREATE TABLE IF NOT EXISTS public.transactions
(
    id text COLLATE pg_catalog."default" NOT NULL,
    "cardId" text COLLATE pg_catalog."default",
    "transactionType" text COLLATE pg_catalog."default",
    amount double precision,
    "creationDate" date,
    CONSTRAINT transactions_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.transactions
    OWNER to postgres;