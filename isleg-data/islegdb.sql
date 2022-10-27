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


--
-- Name: after_delete_category(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.after_delete_category(cat_id uuid)
    LANGUAGE plpgsql
    AS $$
DECLARE category_uuid uuid;
BEGIN
FOR category_uuid IN SELECT id FROM categories WHERE parent_category_id = cat_id
LOOP UPDATE translation_category SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE category_product SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE products SET deleted_at = now() FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = category_uuid;
UPDATE translation_product SET deleted_at = now() FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE main_image SET deleted_at = now() FROM products,category_product WHERE main_image.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE images SET deleted_at = now() FROM products,category_product WHERE images.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE category_shop SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE shops SET deleted_at = now() FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = category_uuid; END LOOP; END;
$$;


ALTER PROCEDURE public.after_delete_category(cat_id uuid) OWNER TO postgres;

--
-- Name: after_insert_language(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.after_insert_language() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE language_id uuid; afisa_uuid uuid; category_uuid uuid; district_uuid uuid; order_date_uuid uuid; product_uuid uuid;
BEGIN
language_id = (SELECT id FROM languages ORDER BY created_at DESC LIMIT 1);
FOR afisa_uuid IN SELECT id FROM afisa
LOOP INSERT INTO translation_afisa (afisa_id,lang_id) VALUES (afisa_uuid,language_id); END LOOP;
FOR category_uuid IN SELECT id FROM categories
LOOP INSERT INTO translation_category (lang_id,category_id) VALUES (language_id,category_uuid); END LOOP;
INSERT INTO company_address (lang_id) VALUES (language_id);
FOR district_uuid IN SELECT id FROM district
LOOP INSERT INTO translation_district (lang_id,district_id) VALUES (language_id,district_uuid); END LOOP;
FOR order_date_uuid IN SELECT id FROM order_dates
LOOP INSERT INTO translation_order_dates (lang_id,order_date_id) VALUES (language_id,order_date_uuid); END LOOP;
INSERT INTO payment_types (lang_id) VALUES (language_id);
FOR product_uuid IN SELECT id FROM products
LOOP INSERT INTO translation_product (lang_id,product_id) VALUES (language_id,product_uuid); END LOOP;
INSERT INTO translation_about (lang_id) VALUES (language_id);
INSERT INTO translation_basket_page (lang_id) VALUES (language_id);
INSERT INTO translation_contact (lang_id) VALUES (language_id);
INSERT INTO translation_footer (lang_id) VALUES (language_id);
INSERT INTO translation_header (lang_id) VALUES (language_id);
INSERT INTO translation_my_information_page (lang_id) VALUES (language_id);
INSERT INTO translation_my_order_page (lang_id) VALUES (language_id);
INSERT INTO translation_order_page (lang_id) VALUES (language_id);
INSERT INTO translation_payment (lang_id) VALUES (language_id);
INSERT INTO translation_secure (lang_id) VALUES (language_id);
INSERT INTO translation_update_password_page (lang_id) VALUES (language_id);
RETURN NEW; END; $$;


ALTER FUNCTION public.after_insert_language() OWNER TO postgres;

--
-- Name: after_restore_category(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.after_restore_category(cat_id uuid)
    LANGUAGE plpgsql
    AS $$
DECLARE category_uuid uuid;
BEGIN
FOR category_uuid IN SELECT id FROM categories WHERE parent_category_id = cat_id
LOOP UPDATE translation_category SET deleted_at = NULL WHERE category_id = category_uuid;
UPDATE category_product SET deleted_at = NULL WHERE category_id = category_uuid;
UPDATE products SET deleted_at = NULL FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = category_uuid;
UPDATE translation_product SET deleted_at = NULL FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE main_image SET deleted_at = NULL FROM products,category_product WHERE main_image.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE images SET deleted_at = NULL FROM products,category_product WHERE images.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE category_shop SET deleted_at = NULL WHERE category_id = category_uuid;
UPDATE shops SET deleted_at = NULL FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = category_uuid;
END LOOP; END; $$;


ALTER PROCEDURE public.after_restore_category(cat_id uuid) OWNER TO postgres;

--
-- Name: delete_afisa(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_afisa(a_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE afisa SET deleted_at = now() WHERE id = a_id;
UPDATE translation_afisa SET deleted_at = now() WHERE afisa_id = a_id;
END; $$;


ALTER PROCEDURE public.delete_afisa(a_id uuid) OWNER TO postgres;

--
-- Name: delete_brend(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_brend(b_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE brends SET deleted_at = now() WHERE id = b_id;
UPDATE products SET deleted_at = now() WHERE brend_id = b_id;
UPDATE translation_product SET deleted_at = now() FROM products WHERE translation_product.product_id=products.id AND products.brend_id = b_id;
UPDATE main_image SET deleted_at = now() FROM products WHERE main_image.product_id=products.id AND products.brend_id = b_id;
UPDATE images SET deleted_at = now() FROM products WHERE images.product_id=products.id AND products.brend_id = b_id;
END; $$;


ALTER PROCEDURE public.delete_brend(b_id uuid) OWNER TO postgres;

--
-- Name: delete_category(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_category(category_uuid uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE categories SET deleted_at = now() WHERE id = category_uuid;
UPDATE translation_category SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE categories SET deleted_at = now() WHERE parent_category_id = category_uuid;
UPDATE category_product SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE products SET deleted_at = now() FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = category_uuid;
UPDATE translation_product SET deleted_at = now() FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE category_shop SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE shops SET deleted_at = now() FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = category_uuid;
UPDATE main_image SET deleted_at = now() FROM products,category_product WHERE main_image.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE images SET deleted_at = now() FROM products,category_product WHERE images.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
END; $$;


ALTER PROCEDURE public.delete_category(category_uuid uuid) OWNER TO postgres;

--
-- Name: delete_language(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_language(language_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE languages SET deleted_at = now() WHERE id = language_id;
UPDATE payment_types SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_order_dates SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_my_order_page SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_order_page SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_basket_page SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_header SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_footer SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_secure SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_payment SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_about SET deleted_at = now() WHERE lang_id = language_id;
UPDATE company_address SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_contact SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_my_information_page SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_update_password_page SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_category SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_product SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_afisa SET deleted_at = now() WHERE lang_id = language_id;
UPDATE translation_district SET deleted_at = now() WHERE lang_id = language_id;
END; $$;


ALTER PROCEDURE public.delete_language(language_id uuid) OWNER TO postgres;

--
-- Name: delete_product(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_product(p_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE products SET deleted_at = now() WHERE id = p_id;
UPDATE category_product SET deleted_at = now() WHERE product_id = p_id;
UPDATE translation_product SET deleted_at = now() WHERE product_id = p_id;
UPDATE main_image SET deleted_at = now() WHERE product_id = p_id;
UPDATE images SET deleted_at = now() WHERE product_id = p_id;
END; $$;


ALTER PROCEDURE public.delete_product(p_id uuid) OWNER TO postgres;

--
-- Name: delete_shop(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_shop(s_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE shops SET deleted_at = now() WHERE id = s_id;
UPDATE category_shop SET deleted_at = now() WHERE shop_id = s_id;
END; $$;


ALTER PROCEDURE public.delete_shop(s_id uuid) OWNER TO postgres;

--
-- Name: restore_afisa(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_afisa(a_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE afisa SET deleted_at = NULL WHERE id = a_id;
UPDATE translation_afisa SET deleted_at = NULL WHERE afisa_id = a_id;
END; $$;


ALTER PROCEDURE public.restore_afisa(a_id uuid) OWNER TO postgres;

--
-- Name: restore_brend(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_brend(b_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE brends SET deleted_at = NULL WHERE id = b_id;
UPDATE products SET deleted_at = NULL WHERE brend_id = b_id;
UPDATE translation_product SET deleted_at = NULL FROM products WHERE translation_product.product_id=products.id AND products.brend_id = b_id;
UPDATE main_image SET deleted_at = NULL FROM products WHERE main_image.product_id=products.id AND products.brend_id = b_id;
UPDATE images SET deleted_at = NULL FROM products WHERE images.product_id=products.id AND products.brend_id = b_id;
END; $$;


ALTER PROCEDURE public.restore_brend(b_id uuid) OWNER TO postgres;

--
-- Name: restore_category(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_category(cat_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE categories SET deleted_at = NULL WHERE id = cat_id;
UPDATE translation_category SET deleted_at = NULL WHERE category_id = cat_id;
UPDATE categories SET deleted_at = NULL WHERE parent_category_id = cat_id;
UPDATE category_product SET deleted_at = NULL WHERE category_id = cat_id;
UPDATE products SET deleted_at = NULL FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = cat_id;
UPDATE translation_product SET deleted_at = NULL FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = cat_id;
UPDATE main_image SET deleted_at = NULL FROM products,category_product WHERE main_image.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = cat_id;
UPDATE images SET deleted_at = NULL FROM products,category_product WHERE images.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = cat_id;
UPDATE category_shop SET deleted_at = NULL WHERE category_id = cat_id;
UPDATE shops SET deleted_at = NULL FROM category_shop WHERE category_shop.shop_id = shops.id AND category_shop.category_id = cat_id;
END; $$;


ALTER PROCEDURE public.restore_category(cat_id uuid) OWNER TO postgres;

--
-- Name: restore_language(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_language(language_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE languages SET deleted_at = NULL WHERE id = language_id;
UPDATE payment_types SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_order_dates SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_my_order_page SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_basket_page SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_order_page SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_header SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_footer SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_secure SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_payment SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_about SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE company_address SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_contact SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_my_information_page SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_update_password_page SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_category SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_product SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_afisa SET deleted_at = NULL WHERE lang_id = language_id;
UPDATE translation_district SET deleted_at = NULL WHERE lang_id = language_id;
END; $$;


ALTER PROCEDURE public.restore_language(language_id uuid) OWNER TO postgres;

--
-- Name: restore_product(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_product(p_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE products SET deleted_at = NULL WHERE id = p_id;
UPDATE category_product SET deleted_at = NULL WHERE product_id = p_id;
UPDATE translation_product SET deleted_at = NULL WHERE product_id = p_id;
UPDATE main_image SET deleted_at = NULL WHERE product_id = p_id;
UPDATE images SET deleted_at = NULL WHERE product_id = p_id;
END; $$;


ALTER PROCEDURE public.restore_product(p_id uuid) OWNER TO postgres;

--
-- Name: restore_shop(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_shop(s_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE shops SET deleted_at = NULL WHERE id = s_id;
UPDATE category_shop SET deleted_at = NULL WHERE shop_id = s_id;
END; $$;


ALTER PROCEDURE public.restore_shop(s_id uuid) OWNER TO postgres;

--
-- Name: update_updated_at(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.update_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$ BEGIN
NEW.updated_at=now();
RETURN NEW;
END; $$;


ALTER FUNCTION public.update_updated_at() OWNER TO postgres;

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
c4bcda34-7332-4ae5-8129-d7538d63fee4	Golden Eagle	uploads/brend/7a425220-7200-4eda-9013-c2d10eca4c89.jpg	2022-08-12 10:36:10.886455+05	2022-10-22 02:53:49.519217+05	\N
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
d154a3f1-7086-439f-b343-3998d6521efa	\N	uploads/category/19b7d8b4-1a33-47aa-be68-6b3bcc6e3c1a.png	f	2022-10-27 12:35:14.821001+05	2022-10-27 12:35:14.821001+05	\N
ab28ad8f-72af-4e9e-841b-38a6e6881a6e	d154a3f1-7086-439f-b343-3998d6521efa		t	2022-10-27 12:37:27.789152+05	2022-10-27 12:37:27.789152+05	\N
d7862d17-0742-4bd5-8fc8-478fd7e868c4	d154a3f1-7086-439f-b343-3998d6521efa		t	2022-10-27 12:38:27.306873+05	2022-10-27 12:38:27.306873+05	\N
71994790-1b7b-41ab-90a8-b3df0d68e3e6	d154a3f1-7086-439f-b343-3998d6521efa		t	2022-10-27 12:38:49.13083+05	2022-10-27 12:38:49.13083+05	\N
75dd289a-f72b-42fa-975e-ee10cd796135	d154a3f1-7086-439f-b343-3998d6521efa		t	2022-10-27 12:39:23.932688+05	2022-10-27 12:39:23.932688+05	\N
\.


--
-- Data for Name: category_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category_product (id, category_id, product_id, created_at, updated_at, deleted_at) FROM stdin;
b3ae4bc3-95b5-4cab-aa0b-1eed3f428463	d154a3f1-7086-439f-b343-3998d6521efa	d085e5a4-8229-4177-b5e1-623e80846017	2022-10-27 12:45:23.542988+05	2022-10-27 12:45:23.542988+05	\N
c01031e1-3874-4134-ad48-cc5a3d3cc1b9	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	d085e5a4-8229-4177-b5e1-623e80846017	2022-10-27 12:45:23.542988+05	2022-10-27 12:45:23.542988+05	\N
690da0d9-69e3-4c5b-9654-dec0018649b6	d154a3f1-7086-439f-b343-3998d6521efa	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	2022-10-27 12:47:51.191216+05	2022-10-27 12:47:51.191216+05	\N
a9f1e3e7-1630-464d-aca6-14e967c009aa	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	2022-10-27 12:47:51.191216+05	2022-10-27 12:47:51.191216+05	\N
a6d711ae-c6c8-46f4-bf05-e10c60d8696c	d154a3f1-7086-439f-b343-3998d6521efa	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	2022-10-27 12:49:37.993256+05	2022-10-27 12:49:37.993256+05	\N
556814d3-0f90-44aa-8e54-534ea6af6c66	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	2022-10-27 12:49:37.993256+05	2022-10-27 12:49:37.993256+05	\N
0b32aa17-21a0-46cc-9d5c-48c3d1c8a8a2	d154a3f1-7086-439f-b343-3998d6521efa	70a75d8b-d570-41d4-95cb-2199f4417542	2022-10-27 13:05:10.732541+05	2022-10-27 13:05:10.732541+05	\N
cc552ed8-cd25-4e3f-8aff-7357c5714a4e	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	70a75d8b-d570-41d4-95cb-2199f4417542	2022-10-27 13:05:10.732541+05	2022-10-27 13:05:10.732541+05	\N
d3baf690-3bf3-434f-9a5e-afca4babe5f6	d154a3f1-7086-439f-b343-3998d6521efa	ee1d67ed-5862-4dfc-8424-52531a240a6c	2022-10-27 13:07:14.125364+05	2022-10-27 13:07:14.125364+05	\N
867b1ca4-2369-405f-8a75-3dff33ca3d71	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	ee1d67ed-5862-4dfc-8424-52531a240a6c	2022-10-27 13:07:14.125364+05	2022-10-27 13:07:14.125364+05	\N
cc01ef82-2dc5-4b20-95e3-dfed4d22a4ef	d154a3f1-7086-439f-b343-3998d6521efa	c14c7f18-77db-4e3c-8939-e6001cb95db0	2022-10-27 13:09:04.181164+05	2022-10-27 13:09:04.181164+05	\N
50844c9a-eafe-4e60-95b4-e566a5cd7923	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	c14c7f18-77db-4e3c-8939-e6001cb95db0	2022-10-27 13:09:04.181164+05	2022-10-27 13:09:04.181164+05	\N
ad17e673-67ea-4a31-bb01-012cc45b9825	d154a3f1-7086-439f-b343-3998d6521efa	793be71f-b0fa-43a2-b527-5fb09236f530	2022-10-27 13:22:11.437897+05	2022-10-27 13:22:11.437897+05	\N
321e5048-ca26-4a4e-8c7f-bfc758e6c099	d7862d17-0742-4bd5-8fc8-478fd7e868c4	793be71f-b0fa-43a2-b527-5fb09236f530	2022-10-27 13:22:11.437897+05	2022-10-27 13:22:11.437897+05	\N
2db6e80f-0bc8-4688-8813-59def8634220	d154a3f1-7086-439f-b343-3998d6521efa	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	2022-10-27 13:24:26.65126+05	2022-10-27 13:24:26.65126+05	\N
3ce334ea-f8e2-41f7-b6c7-3dd71fec7c61	d7862d17-0742-4bd5-8fc8-478fd7e868c4	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	2022-10-27 13:24:26.65126+05	2022-10-27 13:24:26.65126+05	\N
f6840cbd-ae5e-42a4-8746-d90595d0d6a0	d154a3f1-7086-439f-b343-3998d6521efa	3f397126-6d8d-4a0d-982c-01fd00526957	2022-10-27 13:26:05.330899+05	2022-10-27 13:26:05.330899+05	\N
79782ad3-0b26-4081-9c7f-c8ae45f39d0e	d7862d17-0742-4bd5-8fc8-478fd7e868c4	3f397126-6d8d-4a0d-982c-01fd00526957	2022-10-27 13:26:05.330899+05	2022-10-27 13:26:05.330899+05	\N
eced4828-b523-446b-af5b-b5b78e5053b5	d154a3f1-7086-439f-b343-3998d6521efa	ccb43083-1c9e-4e84-bffd-ecb28474165e	2022-10-27 13:29:08.279964+05	2022-10-27 13:29:08.279964+05	\N
2fd6bbe7-0b16-4605-a6e2-c8d82c62eff4	d7862d17-0742-4bd5-8fc8-478fd7e868c4	ccb43083-1c9e-4e84-bffd-ecb28474165e	2022-10-27 13:29:08.279964+05	2022-10-27 13:29:08.279964+05	\N
f961661a-de6d-4cfe-8c52-4686406cce6b	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	ccb43083-1c9e-4e84-bffd-ecb28474165e	2022-10-27 13:29:08.279964+05	2022-10-27 13:29:08.279964+05	\N
395621d1-8269-4fd3-8626-8a6712916d74	d154a3f1-7086-439f-b343-3998d6521efa	83da5c7b-bffe-4450-97c9-0f376441b1d4	2022-10-27 13:30:49.724642+05	2022-10-27 13:30:49.724642+05	\N
94a6f4a3-6d81-4c9e-a99e-b8dbfdc1b71a	d7862d17-0742-4bd5-8fc8-478fd7e868c4	83da5c7b-bffe-4450-97c9-0f376441b1d4	2022-10-27 13:30:49.724642+05	2022-10-27 13:30:49.724642+05	\N
ecc4b4b6-9fb7-4371-8df8-d4c780c53098	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	83da5c7b-bffe-4450-97c9-0f376441b1d4	2022-10-27 13:30:49.724642+05	2022-10-27 13:30:49.724642+05	\N
b83c9c2c-e6f5-4a04-9f4b-67e531a87e64	d154a3f1-7086-439f-b343-3998d6521efa	0946a0f5-d23f-4660-9151-80ef91ae9747	2022-10-27 13:32:30.048139+05	2022-10-27 13:32:30.048139+05	\N
0069bea4-e091-4894-82f2-0057952b2ce7	d7862d17-0742-4bd5-8fc8-478fd7e868c4	0946a0f5-d23f-4660-9151-80ef91ae9747	2022-10-27 13:32:30.048139+05	2022-10-27 13:32:30.048139+05	\N
d9d36d7e-bc22-4b4f-aef9-9375aca69223	d154a3f1-7086-439f-b343-3998d6521efa	03050bc6-6223-49f3-b729-397fd3b6b285	2022-10-27 13:36:08.830146+05	2022-10-27 13:36:08.830146+05	\N
14e3d2d1-eac0-45bd-9c2e-73ad51ad9c07	71994790-1b7b-41ab-90a8-b3df0d68e3e6	03050bc6-6223-49f3-b729-397fd3b6b285	2022-10-27 13:36:08.830146+05	2022-10-27 13:36:08.830146+05	\N
93325946-9d80-40a4-9761-f2d3a8d7393a	d154a3f1-7086-439f-b343-3998d6521efa	8b481e58-cd39-4761-a052-75e30124689a	2022-10-27 13:38:14.209569+05	2022-10-27 13:38:14.209569+05	\N
aeed9d4e-aa12-4874-859f-13cb5f178cfc	71994790-1b7b-41ab-90a8-b3df0d68e3e6	8b481e58-cd39-4761-a052-75e30124689a	2022-10-27 13:38:14.209569+05	2022-10-27 13:38:14.209569+05	\N
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
6cbdf2b6-a0fd-4bdf-a2d6-1899f36f361a	fec81bff-8264-403f-b2ba-58afd53c821b	jkebdhwej	2022-10-24 02:21:24.168988+05	2022-10-24 02:21:24.168988+05	\N	f
fcfb981d-b578-4b41-9c86-7b92f0c5fe14	655c5504-1547-4daf-abe0-80b4116684f0	dcmnsdk	2022-10-24 02:23:26.42299+05	2022-10-24 02:23:26.42299+05	\N	f
243d3411-68a3-413a-ab67-cedef7ad6551	1f5bc917-fc85-46cf-a1b2-7c14cfe940be	Mir 2/2 jay 7 oy 36	2022-10-24 11:16:38.435106+05	2022-10-24 11:16:38.435106+05	\N	f
6a0d59dc-e354-4c60-8115-13f86599d7f9	82bf039d-ac70-49a4-abf3-0be224403fbf	dclkwnekjnwefjk	2022-10-26 01:27:21.468466+05	2022-10-26 01:27:21.468466+05	\N	f
cb16e7d8-4e37-4894-ad24-9ea31cb00590	9f4622b6-33df-4f87-a08b-dde3ba9f99ea	wkednwej	2022-10-26 01:28:10.48565+05	2022-10-26 01:28:10.48565+05	\N	f
04ed2b19-a903-4926-af42-8f42e70e88a3	58291dce-f935-407c-bea2-69e57c819ac9	ewdjnkewjdne	2022-10-26 01:36:02.45576+05	2022-10-26 01:36:02.45576+05	\N	f
7b2137a6-86b9-465c-9226-e2e762f6fcaa	c324f469-bd94-4346-8e8a-22476b44b7b5	wednwkejew	2022-10-26 01:36:41.248104+05	2022-10-26 01:36:41.248104+05	\N	f
80b6d5dc-865f-4c4e-9e1e-dc2c19b606e6	84973c6e-0b4e-4bf0-9762-624c229d8340	wdmwkldnjwd	2022-10-26 02:06:45.09904+05	2022-10-26 02:06:45.09904+05	\N	f
7b186c54-19ca-4139-9d57-156bf5c8e5f3	4c4471d1-d072-46a9-a43a-f12a48475061	wedmwekd	2022-10-26 02:11:29.972831+05	2022-10-26 02:11:29.972831+05	\N	f
0a1e9350-5592-4be5-a0ac-2e54a8b3ff55	598299f6-a2a4-40ce-91d2-64a1855548ca	welkuw4rfnw	2022-10-27 02:26:24.26481+05	2022-10-27 02:26:24.26481+05	\N	f
9e230fb8-956b-4cba-a7be-cf4b0b502700	63c735d5-b68a-48f2-b9c6-c30a9a0b2d3a	ewdmekd	2022-10-27 02:28:14.350437+05	2022-10-27 02:28:14.350437+05	\N	f
da04dc90-fa55-45df-acc2-2cb24b3cb651	6136ad89-9282-4551-a6be-c270b87dd491	wedwedewde	2022-10-27 02:36:50.490061+05	2022-10-27 02:36:50.490061+05	\N	f
41edbf83-adf9-4d38-8f56-fb4334d66243	b8da7f78-f427-4922-bb39-431c9bb7544c	sd,jcnjew	2022-10-27 02:37:49.334867+05	2022-10-27 02:37:49.334867+05	\N	f
ae32a106-8997-4774-bd30-a2a1ef1c1a24	530d0c93-9a83-48f8-a467-6f12a0c1a0d7	sdkmkewmdew	2022-10-27 02:44:47.900788+05	2022-10-27 02:44:47.900788+05	\N	f
f32f9d1c-f893-421e-8ab5-d499adfdb693	84973c6e-0b4e-4bf0-9762-624c229d8340	fknekfjern	2022-10-27 03:00:08.458444+05	2022-10-27 03:00:08.458444+05	\N	f
662220d8-2675-4799-9ca0-113c3ec46aca	dc82cb08-ebf6-4132-8464-0d776d883ba3	hrfwurh4wjh	2022-10-27 03:03:06.364489+05	2022-10-27 03:03:06.364489+05	\N	f
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (id, full_name, phone_number, password, birthday, gender, created_at, updated_at, deleted_at, email, is_register) FROM stdin;
1f5bc917-fc85-46cf-a1b2-7c14cfe940be	Allanur Bayramgeldiyew	+99362420377	$2a$14$4jP3cdo/xCmL.u6LEoFOJuL5vKlOmKg22cybQgnPvnbhY/xdNEi0e	\N	\N	2022-10-17 01:20:15.562427+05	2022-10-17 01:20:15.562427+05	\N	abb@gmail.com	t
f89518a7-cb3d-46fd-81cf-a0d334ca6d4d	Salam	+99363747156	$2a$14$1bmblLxYRUIgPGCc86MYb.MupBJ5eQ6b7I4hnAuInYSCtFDKi0VVq	\N	\N	2022-10-21 02:23:56.055341+05	2022-10-21 02:23:56.055341+05	\N	jsdjewb@gmail.com	t
c0534c60-137d-4e80-a0c2-57e355b0bf35	jdjwnedkjwenf	+99363747157	$2a$14$TZldeIc0GGJpM/x3O3Sh7u.pgdxP9gCD.U6UwGwphHptecq/Gu/da	\N	\N	2022-10-21 02:26:35.510969+05	2022-10-21 02:26:35.510969+05	\N	wednwefn@gmail.com	t
c403dd72-1383-45f9-881c-7f730858e677	ekdnwkej	+99363747158	$2a$14$sNajqJ9v2fwEly4zTAsJ5urAye6BfNJnSSG5BaW6Q55Rc1RdPyNo6	\N	\N	2022-10-21 02:29:44.14913+05	2022-10-21 02:29:44.14913+05	\N	wefjknwef@gmial.com	t
e24f8d12-d3f2-4a90-9038-66c036e2c336	wejdnkewjfkwjef	+99366666666	$2a$14$EbhbjK7bGpRJ/Z33Qk9Tg.VF8qG0ivaNLSxUZky3zbm/vBtESRm1u	\N	\N	2022-10-21 02:32:04.998568+05	2022-10-21 02:32:04.998568+05	\N	wefmnbwef@gmail.com	t
75345598-64cc-401e-8e52-16f092a29fec	wem, wef	+99363747159	$2a$14$p.vyA6WyfkQLPrwJ92oE0eq/xfAn7yPedJp5ifKN.Xtv300uoUMxm	\N	\N	2022-10-21 02:38:03.024387+05	2022-10-21 02:38:03.024387+05	\N	wednwef@gmail.com	t
d3aa02fb-12ca-4615-aa91-b6f5247dd3d6	ewdjbew	+99363747153	$2a$14$0DZ0MU7ZltOItfZkwP.hZetIA5ADFxqV3rK0X8oN.TObtL7EXTUp2	\N	\N	2022-10-21 02:44:42.933799+05	2022-10-21 02:44:42.933799+05	\N	wedwe@gmail.com	t
2c222c2d-4d34-4023-88fb-0018e56a96f7	edbewjh@gmail.com	+99364454465	$2a$14$7uTwETTKruVStLdsaE..SOG4Hd7K7lvzqgw4r68A5hBEf0lT02uQe	\N	\N	2022-10-21 02:56:36.54818+05	2022-10-21 02:56:36.54818+05	\N	123456@gmail.com	t
fec81bff-8264-403f-b2ba-58afd53c821b	jkwedhw3	+99362323232	\N	\N	\N	2022-10-24 02:21:23.927414+05	2022-10-24 02:21:23.927414+05	\N	\N	f
655c5504-1547-4daf-abe0-80b4116684f0	jsdfne	+99362321231	\N	\N	\N	2022-10-24 02:23:26.392303+05	2022-10-24 02:23:26.392303+05	\N	\N	f
f25a66d3-93ac-4da4-b237-d34867d5ca8f	Salam	+99362111111	$2a$14$L0vGymLsVyGTFCMUUo0poOMWzamcBDO0ajua.A2bMih4FcvnPyIjS	1998-06-24	\N	2022-10-17 00:50:48.568182+05	2022-10-24 12:00:37.896099+05	\N	salam@gmail.com	t
48c02978-22fe-42e7-a292-1190f407e08a	Allanur Bayramgeldiyew	+99365259874	$2a$14$xKig.NqzRICqV6SCK2YsoO1AUWUE5et8caH/.0lagl2hf5c7YmpXa	\N	\N	2022-10-24 14:05:05.267965+05	2022-10-24 14:05:05.267965+05	\N	df7wef84e@gmail.com	t
82bf039d-ac70-49a4-abf3-0be224403fbf	wednkwe	+99366545646	\N	\N	\N	2022-10-26 01:27:21.348331+05	2022-10-26 01:27:21.348331+05	\N	\N	f
9f4622b6-33df-4f87-a08b-dde3ba9f99ea	wednkewj	+99364545455	\N	\N	\N	2022-10-26 01:28:10.464642+05	2022-10-26 01:28:10.464642+05	\N	\N	f
58291dce-f935-407c-bea2-69e57c819ac9	djkbehjdewb	+99363454545	\N	\N	\N	2022-10-26 01:36:02.428415+05	2022-10-26 01:36:02.428415+05	\N	\N	f
c324f469-bd94-4346-8e8a-22476b44b7b5	ewndknwek	+99362222222	\N	\N	\N	2022-10-26 01:36:41.230748+05	2022-10-26 01:36:41.230748+05	\N	\N	f
84973c6e-0b4e-4bf0-9762-624c229d8340	jdbwehjb	+99363333333	\N	\N	\N	2022-10-26 02:06:45.043124+05	2022-10-26 02:06:45.043124+05	\N	\N	f
4c4471d1-d072-46a9-a43a-f12a48475061	wdkwkednwek	+99364545454	\N	\N	\N	2022-10-26 02:11:29.947804+05	2022-10-26 02:11:29.947804+05	\N	\N	f
598299f6-a2a4-40ce-91d2-64a1855548ca	Muhammet	+99363747155	$2a$14$sW8sjPnhdlxEqPwPGTrTquwV4pY78wpT6D1unW3LKMdx1Yy/Jk9BS	\N	\N	2022-10-26 02:18:35.297044+05	2022-10-26 02:18:35.297044+05	\N	ednejwn@gmail.com	t
63c735d5-b68a-48f2-b9c6-c30a9a0b2d3a	,jwenkjewebjwehb dweh	+99364565454	\N	\N	\N	2022-10-27 02:28:14.20651+05	2022-10-27 02:28:14.20651+05	\N	\N	f
6136ad89-9282-4551-a6be-c270b87dd491	jwebwhj	+99362121212	\N	\N	\N	2022-10-27 02:36:49.743094+05	2022-10-27 02:36:49.743094+05	\N	\N	f
b8da7f78-f427-4922-bb39-431c9bb7544c	we,fnkwefj	+99362131232	\N	\N	\N	2022-10-27 02:37:49.269535+05	2022-10-27 02:37:49.269535+05	\N	\N	f
530d0c93-9a83-48f8-a467-6f12a0c1a0d7	edmwekd	+99363222323	\N	\N	\N	2022-10-27 02:44:47.883144+05	2022-10-27 02:44:47.883144+05	\N	\N	f
dc82cb08-ebf6-4132-8464-0d776d883ba3	efbewhjbew	+99366565656	\N	\N	\N	2022-10-27 03:03:06.341747+05	2022-10-27 03:03:06.341747+05	\N	\N	f
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
76ac5895-6c11-4383-a187-65e4a195449b	d085e5a4-8229-4177-b5e1-623e80846017	uploads/product/efb6fc81-c9b0-42f7-9ce0-1f414f2c715e.jpg	uploads/product/5fd3bef7-09d0-4b51-a4fb-59a6c8537d6a.jpg	2022-10-27 12:45:23.487097+05	2022-10-27 12:45:23.487097+05	\N
87f3cebb-c979-42db-a4bb-77b74398edb4	d085e5a4-8229-4177-b5e1-623e80846017	uploads/product/fd40bda0-de78-47b5-af32-191063f68243.jpg	uploads/product/8872c8a5-1be6-4cac-9b96-94a579dc3458.jpg	2022-10-27 12:45:23.487097+05	2022-10-27 12:45:23.487097+05	\N
c77af923-69c8-4cfa-8f66-1ef682b5c53c	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	uploads/product/fe37a494-a709-4d84-9eaf-44fba5ed9b0a.jpg	uploads/product/f5a137fa-a069-460b-8b6c-653b663f6b6a.jpg	2022-10-27 12:47:51.157834+05	2022-10-27 12:47:51.157834+05	\N
7dd56932-41ee-45f8-a69b-cf86ded111e1	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	uploads/product/aeec2982-e369-4663-a55a-049811b002e6.jpg	uploads/product/ba67d67d-bf0c-4a0e-adf5-6d54d6a05eb8.jpg	2022-10-27 12:47:51.157834+05	2022-10-27 12:47:51.157834+05	\N
418b0c99-69e3-41f8-bb88-2d3c187e827d	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	uploads/product/d22424be-72d3-44ea-bbfa-0de2e12c19c8.jpg	uploads/product/13c91aca-a93b-4b9a-9af1-4ef2357c6be4.jpg	2022-10-27 12:49:37.905319+05	2022-10-27 12:49:37.905319+05	\N
be8c51d7-ed1a-41ed-b237-30405b801211	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	uploads/product/4721e6b4-43d6-48ae-aec4-e96a2f45e7b7.jpg	uploads/product/21cc543f-79b8-4bfb-8197-0fe6141346ae.jpg	2022-10-27 12:49:37.905319+05	2022-10-27 12:49:37.905319+05	\N
c5e41ca8-f8e3-4383-bf82-3031ee3dfeb1	70a75d8b-d570-41d4-95cb-2199f4417542	uploads/product/05745cf6-6299-4d35-81e0-517930cb843e.jpg	uploads/product/2ed5df2e-1fe9-4e89-94e9-4a57199a2e75.jpg	2022-10-27 13:05:10.321846+05	2022-10-27 13:05:10.321846+05	\N
7f74e7cd-21ee-4400-896e-d93cdc88fb59	70a75d8b-d570-41d4-95cb-2199f4417542	uploads/product/6348f985-5db5-4d9e-b927-ea5b76b975ca.jpg	uploads/product/4d560c92-1fb7-4cb1-8090-447e94966c31.jpg	2022-10-27 13:05:10.321846+05	2022-10-27 13:05:10.321846+05	\N
5889a8ca-d478-40bd-a148-fdeb6b24237a	ee1d67ed-5862-4dfc-8424-52531a240a6c	uploads/product/fb1af5c5-3c6f-480a-83a4-d7a324b27ef1.jpg	uploads/product/4444ce85-b4af-4126-a078-d91970788872.jpg	2022-10-27 13:07:14.091292+05	2022-10-27 13:07:14.091292+05	\N
136f7f16-a271-46da-8577-e756f3aa6847	ee1d67ed-5862-4dfc-8424-52531a240a6c	uploads/product/fae54003-461b-43d7-84aa-c6ae5128c4af.jpg	uploads/product/31d00bf3-7de5-463b-b591-b5775638cc75.jpg	2022-10-27 13:07:14.091292+05	2022-10-27 13:07:14.091292+05	\N
5850c6bd-8c14-4fc0-bc76-ff909cb4f2ff	c14c7f18-77db-4e3c-8939-e6001cb95db0	uploads/product/304771ec-b010-40f3-92df-82b6a5b934de.jpg	uploads/product/2df895c6-4267-4232-82df-53212d785429.jpg	2022-10-27 13:09:04.147542+05	2022-10-27 13:09:04.147542+05	\N
b4cd94cc-0c7f-4f92-8360-89ff7d73fe8f	c14c7f18-77db-4e3c-8939-e6001cb95db0	uploads/product/cbc2c8cd-052a-4d6f-b2ae-ccf785bf31fb.jpg	uploads/product/3b4a1295-860c-4066-8e87-d5515cb81d5c.jpg	2022-10-27 13:09:04.147542+05	2022-10-27 13:09:04.147542+05	\N
8c1c80a9-64cb-4d9a-9986-bb056697eaa3	793be71f-b0fa-43a2-b527-5fb09236f530	uploads/product/f8f1e4f7-be4a-4b6d-a571-c632f36add50.jpg	uploads/product/7b417a83-b629-497d-b757-b03beb029f85.jpg	2022-10-27 13:22:11.393609+05	2022-10-27 13:22:11.393609+05	\N
e0eaf897-245e-421b-af9e-e899c6617e0e	793be71f-b0fa-43a2-b527-5fb09236f530	uploads/product/fa161b81-e4e1-4c22-bfa6-f5f7c7ae1756.jpg	uploads/product/36314b1f-a86d-4bf0-ae1d-f9a80745cdbf.jpg	2022-10-27 13:22:11.393609+05	2022-10-27 13:22:11.393609+05	\N
580ab999-7569-4e88-b345-27543cc96e8b	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	uploads/product/0150db4e-e9bf-4cc8-b8f6-7f7f62ef898f.jpg	uploads/product/cf067d6f-ffcf-44ec-a125-758f35d85b37.jpg	2022-10-27 13:24:26.617604+05	2022-10-27 13:24:26.617604+05	\N
865750f3-5d69-47a0-b312-9c2ab0cc2289	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	uploads/product/a724631a-74e7-48ab-b976-7b89f808388d.jpg	uploads/product/4402bb7b-6d14-444d-a18e-d2c34b72a727.jpg	2022-10-27 13:24:26.617604+05	2022-10-27 13:24:26.617604+05	\N
aecb0fb0-5a38-40a1-8af2-d09614a9b6bd	3f397126-6d8d-4a0d-982c-01fd00526957	uploads/product/2ea5faef-2ce3-44b6-80d1-7a7c15701f8c.jpg	uploads/product/9a0b4714-1ee7-4077-be1c-edd2018ade1c.jpg	2022-10-27 13:26:05.297633+05	2022-10-27 13:26:05.297633+05	\N
4eb50f79-8829-4079-b0cd-8331435823d5	3f397126-6d8d-4a0d-982c-01fd00526957	uploads/product/6109cab7-8d18-4e29-8223-710a9b721ca8.jpg	uploads/product/281fe419-0068-4fb2-bdfc-072a82c0c082.jpg	2022-10-27 13:26:05.297633+05	2022-10-27 13:26:05.297633+05	\N
6c1d1042-f8b2-4404-aa6e-8c8a64532545	ccb43083-1c9e-4e84-bffd-ecb28474165e	uploads/product/42f9cfe5-3c5d-4b69-a8e4-a3d9867a0f5b.jpg	uploads/product/c97ca423-724f-4713-b799-a011287c6aa5.jpg	2022-10-27 13:29:08.23347+05	2022-10-27 13:29:08.23347+05	\N
987bfb34-af2c-40f1-b540-f3ab8430326d	ccb43083-1c9e-4e84-bffd-ecb28474165e	uploads/product/3127d757-72a9-44d9-9595-b7ed97963f3b.jpg	uploads/product/2275f3ec-9942-41de-9caa-9ef406ecfd00.jpg	2022-10-27 13:29:08.23347+05	2022-10-27 13:29:08.23347+05	\N
abadfa0d-6999-40b1-b27b-b0084da7fcec	83da5c7b-bffe-4450-97c9-0f376441b1d4	uploads/product/162dec1d-64d9-4dda-84d2-0d065b7e0fe2.jpg	uploads/product/fda9bc4e-3727-48e8-aee2-b46737f0bf36.jpg	2022-10-27 13:30:49.679942+05	2022-10-27 13:30:49.679942+05	\N
f1eedc22-c922-4aea-93e0-c5b44449a23e	83da5c7b-bffe-4450-97c9-0f376441b1d4	uploads/product/bc01f470-6db7-4e73-bfe9-f6ea46ba34c6.jpg	uploads/product/b152ac0e-e7f2-4668-a3d8-abdcd118a2b2.jpg	2022-10-27 13:30:49.679942+05	2022-10-27 13:30:49.679942+05	\N
b6c7be30-57c3-4406-963e-a0c74d0e731b	0946a0f5-d23f-4660-9151-80ef91ae9747	uploads/product/583d21a2-5b92-4d30-9062-a7b646e981ca.jpg	uploads/product/4783671a-16a8-4858-b7f6-17ba9f0f8305.jpg	2022-10-27 13:32:30.014698+05	2022-10-27 13:32:30.014698+05	\N
80a99f3b-1586-4305-ab97-0af6b57785bd	0946a0f5-d23f-4660-9151-80ef91ae9747	uploads/product/3600b0d8-27ac-4bc8-adb1-e5826df56cff.jpg	uploads/product/802e22e1-b0a9-4f6b-9456-47e9dc309442.jpg	2022-10-27 13:32:30.014698+05	2022-10-27 13:32:30.014698+05	\N
b93797c8-b747-4e9b-bc2e-6c83bc6b9b8b	03050bc6-6223-49f3-b729-397fd3b6b285	uploads/product/ee8ef999-a666-4c66-a775-61e3acabda1d.jpg	uploads/product/db4aec4d-e284-4eda-9441-77926a80cc35.jpg	2022-10-27 13:36:08.795988+05	2022-10-27 13:36:08.795988+05	\N
615494e5-4be4-47c9-96f8-06dafaa1f898	03050bc6-6223-49f3-b729-397fd3b6b285	uploads/product/085cff96-14ce-4b4b-be05-71ca8590867f.jpg	uploads/product/c547330c-c275-4575-aa7a-543c7811d02c.jpg	2022-10-27 13:36:08.795988+05	2022-10-27 13:36:08.795988+05	\N
95e50227-58e1-438e-ac2e-f1288215cc73	8b481e58-cd39-4761-a052-75e30124689a	uploads/product/51b146be-a1a9-410e-8314-9ed26680d363.jpg	uploads/product/869b8b7c-5643-4596-b8db-468af8a4f312.jpg	2022-10-27 13:38:14.176348+05	2022-10-27 13:38:14.176348+05	\N
c92f6bfa-9139-433c-8177-52370121605e	8b481e58-cd39-4761-a052-75e30124689a	uploads/product/70598020-bd76-4bb8-b021-af48e70bf5d6.jpg	uploads/product/92eb371e-2832-41b8-bdbb-df223fdcd9a9.jpg	2022-10-27 13:38:14.176348+05	2022-10-27 13:38:14.176348+05	\N
\.


--
-- Data for Name: languages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.languages (id, name_short, flag, created_at, updated_at, deleted_at) FROM stdin;
aea98b93-7bdf-455b-9ad4-a259d69dc76e	ru	uploads/language1c24e3a6-173e-4264-a631-f099d15495dd.jpeg	2022-06-15 19:53:21.29491+05	2022-06-15 19:53:21.29491+05	\N
8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	tm	uploads/language17b99bd1-f52d-41db-b4e6-1ecff03e0fd0.jpeg	2022-06-15 19:53:06.041686+05	2022-10-16 18:53:27.82538+05	\N
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
1ecfc54d-0c07-4577-97e1-29503ca60c86	d085e5a4-8229-4177-b5e1-623e80846017	uploads/product/6a2b2732-6585-428e-b6d0-52e94fb9a4f2.jpg	uploads/product/810f948e-6a52-4319-84f8-1cd25b3f8bb7.jpg	uploads/product/6ca77765-43e0-48cb-86ca-573568ab4201.jpg	2022-10-27 12:45:23.472753+05	2022-10-27 12:45:23.472753+05	\N
4de19867-d401-4538-b7ce-015967d0c6f0	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	uploads/product/fd224ec4-c237-4079-811c-f5fa39e3e885.jpg	uploads/product/6c4ae0ab-58bb-4405-9cd4-657605995d1d.jpg	uploads/product/a301fd4e-2e75-4503-8530-8fdcf80b15b1.jpg	2022-10-27 12:47:51.142922+05	2022-10-27 12:47:51.142922+05	\N
b3460c07-2e1f-4110-94c2-14750748eec6	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	uploads/product/f498f90e-89a0-43e9-9116-0216348d10ab.jpg	uploads/product/4d44c079-787a-45a2-8b90-6a7d83d9ef2c.jpg	uploads/product/483dbc69-d63c-4202-9d55-c7f283837ae9.jpg	2022-10-27 12:49:37.845528+05	2022-10-27 12:49:37.845528+05	\N
258471e1-e002-4e3d-9be8-4a442a6d6d50	70a75d8b-d570-41d4-95cb-2199f4417542	uploads/product/1ec30d87-b8a3-4de8-a71f-1749328bf20e.jpg	uploads/product/88cf205f-c937-4f89-a29e-be0cc8b0d99e.jpg	uploads/product/82ad8ac8-251a-4432-9b07-19e3596198fb.jpg	2022-10-27 13:05:10.30663+05	2022-10-27 13:05:10.30663+05	\N
a95f9bba-0dc1-4096-90c3-db2cc308bda7	ee1d67ed-5862-4dfc-8424-52531a240a6c	uploads/product/38457790-da1e-4ffb-bdae-c28c40fc6534.jpg	uploads/product/60eeed69-6577-42d9-a7f8-08000ea7dbc2.jpg	uploads/product/8c6655d0-eb65-4431-81fd-916d35f42850.jpg	2022-10-27 13:07:14.075976+05	2022-10-27 13:07:14.075976+05	\N
bfffc5b6-956d-4460-9683-973289a4a76b	c14c7f18-77db-4e3c-8939-e6001cb95db0	uploads/product/0c009049-b607-44e3-8d25-140040506225.jpg	uploads/product/ed283bde-c753-49c7-8e3a-e40911da1ef6.jpg	uploads/product/48fc7ed3-5e18-4d57-88ee-bcd389377593.jpg	2022-10-27 13:09:04.132921+05	2022-10-27 13:09:04.132921+05	\N
69bbae4e-e69b-4e46-a0f5-8f6a6574843f	793be71f-b0fa-43a2-b527-5fb09236f530	uploads/product/c61c3627-e507-48ee-8287-e023d20a1339.jpg	uploads/product/5bf97605-7983-4767-84fb-6e11c3ac00cc.jpg	uploads/product/b302a6cb-eb01-4afb-991d-a95f2be25d9b.jpg	2022-10-27 13:22:11.378944+05	2022-10-27 13:22:11.378944+05	\N
c325207e-dfbf-4986-96e6-26fd7e17ccd2	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	uploads/product/d6338cc4-f07e-4068-a9f2-656e64886d62.jpg	uploads/product/ce1ced74-7e0d-4a97-bea7-63dfe2026a75.jpg	uploads/product/06ccab77-b6c9-4356-a182-46ff5de5e8d1.jpg	2022-10-27 13:24:26.603629+05	2022-10-27 13:24:26.603629+05	\N
57233981-9813-4f95-b436-7962f56f2889	3f397126-6d8d-4a0d-982c-01fd00526957	uploads/product/4d7d4295-07e6-4756-af1c-789cf6b8512f.jpg	uploads/product/c766e9c4-787f-41c6-b66c-13adab39c4ac.jpg	uploads/product/ce597557-b988-4a31-a2aa-c22b9f6e960a.jpg	2022-10-27 13:26:05.283184+05	2022-10-27 13:26:05.283184+05	\N
45456031-0230-493e-b897-bf218a376fbf	ccb43083-1c9e-4e84-bffd-ecb28474165e	uploads/product/e122a649-5132-45f5-99a7-ef299f683bcf.jpg	uploads/product/25d480da-32bc-42a2-a1ab-0780d1207171.jpg	uploads/product/63ad2aaf-0c61-4380-9b88-113b6b1eac5d.jpg	2022-10-27 13:29:08.219368+05	2022-10-27 13:29:08.219368+05	\N
ef3d46dc-e961-472e-ad8c-e4e61a10e7a2	83da5c7b-bffe-4450-97c9-0f376441b1d4	uploads/product/e4dacbb0-3fe1-456d-bd26-23e382a0d536.jpg	uploads/product/05d9a4d6-fddf-4a72-8c0b-650c6940217f.jpg	uploads/product/0e3b5a4d-daef-4cfa-8847-dd1538ac6ce1.jpg	2022-10-27 13:30:49.665842+05	2022-10-27 13:30:49.665842+05	\N
fa340958-a6c9-437d-9928-1f388892fb56	0946a0f5-d23f-4660-9151-80ef91ae9747	uploads/product/b627f961-60f8-40c9-91fa-21818c09cddb.jpg	uploads/product/09db62e8-d9ee-458a-be46-af4988bfe8cb.jpg	uploads/product/37db2574-b67a-4bca-b606-46273f27d63c.jpg	2022-10-27 13:32:30.000631+05	2022-10-27 13:32:30.000631+05	\N
e948a556-f723-4e50-b9d2-2d94b7c7e619	03050bc6-6223-49f3-b729-397fd3b6b285	uploads/product/bb8433b9-3f0b-4f47-8e67-3acad07820a4.jpg	uploads/product/a8edb71f-4a7c-4940-851f-ea281b3d6039.jpg	uploads/product/2e34230d-7ce7-4403-9646-61581a914d65.jpg	2022-10-27 13:36:08.781822+05	2022-10-27 13:36:08.781822+05	\N
8d74ecec-e8c9-4c36-8b90-361c5665a1b7	8b481e58-cd39-4761-a052-75e30124689a	uploads/product/17f209e5-1d6c-437e-b34d-adf5c62c8433.jpg	uploads/product/a998b549-d68a-4e47-8eb0-7cd345a2183a.jpg	uploads/product/96483dec-8e60-4c7c-9b45-f7f28713bac0.jpg	2022-10-27 13:38:14.095266+05	2022-10-27 13:38:14.095266+05	\N
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
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, customer_id, customer_mark, order_time, payment_type, total_price, created_at, updated_at, deleted_at, order_number) FROM stdin;
ed6f4a78-b4fa-4035-a8e0-d884c65a5889	fec81bff-8264-403f-b2ba-58afd53c821b	ajkedbhewj	18:00 - 21:00	nagt_tm	845	2022-10-24 02:21:24.249433+05	2022-10-24 02:21:24.249433+05	\N	22
2698cad6-a4f2-493d-b86d-d3d1659c1896	655c5504-1547-4daf-abe0-80b4116684f0	jenew	18:00 - 21:00	nagt_tm	845	2022-10-24 02:23:26.441517+05	2022-10-24 02:23:26.441517+05	\N	23
f767105b-51b9-47a2-a161-1646934f5acd	1f5bc917-fc85-46cf-a1b2-7c14cfe940be	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-10-24 11:16:38.47027+05	2022-10-24 11:16:38.47027+05	\N	24
b4c0dcac-8683-44aa-90bf-84aaa8bd811c	1f5bc917-fc85-46cf-a1b2-7c14cfe940be	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-10-24 23:40:17.660318+05	2022-10-24 23:40:17.660318+05	\N	25
8ae67524-a41b-4cea-95ec-4fb9a56baf00	1f5bc917-fc85-46cf-a1b2-7c14cfe940be	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-10-24 23:54:04.990677+05	2022-10-24 23:54:04.990677+05	\N	26
75b0e572-5074-484d-9f88-3eac9f10becc	1f5bc917-fc85-46cf-a1b2-7c14cfe940be	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-10-25 00:05:09.561504+05	2022-10-25 00:05:09.561504+05	\N	27
b9cb6d94-bb7e-43b4-959d-abd93725d60d	82bf039d-ac70-49a4-abf3-0be224403fbf	jenfkwef	09:00 - 12:00	töleg terminaly	1040	2022-10-26 01:27:21.651292+05	2022-10-26 01:27:21.651292+05	\N	28
3444b57e-3020-486d-b4e3-4b9c4c4a0dce	9f4622b6-33df-4f87-a08b-dde3ba9f99ea	wednwejd	09:00 - 12:00	töleg terminaly	1040	2022-10-26 01:28:10.509031+05	2022-10-26 01:28:10.509031+05	\N	29
78dd7fc6-da93-4ff5-a0bc-47dc71f40050	58291dce-f935-407c-bea2-69e57c819ac9	wedjnew	09:00 - 12:00	töleg terminaly	715	2022-10-26 01:36:02.471918+05	2022-10-26 01:36:02.471918+05	\N	30
94f947c3-77c1-444a-8e12-159da89b28f6	c324f469-bd94-4346-8e8a-22476b44b7b5	wednwejnew	09:00 - 12:00	töleg terminaly	780	2022-10-26 01:36:41.26726+05	2022-10-26 01:36:41.26726+05	\N	31
3b66a950-d316-47e4-893e-8af2808994a9	84973c6e-0b4e-4bf0-9762-624c229d8340	dnedkjewjde	18:00 - 21:00	töleg terminaly	260	2022-10-26 02:06:45.116062+05	2022-10-26 02:06:45.116062+05	\N	32
d98dad05-5a08-409c-a479-34634b3bdd02	4c4471d1-d072-46a9-a43a-f12a48475061	ewdjewn	09:00 - 12:00	töleg terminaly	195	2022-10-26 02:11:29.990018+05	2022-10-26 02:11:29.990018+05	\N	33
992812d2-5bb7-470b-af0d-121ce32b378f	598299f6-a2a4-40ce-91d2-64a1855548ca	kewdnwej	09:00 - 12:00	töleg terminaly	260	2022-10-27 02:26:24.342543+05	2022-10-27 02:26:24.342543+05	\N	34
e8852db0-1a76-4d6e-a29b-8d946108e4f2	63c735d5-b68a-48f2-b9c6-c30a9a0b2d3a	ednkewj	09:00 - 12:00	töleg terminaly	390	2022-10-27 02:28:14.442924+05	2022-10-27 02:28:14.442924+05	\N	35
bf63b1ac-88bd-47e8-a929-e7a58c32291a	6136ad89-9282-4551-a6be-c270b87dd491	wedwed	18:00 - 21:00	töleg terminaly	455	2022-10-27 02:36:51.327294+05	2022-10-27 02:36:51.327294+05	\N	36
40f78590-86eb-4cf5-9c0a-8ff10d2ee7ae	b8da7f78-f427-4922-bb39-431c9bb7544c	ewdfnwkefj	18:00 - 21:00	töleg terminaly	455	2022-10-27 02:37:49.558033+05	2022-10-27 02:37:49.558033+05	\N	37
26b27442-8a2d-4e27-886b-b790ae0d6078	530d0c93-9a83-48f8-a467-6f12a0c1a0d7	wedkwek	18:00 - 21:00	töleg terminaly	520	2022-10-27 02:44:47.917788+05	2022-10-27 02:44:47.917788+05	\N	38
06925508-8808-4db8-9766-aa906eab37ef	84973c6e-0b4e-4bf0-9762-624c229d8340	dnekjw	09:00 - 12:00	töleg terminaly	195	2022-10-27 03:00:08.518043+05	2022-10-27 03:00:08.518043+05	\N	39
4d667365-ea58-4227-b08a-9f1b15243249	dc82cb08-ebf6-4132-8464-0d776d883ba3	jewkfhbew	18:00 - 21:00	nagt_tm	325	2022-10-27 03:03:06.380206+05	2022-10-27 03:03:06.380206+05	\N	40
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

COPY public.products (id, brend_id, price, old_price, amount, created_at, updated_at, deleted_at, limit_amount, is_new) FROM stdin;
d085e5a4-8229-4177-b5e1-623e80846017	c4bcda34-7332-4ae5-8129-d7538d63fee4	52.4	61.6	1000	2022-10-27 12:45:23.437479+05	2022-10-27 12:45:23.437479+05	\N	100	f
4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	c4bcda34-7332-4ae5-8129-d7538d63fee4	82.7	91.8	1000	2022-10-27 12:47:51.121005+05	2022-10-27 12:47:51.121005+05	\N	100	f
2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	c4bcda34-7332-4ae5-8129-d7538d63fee4	86.3	158	1000	2022-10-27 12:49:37.648623+05	2022-10-27 12:49:37.648623+05	\N	100	f
70a75d8b-d570-41d4-95cb-2199f4417542	46b13f0a-d584-4ad3-b270-437ecdc51449	74.8	94.4	1000	2022-10-27 13:05:10.286344+05	2022-10-27 13:05:10.286344+05	\N	100	f
ee1d67ed-5862-4dfc-8424-52531a240a6c	f53a27b4-7810-4d8f-bd45-edad405d92b9	74.8	99.3	1000	2022-10-27 13:07:14.050311+05	2022-10-27 13:07:14.050311+05	\N	100	f
c14c7f18-77db-4e3c-8939-e6001cb95db0	f53a27b4-7810-4d8f-bd45-edad405d92b9	68.9	74.8	1000	2022-10-27 13:09:04.112256+05	2022-10-27 13:09:04.112256+05	\N	100	f
793be71f-b0fa-43a2-b527-5fb09236f530	fdd259c2-794a-42b9-a3ad-9e91502af23e	72.5	0	1000	2022-10-27 13:22:11.35263+05	2022-10-27 13:22:11.35263+05	\N	100	f
e34a20fa-3aef-4ba6-92ba-79d3649c61a6	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	109	0	1000	2022-10-27 13:24:26.583694+05	2022-10-27 13:24:26.583694+05	\N	100	f
3f397126-6d8d-4a0d-982c-01fd00526957	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	95	0	1000	2022-10-27 13:26:05.260347+05	2022-10-27 13:26:05.260347+05	\N	100	f
ccb43083-1c9e-4e84-bffd-ecb28474165e	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	92	96.9	1000	2022-10-27 13:29:08.198279+05	2022-10-27 13:29:08.198279+05	\N	100	f
83da5c7b-bffe-4450-97c9-0f376441b1d4	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	50.1	85.8	1000	2022-10-27 13:30:49.646484+05	2022-10-27 13:30:49.646484+05	\N	100	f
0946a0f5-d23f-4660-9151-80ef91ae9747	214be879-65c3-4710-86b4-3fc3bce2e974	141.2	0	1000	2022-10-27 13:32:29.977781+05	2022-10-27 13:32:29.977781+05	\N	100	f
03050bc6-6223-49f3-b729-397fd3b6b285	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	115	0	1000	2022-10-27 13:36:08.764157+05	2022-10-27 13:36:08.764157+05	\N	100	f
8b481e58-cd39-4761-a052-75e30124689a	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	161	0	1000	2022-10-27 13:38:14.068965+05	2022-10-27 13:38:14.068965+05	\N	100	f
\.


--
-- Data for Name: shops; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shops (id, owner_name, address, phone_number, running_time, created_at, updated_at, deleted_at) FROM stdin;
263a86a7-ee3f-4ec1-9933-d925bbb9e206	Owez Myradow	Asgabat saher Mir 2/2 jay 2 magazyn 23	+99362420387	7:00-21:00	2022-10-24 13:22:34.682812+05	2022-10-24 13:22:34.682812+05	\N
4cbec734-efd3-40cf-81a6-f145769da9a0	Owez Myradow	Asgabat saher Mir 2/2 jay 2 magazyn 23	+99362420387	7:00-21:00	2022-10-24 13:23:08.286246+05	2022-10-24 13:23:08.286246+05	\N
a789cd15-3fa4-49d3-bda0-eb47bafdc61f	Owez Myradow	Asgabat saher Mir 2/2 jay 2 magazyn 23	+99362420387	7:00-21:00	2022-10-24 13:23:29.106899+05	2022-10-24 13:23:29.106899+05	\N
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
5ef15568-e39e-4a3a-bf80-3695ae6e5367	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d154a3f1-7086-439f-b343-3998d6521efa	Arzanladyş we Aksiýalar	2022-10-27 12:35:14.838875+05	2022-10-27 12:35:14.838875+05	\N
00638721-b67f-41fd-b332-cdd96f66bf0c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d154a3f1-7086-439f-b343-3998d6521efa	Распродажи и Акции	2022-10-27 12:35:14.853749+05	2022-10-27 12:35:14.853749+05	\N
02d50a95-acb2-45d9-a9a7-edd3dd2d68e9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	Arzanladyşdaky harytlar	2022-10-27 12:37:27.851772+05	2022-10-27 12:37:27.851772+05	\N
23bf9d24-c3e3-4eb2-8900-3c7568010602	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	Продукция со скидкой	2022-10-27 12:37:27.866721+05	2022-10-27 12:37:27.866721+05	\N
c6086874-26c3-4ea3-bb70-3c88bf67643b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d7862d17-0742-4bd5-8fc8-478fd7e868c4	Sowgatlyk toplumlar	2022-10-27 12:38:27.36353+05	2022-10-27 12:38:27.36353+05	\N
c9f9ac63-172d-450f-80c0-eecfba4284d1	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d7862d17-0742-4bd5-8fc8-478fd7e868c4	Подарочные наборы	2022-10-27 12:38:27.377387+05	2022-10-27 12:38:27.377387+05	\N
0cf94c44-aab0-4180-8e07-b076babf5865	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	71994790-1b7b-41ab-90a8-b3df0d68e3e6	Aksiýadaky harytlar	2022-10-27 12:38:49.186311+05	2022-10-27 12:38:49.186311+05	\N
fc576a2d-3ab3-420c-856b-704daf0cc3ed	aea98b93-7bdf-455b-9ad4-a259d69dc76e	71994790-1b7b-41ab-90a8-b3df0d68e3e6	Продукция в категории Акции	2022-10-27 12:38:49.200519+05	2022-10-27 12:38:49.200519+05	\N
3a7f44fb-5e22-416b-9bdc-f20be8485b1b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	75dd289a-f72b-42fa-975e-ee10cd796135	Täze harytlar	2022-10-27 12:39:23.997701+05	2022-10-27 12:39:23.997701+05	\N
25b0fa42-8466-4cd0-a322-b1d02332b918	aea98b93-7bdf-455b-9ad4-a259d69dc76e	75dd289a-f72b-42fa-975e-ee10cd796135	Новые продукты	2022-10-27 12:39:24.012575+05	2022-10-27 12:39:24.012575+05	\N
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
9154e800-2a92-47de-b4ff-1e63b213e5f7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	поиск	tелефон	пароль	забыл пароль	войти	зарегистрироваться	имя	Подтвердить Пароль	Я прочитал и принимаю Условия Обслуживания и Политика Конфиденциальности	моя информация	мои любимые	мои заказы	выйти	2022-06-16 04:48:26.491672+05	2022-10-26 12:29:21.210919+05	\N	корзина	uytget_ru	uytget_ru
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

COPY public.translation_order_page (id, lang_id, content, type_of_payment, choose_a_delivery_time, your_address, mark, to_order, created_at, updated_at, deleted_at) FROM stdin;
474a15e9-1a05-49aa-9a61-c92837d9c9a8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	content_ru	type_of_payment_ru	choose_a_delivery_time_ru	your_address_ru	mark_ru	to_order_ru	2022-09-01 12:47:16.802639+05	2022-09-01 12:47:16.802639+05	\N
75810722-07fd-400e-94b4-cd230de08cbf	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	content	type_of_payment	choose_a_delivery_time	your_address	mark	to_order	2022-09-01 12:47:16.720956+05	2022-09-01 12:55:25.638676+05	\N
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
fdecb62e-e026-400c-8146-269552268363	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d085e5a4-8229-4177-b5e1-623e80846017	Tebigy ereýän kofe Maxwell House 150 gr	Tebigy ereýän kofe Maxwell House 150 gr	2022-10-27 12:45:23.512856+05	2022-10-27 12:45:23.512856+05	\N	tebigy-ereyan-kofe-maxwell-house-150-gr
b5441981-2b60-4d64-8ded-f6cd68e0bec7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d085e5a4-8229-4177-b5e1-623e80846017	Растворимый кофе Maxwell House 150 г	Растворимый кофе Maxwell House 150 г	2022-10-27 12:45:23.533062+05	2022-10-27 12:45:23.533062+05	\N	rastvorimyi-kofe-maxwell-house-150-g
cc5a5be3-8749-4d81-8a96-fc25163d045b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	Tebigy ereýän kofe Maxwell House 200+50 gr	Tebigy ereýän kofe Maxwell House 200+50 gr	2022-10-27 12:47:51.168874+05	2022-10-27 12:47:51.168874+05	\N	tebigy-ereyan-kofe-maxwell-house-200-50-gr
0d87cb5f-4d1f-49e8-a1b4-c26353bd8c8f	aea98b93-7bdf-455b-9ad4-a259d69dc76e	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	Растворимый кофе Maxwell House 200+50 г	Растворимый кофе Maxwell House 200+50 г	2022-10-27 12:47:51.1792+05	2022-10-27 12:47:51.1792+05	\N	rastvorimyi-kofe-maxwell-house-200-50-g
4c993d4a-f4d6-4d33-afca-8cf9fbf652c0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	Kofe Mövenpick "Gold Original 100% Arabica" 100 gr	Kofe Mövenpick "Gold Original 100% Arabica" 100 gr	2022-10-27 12:49:37.972094+05	2022-10-27 12:49:37.972094+05	\N	kofe-movenpick-gold-original-100-arabica-100-gr
6ba071ce-7548-48e4-ad83-444f175281ba	aea98b93-7bdf-455b-9ad4-a259d69dc76e	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	Кофе Mövenpick "Gold Original 100% Arabica" 100 г	Кофе Mövenpick "Gold Original 100% Arabica" 100 г	2022-10-27 12:49:37.981834+05	2022-10-27 12:49:37.981834+05	\N	kofe-movenpick-gold-original-100-arabica-100-g
c7344a8d-9a89-4acc-8cf3-be8643124045	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	70a75d8b-d570-41d4-95cb-2199f4417542	Üwelen kofe Mövenpick Der Himmlische 250 gr	Üwelen kofe Mövenpick Der Himmlische 250 gr	2022-10-27 13:05:10.332333+05	2022-10-27 13:05:10.332333+05	\N	uwelen-kofe-movenpick-der-himmlische-250-gr
e0565164-63dd-4bf6-8bc2-d1399961c610	aea98b93-7bdf-455b-9ad4-a259d69dc76e	70a75d8b-d570-41d4-95cb-2199f4417542	Молотый кофе Mövenpick Der Himmlische 250 гр	Молотый кофе Mövenpick Der Himmlische 250 гр	2022-10-27 13:05:10.705616+05	2022-10-27 13:05:10.705616+05	\N	molotyi-kofe-movenpick-der-himmlische-250-gr
8f518241-a80c-467d-bc40-c1c0f3ab8585	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ee1d67ed-5862-4dfc-8424-52531a240a6c	Kofe Idee Kaffee "100% Arabica" 250 gr	Kofe Idee Kaffee "100% Arabica" 250 gr	2022-10-27 13:07:14.101052+05	2022-10-27 13:07:14.101052+05	\N	kofe-idee-kaffee-100-arabica-250-gr
7f4a55fd-9237-4af0-ae60-2fc5191cca08	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ee1d67ed-5862-4dfc-8424-52531a240a6c	Кофе Idee Kaffee "100% Arabica" молотый 250 г	Кофе Idee Kaffee "100% Arabica" молотый 250 г	2022-10-27 13:07:14.112977+05	2022-10-27 13:07:14.112977+05	\N	kofe-idee-kaffee-100-arabica-molotyi-250-g
8bf6e83c-0394-4e3d-8e65-5cd87a82b8a9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c14c7f18-77db-4e3c-8939-e6001cb95db0	Kofe Nescafe Gold "Sumatra" çüýşe gapda 85 gr	Kofe Nescafe Gold "Sumatra" çüýşe gapda 85 gr	2022-10-27 13:09:04.159628+05	2022-10-27 13:09:04.159628+05	\N	kofe-nescafe-gold-sumatra-cuyse-gapda-85-gr
e504255e-51cb-4583-b37d-d91624aa94b1	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c14c7f18-77db-4e3c-8939-e6001cb95db0	Кофе Nescafe Gold "Sumatra" банка 85 гр	Кофе Nescafe Gold "Sumatra" банка 85 гр	2022-10-27 13:09:04.169875+05	2022-10-27 13:09:04.169875+05	\N	kofe-nescafe-gold-sumatra-banka-85-gr
6899fc36-f2b9-4e2e-81a9-7f517cd95079	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	793be71f-b0fa-43a2-b527-5fb09236f530	Sowgatlyk toplumy MEN DEEP CLEANING duş üçin krem-gel 300 ml+ýuwmak üçin gel HYDRO ENERgETIC 150 ml	Sowgatlyk toplumy MEN DEEP CLEANING duş üçin krem-gel 300 ml+ýuwmak üçin gel HYDRO ENERgETIC 150 ml	2022-10-27 13:22:11.416525+05	2022-10-27 13:22:11.416525+05	\N	sowgatlyk-toplumy-men-deep-cleaning-dus-ucin-krem-gel-300-ml-yuwmak-ucin-gel-hydro-energetic-150-ml
6d8cc27d-aa60-490a-b282-03f413bfa908	aea98b93-7bdf-455b-9ad4-a259d69dc76e	793be71f-b0fa-43a2-b527-5fb09236f530	Подарочный набор MEN DEEP CLEANINg крем-грель для душа 300 мл+грель для умыв HYDRO ENERgETIC 150 мл	Подарочный набор MEN DEEP CLEANINg крем-грель для душа 300 мл+грель для умыв HYDRO ENERgETIC 150 мл	2022-10-27 13:22:11.428036+05	2022-10-27 13:22:11.428036+05	\N	podarochnyi-nabor-men-deep-cleaning-krem-grel-dlia-dusha-300-ml-grel-dlia-umyv-hydro-energetic-150-ml
950aa440-28cf-4c13-8685-d8baf19c2bf8	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	Sowgatlyk toplumy UFC x EXXE sakgal syrmak üçin köpürjik + sakgal syrmak üçin gel + duş geli Ultimate freshness	Sowgatlyk toplumy UFC x EXXE sakgal syrmak üçin köpürjik + sakgal syrmak üçin gel + duş geli Ultimate freshness	2022-10-27 13:24:26.628997+05	2022-10-27 13:24:26.628997+05	\N	sowgatlyk-toplumy-ufc-x-exxe-sakgal-syrmak-ucin-kopurjik-sakgal-syrmak-ucin-gel-dus-geli-ultimate-freshness
42c64158-3179-4b2e-9d3b-8766e3586931	aea98b93-7bdf-455b-9ad4-a259d69dc76e	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	Подарочный набор UFC x EXXE пена для бритья + крем-бальзам после бритья + гель для душа Ultimate Freshness	Подарочный набор UFC x EXXE пена для бритья + крем-бальзам после бритья + гель для душа Ultimate Freshness	2022-10-27 13:24:26.641508+05	2022-10-27 13:24:26.641508+05	\N	podarochnyi-nabor-ufc-x-exxe-pena-dlia-brit-ia-krem-bal-zam-posle-brit-ia-gel-dlia-dusha-ultimate-freshness
5c87b6d0-7b5d-40df-a32d-9a12f8205e69	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	3f397126-6d8d-4a0d-982c-01fd00526957	Sowgatlyk toplumy UFC x EXXE duş geli + şampun Carbon Hit	Sowgatlyk toplumy UFC x EXXE duş geli + şampun Carbon Hit	2022-10-27 13:26:05.309306+05	2022-10-27 13:26:05.309306+05	\N	sowgatlyk-toplumy-ufc-x-exxe-dus-geli-sampun-carbon-hit
6caa3ccc-e904-413b-8730-80dbddf15790	aea98b93-7bdf-455b-9ad4-a259d69dc76e	3f397126-6d8d-4a0d-982c-01fd00526957	Подарочный набор UFC x EXXE гель для душа + шампунь Carbon Hit	Подарочный набор UFC x EXXE гель для душа + шампунь Carbon Hit	2022-10-27 13:26:05.319675+05	2022-10-27 13:26:05.319675+05	\N	podarochnyi-nabor-ufc-x-exxe-gel-dlia-dusha-shampun-carbon-hit
a440c88c-0030-445b-b675-94029afeacbc	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ccb43083-1c9e-4e84-bffd-ecb28474165e	Sowgatlyk toplumy UFC x EXXE duş geli + dezodorant Ultimate Freshness	Sowgatlyk toplumy UFC x EXXE duş geli + dezodorant Ultimate Freshness	2022-10-27 13:29:08.24936+05	2022-10-27 13:29:08.24936+05	\N	sowgatlyk-toplumy-ufc-x-exxe-dus-geli-dezodorant-ultimate-freshness
f2d39fe8-264d-4394-964a-88b04f133187	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ccb43083-1c9e-4e84-bffd-ecb28474165e	Подарочный набор UFC x EXXE гель для душа + дезодорант Ultimate Freshness	Подарочный набор UFC x EXXE гель для душа + дезодорант Ultimate Freshness	2022-10-27 13:29:08.266915+05	2022-10-27 13:29:08.266915+05	\N	podarochnyi-nabor-ufc-x-exxe-gel-dlia-dusha-dezodorant-ultimate-freshness
db6e8eac-1e4b-48ab-9082-040f0635ebe1	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	83da5c7b-bffe-4450-97c9-0f376441b1d4	Sowgatlyk toplum Le Petit Marseillais erkekler üçin gel-şampun "Narpyz we Laým" 250 ml	Sowgatlyk toplum Le Petit Marseillais erkekler üçin gel-şampun "Narpyz we Laým" 250 ml	2022-10-27 13:30:49.691056+05	2022-10-27 13:30:49.691056+05	\N	sowgatlyk-toplum-le-petit-marseillais-erkekler-ucin-gel-sampun-narpyz-we-laym-250-ml
e6457a78-48e1-4f73-b380-75f76892d146	aea98b93-7bdf-455b-9ad4-a259d69dc76e	83da5c7b-bffe-4450-97c9-0f376441b1d4	Подарочный набор Le Petit Marseillais гель-шампунь для мужчин "Мята и Лайм" 250 мл	Подарочный набор Le Petit Marseillais гель-шампунь для мужчин "Мята и Лайм" 250 мл	2022-10-27 13:30:49.702691+05	2022-10-27 13:30:49.702691+05	\N	podarochnyi-nabor-le-petit-marseillais-gel-shampun-dlia-muzhchin-miata-i-laim-250-ml
78a56e75-aa39-4cbd-9108-489192e01286	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0946a0f5-d23f-4660-9151-80ef91ae9747	Sowgatlyk toplumy Head & Shoulders "Saç üçin balzam 275 ml + Goňaga garşy şampun 400 ml	Sowgatlyk toplumy Head & Shoulders "Saç üçin balzam 275 ml + Goňaga garşy şampun 400 ml	2022-10-27 13:32:30.026538+05	2022-10-27 13:32:30.026538+05	\N	sowgatlyk-toplumy-head-and-shoulders-sac-ucin-balzam-275-ml-gonaga-garsy-sampun-400-ml
9f1877ae-af11-4021-a62c-8e7ff715a6f4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0946a0f5-d23f-4660-9151-80ef91ae9747	Подарочный Набор Head & Shoulders "Бальзам-ополаскиватель для волос 275 мл + Шампунь против перхоти 400 мл	Подарочный Набор Head & Shoulders "Бальзам-ополаскиватель для волос 275 мл + Шампунь против перхоти 400 мл	2022-10-27 13:32:30.038394+05	2022-10-27 13:32:30.038394+05	\N	podarochnyi-nabor-head-and-shoulders-bal-zam-opolaskivatel-dlia-volos-275-ml-shampun-protiv-perkhoti-400-ml
e8b956d7-8e65-44bc-a7a1-02fd3df23338	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	03050bc6-6223-49f3-b729-397fd3b6b285	(Aksiýa!) Kofe Nescafe Gold, paket gapda 220 gr + Kofe Nescafe Classic, 3x1 kiçi paket 14.5 gr	(Aksiýa!) Kofe Nescafe Gold, paket gapda 220 gr + Kofe Nescafe Classic, 3x1 kiçi paket 14.5 gr	2022-10-27 13:36:08.808395+05	2022-10-27 13:36:08.808395+05	\N	aksiya-kofe-nescafe-gold-paket-gapda-220-gr-kofe-nescafe-classic-3x1-kici-paket-14-5-gr
ee30076c-a648-4de4-9f62-238333d88481	aea98b93-7bdf-455b-9ad4-a259d69dc76e	03050bc6-6223-49f3-b729-397fd3b6b285	(Акция!) Кофе Nescafe Gold, пакет 220 г + Кофе Nescafe Classic 3в1, стик 14.5 гр	(Акция!) Кофе Nescafe Gold, пакет 220 г + Кофе Nescafe Classic 3в1, стик 14.5 гр	2022-10-27 13:36:08.818396+05	2022-10-27 13:36:08.818396+05	\N	aktsiia-kofe-nescafe-gold-paket-220-g-kofe-nescafe-classic-3v1-stik-14-5-gr
015b37fc-77f0-49ff-818d-a4443e60a454	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	8b481e58-cd39-4761-a052-75e30124689a	(4+1) Kofe Jacobs Monarch 47.5 gr (5 sany)	(4+1) Kofe Jacobs Monarch 47.5 gr (5 sany)	2022-10-27 13:38:14.188215+05	2022-10-27 13:38:14.188215+05	\N	4-1-kofe-jacobs-monarch-47-5-gr-5-sany
7bed1227-d2e6-4ba7-844f-3ffc70027a90	aea98b93-7bdf-455b-9ad4-a259d69dc76e	8b481e58-cd39-4761-a052-75e30124689a	(4+1) Кофе Jacobs Monarch 47.5 г (5 шт)	(4+1) Кофе Jacobs Monarch 47.5 г (5 шт)	2022-10-27 13:38:14.199712+05	2022-10-27 13:38:14.199712+05	\N	4-1-kofe-jacobs-monarch-47-5-g-5-sht
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

SELECT pg_catalog.setval('public.orders_order_number_seq', 40, true);


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
-- Name: languages after_insert_language; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER after_insert_language AFTER INSERT ON public.languages FOR EACH ROW EXECUTE FUNCTION public.after_insert_language();


--
-- Name: afisa updated_afisa_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_afisa_updated_at BEFORE UPDATE ON public.afisa FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: banner updated_banner_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_banner_updated_at BEFORE UPDATE ON public.banner FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: brends updated_brends_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_brends_updated_at BEFORE UPDATE ON public.brends FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: cart updated_cart_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_cart_updated_at BEFORE UPDATE ON public.cart FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: categories updated_categories_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_categories_updated_at BEFORE UPDATE ON public.categories FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: category_product updated_category_product_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_category_product_updated_at BEFORE UPDATE ON public.category_product FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: category_shop updated_category_shop_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_category_shop_updated_at BEFORE UPDATE ON public.category_shop FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: company_address updated_company_address_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_company_address_updated_at BEFORE UPDATE ON public.company_address FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: company_phone updated_company_phone_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_company_phone_updated_at BEFORE UPDATE ON public.company_phone FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: company_setting updated_company_setting_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_company_setting_updated_at BEFORE UPDATE ON public.company_setting FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: customer_address updated_customer_address_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_customer_address_updated_at BEFORE UPDATE ON public.customer_address FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: customers updated_customers_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_customers_updated_at BEFORE UPDATE ON public.customers FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: district updated_district_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_district_updated_at BEFORE UPDATE ON public.district FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: images updated_images_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_images_updated_at BEFORE UPDATE ON public.images FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: languages updated_languages_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_languages_updated_at BEFORE UPDATE ON public.languages FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: likes updated_likes_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_likes_updated_at BEFORE UPDATE ON public.likes FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: main_image updated_main_image_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_main_image_updated_at BEFORE UPDATE ON public.main_image FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: order_dates updated_order_dates_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_order_dates_updated_at BEFORE UPDATE ON public.order_dates FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: order_times updated_order_times_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_order_times_updated_at BEFORE UPDATE ON public.order_times FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: ordered_products updated_ordered_products_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_ordered_products_updated_at BEFORE UPDATE ON public.ordered_products FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: orders updated_orders_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_orders_updated_at BEFORE UPDATE ON public.orders FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: payment_types updated_payment_types_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_payment_types_updated_at BEFORE UPDATE ON public.payment_types FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: products updated_products_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_products_updated_at BEFORE UPDATE ON public.products FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: shops updated_shops_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_shops_updated_at BEFORE UPDATE ON public.shops FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_about updated_translation_about_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_about_updated_at BEFORE UPDATE ON public.translation_about FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_afisa updated_translation_afisa_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_afisa_updated_at BEFORE UPDATE ON public.translation_afisa FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_basket_page updated_translation_basket_page_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_basket_page_updated_at BEFORE UPDATE ON public.translation_basket_page FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_category updated_translation_category_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_category_updated_at BEFORE UPDATE ON public.translation_category FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_contact updated_translation_contact_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_contact_updated_at BEFORE UPDATE ON public.translation_contact FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_district updated_translation_district_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_district_updated_at BEFORE UPDATE ON public.translation_district FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_footer updated_translation_footer_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_footer_updated_at BEFORE UPDATE ON public.translation_footer FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_header updated_translation_header_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_header_updated_at BEFORE UPDATE ON public.translation_header FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_my_information_page updated_translation_my_information_page_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_my_information_page_updated_at BEFORE UPDATE ON public.translation_my_information_page FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_my_order_page updated_translation_my_order_page_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_my_order_page_updated_at BEFORE UPDATE ON public.translation_my_order_page FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_order_dates updated_translation_order_dates_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_order_dates_updated_at BEFORE UPDATE ON public.translation_order_dates FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_order_page updated_translation_order_page_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_order_page_updated_at BEFORE UPDATE ON public.translation_order_page FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_payment updated_translation_payment_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_payment_updated_at BEFORE UPDATE ON public.translation_payment FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_product updated_translation_product_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_product_updated_at BEFORE UPDATE ON public.translation_product FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_secure updated_translation_secure_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_secure_updated_at BEFORE UPDATE ON public.translation_secure FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


--
-- Name: translation_update_password_page updated_translation_update_password_page_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_update_password_page_updated_at BEFORE UPDATE ON public.translation_update_password_page FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


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

