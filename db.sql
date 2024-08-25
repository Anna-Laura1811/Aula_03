--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

-- Started on 2024-08-19 21:46:20

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE postgres;
--
-- TOC entry 4855 (class 1262 OID 5)
-- Name: postgres; Type: DATABASE; Schema: -; Owner: -
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Portuguese_Brazil.1252';


\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4856 (class 0 OID 0)
-- Dependencies: 4855
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- TOC entry 5 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


--
-- TOC entry 4857 (class 0 OID 0)
-- Dependencies: 5
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 32778)
-- Name: battle; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.battle (
    id uuid NOT NULL,
    enemyid uuid NOT NULL,
    playerid uuid NOT NULL,
    dicethrown integer NOT NULL,
    playername character varying(255) NOT NULL,
    enemyname character varying(255) NOT NULL,
    result character varying(255)
);


--
-- TOC entry 217 (class 1259 OID 32773)
-- Name: enemy; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.enemy (
    id uuid NOT NULL,
    nickname character varying(255) NOT NULL,
    life integer NOT NULL,
    attack integer NOT NULL,
    defesa integer NOT NULL,
    heal integer -- Atributo de cura adicionado
);


--
-- TOC entry 216 (class 1259 OID 32768)
-- Name: player; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.player (
    id uuid NOT NULL,
    nickname character varying(255) NOT NULL,
    life integer NOT NULL,
    attack integer NOT NULL,
    defesa integer NOT NULL,
    heal integer, -- Atributo de cura adicionado
    defensa integer
);


--
-- TOC entry 4849 (class 0 OID 32778)
-- Dependencies: 218
-- Data for Name: battle; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.battle VALUES ('351919e6-5873-4b71-934c-c2de64dc8a36', '7cb676cb-0b6f-43b7-b647-63a531bff0c3', '43e104ee-3372-485d-9536-a7b50bc6c7aa', 4, 'Luiz', 'Mori', 'Player won');
INSERT INTO public.battle VALUES ('7bcd3fd8-bb4f-45a4-9736-818790a19fab', '92ce4297-550c-4eea-acc8-7b461ee6deb2', '72cd1661-29d0-4789-847b-1d632f06e618', 5, 'God', 'Monster', 'Player won');
INSERT INTO public.battle VALUES ('5b835ac6-74a0-474d-b52b-b3ce21edd51d', '92ce4297-550c-4eea-acc8-7b461ee6deb2', '72cd1661-29d0-4789-847b-1d632f06e618', 2, 'God', 'Monster', 'Enemy won');
INSERT INTO public.battle VALUES ('f8adbedc-64c1-451a-8114-fbd1fb55f1f9', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 3, 'Felipe', 'Rodrigo', 'Enemy won');
INSERT INTO public.battle VALUES ('91837344-b9f7-4088-a782-1f50e73fb8fb', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 2, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('1eac8646-94f5-408e-aed4-bfa74d7fe050', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 3, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('33b8acaa-7393-4e92-8be3-8fb34174bfb9', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 4, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('1cd54779-c15d-4319-9137-c535019ea472', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 1, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('3026b090-5fe0-4fff-8009-d3e0e2050174', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 6, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('6332213e-b161-460a-9219-cf7ccecab9d5', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 2, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('b0109bd4-aeb1-4a93-a693-c3e088196683', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 2, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('0df3eb85-9921-407e-a112-641e40a002bc', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 1, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('1f06ab99-e1ad-4041-9ef3-aae1b8b4be44', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40fd-9b9e-97430a3108fe', 2, 'Felipe', 'Rodrigo', '');
INSERT INTO public.battle VALUES ('1cf300fd-7b0c-431c-858c-136ec46ff3a0', '44b5f134-2def-46ea-ba30-e5e6bb85a631', '1b57b244-45ce-40
