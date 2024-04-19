--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Drop databases (except postgres and template1)
--

DROP DATABASE group_db;




--
-- Drop roles
--

DROP ROLE root;


--
-- Roles
--

CREATE ROLE root;
ALTER ROLE root WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'SCRAM-SHA-256$4096:NEprhOvTdW2hRFdXAe/iwg==$v4hYNEesnd/U63aSIYlyEacYwfPRFAPTjcfaTgKSJyw=:o4pRj8No6xFv4Dw1Yzmuz74F9rWs4Kwg19ymMix/xh4=';

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
-- Database "group_db" dump
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
-- Name: group_db; Type: DATABASE; Schema: -; Owner: root
--

CREATE DATABASE group_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE group_db OWNER TO root;

\connect group_db

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
-- Name: group_members; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.group_members (
    group_id integer NOT NULL,
    member_id integer NOT NULL,
    role character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.group_members OWNER TO root;

--
-- Name: group_requests; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.group_requests (
    group_id integer NOT NULL,
    member_id integer NOT NULL,
    description character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.group_requests OWNER TO root;

--
-- Name: groups; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.groups (
    group_id integer NOT NULL,
    group_name character varying NOT NULL,
    group_creator_id integer NOT NULL,
    description character varying,
    photo_url character varying,
    tag_id integer NOT NULL,
    frequency integer NOT NULL,
    max_members integer DEFAULT 25 NOT NULL,
    group_type character varying DEFAULT 'social'::character varying NOT NULL,
    visibility boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.groups OWNER TO root;

--
-- Name: groups_group_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.groups_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.groups_group_id_seq OWNER TO root;

--
-- Name: groups_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.groups_group_id_seq OWNED BY public.groups.group_id;


--
-- Name: post_comments; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.post_comments (
    comment_id integer NOT NULL,
    post_id integer NOT NULL,
    member_id integer NOT NULL,
    comment character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.post_comments OWNER TO root;

--
-- Name: post_comments_comment_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.post_comments_comment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.post_comments_comment_id_seq OWNER TO root;

--
-- Name: post_comments_comment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.post_comments_comment_id_seq OWNED BY public.post_comments.comment_id;


--
-- Name: post_reactions; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.post_reactions (
    reaction_id integer NOT NULL,
    post_id integer NOT NULL,
    member_id integer NOT NULL,
    reaction character varying NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.post_reactions OWNER TO root;

--
-- Name: post_reactions_post_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.post_reactions_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.post_reactions_post_id_seq OWNER TO root;

--
-- Name: post_reactions_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.post_reactions_post_id_seq OWNED BY public.post_reactions.post_id;


--
-- Name: post_reactions_reaction_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.post_reactions_reaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.post_reactions_reaction_id_seq OWNER TO root;

--
-- Name: post_reactions_reaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.post_reactions_reaction_id_seq OWNED BY public.post_reactions.reaction_id;


--
-- Name: posts; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.posts (
    post_id integer NOT NULL,
    member_id integer NOT NULL,
    group_id integer NOT NULL,
    photo_url character varying,
    description character varying,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.posts OWNER TO root;

--
-- Name: posts_group_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.posts_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.posts_group_id_seq OWNER TO root;

--
-- Name: posts_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.posts_group_id_seq OWNED BY public.posts.group_id;


--
-- Name: posts_post_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.posts_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.posts_post_id_seq OWNER TO root;

--
-- Name: posts_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.posts_post_id_seq OWNED BY public.posts.post_id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO root;

--
-- Name: streak_set; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.streak_set (
    streak_set_id integer NOT NULL,
    group_id integer NOT NULL,
    member_id integer NOT NULL,
    end_date timestamp without time zone NOT NULL,
    start_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.streak_set OWNER TO root;

--
-- Name: streak_set_streak_set_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.streak_set_streak_set_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.streak_set_streak_set_id_seq OWNER TO root;

--
-- Name: streak_set_streak_set_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.streak_set_streak_set_id_seq OWNED BY public.streak_set.streak_set_id;


--
-- Name: streaks; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.streaks (
    streak_id integer NOT NULL,
    streak_set_id integer NOT NULL,
    max_streak_count integer DEFAULT 0 NOT NULL,
    total_streak_count integer DEFAULT 0 NOT NULL,
    weekly_streak_count integer DEFAULT 0 NOT NULL,
    completed boolean DEFAULT false NOT NULL,
    recent_date_added timestamp without time zone,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.streaks OWNER TO root;

--
-- Name: streaks_streak_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.streaks_streak_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.streaks_streak_id_seq OWNER TO root;

--
-- Name: streaks_streak_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.streaks_streak_id_seq OWNED BY public.streaks.streak_id;


--
-- Name: tag_category; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.tag_category (
    category_id integer NOT NULL,
    category_name character varying NOT NULL
);


ALTER TABLE public.tag_category OWNER TO root;

--
-- Name: tag_category_category_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.tag_category_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tag_category_category_id_seq OWNER TO root;

--
-- Name: tag_category_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.tag_category_category_id_seq OWNED BY public.tag_category.category_id;


--
-- Name: tags; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.tags (
    tag_id integer NOT NULL,
    tag_name character varying NOT NULL,
    icon_url character varying,
    category_id integer
);


ALTER TABLE public.tags OWNER TO root;

--
-- Name: tags_tag_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.tags_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.tags_tag_id_seq OWNER TO root;

--
-- Name: tags_tag_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.tags_tag_id_seq OWNED BY public.tags.tag_id;


--
-- Name: groups group_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.groups ALTER COLUMN group_id SET DEFAULT nextval('public.groups_group_id_seq'::regclass);


--
-- Name: post_comments comment_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post_comments ALTER COLUMN comment_id SET DEFAULT nextval('public.post_comments_comment_id_seq'::regclass);


--
-- Name: post_reactions reaction_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post_reactions ALTER COLUMN reaction_id SET DEFAULT nextval('public.post_reactions_reaction_id_seq'::regclass);


--
-- Name: post_reactions post_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post_reactions ALTER COLUMN post_id SET DEFAULT nextval('public.post_reactions_post_id_seq'::regclass);


--
-- Name: posts post_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts ALTER COLUMN post_id SET DEFAULT nextval('public.posts_post_id_seq'::regclass);


--
-- Name: posts group_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts ALTER COLUMN group_id SET DEFAULT nextval('public.posts_group_id_seq'::regclass);


--
-- Name: streak_set streak_set_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.streak_set ALTER COLUMN streak_set_id SET DEFAULT nextval('public.streak_set_streak_set_id_seq'::regclass);


--
-- Name: streaks streak_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.streaks ALTER COLUMN streak_id SET DEFAULT nextval('public.streaks_streak_id_seq'::regclass);


--
-- Name: tag_category category_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tag_category ALTER COLUMN category_id SET DEFAULT nextval('public.tag_category_category_id_seq'::regclass);


--
-- Name: tags tag_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tags ALTER COLUMN tag_id SET DEFAULT nextval('public.tags_tag_id_seq'::regclass);


--
-- Data for Name: group_members; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.group_members (group_id, member_id, role, created_at) FROM stdin;
1	1	creator	2024-04-18 22:57:06.777899
2	2	creator	2024-04-18 22:57:06.780908
3	5	creator	2024-04-18 22:57:06.782542
4	10	creator	2024-04-18 22:57:06.784682
5	5	creator	2024-04-18 22:57:06.786699
6	8	creator	2024-04-18 22:57:06.788411
7	9	creator	2024-04-18 22:57:06.790247
8	4	creator	2024-04-18 22:57:06.791926
9	6	creator	2024-04-18 22:57:06.794324
10	11	creator	2024-04-18 22:57:06.796214
1	2	member	2024-04-18 23:01:56.653952
1	14	member	2024-04-18 23:01:56.66657
1	10	member	2024-04-18 23:02:39.58649
1	4	member	2024-04-18 23:02:39.595255
1	8	member	2024-04-18 23:02:39.602456
1	11	member	2024-04-18 23:02:39.610149
1	3	member	2024-04-18 23:02:39.615531
1	13	member	2024-04-18 23:02:39.620863
1	6	member	2024-04-18 23:02:39.628867
1	7	member	2024-04-18 23:02:39.633341
2	3	member	2024-04-18 23:02:39.639196
2	1	member	2024-04-18 23:02:39.644503
2	9	member	2024-04-18 23:02:39.649016
2	6	member	2024-04-18 23:02:49.361099
2	4	member	2024-04-18 23:02:49.371445
2	8	member	2024-04-18 23:02:49.378693
2	10	member	2024-04-18 23:02:49.383991
2	14	member	2024-04-18 23:02:49.390619
2	12	member	2024-04-18 23:02:49.395751
2	11	member	2024-04-18 23:02:49.400713
2	7	member	2024-04-18 23:02:49.406073
3	1	member	2024-04-18 23:02:49.412576
3	2	member	2024-04-18 23:02:49.41695
3	14	member	2024-04-18 23:02:49.421584
3	3	member	2024-04-18 23:02:49.426395
3	12	member	2024-04-18 23:02:49.430722
3	6	member	2024-04-18 23:02:49.434482
3	7	member	2024-04-18 23:02:49.438716
3	13	member	2024-04-18 23:02:49.442631
3	9	member	2024-04-18 23:02:49.446684
4	2	member	2024-04-18 23:03:08.364279
4	3	member	2024-04-18 23:03:08.373716
4	1	member	2024-04-18 23:03:30.337174
4	12	member	2024-04-18 23:03:30.345449
4	5	member	2024-04-18 23:03:30.351004
4	7	member	2024-04-18 23:03:30.35829
4	13	member	2024-04-18 23:03:30.363692
4	4	member	2024-04-18 23:03:30.369485
4	8	member	2024-04-18 23:03:30.374713
5	3	member	2024-04-18 23:03:30.380942
5	14	member	2024-04-18 23:03:30.385391
5	1	member	2024-04-18 23:03:30.390116
5	6	member	2024-04-18 23:03:30.394921
5	9	member	2024-04-18 23:03:30.399153
5	2	member	2024-04-18 23:03:30.403225
5	12	member	2024-04-18 23:03:42.979421
5	10	member	2024-04-18 23:03:42.988283
5	11	member	2024-04-18 23:03:42.996704
5	13	member	2024-04-18 23:03:43.000262
6	3	member	2024-04-18 23:03:43.006831
6	1	member	2024-04-18 23:03:55.305347
6	10	member	2024-04-18 23:03:55.314086
6	4	member	2024-04-18 23:03:55.319702
6	14	member	2024-04-18 23:03:55.325741
6	6	member	2024-04-18 23:03:55.331523
6	7	member	2024-04-18 23:03:55.335853
6	2	member	2024-04-18 23:03:55.341158
6	11	member	2024-04-18 23:03:55.345979
7	14	member	2024-04-18 23:03:55.351081
7	8	member	2024-04-18 23:03:55.355682
7	4	member	2024-04-18 23:03:55.360676
7	12	member	2024-04-18 23:03:55.364182
7	5	member	2024-04-18 23:03:55.368047
7	2	member	2024-04-18 23:03:55.371979
7	1	member	2024-04-18 23:03:55.376012
7	13	member	2024-04-18 23:03:55.379884
7	3	member	2024-04-18 23:03:55.383251
7	10	member	2024-04-18 23:03:55.386988
8	1	member	2024-04-18 23:03:55.391123
8	9	member	2024-04-18 23:03:55.394511
8	6	member	2024-04-18 23:03:55.398072
8	2	member	2024-04-18 23:03:55.40107
8	3	member	2024-04-18 23:03:55.403774
8	7	member	2024-04-18 23:03:55.406359
8	12	member	2024-04-18 23:03:55.40902
8	5	member	2024-04-18 23:03:55.411868
9	2	member	2024-04-18 23:03:55.415364
9	4	member	2024-04-18 23:03:55.418104
9	11	member	2024-04-18 23:03:55.420776
9	7	member	2024-04-18 23:03:55.423417
9	14	member	2024-04-18 23:04:05.303368
9	1	member	2024-04-18 23:04:05.312561
9	8	member	2024-04-18 23:04:05.317138
9	10	member	2024-04-18 23:04:05.322549
9	3	member	2024-04-18 23:04:05.327378
10	3	member	2024-04-18 23:04:05.332815
10	5	member	2024-04-18 23:04:05.337309
10	2	member	2024-04-18 23:04:05.34145
10	10	member	2024-04-18 23:04:05.346014
10	12	member	2024-04-18 23:04:05.349957
10	13	member	2024-04-18 23:04:05.354369
10	6	member	2024-04-18 23:04:05.358021
10	1	member	2024-04-18 23:04:05.361645
10	4	member	2024-04-18 23:04:14.458952
10	8	member	2024-04-18 23:04:14.462774
\.


--
-- Data for Name: group_requests; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.group_requests (group_id, member_id, description, created_at) FROM stdin;
\.


--
-- Data for Name: groups; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.groups (group_id, group_name, group_creator_id, description, photo_url, tag_id, frequency, max_members, group_type, visibility, created_at) FROM stdin;
1	ISE football club	1	Weekly football at BBB football club, ma join gunn	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Ffootball.jpeg?alt=media&token=2ab15bda-2a4f-47e2-88c7-00c7a8597290	19	0	50	social	t	2024-04-18 22:57:06.773416
3	Coursera ganag	5	Enhance your soft skills with us!	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fonlinecourse.jpeg?alt=media&token=52d021cc-d775-4985-beea-226c5a7afd4f	43	0	30	social	t	2024-04-18 22:57:06.781807
4	Code nerdyy	10	Commit the code 5 times a week and you will receive nerdy trophy from us	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fprogramming.png?alt=media&token=210e448e-6764-4a7b-8934-71380dc6c26f	46	0	30	social	t	2024-04-18 22:57:06.783464
5	Midterm try hard gang	5	Study hard, no fail, no cry, no F, happy life	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fmidtermexam.jpeg?alt=media&token=4e7dbe16-8ff8-4db4-aed0-50349b46b969	51	0	30	social	t	2024-04-18 22:57:06.785748
6	ISE bodybuilder	8	Get up from your bed and start working out <3	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fgym.png?alt=media&token=1c78b800-1c3e-407e-b038-f82873c614c1	53	0	30	social	t	2024-04-18 22:57:06.787654
7	Serious movie discussion	9	Watch movie twice a week and share them here!	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fmovies.png?alt=media&token=02e85840-0d42-4065-a456-01bed8bd1015	2	0	30	social	t	2024-04-18 22:57:06.789215
8	K-drama stands	4	Put down all the work and enjoy K-drama once a week!	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fseries.jpeg?alt=media&token=6758213e-394f-4df3-8d2c-f222080ffd8b	5	0	30	social	t	2024-04-18 22:57:06.790998
9	Intania Music Club	6	Post your music here or you will be cursed by spotify devil	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fmusic.jpeg?alt=media&token=4a034376-5361-4cab-a09b-c59d3027083c	7	0	100	social	t	2024-04-18 22:57:06.793107
10	Travelling & Hanging out	11	Share your beautiful journey here!	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Ftravelling.png?alt=media&token=03b1961f-a6a9-4143-ae20-5038d79ca813	17	0	100	social	t	2024-04-18 22:57:06.795661
2	Intania Badminton	2	Let's join our badminton squad from Engineering Faculty! We encourage 2 times a week!	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/groupphoto%2FMockGroupPhoto%2Fbadminton.png?alt=media&token=2a532ca2-e441-4dd1-a1f5-b9692e2915e6	25	0	30	social	t	2024-04-18 22:57:06.779553
\.


--
-- Data for Name: post_comments; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.post_comments (comment_id, post_id, member_id, comment, created_at) FROM stdin;
1	1	1	Goodjob	2024-04-18 23:52:41.451628
2	2	3	Impressive work!	2024-04-18 23:52:41.462028
3	2	10	Fantastic!	2024-04-18 23:52:41.463742
4	10	4	Outstanding!	2024-04-18 23:52:41.465186
5	4	5	Brilliant!	2024-04-18 23:52:41.466726
6	2	2	Excellent!	2024-04-18 23:52:41.467851
7	5	4	Amazing!	2024-04-18 23:52:41.468772
8	7	6	Awesome!	2024-04-18 23:52:41.469896
9	3	2	Goodjob	2024-04-18 23:52:41.470721
10	4	1	Impressive work!	2024-04-18 23:52:41.471639
11	2	2	Fantastic!	2024-04-18 23:52:41.47242
12	2	3	Outstanding!	2024-04-18 23:52:41.473374
13	2	7	Brilliant!	2024-04-18 23:52:41.474517
14	1	3	Excellent!	2024-04-18 23:52:41.475634
15	2	2	Amazing!	2024-04-18 23:52:41.476693
16	5	14	Awesome!	2024-04-18 23:52:41.47773
17	1	1	Goodjob	2024-04-18 23:52:41.478771
18	2	2	Impressive work!	2024-04-18 23:52:41.479803
19	2	10	Fantastic!	2024-04-18 23:52:41.480984
20	10	4	Outstanding!	2024-04-18 23:52:41.482088
21	4	1	Brilliant!	2024-04-18 23:52:41.483144
22	2	2	Excellent!	2024-04-18 23:52:41.484133
23	5	5	Amazing!	2024-04-18 23:52:41.484976
24	7	3	Awesome!	2024-04-18 23:52:41.485866
25	3	2	Goodjob	2024-04-18 23:52:41.486691
26	4	1	Impressive work!	2024-04-18 23:52:41.48767
27	2	4	Fantastic!	2024-04-18 23:52:41.488509
28	2	3	Outstanding!	2024-04-18 23:52:41.489031
29	2	7	Brilliant!	2024-04-18 23:52:41.489541
30	1	9	Excellent!	2024-04-18 23:52:41.490193
31	2	2	Amazing!	2024-04-18 23:52:41.49086
32	5	8	Awesome!	2024-04-18 23:52:41.491663
33	1	1	Goodjob	2024-04-18 23:52:41.492433
34	2	2	Impressive work!	2024-04-18 23:52:41.493207
35	2	10	Fantastic!	2024-04-18 23:52:41.493863
36	10	4	Outstanding!	2024-04-18 23:52:41.494624
37	4	1	Brilliant!	2024-04-18 23:52:41.495369
38	3	2	Excellent!	2024-04-18 23:52:41.496219
39	5	5	Amazing!	2024-04-18 23:52:41.497031
40	7	3	Awesome!	2024-04-18 23:52:41.497905
41	2	2	Goodjob	2024-04-18 23:52:41.498703
42	4	1	Impressive work!	2024-04-18 23:52:41.499519
43	9	4	Fantastic!	2024-04-18 23:52:41.500362
44	2	3	Outstanding!	2024-04-18 23:52:41.501013
45	7	7	Brilliant!	2024-04-18 23:52:41.50162
46	1	9	Excellent!	2024-04-18 23:52:41.502191
47	2	2	Amazing!	2024-04-18 23:52:41.502785
48	5	8	Awesome!	2024-04-18 23:52:41.503329
\.


--
-- Data for Name: post_reactions; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.post_reactions (reaction_id, post_id, member_id, reaction, created_at) FROM stdin;
1	1	1	love	2024-04-19 00:03:15.971067
2	1	2	haha	2024-04-19 00:03:15.977077
3	1	14	wow	2024-04-19 00:03:15.978534
4	1	10	wow	2024-04-19 00:03:15.979697
5	1	4	love	2024-04-19 00:03:15.980741
6	1	8	like	2024-04-19 00:03:15.982216
7	1	11	angry	2024-04-19 00:03:15.983731
8	1	3	like	2024-04-19 00:03:15.985184
9	1	13	sad	2024-04-19 00:03:15.98628
10	1	6	wow	2024-04-19 00:03:15.987356
11	1	7	like	2024-04-19 00:03:15.988466
12	2	1	wow	2024-04-19 00:03:15.989494
13	2	2	wow	2024-04-19 00:03:15.990571
14	2	14	wow	2024-04-19 00:03:15.991469
15	2	10	wow	2024-04-19 00:03:15.991926
16	2	4	haha	2024-04-19 00:03:15.992785
17	2	8	haha	2024-04-19 00:03:15.993723
18	2	11	haha	2024-04-19 00:03:15.994645
19	2	3	like	2024-04-19 00:03:15.995465
20	2	13	sad	2024-04-19 00:03:15.996315
21	2	6	haha	2024-04-19 00:03:15.997113
22	2	7	like	2024-04-19 00:03:15.997887
23	3	1	angry	2024-04-19 00:03:15.998688
24	3	2	love	2024-04-19 00:03:15.999509
25	3	14	love	2024-04-19 00:03:16.000377
26	3	10	love	2024-04-19 00:03:16.001182
27	3	4	angry	2024-04-19 00:03:16.002058
28	3	8	love	2024-04-19 00:03:16.002778
29	3	11	haha	2024-04-19 00:03:16.00351
30	3	3	angry	2024-04-19 00:03:16.004222
31	3	13	love	2024-04-19 00:03:16.004957
32	3	6	wow	2024-04-19 00:03:16.005808
33	3	7	like	2024-04-19 00:03:16.006612
34	4	1	angry	2024-04-19 00:03:16.007398
35	4	2	like	2024-04-19 00:03:16.008133
36	4	14	wow	2024-04-19 00:03:16.00869
37	4	10	love	2024-04-19 00:03:16.009313
38	4	4	like	2024-04-19 00:03:16.010034
39	4	8	love	2024-04-19 00:03:16.010779
40	4	11	haha	2024-04-19 00:03:16.011452
41	4	3	wow	2024-04-19 00:03:16.012214
42	4	13	angry	2024-04-19 00:03:16.01281
43	4	6	like	2024-04-19 00:03:16.013458
44	4	7	angry	2024-04-19 00:03:16.01409
45	5	1	haha	2024-04-19 00:03:16.014692
46	5	2	sad	2024-04-19 00:03:16.015339
47	5	14	love	2024-04-19 00:03:16.015902
48	5	10	wow	2024-04-19 00:03:16.016488
49	5	4	sad	2024-04-19 00:03:16.017128
50	5	8	angry	2024-04-19 00:03:16.017758
51	5	11	angry	2024-04-19 00:03:16.018342
52	5	3	love	2024-04-19 00:03:16.018994
53	5	13	angry	2024-04-19 00:03:16.019616
54	5	6	wow	2024-04-19 00:03:16.020325
55	5	7	love	2024-04-19 00:03:16.020885
56	6	1	love	2024-04-19 00:03:16.021433
57	6	2	wow	2024-04-19 00:03:16.022492
58	6	14	sad	2024-04-19 00:03:16.023041
59	6	10	wow	2024-04-19 00:03:16.023598
60	6	4	love	2024-04-19 00:03:16.024193
61	6	8	sad	2024-04-19 00:03:16.024743
62	6	11	sad	2024-04-19 00:03:16.025284
63	6	3	wow	2024-04-19 00:03:16.025861
64	6	13	like	2024-04-19 00:03:16.026332
65	6	6	wow	2024-04-19 00:03:16.026905
66	6	7	haha	2024-04-19 00:03:16.027448
67	7	1	like	2024-04-19 00:03:16.02804
68	7	2	like	2024-04-19 00:03:16.028606
69	7	14	haha	2024-04-19 00:03:16.029107
70	7	10	love	2024-04-19 00:03:16.029643
71	7	4	sad	2024-04-19 00:03:16.030184
72	7	8	haha	2024-04-19 00:03:16.030729
73	7	11	wow	2024-04-19 00:03:16.031236
74	7	3	like	2024-04-19 00:03:16.031805
75	7	13	wow	2024-04-19 00:03:16.032378
76	7	6	like	2024-04-19 00:03:16.032826
77	7	7	angry	2024-04-19 00:03:16.03325
78	8	1	wow	2024-04-19 00:03:16.033763
79	8	2	angry	2024-04-19 00:03:16.034305
80	8	14	haha	2024-04-19 00:03:16.034819
81	8	10	like	2024-04-19 00:03:16.035336
82	8	4	like	2024-04-19 00:03:16.03585
83	8	8	sad	2024-04-19 00:03:16.036345
84	8	11	like	2024-04-19 00:03:16.036821
85	8	3	love	2024-04-19 00:03:16.037318
86	8	13	angry	2024-04-19 00:03:16.037812
87	8	6	love	2024-04-19 00:03:16.038333
88	8	7	angry	2024-04-19 00:03:16.038759
89	9	1	angry	2024-04-19 00:03:16.039143
90	9	2	wow	2024-04-19 00:03:16.03967
91	9	14	haha	2024-04-19 00:03:16.04016
92	9	10	love	2024-04-19 00:03:16.040635
93	9	4	like	2024-04-19 00:03:16.041091
94	9	8	like	2024-04-19 00:03:16.041476
95	9	11	sad	2024-04-19 00:03:16.041896
96	9	3	haha	2024-04-19 00:03:16.042334
97	9	13	like	2024-04-19 00:03:16.042756
98	9	6	like	2024-04-19 00:03:16.0432
99	9	7	wow	2024-04-19 00:03:16.043673
100	10	1	sad	2024-04-19 00:03:16.0441
101	10	2	sad	2024-04-19 00:03:16.04452
102	10	14	love	2024-04-19 00:03:16.044957
103	10	10	wow	2024-04-19 00:03:16.045422
104	10	4	angry	2024-04-19 00:03:16.045949
105	10	8	haha	2024-04-19 00:03:16.046421
106	10	11	angry	2024-04-19 00:03:16.046892
107	10	3	love	2024-04-19 00:03:16.047271
108	10	13	wow	2024-04-19 00:03:16.047632
109	10	6	love	2024-04-19 00:03:16.047992
110	10	7	like	2024-04-19 00:03:16.048345
111	1	2	like	2024-04-19 00:03:16.049039
112	1	3	wow	2024-04-19 00:03:16.049408
113	1	1	angry	2024-04-19 00:03:16.049766
114	1	9	like	2024-04-19 00:03:16.05022
115	1	6	like	2024-04-19 00:03:16.050653
116	1	4	sad	2024-04-19 00:03:16.051065
117	1	8	haha	2024-04-19 00:03:16.051529
118	1	10	haha	2024-04-19 00:03:16.051968
119	1	14	like	2024-04-19 00:03:16.052417
120	1	12	like	2024-04-19 00:03:16.052764
121	1	11	love	2024-04-19 00:03:16.053114
122	1	7	love	2024-04-19 00:03:16.053479
123	2	2	angry	2024-04-19 00:03:16.054075
124	2	3	wow	2024-04-19 00:03:16.054499
125	2	1	haha	2024-04-19 00:03:16.05493
126	2	9	angry	2024-04-19 00:03:16.055455
127	2	6	love	2024-04-19 00:03:16.055959
128	2	4	haha	2024-04-19 00:03:16.056432
129	2	8	angry	2024-04-19 00:03:16.056878
130	2	10	haha	2024-04-19 00:03:16.057334
131	2	14	sad	2024-04-19 00:03:16.057759
132	2	12	sad	2024-04-19 00:03:16.058184
133	2	11	sad	2024-04-19 00:03:16.058538
134	2	7	like	2024-04-19 00:03:16.058904
135	3	2	like	2024-04-19 00:03:16.059303
136	3	3	like	2024-04-19 00:03:16.059719
137	3	1	wow	2024-04-19 00:03:16.060119
138	3	9	love	2024-04-19 00:03:16.060592
139	3	6	like	2024-04-19 00:03:16.061085
140	3	4	wow	2024-04-19 00:03:16.061513
141	3	8	wow	2024-04-19 00:03:16.061899
142	3	10	sad	2024-04-19 00:03:16.062316
143	3	14	like	2024-04-19 00:03:16.062781
144	3	12	love	2024-04-19 00:03:16.063662
145	3	11	angry	2024-04-19 00:03:16.064087
146	3	7	angry	2024-04-19 00:03:16.064544
147	4	2	sad	2024-04-19 00:03:16.064929
148	4	3	like	2024-04-19 00:03:16.065333
149	4	1	like	2024-04-19 00:03:16.06585
150	4	9	love	2024-04-19 00:03:16.066308
151	4	6	angry	2024-04-19 00:03:16.066782
152	4	4	haha	2024-04-19 00:03:16.067122
153	4	8	like	2024-04-19 00:03:16.067463
154	4	10	haha	2024-04-19 00:03:16.067948
155	4	14	angry	2024-04-19 00:03:16.068396
156	4	12	angry	2024-04-19 00:03:16.068868
157	4	11	haha	2024-04-19 00:03:16.069305
158	4	7	love	2024-04-19 00:03:16.069781
159	5	2	love	2024-04-19 00:03:16.070244
160	5	3	wow	2024-04-19 00:03:16.070704
161	5	1	like	2024-04-19 00:03:16.071157
162	5	9	like	2024-04-19 00:03:16.071637
163	5	6	haha	2024-04-19 00:03:16.072083
164	5	4	love	2024-04-19 00:03:16.07257
165	5	8	sad	2024-04-19 00:03:16.073009
166	5	10	angry	2024-04-19 00:03:16.073499
167	5	14	angry	2024-04-19 00:03:16.073933
168	5	12	haha	2024-04-19 00:03:16.074454
169	5	11	angry	2024-04-19 00:03:16.074828
170	5	7	haha	2024-04-19 00:03:16.075189
171	6	2	wow	2024-04-19 00:03:16.075733
172	6	3	love	2024-04-19 00:03:16.076186
173	6	1	angry	2024-04-19 00:03:16.076628
174	6	9	like	2024-04-19 00:03:16.077061
175	6	6	haha	2024-04-19 00:03:16.077512
176	6	4	angry	2024-04-19 00:03:16.078039
177	6	8	love	2024-04-19 00:03:16.078578
178	6	10	sad	2024-04-19 00:03:16.079089
179	6	14	sad	2024-04-19 00:03:16.079519
180	6	12	angry	2024-04-19 00:03:16.079949
181	6	11	love	2024-04-19 00:03:16.080403
182	6	7	love	2024-04-19 00:03:16.080854
183	7	2	like	2024-04-19 00:03:16.081288
184	7	3	haha	2024-04-19 00:03:16.081727
185	7	1	angry	2024-04-19 00:03:16.082108
186	7	9	wow	2024-04-19 00:03:16.082526
187	7	6	love	2024-04-19 00:03:16.082994
188	7	4	love	2024-04-19 00:03:16.08334
189	7	8	love	2024-04-19 00:03:16.083915
190	7	10	wow	2024-04-19 00:03:16.084319
191	7	14	love	2024-04-19 00:03:16.08476
192	7	12	haha	2024-04-19 00:03:16.08516
193	7	11	haha	2024-04-19 00:03:16.085554
194	7	7	like	2024-04-19 00:03:16.085954
195	8	2	sad	2024-04-19 00:03:16.086348
196	8	3	sad	2024-04-19 00:03:16.086739
197	8	1	angry	2024-04-19 00:03:16.087118
198	8	9	haha	2024-04-19 00:03:16.087535
199	8	6	haha	2024-04-19 00:03:16.087981
200	8	4	like	2024-04-19 00:03:16.088393
201	8	8	sad	2024-04-19 00:03:16.088785
202	8	10	sad	2024-04-19 00:03:16.089169
203	8	14	haha	2024-04-19 00:03:16.089562
204	8	12	haha	2024-04-19 00:03:16.089976
205	8	11	love	2024-04-19 00:03:16.090361
206	8	7	haha	2024-04-19 00:03:16.090783
207	9	2	haha	2024-04-19 00:03:16.091162
208	9	3	haha	2024-04-19 00:03:16.091588
209	9	1	sad	2024-04-19 00:03:16.091976
210	9	9	wow	2024-04-19 00:03:16.092364
211	9	6	wow	2024-04-19 00:03:16.092735
212	9	4	sad	2024-04-19 00:03:16.093133
213	9	8	sad	2024-04-19 00:03:16.093534
214	9	10	haha	2024-04-19 00:03:16.093902
215	9	14	like	2024-04-19 00:03:16.094257
216	9	12	love	2024-04-19 00:03:16.094638
217	9	11	like	2024-04-19 00:03:16.095144
218	9	7	love	2024-04-19 00:03:16.095589
219	10	2	love	2024-04-19 00:03:16.095982
220	10	3	love	2024-04-19 00:03:16.096373
221	10	1	like	2024-04-19 00:03:16.096766
222	10	9	haha	2024-04-19 00:03:16.097122
223	10	6	like	2024-04-19 00:03:16.097474
224	10	4	wow	2024-04-19 00:03:16.09783
225	10	8	sad	2024-04-19 00:03:16.09819
226	10	10	sad	2024-04-19 00:03:16.098567
227	10	14	like	2024-04-19 00:03:16.098951
228	10	12	sad	2024-04-19 00:03:16.099326
229	10	11	angry	2024-04-19 00:03:16.099703
230	10	7	angry	2024-04-19 00:03:16.100058
231	1	5	sad	2024-04-19 00:03:16.100714
232	1	1	haha	2024-04-19 00:03:16.10107
233	1	2	wow	2024-04-19 00:03:16.101461
234	1	14	love	2024-04-19 00:03:16.101897
235	1	3	angry	2024-04-19 00:03:16.10226
236	1	12	like	2024-04-19 00:03:16.102603
237	1	6	wow	2024-04-19 00:03:16.102959
238	1	7	haha	2024-04-19 00:03:16.103333
239	1	13	love	2024-04-19 00:03:16.10368
240	1	9	love	2024-04-19 00:03:16.104031
241	2	5	like	2024-04-19 00:03:16.104397
242	2	1	wow	2024-04-19 00:03:16.104725
243	2	2	angry	2024-04-19 00:03:16.105008
244	2	14	wow	2024-04-19 00:03:16.105293
245	2	3	wow	2024-04-19 00:03:16.105616
246	2	12	haha	2024-04-19 00:03:16.105905
247	2	6	sad	2024-04-19 00:03:16.106201
248	2	7	like	2024-04-19 00:03:16.106491
249	2	13	like	2024-04-19 00:03:16.106779
250	2	9	sad	2024-04-19 00:03:16.1071
251	3	5	sad	2024-04-19 00:03:16.107479
252	3	1	angry	2024-04-19 00:03:16.107764
253	3	2	like	2024-04-19 00:03:16.108068
254	3	14	like	2024-04-19 00:03:16.108546
255	3	3	angry	2024-04-19 00:03:16.108848
256	3	12	sad	2024-04-19 00:03:16.109141
257	3	6	sad	2024-04-19 00:03:16.109441
258	3	7	haha	2024-04-19 00:03:16.10977
259	3	13	love	2024-04-19 00:03:16.110063
260	3	9	love	2024-04-19 00:03:16.110366
261	4	5	angry	2024-04-19 00:03:16.110667
262	4	1	love	2024-04-19 00:03:16.110976
263	4	2	angry	2024-04-19 00:03:16.111267
264	4	14	angry	2024-04-19 00:03:16.111583
265	4	3	haha	2024-04-19 00:03:16.111903
266	4	12	like	2024-04-19 00:03:16.112204
267	4	6	haha	2024-04-19 00:03:16.112512
268	4	7	like	2024-04-19 00:03:16.112838
269	4	13	love	2024-04-19 00:03:16.113132
270	4	9	sad	2024-04-19 00:03:16.113455
271	5	5	wow	2024-04-19 00:03:16.11373
272	5	1	like	2024-04-19 00:03:16.114014
273	5	2	sad	2024-04-19 00:03:16.114306
274	5	14	like	2024-04-19 00:03:16.114601
275	5	3	love	2024-04-19 00:03:16.114902
276	5	12	haha	2024-04-19 00:03:16.11519
277	5	6	love	2024-04-19 00:03:16.115486
278	5	7	love	2024-04-19 00:03:16.115806
279	5	13	like	2024-04-19 00:03:16.116115
280	5	9	wow	2024-04-19 00:03:16.116425
281	6	5	like	2024-04-19 00:03:16.116701
282	6	1	haha	2024-04-19 00:03:16.116994
283	6	2	love	2024-04-19 00:03:16.117287
284	6	14	angry	2024-04-19 00:03:16.117668
285	6	3	angry	2024-04-19 00:03:16.117952
286	6	12	wow	2024-04-19 00:03:16.118254
287	6	6	sad	2024-04-19 00:03:16.118595
288	6	7	wow	2024-04-19 00:03:16.118874
289	6	13	sad	2024-04-19 00:03:16.119192
290	6	9	haha	2024-04-19 00:03:16.119474
291	7	5	like	2024-04-19 00:03:16.119762
292	7	1	like	2024-04-19 00:03:16.120048
293	7	2	sad	2024-04-19 00:03:16.120333
294	7	14	angry	2024-04-19 00:03:16.120619
295	7	3	sad	2024-04-19 00:03:16.120894
296	7	12	angry	2024-04-19 00:03:16.121172
297	7	6	wow	2024-04-19 00:03:16.12147
298	7	7	like	2024-04-19 00:03:16.121749
299	7	13	love	2024-04-19 00:03:16.122044
300	7	9	wow	2024-04-19 00:03:16.122346
301	8	5	sad	2024-04-19 00:03:16.122653
302	8	1	sad	2024-04-19 00:03:16.122934
303	8	2	love	2024-04-19 00:03:16.123236
304	8	14	angry	2024-04-19 00:03:16.123527
305	8	3	haha	2024-04-19 00:03:16.123806
306	8	12	love	2024-04-19 00:03:16.124096
307	8	6	sad	2024-04-19 00:03:16.124409
308	8	7	haha	2024-04-19 00:03:16.124716
309	8	13	haha	2024-04-19 00:03:16.125002
310	8	9	sad	2024-04-19 00:03:16.125289
311	9	5	haha	2024-04-19 00:03:16.125584
312	9	1	like	2024-04-19 00:03:16.125872
313	9	2	wow	2024-04-19 00:03:16.126154
314	9	14	angry	2024-04-19 00:03:16.126466
315	9	3	sad	2024-04-19 00:03:16.126777
316	9	12	like	2024-04-19 00:03:16.127077
317	9	6	sad	2024-04-19 00:03:16.127371
318	9	7	wow	2024-04-19 00:03:16.127679
319	9	13	angry	2024-04-19 00:03:16.127963
320	9	9	like	2024-04-19 00:03:16.128509
321	10	5	angry	2024-04-19 00:03:16.128818
322	10	1	angry	2024-04-19 00:03:16.129128
323	10	2	sad	2024-04-19 00:03:16.129431
324	10	14	sad	2024-04-19 00:03:16.129707
325	10	3	like	2024-04-19 00:03:16.130007
326	10	12	haha	2024-04-19 00:03:16.130308
327	10	6	haha	2024-04-19 00:03:16.13061
328	10	7	wow	2024-04-19 00:03:16.130937
329	10	13	wow	2024-04-19 00:03:16.131222
330	10	9	love	2024-04-19 00:03:16.131507
331	1	10	wow	2024-04-19 00:03:16.132054
332	1	2	love	2024-04-19 00:03:16.132372
333	1	3	love	2024-04-19 00:03:16.132662
334	1	1	haha	2024-04-19 00:03:16.132955
335	1	12	like	2024-04-19 00:03:16.133237
336	1	5	love	2024-04-19 00:03:16.133544
337	1	7	like	2024-04-19 00:03:16.133855
338	1	13	wow	2024-04-19 00:03:16.13418
339	1	4	haha	2024-04-19 00:03:16.134481
340	1	8	angry	2024-04-19 00:03:16.134772
341	2	10	haha	2024-04-19 00:03:16.135052
342	2	2	like	2024-04-19 00:03:16.135338
343	2	3	wow	2024-04-19 00:03:16.135633
344	2	1	wow	2024-04-19 00:03:16.135941
345	2	12	wow	2024-04-19 00:03:16.13626
346	2	5	angry	2024-04-19 00:03:16.136567
347	2	7	angry	2024-04-19 00:03:16.136891
348	2	13	sad	2024-04-19 00:03:16.137188
349	2	4	sad	2024-04-19 00:03:16.137461
350	2	8	haha	2024-04-19 00:03:16.137736
351	3	10	haha	2024-04-19 00:03:16.138022
352	3	2	haha	2024-04-19 00:03:16.138313
353	3	3	like	2024-04-19 00:03:16.138648
354	3	1	wow	2024-04-19 00:03:16.13892
355	3	12	sad	2024-04-19 00:03:16.139193
356	3	5	love	2024-04-19 00:03:16.139491
357	3	7	love	2024-04-19 00:03:16.139787
358	3	13	sad	2024-04-19 00:03:16.140059
359	3	4	haha	2024-04-19 00:03:16.140333
360	3	8	like	2024-04-19 00:03:16.140627
361	4	10	angry	2024-04-19 00:03:16.140931
362	4	2	love	2024-04-19 00:03:16.141237
363	4	3	haha	2024-04-19 00:03:16.141551
364	4	1	love	2024-04-19 00:03:16.141857
365	4	12	like	2024-04-19 00:03:16.142136
366	4	5	haha	2024-04-19 00:03:16.142431
367	4	7	sad	2024-04-19 00:03:16.142726
368	4	13	love	2024-04-19 00:03:16.143025
369	4	4	angry	2024-04-19 00:03:16.143322
370	4	8	love	2024-04-19 00:03:16.143609
371	5	10	angry	2024-04-19 00:03:16.143877
372	5	2	sad	2024-04-19 00:03:16.144156
373	5	3	love	2024-04-19 00:03:16.144436
374	5	1	like	2024-04-19 00:03:16.144722
375	5	12	sad	2024-04-19 00:03:16.145002
376	5	5	sad	2024-04-19 00:03:16.145283
377	5	7	haha	2024-04-19 00:03:16.145562
378	5	13	love	2024-04-19 00:03:16.145869
379	5	4	sad	2024-04-19 00:03:16.146167
380	5	8	like	2024-04-19 00:03:16.14646
381	6	10	angry	2024-04-19 00:03:16.146749
382	6	2	like	2024-04-19 00:03:16.147032
383	6	3	love	2024-04-19 00:03:16.147304
384	6	1	angry	2024-04-19 00:03:16.147591
385	6	12	love	2024-04-19 00:03:16.147898
386	6	5	love	2024-04-19 00:03:16.148362
387	6	7	sad	2024-04-19 00:03:16.148653
388	6	13	love	2024-04-19 00:03:16.148947
389	6	4	wow	2024-04-19 00:03:16.149247
390	6	8	love	2024-04-19 00:03:16.149529
391	7	10	sad	2024-04-19 00:03:16.149811
392	7	2	haha	2024-04-19 00:03:16.150093
393	7	3	like	2024-04-19 00:03:16.150384
394	7	1	love	2024-04-19 00:03:16.150663
395	7	12	wow	2024-04-19 00:03:16.150952
396	7	5	love	2024-04-19 00:03:16.151227
397	7	7	like	2024-04-19 00:03:16.151513
398	7	13	haha	2024-04-19 00:03:16.151789
399	7	4	love	2024-04-19 00:03:16.152077
400	7	8	angry	2024-04-19 00:03:16.152366
401	8	10	like	2024-04-19 00:03:16.152659
402	8	2	haha	2024-04-19 00:03:16.152951
403	8	3	love	2024-04-19 00:03:16.153255
404	8	1	haha	2024-04-19 00:03:16.153536
405	8	12	wow	2024-04-19 00:03:16.153844
406	8	5	like	2024-04-19 00:03:16.154145
407	8	7	sad	2024-04-19 00:03:16.154444
408	8	13	sad	2024-04-19 00:03:16.154739
409	8	4	like	2024-04-19 00:03:16.155224
410	8	8	haha	2024-04-19 00:03:16.155509
411	9	10	like	2024-04-19 00:03:16.155815
412	9	2	sad	2024-04-19 00:03:16.156102
413	9	3	love	2024-04-19 00:03:16.156433
414	9	1	wow	2024-04-19 00:03:16.156759
415	9	12	haha	2024-04-19 00:03:16.157105
416	9	5	love	2024-04-19 00:03:16.157402
417	9	7	love	2024-04-19 00:03:16.157699
418	9	13	sad	2024-04-19 00:03:16.158746
419	9	4	love	2024-04-19 00:03:16.159088
420	9	8	haha	2024-04-19 00:03:16.159389
421	10	10	wow	2024-04-19 00:03:16.159708
422	10	2	angry	2024-04-19 00:03:16.160023
423	10	3	angry	2024-04-19 00:03:16.16033
424	10	1	sad	2024-04-19 00:03:16.160647
425	10	12	haha	2024-04-19 00:03:16.160969
426	10	5	love	2024-04-19 00:03:16.161363
427	10	7	wow	2024-04-19 00:03:16.161667
428	10	13	angry	2024-04-19 00:03:16.161971
429	10	4	haha	2024-04-19 00:03:16.162265
430	10	8	angry	2024-04-19 00:03:16.162569
431	1	5	love	2024-04-19 00:03:16.163131
432	1	3	wow	2024-04-19 00:03:16.163424
433	1	14	angry	2024-04-19 00:03:16.163718
434	1	1	haha	2024-04-19 00:03:16.164016
435	1	6	love	2024-04-19 00:03:16.164309
436	1	9	wow	2024-04-19 00:03:16.164584
437	1	2	love	2024-04-19 00:03:16.16489
438	1	12	love	2024-04-19 00:03:16.165187
439	1	10	wow	2024-04-19 00:03:16.165471
440	1	11	haha	2024-04-19 00:03:16.16581
441	1	13	angry	2024-04-19 00:03:16.166136
442	2	5	angry	2024-04-19 00:03:16.166407
443	2	3	haha	2024-04-19 00:03:16.166705
444	2	14	love	2024-04-19 00:03:16.166991
445	2	1	angry	2024-04-19 00:03:16.167284
446	2	6	haha	2024-04-19 00:03:16.167572
447	2	9	angry	2024-04-19 00:03:16.167862
448	2	2	angry	2024-04-19 00:03:16.168502
449	2	12	love	2024-04-19 00:03:16.1688
450	2	10	haha	2024-04-19 00:03:16.169097
451	2	11	sad	2024-04-19 00:03:16.169371
452	2	13	wow	2024-04-19 00:03:16.16966
453	3	5	wow	2024-04-19 00:03:16.16994
454	3	3	haha	2024-04-19 00:03:16.17021
455	3	14	angry	2024-04-19 00:03:16.170495
456	3	1	like	2024-04-19 00:03:16.170781
457	3	6	wow	2024-04-19 00:03:16.171065
458	3	9	haha	2024-04-19 00:03:16.17134
459	3	2	like	2024-04-19 00:03:16.171668
460	3	12	angry	2024-04-19 00:03:16.17196
461	3	10	angry	2024-04-19 00:03:16.172256
462	3	11	sad	2024-04-19 00:03:16.172545
463	3	13	angry	2024-04-19 00:03:16.172834
464	4	5	angry	2024-04-19 00:03:16.173119
465	4	3	angry	2024-04-19 00:03:16.1734
466	4	14	angry	2024-04-19 00:03:16.173662
467	4	1	sad	2024-04-19 00:03:16.17393
468	4	6	angry	2024-04-19 00:03:16.174231
469	4	9	sad	2024-04-19 00:03:16.174546
470	4	2	love	2024-04-19 00:03:16.174873
471	4	12	angry	2024-04-19 00:03:16.175163
472	4	10	love	2024-04-19 00:03:16.175471
473	4	11	sad	2024-04-19 00:03:16.175777
474	4	13	like	2024-04-19 00:03:16.17609
475	5	5	haha	2024-04-19 00:03:16.176379
476	5	3	wow	2024-04-19 00:03:16.176694
477	5	14	haha	2024-04-19 00:03:16.176975
478	5	1	haha	2024-04-19 00:03:16.177276
479	5	6	haha	2024-04-19 00:03:16.177575
480	5	9	wow	2024-04-19 00:03:16.177858
481	5	2	sad	2024-04-19 00:03:16.178164
482	5	12	wow	2024-04-19 00:03:16.178464
483	5	10	sad	2024-04-19 00:03:16.178752
484	5	11	haha	2024-04-19 00:03:16.179037
485	5	13	haha	2024-04-19 00:03:16.179334
486	6	5	love	2024-04-19 00:03:16.179628
487	6	3	angry	2024-04-19 00:03:16.179962
488	6	14	haha	2024-04-19 00:03:16.18028
489	6	1	sad	2024-04-19 00:03:16.180705
490	6	6	sad	2024-04-19 00:03:16.183351
491	6	9	love	2024-04-19 00:03:16.183672
492	6	2	haha	2024-04-19 00:03:16.183972
493	6	12	sad	2024-04-19 00:03:16.184281
494	6	10	angry	2024-04-19 00:03:16.184648
495	6	11	haha	2024-04-19 00:03:16.18582
496	6	13	like	2024-04-19 00:03:16.186139
497	7	5	wow	2024-04-19 00:03:16.186473
498	7	3	wow	2024-04-19 00:03:16.186822
499	7	14	wow	2024-04-19 00:03:16.189334
500	7	1	like	2024-04-19 00:03:16.189653
501	7	6	like	2024-04-19 00:03:16.189949
502	7	9	angry	2024-04-19 00:03:16.190255
503	7	2	sad	2024-04-19 00:03:16.190593
504	7	12	love	2024-04-19 00:03:16.190965
505	7	10	like	2024-04-19 00:03:16.193339
506	7	11	wow	2024-04-19 00:03:16.193637
507	7	13	like	2024-04-19 00:03:16.193952
508	8	5	like	2024-04-19 00:03:16.194238
509	8	3	wow	2024-04-19 00:03:16.194538
510	8	14	haha	2024-04-19 00:03:16.194855
511	8	1	sad	2024-04-19 00:03:16.195161
512	8	6	love	2024-04-19 00:03:16.195515
513	8	9	haha	2024-04-19 00:03:16.196344
514	8	2	wow	2024-04-19 00:03:16.197893
515	8	12	angry	2024-04-19 00:03:16.200443
516	8	10	angry	2024-04-19 00:03:16.203371
517	8	11	haha	2024-04-19 00:03:16.203696
518	8	13	haha	2024-04-19 00:03:16.203981
519	9	5	like	2024-04-19 00:03:16.204253
520	9	3	sad	2024-04-19 00:03:16.204555
521	9	14	like	2024-04-19 00:03:16.204837
522	9	1	haha	2024-04-19 00:03:16.205142
523	9	6	angry	2024-04-19 00:03:16.205415
524	9	9	haha	2024-04-19 00:03:16.205723
525	9	2	like	2024-04-19 00:03:16.206005
526	9	12	sad	2024-04-19 00:03:16.206309
527	9	10	like	2024-04-19 00:03:16.206599
528	9	11	haha	2024-04-19 00:03:16.206886
529	9	13	haha	2024-04-19 00:03:16.207174
530	10	5	like	2024-04-19 00:03:16.207483
531	10	3	angry	2024-04-19 00:03:16.207814
532	10	14	sad	2024-04-19 00:03:16.208102
533	10	1	wow	2024-04-19 00:03:16.208385
534	10	6	love	2024-04-19 00:03:16.20868
535	10	9	like	2024-04-19 00:03:16.208992
536	10	2	haha	2024-04-19 00:03:16.20931
537	10	12	haha	2024-04-19 00:03:16.209617
538	10	10	haha	2024-04-19 00:03:16.209911
539	10	11	angry	2024-04-19 00:03:16.210199
540	10	13	like	2024-04-19 00:03:16.210494
541	1	8	haha	2024-04-19 00:03:16.211028
542	1	3	angry	2024-04-19 00:03:16.211322
543	1	1	love	2024-04-19 00:03:16.211633
544	1	10	wow	2024-04-19 00:03:16.211954
545	1	4	love	2024-04-19 00:03:16.212267
546	1	14	wow	2024-04-19 00:03:16.212602
547	1	6	angry	2024-04-19 00:03:16.212903
548	1	7	haha	2024-04-19 00:03:16.213185
549	1	2	sad	2024-04-19 00:03:16.213477
550	1	11	sad	2024-04-19 00:03:16.213777
551	2	8	love	2024-04-19 00:03:16.214091
552	2	3	like	2024-04-19 00:03:16.214393
553	2	1	wow	2024-04-19 00:03:16.214739
554	2	10	haha	2024-04-19 00:03:16.215055
555	2	4	angry	2024-04-19 00:03:16.215356
556	2	14	love	2024-04-19 00:03:16.21568
557	2	6	haha	2024-04-19 00:03:16.215996
558	2	7	wow	2024-04-19 00:03:16.216318
559	2	2	love	2024-04-19 00:03:16.216609
560	2	11	like	2024-04-19 00:03:16.216961
561	3	8	haha	2024-04-19 00:03:16.217271
562	3	3	haha	2024-04-19 00:03:16.21758
563	3	1	angry	2024-04-19 00:03:16.218099
564	3	10	wow	2024-04-19 00:03:16.218444
565	3	4	love	2024-04-19 00:03:16.218766
566	3	14	haha	2024-04-19 00:03:16.21915
567	3	6	angry	2024-04-19 00:03:16.219574
568	3	7	like	2024-04-19 00:03:16.219954
569	3	2	angry	2024-04-19 00:03:16.220314
570	3	11	love	2024-04-19 00:03:16.220644
571	4	8	love	2024-04-19 00:03:16.22093
572	4	3	haha	2024-04-19 00:03:16.2213
573	4	1	angry	2024-04-19 00:03:16.221667
574	4	10	angry	2024-04-19 00:03:16.221931
575	4	4	haha	2024-04-19 00:03:16.222271
576	4	14	wow	2024-04-19 00:03:16.222654
577	4	6	sad	2024-04-19 00:03:16.222904
578	4	7	love	2024-04-19 00:03:16.223165
579	4	2	angry	2024-04-19 00:03:16.223793
580	4	11	haha	2024-04-19 00:03:16.224124
581	5	8	sad	2024-04-19 00:03:16.224628
582	5	3	love	2024-04-19 00:03:16.224974
583	5	1	wow	2024-04-19 00:03:16.225257
584	5	10	haha	2024-04-19 00:03:16.225554
585	5	4	love	2024-04-19 00:03:16.225813
586	5	14	sad	2024-04-19 00:03:16.226051
587	5	6	angry	2024-04-19 00:03:16.226285
588	5	7	like	2024-04-19 00:03:16.226525
589	5	2	wow	2024-04-19 00:03:16.226812
590	5	11	sad	2024-04-19 00:03:16.227069
591	6	8	sad	2024-04-19 00:03:16.227318
592	6	3	wow	2024-04-19 00:03:16.227992
593	6	1	angry	2024-04-19 00:03:16.228316
594	6	10	angry	2024-04-19 00:03:16.228699
595	6	4	wow	2024-04-19 00:03:16.228953
596	6	14	haha	2024-04-19 00:03:16.229223
597	6	6	haha	2024-04-19 00:03:16.229457
598	6	7	wow	2024-04-19 00:03:16.229784
599	6	2	angry	2024-04-19 00:03:16.230175
600	6	11	like	2024-04-19 00:03:16.230469
601	7	8	love	2024-04-19 00:03:16.231169
602	7	3	angry	2024-04-19 00:03:16.231552
603	7	1	wow	2024-04-19 00:03:16.231958
604	7	10	haha	2024-04-19 00:03:16.232261
605	7	4	like	2024-04-19 00:03:16.232588
606	7	14	wow	2024-04-19 00:03:16.232917
607	7	6	angry	2024-04-19 00:03:16.233197
608	7	7	haha	2024-04-19 00:03:16.233474
609	7	2	wow	2024-04-19 00:03:16.233727
610	7	11	sad	2024-04-19 00:03:16.233967
611	8	8	love	2024-04-19 00:03:16.234226
612	8	3	wow	2024-04-19 00:03:16.234532
613	8	1	love	2024-04-19 00:03:16.234854
614	8	10	love	2024-04-19 00:03:16.235109
615	8	4	sad	2024-04-19 00:03:16.235347
616	8	14	love	2024-04-19 00:03:16.235588
617	8	6	love	2024-04-19 00:03:16.235841
618	8	7	haha	2024-04-19 00:03:16.2361
619	8	2	wow	2024-04-19 00:03:16.236355
620	8	11	angry	2024-04-19 00:03:16.236589
621	9	8	haha	2024-04-19 00:03:16.236825
622	9	3	sad	2024-04-19 00:03:16.237061
623	9	1	like	2024-04-19 00:03:16.237299
624	9	10	like	2024-04-19 00:03:16.237532
625	9	4	sad	2024-04-19 00:03:16.237802
626	9	14	wow	2024-04-19 00:03:16.238035
627	9	6	angry	2024-04-19 00:03:16.238295
628	9	7	like	2024-04-19 00:03:16.238548
629	9	2	like	2024-04-19 00:03:16.238806
630	9	11	angry	2024-04-19 00:03:16.239036
631	10	8	angry	2024-04-19 00:03:16.239264
632	10	3	like	2024-04-19 00:03:16.239517
633	10	1	angry	2024-04-19 00:03:16.239785
634	10	10	angry	2024-04-19 00:03:16.240046
635	10	4	love	2024-04-19 00:03:16.240285
636	10	14	haha	2024-04-19 00:03:16.24054
637	10	6	love	2024-04-19 00:03:16.240807
638	10	7	wow	2024-04-19 00:03:16.24108
639	10	2	wow	2024-04-19 00:03:16.241384
640	10	11	love	2024-04-19 00:03:16.241647
641	1	9	angry	2024-04-19 00:03:16.24213
642	1	14	like	2024-04-19 00:03:16.242389
643	1	8	love	2024-04-19 00:03:16.242658
644	1	4	angry	2024-04-19 00:03:16.242914
645	1	12	love	2024-04-19 00:03:16.243304
646	1	5	haha	2024-04-19 00:03:16.243548
647	1	2	wow	2024-04-19 00:03:16.243829
648	1	1	like	2024-04-19 00:03:16.244092
649	1	13	like	2024-04-19 00:03:16.244367
650	1	3	angry	2024-04-19 00:03:16.244614
651	1	10	wow	2024-04-19 00:03:16.24486
652	2	9	like	2024-04-19 00:03:16.245124
653	2	14	wow	2024-04-19 00:03:16.245374
654	2	8	angry	2024-04-19 00:03:16.245648
655	2	4	sad	2024-04-19 00:03:16.245911
656	2	12	wow	2024-04-19 00:03:16.246165
657	2	5	like	2024-04-19 00:03:16.246449
658	2	2	like	2024-04-19 00:03:16.246719
659	2	1	love	2024-04-19 00:03:16.24697
660	2	13	angry	2024-04-19 00:03:16.247218
661	2	3	love	2024-04-19 00:03:16.247494
662	2	10	sad	2024-04-19 00:03:16.24776
663	3	9	love	2024-04-19 00:03:16.248015
664	3	14	like	2024-04-19 00:03:16.248281
665	3	8	sad	2024-04-19 00:03:16.248544
666	3	4	like	2024-04-19 00:03:16.248807
667	3	12	haha	2024-04-19 00:03:16.249074
668	3	5	haha	2024-04-19 00:03:16.249331
669	3	2	like	2024-04-19 00:03:16.249583
670	3	1	sad	2024-04-19 00:03:16.249817
671	3	13	wow	2024-04-19 00:03:16.250063
672	3	3	haha	2024-04-19 00:03:16.250323
673	3	10	haha	2024-04-19 00:03:16.250576
674	4	9	haha	2024-04-19 00:03:16.250821
675	4	14	wow	2024-04-19 00:03:16.251075
676	4	8	like	2024-04-19 00:03:16.251323
677	4	4	haha	2024-04-19 00:03:16.251579
678	4	12	like	2024-04-19 00:03:16.251864
679	4	5	like	2024-04-19 00:03:16.252126
680	4	2	love	2024-04-19 00:03:16.252451
681	4	1	love	2024-04-19 00:03:16.252757
682	4	13	sad	2024-04-19 00:03:16.253061
683	4	3	wow	2024-04-19 00:03:16.253342
684	4	10	angry	2024-04-19 00:03:16.253672
685	5	9	love	2024-04-19 00:03:16.254
686	5	14	like	2024-04-19 00:03:16.254299
687	5	8	like	2024-04-19 00:03:16.254585
688	5	4	love	2024-04-19 00:03:16.25488
689	5	12	haha	2024-04-19 00:03:16.255176
690	5	5	angry	2024-04-19 00:03:16.255485
691	5	2	love	2024-04-19 00:03:16.255798
692	5	1	love	2024-04-19 00:03:16.256108
693	5	13	sad	2024-04-19 00:03:16.256413
694	5	3	sad	2024-04-19 00:03:16.256724
695	5	10	love	2024-04-19 00:03:16.257018
696	6	9	sad	2024-04-19 00:03:16.257322
697	6	14	angry	2024-04-19 00:03:16.257649
698	6	8	haha	2024-04-19 00:03:16.257979
699	6	4	sad	2024-04-19 00:03:16.258321
700	6	12	wow	2024-04-19 00:03:16.258637
701	6	5	haha	2024-04-19 00:03:16.258953
702	6	2	like	2024-04-19 00:03:16.259254
703	6	1	like	2024-04-19 00:03:16.259569
704	6	13	angry	2024-04-19 00:03:16.259891
705	6	3	angry	2024-04-19 00:03:16.260195
706	6	10	love	2024-04-19 00:03:16.260475
707	7	9	sad	2024-04-19 00:03:16.260775
708	7	14	angry	2024-04-19 00:03:16.261181
709	7	8	wow	2024-04-19 00:03:16.261491
710	7	4	angry	2024-04-19 00:03:16.261775
711	7	12	sad	2024-04-19 00:03:16.262258
712	7	5	like	2024-04-19 00:03:16.262582
713	7	2	wow	2024-04-19 00:03:16.262879
714	7	1	love	2024-04-19 00:03:16.263169
715	7	13	wow	2024-04-19 00:03:16.263464
716	7	3	wow	2024-04-19 00:03:16.263778
717	7	10	sad	2024-04-19 00:03:16.264101
718	8	9	wow	2024-04-19 00:03:16.264392
719	8	14	sad	2024-04-19 00:03:16.264677
720	8	8	angry	2024-04-19 00:03:16.264969
721	8	4	haha	2024-04-19 00:03:16.26525
722	8	12	like	2024-04-19 00:03:16.265541
723	8	5	sad	2024-04-19 00:03:16.265885
724	8	2	love	2024-04-19 00:03:16.266208
725	8	1	like	2024-04-19 00:03:16.2665
726	8	13	sad	2024-04-19 00:03:16.266854
727	8	3	like	2024-04-19 00:03:16.267216
728	8	10	sad	2024-04-19 00:03:16.267494
729	9	9	haha	2024-04-19 00:03:16.26781
730	9	14	like	2024-04-19 00:03:16.26812
731	9	8	haha	2024-04-19 00:03:16.268437
732	9	4	love	2024-04-19 00:03:16.268725
733	9	12	sad	2024-04-19 00:03:16.269027
734	9	5	wow	2024-04-19 00:03:16.269321
735	9	2	sad	2024-04-19 00:03:16.269643
736	9	1	love	2024-04-19 00:03:16.269957
737	9	13	angry	2024-04-19 00:03:16.270254
738	9	3	haha	2024-04-19 00:03:16.270571
739	9	10	love	2024-04-19 00:03:16.270861
740	10	9	angry	2024-04-19 00:03:16.27116
741	10	14	angry	2024-04-19 00:03:16.271461
742	10	8	angry	2024-04-19 00:03:16.271776
743	10	4	wow	2024-04-19 00:03:16.272074
744	10	12	like	2024-04-19 00:03:16.272415
745	10	5	wow	2024-04-19 00:03:16.272715
746	10	2	love	2024-04-19 00:03:16.273012
747	10	1	sad	2024-04-19 00:03:16.273311
748	10	13	sad	2024-04-19 00:03:16.273615
749	10	3	sad	2024-04-19 00:03:16.273942
750	10	10	haha	2024-04-19 00:03:16.274226
751	1	4	wow	2024-04-19 00:03:16.274779
752	1	1	wow	2024-04-19 00:03:16.275077
753	1	9	haha	2024-04-19 00:03:16.275368
754	1	6	haha	2024-04-19 00:03:16.27565
755	1	2	like	2024-04-19 00:03:16.275937
756	1	3	haha	2024-04-19 00:03:16.27623
757	1	7	haha	2024-04-19 00:03:16.276523
758	1	12	sad	2024-04-19 00:03:16.276818
759	1	5	wow	2024-04-19 00:03:16.277124
760	2	4	wow	2024-04-19 00:03:16.277427
761	2	1	sad	2024-04-19 00:03:16.27775
762	2	9	haha	2024-04-19 00:03:16.278048
763	2	6	sad	2024-04-19 00:03:16.278326
764	2	2	sad	2024-04-19 00:03:16.278634
765	2	3	love	2024-04-19 00:03:16.278932
766	2	7	sad	2024-04-19 00:03:16.279223
767	2	12	haha	2024-04-19 00:03:16.279514
768	2	5	haha	2024-04-19 00:03:16.27979
769	3	4	like	2024-04-19 00:03:16.280102
770	3	1	angry	2024-04-19 00:03:16.280404
771	3	9	haha	2024-04-19 00:03:16.280693
772	3	6	haha	2024-04-19 00:03:16.280973
773	3	2	wow	2024-04-19 00:03:16.28126
774	3	3	love	2024-04-19 00:03:16.281546
775	3	7	sad	2024-04-19 00:03:16.282117
776	3	12	like	2024-04-19 00:03:16.282416
777	3	5	angry	2024-04-19 00:03:16.282708
778	4	4	haha	2024-04-19 00:03:16.282999
779	4	1	wow	2024-04-19 00:03:16.283284
780	4	9	love	2024-04-19 00:03:16.28358
781	4	6	haha	2024-04-19 00:03:16.283868
782	4	2	like	2024-04-19 00:03:16.284154
783	4	3	love	2024-04-19 00:03:16.284457
784	4	7	sad	2024-04-19 00:03:16.284757
785	4	12	love	2024-04-19 00:03:16.285065
786	4	5	sad	2024-04-19 00:03:16.285374
787	5	4	wow	2024-04-19 00:03:16.285684
788	5	1	wow	2024-04-19 00:03:16.285994
789	5	9	haha	2024-04-19 00:03:16.286292
790	5	6	wow	2024-04-19 00:03:16.2866
791	5	2	like	2024-04-19 00:03:16.286897
792	5	3	love	2024-04-19 00:03:16.287202
793	5	7	angry	2024-04-19 00:03:16.287523
794	5	12	like	2024-04-19 00:03:16.287815
795	5	5	haha	2024-04-19 00:03:16.28811
796	6	4	wow	2024-04-19 00:03:16.288405
797	6	1	like	2024-04-19 00:03:16.288709
798	6	9	angry	2024-04-19 00:03:16.288986
799	6	6	wow	2024-04-19 00:03:16.289271
800	6	2	angry	2024-04-19 00:03:16.289544
801	6	3	like	2024-04-19 00:03:16.289837
802	6	7	sad	2024-04-19 00:03:16.290114
803	6	12	angry	2024-04-19 00:03:16.29042
804	6	5	sad	2024-04-19 00:03:16.290712
805	7	4	haha	2024-04-19 00:03:16.291002
806	7	1	angry	2024-04-19 00:03:16.291353
807	7	9	like	2024-04-19 00:03:16.29168
808	7	6	love	2024-04-19 00:03:16.291983
809	7	2	love	2024-04-19 00:03:16.292276
810	7	3	angry	2024-04-19 00:03:16.292557
811	7	7	like	2024-04-19 00:03:16.292934
812	7	12	wow	2024-04-19 00:03:16.293223
813	7	5	sad	2024-04-19 00:03:16.293523
814	8	4	love	2024-04-19 00:03:16.293814
815	8	1	wow	2024-04-19 00:03:16.294112
816	8	9	like	2024-04-19 00:03:16.294391
817	8	6	wow	2024-04-19 00:03:16.294704
818	8	2	love	2024-04-19 00:03:16.295024
819	8	3	sad	2024-04-19 00:03:16.295305
820	8	7	haha	2024-04-19 00:03:16.295599
821	8	12	like	2024-04-19 00:03:16.2959
822	8	5	love	2024-04-19 00:03:16.296189
823	9	4	wow	2024-04-19 00:03:16.296484
824	9	1	haha	2024-04-19 00:03:16.296762
825	9	9	sad	2024-04-19 00:03:16.297063
826	9	6	haha	2024-04-19 00:03:16.297333
827	9	2	like	2024-04-19 00:03:16.297627
828	9	3	wow	2024-04-19 00:03:16.297911
829	9	7	like	2024-04-19 00:03:16.298203
830	9	12	like	2024-04-19 00:03:16.298482
831	9	5	angry	2024-04-19 00:03:16.298745
832	10	4	sad	2024-04-19 00:03:16.29904
833	10	1	haha	2024-04-19 00:03:16.299363
834	10	9	haha	2024-04-19 00:03:16.299655
835	10	6	love	2024-04-19 00:03:16.299939
836	10	2	sad	2024-04-19 00:03:16.300206
837	10	3	sad	2024-04-19 00:03:16.300484
838	10	7	like	2024-04-19 00:03:16.300758
839	10	12	wow	2024-04-19 00:03:16.301305
840	10	5	haha	2024-04-19 00:03:16.301575
841	1	6	sad	2024-04-19 00:03:16.302081
842	1	2	haha	2024-04-19 00:03:16.302385
843	1	4	angry	2024-04-19 00:03:16.302685
844	1	11	love	2024-04-19 00:03:16.302994
845	1	7	haha	2024-04-19 00:03:16.303306
846	1	14	haha	2024-04-19 00:03:16.303608
847	1	1	like	2024-04-19 00:03:16.303888
848	1	8	wow	2024-04-19 00:03:16.304166
849	1	10	sad	2024-04-19 00:03:16.304455
850	1	3	wow	2024-04-19 00:03:16.304738
851	2	6	love	2024-04-19 00:03:16.305086
852	2	2	haha	2024-04-19 00:03:16.305364
853	2	4	love	2024-04-19 00:03:16.305668
854	2	11	sad	2024-04-19 00:03:16.30596
855	2	7	sad	2024-04-19 00:03:16.306227
856	2	14	wow	2024-04-19 00:03:16.306519
857	2	1	angry	2024-04-19 00:03:16.306815
858	2	8	angry	2024-04-19 00:03:16.307093
859	2	10	haha	2024-04-19 00:03:16.307375
860	2	3	haha	2024-04-19 00:03:16.307683
861	3	6	like	2024-04-19 00:03:16.307972
862	3	2	like	2024-04-19 00:03:16.308256
863	3	4	like	2024-04-19 00:03:16.308536
864	3	11	haha	2024-04-19 00:03:16.30883
865	3	7	sad	2024-04-19 00:03:16.309121
866	3	14	love	2024-04-19 00:03:16.309401
867	3	1	angry	2024-04-19 00:03:16.309676
868	3	8	sad	2024-04-19 00:03:16.309955
869	3	10	love	2024-04-19 00:03:16.310243
870	3	3	wow	2024-04-19 00:03:16.310546
871	4	6	haha	2024-04-19 00:03:16.310845
872	4	2	haha	2024-04-19 00:03:16.311155
873	4	4	haha	2024-04-19 00:03:16.31145
874	4	11	haha	2024-04-19 00:03:16.311753
875	4	7	angry	2024-04-19 00:03:16.312034
876	4	14	wow	2024-04-19 00:03:16.312302
877	4	1	like	2024-04-19 00:03:16.312585
878	4	8	wow	2024-04-19 00:03:16.31288
879	4	10	haha	2024-04-19 00:03:16.313167
880	4	3	angry	2024-04-19 00:03:16.313503
881	5	6	wow	2024-04-19 00:03:16.313802
882	5	2	angry	2024-04-19 00:03:16.31411
883	5	4	wow	2024-04-19 00:03:16.314397
884	5	11	like	2024-04-19 00:03:16.314693
885	5	7	haha	2024-04-19 00:03:16.314976
886	5	14	love	2024-04-19 00:03:16.315264
887	5	1	angry	2024-04-19 00:03:16.315549
888	5	8	haha	2024-04-19 00:03:16.315859
889	5	10	angry	2024-04-19 00:03:16.316155
890	5	3	love	2024-04-19 00:03:16.316458
891	6	6	wow	2024-04-19 00:03:16.31677
892	6	2	wow	2024-04-19 00:03:16.317096
893	6	4	angry	2024-04-19 00:03:16.317388
894	6	11	like	2024-04-19 00:03:16.31769
895	6	7	like	2024-04-19 00:03:16.317979
896	6	14	haha	2024-04-19 00:03:16.318276
897	6	1	love	2024-04-19 00:03:16.318571
898	6	8	wow	2024-04-19 00:03:16.318869
899	6	10	haha	2024-04-19 00:03:16.319175
900	6	3	haha	2024-04-19 00:03:16.31946
901	7	6	haha	2024-04-19 00:03:16.319777
902	7	2	like	2024-04-19 00:03:16.320057
903	7	4	angry	2024-04-19 00:03:16.32034
904	7	11	haha	2024-04-19 00:03:16.320625
905	7	7	wow	2024-04-19 00:03:16.321141
906	7	14	sad	2024-04-19 00:03:16.321442
907	7	1	wow	2024-04-19 00:03:16.32174
908	7	8	haha	2024-04-19 00:03:16.322009
909	7	10	love	2024-04-19 00:03:16.322299
910	7	3	love	2024-04-19 00:03:16.322602
911	8	6	angry	2024-04-19 00:03:16.322889
912	8	2	wow	2024-04-19 00:03:16.323164
913	8	4	love	2024-04-19 00:03:16.323461
914	8	11	wow	2024-04-19 00:03:16.323801
915	8	7	love	2024-04-19 00:03:16.324104
916	8	14	like	2024-04-19 00:03:16.324401
917	8	1	sad	2024-04-19 00:03:16.324699
918	8	8	like	2024-04-19 00:03:16.324966
919	8	10	angry	2024-04-19 00:03:16.325258
920	8	3	sad	2024-04-19 00:03:16.325552
921	9	6	sad	2024-04-19 00:03:16.325854
922	9	2	like	2024-04-19 00:03:16.326137
923	9	4	haha	2024-04-19 00:03:16.326428
924	9	11	like	2024-04-19 00:03:16.32671
925	9	7	haha	2024-04-19 00:03:16.327
926	9	14	like	2024-04-19 00:03:16.327293
927	9	1	angry	2024-04-19 00:03:16.327589
928	9	8	haha	2024-04-19 00:03:16.3279
929	9	10	angry	2024-04-19 00:03:16.328222
930	9	3	like	2024-04-19 00:03:16.328527
931	10	6	like	2024-04-19 00:03:16.328819
932	10	2	sad	2024-04-19 00:03:16.329097
933	10	4	wow	2024-04-19 00:03:16.329401
934	10	11	love	2024-04-19 00:03:16.329689
935	10	7	haha	2024-04-19 00:03:16.329981
936	10	14	haha	2024-04-19 00:03:16.330252
937	10	1	haha	2024-04-19 00:03:16.330537
938	10	8	love	2024-04-19 00:03:16.330853
939	10	10	love	2024-04-19 00:03:16.331136
940	10	3	love	2024-04-19 00:03:16.331448
941	1	11	like	2024-04-19 00:03:16.331939
942	1	3	like	2024-04-19 00:03:16.332231
943	1	5	like	2024-04-19 00:03:16.332564
944	1	2	like	2024-04-19 00:03:16.332874
945	1	10	sad	2024-04-19 00:03:16.333158
946	1	12	haha	2024-04-19 00:03:16.333441
947	1	13	wow	2024-04-19 00:03:16.333722
948	1	6	angry	2024-04-19 00:03:16.334012
949	1	1	angry	2024-04-19 00:03:16.334307
950	1	4	haha	2024-04-19 00:03:16.334625
951	1	8	love	2024-04-19 00:03:16.334915
952	2	11	love	2024-04-19 00:03:16.335205
953	2	3	love	2024-04-19 00:03:16.335489
954	2	5	sad	2024-04-19 00:03:16.335795
955	2	2	love	2024-04-19 00:03:16.336089
956	2	10	haha	2024-04-19 00:03:16.336376
957	2	12	sad	2024-04-19 00:03:16.33666
958	2	13	love	2024-04-19 00:03:16.336942
959	2	6	sad	2024-04-19 00:03:16.337255
960	2	1	sad	2024-04-19 00:03:16.337543
961	2	4	love	2024-04-19 00:03:16.337834
962	2	8	wow	2024-04-19 00:03:16.338105
963	3	11	sad	2024-04-19 00:03:16.338388
964	3	3	haha	2024-04-19 00:03:16.338658
965	3	5	haha	2024-04-19 00:03:16.338936
966	3	2	wow	2024-04-19 00:03:16.339243
967	3	10	wow	2024-04-19 00:03:16.339515
968	3	12	angry	2024-04-19 00:03:16.339821
969	3	13	wow	2024-04-19 00:03:16.340103
970	3	6	love	2024-04-19 00:03:16.340362
971	3	1	love	2024-04-19 00:03:16.340822
972	3	4	haha	2024-04-19 00:03:16.341116
973	3	8	wow	2024-04-19 00:03:16.341417
974	4	11	sad	2024-04-19 00:03:16.341718
975	4	3	love	2024-04-19 00:03:16.342003
976	4	5	like	2024-04-19 00:03:16.342287
977	4	2	wow	2024-04-19 00:03:16.342573
978	4	10	love	2024-04-19 00:03:16.342866
979	4	12	love	2024-04-19 00:03:16.343164
980	4	13	love	2024-04-19 00:03:16.343459
981	4	6	sad	2024-04-19 00:03:16.343778
982	4	1	haha	2024-04-19 00:03:16.344067
983	4	4	wow	2024-04-19 00:03:16.344337
984	4	8	sad	2024-04-19 00:03:16.344622
985	5	11	haha	2024-04-19 00:03:16.344911
986	5	3	like	2024-04-19 00:03:16.345186
987	5	5	wow	2024-04-19 00:03:16.345474
988	5	2	wow	2024-04-19 00:03:16.345789
989	5	10	love	2024-04-19 00:03:16.346078
990	5	12	like	2024-04-19 00:03:16.346359
991	5	13	love	2024-04-19 00:03:16.346645
992	5	6	haha	2024-04-19 00:03:16.346985
993	5	1	haha	2024-04-19 00:03:16.347273
994	5	4	wow	2024-04-19 00:03:16.347546
995	5	8	wow	2024-04-19 00:03:16.347841
996	6	11	haha	2024-04-19 00:03:16.348126
997	6	3	angry	2024-04-19 00:03:16.348402
998	6	5	wow	2024-04-19 00:03:16.348674
999	6	2	sad	2024-04-19 00:03:16.34899
1000	6	10	love	2024-04-19 00:03:16.349297
1001	6	12	haha	2024-04-19 00:03:16.349585
1002	6	13	wow	2024-04-19 00:03:16.349873
1003	6	6	sad	2024-04-19 00:03:16.350176
1004	6	1	love	2024-04-19 00:03:16.350457
1005	6	4	wow	2024-04-19 00:03:16.350742
1006	6	8	angry	2024-04-19 00:03:16.351032
1007	7	11	haha	2024-04-19 00:03:16.351307
1008	7	3	sad	2024-04-19 00:03:16.351609
1009	7	5	haha	2024-04-19 00:03:16.351895
1010	7	2	wow	2024-04-19 00:03:16.352171
1011	7	10	wow	2024-04-19 00:03:16.352467
1012	7	12	haha	2024-04-19 00:03:16.352728
1013	7	13	wow	2024-04-19 00:03:16.353036
1014	7	6	sad	2024-04-19 00:03:16.35333
1015	7	1	wow	2024-04-19 00:03:16.353626
1016	7	4	angry	2024-04-19 00:03:16.353918
1017	7	8	wow	2024-04-19 00:03:16.354213
1018	8	11	like	2024-04-19 00:03:16.354508
1019	8	3	haha	2024-04-19 00:03:16.354798
1020	8	5	wow	2024-04-19 00:03:16.355087
1021	8	2	sad	2024-04-19 00:03:16.355372
1022	8	10	wow	2024-04-19 00:03:16.355661
1023	8	12	love	2024-04-19 00:03:16.355954
1024	8	13	haha	2024-04-19 00:03:16.356225
1025	8	6	wow	2024-04-19 00:03:16.356508
1026	8	1	like	2024-04-19 00:03:16.356824
1027	8	4	sad	2024-04-19 00:03:16.357102
1028	8	8	wow	2024-04-19 00:03:16.357384
1029	9	11	wow	2024-04-19 00:03:16.357685
1030	9	3	love	2024-04-19 00:03:16.357983
1031	9	5	love	2024-04-19 00:03:16.35829
1032	9	2	sad	2024-04-19 00:03:16.358582
1033	9	10	like	2024-04-19 00:03:16.358878
1034	9	12	sad	2024-04-19 00:03:16.359156
1035	9	13	angry	2024-04-19 00:03:16.359436
1036	9	6	haha	2024-04-19 00:03:16.359729
1037	9	1	haha	2024-04-19 00:03:16.360241
1038	9	4	haha	2024-04-19 00:03:16.360537
1039	9	8	sad	2024-04-19 00:03:16.360813
1040	10	11	love	2024-04-19 00:03:16.361129
1041	10	3	wow	2024-04-19 00:03:16.361433
1042	10	5	wow	2024-04-19 00:03:16.361724
1043	10	2	haha	2024-04-19 00:03:16.362003
1044	10	10	love	2024-04-19 00:03:16.362293
1045	10	12	wow	2024-04-19 00:03:16.362569
1046	10	13	like	2024-04-19 00:03:16.36285
1047	10	6	haha	2024-04-19 00:03:16.363156
1048	10	1	sad	2024-04-19 00:03:16.363447
1049	10	4	sad	2024-04-19 00:03:16.363728
1050	10	8	love	2024-04-19 00:03:16.364002
\.


--
-- Data for Name: posts; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.posts (post_id, member_id, group_id, photo_url, description, created_at) FROM stdin;
1	8	1	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Ffootball.jpg?alt=media&token=47e76d94-1a3c-4c64-ac93-be9c20bdc15f	1st day practise	2024-04-18 23:37:19.915335
2	3	2	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fbadminton.jpg?alt=media&token=65167afe-1e56-4e1b-8fe7-9cfae50ae484	Enjoy makk	2024-04-18 23:37:19.931423
3	1	3	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fonlinecourse.jpeg?alt=media&token=717da4a4-9049-41d2-8ffc-448d81f31c86	Wanna sleep T_T	2024-04-18 23:37:19.937842
4	1	4	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fprogramming.JPG?alt=media&token=dbd8f503-2fea-43ec-9599-c1df1ac10515	Grinding	2024-04-18 23:37:19.944845
5	6	5	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fmidterm.jpeg?alt=media&token=2a767e05-667c-4898-9175-e8729fe20c9b	Cheat sheet done!	2024-04-18 23:37:19.950624
6	10	6	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Ffitness.JPG?alt=media&token=73336d8e-a753-44dc-89ff-d751d7aacf69	Six packs is coming	2024-04-18 23:37:19.954203
7	3	7	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fmovies.jpeg?alt=media&token=44906b86-8c27-4ab1-b46f-96dcfe9cfab3	I cried so hard T_T	2024-04-18 23:37:19.95946
8	5	8	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Fseries.jpeg?alt=media&token=84f76ffd-fdc7-4af1-9ede-c12dc425b9cb	Today I skipped K-drama na	2024-04-18 23:37:19.963526
9	8	9	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2FMusic.jpg?alt=media&token=2184d287-59d2-4cef-8c31-33b92e44a8b5	SRV is my tune	2024-04-18 23:37:19.967681
10	13	10	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/postphoto%2Fmockpost%2Ftravel.jpeg?alt=media&token=6e02713a-d086-4eca-bf69-257a3ebab442	The weather is so fresh here!	2024-04-18 23:37:19.971236
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: streak_set; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.streak_set (streak_set_id, group_id, member_id, end_date, start_date) FROM stdin;
1	1	2	2024-04-26 06:59:59	2024-04-18 23:01:56.659818
2	1	14	2024-04-26 06:59:59	2024-04-18 23:01:56.669902
3	1	10	2024-04-26 06:59:59	2024-04-18 23:02:39.591351
4	1	4	2024-04-26 06:59:59	2024-04-18 23:02:39.597942
5	1	8	2024-04-26 06:59:59	2024-04-18 23:02:39.606156
6	1	11	2024-04-26 06:59:59	2024-04-18 23:02:39.613036
7	1	3	2024-04-26 06:59:59	2024-04-18 23:02:39.617895
8	1	13	2024-04-26 06:59:59	2024-04-18 23:02:39.62577
9	1	6	2024-04-26 06:59:59	2024-04-18 23:02:39.630721
10	1	7	2024-04-26 06:59:59	2024-04-18 23:02:39.63537
11	2	3	2024-04-26 06:59:59	2024-04-18 23:02:39.641722
12	2	1	2024-04-26 06:59:59	2024-04-18 23:02:39.646702
13	2	9	2024-04-26 06:59:59	2024-04-18 23:02:39.651028
14	2	6	2024-04-26 06:59:59	2024-04-18 23:02:49.366821
15	2	4	2024-04-26 06:59:59	2024-04-18 23:02:49.374088
16	2	8	2024-04-26 06:59:59	2024-04-18 23:02:49.380731
17	2	10	2024-04-26 06:59:59	2024-04-18 23:02:49.387305
18	2	14	2024-04-26 06:59:59	2024-04-18 23:02:49.393195
19	2	12	2024-04-26 06:59:59	2024-04-18 23:02:49.398141
20	2	11	2024-04-26 06:59:59	2024-04-18 23:02:49.403206
21	2	7	2024-04-26 06:59:59	2024-04-18 23:02:49.408511
22	3	1	2024-04-26 06:59:59	2024-04-18 23:02:49.414627
23	3	2	2024-04-26 06:59:59	2024-04-18 23:02:49.41904
24	3	14	2024-04-26 06:59:59	2024-04-18 23:02:49.423881
25	3	3	2024-04-26 06:59:59	2024-04-18 23:02:49.428561
26	3	12	2024-04-26 06:59:59	2024-04-18 23:02:49.432368
27	3	6	2024-04-26 06:59:59	2024-04-18 23:02:49.436571
28	3	7	2024-04-26 06:59:59	2024-04-18 23:02:49.440723
29	3	13	2024-04-26 06:59:59	2024-04-18 23:02:49.444586
30	3	9	2024-04-26 06:59:59	2024-04-18 23:02:49.448185
31	4	2	2024-04-26 06:59:59	2024-04-18 23:03:08.368643
32	4	3	2024-04-26 06:59:59	2024-04-18 23:03:08.378105
33	4	1	2024-04-26 06:59:59	2024-04-18 23:03:30.340597
34	4	12	2024-04-26 06:59:59	2024-04-18 23:03:30.348105
35	4	5	2024-04-26 06:59:59	2024-04-18 23:03:30.35439
36	4	7	2024-04-26 06:59:59	2024-04-18 23:03:30.360771
37	4	13	2024-04-26 06:59:59	2024-04-18 23:03:30.365571
38	4	4	2024-04-26 06:59:59	2024-04-18 23:03:30.371886
39	4	8	2024-04-26 06:59:59	2024-04-18 23:03:30.377245
40	5	3	2024-04-26 06:59:59	2024-04-18 23:03:30.382985
41	5	14	2024-04-26 06:59:59	2024-04-18 23:03:30.387776
42	5	1	2024-04-26 06:59:59	2024-04-18 23:03:30.392707
43	5	6	2024-04-26 06:59:59	2024-04-18 23:03:30.396818
44	5	9	2024-04-26 06:59:59	2024-04-18 23:03:30.401191
45	5	2	2024-04-26 06:59:59	2024-04-18 23:03:30.404982
46	5	12	2024-04-26 06:59:59	2024-04-18 23:03:42.984124
47	5	10	2024-04-26 06:59:59	2024-04-18 23:03:42.992357
48	5	11	2024-04-26 06:59:59	2024-04-18 23:03:42.998536
49	5	13	2024-04-26 06:59:59	2024-04-18 23:03:43.002387
50	6	3	2024-04-26 06:59:59	2024-04-18 23:03:43.009227
51	6	1	2024-04-26 06:59:59	2024-04-18 23:03:55.309237
52	6	10	2024-04-26 06:59:59	2024-04-18 23:03:55.316204
53	6	4	2024-04-26 06:59:59	2024-04-18 23:03:55.322317
54	6	14	2024-04-26 06:59:59	2024-04-18 23:03:55.32845
55	6	6	2024-04-26 06:59:59	2024-04-18 23:03:55.333869
56	6	7	2024-04-26 06:59:59	2024-04-18 23:03:55.338335
57	6	2	2024-04-26 06:59:59	2024-04-18 23:03:55.343683
58	6	11	2024-04-26 06:59:59	2024-04-18 23:03:55.347887
59	7	14	2024-04-26 06:59:59	2024-04-18 23:03:55.353093
60	7	8	2024-04-26 06:59:59	2024-04-18 23:03:55.357805
61	7	4	2024-04-26 06:59:59	2024-04-18 23:03:55.362112
62	7	12	2024-04-26 06:59:59	2024-04-18 23:03:55.366004
63	7	5	2024-04-26 06:59:59	2024-04-18 23:03:55.369729
64	7	2	2024-04-26 06:59:59	2024-04-18 23:03:55.373801
65	7	1	2024-04-26 06:59:59	2024-04-18 23:03:55.377887
66	7	13	2024-04-26 06:59:59	2024-04-18 23:03:55.381349
67	7	3	2024-04-26 06:59:59	2024-04-18 23:03:55.385099
68	7	10	2024-04-26 06:59:59	2024-04-18 23:03:55.388551
69	8	1	2024-04-26 06:59:59	2024-04-18 23:03:55.392716
70	8	9	2024-04-26 06:59:59	2024-04-18 23:03:55.396111
71	8	6	2024-04-26 06:59:59	2024-04-18 23:03:55.399455
72	8	2	2024-04-26 06:59:59	2024-04-18 23:03:55.402349
73	8	3	2024-04-26 06:59:59	2024-04-18 23:03:55.404989
74	8	7	2024-04-26 06:59:59	2024-04-18 23:03:55.407514
75	8	12	2024-04-26 06:59:59	2024-04-18 23:03:55.410314
76	8	5	2024-04-26 06:59:59	2024-04-18 23:03:55.41318
77	9	2	2024-04-26 06:59:59	2024-04-18 23:03:55.416591
78	9	4	2024-04-26 06:59:59	2024-04-18 23:03:55.419325
79	9	11	2024-04-26 06:59:59	2024-04-18 23:03:55.421989
80	9	7	2024-04-26 06:59:59	2024-04-18 23:03:55.424585
81	9	14	2024-04-26 06:59:59	2024-04-18 23:04:05.308332
82	9	1	2024-04-26 06:59:59	2024-04-18 23:04:05.314969
83	9	8	2024-04-26 06:59:59	2024-04-18 23:04:05.319554
84	9	10	2024-04-26 06:59:59	2024-04-18 23:04:05.324526
85	9	3	2024-04-26 06:59:59	2024-04-18 23:04:05.329602
86	10	3	2024-04-26 06:59:59	2024-04-18 23:04:05.334918
87	10	5	2024-04-26 06:59:59	2024-04-18 23:04:05.339325
88	10	2	2024-04-26 06:59:59	2024-04-18 23:04:05.343776
89	10	10	2024-04-26 06:59:59	2024-04-18 23:04:05.34775
90	10	12	2024-04-26 06:59:59	2024-04-18 23:04:05.351918
91	10	13	2024-04-26 06:59:59	2024-04-18 23:04:05.356153
92	10	6	2024-04-26 06:59:59	2024-04-18 23:04:05.359596
93	10	1	2024-04-26 06:59:59	2024-04-18 23:04:05.36326
94	10	4	2024-04-26 06:59:59	2024-04-18 23:04:14.460856
95	10	8	2024-04-26 06:59:59	2024-04-18 23:04:14.464453
\.


--
-- Data for Name: streaks; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.streaks (streak_id, streak_set_id, max_streak_count, total_streak_count, weekly_streak_count, completed, recent_date_added, created_at) FROM stdin;
1	1	0	0	0	f	\N	2024-04-18 23:01:56.662094
2	2	0	0	0	f	\N	2024-04-18 23:01:56.670518
3	3	0	0	0	f	\N	2024-04-18 23:02:39.592784
4	4	0	0	0	f	\N	2024-04-18 23:02:39.599158
6	6	0	0	0	f	\N	2024-04-18 23:02:39.613585
7	7	0	0	0	f	\N	2024-04-18 23:02:39.618609
8	8	0	0	0	f	\N	2024-04-18 23:02:39.626498
9	9	0	0	0	f	\N	2024-04-18 23:02:39.631344
10	10	0	0	0	f	\N	2024-04-18 23:02:39.636044
12	12	0	0	0	f	\N	2024-04-18 23:02:39.647277
13	13	0	0	0	f	\N	2024-04-18 23:02:39.651581
14	14	0	0	0	f	\N	2024-04-18 23:02:49.368215
15	15	0	0	0	f	\N	2024-04-18 23:02:49.375043
16	16	0	0	0	f	\N	2024-04-18 23:02:49.381518
17	17	0	0	0	f	\N	2024-04-18 23:02:49.388219
18	18	0	0	0	f	\N	2024-04-18 23:02:49.394081
19	19	0	0	0	f	\N	2024-04-18 23:02:49.398827
20	20	0	0	0	f	\N	2024-04-18 23:02:49.403922
21	21	0	0	0	f	\N	2024-04-18 23:02:49.409239
23	23	0	0	0	f	\N	2024-04-18 23:02:49.419643
24	24	0	0	0	f	\N	2024-04-18 23:02:49.424485
25	25	0	0	0	f	\N	2024-04-18 23:02:49.429105
26	26	0	0	0	f	\N	2024-04-18 23:02:49.432885
27	27	0	0	0	f	\N	2024-04-18 23:02:49.437127
28	28	0	0	0	f	\N	2024-04-18 23:02:49.44122
29	29	0	0	0	f	\N	2024-04-18 23:02:49.445102
30	30	0	0	0	f	\N	2024-04-18 23:02:49.448548
31	31	0	0	0	f	\N	2024-04-18 23:03:08.369781
32	32	0	0	0	f	\N	2024-04-18 23:03:08.379037
34	34	0	0	0	f	\N	2024-04-18 23:03:30.348839
35	35	0	0	0	f	\N	2024-04-18 23:03:30.355372
36	36	0	0	0	f	\N	2024-04-18 23:03:30.361718
37	37	0	0	0	f	\N	2024-04-18 23:03:30.366865
38	38	0	0	0	f	\N	2024-04-18 23:03:30.372659
39	39	0	0	0	f	\N	2024-04-18 23:03:30.37793
40	40	0	0	0	f	\N	2024-04-18 23:03:30.383484
41	41	0	0	0	f	\N	2024-04-18 23:03:30.388356
42	42	0	0	0	f	\N	2024-04-18 23:03:30.393265
44	44	0	0	0	f	\N	2024-04-18 23:03:30.401713
45	45	0	0	0	f	\N	2024-04-18 23:03:30.405479
46	46	0	0	0	f	\N	2024-04-18 23:03:42.98578
47	47	0	0	0	f	\N	2024-04-18 23:03:42.99348
48	48	0	0	0	f	\N	2024-04-18 23:03:42.99901
49	49	0	0	0	f	\N	2024-04-18 23:03:43.003202
50	50	0	0	0	f	\N	2024-04-18 23:03:43.009756
51	51	0	0	0	f	\N	2024-04-18 23:03:55.310596
53	53	0	0	0	f	\N	2024-04-18 23:03:55.323121
54	54	0	0	0	f	\N	2024-04-18 23:03:55.329173
55	55	0	0	0	f	\N	2024-04-18 23:03:55.334455
56	56	0	0	0	f	\N	2024-04-18 23:03:55.339021
57	57	0	0	0	f	\N	2024-04-18 23:03:55.344307
58	58	0	0	0	f	\N	2024-04-18 23:03:55.348366
59	59	0	0	0	f	\N	2024-04-18 23:03:55.353751
60	60	0	0	0	f	\N	2024-04-18 23:03:55.358414
61	61	0	0	0	f	\N	2024-04-18 23:03:55.362599
62	62	0	0	0	f	\N	2024-04-18 23:03:55.366516
63	63	0	0	0	f	\N	2024-04-18 23:03:55.370278
64	64	0	0	0	f	\N	2024-04-18 23:03:55.374393
65	65	0	0	0	f	\N	2024-04-18 23:03:55.378382
66	66	0	0	0	f	\N	2024-04-18 23:03:55.381731
68	68	0	0	0	f	\N	2024-04-18 23:03:55.389001
69	69	0	0	0	f	\N	2024-04-18 23:03:55.393145
70	70	0	0	0	f	\N	2024-04-18 23:03:55.396568
71	71	0	0	0	f	\N	2024-04-18 23:03:55.399842
72	72	0	0	0	f	\N	2024-04-18 23:03:55.402678
73	73	0	0	0	f	\N	2024-04-18 23:03:55.405363
74	74	0	0	0	f	\N	2024-04-18 23:03:55.407851
75	75	0	0	0	f	\N	2024-04-18 23:03:55.41068
77	77	0	0	0	f	\N	2024-04-18 23:03:55.416978
78	78	0	0	0	f	\N	2024-04-18 23:03:55.41969
79	79	0	0	0	f	\N	2024-04-18 23:03:55.422348
80	80	0	0	0	f	\N	2024-04-18 23:03:55.424921
81	81	0	0	0	f	\N	2024-04-18 23:04:05.309516
82	82	0	0	0	f	\N	2024-04-18 23:04:05.315531
84	84	0	0	0	f	\N	2024-04-18 23:04:05.325274
85	85	0	0	0	f	\N	2024-04-18 23:04:05.330198
86	86	0	0	0	f	\N	2024-04-18 23:04:05.33553
87	87	0	0	0	f	\N	2024-04-18 23:04:05.339859
88	88	0	0	0	f	\N	2024-04-18 23:04:05.344401
89	89	0	0	0	f	\N	2024-04-18 23:04:05.348276
90	90	0	0	0	f	\N	2024-04-18 23:04:05.352657
92	92	0	0	0	f	\N	2024-04-18 23:04:05.360243
93	93	0	0	0	f	\N	2024-04-18 23:04:05.363691
94	94	0	0	0	f	\N	2024-04-18 23:04:14.461349
95	95	0	0	0	f	\N	2024-04-18 23:04:14.465018
52	52	2	2	2	t	2024-04-18 23:37:19.955784	2024-04-18 23:03:55.317089
67	67	2	2	2	t	2024-04-18 23:37:19.961105	2024-04-18 23:03:55.385586
76	76	2	2	2	t	2024-04-18 23:37:19.964968	2024-04-18 23:03:55.413537
83	83	2	2	2	t	2024-04-18 23:37:19.968699	2024-04-18 23:04:05.320431
91	91	2	2	2	t	2024-04-18 23:37:19.972657	2024-04-18 23:04:05.356579
5	5	6	6	6	t	2024-04-18 23:37:19.925352	2024-04-18 23:02:39.607249
11	11	2	2	2	t	2024-04-18 23:37:19.933544	2024-04-18 23:02:39.642423
22	22	1	1	1	t	2024-04-18 23:37:19.941016	2024-04-18 23:02:49.415218
33	33	2	2	2	t	2024-04-18 23:37:19.947317	2024-04-18 23:03:30.341614
43	43	2	2	2	t	2024-04-18 23:37:19.951683	2024-04-18 23:03:30.397355
\.


--
-- Data for Name: tag_category; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.tag_category (category_id, category_name) FROM stdin;
1	Sports and Fitness
2	Learning and Development
3	Health and Wellness
4	Entertainment and Media
5	Hobbies and Leisure
6	Others
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.tags (tag_id, tag_name, icon_url, category_id) FROM stdin;
1	Dietary	\N	3
2	Bulking	\N	3
3	Vegan and J	\N	3
4	Meditation	\N	3
8	Theater	\N	4
9	Podcasts	\N	4
10	Cooking	\N	5
11	Baking	\N	5
12	Gardening	\N	5
13	Planting	\N	5
14	Knitting	\N	5
15	Pottery	\N	5
16	Caligraphy	\N	5
18	Board games	\N	5
20	Rock Climbing	\N	1
21	Basketball	\N	1
22	Volleyball	\N	1
23	Golf	\N	1
24	Boxing	\N	1
26	Bowling	\N	1
27	Ice skating	\N	1
28	Racquet	\N	1
29	Tennis	\N	1
30	Table tennis	\N	1
31	Snooker	\N	1
32	Pool	\N	1
33	Swimming	\N	1
34	Running	\N	1
35	Yoga and Pilates	\N	1
36	Karate	\N	1
37	Taekwondo	\N	1
38	Hiking	\N	1
39	Cycling	\N	1
40	Hockey	\N	1
41	Figure Skating	\N	1
42	Skiing	\N	1
44	Exam prep	\N	2
45	Investing	\N	2
47	Language	\N	2
48	Public speaking	\N	2
49	SAT	\N	2
50	IELTS	\N	2
52	Final exam	\N	2
43	Online courses	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2Fcover_hu3d03a01dcc18bc5be0e67db3d8d209a6_184197_640x0_resize_q90_lanczos.ea326e175ea41a8a8013658ca8ec88b43ed1a14e0e9902f5ab052f97182095a2.jpg?alt=media&token=732287d8-9f5d-4329-bd89-1296020eab05	2
46	Programming	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2Fcomputer-coding.webp?alt=media&token=d1e25f86-f955-427b-a0d0-ebea0e46ea88	2
51	Midterm exam	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2Fimages.jpeg?alt=media&token=a608a396-ec92-4d43-802d-0951dd34b7ad	2
19	Football	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2FFootball_Pallo_valmiina-cropped.jpg?alt=media&token=81b5f42a-7fdc-41c8-be33-e553ea2aa9c6	1
25	Badminton	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2Fbadminton-iStock-1.jpg?alt=media&token=64fb3a22-1f10-4a0d-900a-8076ce0e5283	1
53	Fitness and Gym	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2Flimited-budget-gym-01.png?alt=media&token=d06773ca-0a63-42ee-a1df-78920a20d86f	3
6	Series	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2FMain_Thumb_0708.jpg?alt=media&token=d5892861-92a2-4659-817d-ac14bf11e5d2	4
5	Movies	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2Fcinema-rex.jpg?alt=media&token=5d3f42dc-631d-4598-b6df-8984e18601f1	4
7	Music	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2Fimages%20(1).jpeg?alt=media&token=4923773a-380b-45ad-b943-7b81e264a0c9	4
17	Travelling	https://firebasestorage.googleapis.com/v0/b/gleam-firebase-6925b.appspot.com/o/tagphoto%2Fimages%20(2).jpeg?alt=media&token=d9cac3b0-db70-47da-839e-766d48a05d50	5
\.


--
-- Name: groups_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.groups_group_id_seq', 10, true);


--
-- Name: post_comments_comment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.post_comments_comment_id_seq', 48, true);


--
-- Name: post_reactions_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.post_reactions_post_id_seq', 1, false);


--
-- Name: post_reactions_reaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.post_reactions_reaction_id_seq', 1050, true);


--
-- Name: posts_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.posts_group_id_seq', 1, false);


--
-- Name: posts_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.posts_post_id_seq', 10, true);


--
-- Name: streak_set_streak_set_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.streak_set_streak_set_id_seq', 95, true);


--
-- Name: streaks_streak_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.streaks_streak_id_seq', 95, true);


--
-- Name: tag_category_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.tag_category_category_id_seq', 6, true);


--
-- Name: tags_tag_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.tags_tag_id_seq', 53, true);


--
-- Name: group_members group_members_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.group_members
    ADD CONSTRAINT group_members_pkey PRIMARY KEY (group_id, member_id);


--
-- Name: group_requests group_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.group_requests
    ADD CONSTRAINT group_requests_pkey PRIMARY KEY (group_id, member_id);


--
-- Name: groups groups_group_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_group_name_key UNIQUE (group_name);


--
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (group_id);


--
-- Name: post_comments post_comments_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post_comments
    ADD CONSTRAINT post_comments_pkey PRIMARY KEY (comment_id);


--
-- Name: post_reactions post_reactions_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post_reactions
    ADD CONSTRAINT post_reactions_pkey PRIMARY KEY (reaction_id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (post_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: streak_set streak_set_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.streak_set
    ADD CONSTRAINT streak_set_pkey PRIMARY KEY (streak_set_id);


--
-- Name: streaks streaks_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.streaks
    ADD CONSTRAINT streaks_pkey PRIMARY KEY (streak_id);


--
-- Name: tag_category tag_category_category_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tag_category
    ADD CONSTRAINT tag_category_category_name_key UNIQUE (category_name);


--
-- Name: tag_category tag_category_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tag_category
    ADD CONSTRAINT tag_category_pkey PRIMARY KEY (category_id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (tag_id);


--
-- Name: tags tags_tag_name_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_tag_name_key UNIQUE (tag_name);


--
-- Name: group_members group_members_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.group_members
    ADD CONSTRAINT group_members_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id) ON DELETE CASCADE;


--
-- Name: group_requests group_requests_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.group_requests
    ADD CONSTRAINT group_requests_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id) ON DELETE CASCADE;


--
-- Name: groups groups_tag_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES public.tags(tag_id);


--
-- Name: post_comments post_comments_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post_comments
    ADD CONSTRAINT post_comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON DELETE CASCADE;


--
-- Name: post_reactions post_reactions_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.post_reactions
    ADD CONSTRAINT post_reactions_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(post_id) ON DELETE CASCADE;


--
-- Name: posts posts_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id) ON DELETE CASCADE;


--
-- Name: streak_set streak_set_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.streak_set
    ADD CONSTRAINT streak_set_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(group_id) ON DELETE CASCADE;


--
-- Name: streaks streaks_streak_set_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.streaks
    ADD CONSTRAINT streaks_streak_set_id_fkey FOREIGN KEY (streak_set_id) REFERENCES public.streak_set(streak_set_id) ON DELETE CASCADE;


--
-- Name: tags tags_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.tag_category(category_id);


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
-- PostgreSQL database cluster dump complete
--

