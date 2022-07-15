--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Ubuntu 14.4-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.4 (Ubuntu 14.4-1.pgdg20.04+1)

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: afisa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.afisa (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    image_path character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.afisa OWNER TO postgres;

--
-- Name: banner; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.banner (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    image_path character varying,
    url text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.banner OWNER TO postgres;

--
-- Name: brends; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brends (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying,
    image_path character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.brends OWNER TO postgres;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    parent_category_id uuid,
    image_path character varying,
    is_home_category boolean,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- Name: category_product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category_product (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    category_id uuid,
    product_id uuid,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.category_product OWNER TO postgres;

--
-- Name: category_shop; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.category_shop (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    category_id uuid,
    shop_id uuid,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.category_shop OWNER TO postgres;

--
-- Name: company_address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.company_address (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    address character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.company_address OWNER TO postgres;

--
-- Name: company_phone; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.company_phone (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    phone character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.company_phone OWNER TO postgres;

--
-- Name: company_setting; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.company_setting (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    logo_path character varying,
    favicon_path character varying,
    email character varying,
    instagram character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp without time zone
);


ALTER TABLE public.company_setting OWNER TO postgres;

--
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    full_name character varying,
    phone_number character varying,
    password character varying,
    birthday date,
    gender character varying,
    addresses character varying[],
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- Name: district; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.district (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    price numeric,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.district OWNER TO postgres;

--
-- Name: languages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.languages (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name_short character varying(5),
    flag_path character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.languages OWNER TO postgres;

--
-- Name: likes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.likes (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    product_id uuid,
    customer_id uuid,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.likes OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    brend_id uuid,
    price numeric,
    old_price numeric,
    amount bigint,
    product_code character varying,
    main_image_path character varying,
    image_paths character varying[],
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: shops; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shops (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    owner_name character varying,
    address character varying,
    phone_number character varying,
    number_of_goods integer,
    running_time character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.shops OWNER TO postgres;

--
-- Name: translation_about; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_about (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    title character varying DEFAULT 'uytget'::character varying,
    content text DEFAULT 'uytget'::text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_about OWNER TO postgres;

--
-- Name: translation_afisa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_afisa (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    afisa_id uuid,
    lang_id uuid,
    title character varying DEFAULT 'uytget'::character varying,
    description text DEFAULT 'uytget'::text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_afisa OWNER TO postgres;

--
-- Name: translation_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_category (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    category_id uuid,
    name character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_category OWNER TO postgres;

--
-- Name: translation_contact; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_contact (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    full_name character varying DEFAULT 'uytget'::character varying,
    email character varying DEFAULT 'uytget'::character varying,
    phone character varying DEFAULT 'uytget'::character varying,
    letter character varying DEFAULT 'uytget'::character varying,
    company_phone character varying DEFAULT 'uytget'::character varying,
    imo character varying DEFAULT 'uytget'::character varying,
    company_email character varying DEFAULT 'uytget'::character varying,
    instagram character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    button_text character varying
);


ALTER TABLE public.translation_contact OWNER TO postgres;

--
-- Name: translation_district; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_district (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    district_id uuid,
    name character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_district OWNER TO postgres;

--
-- Name: translation_footer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_footer (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    about character varying DEFAULT 'uytget'::character varying,
    payment character varying DEFAULT 'uytget'::character varying,
    contact character varying DEFAULT 'uytget'::character varying,
    secure character varying DEFAULT 'uytget'::character varying,
    word character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_footer OWNER TO postgres;

--
-- Name: translation_header; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_header (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid DEFAULT public.uuid_generate_v4(),
    research character varying DEFAULT 'uytget'::character varying,
    phone character varying DEFAULT 'uytget'::character varying,
    password character varying DEFAULT 'uytget'::character varying,
    forgot_password character varying DEFAULT 'uytget'::character varying,
    sign_in character varying DEFAULT 'uytget'::character varying,
    sign_up character varying DEFAULT 'uytget'::character varying,
    name character varying DEFAULT 'uytget'::character varying,
    password_verification character varying DEFAULT 'uytget'::character varying,
    verify_secure character varying DEFAULT 'uytget'::character varying,
    my_information character varying DEFAULT 'uytget'::character varying,
    my_favorites character varying DEFAULT 'uytget'::character varying,
    my_orders character varying DEFAULT 'uytget'::character varying,
    log_out character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_header OWNER TO postgres;

--
-- Name: translation_my_information_page; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_my_information_page (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    address character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    birthday character varying DEFAULT 'uytget'::character varying,
    update_password character varying DEFAULT 'uytegt'::character varying,
    save character varying DEFAULT 'uytegt'::character varying
);


ALTER TABLE public.translation_my_information_page OWNER TO postgres;

--
-- Name: translation_payment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_payment (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    title character varying DEFAULT 'uytget'::character varying,
    content text DEFAULT 'uytget'::text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_payment OWNER TO postgres;

--
-- Name: translation_product; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_product (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    product_id uuid,
    name character varying DEFAULT 'uytget'::character varying,
    description text DEFAULT 'uytget'::text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_product OWNER TO postgres;

--
-- Name: translation_secure; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_secure (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    title character varying DEFAULT 'uytget'::character varying,
    content text DEFAULT 'uytget'::text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_secure OWNER TO postgres;

--
-- Name: translation_update_password_page; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_update_password_page (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    title character varying DEFAULT 'uytget'::character varying,
    verify_password character varying DEFAULT 'uytget'::character varying,
    explanation character varying DEFAULT 'uytget'::character varying,
    save character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    password character varying DEFAULT 'uytget'::character varying
);


ALTER TABLE public.translation_update_password_page OWNER TO postgres;

--
-- Data for Name: afisa; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.afisa (id, image_path, created_at, updated_at, deleted_at) FROM stdin;
cb670531-8ca3-4d74-8ea4-f7853aae4132		2022-06-23 18:04:26.258751+05	2022-06-23 18:04:26.258751+05	\N
\.


--
-- Data for Name: banner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banner (id, image_path, url, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: brends; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.brends (id, name, image_path, created_at, updated_at, deleted_at) FROM stdin;
c2f11fea-8057-4a7c-b290-c97d1d90ac59	Markow	uploads/brend9be78774-985a-4d82-bd2c-e73fd7605c2e.jpeg	2022-06-16 14:13:06.09796+05	2022-06-16 14:13:06.09796+05	\N
6a31c50a-704f-4b0d-80ae-240ca3094cda	Algida	uploads/brendbf1a1059-508d-48a0-9cab-e9c0ff52ea82.jpg	2022-06-16 14:13:49.98051+05	2022-06-16 14:13:49.98051+05	\N
214be879-65c3-4710-86b4-3fc3bce2e974	Arcalyk	uploads/brend24badfac-613d-4aa3-881b-952bd14994b5.jpeg	2022-06-16 14:14:05.416191+05	2022-06-16 14:14:05.416191+05	\N
ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	Tut	uploads/brend4f68381a-aa73-4168-90b3-66c1a17cd5c5.jpeg	2022-06-16 14:14:25.908903+05	2022-06-16 14:14:25.908903+05	\N
fdd259c2-794a-42b9-a3ad-9e91502af23e	Koka Kola	uploads/brend75f655c6-bcf5-47b2-ba01-d112cba64e81.jpg	2022-07-12 17:54:39.242004+05	2022-07-12 17:54:39.242004+05	\N
f53a27b4-7810-4d8f-bd45-edad405d92b9	Maral Koke	uploads/brend7827fcfe-f8a9-4747-8c34-b55af2488b29.jpeg	2022-07-12 17:57:46.472194+05	2022-07-12 17:57:46.472194+05	\N
46b13f0a-d584-4ad3-b270-437ecdc51449	Taze Ay	uploads/brend993b6484-657d-4662-abe2-922170abe75b.jpeg	2022-07-12 18:16:12.889441+05	2022-07-12 18:16:12.889441+05	\N
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, parent_category_id, image_path, is_home_category, created_at, updated_at, deleted_at) FROM stdin;
3d0851a3-45cf-4662-b315-67a98b9273c2	\N	uploads/category302c94f3-d485-4cb9-8bbb-2dca8eefb87a.jpeg	f	2022-06-16 12:59:15.624996+05	2022-06-16 12:59:15.624996+05	\N
1b94e399-f0f1-435f-92c5-f855889c1683	3d0851a3-45cf-4662-b315-67a98b9273c2		f	2022-06-16 13:32:50.82101+05	2022-06-16 13:32:50.82101+05	\N
02bd4413-8586-49ab-802e-16304e756a8b	\N	uploads/category8329e720-4169-4564-8abc-82ef79fbcbfe.jpeg	f	2022-06-16 13:43:22.644619+05	2022-06-16 13:43:22.644619+05	\N
b982bd86-0a0f-4950-baad-5a131e9b728e	02bd4413-8586-49ab-802e-16304e756a8b		f	2022-06-16 13:44:16.430875+05	2022-06-16 13:44:16.430875+05	\N
33cf0893-ff6e-40b3-b50f-2a3e926eca70	b982bd86-0a0f-4950-baad-5a131e9b728e		f	2022-06-16 13:44:51.155881+05	2022-06-16 13:44:51.155881+05	\N
f745d171-68e6-42e2-b339-cb3c210cda55	b982bd86-0a0f-4950-baad-5a131e9b728e		f	2022-06-16 13:45:48.828786+05	2022-06-16 13:45:48.828786+05	\N
5bb9a4e7-9992-418f-b551-537844d371da	02bd4413-8586-49ab-802e-16304e756a8b		f	2022-06-16 13:46:44.575803+05	2022-06-16 13:46:44.575803+05	\N
fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	5bb9a4e7-9992-418f-b551-537844d371da		f	2022-06-16 13:47:18.854741+05	2022-06-16 13:47:18.854741+05	\N
d4cb1359-6c23-4194-8e3c-21ed8cec8373	5bb9a4e7-9992-418f-b551-537844d371da		f	2022-06-16 13:48:04.517774+05	2022-06-16 13:48:04.517774+05	\N
f2e02bbb-c554-4989-b315-8d6aa0575bfa	3d0851a3-45cf-4662-b315-67a98b9273c2		f	2022-06-16 13:48:59.276619+05	2022-06-16 13:48:59.276619+05	\N
4af3388b-2738-4ff6-b42e-927cb0ff897f	\N	uploads/categoryf3278014-450d-4d6f-a1f6-8f9b60652821.jpeg	f	2022-06-16 13:49:43.627092+05	2022-06-16 13:49:43.627092+05	\N
56b86071-1c45-490b-a683-a8898546f179	4af3388b-2738-4ff6-b42e-927cb0ff897f		f	2022-06-16 13:50:35.295305+05	2022-06-16 13:50:35.295305+05	\N
ec0b10ac-bf81-4ae3-881e-ef616ea13d7f	56b86071-1c45-490b-a683-a8898546f179		f	2022-06-16 13:51:05.800143+05	2022-06-16 13:51:05.800143+05	\N
0a1963a2-4084-403e-871d-763ae4825fab	56b86071-1c45-490b-a683-a8898546f179		f	2022-06-16 13:51:55.953017+05	2022-06-16 13:51:55.953017+05	\N
5d877898-9ef4-4b91-8518-193b431228a8	4af3388b-2738-4ff6-b42e-927cb0ff897f		f	2022-06-16 13:52:45.112373+05	2022-06-16 13:52:45.112373+05	\N
fc87c4c5-d7cb-4def-a0e0-11cd5751e04b	5d877898-9ef4-4b91-8518-193b431228a8		f	2022-06-16 13:53:34.255795+05	2022-06-16 13:53:34.255795+05	\N
38d92a87-4a9c-4860-94e6-e568f21ecd8e	5d877898-9ef4-4b91-8518-193b431228a8		f	2022-06-16 13:54:05.320424+05	2022-06-16 13:54:05.320424+05	\N
7f453dd0-7b2e-480d-a8be-fcfa23bd863e	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:43:07.336084+05	2022-06-20 09:43:07.336084+05	\N
29ed85bb-11eb-4458-bbf3-5a5644d167d6	\N	uploads/categoryeaae1626-7e9f-4db9-abf6-f454ade813d3.jpeg	f	2022-06-20 09:41:17.575565+05	2022-06-20 09:41:17.575565+05	\N
66772380-c161-4c45-9350-a45e765193e2	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:45:34.38667+05	2022-06-20 09:45:34.38667+05	\N
338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:46:01.119337+05	2022-06-20 09:46:01.119337+05	\N
45765130-7f97-4f0c-b886-f70b75e02610	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 10:11:06.648938+05	2022-06-20 10:11:06.648938+05	\N
\.


--
-- Data for Name: category_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category_product (id, category_id, product_id, created_at, updated_at, deleted_at) FROM stdin;
c5878350-4b13-47d9-a589-802d69204c7b	29ed85bb-11eb-4458-bbf3-5a5644d167d6	525af569-06b6-440a-ab5a-6ee0b39cf51d	2022-06-20 12:37:45.324642+05	2022-06-20 12:37:45.324642+05	\N
dd6045cb-907a-4a78-9433-a819b6b20bae	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	525af569-06b6-440a-ab5a-6ee0b39cf51d	2022-06-20 12:37:45.335485+05	2022-06-20 12:37:45.335485+05	\N
56b0f1ef-34c7-4389-a795-c716f7235c8b	29ed85bb-11eb-4458-bbf3-5a5644d167d6	0d356bb1-c695-4b29-b199-bdff967abfe2	2022-06-20 12:40:33.305124+05	2022-06-20 12:40:33.305124+05	\N
7261d637-0cff-4d64-b3b0-adb7bc3b6215	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	0d356bb1-c695-4b29-b199-bdff967abfe2	2022-06-20 12:40:33.316715+05	2022-06-20 12:40:33.316715+05	\N
c69ef881-0203-4fb4-9ec7-c53ed3a5be1c	29ed85bb-11eb-4458-bbf3-5a5644d167d6	e3f8aebb-1451-43a3-9e9b-582da01c8d08	2022-06-20 12:41:56.383576+05	2022-06-20 12:41:56.383576+05	\N
86e4b4d6-33e8-44a5-8aad-d9a53bbe9652	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	e3f8aebb-1451-43a3-9e9b-582da01c8d08	2022-06-20 12:41:56.396734+05	2022-06-20 12:41:56.396734+05	\N
bc954fab-1cfc-4a14-a2b2-62c8d6c68084	29ed85bb-11eb-4458-bbf3-5a5644d167d6	3b6d2d59-7ad4-4392-b7de-5a1d6bd003e1	2022-06-20 12:43:47.851536+05	2022-06-20 12:43:47.851536+05	\N
5fc165c1-018a-480c-b95d-8124df70d549	66772380-c161-4c45-9350-a45e765193e2	3b6d2d59-7ad4-4392-b7de-5a1d6bd003e1	2022-06-20 12:43:47.86481+05	2022-06-20 12:43:47.86481+05	\N
cf53f807-d2a1-4c0e-9191-0654e6ba8623	29ed85bb-11eb-4458-bbf3-5a5644d167d6	4a6b57e1-18d2-4aac-8346-576d3897967e	2022-06-20 12:47:57.011271+05	2022-06-20 12:47:57.011271+05	\N
1fdb393d-8db0-4cd7-ae92-08f74fc0b04f	66772380-c161-4c45-9350-a45e765193e2	4a6b57e1-18d2-4aac-8346-576d3897967e	2022-06-20 12:47:57.022952+05	2022-06-20 12:47:57.022952+05	\N
eeb60ea2-2cd5-4482-b806-e5a97511dcc2	29ed85bb-11eb-4458-bbf3-5a5644d167d6	9a0572df-3006-426e-a623-11c0cbc930ea	2022-06-20 12:49:20.557364+05	2022-06-20 12:49:20.557364+05	\N
3eb851df-c436-4d58-8cf7-9d27efbfae92	66772380-c161-4c45-9350-a45e765193e2	9a0572df-3006-426e-a623-11c0cbc930ea	2022-06-20 12:49:20.568728+05	2022-06-20 12:49:20.568728+05	\N
aae6c5c5-318c-4ee1-a3e1-9c07aa2be82e	29ed85bb-11eb-4458-bbf3-5a5644d167d6	1a8935fd-c6ab-4656-b173-826c487a2274	2022-06-21 10:15:31.648119+05	2022-06-21 10:15:31.648119+05	\N
4235a58d-a31b-4d92-96c3-e03b292b8917	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	1a8935fd-c6ab-4656-b173-826c487a2274	2022-06-21 10:15:31.681382+05	2022-06-21 10:15:31.681382+05	\N
a9181fa3-4fdc-4e26-a59a-9c07db3b22df	29ed85bb-11eb-4458-bbf3-5a5644d167d6	b4499063-096e-4fa6-9e21-a47185afd829	2022-06-21 10:17:07.728294+05	2022-06-21 10:17:07.728294+05	\N
d72e1b64-5daf-42da-a944-632f6929e59a	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	b4499063-096e-4fa6-9e21-a47185afd829	2022-06-21 10:17:07.739337+05	2022-06-21 10:17:07.739337+05	\N
a920c76c-50ad-4695-a342-8d379756ac79	29ed85bb-11eb-4458-bbf3-5a5644d167d6	538f0688-30ce-497b-9a0e-cd53d0d5239d	2022-06-21 10:21:35.566759+05	2022-06-21 10:21:35.566759+05	\N
317fa075-05a3-4f6f-b691-01b8e3035b54	45765130-7f97-4f0c-b886-f70b75e02610	538f0688-30ce-497b-9a0e-cd53d0d5239d	2022-06-21 10:21:35.579531+05	2022-06-21 10:21:35.579531+05	\N
ffb0f490-3ed7-44d7-b1b2-8447e0eacce7	29ed85bb-11eb-4458-bbf3-5a5644d167d6	0dc06a1f-e25a-4c3d-8310-09985e905262	2022-06-21 10:23:26.481104+05	2022-06-21 10:23:26.481104+05	\N
97f9ce4a-c0fe-41d4-abcd-6e3551893b71	45765130-7f97-4f0c-b886-f70b75e02610	0dc06a1f-e25a-4c3d-8310-09985e905262	2022-06-21 10:23:26.491473+05	2022-06-21 10:23:26.491473+05	\N
06d81ebd-9b32-411b-8165-f9de45a34d48	3d0851a3-45cf-4662-b315-67a98b9273c2	ec4963db-c429-4135-9790-d3860c350bc5	2022-06-21 10:28:00.506929+05	2022-06-21 10:28:00.506929+05	\N
22d06f76-ef62-47ce-bbd6-8bc0ccc00344	1b94e399-f0f1-435f-92c5-f855889c1683	ec4963db-c429-4135-9790-d3860c350bc5	2022-06-21 10:28:00.517138+05	2022-06-21 10:28:00.517138+05	\N
1cc6b8a6-698d-41db-ae8c-fc74dfd01f41	02bd4413-8586-49ab-802e-16304e756a8b	1fa25151-9c63-4554-a79d-faf6cc78ef69	2022-06-21 10:33:59.011702+05	2022-06-21 10:33:59.011702+05	\N
fee82dae-4c27-4e8f-86dd-cdbaf6243cdb	33cf0893-ff6e-40b3-b50f-2a3e926eca70	1fa25151-9c63-4554-a79d-faf6cc78ef69	2022-06-21 10:33:59.02302+05	2022-06-21 10:33:59.02302+05	\N
08a138a4-4f36-4397-b6f3-4dcc931de242	b982bd86-0a0f-4950-baad-5a131e9b728e	1fa25151-9c63-4554-a79d-faf6cc78ef69	2022-06-21 10:33:59.035017+05	2022-06-21 10:33:59.035017+05	\N
b6696fdd-7a08-4c87-a9b8-e209332c503b	02bd4413-8586-49ab-802e-16304e756a8b	d95aabd1-5a3a-47cc-aab5-9c6025e12280	2022-06-21 10:35:13.169478+05	2022-06-21 10:35:13.169478+05	\N
66bfc1f5-5b52-4722-acf0-cf46e1640e57	33cf0893-ff6e-40b3-b50f-2a3e926eca70	d95aabd1-5a3a-47cc-aab5-9c6025e12280	2022-06-21 10:35:13.181404+05	2022-06-21 10:35:13.181404+05	\N
089c20d9-b1cc-4da5-b066-c9fd63fedd01	b982bd86-0a0f-4950-baad-5a131e9b728e	d95aabd1-5a3a-47cc-aab5-9c6025e12280	2022-06-21 10:35:13.191772+05	2022-06-21 10:35:13.191772+05	\N
3dd089b3-5047-4fc6-9a8f-f625040abb31	02bd4413-8586-49ab-802e-16304e756a8b	d59506eb-aa84-4127-a411-5c5f95350d15	2022-06-21 10:37:39.683308+05	2022-06-21 10:37:39.683308+05	\N
88123309-697f-4af2-8a7a-8e0737cf98aa	33cf0893-ff6e-40b3-b50f-2a3e926eca70	d59506eb-aa84-4127-a411-5c5f95350d15	2022-06-21 10:37:39.693643+05	2022-06-21 10:37:39.693643+05	\N
7ea81d46-2460-413c-a85c-1931d909758f	b982bd86-0a0f-4950-baad-5a131e9b728e	d59506eb-aa84-4127-a411-5c5f95350d15	2022-06-21 10:37:39.704719+05	2022-06-21 10:37:39.704719+05	\N
cedf5aac-94d4-4a4b-9a8e-4ae79f0657a3	02bd4413-8586-49ab-802e-16304e756a8b	ab6ba3a4-0d3a-4510-acd0-feb4fe48fc19	2022-06-21 10:39:18.429355+05	2022-06-21 10:39:18.429355+05	\N
32f4f66c-d477-4f88-b295-0773a92d8668	b982bd86-0a0f-4950-baad-5a131e9b728e	ab6ba3a4-0d3a-4510-acd0-feb4fe48fc19	2022-06-21 10:39:18.441262+05	2022-06-21 10:39:18.441262+05	\N
8bb02e9a-60b1-459c-a5d9-1e6847a667a7	f745d171-68e6-42e2-b339-cb3c210cda55	ab6ba3a4-0d3a-4510-acd0-feb4fe48fc19	2022-06-21 10:39:18.451708+05	2022-06-21 10:39:18.451708+05	\N
a67bcc04-5d5b-49e2-b7e6-16721ade24f8	02bd4413-8586-49ab-802e-16304e756a8b	ce76ca4c-0ffb-4dd7-a252-3d3eaa6da732	2022-06-21 10:40:32.442941+05	2022-06-21 10:40:32.442941+05	\N
28fc71a8-c952-4d52-9556-87ab9f1eb8d7	b982bd86-0a0f-4950-baad-5a131e9b728e	ce76ca4c-0ffb-4dd7-a252-3d3eaa6da732	2022-06-21 10:40:32.452476+05	2022-06-21 10:40:32.452476+05	\N
b28d2d48-5617-4cbb-81e3-546703fed513	f745d171-68e6-42e2-b339-cb3c210cda55	ce76ca4c-0ffb-4dd7-a252-3d3eaa6da732	2022-06-21 10:40:32.463348+05	2022-06-21 10:40:32.463348+05	\N
fe115b96-a73f-4173-9bd6-11d1a94c7bc7	02bd4413-8586-49ab-802e-16304e756a8b	2072a0fb-bbc4-4231-a7a4-dad00bb0a892	2022-06-21 10:41:30.487167+05	2022-06-21 10:41:30.487167+05	\N
da5c41a7-b060-4ad1-a0eb-28689bd1f511	b982bd86-0a0f-4950-baad-5a131e9b728e	2072a0fb-bbc4-4231-a7a4-dad00bb0a892	2022-06-21 10:41:30.499603+05	2022-06-21 10:41:30.499603+05	\N
f236bca4-1ed8-4718-952d-6dc1e0ab6c50	f745d171-68e6-42e2-b339-cb3c210cda55	2072a0fb-bbc4-4231-a7a4-dad00bb0a892	2022-06-21 10:41:30.510054+05	2022-06-21 10:41:30.510054+05	\N
5fb6454d-dd53-4632-81eb-b01bd63fc6fe	02bd4413-8586-49ab-802e-16304e756a8b	86f78ca3-177d-4c89-8693-7678066d7389	2022-06-21 10:47:21.582616+05	2022-06-21 10:47:21.582616+05	\N
8b1085d7-38a2-4602-bdd8-3cc9e988b72c	5bb9a4e7-9992-418f-b551-537844d371da	86f78ca3-177d-4c89-8693-7678066d7389	2022-06-21 10:47:21.593583+05	2022-06-21 10:47:21.593583+05	\N
37232936-6786-48d0-82f7-ee9195dcdc2e	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	86f78ca3-177d-4c89-8693-7678066d7389	2022-06-21 10:47:21.605456+05	2022-06-21 10:47:21.605456+05	\N
828ef782-1826-44c0-be60-4371e6c98d0d	02bd4413-8586-49ab-802e-16304e756a8b	0a6863e2-7ed9-4fcd-9875-270fb778b33e	2022-06-21 10:48:09.349937+05	2022-06-21 10:48:09.349937+05	\N
fa8a3759-876c-407e-a633-59bb311b5b7f	5bb9a4e7-9992-418f-b551-537844d371da	0a6863e2-7ed9-4fcd-9875-270fb778b33e	2022-06-21 10:48:09.361646+05	2022-06-21 10:48:09.361646+05	\N
c038ad39-b797-436c-b60b-a4c38f70a2bb	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	0a6863e2-7ed9-4fcd-9875-270fb778b33e	2022-06-21 10:48:09.372502+05	2022-06-21 10:48:09.372502+05	\N
2eb9826e-9cb5-4cd0-8d10-ec9fd6d8f576	02bd4413-8586-49ab-802e-16304e756a8b	49381c4e-298d-43b7-8ae4-8dbe6e7da581	2022-06-21 10:49:08.707879+05	2022-06-21 10:49:08.707879+05	\N
eec559c6-45fa-4d53-a852-3f173e959e13	5bb9a4e7-9992-418f-b551-537844d371da	49381c4e-298d-43b7-8ae4-8dbe6e7da581	2022-06-21 10:49:08.719732+05	2022-06-21 10:49:08.719732+05	\N
e46d693e-b833-4a86-ad7f-e33c2061c6fb	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	49381c4e-298d-43b7-8ae4-8dbe6e7da581	2022-06-21 10:49:08.731181+05	2022-06-21 10:49:08.731181+05	\N
4d0a9d83-5dec-4702-8a7a-35613bb819e5	02bd4413-8586-49ab-802e-16304e756a8b	c1f8c7cb-081e-4f99-aeb3-0bc84153295e	2022-06-21 10:49:54.786032+05	2022-06-21 10:49:54.786032+05	\N
ecf29934-ed99-4313-be6d-c5f7fb8b5fe3	5bb9a4e7-9992-418f-b551-537844d371da	c1f8c7cb-081e-4f99-aeb3-0bc84153295e	2022-06-21 10:49:54.796938+05	2022-06-21 10:49:54.796938+05	\N
84823a12-9718-4c60-b430-d4cf0ceaa760	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	c1f8c7cb-081e-4f99-aeb3-0bc84153295e	2022-06-21 10:49:54.807387+05	2022-06-21 10:49:54.807387+05	\N
c97961e5-5aec-468a-9895-beed74d0ca7c	02bd4413-8586-49ab-802e-16304e756a8b	0cbe2487-c709-403f-a6c4-4f1a73fd3f78	2022-06-21 10:50:40.643457+05	2022-06-21 10:50:40.643457+05	\N
8800f712-2d38-4ff9-ae83-724c9d4e899b	02bd4413-8586-49ab-802e-16304e756a8b	cbb0047a-e543-41a8-845b-8439d11638f4	2022-06-21 10:54:40.613895+05	2022-06-21 10:54:40.613895+05	\N
0f6afd75-7c25-46bb-922a-625dc509c8ff	5bb9a4e7-9992-418f-b551-537844d371da	cbb0047a-e543-41a8-845b-8439d11638f4	2022-06-21 10:54:40.625302+05	2022-06-21 10:54:40.625302+05	\N
8e065fcf-6e52-4075-a85d-1a76d663d03e	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	cbb0047a-e543-41a8-845b-8439d11638f4	2022-06-21 10:54:40.635847+05	2022-06-21 10:54:40.635847+05	\N
66532212-2f1a-49a6-8522-6b487bdf552c	02bd4413-8586-49ab-802e-16304e756a8b	c5520db8-19de-4209-b99c-826a342210c3	2022-06-21 10:59:22.829853+05	2022-06-21 10:59:22.829853+05	\N
370b759a-eaa7-4126-a00b-ccb9daa53e80	5bb9a4e7-9992-418f-b551-537844d371da	c5520db8-19de-4209-b99c-826a342210c3	2022-06-21 10:59:22.84146+05	2022-06-21 10:59:22.84146+05	\N
6e800947-6611-44c1-a67f-30d895ec4f18	d4cb1359-6c23-4194-8e3c-21ed8cec8373	c5520db8-19de-4209-b99c-826a342210c3	2022-06-21 10:59:22.852847+05	2022-06-21 10:59:22.852847+05	\N
b80074c7-8009-4cea-b35e-c9fae0e4e7c5	02bd4413-8586-49ab-802e-16304e756a8b	ebc34352-64f7-4ad5-aa00-b1777efb3e56	2022-06-21 11:01:08.186259+05	2022-06-21 11:01:08.186259+05	\N
6063e64a-fc72-4d29-ac35-86d094a1dfe5	5bb9a4e7-9992-418f-b551-537844d371da	ebc34352-64f7-4ad5-aa00-b1777efb3e56	2022-06-21 11:01:08.197551+05	2022-06-21 11:01:08.197551+05	\N
679ba84d-7f82-4593-a842-2fa1c1f7fa2f	d4cb1359-6c23-4194-8e3c-21ed8cec8373	ebc34352-64f7-4ad5-aa00-b1777efb3e56	2022-06-21 11:01:08.209084+05	2022-06-21 11:01:08.209084+05	\N
66044d96-50be-4e7c-824c-c46a224ab9bf	3d0851a3-45cf-4662-b315-67a98b9273c2	88753f91-4e73-4478-91c5-37b278984294	2022-06-21 11:03:19.588643+05	2022-06-21 11:03:19.588643+05	\N
e643b147-ed19-4fb1-8a66-842bd7bc4523	f2e02bbb-c554-4989-b315-8d6aa0575bfa	88753f91-4e73-4478-91c5-37b278984294	2022-06-21 11:03:19.599137+05	2022-06-21 11:03:19.599137+05	\N
c6d52d3c-380f-4a9a-b774-59fd1dc45326	3d0851a3-45cf-4662-b315-67a98b9273c2	93096765-14be-4093-8e53-81caba6de3aa	2022-06-21 11:04:41.879148+05	2022-06-21 11:04:41.879148+05	\N
07f4772d-29ff-4106-974c-61079b8b50f2	f2e02bbb-c554-4989-b315-8d6aa0575bfa	93096765-14be-4093-8e53-81caba6de3aa	2022-06-21 11:04:41.889999+05	2022-06-21 11:04:41.889999+05	\N
57798dec-e841-4cc0-b4b4-e8fcbaab08dd	4af3388b-2738-4ff6-b42e-927cb0ff897f	070d7096-2fdd-4327-b0b6-99b13af1570f	2022-06-21 11:06:35.203416+05	2022-06-21 11:06:35.203416+05	\N
96eb2f34-6bd0-4d73-9628-5d40401d20eb	56b86071-1c45-490b-a683-a8898546f179	070d7096-2fdd-4327-b0b6-99b13af1570f	2022-06-21 11:06:35.214182+05	2022-06-21 11:06:35.214182+05	\N
f6945517-dacd-460c-a738-f3e1fd7bf6c9	ec0b10ac-bf81-4ae3-881e-ef616ea13d7f	070d7096-2fdd-4327-b0b6-99b13af1570f	2022-06-21 11:06:35.224811+05	2022-06-21 11:06:35.224811+05	\N
0c9daa52-fd4a-444b-bdbd-eccb56fc4b06	4af3388b-2738-4ff6-b42e-927cb0ff897f	aee7abe3-c6cc-4562-bf67-3f87e952611b	2022-06-21 11:07:41.781274+05	2022-06-21 11:07:41.781274+05	\N
21a919b5-2f0a-4704-8e89-c28f6cc2ba91	56b86071-1c45-490b-a683-a8898546f179	aee7abe3-c6cc-4562-bf67-3f87e952611b	2022-06-21 11:07:41.794174+05	2022-06-21 11:07:41.794174+05	\N
f2702c2c-5781-4fc0-91be-0ffbf4fb5f6f	ec0b10ac-bf81-4ae3-881e-ef616ea13d7f	aee7abe3-c6cc-4562-bf67-3f87e952611b	2022-06-21 11:07:41.804776+05	2022-06-21 11:07:41.804776+05	\N
9fb1a232-d25a-485c-8f88-2788fbf28fb8	4af3388b-2738-4ff6-b42e-927cb0ff897f	205b50c5-da4b-4edf-adac-54f93dc99253	2022-06-21 11:10:49.56247+05	2022-06-21 11:10:49.56247+05	\N
6aed71c9-8913-46d4-a3d8-fa233fd00d76	56b86071-1c45-490b-a683-a8898546f179	205b50c5-da4b-4edf-adac-54f93dc99253	2022-06-21 11:10:49.573395+05	2022-06-21 11:10:49.573395+05	\N
58b91db6-a126-4c89-a767-19959f409b52	0a1963a2-4084-403e-871d-763ae4825fab	205b50c5-da4b-4edf-adac-54f93dc99253	2022-06-21 11:10:49.584572+05	2022-06-21 11:10:49.584572+05	\N
0e050935-6f44-4b3d-83ec-877da06279b5	4af3388b-2738-4ff6-b42e-927cb0ff897f	c9307e74-88a2-4d96-96ec-6f04e42ad0cb	2022-06-21 11:11:48.019687+05	2022-06-21 11:11:48.019687+05	\N
b4c4308a-ba59-416f-8cc5-34e2c89de0e4	56b86071-1c45-490b-a683-a8898546f179	c9307e74-88a2-4d96-96ec-6f04e42ad0cb	2022-06-21 11:11:48.029964+05	2022-06-21 11:11:48.029964+05	\N
31b4d04c-112b-4cd3-b926-33e242ea92b1	0a1963a2-4084-403e-871d-763ae4825fab	c9307e74-88a2-4d96-96ec-6f04e42ad0cb	2022-06-21 11:11:48.041018+05	2022-06-21 11:11:48.041018+05	\N
08cd7a3c-2549-4059-9baf-968578611403	4af3388b-2738-4ff6-b42e-927cb0ff897f	f3208845-80d9-4ccb-9ad2-07a8ee2832c6	2022-06-21 11:13:34.487137+05	2022-06-21 11:13:34.487137+05	\N
f9cac9ef-f033-4d06-9fd1-8e56042a9de8	5d877898-9ef4-4b91-8518-193b431228a8	f3208845-80d9-4ccb-9ad2-07a8ee2832c6	2022-06-21 11:13:34.498485+05	2022-06-21 11:13:34.498485+05	\N
bc1aded9-87eb-4cbc-ac59-e73d98000ce8	fc87c4c5-d7cb-4def-a0e0-11cd5751e04b	f3208845-80d9-4ccb-9ad2-07a8ee2832c6	2022-06-21 11:13:34.509851+05	2022-06-21 11:13:34.509851+05	\N
813f8092-03a5-4d12-8c03-985dce382c44	4af3388b-2738-4ff6-b42e-927cb0ff897f	7bab1a39-0c66-4c1e-9f9c-7f25e050daa5	2022-06-21 11:14:39.177767+05	2022-06-21 11:14:39.177767+05	\N
381ecfff-c9d3-4a12-8bfe-1e1b3f40e921	5d877898-9ef4-4b91-8518-193b431228a8	7bab1a39-0c66-4c1e-9f9c-7f25e050daa5	2022-06-21 11:14:39.188324+05	2022-06-21 11:14:39.188324+05	\N
d684ca99-d6dc-46d5-ba75-92b378c652b3	4af3388b-2738-4ff6-b42e-927cb0ff897f	c182ee68-2717-4604-b0ab-0e6994e61ff0	2022-06-21 11:17:33.360363+05	2022-06-21 11:17:33.360363+05	\N
8ac7acf6-23cb-473f-b9bc-0f469247c687	5d877898-9ef4-4b91-8518-193b431228a8	c182ee68-2717-4604-b0ab-0e6994e61ff0	2022-06-21 11:17:33.371494+05	2022-06-21 11:17:33.371494+05	\N
0306adf4-1901-45c1-a22c-e3375913652f	38d92a87-4a9c-4860-94e6-e568f21ecd8e	c182ee68-2717-4604-b0ab-0e6994e61ff0	2022-06-21 11:17:33.382896+05	2022-06-21 11:17:33.382896+05	\N
d4526216-2cbe-4c57-bcf0-3c049d403c7f	4af3388b-2738-4ff6-b42e-927cb0ff897f	4ae4d83c-56ad-4d99-9d6f-e0dd77f9c982	2022-06-21 11:18:25.083615+05	2022-06-21 11:18:25.083615+05	\N
f79d9cc4-d15d-4b74-956b-e24f5b924354	5d877898-9ef4-4b91-8518-193b431228a8	4ae4d83c-56ad-4d99-9d6f-e0dd77f9c982	2022-06-21 11:18:25.095281+05	2022-06-21 11:18:25.095281+05	\N
3158a807-0160-467d-8ad9-ec364b81c225	38d92a87-4a9c-4860-94e6-e568f21ecd8e	4ae4d83c-56ad-4d99-9d6f-e0dd77f9c982	2022-06-21 11:18:25.105438+05	2022-06-21 11:18:25.105438+05	\N
\.


--
-- Data for Name: category_shop; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category_shop (id, category_id, shop_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: company_address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_address (id, lang_id, address, created_at, updated_at, deleted_at) FROM stdin;
d2c66808-e5fe-435f-ba01-cb717f80d9e0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Aşgabat şäheri Azady köçesiniň 23-nji jaýy	2022-06-22 18:44:50.21776+05	2022-06-22 18:44:50.21776+05	\N
75706251-06ea-41c1-905f-95ed8b4132f8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Улица Азади 23, Ашхабад	2022-06-22 18:44:50.239558+05	2022-06-22 18:44:50.239558+05	\N
\.


--
-- Data for Name: company_phone; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_phone (id, phone, created_at, updated_at, deleted_at) FROM stdin;
3060bc25-2a55-4ee0-894d-f87f887e1fc4	+99361899737	2022-06-22 18:21:06.98191+05	2022-06-22 18:21:06.98191+05	\N
\.


--
-- Data for Name: company_setting; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_setting (id, logo_path, favicon_path, email, instagram, created_at, updated_at, deleted_at) FROM stdin;
7d193677-e0b1-4df0-be88-dc6e16a47ca7	uploads/logo6e25c0ce-5f6f-45c0-9a52-1eab10edb892.jpg	uploads/favicona3dd02ae-759a-4ca4-8869-0c62ea6a14aa.jpg	isleg@gmail.com	@isleginstagram	2022-06-15 19:57:04.54457+05	2022-06-15 19:57:04.54457+05	\N
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (id, full_name, phone_number, password, birthday, gender, addresses, created_at, updated_at, deleted_at) FROM stdin;
20b26aa1-2247-4ed8-a6ef-3ec6bb9f64d7	Asyr Berdiyev	+99365453298	$2a$14$xpoVZw3GhVw05cx/iZmYu.0iqiRyMH46x58wcNuIWTxM/rQGjYSDu	1998-06-24	1	{"Mir 2/2 jay 7 oy 36","Mir 3/2 jay 5 oy 4"}	2022-07-12 12:43:16.090985+05	2022-07-12 12:43:16.090985+05	\N
e53e0ef8-a3a6-485f-8a5f-d0ba8327b3d5	Serdar Bayramow	+99365453294	$2a$14$EqECyFqszVXzcX5q4jwnqOe2ys8uTN.V.GwfIJkXq6ZjlbLOiXU2C	1998-09-03	1	{"Mir 4/2 jay 1 oy 5","Mir 1/2 jay 5 oy 4"}	2022-07-13 11:01:48.284329+05	2022-07-13 11:01:48.284329+05	\N
e6709e8f-0c1f-48b8-aeb5-c25f890f3c4e	Guljemal Bayramowa	+99363456742	$2a$14$xoBQ1kA5zR8COGnbVQ2mD.DOUmYYgPwRvXewMnJmCoXx7rQFetlc2	1998-09-03	0	{"Mir 4/2 jay 1 oy 5","Mir 1/2 jay 5 oy 4"}	2022-07-13 11:27:02.415977+05	2022-07-13 11:27:02.415977+05	\N
\.


--
-- Data for Name: district; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.district (id, price, created_at, updated_at, deleted_at) FROM stdin;
a58294d3-efe5-4cb7-82d3-8df8c37563c5	15	2022-06-25 10:23:25.640364+05	2022-06-25 10:23:25.640364+05	\N
\.


--
-- Data for Name: languages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.languages (id, name_short, flag_path, created_at, updated_at, deleted_at) FROM stdin;
8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	tm	uploads/language17b99bd1-f52d-41db-b4e6-1ecff03e0fd0.jpeg	2022-06-15 19:53:06.041686+05	2022-06-15 19:53:06.041686+05	\N
aea98b93-7bdf-455b-9ad4-a259d69dc76e	ru	uploads/language92b53cfe-d5a5-4686-9082-86fed42ffac1.jpeg	2022-06-15 19:53:21.29491+05	2022-06-15 19:53:21.29491+05	\N
\.


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.likes (id, product_id, customer_id, created_at, updated_at, deleted_at) FROM stdin;
4df3f9dd-8253-47ee-88bf-9da4934e4ae6	525af569-06b6-440a-ab5a-6ee0b39cf51d	20b26aa1-2247-4ed8-a6ef-3ec6bb9f64d7	2022-07-14 13:12:42.138482+05	2022-07-14 13:12:42.138482+05	\N
11bca126-ad3c-45d2-8b8b-75198a06747c	0d356bb1-c695-4b29-b199-bdff967abfe2	20b26aa1-2247-4ed8-a6ef-3ec6bb9f64d7	2022-07-15 11:10:50.064022+05	2022-07-15 11:10:50.064022+05	\N
36d05606-b393-42a8-a83f-d2aa2c77f403	0d356bb1-c695-4b29-b199-bdff967abfe2	e53e0ef8-a3a6-485f-8a5f-d0ba8327b3d5	2022-07-15 11:11:13.357965+05	2022-07-15 11:11:13.357965+05	\N
5677d729-d3bb-4e59-9acb-7046dd5df438	9a0572df-3006-426e-a623-11c0cbc930ea	e53e0ef8-a3a6-485f-8a5f-d0ba8327b3d5	2022-07-15 11:11:26.052559+05	2022-07-15 11:11:26.052559+05	\N
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, brend_id, price, old_price, amount, product_code, main_image_path, image_paths, created_at, updated_at, deleted_at) FROM stdin;
525af569-06b6-440a-ab5a-6ee0b39cf51d	c2f11fea-8057-4a7c-b290-c97d1d90ac59	46	48.6	23	we6dew6	uploads/productMain710f2792-d815-4606-9811-7bad5fb2c12e.jpeg	{uploads/product12b0eb77-8b58-4c01-944e-25ca0257dbd6.jpeg,uploads/product72235745-3db8-46bd-923b-b0a7294180dd.jpeg}	2022-06-20 12:37:45.259547+05	2022-06-20 12:37:45.259547+05	\N
0d356bb1-c695-4b29-b199-bdff967abfe2	c2f11fea-8057-4a7c-b290-c97d1d90ac59	23	22.5	45	s6fs6	uploads/productMain4dee470a-7301-4bfd-9d90-4bdf4fdfbccf.jpeg	\N	2022-06-20 12:40:33.161026+05	2022-06-20 12:40:33.161026+05	\N
e3f8aebb-1451-43a3-9e9b-582da01c8d08	6a31c50a-704f-4b0d-80ae-240ca3094cda	85	80.5	12	s6fs6	uploads/productMain32fed5d0-1542-4f92-86d4-d52beb535494.jpeg	{uploads/productac6abc8b-1f29-4c6a-9e6d-640a5c0612e4.jpg}	2022-06-20 12:41:56.290535+05	2022-06-20 12:41:56.290535+05	\N
3b6d2d59-7ad4-4392-b7de-5a1d6bd003e1	6a31c50a-704f-4b0d-80ae-240ca3094cda	24	23.5	128	s6fs66	uploads/productMaind1d6a321-1d10-4355-be39-7ff7001578c5.jpeg	{uploads/productc0fb523c-0ff7-45fc-b37b-18bfd82f0156.jpeg}	2022-06-20 12:43:47.742025+05	2022-06-20 12:43:47.742025+05	\N
4a6b57e1-18d2-4aac-8346-576d3897967e	c2f11fea-8057-4a7c-b290-c97d1d90ac59	65	63	84	s6fs66	uploads/productMain0f119bc5-9bbb-469d-a21d-923623bb9f71.jpg	{uploads/productc9dcea03-87e1-4463-8237-12cc6bbd751a.jpeg}	2022-06-20 12:47:56.91256+05	2022-06-20 12:47:56.91256+05	\N
9a0572df-3006-426e-a623-11c0cbc930ea	c2f11fea-8057-4a7c-b290-c97d1d90ac59	82	80.5	54	s6fs66	uploads/productMain07a3bcc6-b367-4c0f-abff-12c89ca9914c.jpeg	{uploads/productdeb7bb37-a63f-4d63-ab86-7fb512723f7a.jpeg,uploads/productcaafe599-c110-47fe-acc4-d47a3cac70db.jpeg}	2022-06-20 12:49:20.499041+05	2022-06-20 12:49:20.499041+05	\N
1a8935fd-c6ab-4656-b173-826c487a2274	214be879-65c3-4710-86b4-3fc3bce2e974	21	18.5	23	s6fs666we1	uploads/productMain7d6a6790-6d6e-4b8d-a0fb-f3ef822f64e8.jpeg	{uploads/product77cd7b93-8594-43e8-9eb8-880cd0f2283b.jpg,uploads/product854aae73-3803-48cc-a69b-4c9ed961e664.jpeg}	2022-06-21 10:15:31.506167+05	2022-06-21 10:15:31.506167+05	\N
b4499063-096e-4fa6-9e21-a47185afd829	214be879-65c3-4710-86b4-3fc3bce2e974	28	25.5	45	s6fs66	uploads/productMain07c5ce1d-b8b1-4893-9825-17d218832484.jpeg	{uploads/product4c118ed9-2986-484c-86c6-7540c11d5351.jpeg}	2022-06-21 10:17:07.683256+05	2022-06-21 10:17:07.683256+05	\N
538f0688-30ce-497b-9a0e-cd53d0d5239d	214be879-65c3-4710-86b4-3fc3bce2e974	23.5	25	45	s6fs66	uploads/productMain7377551a-91b3-4996-a5ae-f03ddf5530ac.jpeg	{uploads/product8b43004a-ad22-4241-8ad5-bd12e644351e.jpeg}	2022-06-21 10:21:35.476766+05	2022-06-21 10:21:35.476766+05	\N
0dc06a1f-e25a-4c3d-8310-09985e905262	214be879-65c3-4710-86b4-3fc3bce2e974	46	0	58	s6fs666516	uploads/productMainb95cac81-2dc7-4386-91ce-476b2f5763e4.jpeg	{uploads/product50e40976-8376-4028-a389-2557814cfa48.jpeg}	2022-06-21 10:23:26.430614+05	2022-06-21 10:23:26.430614+05	\N
ec4963db-c429-4135-9790-d3860c350bc5	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	34	35	89	6s5f6	uploads/productMainca02c36d-c390-41bf-b380-bd58e1879939.jpg	{uploads/producta0f75297-de8c-48ca-b923-a996926150ec.jpeg,uploads/producte977c95c-343e-41f8-9ee1-ec7c21c7ac35.jpeg}	2022-06-21 10:28:00.458614+05	2022-06-21 10:28:00.458614+05	\N
1fa25151-9c63-4554-a79d-faf6cc78ef69	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	74.3	0	456	w5we	uploads/productMain08f71540-c011-4251-a7eb-4c8dccc28c08.jpeg	{uploads/productcd6cb444-b9d1-41cf-ad2d-272c80384c80.jpeg}	2022-06-21 10:33:58.914208+05	2022-06-21 10:33:58.914208+05	\N
d95aabd1-5a3a-47cc-aab5-9c6025e12280	c2f11fea-8057-4a7c-b290-c97d1d90ac59	38.5	0	65	6ef51e65	uploads/productMain32404a31-0a04-4f10-9888-b75750698b60.jpeg	{uploads/product376a5f90-a152-48d5-b5eb-ebff1e7ebf87.jpeg,uploads/productd06b75f1-80de-4ce4-b474-192884784290.jpeg}	2022-06-21 10:35:13.123633+05	2022-06-21 10:35:13.123633+05	\N
d59506eb-aa84-4127-a411-5c5f95350d15	c2f11fea-8057-4a7c-b290-c97d1d90ac59	104.2	0	568	we6d1we6	uploads/productMain7c35582a-6ea3-4a9f-a86c-c5fed6da87f8.jpeg	{uploads/product32ded633-36d9-4999-957b-b6904f80775e.jpeg,uploads/productdd77dcc8-de52-4ca8-9582-461216ae4d00.jpeg}	2022-06-21 10:37:39.623946+05	2022-06-21 10:37:39.623946+05	\N
ab6ba3a4-0d3a-4510-acd0-feb4fe48fc19	c2f11fea-8057-4a7c-b290-c97d1d90ac59	3.5	0	61	w16we	uploads/productMain20e75fae-da0b-4149-b7b5-0e34a99f4522.jpeg	{uploads/productdf63152c-5b0f-4bc3-a35e-dea9dc0f12aa.jpeg,uploads/product770d0c7f-914a-4b00-87ea-37bc84df2ccc.jpeg}	2022-06-21 10:39:18.368227+05	2022-06-21 10:39:18.368227+05	\N
ce76ca4c-0ffb-4dd7-a252-3d3eaa6da732	6a31c50a-704f-4b0d-80ae-240ca3094cda	1.5	1.7	2684	w6dwed	uploads/productMain1ed63766-b427-4f3b-b796-a593d2a5397e.jpeg	{uploads/product104c0eee-f261-4ddc-8f05-6821a08675c5.jpeg,uploads/product213eb583-2b3c-4f2e-a6f7-1454e080ce74.jpeg}	2022-06-21 10:40:32.397262+05	2022-06-21 10:40:32.397262+05	\N
2072a0fb-bbc4-4231-a7a4-dad00bb0a892	6a31c50a-704f-4b0d-80ae-240ca3094cda	7	0	264	1w6dew	uploads/productMain6ee2f9be-6e64-4aed-8631-ef4b81aa1701.jpeg	{uploads/product21c93b2f-1f1b-4fdc-a33c-b1e96aa3a119.jpeg,uploads/producteeaa8cfe-2f98-4ef7-b9cc-994991cea1d0.jpeg}	2022-06-21 10:41:30.430549+05	2022-06-21 10:41:30.430549+05	\N
86f78ca3-177d-4c89-8693-7678066d7389	6a31c50a-704f-4b0d-80ae-240ca3094cda	18.6	0	56	618ew	uploads/productMainb54c5441-9a0a-41bb-9c3a-64034fb3912d.jpeg	{uploads/product9e5080f3-d3bc-498d-b0fe-5621c1f4e0ac.jpeg,uploads/product4113739e-2e42-4f76-a46c-5d965891be60.jpeg}	2022-06-21 10:47:21.505128+05	2022-06-21 10:47:21.505128+05	\N
0a6863e2-7ed9-4fcd-9875-270fb778b33e	6a31c50a-704f-4b0d-80ae-240ca3094cda	17.7	0	85	61d9we8	uploads/productMain2996089c-4c86-48d2-815c-10395b92a9cb.jpeg	{uploads/productfe827c62-1954-47b3-8416-f78817dbdcc0.jpeg,uploads/product1e9643be-36b3-47f8-82d0-9ad23a940b98.jpeg}	2022-06-21 10:48:09.289024+05	2022-06-21 10:48:09.289024+05	\N
49381c4e-298d-43b7-8ae4-8dbe6e7da581	214be879-65c3-4710-86b4-3fc3bce2e974	13.9	0	89	6wd4we98	uploads/productMain75ac8d63-7d76-4758-8a60-375bcf30ff7f.jpeg	{uploads/producte48f5ec1-8d0f-4781-b1f8-50a52fc3e70a.jpeg,uploads/product1729d3ed-9e53-409a-aa99-06fa37e1c1f4.jpeg}	2022-06-21 10:49:08.655229+05	2022-06-21 10:49:08.655229+05	\N
c1f8c7cb-081e-4f99-aeb3-0bc84153295e	214be879-65c3-4710-86b4-3fc3bce2e974	22.7	0	268	w6dw9	uploads/productMain93cf7211-66b8-47d1-aa7f-1580fe8137fc.jpeg	{uploads/product6f619bc9-936e-453d-a339-37ab850f64f9.jpeg,uploads/product65530cd8-c8c2-4ce2-8c37-7327a6aaec18.jpeg}	2022-06-21 10:49:54.728976+05	2022-06-21 10:49:54.728976+05	\N
0cbe2487-c709-403f-a6c4-4f1a73fd3f78	214be879-65c3-4710-86b4-3fc3bce2e974	13.9	0	68	ww6	uploads/productMain389701f5-89f1-4996-9a1d-e327c78e2536.jpeg	{uploads/productc0c9898e-6404-4586-a786-0d93d2816957.jpeg,uploads/productf577b09a-d1e1-4ac3-90fa-183ecfc5840c.jpeg}	2022-06-21 10:50:40.591888+05	2022-06-21 10:50:40.591888+05	\N
cbb0047a-e543-41a8-845b-8439d11638f4	214be879-65c3-4710-86b4-3fc3bce2e974	13.9	0	68	ww6	uploads/productMain0174f727-509f-4b7c-b217-4f913275510b.jpeg	{uploads/product974ab765-cdf9-4e93-9ff8-c8caae136b2d.jpeg,uploads/producta8a2abc9-fe81-44c9-a1d5-3476cec59b3b.jpeg}	2022-06-21 10:54:40.567153+05	2022-06-21 10:54:40.567153+05	\N
c5520db8-19de-4209-b99c-826a342210c3	c2f11fea-8057-4a7c-b290-c97d1d90ac59	19	0	68	16dew1	uploads/productMain8af87cd6-8fcf-4a79-bcce-c8de075aa44b.jpeg	{uploads/productac29a46d-fd24-4a3e-b4db-d9ab9b9ccfbc.jpeg,uploads/product369d6008-82a7-4f1b-878e-59f0360b00b6.jpeg}	2022-06-21 10:59:22.772281+05	2022-06-21 10:59:22.772281+05	\N
ebc34352-64f7-4ad5-aa00-b1777efb3e56	c2f11fea-8057-4a7c-b290-c97d1d90ac59	13	0	688	w1e6w6	uploads/productMain60aa5e71-b7ee-4ef6-9d64-b59819d2df51.jpeg	{uploads/product2e062651-7f86-43fe-8d0d-a669cf3cbc9c.jpeg,uploads/product84c04d6e-fba4-406a-8640-8419d7357976.jpeg}	2022-06-21 11:01:08.142166+05	2022-06-21 11:01:08.142166+05	\N
88753f91-4e73-4478-91c5-37b278984294	c2f11fea-8057-4a7c-b290-c97d1d90ac59	7.1	0	65	w1ef6we4	uploads/productMaind55e28bc-ef9f-43df-8aef-84cb81d36cbc.jpeg	{uploads/product459c0523-e4c4-4f00-9a97-8c84e3ba16c9.jpeg,uploads/productec026133-c23b-4ef8-928e-edb4fe22db57.jpeg}	2022-06-21 11:03:19.525658+05	2022-06-21 11:03:19.525658+05	\N
93096765-14be-4093-8e53-81caba6de3aa	6a31c50a-704f-4b0d-80ae-240ca3094cda	4.5	0	85	1yku4i9k84	uploads/productMain8b11acad-8d6a-4530-a484-4b61c1cec60f.jpeg	{uploads/productd04e2ca7-207d-42fd-b48c-c40bf02ffecc.jpeg,uploads/productb56d9e45-4ce6-4d70-9021-00a00710fcf4.jpeg}	2022-06-21 11:04:41.828533+05	2022-06-21 11:04:41.828533+05	\N
070d7096-2fdd-4327-b0b6-99b13af1570f	6a31c50a-704f-4b0d-80ae-240ca3094cda	94.8	98.6	92	ehrbterbi	uploads/productMain2037e05e-4c71-497d-847c-1736795da2c1.jpeg	{uploads/product419cf127-92d6-4f0d-aa59-440c4cecdb89.jpeg,uploads/productb5813ce9-4c4a-4145-9268-855478b7a9e3.jpeg}	2022-06-21 11:06:35.148892+05	2022-06-21 11:06:35.148892+05	\N
aee7abe3-c6cc-4562-bf67-3f87e952611b	214be879-65c3-4710-86b4-3fc3bce2e974	56.1	58.1	34	odewfiobe	uploads/productMain6e6f4f07-c035-4449-b1ae-6671189e35f7.jpeg	{uploads/product97dc2e51-919f-41b5-8295-55d23c980dd4.jpeg,uploads/productd2745857-1f33-45e4-acdf-633c0437c87b.jpeg}	2022-06-21 11:07:41.736902+05	2022-06-21 11:07:41.736902+05	\N
205b50c5-da4b-4edf-adac-54f93dc99253	214be879-65c3-4710-86b4-3fc3bce2e974	18.3	18.8	64	f6h51yuj7	uploads/productMain2b4d0e4b-6a4a-45e5-a9a8-bcd4c15f1e17.jpeg	{uploads/producta7c24eb9-99e2-4cc3-96b7-7bbaa61d2d86.jpeg,uploads/product9e85132c-cb21-4051-85e3-4c5e3b1931c4.jpeg}	2022-06-21 11:10:49.511294+05	2022-06-21 11:10:49.511294+05	\N
c9307e74-88a2-4d96-96ec-6f04e42ad0cb	214be879-65c3-4710-86b4-3fc3bce2e974	40.3	0	95	656greg	uploads/productMainadf07ba4-eed4-4897-bbc9-ed78eb733877.jpeg	{uploads/productbcd14c08-5329-498b-9364-75a2333e6797.jpeg,uploads/product71217e34-1548-4414-8502-e5885257a323.jpeg}	2022-06-21 11:11:47.929311+05	2022-06-21 11:11:47.929311+05	\N
f3208845-80d9-4ccb-9ad2-07a8ee2832c6	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	13.9	15.4	485	14sfsffew654	uploads/productMain030a48cc-5c4d-4be8-b803-6107234be675.jpeg	{uploads/product8156e8d7-af24-4ea3-8135-19c2b661f1d5.jpeg,uploads/productccbe9505-f0e0-4eda-a23b-6c8538e1f54c.jpeg}	2022-06-21 11:13:34.425672+05	2022-06-21 11:13:34.425672+05	\N
7bab1a39-0c66-4c1e-9f9c-7f25e050daa5	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	11.5	15.8	72	oeiwfwoefo	uploads/productMain29caf1dd-47cd-470f-a27a-775fe6d1f3ad.jpeg	{uploads/product2b9ca336-235f-4fa5-a00c-cf141503d86d.jpeg,uploads/productc6d63afb-9cd3-4e04-a3f1-b7bc7634ecbe.jpeg}	2022-06-21 11:14:39.128719+05	2022-06-21 11:14:39.128719+05	\N
c182ee68-2717-4604-b0ab-0e6994e61ff0	6a31c50a-704f-4b0d-80ae-240ca3094cda	31.7	0	72	6ef987e8	uploads/productMaind1c1e62c-af97-4983-92ca-d27044d9b94e.jpeg	{uploads/product14bb0731-ba19-483d-9bf7-e59230364b63.jpeg,uploads/product951f3c30-e33b-4c8f-8e37-273e303edac9.jpeg}	2022-06-21 11:17:33.304681+05	2022-06-21 11:17:33.304681+05	\N
4ae4d83c-56ad-4d99-9d6f-e0dd77f9c982	6a31c50a-704f-4b0d-80ae-240ca3094cda	35	0	19	j6yukuy	uploads/productMain2ea4ba35-5068-4b76-bd96-41f87ef6ef6d.jpeg	{uploads/product3f3d03da-31a1-4da5-83c2-91ec4f262b31.jpeg,uploads/product9a7162dd-1616-404d-b5c5-2833d7e8febc.jpeg}	2022-06-21 11:18:25.021332+05	2022-06-21 11:18:25.021332+05	\N
\.


--
-- Data for Name: shops; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shops (id, owner_name, address, phone_number, number_of_goods, running_time, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: translation_about; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_about (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
7abeb5cf-2fbb-43b9-94ca-251dd5f40d5a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Sizi Isleg onlaýn marketimizde hoş gördük!	Onlaýn marketimiz 2019-njy ýylyň iýul aýyndan bäri hyzmat berýär. Häzirki wagtda Size ýüzlerçe brendlere degişli bolan müňlerçe haryt görnüşlerini hödürleýäris! Haryt görnüşlerimizi sizden gelýän isleg we teklipleriň esasynda köpeltmäge dowam edýäris. Biziň maksadymyz müşderilerimize ýokary hilli hyzmat bermek bolup durýar. Indi Siz öýüňizden çykmazdan özüňizi gerekli zatlar bilen üpjün edip bilersiňiz! Munuň bilen bir hatarda Siz wagtyňyzy we transport çykdajylaryny hem tygşytlaýarsyňyz. Tölegi harytlar size gowuşandan soňra nagt ýa-da bank kartlarynyň üsti bilen amala aşyryp bilersiňiz!\n\nBiziň gapymyz hyzmatdaşlyklara we tekliplere hemişe açyk!	2022-06-25 12:07:15.62033+05	2022-06-25 12:07:15.62033+05	\N
e50bb3d1-14a1-400e-83d9-8bc15969b914	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Рады приветствовать Вас в интернет-маркете Isleg!	Мы начали работу в июле 2019 года и на сегодняшний день мы предлагаем Вам тысячи видов товаров, которые принадлежат сотням брендам. Каждый день мы работаем над увеличением ассортимента, привлечением новых компаний к сотрудничеству. Целью нашей работы является создание выгодных условий для наших клиентов-экономия времени на походы в магазины, оплата наличными или картой, доставка в удобное время, и конечно же качественная продукция по лучшим ценам!\n\nМы открыты для сотрудничества и пожеланий!	2022-06-25 12:07:15.653744+05	2022-06-25 12:07:15.653744+05	\N
\.


--
-- Data for Name: translation_afisa; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_afisa (id, afisa_id, lang_id, title, description, created_at, updated_at, deleted_at) FROM stdin;
5b00d76a-a295-4ded-b0ed-a8e29d6ea113	cb670531-8ca3-4d74-8ea4-f7853aae4132	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Aksiya		2022-06-23 18:04:26.357443+05	2022-06-23 18:04:26.357443+05	\N
ff6ce1a2-0cdf-440f-9175-35bda6750e42	cb670531-8ca3-4d74-8ea4-f7853aae4132	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Действие		2022-06-23 18:04:26.419551+05	2022-06-23 18:04:26.419551+05	\N
\.


--
-- Data for Name: translation_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_category (id, lang_id, category_id, name, created_at, updated_at, deleted_at) FROM stdin;
97c675b2-a2a5-459d-b785-27bd9b65f976	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	3d0851a3-45cf-4662-b315-67a98b9273c2	Gök we bakja önümler	2022-06-16 12:59:15.640472+05	2022-06-16 12:59:15.640472+05	\N
33525d63-123e-45e0-bbab-34b09ebf22a2	aea98b93-7bdf-455b-9ad4-a259d69dc76e	3d0851a3-45cf-4662-b315-67a98b9273c2	Фрукты и овощи	2022-06-16 12:59:15.65105+05	2022-06-16 12:59:15.65105+05	\N
f66de70c-bfa5-458e-921d-94810fa573cd	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	1b94e399-f0f1-435f-92c5-f855889c1683	Miweler	2022-06-16 13:32:50.889423+05	2022-06-16 13:32:50.889423+05	\N
e96a4f88-818a-4285-ab39-d0f4dacc1115	aea98b93-7bdf-455b-9ad4-a259d69dc76e	1b94e399-f0f1-435f-92c5-f855889c1683	Фрукты	2022-06-16 13:32:50.90619+05	2022-06-16 13:32:50.90619+05	\N
bff34c21-04c1-4cea-bfaf-c8f9ce7e2bfe	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	02bd4413-8586-49ab-802e-16304e756a8b	Iýmit önümleri	2022-06-16 13:43:22.674562+05	2022-06-16 13:43:22.674562+05	\N
0e400414-a80c-449d-8842-dd6667b45c73	aea98b93-7bdf-455b-9ad4-a259d69dc76e	02bd4413-8586-49ab-802e-16304e756a8b	Продовольственная продукция	2022-06-16 13:43:22.681932+05	2022-06-16 13:43:22.681932+05	\N
4eef5d40-9aad-4101-b36b-9026dd3dfb51	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b982bd86-0a0f-4950-baad-5a131e9b728e	Kofe we Kakao	2022-06-16 13:44:16.499713+05	2022-06-16 13:44:16.499713+05	\N
10a8b5ec-a3ca-448d-975b-83b3a7a8c0d2	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b982bd86-0a0f-4950-baad-5a131e9b728e	Кофе и Какао	2022-06-16 13:44:16.515874+05	2022-06-16 13:44:16.515874+05	\N
d8a96324-9b81-4b09-914f-f77a0915e35b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	33cf0893-ff6e-40b3-b50f-2a3e926eca70	Tebigy ereýän kofeler	2022-06-16 13:44:51.233713+05	2022-06-16 13:44:51.233713+05	\N
b88bb26f-942f-4638-b189-02bad933b730	aea98b93-7bdf-455b-9ad4-a259d69dc76e	33cf0893-ff6e-40b3-b50f-2a3e926eca70	Натуральный растворимый Кофе	2022-06-16 13:44:51.250133+05	2022-06-16 13:44:51.250133+05	\N
21520180-13e2-4c2b-a5f9-866c2e59ba87	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f745d171-68e6-42e2-b339-cb3c210cda55	Kiçi paket kofeler	2022-06-16 13:45:48.889727+05	2022-06-16 13:45:48.889727+05	\N
ab35a97a-dfd1-4100-8e84-d34e74e9a02e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f745d171-68e6-42e2-b339-cb3c210cda55	Кофе в пакетиках	2022-06-16 13:45:48.906024+05	2022-06-16 13:45:48.906024+05	\N
e099e7f6-1b97-4f70-8f29-f586ab6697d0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	5bb9a4e7-9992-418f-b551-537844d371da	Şokolad we Keksler	2022-06-16 13:46:44.657849+05	2022-06-16 13:46:44.657849+05	\N
415a0711-2482-44b3-8f03-923dca28bd5d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	5bb9a4e7-9992-418f-b551-537844d371da	Шоколады и Кексы	2022-06-16 13:46:44.673892+05	2022-06-16 13:46:44.673892+05	\N
4eb6bcbf-91f2-4505-a27e-cc3f96f2b829	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	Plitkalar	2022-06-16 13:47:18.888998+05	2022-06-16 13:47:18.888998+05	\N
53fb44c7-45fb-49f0-a433-aaed23b2dfc0	aea98b93-7bdf-455b-9ad4-a259d69dc76e	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	Плитки	2022-06-16 13:47:18.942159+05	2022-06-16 13:47:18.942159+05	\N
ee2f97fb-8c6c-4e38-bdb3-bf769bc95d3b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d4cb1359-6c23-4194-8e3c-21ed8cec8373	Batonçikler	2022-06-16 13:48:04.581888+05	2022-06-16 13:48:04.581888+05	\N
ea104eaf-c3fd-4f2d-88bf-dffc14d48dc5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d4cb1359-6c23-4194-8e3c-21ed8cec8373	Батончики	2022-06-16 13:48:04.597499+05	2022-06-16 13:48:04.597499+05	\N
0e6d1662-02bf-49e7-913f-0b3ff19102e8	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f2e02bbb-c554-4989-b315-8d6aa0575bfa	Gök önümler	2022-06-16 13:48:59.304183+05	2022-06-16 13:48:59.304183+05	\N
7bde654c-6ae2-4de8-be40-7379a84e66ea	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f2e02bbb-c554-4989-b315-8d6aa0575bfa	Овощи	2022-06-16 13:48:59.365879+05	2022-06-16 13:48:59.365879+05	\N
5cfeab53-7e44-4001-8310-ddbf1779d4c6	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	4af3388b-2738-4ff6-b42e-927cb0ff897f	Arassaçylyk we Hojalyk	2022-06-16 13:49:43.658705+05	2022-06-16 13:49:43.658705+05	\N
2b31f071-4d17-49a6-96b1-ca7bf2121083	aea98b93-7bdf-455b-9ad4-a259d69dc76e	4af3388b-2738-4ff6-b42e-927cb0ff897f	Уборка и Дом	2022-06-16 13:49:43.665449+05	2022-06-16 13:49:43.665449+05	\N
9d79f031-a5d1-4827-8547-ff44e8ee9ec7	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	56b86071-1c45-490b-a683-a8898546f179	Kir ýuwujy serişdeler	2022-06-16 13:50:35.361634+05	2022-06-16 13:50:35.361634+05	\N
c4dc2d27-8966-460d-9992-f9fcf2ca6c0c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	56b86071-1c45-490b-a683-a8898546f179	Моющие cредства	2022-06-16 13:50:35.378397+05	2022-06-16 13:50:35.378397+05	\N
edf0966a-559e-4b49-a3f7-9ad28f5d26cb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ec0b10ac-bf81-4ae3-881e-ef616ea13d7f	Awtomat üçin	2022-06-16 13:51:05.861192+05	2022-06-16 13:51:05.861192+05	\N
d0b77bd5-7bcf-4f05-a04d-299eedaba57d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ec0b10ac-bf81-4ae3-881e-ef616ea13d7f	Для автоматической стирки	2022-06-16 13:51:05.877721+05	2022-06-16 13:51:05.877721+05	\N
5098490c-af17-49ce-8fb4-2742c152b25d	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0a1963a2-4084-403e-871d-763ae4825fab	Elde ýuwmak üçin	2022-06-16 13:51:56.030185+05	2022-06-16 13:51:56.030185+05	\N
266b5dfd-e894-48d0-9b28-8a905f631cc2	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0a1963a2-4084-403e-871d-763ae4825fab	Для ручной стирки	2022-06-16 13:51:56.046414+05	2022-06-16 13:51:56.046414+05	\N
250873f0-eb5e-4484-8e83-421c83f571a2	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	5d877898-9ef4-4b91-8518-193b431228a8	Sabynlar	2022-06-16 13:52:45.186062+05	2022-06-16 13:52:45.186062+05	\N
c670fa26-0df8-4eaf-937a-30e1eab17846	aea98b93-7bdf-455b-9ad4-a259d69dc76e	5d877898-9ef4-4b91-8518-193b431228a8	Мыла	2022-06-16 13:52:45.202046+05	2022-06-16 13:52:45.202046+05	\N
2f99f616-59dd-4499-8ad2-6efe92a2928a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	fc87c4c5-d7cb-4def-a0e0-11cd5751e04b	Adaty Sabynlar	2022-06-16 13:53:34.342238+05	2022-06-16 13:53:34.342238+05	\N
49752597-6eb4-44c3-8714-92ad883fca65	aea98b93-7bdf-455b-9ad4-a259d69dc76e	fc87c4c5-d7cb-4def-a0e0-11cd5751e04b	Обычные мыла	2022-06-16 13:53:34.359039+05	2022-06-16 13:53:34.359039+05	\N
4c11f4ea-46f4-43e8-9359-7c34118109bd	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	38d92a87-4a9c-4860-94e6-e568f21ecd8e	Suwuk Sabynlar	2022-06-16 13:54:05.397857+05	2022-06-16 13:54:05.397857+05	\N
093e32f6-c04e-42a2-b446-413982903718	aea98b93-7bdf-455b-9ad4-a259d69dc76e	38d92a87-4a9c-4860-94e6-e568f21ecd8e	Жидкие мыла	2022-06-16 13:54:05.41627+05	2022-06-16 13:54:05.41627+05	\N
85469cf2-f48a-4e73-800d-ebf599aaeaba	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	29ed85bb-11eb-4458-bbf3-5a5644d167d6	Arzanladyş we Aksiýalar	2022-06-20 09:41:17.756928+05	2022-06-20 09:41:17.756928+05	\N
bbdd06a4-2dce-4c99-bf05-cf4e911776c7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	29ed85bb-11eb-4458-bbf3-5a5644d167d6	Распродажи и Акции	2022-06-20 09:41:17.941489+05	2022-06-20 09:41:17.941489+05	\N
8a91bcb0-fcce-4a4f-80ff-a2896c0cc36a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	Arzanladyşdaky harytlar	2022-06-20 09:43:07.368782+05	2022-06-20 09:43:07.368782+05	\N
ce573dfd-6af8-4e64-8260-8746a090acd7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	Продукция со скидкой	2022-06-20 09:43:07.377729+05	2022-06-20 09:43:07.377729+05	\N
34f4cdb5-04b9-48c0-b5b0-0045a02aa094	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	66772380-c161-4c45-9350-a45e765193e2	Aksiýadaky harytlar	2022-06-20 09:45:34.450534+05	2022-06-20 09:45:34.450534+05	\N
713cc05f-6a9d-4dae-88b5-dde2e564480c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	66772380-c161-4c45-9350-a45e765193e2	Продукция в категории Акции	2022-06-20 09:45:34.466904+05	2022-06-20 09:45:34.466904+05	\N
e224ecfc-6daa-4df5-8112-74846fc44867	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	Sowgatlyk toplumlar	2022-06-20 09:46:01.148565+05	2022-06-20 09:46:01.148565+05	\N
53959762-0b63-4100-ae13-4bbf8c015fec	aea98b93-7bdf-455b-9ad4-a259d69dc76e	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	Подарочные наборы	2022-06-20 09:46:01.408239+05	2022-06-20 09:46:01.408239+05	\N
3b756a33-bf2c-4d04-af57-962a3226d00b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	45765130-7f97-4f0c-b886-f70b75e02610	Täze harytlar	2022-06-20 10:11:06.719528+05	2022-06-20 10:11:06.719528+05	\N
2d22961c-ef08-4238-ae54-c00593c0073c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	45765130-7f97-4f0c-b886-f70b75e02610	Новые продукты	2022-06-20 10:11:06.735056+05	2022-06-20 10:11:06.735056+05	\N
\.


--
-- Data for Name: translation_contact; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_contact (id, lang_id, full_name, email, phone, letter, company_phone, imo, company_email, instagram, created_at, updated_at, deleted_at, button_text) FROM stdin;
73253999-7355-42b4-8700-94de76f0058a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	at_tm	rmail_tm	phone_tm	letter_tm	cp tm	imo tm	ce tm	instagram tm	2022-06-27 11:29:47.914891+05	2022-06-27 11:29:47.914891+05	\N	ugrat
f1693167-0c68-4a54-9831-56f124d629a3	aea98b93-7bdf-455b-9ad4-a259d69dc76e	at_ru	mail_ru	phone_ru	letter ru	cp ru	imo ru	ce ru	instagram ru	2022-06-27 11:29:48.050553+05	2022-06-27 11:29:48.050553+05	\N	Отправить
\.


--
-- Data for Name: translation_district; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_district (id, lang_id, district_id, name, created_at, updated_at, deleted_at) FROM stdin;
ad9f94d3-05e7-43b3-aa77-7b7f3754d003	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	a58294d3-efe5-4cb7-82d3-8df8c37563c5	Parahat 2	2022-06-25 10:23:25.712337+05	2022-06-25 10:23:25.712337+05	\N
aa1cfa48-3132-4dd4-abfb-070a2986690b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	a58294d3-efe5-4cb7-82d3-8df8c37563c5	Mir 2	2022-06-25 10:23:25.774504+05	2022-06-25 10:23:25.774504+05	\N
\.


--
-- Data for Name: translation_footer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_footer (id, lang_id, about, payment, contact, secure, word, created_at, updated_at, deleted_at) FROM stdin;
12dc4c16-5712-4bff-a957-8e16d450b4fb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Biz Barada	Eltip bermek we töleg tertibi	Aragatnaşyk	Ulanyş düzgünleri we gizlinlik şertnamasy	Ähli hukuklary goraglydyr	2022-06-22 15:23:32.716064+05	2022-06-22 15:23:32.716064+05	\N
84b5504f-1056-4b44-94dd-a7819148da66	aea98b93-7bdf-455b-9ad4-a259d69dc76e	О нас	Порядок доставки и оплаты	Коммуникация	Обслуживания и Политика Конфиденциальности	Все права защищены	2022-06-22 15:23:32.793161+05	2022-06-22 15:23:32.793161+05	\N
\.


--
-- Data for Name: translation_header; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_header (id, lang_id, research, phone, password, forgot_password, sign_in, sign_up, name, password_verification, verify_secure, my_information, my_favorites, my_orders, log_out, created_at, updated_at, deleted_at) FROM stdin;
eaf206e6-d515-4bdb-9323-a047cd0edae5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	gözleg	telefon	parol	Acar sozumi unutdym	ulgama girmek	agza bolmak	Ady	Acar sozi tassyklamak	Ulanyş Düzgünlerini we Gizlinlik Şertnamasyny okadym we kabul edýärin	maglumatym	halanlarym	sargytlarym	cykmak	2022-06-16 04:48:26.460534+05	2022-06-16 04:48:26.460534+05	\N
9154e800-2a92-47de-b4ff-1e63b213e5f7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	поиск	tелефон	пароль	забыл пароль	войти	зарегистрироваться	имя	Подтвердить Пароль	Я прочитал и принимаю Условия Обслуживания и Политика Конфиденциальности	моя информация	мои любимые	мои заказы	выйти	2022-06-16 04:48:26.491672+05	2022-06-16 04:48:26.491672+05	\N
\.


--
-- Data for Name: translation_my_information_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_my_information_page (id, lang_id, address, created_at, updated_at, deleted_at, birthday, update_password, save) FROM stdin;
d294138e-b808-41ae-9ac5-1826751fda3d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ваш адрес	2022-07-04 19:28:46.603058+05	2022-07-04 19:28:46.603058+05	\N	дата рождения	изменить пароль	запомнить
11074158-69f2-473a-b4fe-94304ff0d8a7	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	salgyňyz	2022-07-04 19:28:46.529935+05	2022-07-04 19:28:46.529935+05	\N	doglan senäň	açar sözi üýtget	ýatda sakla
\.


--
-- Data for Name: translation_payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_payment (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
5748ec03-5278-425c-babf-f7f2bf8d2efa	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Eltip bermek we töleg tertibi	Eltip bermek hyzmaty Aşgabat şäheriniň çägi bilen bir hatarda Büzmeýine we Änew şäherine hem elýeterlidir. Hyzmat mugt amala aşyrylýar;\nHer bir sargydyň jemi bahasy azyndan 150 manat bolmalydyr;\nSaýtdan sargyt edeniňizden soňra operator size jaň edip sargydy tassyklar (eger hemişelik müşderi bolsaňyz sargytlaryňyz islegiňize görä awtomatik usulda hem tassyklanýar);\nGirizen salgyňyz we telefon belgiňiz esasynda hyzmat amala aşyrylýar;\nSargyt tassyklanmadyk ýagdaýynda ol hasaba alynmaýar we ýerine ýetirilmeýär. Sargydyň tassyklanmagy üçin girizen telefon belgiňizden jaň kabul edip bilýändigiňize göz ýetiriň. Şeýle hem girizen salgyňyzyň dogrulygyny barlaň;\nSargydy barlap alanyňyzdan soňra töleg amala aşyrylýar. Eltip berijiniň size gowşurýan töleg resminamasynda siziň tölemeli puluňyz bellenendir. Töleg nagt we nagt däl görnüşde milli manatda amala aşyrylýar. Kabul edip tölegini geçiren harydyňyz yzyna alynmaýar;\nSargyt tassyklanandan soňra 24 sagadyň dowamynda eýesi tapylmasa ol güýjüni ýitirýär;	2022-06-25 11:37:47.362666+05	2022-06-25 11:37:47.362666+05	\N
ea7f4c0c-4b1a-41d3-94eb-e058aba9c99f	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Порядок доставки и оплаты	В настоящее время услуга по доставке осуществляется по городу Ашхабад, Бюзмеин и Анау. Услуга предоставляется бесплатно.\nМинимальный заказ должен составлять не менее 150 манат;\nПосле Вашего заказа по сайту, оператор позвонит Вам для подтверждения заказа (постоянным клиентам по их желанию подтверждение осуществляется автоматизированно);\nУслуга доставки выполняется по указанному Вами адресу и номеру телефона;\nЕсли заказ не подтвержден то данный заказ не регистрируется и не выполняется. Для подтверждения заказа, удостоверьтесь, что можете принять звонок по указанному Вами номеру телефона. Также проверьте правильность указанного Вами адреса;\nОплата выполняется после того, как Вы проверите и примите заказ. На платежном документе курьера указана сумма Вашей оплаты. Оплата выполняется наличными и через карту в национальной валюте. Принятый и оплаченный товар возврату не подлежит;\nЕсли не удается найти владельца заказа в течение 24 часов после подтверждения заказа, то данный заказ аннулируется;	2022-06-25 11:37:47.39047+05	2022-06-25 11:37:47.39047+05	\N
\.


--
-- Data for Name: translation_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_product (id, lang_id, product_id, name, description, created_at, updated_at, deleted_at) FROM stdin;
e7a616fc-650e-429c-a201-a513d7efe8d1	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	525af569-06b6-440a-ab5a-6ee0b39cf51d	Agardyjy we tegmil aýryjy serişde "G-OXI spray white" 600 ml	Agardyjy we tegmil aýryjy serişde "G-OXI spray white" 600 ml	2022-06-20 12:37:45.296265+05	2022-06-20 12:37:45.296265+05	\N
ce316f6c-cd44-4cfd-a8ed-a49f84198b81	aea98b93-7bdf-455b-9ad4-a259d69dc76e	525af569-06b6-440a-ab5a-6ee0b39cf51d	Пятновыводитель -отбеливатель "G-OXI spray white" 600 мл	Пятновыводитель -отбеливатель "G-OXI spray white" 600 мл	2022-06-20 12:37:45.31285+05	2022-06-20 12:37:45.31285+05	\N
5c0af669-683a-41d1-a301-fbd5c811d6a4	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0d356bb1-c695-4b29-b199-bdff967abfe2	Tegmil aýryjy serişde reňkli eşikler üçin"G-OXI spray white"600 ml	Tegmil aýryjy serişde reňkli eşikler üçin"G-OXI spray white"600 ml	2022-06-20 12:40:33.2306+05	2022-06-20 12:40:33.2306+05	\N
aa4f3f0a-0f35-4ef5-81d3-843a8e18e529	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0d356bb1-c695-4b29-b199-bdff967abfe2	Пятновыводитель для цветных вещей "G-OXI spray color" 600 мл	Пятновыводитель для цветных вещей "G-OXI spray color" 600 мл	2022-06-20 12:40:33.294199+05	2022-06-20 12:40:33.294199+05	\N
bebf6d07-edfc-4c1b-99bd-5fc8524c3269	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	e3f8aebb-1451-43a3-9e9b-582da01c8d08	Duş geli Nivea Men "Ekstrim serginlik" 250 ml	Duş geli Nivea Men "Ekstrim serginlik" 250 ml	2022-06-20 12:41:56.355507+05	2022-06-20 12:41:56.355507+05	\N
cd875a75-f302-42ae-9529-4a939b732907	aea98b93-7bdf-455b-9ad4-a259d69dc76e	e3f8aebb-1451-43a3-9e9b-582da01c8d08	Гель для душа Nivea Men "Эксремальная свежесть" 250 мл	Гель для душа Nivea Men "Эксремальная свежесть" 250 мл	2022-06-20 12:41:56.372671+05	2022-06-20 12:41:56.372671+05	\N
1ee48372-6d39-4baa-86d5-a5092dc11dcf	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	3b6d2d59-7ad4-4392-b7de-5a1d6bd003e1	(2+1) Süýt kremli sandwiç köke Ülker "Saklıköy" 100 gr (3 sany)	(2+1) Süýt kremli sandwiç köke Ülker "Saklıköy" 100 gr (3 sany)	2022-06-20 12:43:47.824085+05	2022-06-20 12:43:47.824085+05	\N
7aa7eb37-c2d9-4379-a715-0b96cf81acbf	aea98b93-7bdf-455b-9ad4-a259d69dc76e	3b6d2d59-7ad4-4392-b7de-5a1d6bd003e1	(2+1) Печенье сэндвич с молочным кремом Ülker "Saklıköy" 100 г (3 шт)	(2+1) Печенье сэндвич с молочным кремом Ülker "Saklıköy" 100 г (3 шт)	2022-06-20 12:43:47.840415+05	2022-06-20 12:43:47.840415+05	\N
c2d4374e-7e61-49f8-8220-fa07320fc59e	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	4a6b57e1-18d2-4aac-8346-576d3897967e	Naharhana üçin mikrofiber süpürgiç Parex (mämişi) + Gap-gaç ýuwmak üçin gubka "Viomax" (5 sany)	Naharhana üçin mikrofiber süpürgiç Parex (mämişi) + Gap-gaç ýuwmak üçin gubka "Viomax" (5 sany)	2022-06-20 12:47:56.984846+05	2022-06-20 12:47:56.984846+05	\N
d5682400-52be-4703-94c4-0865ad136dc1	aea98b93-7bdf-455b-9ad4-a259d69dc76e	4a6b57e1-18d2-4aac-8346-576d3897967e	Тряпка для кухни из микрофибры Parex (оранжевый) + Губки для посуды "Viomax" (5 шт)	Тряпка для кухни из микрофибры Parex (оранжевый) + Губки для посуды "Viomax" (5 шт)	2022-06-20 12:47:57.000678+05	2022-06-20 12:47:57.000678+05	\N
33685f9d-3701-42e2-9505-359c8dc5e81f	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	9a0572df-3006-426e-a623-11c0cbc930ea	Agyz boşlugyny çaýkamak üçin serişde Listerine Expert "Gijeki dikeltme" 400 ml + "Narpyzyň serginligi" 250 ml	Agyz boşlugyny çaýkamak üçin serişde Listerine Expert "Gijeki dikeltme" 400 ml + "Narpyzyň serginligi" 250 ml	2022-06-20 12:49:20.527635+05	2022-06-20 12:49:20.527635+05	\N
7bab00a3-b19f-4be9-88c3-267790830662	aea98b93-7bdf-455b-9ad4-a259d69dc76e	9a0572df-3006-426e-a623-11c0cbc930ea	Ополаскиватель для полости рта Listerine® Expert "Ночное Восстановление" 400 мл + "Свежая Мята" 250 мл	Ополаскиватель для полости рта Listerine® Expert "Ночное Восстановление" 400 мл + "Свежая Мята" 250 мл	2022-06-20 12:49:20.546448+05	2022-06-20 12:49:20.546448+05	\N
6a052dc2-9999-4052-a083-89a03cc84b6a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	1a8935fd-c6ab-4656-b173-826c487a2274	Sowgatlyk toplum çagalar üçin duş geli Johnson's 300 ml + çagalaryň agyzyny çaýkamak üçin serişde Listerine 250 ml	Sowgatlyk toplum çagalar üçin duş geli Johnson's 300 ml + çagalaryň agyzyny çaýkamak üçin serişde Listerine 250 ml	2022-06-21 10:15:31.585324+05	2022-06-21 10:15:31.585324+05	\N
16a440c8-d3e2-4dc9-8d48-9f3a733bb18c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	1a8935fd-c6ab-4656-b173-826c487a2274	Подарочный набор Johnson's детский гель для душа Johnson's 300 мл + детский ополаскиватель полости рта Listerine 250 мл	Подарочный набор Johnson's детский гель для душа Johnson's 300 мл + детский ополаскиватель полости рта Listerine 250 мл	2022-06-21 10:15:31.626337+05	2022-06-21 10:15:31.626337+05	\N
a62656f5-de15-4463-95e7-8a608c7f8469	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b4499063-096e-4fa6-9e21-a47185afd829	Sowgatlyk toplumy Head & Shoulders "Saç üçin balzam 275 ml + Goňaga garşy şampun 400 ml	Sowgatlyk toplumy Head & Shoulders "Saç üçin balzam 275 ml + Goňaga garşy şampun 400 ml	2022-06-21 10:17:07.70902+05	2022-06-21 10:17:07.70902+05	\N
f1d0c111-921c-4420-9460-7a64562500ce	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b4499063-096e-4fa6-9e21-a47185afd829	Подарочный Набор Head & Shoulders "Бальзам-ополаскиватель для волос 275 мл + Шампунь против перхоти 400 мл	Подарочный Набор Head & Shoulders "Бальзам-ополаскиватель для волос 275 мл + Шампунь против перхоти 400 мл	2022-06-21 10:17:07.718729+05	2022-06-21 10:17:07.718729+05	\N
7ed42c42-e90a-44c1-a079-44628ff773ab	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	538f0688-30ce-497b-9a0e-cd53d0d5239d	Duş üçin şampun-gel Faberlic "Calming Peak" 3x1, 380 ml	Duş üçin şampun-gel Faberlic "Calming Peak" 3x1, 380 ml	2022-06-21 10:21:35.539288+05	2022-06-21 10:21:35.539288+05	\N
c8163361-10c7-4402-9dfd-bc66277fcc8e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	538f0688-30ce-497b-9a0e-cd53d0d5239d	Шампунь-гель для душ "Calming Peak" 3в1, 380 мл	Шампунь-гель для душ "Calming Peak" 3в1, 380 мл	2022-06-21 10:21:35.556254+05	2022-06-21 10:21:35.556254+05	\N
a028180a-939c-4c4b-9c65-41f3e071a696	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0dc06a1f-e25a-4c3d-8310-09985e905262	Mikrofiber süpürgiç Mikrosan "Güderi" 40x50 sm (1 sany)	Mikrofiber süpürgiç Mikrosan "Güderi" 40x50 sm (1 sany)	2022-06-21 10:23:26.461681+05	2022-06-21 10:23:26.461681+05	\N
15acb070-ded8-4b88-9278-8026e2db07a4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0dc06a1f-e25a-4c3d-8310-09985e905262	Салфетка из микрофибры Mikrosan "Güderi" 40x50 см (1 шт)	Салфетка из микрофибры Mikrosan "Güderi" 40x50 см (1 шт)	2022-06-21 10:23:26.470717+05	2022-06-21 10:23:26.470717+05	\N
50f78e2b-e84f-4c20-be2d-01189e0d3dea	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ec4963db-c429-4135-9790-d3860c350bc5	Kiwi (500 gr)	Kiwi (500 gr)	2022-06-21 10:28:00.487529+05	2022-06-21 10:28:00.487529+05	\N
a7be771a-2f53-4ad7-878b-ca54fd302f2a	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ec4963db-c429-4135-9790-d3860c350bc5	Киви (500 г)	Киви (500 г)	2022-06-21 10:28:00.49526+05	2022-06-21 10:28:00.49526+05	\N
0856032c-b195-4f2d-a267-aabb59696d02	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	1fa25151-9c63-4554-a79d-faf6cc78ef69	Kofe Carte Noire, paket gapda 75 gr	Kofe Carte Noire, paket gapda 75 gr	2022-06-21 10:33:58.984686+05	2022-06-21 10:33:58.984686+05	\N
1dd4e733-0808-4aae-a477-991b52e2fd6d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	1fa25151-9c63-4554-a79d-faf6cc78ef69	Кофе Carte Noire, пакет 75 г	Кофе Carte Noire, пакет 75 г	2022-06-21 10:33:59.001019+05	2022-06-21 10:33:59.001019+05	\N
684f9b9e-3258-48ae-bb1e-2e3974a9924f	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d95aabd1-5a3a-47cc-aab5-9c6025e12280	Kofe Jacobs Monarch, çüýşe gapda 47.5 gr	Kofe Jacobs Monarch, çüýşe gapda 47.5 gr	2022-06-21 10:35:13.14977+05	2022-06-21 10:35:13.14977+05	\N
6c8e92ac-6a8a-48c0-ae3d-7719ba8cb142	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d95aabd1-5a3a-47cc-aab5-9c6025e12280	Кофе Jacobs Monarch, стеклянная банка 47.5 г	Кофе Jacobs Monarch, стеклянная банка 47.5 г	2022-06-21 10:35:13.159508+05	2022-06-21 10:35:13.159508+05	\N
78ae3973-8e6c-4fbf-a7c4-b3d79ac5d893	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d59506eb-aa84-4127-a411-5c5f95350d15	Kofe Jacobs Monarch, paket gapda 190 gr	Kofe Jacobs Monarch, paket gapda 190 gr	2022-06-21 10:37:39.655615+05	2022-06-21 10:37:39.655615+05	\N
f35164cc-3f9e-40fa-8bc3-ceab0c83f0d5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d59506eb-aa84-4127-a411-5c5f95350d15	Кофе Jacobs Monarch, пакет 190 г	Кофе Jacobs Monarch, пакет 190 г	2022-06-21 10:37:39.67119+05	2022-06-21 10:37:39.67119+05	\N
cbc5b906-39d0-4a33-aa8f-52c9e59d18a7	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ab6ba3a4-0d3a-4510-acd0-feb4fe48fc19	Kofe Jacobs Hazelnut (tokaý hozy), 3x1 kiçi paket 15 gr	Kofe Jacobs Hazelnut (tokaý hozy), 3x1 kiçi paket 15 gr	2022-06-21 10:39:18.399267+05	2022-06-21 10:39:18.399267+05	\N
58d42832-041a-46e7-9872-4a7d31bb447c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ab6ba3a4-0d3a-4510-acd0-feb4fe48fc19	Кофе Jacobs Hazelnut (лесной орех) 3в1, стик 15 г	Кофе Jacobs Hazelnut (лесной орех) 3в1, стик 15 г	2022-06-21 10:39:18.406658+05	2022-06-21 10:39:18.406658+05	\N
c25c372b-2802-4977-a8e3-333d5a364a16	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ce76ca4c-0ffb-4dd7-a252-3d3eaa6da732	Kofe Nescafe Classic, kiçi paket 2 gr	Kofe Nescafe Classic, kiçi paket 2 gr	2022-06-21 10:40:32.422306+05	2022-06-21 10:40:32.422306+05	\N
cd29576b-3cd2-4661-8060-bb14619ea840	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ce76ca4c-0ffb-4dd7-a252-3d3eaa6da732	Кофе Nescafe Classic, стик 2 гр	Кофе Nescafe Classic, стик 2 гр	2022-06-21 10:40:32.431712+05	2022-06-21 10:40:32.431712+05	\N
6feb3554-183c-4f16-a675-441da66eac95	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	2072a0fb-bbc4-4231-a7a4-dad00bb0a892	Gyzgyn şokolad "Kentcafe" karamelli 19 gr	Gyzgyn şokolad "Kentcafe" karamelli 19 gr	2022-06-21 10:41:30.458042+05	2022-06-21 10:41:30.458042+05	\N
83996676-1a1f-47aa-9e9c-25609fb714e7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	2072a0fb-bbc4-4231-a7a4-dad00bb0a892	Горячий шоколад "Kentcafe" карамель 19 г	Горячий шоколад "Kentcafe" карамель 19 г	2022-06-21 10:41:30.476604+05	2022-06-21 10:41:30.476604+05	\N
c8ad04e9-946b-4a9a-95ed-b940970635fb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	86f78ca3-177d-4c89-8693-7678066d7389	Şokolad Alpen Gold Oreo 90 gr	Şokolad Alpen Gold Oreo 90 gr	2022-06-21 10:47:21.563876+05	2022-06-21 10:47:21.563876+05	\N
a478458c-c871-4dae-9b6f-fc3add7b1686	aea98b93-7bdf-455b-9ad4-a259d69dc76e	86f78ca3-177d-4c89-8693-7678066d7389	Шоколад Alpen Gold Oreo 90 гр	Шоколад Alpen Gold Oreo 90 гр	2022-06-21 10:47:21.572369+05	2022-06-21 10:47:21.572369+05	\N
15744a2b-52c6-4d43-aa7a-b07d6a4313f2	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0a6863e2-7ed9-4fcd-9875-270fb778b33e	Süýtli şokolad Eti "Adicto" 70 gr	Süýtli şokolad Eti "Adicto" 70 gr	2022-06-21 10:48:09.320471+05	2022-06-21 10:48:09.320471+05	\N
503036b8-f987-4dd3-be4e-6a877f81c3f2	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0a6863e2-7ed9-4fcd-9875-270fb778b33e	Молочный шоколад Eti "Adicto" 70 гр	Молочный шоколад Eti "Adicto" 70 гр	2022-06-21 10:48:09.33918+05	2022-06-21 10:48:09.33918+05	\N
c26e5272-6bec-44da-9e51-40fd85339369	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	49381c4e-298d-43b7-8ae4-8dbe6e7da581	Şokolad Alpen Gold, Ajy 70% kakaoly 80 gr	Şokolad Alpen Gold, Ajy 70% kakaoly 80 gr	2022-06-21 10:49:08.688678+05	2022-06-21 10:49:08.688678+05	\N
c584d84a-ff88-4dc9-92d1-ec3e1199db3f	aea98b93-7bdf-455b-9ad4-a259d69dc76e	49381c4e-298d-43b7-8ae4-8dbe6e7da581	Шоколад Alpen Gold горький, 70% какао 80 гр	Шоколад Alpen Gold горький, 70% какао 80 гр	2022-06-21 10:49:08.696641+05	2022-06-21 10:49:08.696641+05	\N
96c55330-426e-46b5-ad5f-6767257c5a11	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c1f8c7cb-081e-4f99-aeb3-0bc84153295e	Garaňky şokoladly wafli KitKat Senses "Double Chocolate" 112 gr	Garaňky şokoladly wafli KitKat Senses "Double Chocolate" 112 gr	2022-06-21 10:49:54.757512+05	2022-06-21 10:49:54.757512+05	\N
c23b169b-90db-4185-bca6-3efef801d817	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c1f8c7cb-081e-4f99-aeb3-0bc84153295e	Шоколад KitKat Senses "Double Chocolate" 112 gr	Шоколад KitKat Senses "Double Chocolate" 112 gr	2022-06-21 10:49:54.775079+05	2022-06-21 10:49:54.775079+05	\N
913898de-64da-4b84-b161-1354d9df0708	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0cbe2487-c709-403f-a6c4-4f1a73fd3f78	Gara gözenekli şokolad Alpen Gold "Aerated" 80 gr	Gara gözenekli şokolad Alpen Gold "Aerated" 80 gr	2022-06-21 10:50:40.624338+05	2022-06-21 10:50:40.624338+05	\N
852ec6e7-4492-4bf9-af62-e32b4813774c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0cbe2487-c709-403f-a6c4-4f1a73fd3f78	Темный пористый шоколад Alpen Gold "Aerated" 80 г	Темный пористый шоколад Alpen Gold "Aerated" 80 г	2022-06-21 10:50:40.63247+05	2022-06-21 10:50:40.63247+05	\N
df9b6ed6-6cae-4f39-9dd4-2c4397c6f034	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	cbb0047a-e543-41a8-845b-8439d11638f4	Gara gözenekli şokolad Alpen Gold "Aerated" 80 gr	Gara gözenekli şokolad Alpen Gold "Aerated" 80 gr	2022-06-21 10:54:40.595161+05	2022-06-21 10:54:40.595161+05	\N
0c355cb3-3f46-4a0c-81b2-fa995dc3b371	aea98b93-7bdf-455b-9ad4-a259d69dc76e	cbb0047a-e543-41a8-845b-8439d11638f4	Темный пористый шоколад Alpen Gold "Aerated" 80 г	Темный пористый шоколад Alpen Gold "Aerated" 80 г	2022-06-21 10:54:40.602784+05	2022-06-21 10:54:40.602784+05	\N
828d8dc4-88eb-4104-9b15-e7b3a840be3d	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c5520db8-19de-4209-b99c-826a342210c3	Şokoladly wafli Kinder "Bueno" 43 gr	Şokoladly wafli Kinder "Bueno" 43 gr	2022-06-21 10:59:22.800629+05	2022-06-21 10:59:22.800629+05	\N
5cd26037-b4e0-45b3-8fcc-60b11ff48499	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c5520db8-19de-4209-b99c-826a342210c3	Вафли Kinder "Bueno" 43 г	Вафли Kinder "Bueno" 43 г	2022-06-21 10:59:22.818116+05	2022-06-21 10:59:22.818116+05	\N
d2f44cce-f945-45bf-8519-8de6034775e6	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ebc34352-64f7-4ad5-aa00-b1777efb3e56	Şokoladly wafli KitKat "Duo" 58 gr	Şokoladly wafli KitKat "Duo" 58 gr	2022-06-21 11:01:08.167983+05	2022-06-21 11:01:08.167983+05	\N
fd0acb7b-9fe2-4d15-896f-2309b337e241	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ebc34352-64f7-4ad5-aa00-b1777efb3e56	Батончик KitKat "Duo" 58 гр	Батончик KitKat "Duo" 58 гр	2022-06-21 11:01:08.175943+05	2022-06-21 11:01:08.175943+05	\N
942b30a8-13e5-446e-a8a7-dfbcdcb93b45	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	88753f91-4e73-4478-91c5-37b278984294	Hyýar Arzuw (1 kg)	Hyýar Arzuw (1 kg)	2022-06-21 11:03:19.559717+05	2022-06-21 11:03:19.559717+05	\N
e34d1c47-665c-40a0-a51d-a01e3054a6e7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	88753f91-4e73-4478-91c5-37b278984294	Огурцы арзув (1 кг)	Огурцы арзув (1 кг)	2022-06-21 11:03:19.577287+05	2022-06-21 11:03:19.577287+05	\N
5b77f25f-4a51-446e-b48d-20d1d786549b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	93096765-14be-4093-8e53-81caba6de3aa	Kelem (1-1.5 kg)	Kelem (1-1.5 kg)	2022-06-21 11:04:41.859527+05	2022-06-21 11:04:41.859527+05	\N
7c01aea8-6a81-49f5-8604-35116a74fb7b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	93096765-14be-4093-8e53-81caba6de3aa	Капуста ( 1-1.5 кг)	Капуста ( 1-1.5 кг)	2022-06-21 11:04:41.868596+05	2022-06-21 11:04:41.868596+05	\N
77b14c86-9ae2-445d-a6b5-c8d6e00d4c99	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	070d7096-2fdd-4327-b0b6-99b13af1570f	Kir ýuwujy serişde Persil "Premium" Color Gel 1.17 lt	Kir ýuwujy serişde Persil "Premium" Color Gel 1.17 lt	2022-06-21 11:06:35.18451+05	2022-06-21 11:06:35.18451+05	\N
6bffd071-c3bb-4bb5-b89f-8c1354abe2bf	aea98b93-7bdf-455b-9ad4-a259d69dc76e	070d7096-2fdd-4327-b0b6-99b13af1570f	Моющее средство Persil "Премиум" Color Gel 1.17 л	Моющее средство Persil "Премиум" Color Gel 1.17 л	2022-06-21 11:06:35.191806+05	2022-06-21 11:06:35.191806+05	\N
e1c575a4-fe91-4dec-86e6-ac3ccae7cf99	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	aee7abe3-c6cc-4562-bf67-3f87e952611b	Kir ýuwujy serişde Persil "Color" 1.5 kg	Kir ýuwujy serişde Persil "Color" 1.5 kg	2022-06-21 11:07:41.763048+05	2022-06-21 11:07:41.763048+05	\N
2fbed026-477a-4ac7-935c-3236611f645c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	aee7abe3-c6cc-4562-bf67-3f87e952611b	Моющее средство Persil "Color" 1.5 кг	Моющее средство Persil "Color" 1.5 кг	2022-06-21 11:07:41.770219+05	2022-06-21 11:07:41.770219+05	\N
facdee6c-25cb-4189-a172-c61f8f8fc406	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	205b50c5-da4b-4edf-adac-54f93dc99253	Kir ýuwujy serişde Persil "Свежесть от Vernel" elde ýuwmak üçin, 410 gr	Kir ýuwujy serişde Persil "Свежесть от Vernel" elde ýuwmak üçin, 410 gr	2022-06-21 11:10:49.543198+05	2022-06-21 11:10:49.543198+05	\N
efb37f46-e0b6-4edf-a8e0-75cd07d9e9d4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	205b50c5-da4b-4edf-adac-54f93dc99253	Моющее средство Persil для ручной стирки "360° Свежесть от Vernel" 410 г	Моющее средство Persil для ручной стирки "360° Свежесть от Vernel" 410 г	2022-06-21 11:10:49.551943+05	2022-06-21 11:10:49.551943+05	\N
89168c0e-971d-415a-b3f4-f01fb7096c14	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c9307e74-88a2-4d96-96ec-6f04e42ad0cb	Eşik ýuwmak üçin gel Qualita uniwersal 1000 ml (doy-pack)	Eşik ýuwmak üçin gel Qualita uniwersal 1000 ml (doy-pack)	2022-06-21 11:11:48.000826+05	2022-06-21 11:11:48.000826+05	\N
985569b5-a0cc-4690-9c88-301718afe3f5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c9307e74-88a2-4d96-96ec-6f04e42ad0cb	Гель для стирки Qualita универсальный 1000 мл (дойпак)	Гель для стирки Qualita универсальный 1000 мл (дойпак)	2022-06-21 11:11:48.008194+05	2022-06-21 11:11:48.008194+05	\N
6db83240-5656-47cd-a482-a7dca34edc3d	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f3208845-80d9-4ccb-9ad2-07a8ee2832c6	Sabyn Nivea "Ýertudana we süýt" 90 gr	Sabyn Nivea "Ýertudana we süýt" 90 gr	2022-06-21 11:13:34.458623+05	2022-06-21 11:13:34.458623+05	\N
1a2d4b3e-5845-4fc7-9474-fd7f5d2105c0	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f3208845-80d9-4ccb-9ad2-07a8ee2832c6	Мыло Nivea "Клубника и молоко" 90 гр	Мыло Nivea "Клубника и молоко" 90 гр	2022-06-21 11:13:34.476113+05	2022-06-21 11:13:34.476113+05	\N
8db4cde0-bef7-4a5b-8917-fc5086bc2b84	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7bab1a39-0c66-4c1e-9f9c-7f25e050daa5	Antibakterial sabyn Protex "Aloe" 90 g	Antibakterial sabyn Protex "Aloe" 90 g	2022-06-21 11:14:39.149524+05	2022-06-21 11:14:39.149524+05	\N
08c3c457-9fc4-43a6-83da-b6baa82fc56c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7bab1a39-0c66-4c1e-9f9c-7f25e050daa5	Антибактериальное туалетое мыло Protex "Aloe" 90 г	Антибактериальное туалетое мыло Protex "Aloe" 90 г	2022-06-21 11:14:39.166369+05	2022-06-21 11:14:39.166369+05	\N
f8d6e850-d989-4682-905e-b84878685f03	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c182ee68-2717-4604-b0ab-0e6994e61ff0	Suwuk sabyn Fa "Laým aromaty" 250 ml	Suwuk sabyn Fa "Laým aromaty" 250 ml	2022-06-21 11:17:33.339828+05	2022-06-21 11:17:33.339828+05	\N
80c8a381-1180-433d-acdc-c1ba48f14605	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c182ee68-2717-4604-b0ab-0e6994e61ff0	Жидкое мыло Fa "Аромат лайма" 250 мл	Жидкое мыло Fa "Аромат лайма" 250 мл	2022-06-21 11:17:33.349547+05	2022-06-21 11:17:33.349547+05	\N
6d25b5b4-def9-4d2a-83bf-cb5e786075f6	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	4ae4d83c-56ad-4d99-9d6f-e0dd77f9c982	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-06-21 11:18:25.053667+05	2022-06-21 11:18:25.053667+05	\N
397aea3c-6e37-43d8-b254-c780c2f8d248	aea98b93-7bdf-455b-9ad4-a259d69dc76e	4ae4d83c-56ad-4d99-9d6f-e0dd77f9c982	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-06-21 11:18:25.061144+05	2022-06-21 11:18:25.061144+05	\N
\.


--
-- Data for Name: translation_secure; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_secure (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
5988b64a-82ad-4ed0-bd1b-bdd0b3b05912	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ÖZARA YLALAŞYGY	Ynamdar - Internet Marketi (Mundan beýläk – “Ynamdar”) we www.ynamdar.com internet saýty (Mundan beýläk – “Saýt”) bilen, onuň agzasynyň (“Agza”) arasynda aşakdaky şertleri ýerine ýetirmek barada ylalaşyga gelindi.	2022-06-25 10:46:54.190131+05	2022-06-25 10:46:54.190131+05	\N
3579a847-ce74-4fbe-b10d-8aba83867857	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Пользовательское соглашение	Между Ынамдар – Интернет Маркетом (далее – “Ынамдар”) и интернет сайтом www.ynamdar.com (далее – “Сайт”), а также его клиентом (далее - “Клиент”) достигнуто соглашение по нижеследующим условиям.\n	2022-06-25 10:46:54.221498+05	2022-06-25 10:46:54.221498+05	\N
\.


--
-- Data for Name: translation_update_password_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_update_password_page (id, lang_id, title, verify_password, explanation, save, created_at, updated_at, deleted_at, password) FROM stdin;
de12082b-baab-4b83-ac07-119df09d1230	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	açar sözi üýtgetmek	açar sözi tassykla	siziň açar sözüňiz 5-20 uzynlygynda harp ýa-da sandan ybarat bolmalydyr	ýatda sakla	2022-07-05 10:35:08.867617+05	2022-07-05 10:35:08.867617+05	\N	açar sözi
5190ca93-7007-4db4-8105-65cc3b1af868	aea98b93-7bdf-455b-9ad4-a259d69dc76e	изменить пароль	Подтвердить Пароль	ключевое слово должно быть буквой или цифрой длиной от 5 до 20	запомнить	2022-07-05 10:35:08.984141+05	2022-07-05 10:35:08.984141+05	\N	ключевое слово
\.


--
-- Name: afisa afisa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.afisa
    ADD CONSTRAINT afisa_pkey PRIMARY KEY (id);


--
-- Name: banner banner_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.banner
    ADD CONSTRAINT banner_pkey PRIMARY KEY (id);


--
-- Name: brends brends_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brends
    ADD CONSTRAINT brends_pkey PRIMARY KEY (id);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: category_product category_product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_product
    ADD CONSTRAINT category_product_pkey PRIMARY KEY (id);


--
-- Name: category_shop category_shop_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_shop
    ADD CONSTRAINT category_shop_pkey PRIMARY KEY (id);


--
-- Name: company_address company_address_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.company_address
    ADD CONSTRAINT company_address_pkey PRIMARY KEY (id);


--
-- Name: company_phone company_phone_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.company_phone
    ADD CONSTRAINT company_phone_pkey PRIMARY KEY (id);


--
-- Name: company_setting company_setting_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.company_setting
    ADD CONSTRAINT company_setting_pkey PRIMARY KEY (id);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- Name: district district_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.district
    ADD CONSTRAINT district_pkey PRIMARY KEY (id);


--
-- Name: languages languages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_pkey PRIMARY KEY (id);


--
-- Name: likes likes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT likes_pkey PRIMARY KEY (id);


--
-- Name: products product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- Name: shops shops_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shops
    ADD CONSTRAINT shops_pkey PRIMARY KEY (id);


--
-- Name: translation_about translation_about_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_about
    ADD CONSTRAINT translation_about_pkey PRIMARY KEY (id);


--
-- Name: translation_afisa translation_afisa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_afisa
    ADD CONSTRAINT translation_afisa_pkey PRIMARY KEY (id);


--
-- Name: translation_category translation_category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_category
    ADD CONSTRAINT translation_category_pkey PRIMARY KEY (id);


--
-- Name: translation_contact translation_contact_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_contact
    ADD CONSTRAINT translation_contact_pkey PRIMARY KEY (id);


--
-- Name: translation_district translation_district_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_district
    ADD CONSTRAINT translation_district_pkey PRIMARY KEY (id);


--
-- Name: translation_footer translation_footer_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_footer
    ADD CONSTRAINT translation_footer_pkey PRIMARY KEY (id);


--
-- Name: translation_header translation_header_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_header
    ADD CONSTRAINT translation_header_pkey PRIMARY KEY (id);


--
-- Name: translation_my_information_page translation_my_information_page_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_my_information_page
    ADD CONSTRAINT translation_my_information_page_pkey PRIMARY KEY (id);


--
-- Name: translation_payment translation_payment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_payment
    ADD CONSTRAINT translation_payment_pkey PRIMARY KEY (id);


--
-- Name: translation_product translation_product_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_product
    ADD CONSTRAINT translation_product_pkey PRIMARY KEY (id);


--
-- Name: translation_secure translation_secure_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_secure
    ADD CONSTRAINT translation_secure_pkey PRIMARY KEY (id);


--
-- Name: translation_update_password_page translation_update_password_page_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_update_password_page
    ADD CONSTRAINT translation_update_password_page_pkey PRIMARY KEY (id);


--
-- Name: translation_afisa fk_afisa_translation_afisa; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_afisa
    ADD CONSTRAINT fk_afisa_translation_afisa FOREIGN KEY (afisa_id) REFERENCES public.afisa(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: products fk_brend_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_brend_product FOREIGN KEY (brend_id) REFERENCES public.brends(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: category_product fk_category_category_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_product
    ADD CONSTRAINT fk_category_category_product FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: category_shop fk_category_category_shop; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_shop
    ADD CONSTRAINT fk_category_category_shop FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: categories fk_category_child_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT fk_category_child_category FOREIGN KEY (parent_category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_category fk_category_translation_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_category
    ADD CONSTRAINT fk_category_translation_category FOREIGN KEY (category_id) REFERENCES public.categories(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: likes fk_customer_like; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT fk_customer_like FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_district fk_district_translation_district; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_district
    ADD CONSTRAINT fk_district_translation_district FOREIGN KEY (district_id) REFERENCES public.district(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: company_address fk_language_company_address; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.company_address
    ADD CONSTRAINT fk_language_company_address FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_about fk_language_translation_about; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_about
    ADD CONSTRAINT fk_language_translation_about FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_afisa fk_language_translation_afisa; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_afisa
    ADD CONSTRAINT fk_language_translation_afisa FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_category fk_language_translation_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_category
    ADD CONSTRAINT fk_language_translation_category FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_contact fk_language_translation_contact; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_contact
    ADD CONSTRAINT fk_language_translation_contact FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_district fk_language_translation_district; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_district
    ADD CONSTRAINT fk_language_translation_district FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_header fk_language_translation_header; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_header
    ADD CONSTRAINT fk_language_translation_header FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_my_information_page fk_language_translation_my_information_page; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_my_information_page
    ADD CONSTRAINT fk_language_translation_my_information_page FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_payment fk_language_translation_payment; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_payment
    ADD CONSTRAINT fk_language_translation_payment FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_product fk_language_translation_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_product
    ADD CONSTRAINT fk_language_translation_product FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_secure fk_language_translation_secure; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_secure
    ADD CONSTRAINT fk_language_translation_secure FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_update_password_page fk_language_translation_update_password_page; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_update_password_page
    ADD CONSTRAINT fk_language_translation_update_password_page FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_footer fk_languages_translation_footer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_footer
    ADD CONSTRAINT fk_languages_translation_footer FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: category_product fk_product_category_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_product
    ADD CONSTRAINT fk_product_category_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: likes fk_product_like; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.likes
    ADD CONSTRAINT fk_product_like FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_product fk_product_translation_product; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_product
    ADD CONSTRAINT fk_product_translation_product FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: category_shop fk_shop_category_shop; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.category_shop
    ADD CONSTRAINT fk_shop_category_shop FOREIGN KEY (shop_id) REFERENCES public.shops(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

