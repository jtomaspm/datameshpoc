-- Table: public.cards

-- DROP TABLE IF EXISTS public.cards;

CREATE TABLE IF NOT EXISTS public.cards
(
    id text COLLATE pg_catalog."default" NOT NULL,
    "clientId" text COLLATE pg_catalog."default",
    "cardNumber" text COLLATE pg_catalog."default",
    "creationDate" date,
    CONSTRAINT clients_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.cards
    OWNER to postgres;