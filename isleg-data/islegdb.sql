--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12 (Ubuntu 12.12-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.12 (Ubuntu 12.12-0ubuntu0.20.04.1)

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
    image character varying,
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
    image character varying,
    url text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.banner OWNER TO postgres;

--
-- Name: basket; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.basket (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    product_id uuid,
    customer_id uuid,
    quantity_of_product bigint,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.basket OWNER TO postgres;

--
-- Name: brends; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brends (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying,
    image character varying,
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
    image character varying,
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
    logo character varying,
    favicon character varying,
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
    deleted_at timestamp with time zone,
    email character varying
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
-- Name: images; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.images (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    product_id uuid,
    small character varying,
    medium character varying,
    large character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.images OWNER TO postgres;

--
-- Name: languages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.languages (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name_short character varying(5),
    flag character varying,
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
-- Name: main_image; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.main_image (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    product_id uuid,
    small character varying,
    medium character varying,
    large character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.main_image OWNER TO postgres;

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
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    limit_amount bigint,
    is_new boolean DEFAULT false
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
-- Name: translation_basket_page; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_basket_page (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    quantity_of_goods character varying DEFAULT 'uytget'::character varying,
    total_price character varying DEFAULT 'uytget'::character varying,
    discount character varying DEFAULT 'uytget'::character varying,
    delivery character varying DEFAULT 'uytget'::character varying,
    total character varying DEFAULT 'uytget'::character varying,
    currency character varying DEFAULT 'uytget'::character varying,
    to_order character varying DEFAULT 'uytget'::character varying,
    your_basket character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_basket_page OWNER TO postgres;

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
    deleted_at timestamp with time zone,
    basket character varying DEFAULT 'uytget'::character varying,
    email character varying DEFAULT 'uytget'::character varying,
    add_to_basket character varying DEFAULT 'uytget'::character varying
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
-- Name: translation_my_order_page; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_my_order_page (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    orders character varying DEFAULT 'uytget'::character varying,
    date character varying DEFAULT 'uytget'::character varying,
    price character varying DEFAULT 'uytget'::character varying,
    currency character varying DEFAULT 'uytget'::character varying,
    image character varying DEFAULT 'uytget'::character varying,
    name character varying DEFAULT 'uytget'::character varying,
    brend character varying DEFAULT 'uytget'::character varying,
    code character varying DEFAULT 'uytget'::character varying,
    amount character varying DEFAULT 'uytget'::character varying,
    total_price character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_my_order_page OWNER TO postgres;

--
-- Name: translation_order_page; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_order_page (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    content character varying DEFAULT 'uytget'::character varying,
    type_of_payment character varying DEFAULT 'uytget'::character varying,
    choose_a_delivery_time character varying DEFAULT 'uytget'::character varying,
    your_address character varying DEFAULT 'uytget'::character varying,
    mark character varying DEFAULT 'uytget'::character varying,
    to_order character varying DEFAULT 'uytget'::character varying,
    tomorrow character varying DEFAULT 'uytget'::character varying,
    cash character varying DEFAULT 'uytget'::character varying,
    payment_terminal character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_order_page OWNER TO postgres;

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

COPY public.afisa (id, image, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: banner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banner (id, image, url, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: basket; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.basket (id, product_id, customer_id, quantity_of_product, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: brends; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.brends (id, name, image, created_at, updated_at, deleted_at) FROM stdin;
214be879-65c3-4710-86b4-3fc3bce2e974	Arcalyk	uploads/brend24badfac-613d-4aa3-881b-952bd14994b5.jpeg	2022-06-16 14:14:05.416191+05	2022-06-16 14:14:05.416191+05	\N
ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	Tut	uploads/brend4f68381a-aa73-4168-90b3-66c1a17cd5c5.jpeg	2022-06-16 14:14:25.908903+05	2022-06-16 14:14:25.908903+05	\N
fdd259c2-794a-42b9-a3ad-9e91502af23e	Koka Kola	uploads/brend75f655c6-bcf5-47b2-ba01-d112cba64e81.jpg	2022-07-12 17:54:39.242004+05	2022-07-12 17:54:39.242004+05	\N
f53a27b4-7810-4d8f-bd45-edad405d92b9	Maral Koke	uploads/brend7827fcfe-f8a9-4747-8c34-b55af2488b29.jpeg	2022-07-12 17:57:46.472194+05	2022-07-12 17:57:46.472194+05	\N
46b13f0a-d584-4ad3-b270-437ecdc51449	Taze Ay	uploads/brend993b6484-657d-4662-abe2-922170abe75b.jpeg	2022-07-12 18:16:12.889441+05	2022-07-12 18:16:12.889441+05	\N
c4bcda34-7332-4ae5-8129-d7538d63fee4	Buzz	uploads/brend/67f6bc90-a0ef-4828-b17b-8b00e930f1f1.jpeg	2022-08-12 10:36:10.886455+05	2022-08-12 10:36:10.886455+05	\N
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, parent_category_id, image, is_home_category, created_at, updated_at, deleted_at) FROM stdin;
f745d171-68e6-42e2-b339-cb3c210cda55	b982bd86-0a0f-4950-baad-5a131e9b728e		f	2022-06-16 13:45:48.828786+05	2022-06-16 13:45:48.828786+05	\N
d4cb1359-6c23-4194-8e3c-21ed8cec8373	5bb9a4e7-9992-418f-b551-537844d371da		f	2022-06-16 13:48:04.517774+05	2022-06-16 13:48:04.517774+05	\N
7f453dd0-7b2e-480d-a8be-fcfa23bd863e	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:43:07.336084+05	2022-06-20 09:43:07.336084+05	\N
29ed85bb-11eb-4458-bbf3-5a5644d167d6	\N	uploads/categoryeaae1626-7e9f-4db9-abf6-f454ade813d3.jpeg	f	2022-06-20 09:41:17.575565+05	2022-06-20 09:41:17.575565+05	\N
66772380-c161-4c45-9350-a45e765193e2	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:45:34.38667+05	2022-06-20 09:45:34.38667+05	\N
338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:46:01.119337+05	2022-06-20 09:46:01.119337+05	\N
45765130-7f97-4f0c-b886-f70b75e02610	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 10:11:06.648938+05	2022-06-20 10:11:06.648938+05	\N
fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	5bb9a4e7-9992-418f-b551-537844d371da		f	2022-06-16 13:47:18.854741+05	2022-06-16 13:47:18.854741+05	\N
02bd4413-8586-49ab-802e-16304e756a8b	\N	uploads/category0684921b-251d-405f-8b30-30964be0b3d2.jpeg	f	2022-06-16 13:43:22.644619+05	2022-06-16 13:43:22.644619+05	\N
5bb9a4e7-9992-418f-b551-537844d371da	02bd4413-8586-49ab-802e-16304e756a8b		f	2022-06-16 13:46:44.575803+05	2022-06-16 13:46:44.575803+05	\N
b982bd86-0a0f-4950-baad-5a131e9b728e	02bd4413-8586-49ab-802e-16304e756a8b		f	2022-06-16 13:44:16.430875+05	2022-06-16 13:44:16.430875+05	\N
\.


--
-- Data for Name: category_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category_product (id, category_id, product_id, created_at, updated_at, deleted_at) FROM stdin;
08113084-23d4-40d4-963a-2a3e997e3760	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	440507af-648f-4b56-b126-ca75d0370731	2022-09-09 15:05:30.567827+05	2022-09-09 15:05:30.567827+05	\N
795e3137-6e9d-4d02-bc66-b29bba82eb50	f745d171-68e6-42e2-b339-cb3c210cda55	440507af-648f-4b56-b126-ca75d0370731	2022-09-09 15:05:30.577487+05	2022-09-09 15:05:30.577487+05	\N
bc7a8442-b53b-489f-bba5-0cfe2ebccd0f	d4cb1359-6c23-4194-8e3c-21ed8cec8373	440507af-648f-4b56-b126-ca75d0370731	2022-09-09 15:05:30.587125+05	2022-09-09 15:05:30.587125+05	\N
2b8f5a3b-d43c-4259-b2bc-ae9f70791199	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	3ca5ade3-d2ce-40ba-9be6-601105b5205a	2022-09-09 15:06:14.377898+05	2022-09-09 15:06:14.377898+05	\N
471b1f48-9ac1-4d58-a48d-d4109829a26b	f745d171-68e6-42e2-b339-cb3c210cda55	3ca5ade3-d2ce-40ba-9be6-601105b5205a	2022-09-09 15:06:14.388538+05	2022-09-09 15:06:14.388538+05	\N
83416c89-96f4-4b33-9723-7291141b1a7b	d4cb1359-6c23-4194-8e3c-21ed8cec8373	3ca5ade3-d2ce-40ba-9be6-601105b5205a	2022-09-09 15:06:14.399242+05	2022-09-09 15:06:14.399242+05	\N
dde9d428-0222-4472-ba81-ae036c2c243b	d4cb1359-6c23-4194-8e3c-21ed8cec8373	b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	2022-09-10 14:18:39.434622+05	2022-09-10 14:18:39.434622+05	\N
17c3c22e-3833-4801-b6ae-6779a54184c0	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	2022-09-10 14:18:39.600908+05	2022-09-10 14:18:39.600908+05	\N
6598ac97-f383-4db2-8293-62cf6780c2b1	29ed85bb-11eb-4458-bbf3-5a5644d167d6	b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	2022-09-10 14:18:39.645834+05	2022-09-10 14:18:39.645834+05	\N
791a4383-51a7-402c-9611-63fd9ad02a39	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	ba935176-cf8c-4684-ab24-3ff11e5f176a	2022-09-10 14:19:35.523965+05	2022-09-10 14:19:35.523965+05	\N
7fe28b15-ef1b-4642-9404-814d0754c06f	29ed85bb-11eb-4458-bbf3-5a5644d167d6	ba935176-cf8c-4684-ab24-3ff11e5f176a	2022-09-10 14:19:35.536176+05	2022-09-10 14:19:35.536176+05	\N
ec5e5f44-c54c-4e39-a200-22ded9e67304	66772380-c161-4c45-9350-a45e765193e2	ba935176-cf8c-4684-ab24-3ff11e5f176a	2022-09-10 14:19:35.545838+05	2022-09-10 14:19:35.545838+05	\N
26a34503-f790-4374-80d8-dc69b47e0d2a	29ed85bb-11eb-4458-bbf3-5a5644d167d6	624dbb53-dba3-424e-9d3d-af4e12b6c067	2022-09-10 14:20:22.202099+05	2022-09-10 14:20:22.202099+05	\N
5a24eff4-f045-4ba2-8792-9d2742b368a3	66772380-c161-4c45-9350-a45e765193e2	624dbb53-dba3-424e-9d3d-af4e12b6c067	2022-09-10 14:20:22.212281+05	2022-09-10 14:20:22.212281+05	\N
aeb3b540-c5cc-4094-a202-fb26c6b3d567	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	624dbb53-dba3-424e-9d3d-af4e12b6c067	2022-09-10 14:20:22.224897+05	2022-09-10 14:20:22.224897+05	\N
798b15b0-8359-4bb3-a4fd-1d572cb600f9	66772380-c161-4c45-9350-a45e765193e2	32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	2022-09-10 14:21:26.858504+05	2022-09-10 14:21:26.858504+05	\N
8ca97452-c660-406b-9172-687b65959f17	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	2022-09-10 14:21:26.869891+05	2022-09-10 14:21:26.869891+05	\N
c7065d31-3169-4328-b954-789c1ca3b0f3	45765130-7f97-4f0c-b886-f70b75e02610	32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	2022-09-10 14:21:26.880201+05	2022-09-10 14:21:26.880201+05	\N
41339f48-e2f1-4e9c-97aa-735231af8f84	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	abc05e23-5d72-41db-969a-662442da399f	2022-09-10 14:22:00.759274+05	2022-09-10 14:22:00.759274+05	\N
d6ecd238-1c31-40a9-9b4b-a1d1512d6d09	45765130-7f97-4f0c-b886-f70b75e02610	abc05e23-5d72-41db-969a-662442da399f	2022-09-10 14:22:00.771022+05	2022-09-10 14:22:00.771022+05	\N
90c71ee4-f6b8-4006-9ae5-8a12d358cd9f	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	abc05e23-5d72-41db-969a-662442da399f	2022-09-10 14:22:00.781172+05	2022-09-10 14:22:00.781172+05	\N
bd3cd1b2-1340-44b0-b52b-e29f762b9bbb	45765130-7f97-4f0c-b886-f70b75e02610	8cb36bb0-9103-4031-b83b-4180552d74ca	2022-09-10 14:23:40.050661+05	2022-09-10 14:23:40.050661+05	\N
e3791c61-f004-410b-93c0-fced77708c51	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	8cb36bb0-9103-4031-b83b-4180552d74ca	2022-09-10 14:23:40.061226+05	2022-09-10 14:23:40.061226+05	\N
1b1ef6f3-44b9-4c27-926a-a720eb41c547	5bb9a4e7-9992-418f-b551-537844d371da	8cb36bb0-9103-4031-b83b-4180552d74ca	2022-09-10 14:23:40.071302+05	2022-09-10 14:23:40.071302+05	\N
a05d1a55-5032-4bd9-8663-2b3c401dc458	b982bd86-0a0f-4950-baad-5a131e9b728e	113e11af-f785-4d08-b1d0-c63ad65a65a8	2022-09-10 14:24:15.984478+05	2022-09-10 14:24:15.984478+05	\N
7641db5b-b5e9-4d6f-9ca2-90cd21d9f34c	5bb9a4e7-9992-418f-b551-537844d371da	113e11af-f785-4d08-b1d0-c63ad65a65a8	2022-09-10 14:24:15.99574+05	2022-09-10 14:24:15.99574+05	\N
368e5ad9-695f-4539-b7c5-1c6ccfb8384f	02bd4413-8586-49ab-802e-16304e756a8b	113e11af-f785-4d08-b1d0-c63ad65a65a8	2022-09-10 14:24:16.006798+05	2022-09-10 14:24:16.006798+05	\N
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
75706251-06ea-41c1-905f-95ed8b4132f8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Улица Азади 23, Ашхабад	2022-06-22 18:44:50.239558+05	2022-06-22 18:44:50.239558+05	\N
d2c66808-e5fe-435f-ba01-cb717f80d9e0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	adres_tm	2022-06-22 18:44:50.21776+05	2022-08-22 09:33:42.14835+05	\N
\.


--
-- Data for Name: company_phone; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_phone (id, phone, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: company_setting; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_setting (id, logo, favicon, email, instagram, created_at, updated_at, deleted_at) FROM stdin;
7d193677-e0b1-4df0-be88-dc6e16a47ca7	uploads/logode9c4f45-acba-42ce-b435-e744631a98ba.jpeg	uploads/favicon8a413c02-108d-4d2f-8e92-d24a18cea1d3.jpeg	isleg-bazar@gmail.com	@islegbazarinstagram	2022-06-15 19:57:04.54457+05	2022-06-15 19:57:04.54457+05	\N
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (id, full_name, phone_number, password, birthday, gender, addresses, created_at, updated_at, deleted_at, email) FROM stdin;
7e872c52-0d23-4086-8c45-43000b57332e	Muhammetmyrat	+99363747155	$2a$14$1uOYIcXK4lzyBnhm.L/dW.TD8c9ZqTzAiCsOMCCRRzxiKnDAU2gFK	\N	\N	\N	2022-08-02 23:41:59.869254+05	2022-08-02 23:41:59.869254+05	\N	m.bayramov@salam.tm
7fafe6f8-c6b6-4bcc-9063-e98c113902c5	jjednkjwedjed	+99363747156	$2a$14$WPTcXE1j871GQ/n2i2CX9.RjyRIyR4bBqCj6b/vchJB1TjYC6v0XK	\N	\N	\N	2022-08-02 23:52:46.544849+05	2022-08-02 23:52:46.544849+05	\N	ewkdnewj@gmail.com
38615c8c-1af5-424f-b7a3-071d38c42b86	Aly Muhammedow	+99363234587	$2a$14$Ep0/A9EAbgV/BD.UdQ6KQOU0DCpr2C8n6du8li5nPKYz.xIQb2HgC	\N	\N	\N	2022-08-23 19:59:07.331615+05	2022-08-23 19:59:07.331615+05	\N	aly@gmail.com
9b1a0831-9943-4aa9-aa2a-3507743a5de4	Berdi	+99361235698	$2a$14$S5nCg8mlGD..q3didZiCFuaocaEPA35ugIfpovdEoM7p5I8TOTX0K	\N	\N	\N	2022-08-23 20:31:28.830324+05	2022-08-23 20:31:28.830324+05	\N	berdi@gmail.com
\.


--
-- Data for Name: district; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.district (id, price, created_at, updated_at, deleted_at) FROM stdin;
a58294d3-efe5-4cb7-82d3-8df8c37563c5	15	2022-06-25 10:23:25.640364+05	2022-06-25 10:23:25.640364+05	\N
\.


--
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.images (id, product_id, small, medium, large, created_at, updated_at, deleted_at) FROM stdin;
866168e7-ae22-4cba-a51e-7dcd5f2be760	440507af-648f-4b56-b126-ca75d0370731	uploads/product/599ece23-1468-42a8-9cf6-f25f0fa5baf2.jpg	uploads/product/f2f708e7-5ede-4bf0-bd51-f44f67f8c781.jpg	uploads/product/c1be43aa-f3a2-465a-944c-0f816ffc1d8c.jpg	2022-09-09 15:05:30.520393+05	2022-09-09 15:05:30.520393+05	\N
df1f6fa0-d8bb-49ea-84fe-d98f13eaa7bb	440507af-648f-4b56-b126-ca75d0370731	uploads/product/299e1765-4a74-448f-ab96-56172041aee9.jpeg	uploads/product/24a88aad-d500-43d0-b7e8-300a0fade151.jpeg	uploads/product/0dbfe4b2-36b7-4a11-af46-73bb3bdad108.jpeg	2022-09-09 15:05:30.532481+05	2022-09-09 15:05:30.532481+05	\N
5a3738b7-4af2-404b-910a-5ee0a9482824	3ca5ade3-d2ce-40ba-9be6-601105b5205a	uploads/product/8b6fabd8-41c2-4949-99b9-915331e1a117.jpg	uploads/product/ab88bf0d-a30d-47ae-8167-6940ef7a0be4.jpg	uploads/product/5d83d182-16fc-4278-84e3-e0efc98a4464.jpg	2022-09-09 15:06:14.332499+05	2022-09-09 15:06:14.332499+05	\N
fc721d41-9cdc-4dfd-8115-4861ee48f0fa	3ca5ade3-d2ce-40ba-9be6-601105b5205a	uploads/product/372c8fc2-8d40-4afe-903e-20028ab16fd3.jpeg	uploads/product/420e351a-8950-47c5-8c97-fe44d4147a37.jpeg	uploads/product/12cfa8d8-7fa6-4a30-9188-e74bcc2fbe4f.jpeg	2022-09-09 15:06:14.343239+05	2022-09-09 15:06:14.343239+05	\N
c8ff8bab-0ccb-4728-a506-ec502a5fd111	b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	uploads/product/f57a2f66-d5bf-4a26-bdac-ed95c7a1c63f.jpg	uploads/product/63ef0d90-6473-4a38-9eaf-1efbf53c15a3.jpg	uploads/product/0f739540-7a7a-4d9e-86d3-a90bbbf95e2e.jpg	2022-09-10 14:18:39.304893+05	2022-09-10 14:18:39.304893+05	\N
f4caae13-5130-4619-aa45-0611acd02243	b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	uploads/product/873f120e-5e9a-4e6e-9739-923c7bd86a74.jpeg	uploads/product/442910d5-8049-4c36-90a2-e07d49ef774a.jpeg	uploads/product/1dc90d34-a2af-4edd-8034-602263e867a1.jpeg	2022-09-10 14:18:39.333588+05	2022-09-10 14:18:39.333588+05	\N
7427512d-fa9a-4fcc-ad1a-d5aa02dcc289	ba935176-cf8c-4684-ab24-3ff11e5f176a	uploads/product/a6815dcf-db60-4a25-862f-b28eceaf8346.jpg	uploads/product/d6ea36a5-91db-4366-970f-740e3a606da8.jpg	uploads/product/bd6a21f5-8079-4c36-8532-23215c06e8ee.jpg	2022-09-10 14:19:35.479121+05	2022-09-10 14:19:35.479121+05	\N
80f87db4-a2cf-4317-8369-4988dc63bac1	ba935176-cf8c-4684-ab24-3ff11e5f176a	uploads/product/fc5b68fb-d08c-4c3f-83b7-348d0fdc94e2.jpeg	uploads/product/ce0892fb-7cc8-4f0a-9068-4ce3cec9ebc5.jpeg	uploads/product/ee418e78-99cd-4bf2-8b60-bcec3d8906fe.jpeg	2022-09-10 14:19:35.489668+05	2022-09-10 14:19:35.489668+05	\N
3056dc04-7015-4739-9274-03316a9cf29f	624dbb53-dba3-424e-9d3d-af4e12b6c067	uploads/product/0a06f282-70a6-4e43-a472-5675e111071c.jpg	uploads/product/60cc22aa-a96e-4ff3-a388-337c8c23f020.jpg	uploads/product/efda0561-e1b1-40d9-92f2-0cd517d5f64e.jpg	2022-09-10 14:20:22.157337+05	2022-09-10 14:20:22.157337+05	\N
4d0bc935-4f13-4f02-ab33-a60acb385352	624dbb53-dba3-424e-9d3d-af4e12b6c067	uploads/product/df1bf2b2-77fc-4306-97ce-4e0d8aeda81e.jpeg	uploads/product/14988f98-7575-48a5-8bec-f897e0751034.jpeg	uploads/product/4e1bdf44-8de3-4fb9-9a82-7009e34191dc.jpeg	2022-09-10 14:20:22.168297+05	2022-09-10 14:20:22.168297+05	\N
5c3270ae-c1df-41e4-a9a8-88eeb734d58a	32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	uploads/product/66be7084-ac0f-4ec2-a4f9-49e0a01c9755.jpg	uploads/product/4b370996-3443-4abc-9b63-fc3001cca39b.jpg	uploads/product/5152cdb9-b521-4167-a168-6c44bae5308a.jpg	2022-09-10 14:21:26.814456+05	2022-09-10 14:21:26.814456+05	\N
189bb245-f912-4426-acae-ff04c18f5108	32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	uploads/product/300df876-7c66-4df4-a446-c74a5ac20420.jpeg	uploads/product/c671671b-71de-4b21-aed2-2f794d18578f.jpeg	uploads/product/5be4f0f9-9c49-4072-98dd-9757cb46f50a.jpeg	2022-09-10 14:21:26.825802+05	2022-09-10 14:21:26.825802+05	\N
0228f244-d11b-4967-b116-9358dd299d2c	abc05e23-5d72-41db-969a-662442da399f	uploads/product/7408d015-f192-4515-86ef-77b73ce11365.jpg	uploads/product/68de7ba5-076a-4ab0-9da2-5048eadf917b.jpg	uploads/product/b9deb354-4517-4d22-a6d5-6fa62217988c.jpg	2022-09-10 14:22:00.625052+05	2022-09-10 14:22:00.625052+05	\N
465add6a-ce57-464f-ac83-324c4112b5e8	abc05e23-5d72-41db-969a-662442da399f	uploads/product/681695d9-1e5c-4983-b7c5-06569dcfc6f9.jpeg	uploads/product/cfe60b42-52c3-412b-9e51-e41ccb2590f6.jpeg	uploads/product/78879277-bddc-4b00-b105-fda9b903bb80.jpeg	2022-09-10 14:22:00.681609+05	2022-09-10 14:22:00.681609+05	\N
886b3cc5-b704-48ef-b4a8-f1718042cc05	8cb36bb0-9103-4031-b83b-4180552d74ca	uploads/product/3cd994d2-ce6b-4eb0-829d-b4283a9abef1.jpg	uploads/product/d4751956-7a9d-49e5-a227-a8391b099782.jpg	uploads/product/dc01188f-5f06-4c7e-954b-9d4c4728a392.jpg	2022-09-10 14:23:40.004546+05	2022-09-10 14:23:40.004546+05	\N
2aa304aa-2de4-40b6-a047-e0b715c0572d	8cb36bb0-9103-4031-b83b-4180552d74ca	uploads/product/0861dd77-d918-4528-aa02-3fbb3dd5349f.jpeg	uploads/product/3ce22377-ccf3-4ff1-a19d-3cb58b71255e.jpeg	uploads/product/9d4b0f8b-8091-49cb-bb8b-9974c339fb9a.jpeg	2022-09-10 14:23:40.015767+05	2022-09-10 14:23:40.015767+05	\N
a1913f4d-0583-4d22-94e1-190aad16f4e5	113e11af-f785-4d08-b1d0-c63ad65a65a8	uploads/product/0b199f38-8372-4d56-94c4-2256de7168bd.jpg	uploads/product/947bf320-59c1-48da-b3b6-71b264f7f93a.jpg	uploads/product/da78d705-2643-495c-a1e1-a5c353af4ff2.jpg	2022-09-10 14:24:15.938689+05	2022-09-10 14:24:15.938689+05	\N
151a5db7-41ab-4e04-ac02-110f45646f14	113e11af-f785-4d08-b1d0-c63ad65a65a8	uploads/product/49da7ca2-03bf-4ea4-be7d-d045c64aa9c7.jpeg	uploads/product/f280b885-06a1-4e22-bd6a-c44febca24d9.jpeg	uploads/product/7f931e46-a765-4cc1-9d2c-a83e53a3bfe6.jpeg	2022-09-10 14:24:15.949955+05	2022-09-10 14:24:15.949955+05	\N
\.


--
-- Data for Name: languages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.languages (id, name_short, flag, created_at, updated_at, deleted_at) FROM stdin;
8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	tm	uploads/language17b99bd1-f52d-41db-b4e6-1ecff03e0fd0.jpeg	2022-06-15 19:53:06.041686+05	2022-06-15 19:53:06.041686+05	\N
aea98b93-7bdf-455b-9ad4-a259d69dc76e	ru	uploads/language1c24e3a6-173e-4264-a631-f099d15495dd.jpeg	2022-06-15 19:53:21.29491+05	2022-06-15 19:53:21.29491+05	\N
\.


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.likes (id, product_id, customer_id, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: main_image; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.main_image (id, product_id, small, medium, large, created_at, updated_at, deleted_at) FROM stdin;
8000093a-ec1a-4a4a-9dbb-e413a285f2a5	440507af-648f-4b56-b126-ca75d0370731	uploads/product/c281c7c4-78f0-4df3-8ff1-c4ff29046e6f.jpg	uploads/product/bfe00698-d9a5-4ac6-a6d7-c517793d907f.jpg	uploads/product/029b47d6-33be-4604-a15f-051b36cc758e.jpg	2022-09-09 15:05:30.510382+05	2022-09-09 15:05:30.510382+05	\N
16984148-7376-40c2-9d3c-5b19c9667c8d	3ca5ade3-d2ce-40ba-9be6-601105b5205a	uploads/product/c8430183-1739-4e82-86e1-bed7ce2e4039.jpg	uploads/product/1a150d8b-a00c-4c5a-af8f-370fbe08443d.jpg	uploads/product/465cfb6d-c355-415e-aba1-459cbdaaf011.jpg	2022-09-09 15:06:14.323642+05	2022-09-09 15:06:14.323642+05	\N
83011a8e-47f4-44f5-a762-dc6081b629c1	b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	uploads/product/ad7651eb-39e4-4fe5-8bee-4a7fa6d66e1a.jpg	uploads/product/cb785369-a857-4b26-934c-3f0983ed67e0.jpg	uploads/product/3da0e40f-35d2-4c4e-b861-52571addf824.jpg	2022-09-10 14:18:39.211898+05	2022-09-10 14:18:39.211898+05	\N
bce02dd7-97a6-4f3c-be5a-86050eb59600	ba935176-cf8c-4684-ab24-3ff11e5f176a	uploads/product/f9c88be7-871a-4c98-a144-2873b6e6f25d.jpg	uploads/product/729802be-6637-48b5-ae03-dad88c062e1b.jpg	uploads/product/43414a10-1a2b-4b93-a6ab-d19f0a1713b0.jpg	2022-09-10 14:19:35.468254+05	2022-09-10 14:19:35.468254+05	\N
c6f9dca8-641f-4986-859e-6458e224b13b	624dbb53-dba3-424e-9d3d-af4e12b6c067	uploads/product/2c1c06fb-fca3-4e05-a6b3-96f70cf17823.jpg	uploads/product/f9b3531d-ae50-4a8c-b886-3ba9e0d09c18.jpg	uploads/product/76aff707-2af8-4b07-bb4d-91d1659abf16.jpg	2022-09-10 14:20:22.135125+05	2022-09-10 14:20:22.135125+05	\N
1d4d6e25-01e6-4fda-9563-fa77d5c6e6e5	32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	uploads/product/063fe01c-a7e8-4d7c-8ea2-ede107c3f84d.jpg	uploads/product/44a13ecb-a521-431a-96e0-29e07814e5be.jpg	uploads/product/6c436214-d81e-41d6-ab77-878ea1162bdd.jpg	2022-09-10 14:21:26.802947+05	2022-09-10 14:21:26.802947+05	\N
3b5b0c76-cd04-4cec-875e-462384ffcbf0	abc05e23-5d72-41db-969a-662442da399f	uploads/product/c7ae901e-be27-4e1c-940d-735733c66352.jpg	uploads/product/c434aad8-79c4-4122-bc13-0da630a220ba.jpg	uploads/product/230bfb55-5a23-42be-9ed4-bd65d6262d8d.jpg	2022-09-10 14:22:00.570895+05	2022-09-10 14:22:00.570895+05	\N
66a288af-4c57-4658-b51a-10b94a895872	8cb36bb0-9103-4031-b83b-4180552d74ca	uploads/product/d371a5e1-5fcb-4910-a2f2-a84f12c00a2e.jpg	uploads/product/87b76fd1-e242-47ea-b858-348d5c88f39a.jpg	uploads/product/7c02f0e7-b767-4dda-94d0-c48c12bfd816.jpg	2022-09-10 14:23:39.983692+05	2022-09-10 14:23:39.983692+05	\N
020ae5f4-60b9-434e-a20b-a1162d4c18dd	113e11af-f785-4d08-b1d0-c63ad65a65a8	uploads/product/f8b4c39f-fdbe-4bb5-9baa-df35b41d59cb.jpg	uploads/product/a204aad4-22b9-4fdb-a193-d17d94663800.jpg	uploads/product/e5a9e5ee-fe95-438d-b442-66532d5aab65.jpg	2022-09-10 14:24:15.917399+05	2022-09-10 14:24:15.917399+05	\N
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, brend_id, price, old_price, amount, product_code, created_at, updated_at, deleted_at, limit_amount, is_new) FROM stdin;
3ca5ade3-d2ce-40ba-9be6-601105b5205a	214be879-65c3-4710-86b4-3fc3bce2e974	65	68	1000	151fwe51we	2022-09-09 15:06:14.29935+05	2022-09-09 15:06:14.29935+05	\N	100	f
ba935176-cf8c-4684-ab24-3ff11e5f176a	fdd259c2-794a-42b9-a3ad-9e91502af23e	65	68	1000	151fwe51we	2022-09-10 14:19:35.444385+05	2022-09-10 14:19:35.444385+05	\N	100	f
abc05e23-5d72-41db-969a-662442da399f	c4bcda34-7332-4ae5-8129-d7538d63fee4	65	68	1000	151fwe51we	2022-09-10 14:22:00.51937+05	2022-09-10 14:22:00.51937+05	\N	100	f
624dbb53-dba3-424e-9d3d-af4e12b6c067	f53a27b4-7810-4d8f-bd45-edad405d92b9	65	68	1000	151fwe51we	2022-09-10 14:20:22.105259+05	2022-09-10 14:20:22.105259+05	\N	100	t
8cb36bb0-9103-4031-b83b-4180552d74ca	214be879-65c3-4710-86b4-3fc3bce2e974	65	68	1000	151fwe51we	2022-09-10 14:23:39.953886+05	2022-09-10 14:23:39.953886+05	\N	100	f
440507af-648f-4b56-b126-ca75d0370731	214be879-65c3-4710-86b4-3fc3bce2e974	65	68	1000	151fwe51we	2022-09-09 15:05:30.484381+05	2022-09-09 15:05:30.484381+05	\N	100	f
b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	65	68	1000	151fwe51we	2022-09-10 14:18:39.139255+05	2022-09-10 14:18:39.139255+05	\N	100	t
32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	46b13f0a-d584-4ad3-b270-437ecdc51449	65	68	1000	151fwe51we	2022-09-10 14:21:26.773503+05	2022-09-10 14:21:26.773503+05	\N	100	f
113e11af-f785-4d08-b1d0-c63ad65a65a8	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	65	68	1000	151fwe51we	2022-09-10 14:24:15.897483+05	2022-09-10 14:24:15.897483+05	\N	100	f
\.


--
-- Data for Name: shops; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shops (id, owner_name, address, phone_number, running_time, created_at, updated_at, deleted_at) FROM stdin;
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
\.


--
-- Data for Name: translation_basket_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_basket_page (id, lang_id, quantity_of_goods, total_price, discount, delivery, total, currency, to_order, your_basket, created_at, updated_at, deleted_at) FROM stdin;
456dcb5a-fabb-47f8-b216-0cddd3077124	aea98b93-7bdf-455b-9ad4-a259d69dc76e	quantity_of_goods_ru	total_price_ru	discount_ru	delivery_ru	total_ru	currency_ru	to_order_ru	your_basket_ru	2022-08-30 12:36:24.978404+05	2022-08-30 12:36:37.967063+05	\N
51b3699e-1c7b-442a-be7b-6b2ad1f111b4	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	quantity_of_goods	total_price	discount	delivery	total	currency	to_order	your_basket	2022-08-30 12:36:24.978404+05	2022-08-30 12:39:12.849615+05	\N
\.


--
-- Data for Name: translation_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_category (id, lang_id, category_id, name, created_at, updated_at, deleted_at) FROM stdin;
21520180-13e2-4c2b-a5f9-866c2e59ba87	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f745d171-68e6-42e2-b339-cb3c210cda55	Kiçi paket kofeler	2022-06-16 13:45:48.889727+05	2022-06-16 13:45:48.889727+05	\N
ee2f97fb-8c6c-4e38-bdb3-bf769bc95d3b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d4cb1359-6c23-4194-8e3c-21ed8cec8373	Batonçikler	2022-06-16 13:48:04.581888+05	2022-06-16 13:48:04.581888+05	\N
85469cf2-f48a-4e73-800d-ebf599aaeaba	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	29ed85bb-11eb-4458-bbf3-5a5644d167d6	Arzanladyş we Aksiýalar	2022-06-20 09:41:17.756928+05	2022-06-20 09:41:17.756928+05	\N
8a91bcb0-fcce-4a4f-80ff-a2896c0cc36a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	Arzanladyşdaky harytlar	2022-06-20 09:43:07.368782+05	2022-06-20 09:43:07.368782+05	\N
34f4cdb5-04b9-48c0-b5b0-0045a02aa094	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	66772380-c161-4c45-9350-a45e765193e2	Aksiýadaky harytlar	2022-06-20 09:45:34.450534+05	2022-06-20 09:45:34.450534+05	\N
e224ecfc-6daa-4df5-8112-74846fc44867	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	Sowgatlyk toplumlar	2022-06-20 09:46:01.148565+05	2022-06-20 09:46:01.148565+05	\N
3b756a33-bf2c-4d04-af57-962a3226d00b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	45765130-7f97-4f0c-b886-f70b75e02610	Täze harytlar	2022-06-20 10:11:06.719528+05	2022-06-20 10:11:06.719528+05	\N
ab35a97a-dfd1-4100-8e84-d34e74e9a02e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f745d171-68e6-42e2-b339-cb3c210cda55	Кофе в пакетиках	2022-06-16 13:45:48.906024+05	2022-06-16 13:45:48.906024+05	\N
ea104eaf-c3fd-4f2d-88bf-dffc14d48dc5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d4cb1359-6c23-4194-8e3c-21ed8cec8373	Батончики	2022-06-16 13:48:04.597499+05	2022-06-16 13:48:04.597499+05	\N
bbdd06a4-2dce-4c99-bf05-cf4e911776c7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	29ed85bb-11eb-4458-bbf3-5a5644d167d6	Распродажи и Акции	2022-06-20 09:41:17.941489+05	2022-06-20 09:41:17.941489+05	\N
ce573dfd-6af8-4e64-8260-8746a090acd7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	Продукция со скидкой	2022-06-20 09:43:07.377729+05	2022-06-20 09:43:07.377729+05	\N
713cc05f-6a9d-4dae-88b5-dde2e564480c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	66772380-c161-4c45-9350-a45e765193e2	Продукция в категории Акции	2022-06-20 09:45:34.466904+05	2022-06-20 09:45:34.466904+05	\N
53959762-0b63-4100-ae13-4bbf8c015fec	aea98b93-7bdf-455b-9ad4-a259d69dc76e	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	Подарочные наборы	2022-06-20 09:46:01.408239+05	2022-06-20 09:46:01.408239+05	\N
2d22961c-ef08-4238-ae54-c00593c0073c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	45765130-7f97-4f0c-b886-f70b75e02610	Новые продукты	2022-06-20 10:11:06.735056+05	2022-06-20 10:11:06.735056+05	\N
4eef5d40-9aad-4101-b36b-9026dd3dfb51	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b982bd86-0a0f-4950-baad-5a131e9b728e	name_tm	2022-06-16 13:44:16.499713+05	2022-06-16 13:44:16.499713+05	\N
10a8b5ec-a3ca-448d-975b-83b3a7a8c0d2	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b982bd86-0a0f-4950-baad-5a131e9b728e	name_ru	2022-06-16 13:44:16.515874+05	2022-06-16 13:44:16.515874+05	\N
4eb6bcbf-91f2-4505-a27e-cc3f96f2b829	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	Plitkalar	2022-06-16 13:47:18.888998+05	2022-06-16 13:47:18.888998+05	\N
53fb44c7-45fb-49f0-a433-aaed23b2dfc0	aea98b93-7bdf-455b-9ad4-a259d69dc76e	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	Плитки	2022-06-16 13:47:18.942159+05	2022-06-16 13:47:18.942159+05	\N
bff34c21-04c1-4cea-bfaf-c8f9ce7e2bfe	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	02bd4413-8586-49ab-802e-16304e756a8b	name_tm	2022-06-16 13:43:22.674562+05	2022-06-16 13:43:22.674562+05	\N
0e400414-a80c-449d-8842-dd6667b45c73	aea98b93-7bdf-455b-9ad4-a259d69dc76e	02bd4413-8586-49ab-802e-16304e756a8b	name_ru	2022-06-16 13:43:22.681932+05	2022-06-16 13:43:22.681932+05	\N
e099e7f6-1b97-4f70-8f29-f586ab6697d0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	5bb9a4e7-9992-418f-b551-537844d371da	Şokolad we Keksler	2022-06-16 13:46:44.657849+05	2022-06-16 13:46:44.657849+05	\N
415a0711-2482-44b3-8f03-923dca28bd5d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	5bb9a4e7-9992-418f-b551-537844d371da	Шоколады и Кексы	2022-06-16 13:46:44.673892+05	2022-06-16 13:46:44.673892+05	\N
\.


--
-- Data for Name: translation_contact; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_contact (id, lang_id, full_name, email, phone, letter, company_phone, imo, company_email, instagram, created_at, updated_at, deleted_at, button_text) FROM stdin;
f1693167-0c68-4a54-9831-56f124d629a3	aea98b93-7bdf-455b-9ad4-a259d69dc76e	at_ru	mail_ru	phone_ru	letter ru	cp ru	imo ru	ce ru	instagram ru	2022-06-27 11:29:48.050553+05	2022-06-27 11:29:48.050553+05	\N	Отправить
73253999-7355-42b4-8700-94de76f0058a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	at_tm	mail_tm	phone_tm	letter_tm	cp_tm	imo_tm	ce_tm	ins_tm	2022-06-27 11:29:47.914891+05	2022-06-27 11:29:47.914891+05	\N	ugrat
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
84b5504f-1056-4b44-94dd-a7819148da66	aea98b93-7bdf-455b-9ad4-a259d69dc76e	О нас	Порядок доставки и оплаты	Коммуникация	Обслуживания и Политика Конфиденциальности	Все права защищены	2022-06-22 15:23:32.793161+05	2022-06-22 15:23:32.793161+05	\N
12dc4c16-5712-4bff-a957-8e16d450b4fb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Biz Barada	Eltip bermek we töleg tertibi	Aragatnaşyk	Ulanyş düzgünleri we gizlinlik şertnamasy	Ähli hukuklary goraglydyr	2022-06-22 15:23:32.716064+05	2022-06-22 15:23:32.716064+05	\N
\.


--
-- Data for Name: translation_header; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_header (id, lang_id, research, phone, password, forgot_password, sign_in, sign_up, name, password_verification, verify_secure, my_information, my_favorites, my_orders, log_out, created_at, updated_at, deleted_at, basket, email, add_to_basket) FROM stdin;
eaf206e6-d515-4bdb-9323-a047cd0edae5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	gözleg	telefon	parol	Acar sozumi unutdym	ulgama girmek	agza bolmak	Ady	Acar sozi tassyklamak	Ulanyş Düzgünlerini we Gizlinlik Şertnamasyny okadym we kabul edýärin	maglumatym	halanlarym	sargytlarym	cykmak	2022-06-16 04:48:26.460534+05	2022-06-16 04:48:26.460534+05	\N	sebet	uytget	uytget
9154e800-2a92-47de-b4ff-1e63b213e5f7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	поиск	tелефон	пароль	забыл пароль	войти	зарегистрироваться	имя	Подтвердить Пароль	Я прочитал и принимаю Условия Обслуживания и Политика Конфиденциальности	моя информация	мои любимые	мои заказы	выйти	2022-06-16 04:48:26.491672+05	2022-06-16 04:48:26.491672+05	\N	корзина	uytget	uytget
\.


--
-- Data for Name: translation_my_information_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_my_information_page (id, lang_id, address, created_at, updated_at, deleted_at, birthday, update_password, save) FROM stdin;
d294138e-b808-41ae-9ac5-1826751fda3d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ваш адрес	2022-07-04 19:28:46.603058+05	2022-07-04 19:28:46.603058+05	\N	дата рождения	изменить пароль	запомнить
11074158-69f2-473a-b4fe-94304ff0d8a7	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	salgyňyz	2022-07-04 19:28:46.529935+05	2022-07-04 19:28:46.529935+05	\N	doglan senäň	açar sözi üýtget	ýatda sakla
\.


--
-- Data for Name: translation_my_order_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_my_order_page (id, lang_id, orders, date, price, currency, image, name, brend, code, amount, total_price, created_at, updated_at, deleted_at) FROM stdin;
6f30b588-94d8-49f5-a558-a90c2ec9150e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	orders_ru	date_ru	price_ru	currency_ru	image_ru	name_ru	brend_ru	code_ru	amount_ru	total_price_ru	2022-09-02 13:04:39.394714+05	2022-09-02 13:04:39.394714+05	\N
ff43b90d-e22d-4364-b358-6fd56bb3a305	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	orders	date	price	currency	image	name	brend	code	amount	total_price	2022-09-02 13:04:39.36328+05	2022-09-02 13:12:48.119751+05	\N
\.


--
-- Data for Name: translation_order_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_order_page (id, lang_id, content, type_of_payment, choose_a_delivery_time, your_address, mark, to_order, tomorrow, cash, payment_terminal, created_at, updated_at, deleted_at) FROM stdin;
474a15e9-1a05-49aa-9a61-c92837d9c9a8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	content_ru	type_of_payment_ru	choose_a_delivery_time_ru	your_address_ru	mark_ru	to_order_ru	tomorrow_ru	cash_ru	payment_terminal_ru	2022-09-01 12:47:16.802639+05	2022-09-01 12:47:16.802639+05	\N
75810722-07fd-400e-94b4-cd230de08cbf	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	content	type_of_payment	choose_a_delivery_time	your_address	mark	to_order	tomorrow	cash	payment_terminal	2022-09-01 12:47:16.720956+05	2022-09-01 12:55:25.638676+05	\N
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
cd44b671-cc28-4105-98eb-547418cb0fea	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-10 14:18:39.368603+05	2022-09-10 14:18:39.368603+05	\N
01fde8cc-9a7f-431a-a0fe-4e77915ae3e5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b5cfdcaf-02ff-4097-bdb9-2651a0cf366f	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-10 14:18:39.42174+05	2022-09-10 14:18:39.42174+05	\N
5886097a-b173-45aa-b8d3-b41e8139fe6b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ba935176-cf8c-4684-ab24-3ff11e5f176a	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-10 14:19:35.502218+05	2022-09-10 14:19:35.502218+05	\N
707313c7-eb2b-4732-91fa-539a3c4b3a09	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ba935176-cf8c-4684-ab24-3ff11e5f176a	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-10 14:19:35.511601+05	2022-09-10 14:19:35.511601+05	\N
0dc9c666-a44d-4ffa-b509-df7c39e19451	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	624dbb53-dba3-424e-9d3d-af4e12b6c067	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-10 14:20:22.178859+05	2022-09-10 14:20:22.178859+05	\N
bc011ea7-2465-4602-b519-1363843fa955	aea98b93-7bdf-455b-9ad4-a259d69dc76e	624dbb53-dba3-424e-9d3d-af4e12b6c067	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-10 14:20:22.190023+05	2022-09-10 14:20:22.190023+05	\N
ec474669-3d3b-4194-a02e-f9e8a0b84224	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-10 14:21:26.835649+05	2022-09-10 14:21:26.835649+05	\N
9dda6426-0193-4c31-98dc-65b3030851b7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	32cd4c58-c96c-4b05-8158-bc9f7f7d02d4	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-10 14:21:26.846323+05	2022-09-10 14:21:26.846323+05	\N
986453ec-dc59-4a88-b3e1-a25c0da45e3e	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	abc05e23-5d72-41db-969a-662442da399f	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-10 14:22:00.736874+05	2022-09-10 14:22:00.736874+05	\N
b873390c-85a2-450a-8c1a-5543e94709c7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	abc05e23-5d72-41db-969a-662442da399f	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-10 14:22:00.747206+05	2022-09-10 14:22:00.747206+05	\N
6de48174-8bd6-438b-ac78-678da2928572	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	8cb36bb0-9103-4031-b83b-4180552d74ca	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-10 14:23:40.028279+05	2022-09-10 14:23:40.028279+05	\N
6a70a9da-74a1-432c-bd7a-b226de4e7704	aea98b93-7bdf-455b-9ad4-a259d69dc76e	8cb36bb0-9103-4031-b83b-4180552d74ca	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-10 14:23:40.037728+05	2022-09-10 14:23:40.037728+05	\N
205bab01-15f3-492b-81f4-f88f286b69b5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	440507af-648f-4b56-b126-ca75d0370731	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-09 15:05:30.543832+05	2022-09-09 15:05:30.543832+05	\N
0d9cc1f9-3eb7-4200-abcf-cbaa2981c806	aea98b93-7bdf-455b-9ad4-a259d69dc76e	440507af-648f-4b56-b126-ca75d0370731	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-09 15:05:30.553884+05	2022-09-09 15:05:30.553884+05	\N
601d4a31-9e16-41a5-b30b-982e0db95665	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	3ca5ade3-d2ce-40ba-9be6-601105b5205a	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-09 15:06:14.354841+05	2022-09-09 15:06:14.354841+05	\N
baf40ae6-5e0b-41d4-a66e-81b36cc9dd9e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	3ca5ade3-d2ce-40ba-9be6-601105b5205a	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-09 15:06:14.365684+05	2022-09-09 15:06:14.365684+05	\N
c6dd04a9-5f29-4292-9fbf-8cea0b73c808	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	113e11af-f785-4d08-b1d0-c63ad65a65a8	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-10 14:24:15.961132+05	2022-09-10 14:24:15.961132+05	\N
c409998c-54f6-4016-9913-5bd56aa23cc6	aea98b93-7bdf-455b-9ad4-a259d69dc76e	113e11af-f785-4d08-b1d0-c63ad65a65a8	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-10 14:24:15.971943+05	2022-09-10 14:24:15.971943+05	\N
\.


--
-- Data for Name: translation_secure; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_secure (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
3579a847-ce74-4fbe-b10d-8aba83867857	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Пользовательское соглашение	Между Ынамдар – Интернет Маркетом (далее – “Ынамдар”) и интернет сайтом www.ynamdar.com (далее – “Сайт”), а также его клиентом (далее - “Клиент”) достигнуто соглашение по нижеследующим условиям.\n	2022-06-25 10:46:54.221498+05	2022-06-25 10:46:54.221498+05	\N
5988b64a-82ad-4ed0-bd1b-bdd0b3b05912	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ÖZARA YLALAŞYGY	Ynamdar - Internet Marketi (Mundan beýläk – “Ynamdar”) we www.ynamdar.com internet saýty (Mundan beýläk – “Saýt”) bilen, onuň agzasynyň (“Agza”) arasynda aşakdaky şertleri ýerine ýetirmek barada ylalaşyga gelindi.	2022-06-25 10:46:54.190131+05	2022-06-25 10:46:54.190131+05	\N
\.


--
-- Data for Name: translation_update_password_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_update_password_page (id, lang_id, title, verify_password, explanation, save, created_at, updated_at, deleted_at, password) FROM stdin;
5190ca93-7007-4db4-8105-65cc3b1af868	aea98b93-7bdf-455b-9ad4-a259d69dc76e	изменить пароль	Подтвердить Пароль	ключевое слово должно быть буквой или цифрой длиной от 5 до 20	запомнить	2022-07-05 10:35:08.984141+05	2022-07-05 10:35:08.984141+05	\N	ключевое слово
de12082b-baab-4b83-ac07-119df09d1230	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	açar sözi üýtgetmek	açar sözi tassykla	siziň açar sözüňiz 5-20 uzynlygynda harp ýa-da sandan ybarat bolmalydyr	ýatda sakla	2022-07-05 10:35:08.867617+05	2022-07-05 10:35:08.867617+05	\N	açar sözi
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
-- Name: basket basket_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.basket
    ADD CONSTRAINT basket_pkey PRIMARY KEY (id);


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
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


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
-- Name: main_image main_image_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.main_image
    ADD CONSTRAINT main_image_pkey PRIMARY KEY (id);


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
-- Name: translation_basket_page translation_basket_page_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_basket_page
    ADD CONSTRAINT translation_basket_page_pkey PRIMARY KEY (id);


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
-- Name: translation_my_order_page translation_my_order_page_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_my_order_page
    ADD CONSTRAINT translation_my_order_page_pkey PRIMARY KEY (id);


--
-- Name: translation_order_page translation_order_page_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_order_page
    ADD CONSTRAINT translation_order_page_pkey PRIMARY KEY (id);


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
-- Name: basket fk_customer_basket; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.basket
    ADD CONSTRAINT fk_customer_basket FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON UPDATE CASCADE ON DELETE CASCADE;


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
-- Name: basket fk_product_basket; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.basket
    ADD CONSTRAINT fk_product_basket FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


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
-- Name: translation_basket_page language_translation_basket_page; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_basket_page
    ADD CONSTRAINT language_translation_basket_page FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_my_order_page language_translation_my_order_page; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_my_order_page
    ADD CONSTRAINT language_translation_my_order_page FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_order_page language_translation_order_page; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_order_page
    ADD CONSTRAINT language_translation_order_page FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: images products_images; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT products_images FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: main_image products_main_image; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.main_image
    ADD CONSTRAINT products_main_image FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

