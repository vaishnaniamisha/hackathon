-- userid sequence creation
CREATE SEQUENCE public.userid_seq
    INCREMENT 1
    START 1001
    MINVALUE 1001
    MAXVALUE 999999999999999;

ALTER SEQUENCE public.userid_seq
    OWNER TO postgres;

-- User table creation
CREATE TABLE public.users
(
    "ID" integer NOT NULL DEFAULT nextval('userid_seq'::regclass),
    "Name" character varying(50),
    CONSTRAINT users_pkey PRIMARY KEY ("ID")
);

ALTER TABLE public.users
    OWNER to postgres;

-- challenge sequence creation
CREATE SEQUENCE public.challengeid_seq
    INCREMENT 1
    START 1001
    MINVALUE 1001
    MAXVALUE 999999999999999;

ALTER SEQUENCE public.challengeid_seq
    OWNER TO postgres;

-- Challenge table creation 
CREATE TABLE public.challenges
(
    "ID" integer NOT NULL DEFAULT nextval('challengeid_seq'::regclass) ,
    "Title" character varying(256),
    "Description" character varying(1024),
    "Tag" character varying,
    "VoteCount" integer,
    "CreatedBy" integer,
    "CreatedDate" date,
    "IsDeleted" boolean DEFAULT false,
    CONSTRAINT challenge_pkay PRIMARY KEY ("ID"),
    CONSTRAINT challenge_createdby_users_id_fkey FOREIGN KEY ("CreatedBy")
        REFERENCES public.users ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

ALTER TABLE public.challenges
    OWNER to postgres;

-- Tags table creation
CREATE TABLE public."Tags"
(
    "Tag" character varying(50)
);

ALTER TABLE public."Tags"
    OWNER to postgres;

-- collabration sequence creation
CREATE SEQUENCE public.collabrationid_seq
    INCREMENT 1
    START 1001
    MINVALUE 1001
    MAXVALUE 999999999999999999
    CACHE 1;

ALTER SEQUENCE public.collabrationid_seq
    OWNER TO postgres;

-- ChallengeCollabration table creation
CREATE TABLE public."ChallengeCollabration"
(
    "ID" integer NOT NULL DEFAULT nextval('collabrationid_seq'::regclass),
    "UserId" integer,
    "ChallengeId" integer,
    CONSTRAINT challengecollaboration_pkey PRIMARY KEY ("ID"),
    CONSTRAINT collabration_chanllangeid_challenge_id_fk FOREIGN KEY ("ChallengeId")
        REFERENCES public.challenges ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT collabration_userid_users_id_fk FOREIGN KEY ("UserId")
        REFERENCES public.users ("ID") MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

ALTER TABLE public."ChallengeCollabration"
    OWNER to postgres;

-- Insert Users

INSERT INTO public.users("Name") VALUES ('Amisha V');

INSERT INTO public.users("Name") VALUES ('Peter P');

INSERT INTO public.users("Name") VALUES ('Rajesh');