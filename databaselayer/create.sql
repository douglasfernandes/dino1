
-- Table: public.animals

-- DROP TABLE public.animals;

CREATE TABLE public.animals
(
    id integer NOT NULL DEFAULT nextval('animals_id_seq'::regclass),
    animal_type text COLLATE pg_catalog."default",
    nickname text COLLATE pg_catalog."default",
    zone integer,
    age integer,
    CONSTRAINT animals_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.animals
    OWNER to postgres;


INSERT INTO animals(animal_type, nickname, zone, age) VALUES ('Tyrannosaurus Rex', 'rex', 1, 10),('Velociraptor', 'rapto', 2, 15) 