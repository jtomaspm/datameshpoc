-- Table: public.clients

-- DROP TABLE IF EXISTS public.clients;

CREATE TABLE IF NOT EXISTS public.clients
(
    id text COLLATE pg_catalog."default" NOT NULL,
    "personId" text COLLATE pg_catalog."default",
    "creationDate" date,
    CONSTRAINT clients_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.clients
    OWNER to postgres;