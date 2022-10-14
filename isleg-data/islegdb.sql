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
-- Name: cart; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cart (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    product_id uuid,
    customer_id uuid,
    quantity_of_product bigint,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.cart OWNER TO postgres;

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
-- Name: customer_address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer_address (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    customer_id uuid,
    address character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    is_active boolean DEFAULT true
);


ALTER TABLE public.customer_address OWNER TO postgres;

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
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    email character varying,
    is_register boolean DEFAULT true
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
-- Name: order_dates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_dates (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    date character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.order_dates OWNER TO postgres;

--
-- Name: order_times; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_times (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    order_date_id uuid,
    "time" character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.order_times OWNER TO postgres;

--
-- Name: ordered_products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ordered_products (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    product_id uuid,
    quantity_of_product integer,
    order_id uuid,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.ordered_products OWNER TO postgres;

--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    customer_id uuid,
    customer_mark character varying,
    order_time character varying,
    payment_type character varying,
    total_price character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    order_number integer NOT NULL
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: orders_order_number_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_order_number_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_order_number_seq OWNER TO postgres;

--
-- Name: orders_order_number_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.orders_order_number_seq OWNED BY public.orders.order_number;


--
-- Name: payment_types; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.payment_types (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    type character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.payment_types OWNER TO postgres;

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
    deleted_at timestamp with time zone,
    empty_the_basket character varying DEFAULT 'uytget'::character varying
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
-- Name: translation_order_dates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_order_dates (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid,
    order_date_id uuid,
    date character varying DEFAULT 'uytget'::character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_order_dates OWNER TO postgres;

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
    deleted_at timestamp with time zone,
    slug character varying DEFAULT 'uytget'::character varying
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
-- Name: orders order_number; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN order_number SET DEFAULT nextval('public.orders_order_number_seq'::regclass);


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
-- Data for Name: cart; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cart (id, product_id, customer_id, quantity_of_product, created_at, updated_at, deleted_at) FROM stdin;
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
d82042be-0468-446f-a06e-c569fc6967de	f745d171-68e6-42e2-b339-cb3c210cda55	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	2022-09-17 14:54:57.077351+05	2022-09-17 14:54:57.077351+05	\N
715c66a2-32b6-449b-8ed1-2f656ed07c2f	d4cb1359-6c23-4194-8e3c-21ed8cec8373	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	2022-09-17 14:54:57.085055+05	2022-09-17 14:54:57.085055+05	\N
7b67d6b8-6c31-48c7-bb98-e5645af6ad9d	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	2022-09-17 14:54:57.09628+05	2022-09-17 14:54:57.09628+05	\N
703562cb-0dc5-4ea8-bb9b-a632150d9406	d4cb1359-6c23-4194-8e3c-21ed8cec8373	b2b165a3-2261-4d67-8160-0e239ecd99b5	2022-09-17 14:55:35.543212+05	2022-09-17 14:55:35.543212+05	\N
de77f0ce-1f16-4610-b62f-d8bd50339cd8	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	b2b165a3-2261-4d67-8160-0e239ecd99b5	2022-09-17 14:55:35.552501+05	2022-09-17 14:55:35.552501+05	\N
599b6d48-0077-4960-86f9-947addb08210	29ed85bb-11eb-4458-bbf3-5a5644d167d6	b2b165a3-2261-4d67-8160-0e239ecd99b5	2022-09-17 14:55:35.563229+05	2022-09-17 14:55:35.563229+05	\N
8199e176-7f47-4421-8f14-2fd116564d80	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	2022-09-17 14:56:05.486473+05	2022-09-17 14:56:05.486473+05	\N
c43bf0b8-a01d-41fd-9614-2e75cd19b413	29ed85bb-11eb-4458-bbf3-5a5644d167d6	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	2022-09-17 14:56:05.497232+05	2022-09-17 14:56:05.497232+05	\N
3c4747b8-1925-43a0-959c-b7c2279f84fd	66772380-c161-4c45-9350-a45e765193e2	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	2022-09-17 14:56:05.507917+05	2022-09-17 14:56:05.507917+05	\N
1ddb4057-86f2-436e-a0be-f09d1ff807b3	29ed85bb-11eb-4458-bbf3-5a5644d167d6	d731b17a-ae8d-4561-ad67-0f431d5c529b	2022-09-17 14:56:36.24299+05	2022-09-17 14:56:36.24299+05	\N
401e2f12-c159-4a9f-9365-ae3a5729fa07	66772380-c161-4c45-9350-a45e765193e2	d731b17a-ae8d-4561-ad67-0f431d5c529b	2022-09-17 14:56:36.252888+05	2022-09-17 14:56:36.252888+05	\N
76bd693b-0d60-4613-beed-79fee34431b0	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	d731b17a-ae8d-4561-ad67-0f431d5c529b	2022-09-17 14:56:36.263964+05	2022-09-17 14:56:36.263964+05	\N
5b934855-c387-4039-8029-108464f90297	66772380-c161-4c45-9350-a45e765193e2	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	2022-09-17 14:57:07.276145+05	2022-09-17 14:57:07.276145+05	\N
3f5f53c8-b03d-4868-bbc1-7dfe2466000f	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	2022-09-17 14:57:07.288945+05	2022-09-17 14:57:07.288945+05	\N
1d9851c7-5fe4-47c9-9d9b-b300777610dc	45765130-7f97-4f0c-b886-f70b75e02610	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	2022-09-17 14:57:07.298188+05	2022-09-17 14:57:07.298188+05	\N
9664b220-8b44-489d-a4cd-b925b2136a0a	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	d4156225-082e-4f0f-9b2c-85268114433a	2022-09-17 14:57:34.665141+05	2022-09-17 14:57:34.665141+05	\N
8913356d-d03e-4aba-b2d7-c175c82eea3f	45765130-7f97-4f0c-b886-f70b75e02610	d4156225-082e-4f0f-9b2c-85268114433a	2022-09-17 14:57:34.676377+05	2022-09-17 14:57:34.676377+05	\N
7529cbf9-7d44-4e8d-a39f-308e9d85b40f	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	d4156225-082e-4f0f-9b2c-85268114433a	2022-09-17 14:57:34.687184+05	2022-09-17 14:57:34.687184+05	\N
3ab03ec7-2fb0-44e4-b5c9-c0cbf4c42fa3	45765130-7f97-4f0c-b886-f70b75e02610	81b84c5d-9759-4b86-978a-649c8ef79660	2022-09-17 14:58:10.07692+05	2022-09-17 14:58:10.07692+05	\N
0740063b-0d3a-46f5-987c-072fdf0ad3cf	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	81b84c5d-9759-4b86-978a-649c8ef79660	2022-09-17 14:58:10.086907+05	2022-09-17 14:58:10.086907+05	\N
9dca9955-b2eb-4d2c-b4a9-ddbcfd251f56	02bd4413-8586-49ab-802e-16304e756a8b	81b84c5d-9759-4b86-978a-649c8ef79660	2022-09-17 14:58:10.098298+05	2022-09-17 14:58:10.098298+05	\N
80623747-b1b4-49d6-ba3a-eb7703a830c6	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	660071e0-8f17-4c48-9d80-d4cac306de3a	2022-09-17 14:58:40.166259+05	2022-09-17 14:58:40.166259+05	\N
430bff2b-7074-4703-8496-9d3124ba45eb	02bd4413-8586-49ab-802e-16304e756a8b	660071e0-8f17-4c48-9d80-d4cac306de3a	2022-09-17 14:58:40.177106+05	2022-09-17 14:58:40.177106+05	\N
68653c84-8321-4ccf-b180-4fe3a49f2d43	5bb9a4e7-9992-418f-b551-537844d371da	660071e0-8f17-4c48-9d80-d4cac306de3a	2022-09-17 14:58:40.189647+05	2022-09-17 14:58:40.189647+05	\N
7f22a257-7ff9-412e-bde2-8a1246454cb4	02bd4413-8586-49ab-802e-16304e756a8b	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	2022-09-17 14:59:14.122334+05	2022-09-17 14:59:14.122334+05	\N
2339b2f3-c284-4b6d-8196-7c080904a9c6	5bb9a4e7-9992-418f-b551-537844d371da	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	2022-09-17 14:59:14.133375+05	2022-09-17 14:59:14.133375+05	\N
b4afd21c-f52e-4ad1-acb7-74ae96fc9b57	b982bd86-0a0f-4950-baad-5a131e9b728e	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	2022-09-17 14:59:14.143787+05	2022-09-17 14:59:14.143787+05	\N
63415e35-ff9b-47dd-9e54-934380382a4f	5bb9a4e7-9992-418f-b551-537844d371da	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	2022-09-17 14:59:44.92297+05	2022-09-17 14:59:44.92297+05	\N
04d1271e-8951-4d4e-83c5-7d4111bda28e	b982bd86-0a0f-4950-baad-5a131e9b728e	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	2022-09-17 14:59:44.933209+05	2022-09-17 14:59:44.933209+05	\N
d8914b36-6023-448b-a0e4-e14cbe63f5ac	f745d171-68e6-42e2-b339-cb3c210cda55	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	2022-09-17 14:59:44.943874+05	2022-09-17 14:59:44.943874+05	\N
0b5defe3-62ca-40a2-955e-f186820d8f2c	b982bd86-0a0f-4950-baad-5a131e9b728e	8df705a5-2351-4aca-b03e-3357a23840b4	2022-09-17 15:00:15.235939+05	2022-09-17 15:00:15.235939+05	\N
e25fe89a-4c5a-4611-be9e-0007a5c631ff	f745d171-68e6-42e2-b339-cb3c210cda55	8df705a5-2351-4aca-b03e-3357a23840b4	2022-09-17 15:00:15.244281+05	2022-09-17 15:00:15.244281+05	\N
d9e3f447-0a02-4788-8ecb-ef02d266db02	d4cb1359-6c23-4194-8e3c-21ed8cec8373	8df705a5-2351-4aca-b03e-3357a23840b4	2022-09-17 15:00:15.256055+05	2022-09-17 15:00:15.256055+05	\N
c55cc88b-5651-4fa2-91fd-c85c53b2fed7	b982bd86-0a0f-4950-baad-5a131e9b728e	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	2022-10-06 11:07:41.731753+05	2022-10-06 11:07:41.731753+05	\N
9eff8362-8c57-4b7a-a8bd-695cf30d3dcc	f745d171-68e6-42e2-b339-cb3c210cda55	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	2022-10-06 11:07:41.792463+05	2022-10-06 11:07:41.792463+05	\N
d90750b0-0672-4805-a8fc-9fe46f3ff7b4	d4cb1359-6c23-4194-8e3c-21ed8cec8373	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	2022-10-06 11:07:41.802468+05	2022-10-06 11:07:41.802468+05	\N
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
96c30e15-274c-49a0-bcc5-e2f8deac248f	+993 12 227475	2022-09-29 12:49:06.569246+05	2022-09-29 12:49:06.569246+05	\N
\.


--
-- Data for Name: company_setting; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_setting (id, logo, favicon, email, instagram, created_at, updated_at, deleted_at) FROM stdin;
7d193677-e0b1-4df0-be88-dc6e16a47ca7	uploads/logode9c4f45-acba-42ce-b435-e744631a98ba.jpeg	uploads/favicon8a413c02-108d-4d2f-8e92-d24a18cea1d3.jpeg	isleg-bazar@gmail.com	@islegbazarinstagram	2022-06-15 19:57:04.54457+05	2022-06-15 19:57:04.54457+05	\N
\.


--
-- Data for Name: customer_address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customer_address (id, customer_id, address, created_at, updated_at, deleted_at, is_active) FROM stdin;
dad2aaae-1b17-4494-8d85-8a8ee6e3e60f	7e872c52-0d23-4086-8c45-43000b57332e	Mir 2/2 jay 7 oy 36	2022-09-21 21:28:29.708359+05	2022-09-21 21:28:29.708359+05	\N	f
3c4e1c4f-fd51-4dd5-befe-7051b79312a6	89a6ac71-4495-4218-b9f9-3f2a3eab085b	Mir 2/2 jay 7 oy 36	2022-09-22 14:27:08.806836+05	2022-09-22 14:27:08.806836+05	\N	t
ad9422b5-c496-4791-b1b4-7454bc10aefd	89a6ac71-4495-4218-b9f9-3f2a3eab085b	Mir 2/2 jay 7 oy 36	2022-09-22 12:54:24.177072+05	2022-09-22 12:54:24.177072+05	\N	f
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (id, full_name, phone_number, password, birthday, gender, created_at, updated_at, deleted_at, email, is_register) FROM stdin;
7e872c52-0d23-4086-8c45-43000b57332e	Serdar	+99363747155	$2a$14$N11BpwBYOI72mX9nBjMNL.e0.iwFBndE2efV3Nqx0/fHj3OzNZSlq	\N	\N	2022-09-22 12:46:26.731782+05	2022-09-22 12:46:26.731782+05	\N	serdar@gmail.com	t
7fafe6f8-c6b6-4bcc-9063-e98c113902c5	Mahri Wepayewa	+99363747156	$2a$14$N11BpwBYOI72mX9nBjMNL.e0.iwFBndE2efV3Nqx0/fHj3OzNZSlq	\N	\N	2022-09-22 12:46:26.731782+05	2022-09-22 12:46:26.731782+05	\N	mahri@gmail.com	t
38615c8c-1af5-424f-b7a3-071d38c42b86	Aylar Siriyewa	+99363234587	$2a$14$N11BpwBYOI72mX9nBjMNL.e0.iwFBndE2efV3Nqx0/fHj3OzNZSlq	\N	\N	2022-09-22 12:46:26.731782+05	2022-09-22 12:46:26.731782+05	\N	aylar@gmail.com	t
9b1a0831-9943-4aa9-aa2a-3507743a5de4	Rahmet	+99361235698	$2a$14$N11BpwBYOI72mX9nBjMNL.e0.iwFBndE2efV3Nqx0/fHj3OzNZSlq	\N	\N	2022-09-22 12:46:26.731782+05	2022-09-22 12:46:26.731782+05	\N	rahmet@gmail.com	t
eb4d03d3-c201-49e6-867e-a7b6927a414c	Wepa Maksadow	+99363658741	$2a$14$N11BpwBYOI72mX9nBjMNL.e0.iwFBndE2efV3Nqx0/fHj3OzNZSlq	\N	\N	2022-09-22 12:46:26.731782+05	2022-09-22 12:46:26.731782+05	\N	wepa@gmail.com	t
4406f560-b979-4e7a-a296-bad88b20d731	wedkwekfjewf	+99363787878	$2a$14$.przN91vmxTSncM0mhWxNexs1U2Nb9XrpfzBfRTZT0QKAw1DLFNiu	\N	\N	2022-10-12 01:52:24.749747+05	2022-10-12 01:52:24.749747+05	\N	ewkfnewj@gmail.com	t
8409206b-c46a-4ac6-a6fd-285aac8c53c7	wfjknwkejfwe	+99367474747	$2a$14$NagV1Uq8YOSjdSjkd1QIFOPbByfZgwLp9Obc3zsRahyjulxb25xNq	\N	\N	2022-10-12 01:55:36.31034+05	2022-10-12 01:55:36.31034+05	\N	ewdknewnewjfnej@gmail.com	t
8a34b75e-2f2a-4987-85d3-87e98e7f6733	lkruiq34	+99367777777	$2a$14$5k0/YzSxmGu/6PpwxL3xuOnUblRx60KWZ6ogj6oCPT1ao13X595NC	\N	\N	2022-10-12 02:15:28.732008+05	2022-10-12 02:15:28.732008+05	\N	wdkneqj@gmail.com	t
e65f3e71-eecf-4463-8528-2c1ad5dce6df	ewdjnewjnew	+99366262626	$2a$14$JHETMGowAtRdurf3GKu5YOYExT6DCF/b12pyk8lv7yJ4cW6dHx/Ia	\N	\N	2022-10-12 02:19:40.911888+05	2022-10-12 02:19:40.911888+05	\N	wedkwedewejnew@gmail.com	t
1287f95b-fef2-4796-ace5-87465ee8efc7	ewdjnewjnew	+99366262627	$2a$14$4wpPRXEPMY1zk/rLnYhjT.AY7/4.CJTUwXBu5/aWSvMcz0iWer5VS	\N	\N	2022-10-12 02:20:37.457297+05	2022-10-12 02:20:37.457297+05	\N	wedkwedewGHJejnew@gmail.com	t
c6916c0b-8756-43b6-b51d-46500ce04779	kmlkadnead	+99368888888	$2a$14$tEI47M9qymtVD3QSpwuFYOVtA8NIk.6DS8JdD4GtZjaKxupwrkSLm	\N	\N	2022-10-12 02:23:15.77006+05	2022-10-12 02:23:15.77006+05	\N	adkedn@gmail.com	t
5b2b52d9-922a-4731-b3c9-6322695e6908	ewdkwede	+99364545454	$2a$14$CbXcYEFy1d8gen.9eSkiGuR5S3FnD3R8NjgBXCAgWQUDAMK4lEN12	\N	\N	2022-10-12 02:26:12.740487+05	2022-10-12 02:26:12.740487+05	\N	wedkwend@gmail.com	t
7e3431f2-e64e-4081-99c0-861958e2e5fd	ewdkwenf	+99364548789	$2a$14$rW.UsqkclC8qkB4A9a2SL.y6sfLTiAt1EijE0gi2BIgUQEZ/bKltS	\N	\N	2022-10-12 02:28:07.856352+05	2022-10-12 02:28:07.856352+05	\N	ewfklwfelnk@gmail.com	t
856a6c98-8c39-43d6-96b5-3870720c500c	ddedeededmwe	+99364848484	$2a$14$jm01PfnaGBm7v0vnGqOGhO9SOY8fJ.rFxFBeC.8PMPVMqD5dIbxtm	\N	\N	2022-10-12 02:30:50.540961+05	2022-10-12 02:30:50.540961+05	\N	ekwenknef@gmail.com	t
fa3615fb-a0a1-4e82-b4fa-c3dbb48238f5	edkwenio;q43;	+99361234568	$2a$14$SQs0oA45Ogq6xY3BTLGcM.dQN5AqjRt.G5gKF.wm/2CIZL/9f/iqq	\N	\N	2022-10-12 02:35:01.711127+05	2022-10-12 02:35:01.711127+05	\N	ewdmewdjk@gmail.com	t
89a6ac71-4495-4218-b9f9-3f2a3eab085b	Allanur Bayramgeldiyew	+99362420377	$2a$14$itnGwI2Y1ZluXXOsYtqQDOKsc9wbc.ID9cjrDQ4O4en.ss0ALgO2y	1998-06-24	\N	2022-09-22 12:54:39.635618+05	2022-09-22 12:54:39.635618+05	\N	abb@gmail.com	t
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

COPY public.images (id, product_id, small, large, created_at, updated_at, deleted_at) FROM stdin;
b59f96d5-fa50-47cf-873a-dbe1a6e15302	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uploads/product/df8b0be2-7147-4ae6-8ba4-176c90aa817b.jpg	uploads/product/adca811e-07ce-4a65-9514-41a7dead8fc3.jpg	2022-09-17 14:54:57.029637+05	2022-09-17 14:54:57.029637+05	\N
2362835c-9c11-49fe-861b-81c3f29c767f	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uploads/product/0ea9466b-fd9b-4292-b3c9-0104acbaed0c.jpg	uploads/product/81d1781c-4701-40ea-91c4-4609b05de9e7.jpg	2022-09-17 14:54:57.040741+05	2022-09-17 14:54:57.040741+05	\N
ae50cccd-3a73-45ce-8fc8-b95e0f0b91bc	b2b165a3-2261-4d67-8160-0e239ecd99b5	uploads/product/a3cceed2-4af8-4017-9099-bf08231d921d.jpg	uploads/product/2d006bc4-8018-497e-9f4f-0a908fc11eaf.jpg	2022-09-17 14:55:35.496799+05	2022-09-17 14:55:35.496799+05	\N
34e6d0f9-97fc-4590-a59f-5fc2107ccdbf	b2b165a3-2261-4d67-8160-0e239ecd99b5	uploads/product/83b2761e-d5a7-44b0-886d-a2ac7ce1aeae.jpg	uploads/product/de2d464f-b447-4fe5-b33e-8db764ee90ce.jpg	2022-09-17 14:55:35.507566+05	2022-09-17 14:55:35.507566+05	\N
ff4883cb-6f14-4c18-b080-56876b24e0a9	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uploads/product/f18dea7c-1df0-48bc-a887-eb7b3b3b6f41.jpg	uploads/product/48cdeb99-d3b0-4bad-9365-01bdbcb9e88a.jpg	2022-09-17 14:56:05.441685+05	2022-09-17 14:56:05.441685+05	\N
06073ec4-b58e-4fba-877d-6aae1339d084	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uploads/product/3e6a4832-663f-4b21-b3e4-f637245590a2.jpg	uploads/product/b37e6ed9-401c-465f-9d50-f990b74a9e93.jpg	2022-09-17 14:56:05.452041+05	2022-09-17 14:56:05.452041+05	\N
b2964c6f-e49e-4d55-833b-7146dbe8e9f0	d731b17a-ae8d-4561-ad67-0f431d5c529b	uploads/product/fd6516a1-d696-4fcf-a5ee-89c6f5b48cb5.jpg	uploads/product/af06206e-e7de-438b-b135-70a9348eb873.jpg	2022-09-17 14:56:36.186395+05	2022-09-17 14:56:36.186395+05	\N
0f72550f-4127-421f-bf4d-83e7d5d0c667	d731b17a-ae8d-4561-ad67-0f431d5c529b	uploads/product/c148beac-f6f2-43c8-b4c4-8ddae87cf1b8.jpg	uploads/product/2975a2c4-f978-40fa-9798-b30a71c8c908.jpg	2022-09-17 14:56:36.197707+05	2022-09-17 14:56:36.197707+05	\N
541a60b2-e164-45f1-8238-b430f9a4f696	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uploads/product/f9484719-a04c-482a-b9cb-d4477d67171e.jpg	uploads/product/56d9c5d4-89ac-45e7-80a5-d4420a1ba1e4.jpg	2022-09-17 14:57:07.232409+05	2022-09-17 14:57:07.232409+05	\N
01234b03-8ebe-4bd9-9296-5411d8b602d6	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uploads/product/f8a3e4a8-1363-49de-95d8-7ea99578bb36.jpg	uploads/product/7d443731-50ee-4e9b-9552-562282763df1.jpg	2022-09-17 14:57:07.242703+05	2022-09-17 14:57:07.242703+05	\N
f8c2dd64-75f1-4617-8237-22aa6fde2faa	d4156225-082e-4f0f-9b2c-85268114433a	uploads/product/fe06f4f8-13da-49d5-aeec-4a05bd9155c4.jpg	uploads/product/5db94dda-c1f5-4d99-a0f1-a005be127383.jpg	2022-09-17 14:57:34.620044+05	2022-09-17 14:57:34.620044+05	\N
394d0cde-5f52-4228-a615-574acec7e51a	d4156225-082e-4f0f-9b2c-85268114433a	uploads/product/c29e7734-e8cd-453f-8583-74f77b76f425.jpg	uploads/product/ac6306f5-0193-413f-9838-7daed09d7fb5.jpg	2022-09-17 14:57:34.631937+05	2022-09-17 14:57:34.631937+05	\N
b1575796-5322-4983-ae21-bb84725a4f75	81b84c5d-9759-4b86-978a-649c8ef79660	uploads/product/1ff26530-a022-4a3f-90e9-8c4a3a2d5f2e.jpg	uploads/product/10b3062a-71cc-49b3-9eae-4410745a7685.jpg	2022-09-17 14:58:10.031274+05	2022-09-17 14:58:10.031274+05	\N
21f8e93a-ea3a-4596-899a-1982d8956dcb	81b84c5d-9759-4b86-978a-649c8ef79660	uploads/product/3bebc440-4f7f-4c34-abd1-c7151f41823d.jpg	uploads/product/1cfabb12-a0ce-4a11-b937-149f03fc95c2.jpg	2022-09-17 14:58:10.042467+05	2022-09-17 14:58:10.042467+05	\N
843515aa-4fef-4ef3-9570-61bf600a9357	660071e0-8f17-4c48-9d80-d4cac306de3a	uploads/product/d629fe21-e9a8-4c39-941d-ea222b0ce204.jpg	uploads/product/2e752914-f145-4f09-a35d-272db51b3083.jpg	2022-09-17 14:58:40.120625+05	2022-09-17 14:58:40.120625+05	\N
2957da00-07b2-4f26-9504-44f1c50da3e8	660071e0-8f17-4c48-9d80-d4cac306de3a	uploads/product/0b9ffad0-fcd9-4a36-b661-15bcff855593.jpg	uploads/product/fc1d5b8d-bd84-411e-af20-b5a306966579.jpg	2022-09-17 14:58:40.131788+05	2022-09-17 14:58:40.131788+05	\N
e6c8f3db-35cf-404b-a3e9-cc71b7976292	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uploads/product/40330079-629d-4146-a2b3-366e38cb2763.jpg	uploads/product/3344dd28-0808-4c4c-a3ed-ef877924f498.jpg	2022-09-17 14:59:14.077892+05	2022-09-17 14:59:14.077892+05	\N
41835ad5-e5ba-4275-8848-b4f544ade2cf	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uploads/product/50b21aa9-019c-46e5-8d36-ce79f596a5bd.jpg	uploads/product/9f1c4c2a-57f6-4f96-ad9c-8245d4773c88.jpg	2022-09-17 14:59:14.08807+05	2022-09-17 14:59:14.08807+05	\N
fa73d216-36fe-4086-84b7-334ff8945658	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	uploads/product/980c3753-a72c-4515-b29f-a0713c87b224.jpg	uploads/product/da6d93fe-15d4-4dd9-a1ea-541911e6a843.jpg	2022-09-17 14:59:44.877184+05	2022-09-17 14:59:44.877184+05	\N
f47e3ed5-7dc7-4a08-a560-b15401f13b9e	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	uploads/product/6e168766-4d15-4bb2-8fdb-85e04cd4df6f.jpg	uploads/product/94ac9355-777b-4df2-8df1-fa8e5d0f1cda.jpg	2022-09-17 14:59:44.888846+05	2022-09-17 14:59:44.888846+05	\N
d96e186f-8d34-447c-814e-61ef93873f84	8df705a5-2351-4aca-b03e-3357a23840b4	uploads/product/7527c80b-aad3-440c-b605-1dc1f9b856c1.jpg	uploads/product/205684f5-bbeb-47c2-912c-4ad689674547.jpg	2022-09-17 15:00:15.188668+05	2022-09-17 15:00:15.188668+05	\N
81a68f8d-31c4-4b6b-ab15-72aea698abc3	8df705a5-2351-4aca-b03e-3357a23840b4	uploads/product/d2d8dd0c-b452-4566-b98e-a9ffab10fb13.jpg	uploads/product/a940f01a-5e38-43b6-99dc-a742e388ffe5.jpg	2022-09-17 15:00:15.200605+05	2022-09-17 15:00:15.200605+05	\N
83279012-42f0-4367-a44a-6bd59ddb249c	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uploads/product/1d54f24c-599d-46a6-ab14-c2bb565cc059.jpg	uploads/product/5fda4077-e81d-4e72-b899-f33103512859.jpg	2022-10-06 11:07:41.569383+05	2022-10-06 11:07:41.569383+05	\N
68f44447-1c5f-457c-a47a-2a281ea9c0dd	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uploads/product/8db0d212-5797-47a9-97e5-f08ceb7a293f.jpg	uploads/product/042697cc-588d-4286-a226-a94992fe942d.jpg	2022-10-06 11:07:41.646964+05	2022-10-06 11:07:41.646964+05	\N
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
05d3e7ae-d497-4cfd-8b6b-c6b334007f9e	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	7e872c52-0d23-4086-8c45-43000b57332e	2022-10-12 01:47:35.545237+05	2022-10-12 01:47:35.545237+05	\N
bd6ac2c4-1f1c-4e50-b092-6d1332178eda	b2b165a3-2261-4d67-8160-0e239ecd99b5	7e872c52-0d23-4086-8c45-43000b57332e	2022-10-12 01:47:35.72356+05	2022-10-12 01:47:35.72356+05	\N
79713438-920f-4e43-927e-d5be7efb1d4a	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	7e872c52-0d23-4086-8c45-43000b57332e	2022-10-12 01:55:36.634327+05	2022-10-12 01:55:36.634327+05	\N
\.


--
-- Data for Name: main_image; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.main_image (id, product_id, small, medium, large, created_at, updated_at, deleted_at) FROM stdin;
af383593-cacb-4440-8144-4560c1887921	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uploads/product/0bb0fc0e-6b9a-48cb-a005-067818339a9a.jpg	uploads/product/a07bf6d8-837d-4743-9752-3a5a2b178b4d.jpg	uploads/product/3d48f030-4edb-4510-af2f-4688667f4e0b.jpg	2022-09-17 14:54:57.018537+05	2022-09-17 14:54:57.018537+05	\N
045bb7ae-6d64-4366-acdd-7f342a7600a2	b2b165a3-2261-4d67-8160-0e239ecd99b5	uploads/product/a7c19215-5db7-4ebb-b373-a8b58b678def.jpg	uploads/product/58fae77e-a7d4-4f4b-bdfd-399e1e7994e8.jpg	uploads/product/22c2eb3b-cf3e-4369-b08f-92a6f8665bdd.jpg	2022-09-17 14:55:35.485706+05	2022-09-17 14:55:35.485706+05	\N
85d25308-e2a0-40f7-ab57-4dd081e59ed8	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uploads/product/061bb6bc-d603-4a28-8737-246018133728.jpg	uploads/product/1e562ca8-31bf-4712-af03-feb1654bf1f1.jpg	uploads/product/cba805fa-8c99-4992-9315-60071832bb59.jpg	2022-09-17 14:56:05.43168+05	2022-09-17 14:56:05.43168+05	\N
0b9cc77d-87fd-4603-bf2d-e7203adeb4e8	d731b17a-ae8d-4561-ad67-0f431d5c529b	uploads/product/a2d59c2f-e312-4541-b4a7-79fc679c5e0c.jpg	uploads/product/87aa78b0-6f8a-4b37-a7f6-790e77398e8b.jpg	uploads/product/5d7e8980-f050-4936-9198-eaa3f3236370.jpg	2022-09-17 14:56:36.174972+05	2022-09-17 14:56:36.174972+05	\N
7ab33a3b-0195-4024-a035-e54268762d3b	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uploads/product/bf635780-dc34-4478-b936-8e073eefa79e.jpg	uploads/product/ea2ce473-4110-4129-8d5e-0fa456d10530.jpg	uploads/product/a92d33a6-8180-4355-8f4b-bd23079177a4.jpg	2022-09-17 14:57:07.221569+05	2022-09-17 14:57:07.221569+05	\N
7982b9f8-bfca-42f5-b13a-37b1ed41e1a2	d4156225-082e-4f0f-9b2c-85268114433a	uploads/product/706c3355-01f9-4648-850c-eacc3557f435.jpg	uploads/product/e585a68a-d6c3-4a1f-8660-041cade9899e.jpg	uploads/product/03902970-8446-4180-9200-902a6aa7fa23.jpg	2022-09-17 14:57:34.609435+05	2022-09-17 14:57:34.609435+05	\N
286f3f21-8750-499f-b35c-9de67b236316	81b84c5d-9759-4b86-978a-649c8ef79660	uploads/product/bc78c4bb-35f7-4873-88eb-af1c985e9f34.jpg	uploads/product/356b93e2-be40-4ed0-bf77-b229210de3e7.jpg	uploads/product/d7641dc6-6e2a-4d21-9403-98c872f93b25.jpg	2022-09-17 14:58:10.020986+05	2022-09-17 14:58:10.020986+05	\N
9ddabe0c-bad0-493e-b084-a5d6d96e894d	660071e0-8f17-4c48-9d80-d4cac306de3a	uploads/product/64e1692d-bdd3-4081-88d5-74c8cc71ae51.jpg	uploads/product/74278164-f5c5-423f-8d84-ba7a122a8171.jpg	uploads/product/018a0b4a-8593-4287-93b3-37aaa1a04f0f.jpg	2022-09-17 14:58:40.111413+05	2022-09-17 14:58:40.111413+05	\N
489304cb-a16a-4f78-841e-797b341f224b	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uploads/product/139044f5-bb94-4f31-9b5a-c7c28056398f.jpg	uploads/product/a45a8c58-bc17-44ad-807d-719607bdd031.jpg	uploads/product/7d873556-bfe5-4991-8cc7-0ab6609eb45e.jpg	2022-09-17 14:59:14.066244+05	2022-09-17 14:59:14.066244+05	\N
3141d941-c1fa-41f0-b542-44090d4ba2a1	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	uploads/product/ad41dd13-4b3b-442a-92a9-ea7ba39a8ffd.jpg	uploads/product/4b536ae1-bb8a-43f3-bf33-378fb2e53f58.jpg	uploads/product/e50d6762-01dd-4894-a991-0f27bb401630.jpg	2022-09-17 14:59:44.866884+05	2022-09-17 14:59:44.866884+05	\N
24fdc6b1-afb2-4735-b406-70addc0dd8d9	8df705a5-2351-4aca-b03e-3357a23840b4	uploads/product/c0a88fd0-2374-49af-95d0-5c692c626b94.jpg	uploads/product/8a67c3a7-94fc-4831-9f90-80ae409c684f.jpg	uploads/product/c7a5b0ce-d9fc-449c-b039-726d716c62a7.jpg	2022-09-17 15:00:15.178822+05	2022-09-17 15:00:15.178822+05	\N
ea8b26f1-1e34-4587-9b72-3747e29930cf	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uploads/product/5938e42c-700a-48c3-8461-ad3df445dac1.jpg	uploads/product/9ac3b3f6-6d4b-4be0-8a98-4d1b10f4d954.jpg	uploads/product/5516cc1d-8f21-4006-b353-b41518ef1bef.jpg	2022-10-06 11:07:41.522494+05	2022-10-06 11:07:41.522494+05	\N
\.


--
-- Data for Name: order_dates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_dates (id, date, created_at, updated_at, deleted_at) FROM stdin;
32646376-c93f-412b-9e75-b3a5fa70df9e	today	2022-09-28 17:35:33.772335+05	2022-09-28 17:35:33.772335+05	\N
c1f2beca-a6b6-4971-a6a7-ed50079c6912	tomorrow	2022-09-28 17:36:46.804343+05	2022-09-28 17:36:46.804343+05	\N
\.


--
-- Data for Name: order_times; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_times (id, order_date_id, "time", created_at, updated_at, deleted_at) FROM stdin;
9c2261db-a8d3-4c7d-8bcb-302ac1f1f9fb	32646376-c93f-412b-9e75-b3a5fa70df9e	09:00 - 12:00	2022-09-28 17:35:33.802847+05	2022-09-28 17:35:33.802847+05	\N
fabed2a7-f467-4ef5-846f-73c0384755b8	32646376-c93f-412b-9e75-b3a5fa70df9e	18:00 - 21:00	2022-09-28 17:35:33.802847+05	2022-09-28 17:35:33.802847+05	\N
7d47a77a-b8f3-4e96-aa56-5ec7fb328e86	c1f2beca-a6b6-4971-a6a7-ed50079c6912	09:00 - 12:00	2022-09-28 17:36:46.825964+05	2022-09-28 17:36:46.825964+05	\N
de31361b-9fba-48f2-9341-9e3dd08cf9fd	c1f2beca-a6b6-4971-a6a7-ed50079c6912	18:00 - 21:00	2022-09-28 17:36:46.825964+05	2022-09-28 17:36:46.825964+05	\N
\.


--
-- Data for Name: ordered_products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ordered_products (id, product_id, quantity_of_product, order_id, created_at, updated_at, deleted_at) FROM stdin;
777caf46-9fb8-46a8-808a-cd285ea11340	b2b165a3-2261-4d67-8160-0e239ecd99b5	12	3d29c94b-1869-4d4c-a94f-180a5c9eb614	2022-09-21 21:28:30.413438+05	2022-09-21 21:28:30.413438+05	\N
e209c082-ae64-44c1-b235-9245a07cb79d	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	4	3d29c94b-1869-4d4c-a94f-180a5c9eb614	2022-09-21 21:28:30.437102+05	2022-09-21 21:28:30.437102+05	\N
3c25a223-2df3-4f47-a8c5-ad70b67522d1	b2b165a3-2261-4d67-8160-0e239ecd99b5	12	ef468ad0-747b-4c12-a195-c72b9f19dc84	2022-09-22 12:54:24.210828+05	2022-09-22 12:54:24.210828+05	\N
8d1f6c6c-6e9a-4761-9791-7f41fb37347f	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	4	ef468ad0-747b-4c12-a195-c72b9f19dc84	2022-09-22 12:54:24.22149+05	2022-09-22 12:54:24.22149+05	\N
98ee5e13-b07c-4535-9b5b-8651eadd8795	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	25	e9b43b21-2947-4c41-95e6-1b81d925996b	2022-09-22 14:27:08.845012+05	2022-09-22 14:27:08.845012+05	\N
17685559-6510-4f87-92dc-8e3e5ad339e2	8df705a5-2351-4aca-b03e-3357a23840b4	3	e9b43b21-2947-4c41-95e6-1b81d925996b	2022-09-22 14:27:08.855567+05	2022-09-22 14:27:08.855567+05	\N
f4f79932-6d84-43d7-a4dd-aa6a64588f48	d4156225-082e-4f0f-9b2c-85268114433a	25	9f1f4526-7ee7-4347-96fc-6fb97c1b402a	2022-09-22 14:45:46.407399+05	2022-09-22 14:45:46.407399+05	\N
982165f4-c3fa-4922-9d42-9d28b138b438	81b84c5d-9759-4b86-978a-649c8ef79660	3	9f1f4526-7ee7-4347-96fc-6fb97c1b402a	2022-09-22 14:45:46.419832+05	2022-09-22 14:45:46.419832+05	\N
336647b9-abd9-4a2a-a145-9e4326bdcc76	d4156225-082e-4f0f-9b2c-85268114433a	25	7728fdeb-e871-414d-b5fd-d57ee59e879c	2022-09-29 14:05:14.80837+05	2022-09-29 14:05:14.80837+05	\N
c8a62bf5-9dfc-4aef-bdfe-8b17913d770c	81b84c5d-9759-4b86-978a-649c8ef79660	3	7728fdeb-e871-414d-b5fd-d57ee59e879c	2022-09-29 14:05:14.853259+05	2022-09-29 14:05:14.853259+05	\N
7929584a-340c-4ec1-bb0c-8c7d9f858559	d4156225-082e-4f0f-9b2c-85268114433a	25	8930edc5-7b9e-4cef-ac58-97bc8261df3a	2022-09-29 14:06:29.698335+05	2022-09-29 14:06:29.698335+05	\N
93bcf413-e45b-4903-ac39-a444b627d325	81b84c5d-9759-4b86-978a-649c8ef79660	3	8930edc5-7b9e-4cef-ac58-97bc8261df3a	2022-09-29 14:06:29.710636+05	2022-09-29 14:06:29.710636+05	\N
af266256-496c-4f52-849f-ff3471ed3afd	d4156225-082e-4f0f-9b2c-85268114433a	25	3a8f977f-2eb9-4a4a-9bb1-51cd9a5a4fe6	2022-09-29 14:06:57.843088+05	2022-09-29 14:06:57.843088+05	\N
c2b3ead0-63ff-4161-b16b-f6cded8b8a88	81b84c5d-9759-4b86-978a-649c8ef79660	3	3a8f977f-2eb9-4a4a-9bb1-51cd9a5a4fe6	2022-09-29 14:06:57.854411+05	2022-09-29 14:06:57.854411+05	\N
4d7aca2c-6b2b-49bc-a86b-949f12dfa59a	d4156225-082e-4f0f-9b2c-85268114433a	25	7b0af658-e230-405d-a1ed-718f52973eb2	2022-09-29 14:08:09.999618+05	2022-09-29 14:08:09.999618+05	\N
4ae25fa7-04df-43d1-82bd-7f7b91df4252	81b84c5d-9759-4b86-978a-649c8ef79660	3	7b0af658-e230-405d-a1ed-718f52973eb2	2022-09-29 14:08:10.011834+05	2022-09-29 14:08:10.011834+05	\N
97cfde5b-1e67-486e-b81a-9818bb5a1e6e	d4156225-082e-4f0f-9b2c-85268114433a	25	68c886f4-c005-4a22-9219-f643577d0a62	2022-09-29 14:08:56.71182+05	2022-09-29 14:08:56.71182+05	\N
155ce965-827a-48cc-8b79-1fea6b971919	81b84c5d-9759-4b86-978a-649c8ef79660	3	68c886f4-c005-4a22-9219-f643577d0a62	2022-09-29 14:08:56.723832+05	2022-09-29 14:08:56.723832+05	\N
fca009f3-642b-45d0-b4ba-c8b5e0836a99	d4156225-082e-4f0f-9b2c-85268114433a	25	c09ef7ff-4244-422a-b2da-3e160e39f0c8	2022-09-30 16:17:12.503388+05	2022-09-30 16:17:12.503388+05	\N
b0a20b48-5bf6-4324-9c48-71a09b92159d	81b84c5d-9759-4b86-978a-649c8ef79660	3	c09ef7ff-4244-422a-b2da-3e160e39f0c8	2022-09-30 16:17:12.562328+05	2022-09-30 16:17:12.562328+05	\N
3187dd81-6db8-45b2-85d3-266b17e1bca2	8df705a5-2351-4aca-b03e-3357a23840b4	3	c09ef7ff-4244-422a-b2da-3e160e39f0c8	2022-09-30 16:17:12.57442+05	2022-09-30 16:17:12.57442+05	\N
4005e7c3-2aee-4428-bc33-a144904bcffd	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	c09ef7ff-4244-422a-b2da-3e160e39f0c8	2022-09-30 16:17:12.596612+05	2022-09-30 16:17:12.596612+05	\N
6248a3ef-5bbe-423e-bdeb-51358cf79230	d4156225-082e-4f0f-9b2c-85268114433a	25	1499d084-b010-4724-ac40-2451a87986d9	2022-09-30 16:19:46.509965+05	2022-09-30 16:19:46.509965+05	\N
5e311871-6462-4091-afd4-4b6178149250	81b84c5d-9759-4b86-978a-649c8ef79660	3	1499d084-b010-4724-ac40-2451a87986d9	2022-09-30 16:19:46.520869+05	2022-09-30 16:19:46.520869+05	\N
345fefb6-c0b3-4553-8ce7-80633386855e	8df705a5-2351-4aca-b03e-3357a23840b4	3	1499d084-b010-4724-ac40-2451a87986d9	2022-09-30 16:19:46.532044+05	2022-09-30 16:19:46.532044+05	\N
24e41b79-c337-40e0-b04b-1dc2e73b24d0	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	1499d084-b010-4724-ac40-2451a87986d9	2022-09-30 16:19:46.542693+05	2022-09-30 16:19:46.542693+05	\N
4d7663a3-d7ed-45ff-9eca-1272ff093582	d4156225-082e-4f0f-9b2c-85268114433a	25	4271c7d2-eeff-4948-989f-ae5e4518854f	2022-09-30 16:22:21.368957+05	2022-09-30 16:22:21.368957+05	\N
63fd2f2e-7178-44d7-894a-1a93c6cc49d6	81b84c5d-9759-4b86-978a-649c8ef79660	3	4271c7d2-eeff-4948-989f-ae5e4518854f	2022-09-30 16:22:21.379357+05	2022-09-30 16:22:21.379357+05	\N
41470fb9-f636-4320-ba3a-d590f56ae823	8df705a5-2351-4aca-b03e-3357a23840b4	3	4271c7d2-eeff-4948-989f-ae5e4518854f	2022-09-30 16:22:21.390021+05	2022-09-30 16:22:21.390021+05	\N
6b7d5483-a5db-4595-840d-9c229ba8c827	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	4271c7d2-eeff-4948-989f-ae5e4518854f	2022-09-30 16:22:21.402525+05	2022-09-30 16:22:21.402525+05	\N
41c9f8c2-5586-437f-ae42-1201935680de	d4156225-082e-4f0f-9b2c-85268114433a	25	6fda5178-9944-4ffe-92e7-319f41323818	2022-09-30 16:28:44.074707+05	2022-09-30 16:28:44.074707+05	\N
e887a11d-5168-48de-b76d-f071a93157eb	81b84c5d-9759-4b86-978a-649c8ef79660	3	6fda5178-9944-4ffe-92e7-319f41323818	2022-09-30 16:28:44.08603+05	2022-09-30 16:28:44.08603+05	\N
3fff0a4c-e4c6-48d6-8be1-759c9df13268	8df705a5-2351-4aca-b03e-3357a23840b4	3	6fda5178-9944-4ffe-92e7-319f41323818	2022-09-30 16:28:44.097244+05	2022-09-30 16:28:44.097244+05	\N
10471abf-17a6-4350-a358-fccb1a3bcab3	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	6fda5178-9944-4ffe-92e7-319f41323818	2022-09-30 16:28:44.107886+05	2022-09-30 16:28:44.107886+05	\N
c7e09b89-e593-4bff-a3d9-f1724e22d7aa	d4156225-082e-4f0f-9b2c-85268114433a	25	67f2638b-971c-48a9-bada-5d14b50d74bc	2022-09-30 16:29:07.397594+05	2022-09-30 16:29:07.397594+05	\N
7965e167-3b36-422e-ade1-ce0cfb9c20f3	81b84c5d-9759-4b86-978a-649c8ef79660	3	67f2638b-971c-48a9-bada-5d14b50d74bc	2022-09-30 16:29:07.409019+05	2022-09-30 16:29:07.409019+05	\N
2630dd99-7408-4e39-a230-16c25dd4c18f	8df705a5-2351-4aca-b03e-3357a23840b4	3	67f2638b-971c-48a9-bada-5d14b50d74bc	2022-09-30 16:29:07.420871+05	2022-09-30 16:29:07.420871+05	\N
c36fd7f3-9e11-4df2-8147-592a9b555673	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	67f2638b-971c-48a9-bada-5d14b50d74bc	2022-09-30 16:29:07.431699+05	2022-09-30 16:29:07.431699+05	\N
8ee02a8d-ae7a-4490-a92e-7a09b1d87333	d4156225-082e-4f0f-9b2c-85268114433a	25	0be22e6c-0830-49eb-a500-b08b396d9c92	2022-09-30 16:31:39.658076+05	2022-09-30 16:31:39.658076+05	\N
0098fd41-42cc-4ba0-b771-5c593f91c6dd	81b84c5d-9759-4b86-978a-649c8ef79660	3	0be22e6c-0830-49eb-a500-b08b396d9c92	2022-09-30 16:31:39.667322+05	2022-09-30 16:31:39.667322+05	\N
276d10a7-7202-4fc7-8e83-e6e01a9ce83e	8df705a5-2351-4aca-b03e-3357a23840b4	3	0be22e6c-0830-49eb-a500-b08b396d9c92	2022-09-30 16:31:39.677826+05	2022-09-30 16:31:39.677826+05	\N
d579b953-9b2f-42b4-8aff-7c4be2ab150b	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	0be22e6c-0830-49eb-a500-b08b396d9c92	2022-09-30 16:31:39.691084+05	2022-09-30 16:31:39.691084+05	\N
4d419ab4-85c3-459a-b003-aaa151252a6b	d4156225-082e-4f0f-9b2c-85268114433a	25	ad758850-26c1-4405-b20c-2a7d01df3447	2022-09-30 16:33:55.080361+05	2022-09-30 16:33:55.080361+05	\N
f516f085-69cd-45d5-b187-8783a76fe995	81b84c5d-9759-4b86-978a-649c8ef79660	3	ad758850-26c1-4405-b20c-2a7d01df3447	2022-09-30 16:33:55.093023+05	2022-09-30 16:33:55.093023+05	\N
e2e63ade-a3dd-4c68-a0f4-4b4ed414c72e	8df705a5-2351-4aca-b03e-3357a23840b4	3	ad758850-26c1-4405-b20c-2a7d01df3447	2022-09-30 16:33:55.103089+05	2022-09-30 16:33:55.103089+05	\N
374ef930-ce47-409d-b6fa-94c20ee63d3b	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	ad758850-26c1-4405-b20c-2a7d01df3447	2022-09-30 16:33:55.113321+05	2022-09-30 16:33:55.113321+05	\N
8e8707d6-f1cc-4e04-8d3a-24253e7090ee	d4156225-082e-4f0f-9b2c-85268114433a	25	b486486f-cc0c-4aa2-b594-ba21465fafa7	2022-09-30 16:34:12.180261+05	2022-09-30 16:34:12.180261+05	\N
89722c21-dc79-4065-9326-b01ac0df0403	81b84c5d-9759-4b86-978a-649c8ef79660	3	b486486f-cc0c-4aa2-b594-ba21465fafa7	2022-09-30 16:34:12.192195+05	2022-09-30 16:34:12.192195+05	\N
d4b13472-7246-41da-9c31-960588bbcf89	8df705a5-2351-4aca-b03e-3357a23840b4	3	b486486f-cc0c-4aa2-b594-ba21465fafa7	2022-09-30 16:34:12.203999+05	2022-09-30 16:34:12.203999+05	\N
f4ad51c8-1fc5-4388-84b8-0bb9f4751458	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	b486486f-cc0c-4aa2-b594-ba21465fafa7	2022-09-30 16:34:12.213614+05	2022-09-30 16:34:12.213614+05	\N
2c6c9339-12ee-4c5b-a946-b67bd3c97d45	d4156225-082e-4f0f-9b2c-85268114433a	25	aa004448-b5ad-4248-be73-09c46f7aa2fa	2022-09-30 16:37:51.407009+05	2022-09-30 16:37:51.407009+05	\N
36380463-617f-4097-9561-a3be4ca4a8d9	81b84c5d-9759-4b86-978a-649c8ef79660	3	aa004448-b5ad-4248-be73-09c46f7aa2fa	2022-09-30 16:37:51.417225+05	2022-09-30 16:37:51.417225+05	\N
06b7a726-28dd-46ea-8322-430afc6a6210	8df705a5-2351-4aca-b03e-3357a23840b4	3	aa004448-b5ad-4248-be73-09c46f7aa2fa	2022-09-30 16:37:51.428901+05	2022-09-30 16:37:51.428901+05	\N
fa142e21-8039-463d-909f-fa4bf60ec8d7	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	aa004448-b5ad-4248-be73-09c46f7aa2fa	2022-09-30 16:37:51.439923+05	2022-09-30 16:37:51.439923+05	\N
66388d55-7a55-452c-87fa-35c4ea013731	d4156225-082e-4f0f-9b2c-85268114433a	25	3340c300-0d2d-4de9-af6a-244f777cd035	2022-09-30 16:38:29.907342+05	2022-09-30 16:38:29.907342+05	\N
a270da1d-f372-4430-b896-b77db21989ba	81b84c5d-9759-4b86-978a-649c8ef79660	3	3340c300-0d2d-4de9-af6a-244f777cd035	2022-09-30 16:38:29.918482+05	2022-09-30 16:38:29.918482+05	\N
846fe0ae-3e59-4f56-95c4-2a4bdc9ea79e	8df705a5-2351-4aca-b03e-3357a23840b4	3	3340c300-0d2d-4de9-af6a-244f777cd035	2022-09-30 16:38:29.929103+05	2022-09-30 16:38:29.929103+05	\N
82decffc-7185-40fb-8feb-555275331311	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	3340c300-0d2d-4de9-af6a-244f777cd035	2022-09-30 16:38:29.941543+05	2022-09-30 16:38:29.941543+05	\N
7d5bd974-0abe-4afa-8e79-1d4a21af9f20	d4156225-082e-4f0f-9b2c-85268114433a	25	4e06d814-d2ab-4244-81aa-d86c341690ab	2022-09-30 16:39:02.562989+05	2022-09-30 16:39:02.562989+05	\N
043406ae-a5d4-432a-8e60-928ea5ff0c0b	81b84c5d-9759-4b86-978a-649c8ef79660	3	4e06d814-d2ab-4244-81aa-d86c341690ab	2022-09-30 16:39:02.574699+05	2022-09-30 16:39:02.574699+05	\N
6ebc0970-1ed8-43db-9d3e-067b3237e8aa	8df705a5-2351-4aca-b03e-3357a23840b4	3	4e06d814-d2ab-4244-81aa-d86c341690ab	2022-09-30 16:39:02.58548+05	2022-09-30 16:39:02.58548+05	\N
8b6908e1-0b55-4f31-a404-804036107ad2	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	4e06d814-d2ab-4244-81aa-d86c341690ab	2022-09-30 16:39:02.596711+05	2022-09-30 16:39:02.596711+05	\N
8cda470b-541b-47eb-b313-f93b7797037d	d4156225-082e-4f0f-9b2c-85268114433a	25	7bf4fcbc-8526-4211-8666-8a3bf9af63c4	2022-09-30 16:48:58.062478+05	2022-09-30 16:48:58.062478+05	\N
3ae56ef8-08a2-4174-b0b1-7c8198e2e56a	81b84c5d-9759-4b86-978a-649c8ef79660	3	7bf4fcbc-8526-4211-8666-8a3bf9af63c4	2022-09-30 16:48:58.073894+05	2022-09-30 16:48:58.073894+05	\N
a05d1a61-1479-4762-9c85-9d11e0c8b684	8df705a5-2351-4aca-b03e-3357a23840b4	3	7bf4fcbc-8526-4211-8666-8a3bf9af63c4	2022-09-30 16:48:58.083382+05	2022-09-30 16:48:58.083382+05	\N
0428ab7e-5d9b-4cbd-bb62-3dbb3d09f519	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	7bf4fcbc-8526-4211-8666-8a3bf9af63c4	2022-09-30 16:48:58.094764+05	2022-09-30 16:48:58.094764+05	\N
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, customer_id, customer_mark, order_time, payment_type, total_price, created_at, updated_at, deleted_at, order_number) FROM stdin;
3d29c94b-1869-4d4c-a94f-180a5c9eb614	7e872c52-0d23-4086-8c45-43000b57332e	isleg market bet cykypdyr	12:00 - 16:00	nagt	1223.6	2022-09-21 21:28:30.39466+05	2022-09-21 21:28:30.39466+05	\N	1
ef468ad0-747b-4c12-a195-c72b9f19dc84	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1223.6	2022-09-22 12:54:24.199845+05	2022-09-22 12:54:24.199845+05	\N	2
e9b43b21-2947-4c41-95e6-1b81d925996b	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-22 14:27:08.833782+05	2022-09-22 14:27:08.833782+05	\N	3
9f1f4526-7ee7-4347-96fc-6fb97c1b402a	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-22 14:45:46.386254+05	2022-09-22 14:45:46.386254+05	\N	4
7728fdeb-e871-414d-b5fd-d57ee59e879c	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-29 14:05:14.735068+05	2022-09-29 14:05:14.735068+05	\N	5
8930edc5-7b9e-4cef-ac58-97bc8261df3a	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-29 14:06:29.676335+05	2022-09-29 14:06:29.676335+05	\N	6
3a8f977f-2eb9-4a4a-9bb1-51cd9a5a4fe6	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-29 14:06:57.81957+05	2022-09-29 14:06:57.81957+05	\N	7
7b0af658-e230-405d-a1ed-718f52973eb2	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-29 14:08:09.972501+05	2022-09-29 14:08:09.972501+05	\N	8
68c886f4-c005-4a22-9219-f643577d0a62	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-29 14:08:56.681211+05	2022-09-29 14:08:56.681211+05	\N	9
c09ef7ff-4244-422a-b2da-3e160e39f0c8	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:17:12.28516+05	2022-09-30 16:17:12.28516+05	\N	10
1499d084-b010-4724-ac40-2451a87986d9	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:19:46.479759+05	2022-09-30 16:19:46.479759+05	\N	11
4271c7d2-eeff-4948-989f-ae5e4518854f	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:22:21.342127+05	2022-09-30 16:22:21.342127+05	\N	12
6fda5178-9944-4ffe-92e7-319f41323818	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:28:44.051653+05	2022-09-30 16:28:44.051653+05	\N	13
67f2638b-971c-48a9-bada-5d14b50d74bc	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:29:07.367164+05	2022-09-30 16:29:07.367164+05	\N	14
0be22e6c-0830-49eb-a500-b08b396d9c92	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:31:39.632959+05	2022-09-30 16:31:39.632959+05	\N	15
ad758850-26c1-4405-b20c-2a7d01df3447	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:33:55.024097+05	2022-09-30 16:33:55.024097+05	\N	16
b486486f-cc0c-4aa2-b594-ba21465fafa7	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:34:12.15561+05	2022-09-30 16:34:12.15561+05	\N	17
aa004448-b5ad-4248-be73-09c46f7aa2fa	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:37:51.380433+05	2022-09-30 16:37:51.380433+05	\N	18
3340c300-0d2d-4de9-af6a-244f777cd035	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:38:29.875889+05	2022-09-30 16:38:29.875889+05	\N	19
4e06d814-d2ab-4244-81aa-d86c341690ab	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:39:02.54119+05	2022-09-30 16:39:02.54119+05	\N	20
7bf4fcbc-8526-4211-8666-8a3bf9af63c4	89a6ac71-4495-4218-b9f9-3f2a3eab085b	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-09-30 16:48:58.039309+05	2022-09-30 16:48:58.039309+05	\N	21
\.


--
-- Data for Name: payment_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_types (id, lang_id, type, created_at, updated_at, deleted_at) FROM stdin;
83e6589c-0cb6-4267-bcc5-e06cc93b36d8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	наличные	2022-09-20 14:33:50.780468+05	2022-09-20 14:33:50.780468+05	\N
7a6a313d-8fcd-4c56-9fa5-aefb12552b82	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	töleg terminaly	2022-09-20 14:34:46.329459+05	2022-09-20 14:34:46.329459+05	\N
cb7e8cc9-9b2e-4cd8-921f-91b3bb5e5564	aea98b93-7bdf-455b-9ad4-a259d69dc76e	платежный терминал	2022-09-20 14:34:46.359276+05	2022-09-20 14:34:46.359276+05	\N
38696743-82e5-4644-9c86-4a99ae45f912	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	nagt_tm	2022-09-20 14:33:50.755689+05	2022-09-20 14:40:04.959827+05	\N
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, brend_id, price, old_price, amount, product_code, created_at, updated_at, deleted_at, limit_amount, is_new) FROM stdin;
0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	214be879-65c3-4710-86b4-3fc3bce2e974	65	68	1000	151fwe51we	2022-09-17 14:54:56.989242+05	2022-09-17 14:54:56.989242+05	\N	100	f
b2b165a3-2261-4d67-8160-0e239ecd99b5	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	65	68	1000	151fwe51we	2022-09-17 14:55:35.441733+05	2022-09-17 14:55:35.441733+05	\N	100	f
a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	fdd259c2-794a-42b9-a3ad-9e91502af23e	65	68	1000	151fwe51we	2022-09-17 14:56:05.406905+05	2022-09-17 14:56:05.406905+05	\N	100	f
d731b17a-ae8d-4561-ad67-0f431d5c529b	f53a27b4-7810-4d8f-bd45-edad405d92b9	65	68	1000	151fwe51we	2022-09-17 14:56:36.153769+05	2022-09-17 14:56:36.153769+05	\N	100	f
bb6c3bdb-79e2-44b3-98b1-c1cee0976777	46b13f0a-d584-4ad3-b270-437ecdc51449	65	68	1000	151fwe51we	2022-09-17 14:57:07.191142+05	2022-09-17 14:57:07.191142+05	\N	100	f
d4156225-082e-4f0f-9b2c-85268114433a	c4bcda34-7332-4ae5-8129-d7538d63fee4	65	68	1000	151fwe51we	2022-09-17 14:57:34.582228+05	2022-09-17 14:57:34.582228+05	\N	100	f
81b84c5d-9759-4b86-978a-649c8ef79660	214be879-65c3-4710-86b4-3fc3bce2e974	65	68	1000	151fwe51we	2022-09-17 14:58:09.998335+05	2022-09-17 14:58:09.998335+05	\N	100	f
660071e0-8f17-4c48-9d80-d4cac306de3a	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	65	68	1000	151fwe51we	2022-09-17 14:58:40.084476+05	2022-09-17 14:58:40.084476+05	\N	100	f
c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	fdd259c2-794a-42b9-a3ad-9e91502af23e	65	68	1000	151fwe51we	2022-09-17 14:59:14.037001+05	2022-09-17 14:59:14.037001+05	\N	100	f
e3c33ead-3c30-40f1-9d28-7bb8b71b767f	f53a27b4-7810-4d8f-bd45-edad405d92b9	65	68	1000	151fwe51we	2022-09-17 14:59:44.837302+05	2022-09-17 14:59:44.837302+05	\N	100	f
8df705a5-2351-4aca-b03e-3357a23840b4	46b13f0a-d584-4ad3-b270-437ecdc51449	65	68	1000	151fwe51we	2022-09-17 15:00:15.148583+05	2022-09-17 15:00:15.148583+05	\N	100	f
3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	46b13f0a-d584-4ad3-b270-437ecdc51449	67	68	1000	151fwe51we	2022-10-06 11:07:41.410248+05	2022-10-06 11:07:41.410248+05	\N	100	f
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

COPY public.translation_basket_page (id, lang_id, quantity_of_goods, total_price, discount, delivery, total, currency, to_order, your_basket, created_at, updated_at, deleted_at, empty_the_basket) FROM stdin;
456dcb5a-fabb-47f8-b216-0cddd3077124	aea98b93-7bdf-455b-9ad4-a259d69dc76e	quantity_of_goods_ru	total_price_ru	discount_ru	delivery_ru	total_ru	currency_ru	to_order_ru	your_basket_ru	2022-08-30 12:36:24.978404+05	2022-08-30 12:36:37.967063+05	\N	uytget
51b3699e-1c7b-442a-be7b-6b2ad1f111b4	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	quantity_of_goods	total_price	discount	delivery	total	currency	to_order	your_basket	2022-08-30 12:36:24.978404+05	2022-09-19 14:28:12.008122+05	\N	empty_the_basket
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
-- Data for Name: translation_order_dates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_order_dates (id, lang_id, order_date_id, date, created_at, updated_at, deleted_at) FROM stdin;
dcd0c70b-9fa2-4327-8b35-de29bd3febcb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	32646376-c93f-412b-9e75-b3a5fa70df9e	şu gün	2022-09-28 17:35:33.812812+05	2022-09-28 17:35:33.812812+05	\N
3338d831-f091-4574-a0bf-f9cb07dd4893	aea98b93-7bdf-455b-9ad4-a259d69dc76e	32646376-c93f-412b-9e75-b3a5fa70df9e	Cегодня	2022-09-28 17:35:33.82453+05	2022-09-28 17:35:33.82453+05	\N
1aa5185f-9815-4e3f-9c34-718bfb587d91	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c1f2beca-a6b6-4971-a6a7-ed50079c6912	Ertir	2022-09-28 17:36:46.836838+05	2022-09-28 17:36:46.836838+05	\N
9e7a3752-fce2-4b66-bf3e-d915bf463f92	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c1f2beca-a6b6-4971-a6a7-ed50079c6912	Завтра	2022-09-28 17:36:46.847888+05	2022-09-28 17:36:46.847888+05	\N
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

COPY public.translation_product (id, lang_id, product_id, name, description, created_at, updated_at, deleted_at, slug) FROM stdin;
2edb91d0-4d17-4128-9bf8-0eb594418ee5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:54:57.051301+05	2022-09-17 14:54:57.051301+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
2756e684-ad6a-4e95-89f8-75b509f63290	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b2b165a3-2261-4d67-8160-0e239ecd99b5	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:55:35.51906+05	2022-09-17 14:55:35.51906+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
1f038393-3955-4cef-a6a1-a1cf087173c5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:56:05.464418+05	2022-09-17 14:56:05.464418+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
ad74bc57-3cd1-4c50-9287-a6a21b4beca4	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d731b17a-ae8d-4561-ad67-0f431d5c529b	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:56:36.207944+05	2022-09-17 14:56:36.207944+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
e12a74a8-f5c4-4dcd-b8c7-038a8d27624d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:54:57.063296+05	2022-09-17 14:54:57.063296+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
33011809-8da8-4563-8dd0-22ca01e5caee	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b2b165a3-2261-4d67-8160-0e239ecd99b5	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:55:35.530179+05	2022-09-17 14:55:35.530179+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
c499a61c-948d-41d4-9dad-9d12ea7324d4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:56:05.475198+05	2022-09-17 14:56:05.475198+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
ceeb7241-e4b8-4fa0-b99f-80ba9c141589	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d731b17a-ae8d-4561-ad67-0f431d5c529b	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:56:36.219444+05	2022-09-17 14:56:36.219444+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
e156d07d-c7e9-40da-8480-93512f474f80	aea98b93-7bdf-455b-9ad4-a259d69dc76e	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:57:07.264006+05	2022-09-17 14:57:07.264006+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
ce61fdba-2628-4f09-aff2-27ce8ac6b37c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d4156225-082e-4f0f-9b2c-85268114433a	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:57:34.654635+05	2022-09-17 14:57:34.654635+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
66d53b03-1f2b-4a7a-bc86-dd90baaec6ef	aea98b93-7bdf-455b-9ad4-a259d69dc76e	81b84c5d-9759-4b86-978a-649c8ef79660	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:58:10.064659+05	2022-09-17 14:58:10.064659+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
9bc5a5be-72d2-4494-ac02-dd20954a83ab	aea98b93-7bdf-455b-9ad4-a259d69dc76e	660071e0-8f17-4c48-9d80-d4cac306de3a	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:58:40.15465+05	2022-09-17 14:58:40.15465+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
5aa614ae-8c7d-47d7-867e-44c3d4d2015c	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:57:07.254121+05	2022-09-17 14:57:07.254121+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
e008099d-ff03-4182-86ce-d91ca984ca76	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d4156225-082e-4f0f-9b2c-85268114433a	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:57:34.642037+05	2022-09-17 14:57:34.642037+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
8965478c-0afe-4b65-af05-0d151c8dd462	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	81b84c5d-9759-4b86-978a-649c8ef79660	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:58:10.054293+05	2022-09-17 14:58:10.054293+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
3ec435d8-394a-4002-959e-c2d61d242307	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	660071e0-8f17-4c48-9d80-d4cac306de3a	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:58:40.143451+05	2022-09-17 14:58:40.143451+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
8a7fc25f-4776-498c-818d-95b9fb34fd2d	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:59:14.099998+05	2022-09-17 14:59:14.099998+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
8f666b20-35be-41df-b276-472ae2d5dd3d	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:59:44.899731+05	2022-09-17 14:59:44.899731+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
eeccc321-6c74-42c0-8ea7-acda231fc47b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	8df705a5-2351-4aca-b03e-3357a23840b4	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 15:00:15.211143+05	2022-09-17 15:00:15.211143+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
c5e4d0da-dbe9-46c6-87b4-3b70226ca2a9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-10-06 11:07:41.682197+05	2022-10-06 11:07:41.682197+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
925e8b5a-eec6-4f0b-8f46-9718a8f4f653	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:59:14.111821+05	2022-09-17 14:59:14.111821+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
dcedd91a-b7e9-4e9f-ae0a-f58308e6d751	aea98b93-7bdf-455b-9ad4-a259d69dc76e	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:59:44.910924+05	2022-09-17 14:59:44.910924+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
54b061df-a71c-405e-8ddd-0155b034dcd5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	8df705a5-2351-4aca-b03e-3357a23840b4	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 15:00:15.223949+05	2022-09-17 15:00:15.223949+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
ab660f8c-5ba8-45c4-be0b-ad2ef1450d1d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-10-06 11:07:41.692074+05	2022-10-06 11:07:41.692074+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
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
-- Name: orders_order_number_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_order_number_seq', 21, true);


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
-- Name: cart basket_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart
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
-- Name: customer_address customer_address_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer_address
    ADD CONSTRAINT customer_address_pkey PRIMARY KEY (id);


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
-- Name: order_dates order_dates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_dates
    ADD CONSTRAINT order_dates_pkey PRIMARY KEY (id);


--
-- Name: order_times order_times_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_times
    ADD CONSTRAINT order_times_pkey PRIMARY KEY (id);


--
-- Name: ordered_products ordered_products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ordered_products
    ADD CONSTRAINT ordered_products_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: payment_types payment_types_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_types
    ADD CONSTRAINT payment_types_pkey PRIMARY KEY (id);


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
-- Name: translation_order_dates translation_order_dates_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_order_dates
    ADD CONSTRAINT translation_order_dates_pkey PRIMARY KEY (id);


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
-- Name: customer_address customer_customer_address; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer_address
    ADD CONSTRAINT customer_customer_address FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: orders customers_orders; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT customers_orders FOREIGN KEY (customer_id) REFERENCES public.customers(id) ON UPDATE CASCADE ON DELETE CASCADE;


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
-- Name: cart fk_customer_basket; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart
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
-- Name: cart fk_product_basket; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cart
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
-- Name: payment_types languages_payment_types; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.payment_types
    ADD CONSTRAINT languages_payment_types FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_order_dates languages_translation_order_dates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_order_dates
    ADD CONSTRAINT languages_translation_order_dates FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: order_times order_dates_order_times; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.order_times
    ADD CONSTRAINT order_dates_order_times FOREIGN KEY (order_date_id) REFERENCES public.order_dates(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_order_dates order_dates_translation_order_dates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_order_dates
    ADD CONSTRAINT order_dates_translation_order_dates FOREIGN KEY (order_date_id) REFERENCES public.order_dates(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: ordered_products orders_ordered_products; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ordered_products
    ADD CONSTRAINT orders_ordered_products FOREIGN KEY (order_id) REFERENCES public.orders(id) ON UPDATE CASCADE ON DELETE CASCADE;


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
-- Name: ordered_products products_ordered_products; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ordered_products
    ADD CONSTRAINT products_ordered_products FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

