--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE user_db;




--
-- Drop roles
--

DROP ROLE root;


--
-- Roles
--

CREATE ROLE root;
ALTER ROLE root WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:EczG1LMurp4xuYU4WC+IKw==$uR+jVWFsChSct2IhA+TOzx7ybEtCl9OnSnvw7jy6Iyw=:K3HPSj5QyA3TbcW5YWomdGijo5t5uEd8d4Yp3/qYoG8=';

--
-- User Configurations
--








--
-- Databases
--

--
-- Database "template1" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg120+2)

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

UPDATE pg_catalog.pg_database SET datistemplate = false WHERE datname = 'template1';
DROP DATABASE template1;
--
-- Name: template1; Type: DATABASE; Schema: -; Owner: root
--

CREATE DATABASE template1 WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE template1 OWNER TO root;

\connect template1

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
-- Name: DATABASE template1; Type: COMMENT; Schema: -; Owner: root
--

COMMENT ON DATABASE template1 IS 'default template for new databases';


--
-- Name: template1; Type: DATABASE PROPERTIES; Schema: -; Owner: root
--

ALTER DATABASE template1 IS_TEMPLATE = true;


\connect template1

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
-- Name: DATABASE template1; Type: ACL; Schema: -; Owner: root
--

REVOKE CONNECT,TEMPORARY ON DATABASE template1 FROM PUBLIC;
GRANT CONNECT ON DATABASE template1 TO PUBLIC;


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg120+2)

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
-- Name: postgres; Type: DATABASE; Schema: -; Owner: root
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO root;

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
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: root
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- PostgreSQL database dump complete
--

--
-- Database "user_db" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Debian 16.2-1.pgdg120+2)
-- Dumped by pg_dump version 16.2 (Debian 16.2-1.pgdg120+2)

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
-- Name: user_db; Type: DATABASE; Schema: -; Owner: root
--

CREATE DATABASE user_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE user_db OWNER TO root;

\connect user_db

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
-- Name: friends; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.friends (
    id integer NOT NULL,
    user_id1 integer,
    user_id2 integer,
    status character varying DEFAULT 'Pending'::character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.friends OWNER TO root;

--
-- Name: friends_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.friends_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.friends_id_seq OWNER TO root;

--
-- Name: friends_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.friends_id_seq OWNED BY public.friends.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO root;

--
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying NOT NULL,
    email character varying NOT NULL,
    firstname character varying NOT NULL,
    lastname character varying NOT NULL,
    phone_no character varying NOT NULL,
    private_account boolean DEFAULT false NOT NULL,
    nationality character varying NOT NULL,
    birthday date NOT NULL,
    gender character varying NOT NULL,
    photourl character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.users OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO root;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: friends id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.friends ALTER COLUMN id SET DEFAULT nextval('public.friends_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: friends; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.friends (id, user_id1, user_id2, status, created_at) FROM stdin;
1	10	1	Accepted	2024-04-18 21:05:53.288137
2	2	5	Accepted	2024-04-18 21:05:53.398022
3	12	6	Pending	2024-04-18 21:05:53.509377
4	2	9	Pending	2024-04-18 21:05:53.615512
6	7	8	Pending	2024-04-18 21:05:53.827584
7	1	6	Accepted	2024-04-18 21:05:53.93493
8	12	9	Pending	2024-04-18 21:05:54.043543
9	6	10	Accepted	2024-04-18 21:05:54.152335
10	11	2	Pending	2024-04-18 21:05:54.266343
11	9	10	Accepted	2024-04-18 21:05:54.371157
12	12	4	Accepted	2024-04-18 21:05:54.477638
13	7	1	Accepted	2024-04-18 21:05:54.581889
14	12	11	Accepted	2024-04-18 21:05:54.686998
15	2	7	Accepted	2024-04-18 21:05:54.794892
16	9	8	Accepted	2024-04-18 21:05:54.902217
18	9	2	Pending	2024-04-18 21:05:55.120161
19	1	5	Accepted	2024-04-18 21:05:55.226875
21	3	1	Accepted	2024-04-18 21:10:57.104716
22	3	2	Accepted	2024-04-18 21:10:57.104716
24	3	4	Accepted	2024-04-18 21:10:57.104716
25	3	5	Accepted	2024-04-18 21:10:57.104716
26	3	6	Accepted	2024-04-18 21:10:57.104716
27	3	7	Accepted	2024-04-18 21:10:57.104716
28	3	8	Accepted	2024-04-18 21:10:57.104716
29	3	9	Accepted	2024-04-18 21:10:57.104716
30	3	10	Accepted	2024-04-18 21:10:57.104716
31	3	11	Accepted	2024-04-18 21:10:57.104716
32	3	12	Accepted	2024-04-18 21:10:57.104716
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.users (id, username, email, firstname, lastname, phone_no, private_account, nationality, birthday, gender, photourl, created_at) FROM stdin;
1	Bankie888	Bank@gmail.com	Sethanan	BankBank	+123456123	f	TH	2003-01-02	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fbankie.jpeg?alt=media&token=857d7f0b-2858-4666-9388-177337a81502	2024-04-18 21:05:50.892166
2	Betty552	Bethh@gmail.com	Elizabeth	Bethbeth	+9876541231	f	UK	2001-05-15	female	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fbet%26lyly.jpeg?alt=media&token=cc8e2634-159c-4a38-8912-aa8d128c41ab	2024-04-18 21:05:50.898694
3	dunepw	Dune@gmail.com	Pitiphon	Chaicharoen	+98765123121	f	TH	2005-05-15	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fdunepw.jpeg?alt=media&token=cc34ea0f-7a3d-4c93-ade4-031c5e54beb6	2024-04-18 21:05:50.900414
4	jajajedi	Jedi@gmail.com	Theerothai	Sithlord	+987123121	f	JP	2001-10-15	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fjajajedi.jpeg?alt=media&token=80889185-4327-4a0b-b0e7-7d064c5e911e	2024-04-18 21:05:50.901507
5	kaoskywalker	999999999@gmail.com	Thanthai	Kruthong	+12312321	f	US	2007-02-15	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fkaoskywalkerz.jpeg?alt=media&token=938f051a-71ea-460c-887b-550e22d74842	2024-04-18 21:05:50.902511
6	Kri7x	krit@gmail.com	Krit	KritKirtKirt	+1231231241	f	TH	2002-03-15	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fkri7x.jpeg?alt=media&token=346e539d-76f0-42ce-ae5f-5f68b7980e72	2024-04-18 21:05:50.903507
7	GuYzaza888	Guy@gmail.com	Krittin	Guyguy	+423423421	f	JP	2003-05-15	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fkrittineke.jpeg?alt=media&token=a8ab5e13-49fd-4dd1-af5a-e6ed028870f2	2024-04-18 21:05:50.904867
8	mearzwong999	mearzwong@gmail.com	Wongsapat	Wong	+9821231	f	TH	2001-05-15	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fmearzwong.jpeg?alt=media&token=5d832972-66a5-4765-8f73-8287e7c09309	2024-04-18 21:05:50.905967
9	Minniecyp888	mnmn@gmail.com	Chayapa	Minniemouse	+23421341	f	FN	2005-05-15	female	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fminnie.jpeg?alt=media&token=7a405ab1-a9c4-4007-b90b-d3c4954ab201	2024-04-18 21:05:50.907029
10	oatptchy	oat555@gmail.com	Nadech	Koogimiya	+2342351	f	TH	2004-10-15	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Foatptchy.jpeg?alt=media&token=4caccd1e-e964-44b7-8834-aedbdfe5130b	2024-04-18 21:05:50.907995
11	pungdevil66	kanompang@gmail.com	Tassanai	Peesarj	+987654234234221	f	TH	2001-05-15	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fpungmonster.jpeg?alt=media&token=7e1ca9e8-8b14-421c-92cd-d33faa13363c	2024-04-18 21:05:50.908834
12	rushaoosh	menjoo@gmail.com	Napat	Laokai	+23423123151	f	TH	2003-05-15	male	https://example.com/janesmith.jpg	2024-04-18 21:05:50.909616
13	teenoisukiThailand	tee@gmail.com	Tee	Teenoi	+9002304321	f	TH	2005-05-11	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fteenoisuki.jpeg?alt=media&token=e55abfff-74f9-484c-9a1c-b281bc4d2b89	2024-04-18 21:05:50.910542
14	wints	yaitoe@gmail.com	Win	Joetoo	+9894321	f	TH	2004-05-11	male	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/userprofile%2Fmockuserphoto%2Fwin_ts.jpeg?alt=media&token=035ef9f4-9b8d-4cd9-b05f-cff1f895445b	2024-04-18 21:05:50.911772
\.


--
-- Name: friends_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.friends_id_seq', 32, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.users_id_seq', 14, true);


--
-- Name: friends friends_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT friends_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: friends unique_friendship; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT unique_friendship UNIQUE (user_id1, user_id2);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_phone_no_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_no_key UNIQUE (phone_no);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: friends friends_user_id1_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT friends_user_id1_fkey FOREIGN KEY (user_id1) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: friends friends_user_id2_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.friends
    ADD CONSTRAINT friends_user_id2_fkey FOREIGN KEY (user_id2) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--

