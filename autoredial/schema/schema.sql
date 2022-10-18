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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: callhistory; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.callhistory (
    id character varying(255) NOT NULL,
    scammer_id character varying(255) NOT NULL,
    call_at date NOT NULL,
    call_sec bigint NOT NULL,
    result boolean NOT NULL,
    talk_sec bigint
);


--
-- Name: matching; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.matching (
    id character varying(255) NOT NULL,
    created_at date NOT NULL,
    serial_number bigint NOT NULL,
    matched boolean NOT NULL,
    checked boolean NOT NULL,
    matching_at date,
    talk_sec bigint,
    transcript text
);


--
-- Name: matching_serial_number_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.matching_serial_number_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: matching_serial_number_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.matching_serial_number_seq OWNED BY public.matching.serial_number;


--
-- Name: matching_tag; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.matching_tag (
    matching_id character varying(255) NOT NULL,
    tag character varying(255) NOT NULL
);


--
-- Name: metric; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.metric (
    id character varying(255) NOT NULL,
    created_at date NOT NULL,
    calls bigint NOT NULL,
    scammers bigint NOT NULL,
    inactives bigint NOT NULL,
    call_sec bigint NOT NULL,
    talk_sec bigint NOT NULL,
    amount bigint NOT NULL
);


--
-- Name: scammer; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.scammer (
    id character varying(255) NOT NULL,
    tel character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    is_active boolean NOT NULL
);


--
-- Name: scammer_call; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.scammer_call (
    scammer_id character varying(255) NOT NULL,
    call_id character varying(255) NOT NULL
);


--
-- Name: scammer_matching; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.scammer_matching (
    scammer_id character varying(255) NOT NULL,
    matching_id character varying(255) NOT NULL
);


--
-- Name: scammer_tag; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.scammer_tag (
    scammer_id character varying(255) NOT NULL,
    tag character varying(255) NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(255) NOT NULL
);


--
-- Name: matching serial_number; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.matching ALTER COLUMN serial_number SET DEFAULT nextval('public.matching_serial_number_seq'::regclass);


--
-- Name: callhistory callhistory_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.callhistory
    ADD CONSTRAINT callhistory_pkey PRIMARY KEY (id);


--
-- Name: matching matching_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.matching
    ADD CONSTRAINT matching_pkey PRIMARY KEY (id);


--
-- Name: metric metric_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.metric
    ADD CONSTRAINT metric_pkey PRIMARY KEY (id);


--
-- Name: scammer scammer_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scammer
    ADD CONSTRAINT scammer_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: callhistory callhistory_scammer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.callhistory
    ADD CONSTRAINT callhistory_scammer_id_fkey FOREIGN KEY (scammer_id) REFERENCES public.scammer(id);


--
-- Name: matching_tag matching_tag_matching_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.matching_tag
    ADD CONSTRAINT matching_tag_matching_id_fkey FOREIGN KEY (matching_id) REFERENCES public.matching(id);


--
-- Name: scammer_call scammer_call_call_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scammer_call
    ADD CONSTRAINT scammer_call_call_id_fkey FOREIGN KEY (call_id) REFERENCES public.callhistory(id);


--
-- Name: scammer_call scammer_call_scammer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scammer_call
    ADD CONSTRAINT scammer_call_scammer_id_fkey FOREIGN KEY (scammer_id) REFERENCES public.scammer(id);


--
-- Name: scammer_matching scammer_matching_matching_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scammer_matching
    ADD CONSTRAINT scammer_matching_matching_id_fkey FOREIGN KEY (matching_id) REFERENCES public.matching(id);


--
-- Name: scammer_matching scammer_matching_scammer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scammer_matching
    ADD CONSTRAINT scammer_matching_scammer_id_fkey FOREIGN KEY (scammer_id) REFERENCES public.scammer(id);


--
-- Name: scammer_tag scammer_tag_scammer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scammer_tag
    ADD CONSTRAINT scammer_tag_scammer_id_fkey FOREIGN KEY (scammer_id) REFERENCES public.scammer(id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20221012035352');
