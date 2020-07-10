-- Table: public.activities

-- DROP TABLE public.activities;

CREATE TABLE public.activities
(
    id integer NOT NULL DEFAULT nextval('activities_id_seq'::regclass),
    title character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT activities_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.activities
    OWNER to postgres;