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
c4bcda34-7332-4ae5-8129-d7538d63fee4	Golden Eagle	uploads/brend/7a425220-7200-4eda-9013-c2d10eca4c89.jpg	2022-08-12 10:36:10.886455+05	2022-10-22 02:53:49.519217+05	\N
\.


--
-- Data for Name: cart; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cart (id, product_id, customer_id, quantity_of_product, created_at, updated_at, deleted_at) FROM stdin;
d3ab09c6-b976-43bd-95cf-d7584746d540	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	f25a66d3-93ac-4da4-b237-d34867d5ca8f	1	2022-10-22 02:39:50.988017+05	2022-10-22 02:39:50.988017+05	\N
4a52b546-0d4a-4a13-80af-7175d3c135c0	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	f25a66d3-93ac-4da4-b237-d34867d5ca8f	1	2022-10-22 02:39:51.105734+05	2022-10-22 02:39:51.105734+05	\N
f2442927-9039-45d0-8245-9e32aa5d9ac9	d731b17a-ae8d-4561-ad67-0f431d5c529b	f25a66d3-93ac-4da4-b237-d34867d5ca8f	1	2022-10-22 02:39:51.12484+05	2022-10-22 02:39:51.12484+05	\N
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, parent_category_id, image, is_home_category, created_at, updated_at, deleted_at) FROM stdin;
f745d171-68e6-42e2-b339-cb3c210cda55	b982bd86-0a0f-4950-baad-5a131e9b728e		f	2022-06-16 13:45:48.828786+05	2022-06-16 13:45:48.828786+05	\N
d4cb1359-6c23-4194-8e3c-21ed8cec8373	5bb9a4e7-9992-418f-b551-537844d371da		f	2022-06-16 13:48:04.517774+05	2022-06-16 13:48:04.517774+05	\N
fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	5bb9a4e7-9992-418f-b551-537844d371da		f	2022-06-16 13:47:18.854741+05	2022-06-16 13:47:18.854741+05	\N
02bd4413-8586-49ab-802e-16304e756a8b	\N	uploads/category0684921b-251d-405f-8b30-30964be0b3d2.jpeg	f	2022-06-16 13:43:22.644619+05	2022-06-16 13:43:22.644619+05	\N
5bb9a4e7-9992-418f-b551-537844d371da	02bd4413-8586-49ab-802e-16304e756a8b		f	2022-06-16 13:46:44.575803+05	2022-06-16 13:46:44.575803+05	\N
b982bd86-0a0f-4950-baad-5a131e9b728e	02bd4413-8586-49ab-802e-16304e756a8b		f	2022-06-16 13:44:16.430875+05	2022-06-16 13:44:16.430875+05	\N
7605172f-7a12-4781-a892-6e3b5cf11490	\N	uploads/category/3d83e54b-6dcc-4ef6-86e1-85cf6a7906f3.jpg	t	2022-10-18 11:09:26.736735+05	2022-10-18 11:09:26.736735+05	\N
8f6af238-80ef-40f0-8c34-531b9c06b373	\N	uploads/category/7447c247-7ab9-4b84-adb0-219889c365e1.jpg	t	2022-10-21 03:21:07.93801+05	2022-10-21 03:21:07.93801+05	\N
417a385e-6a74-44f3-a536-405eb8251978	\N	uploads/category/fd959b97-7c98-4408-bd8a-85c5be950185.jpg	f	2022-10-21 03:24:33.292071+05	2022-10-21 03:24:33.292071+05	\N
b1bae1ce-4295-4268-bf2d-71c8761e5679	\N	uploads/category/561ae885-d566-4c93-b96b-a276dba26638.jpg	f	2022-10-21 03:26:21.879348+05	2022-10-21 03:26:21.879348+05	\N
d8ded28c-d4fb-4c11-a84c-4d4f81a22e28	\N	uploads/category/e80980fb-d4eb-4353-9f72-47924b203f1e.jpg	f	2022-10-21 03:28:14.644395+05	2022-10-21 03:28:14.644395+05	\N
cdb681a2-98e4-4716-a136-a5e4888e9c32	\N	uploads/category/cbab3616-8f46-44c1-92af-178516cd7dd0.jpg	f	2022-10-21 03:28:30.755427+05	2022-10-21 11:09:53.88461+05	\N
a4277afa-1c92-4f4e-809e-dfbb54ddbc9b	\N	uploads/category/e15eebcd-07fa-4020-be24-447366836e2e.jpg	f	2022-10-21 11:48:26.196782+05	2022-10-21 11:48:26.196782+05	\N
849a1c59-45fb-429b-8fe3-a6e34a6dafaa	7605172f-7a12-4781-a892-6e3b5cf11490	uploads/category/7dc0bb7f-6624-474b-a93c-06362e05bdb5.jpg	f	2022-10-21 11:48:48.606943+05	2022-10-21 12:15:12.803389+05	\N
bdabc7aa-a567-48d5-a1d9-b1ff61c6af4b	417a385e-6a74-44f3-a536-405eb8251978		f	2022-10-21 22:28:26.592037+05	2022-10-21 22:28:26.592037+05	\N
29ed85bb-11eb-4458-bbf3-5a5644d167d6	\N	uploads/categoryeaae1626-7e9f-4db9-abf6-f454ade813d3.jpeg	f	2022-06-20 09:41:17.575565+05	2022-10-22 01:26:07.073419+05	\N
7f453dd0-7b2e-480d-a8be-fcfa23bd863e	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:43:07.336084+05	2022-10-22 01:26:07.073419+05	\N
66772380-c161-4c45-9350-a45e765193e2	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:45:34.38667+05	2022-10-22 01:26:07.073419+05	\N
338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 09:46:01.119337+05	2022-10-22 01:26:07.073419+05	\N
45765130-7f97-4f0c-b886-f70b75e02610	29ed85bb-11eb-4458-bbf3-5a5644d167d6		t	2022-06-20 10:11:06.648938+05	2022-10-22 01:26:07.073419+05	\N
\.


--
-- Data for Name: category_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category_product (id, category_id, product_id, created_at, updated_at, deleted_at) FROM stdin;
d82042be-0468-446f-a06e-c569fc6967de	f745d171-68e6-42e2-b339-cb3c210cda55	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	2022-09-17 14:54:57.077351+05	2022-09-17 14:54:57.077351+05	\N
715c66a2-32b6-449b-8ed1-2f656ed07c2f	d4cb1359-6c23-4194-8e3c-21ed8cec8373	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	2022-09-17 14:54:57.085055+05	2022-09-17 14:54:57.085055+05	\N
703562cb-0dc5-4ea8-bb9b-a632150d9406	d4cb1359-6c23-4194-8e3c-21ed8cec8373	b2b165a3-2261-4d67-8160-0e239ecd99b5	2022-09-17 14:55:35.543212+05	2022-09-17 14:55:35.543212+05	\N
7529cbf9-7d44-4e8d-a39f-308e9d85b40f	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	d4156225-082e-4f0f-9b2c-85268114433a	2022-09-17 14:57:34.687184+05	2022-09-17 14:57:34.687184+05	\N
0740063b-0d3a-46f5-987c-072fdf0ad3cf	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	81b84c5d-9759-4b86-978a-649c8ef79660	2022-09-17 14:58:10.086907+05	2022-09-17 14:58:10.086907+05	\N
9dca9955-b2eb-4d2c-b4a9-ddbcfd251f56	02bd4413-8586-49ab-802e-16304e756a8b	81b84c5d-9759-4b86-978a-649c8ef79660	2022-09-17 14:58:10.098298+05	2022-09-17 14:58:10.098298+05	\N
7f22a257-7ff9-412e-bde2-8a1246454cb4	02bd4413-8586-49ab-802e-16304e756a8b	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	2022-09-17 14:59:14.122334+05	2022-09-17 14:59:14.122334+05	\N
2339b2f3-c284-4b6d-8196-7c080904a9c6	5bb9a4e7-9992-418f-b551-537844d371da	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	2022-09-17 14:59:14.133375+05	2022-09-17 14:59:14.133375+05	\N
b4afd21c-f52e-4ad1-acb7-74ae96fc9b57	b982bd86-0a0f-4950-baad-5a131e9b728e	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	2022-09-17 14:59:14.143787+05	2022-09-17 14:59:14.143787+05	\N
0b5defe3-62ca-40a2-955e-f186820d8f2c	b982bd86-0a0f-4950-baad-5a131e9b728e	8df705a5-2351-4aca-b03e-3357a23840b4	2022-09-17 15:00:15.235939+05	2022-09-17 15:00:15.235939+05	\N
e25fe89a-4c5a-4611-be9e-0007a5c631ff	f745d171-68e6-42e2-b339-cb3c210cda55	8df705a5-2351-4aca-b03e-3357a23840b4	2022-09-17 15:00:15.244281+05	2022-09-17 15:00:15.244281+05	\N
d9e3f447-0a02-4788-8ecb-ef02d266db02	d4cb1359-6c23-4194-8e3c-21ed8cec8373	8df705a5-2351-4aca-b03e-3357a23840b4	2022-09-17 15:00:15.256055+05	2022-09-17 15:00:15.256055+05	\N
c55cc88b-5651-4fa2-91fd-c85c53b2fed7	b982bd86-0a0f-4950-baad-5a131e9b728e	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	2022-10-06 11:07:41.731753+05	2022-10-06 11:07:41.731753+05	\N
9eff8362-8c57-4b7a-a8bd-695cf30d3dcc	f745d171-68e6-42e2-b339-cb3c210cda55	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	2022-10-06 11:07:41.792463+05	2022-10-06 11:07:41.792463+05	\N
d90750b0-0672-4805-a8fc-9fe46f3ff7b4	d4cb1359-6c23-4194-8e3c-21ed8cec8373	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	2022-10-06 11:07:41.802468+05	2022-10-06 11:07:41.802468+05	\N
7b67d6b8-6c31-48c7-bb98-e5645af6ad9d	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	2022-09-17 14:54:57.09628+05	2022-10-22 01:26:07.108522+05	\N
de77f0ce-1f16-4610-b62f-d8bd50339cd8	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	b2b165a3-2261-4d67-8160-0e239ecd99b5	2022-09-17 14:55:35.552501+05	2022-10-22 01:26:07.108522+05	\N
8199e176-7f47-4421-8f14-2fd116564d80	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	2022-09-17 14:56:05.486473+05	2022-10-22 01:26:07.108522+05	\N
3c4747b8-1925-43a0-959c-b7c2279f84fd	66772380-c161-4c45-9350-a45e765193e2	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	2022-09-17 14:56:05.507917+05	2022-10-22 01:26:07.108522+05	\N
401e2f12-c159-4a9f-9365-ae3a5729fa07	66772380-c161-4c45-9350-a45e765193e2	d731b17a-ae8d-4561-ad67-0f431d5c529b	2022-09-17 14:56:36.252888+05	2022-10-22 01:26:07.108522+05	\N
5b934855-c387-4039-8029-108464f90297	66772380-c161-4c45-9350-a45e765193e2	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	2022-09-17 14:57:07.276145+05	2022-10-22 01:26:07.108522+05	\N
76bd693b-0d60-4613-beed-79fee34431b0	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	d731b17a-ae8d-4561-ad67-0f431d5c529b	2022-09-17 14:56:36.263964+05	2022-10-22 01:26:07.108522+05	\N
3f5f53c8-b03d-4868-bbc1-7dfe2466000f	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	2022-09-17 14:57:07.288945+05	2022-10-22 01:26:07.108522+05	\N
9664b220-8b44-489d-a4cd-b925b2136a0a	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	d4156225-082e-4f0f-9b2c-85268114433a	2022-09-17 14:57:34.665141+05	2022-10-22 01:26:07.108522+05	\N
1d9851c7-5fe4-47c9-9d9b-b300777610dc	45765130-7f97-4f0c-b886-f70b75e02610	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	2022-09-17 14:57:07.298188+05	2022-10-22 01:26:07.108522+05	\N
8913356d-d03e-4aba-b2d7-c175c82eea3f	45765130-7f97-4f0c-b886-f70b75e02610	d4156225-082e-4f0f-9b2c-85268114433a	2022-09-17 14:57:34.676377+05	2022-10-22 01:26:07.108522+05	\N
3ab03ec7-2fb0-44e4-b5c9-c0cbf4c42fa3	45765130-7f97-4f0c-b886-f70b75e02610	81b84c5d-9759-4b86-978a-649c8ef79660	2022-09-17 14:58:10.07692+05	2022-10-22 01:26:07.108522+05	\N
599b6d48-0077-4960-86f9-947addb08210	29ed85bb-11eb-4458-bbf3-5a5644d167d6	b2b165a3-2261-4d67-8160-0e239ecd99b5	2022-09-17 14:55:35.563229+05	2022-10-22 01:26:07.073419+05	\N
c43bf0b8-a01d-41fd-9614-2e75cd19b413	29ed85bb-11eb-4458-bbf3-5a5644d167d6	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	2022-09-17 14:56:05.497232+05	2022-10-22 01:26:07.073419+05	\N
1ddb4057-86f2-436e-a0be-f09d1ff807b3	29ed85bb-11eb-4458-bbf3-5a5644d167d6	d731b17a-ae8d-4561-ad67-0f431d5c529b	2022-09-17 14:56:36.24299+05	2022-10-22 01:26:07.073419+05	\N
e46acb68-e4ad-47b1-bda2-f6155cb40488	b982bd86-0a0f-4950-baad-5a131e9b728e	442ffe07-6c0b-459d-80cd-8e12e2147568	2022-10-22 12:47:02.066115+05	2022-10-22 12:47:02.066115+05	\N
efcb9483-ce55-4d24-8bad-c33f977b5174	a4277afa-1c92-4f4e-809e-dfbb54ddbc9b	442ffe07-6c0b-459d-80cd-8e12e2147568	2022-10-22 12:47:02.066115+05	2022-10-22 12:47:02.066115+05	\N
a779244e-743d-4c17-8bcf-c5a22096a5e3	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	442ffe07-6c0b-459d-80cd-8e12e2147568	2022-10-22 12:47:02.066115+05	2022-10-22 12:47:02.066115+05	\N
a6ebfd30-4400-42bf-a0a8-65898687c676	417a385e-6a74-44f3-a536-405eb8251978	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	2022-10-24 00:14:29.387612+05	2022-10-24 00:14:29.387612+05	\N
c2754715-f455-469d-8332-22a4e7a7a8ed	02bd4413-8586-49ab-802e-16304e756a8b	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	2022-10-24 00:14:29.411764+05	2022-10-24 00:14:29.411764+05	\N
80623747-b1b4-49d6-ba3a-eb7703a830c6	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	660071e0-8f17-4c48-9d80-d4cac306de3a	2022-09-17 14:58:40.166259+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05
430bff2b-7074-4703-8496-9d3124ba45eb	02bd4413-8586-49ab-802e-16304e756a8b	660071e0-8f17-4c48-9d80-d4cac306de3a	2022-09-17 14:58:40.177106+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05
68653c84-8321-4ccf-b180-4fe3a49f2d43	5bb9a4e7-9992-418f-b551-537844d371da	660071e0-8f17-4c48-9d80-d4cac306de3a	2022-09-17 14:58:40.189647+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05
\.


--
-- Data for Name: category_shop; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category_shop (id, category_id, shop_id, created_at, updated_at, deleted_at) FROM stdin;
90d117a2-910e-4b74-a278-a722ca51a46c	b982bd86-0a0f-4950-baad-5a131e9b728e	a789cd15-3fa4-49d3-bda0-eb47bafdc61f	2022-10-24 13:23:29.189664+05	2022-10-24 13:23:29.189664+05	\N
f2e6470c-3feb-4e81-82f7-c2d0298f1a7d	cdb681a2-98e4-4716-a136-a5e4888e9c32	a789cd15-3fa4-49d3-bda0-eb47bafdc61f	2022-10-24 13:23:29.189664+05	2022-10-24 13:23:29.189664+05	\N
\.


--
-- Data for Name: company_address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_address (id, lang_id, address, created_at, updated_at, deleted_at) FROM stdin;
75706251-06ea-41c1-905f-95ed8b4132f8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Улица Азади 23, Ашхабад	2022-06-22 18:44:50.239558+05	2022-06-22 18:44:50.239558+05	\N
d2c66808-e5fe-435f-ba01-cb717f80d9e0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	adres_tm	2022-06-22 18:44:50.21776+05	2022-08-22 09:33:42.14835+05	\N
bf030883-dfe6-4836-a889-49f507de037a	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
65c97c72-20d8-4c61-8d6b-b0887aa921dd	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
1df0f34d-91b5-4a39-8ef2-496a8b5e453d	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
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
e6c8f3db-35cf-404b-a3e9-cc71b7976292	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uploads/product/40330079-629d-4146-a2b3-366e38cb2763.jpg	uploads/product/3344dd28-0808-4c4c-a3ed-ef877924f498.jpg	2022-09-17 14:59:14.077892+05	2022-09-17 14:59:14.077892+05	\N
41835ad5-e5ba-4275-8848-b4f544ade2cf	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uploads/product/50b21aa9-019c-46e5-8d36-ce79f596a5bd.jpg	uploads/product/9f1c4c2a-57f6-4f96-ad9c-8245d4773c88.jpg	2022-09-17 14:59:14.08807+05	2022-09-17 14:59:14.08807+05	\N
d96e186f-8d34-447c-814e-61ef93873f84	8df705a5-2351-4aca-b03e-3357a23840b4	uploads/product/7527c80b-aad3-440c-b605-1dc1f9b856c1.jpg	uploads/product/205684f5-bbeb-47c2-912c-4ad689674547.jpg	2022-09-17 15:00:15.188668+05	2022-09-17 15:00:15.188668+05	\N
81a68f8d-31c4-4b6b-ab15-72aea698abc3	8df705a5-2351-4aca-b03e-3357a23840b4	uploads/product/d2d8dd0c-b452-4566-b98e-a9ffab10fb13.jpg	uploads/product/a940f01a-5e38-43b6-99dc-a742e388ffe5.jpg	2022-09-17 15:00:15.200605+05	2022-09-17 15:00:15.200605+05	\N
83279012-42f0-4367-a44a-6bd59ddb249c	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uploads/product/1d54f24c-599d-46a6-ab14-c2bb565cc059.jpg	uploads/product/5fda4077-e81d-4e72-b899-f33103512859.jpg	2022-10-06 11:07:41.569383+05	2022-10-06 11:07:41.569383+05	\N
68f44447-1c5f-457c-a47a-2a281ea9c0dd	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uploads/product/8db0d212-5797-47a9-97e5-f08ceb7a293f.jpg	uploads/product/042697cc-588d-4286-a226-a94992fe942d.jpg	2022-10-06 11:07:41.646964+05	2022-10-06 11:07:41.646964+05	\N
b59f96d5-fa50-47cf-873a-dbe1a6e15302	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uploads/product/df8b0be2-7147-4ae6-8ba4-176c90aa817b.jpg	uploads/product/adca811e-07ce-4a65-9514-41a7dead8fc3.jpg	2022-09-17 14:54:57.029637+05	2022-10-22 01:26:07.108522+05	\N
2362835c-9c11-49fe-861b-81c3f29c767f	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uploads/product/0ea9466b-fd9b-4292-b3c9-0104acbaed0c.jpg	uploads/product/81d1781c-4701-40ea-91c4-4609b05de9e7.jpg	2022-09-17 14:54:57.040741+05	2022-10-22 01:26:07.108522+05	\N
ae50cccd-3a73-45ce-8fc8-b95e0f0b91bc	b2b165a3-2261-4d67-8160-0e239ecd99b5	uploads/product/a3cceed2-4af8-4017-9099-bf08231d921d.jpg	uploads/product/2d006bc4-8018-497e-9f4f-0a908fc11eaf.jpg	2022-09-17 14:55:35.496799+05	2022-10-22 01:26:07.108522+05	\N
34e6d0f9-97fc-4590-a59f-5fc2107ccdbf	b2b165a3-2261-4d67-8160-0e239ecd99b5	uploads/product/83b2761e-d5a7-44b0-886d-a2ac7ce1aeae.jpg	uploads/product/de2d464f-b447-4fe5-b33e-8db764ee90ce.jpg	2022-09-17 14:55:35.507566+05	2022-10-22 01:26:07.108522+05	\N
ff4883cb-6f14-4c18-b080-56876b24e0a9	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uploads/product/f18dea7c-1df0-48bc-a887-eb7b3b3b6f41.jpg	uploads/product/48cdeb99-d3b0-4bad-9365-01bdbcb9e88a.jpg	2022-09-17 14:56:05.441685+05	2022-10-22 01:26:07.108522+05	\N
b2964c6f-e49e-4d55-833b-7146dbe8e9f0	d731b17a-ae8d-4561-ad67-0f431d5c529b	uploads/product/fd6516a1-d696-4fcf-a5ee-89c6f5b48cb5.jpg	uploads/product/af06206e-e7de-438b-b135-70a9348eb873.jpg	2022-09-17 14:56:36.186395+05	2022-10-22 01:26:07.108522+05	\N
0f72550f-4127-421f-bf4d-83e7d5d0c667	d731b17a-ae8d-4561-ad67-0f431d5c529b	uploads/product/c148beac-f6f2-43c8-b4c4-8ddae87cf1b8.jpg	uploads/product/2975a2c4-f978-40fa-9798-b30a71c8c908.jpg	2022-09-17 14:56:36.197707+05	2022-10-22 01:26:07.108522+05	\N
b1575796-5322-4983-ae21-bb84725a4f75	81b84c5d-9759-4b86-978a-649c8ef79660	uploads/product/1ff26530-a022-4a3f-90e9-8c4a3a2d5f2e.jpg	uploads/product/10b3062a-71cc-49b3-9eae-4410745a7685.jpg	2022-09-17 14:58:10.031274+05	2022-10-22 01:26:07.108522+05	\N
f8c2dd64-75f1-4617-8237-22aa6fde2faa	d4156225-082e-4f0f-9b2c-85268114433a	uploads/product/fe06f4f8-13da-49d5-aeec-4a05bd9155c4.jpg	uploads/product/5db94dda-c1f5-4d99-a0f1-a005be127383.jpg	2022-09-17 14:57:34.620044+05	2022-10-22 02:53:49.519217+05	\N
394d0cde-5f52-4228-a615-574acec7e51a	d4156225-082e-4f0f-9b2c-85268114433a	uploads/product/c29e7734-e8cd-453f-8583-74f77b76f425.jpg	uploads/product/ac6306f5-0193-413f-9838-7daed09d7fb5.jpg	2022-09-17 14:57:34.631937+05	2022-10-22 02:53:49.519217+05	\N
06073ec4-b58e-4fba-877d-6aae1339d084	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uploads/product/3e6a4832-663f-4b21-b3e4-f637245590a2.jpg	uploads/product/b37e6ed9-401c-465f-9d50-f990b74a9e93.jpg	2022-09-17 14:56:05.452041+05	2022-10-22 01:26:07.108522+05	\N
21f8e93a-ea3a-4596-899a-1982d8956dcb	81b84c5d-9759-4b86-978a-649c8ef79660	uploads/product/3bebc440-4f7f-4c34-abd1-c7151f41823d.jpg	uploads/product/1cfabb12-a0ce-4a11-b937-149f03fc95c2.jpg	2022-09-17 14:58:10.042467+05	2022-10-22 01:26:07.108522+05	\N
541a60b2-e164-45f1-8238-b430f9a4f696	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uploads/product/f9484719-a04c-482a-b9cb-d4477d67171e.jpg	uploads/product/56d9c5d4-89ac-45e7-80a5-d4420a1ba1e4.jpg	2022-09-17 14:57:07.232409+05	2022-10-22 01:26:07.108522+05	\N
01234b03-8ebe-4bd9-9296-5411d8b602d6	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uploads/product/f8a3e4a8-1363-49de-95d8-7ea99578bb36.jpg	uploads/product/7d443731-50ee-4e9b-9552-562282763df1.jpg	2022-09-17 14:57:07.242703+05	2022-10-22 01:26:07.108522+05	\N
92154468-15dd-4dbd-a968-7c11c40ad77e	939da5d3-2f7a-40e2-b133-0b4113280647	uploads/product/96c93bcd-849c-4175-b0d5-065a0b793aee.jpg	uploads/product/ef676f34-1389-4c7a-828e-9010ffd40813.jpg	2022-10-22 12:46:26.498314+05	2022-10-22 12:46:26.498314+05	\N
27b669b4-152f-44ab-9d5a-fa7982e66f9f	939da5d3-2f7a-40e2-b133-0b4113280647	uploads/product/f8950285-3184-44ca-92ac-483cafeb0b55.jpg	uploads/product/f6a2f688-e34a-4ff0-8a6d-e279c08b9d77.jpg	2022-10-22 12:46:26.498314+05	2022-10-22 12:46:26.498314+05	\N
1187d914-4377-46ab-a25e-056d3c5859d3	442ffe07-6c0b-459d-80cd-8e12e2147568	uploads/product/5adbf006-2791-4308-8719-d39d4c7ac530.jpg	uploads/product/c4f1e04c-e952-4b25-8cd5-ef02d5beb1c7.jpg	2022-10-22 12:47:01.988319+05	2022-10-22 12:47:01.988319+05	\N
00f0ad3d-b75a-4d55-825b-68f19f1c2d3d	442ffe07-6c0b-459d-80cd-8e12e2147568	uploads/product/ec0d3420-dcf1-466c-8b4a-837148f2423f.jpg	uploads/product/b818d70d-4cd1-4f2a-a3dc-2c1237c340af.jpg	2022-10-22 12:47:01.988319+05	2022-10-22 12:47:01.988319+05	\N
f2763091-1ad0-43da-91dd-3af7c29bffce	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	uploads/product/2321f997-3a60-4ef0-8205-e59040557909.jpg	uploads/product/aef46b1a-cb63-4066-8289-ea5705d79d30.jpg	2022-10-24 00:14:29.227295+05	2022-10-24 00:14:29.227295+05	\N
eb33f9cc-e492-4efb-b0cc-2cefa9f8f5cc	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	uploads/product/527b0dae-1759-4caf-9f5d-91b0195db4e9.jpg	uploads/product/8ae83a9c-c402-429f-b858-1a228d8cb1d8.jpg	2022-10-24 00:14:29.251181+05	2022-10-24 00:14:29.251181+05	\N
843515aa-4fef-4ef3-9570-61bf600a9357	660071e0-8f17-4c48-9d80-d4cac306de3a	uploads/product/d629fe21-e9a8-4c39-941d-ea222b0ce204.jpg	uploads/product/2e752914-f145-4f09-a35d-272db51b3083.jpg	2022-09-17 14:58:40.120625+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05
2957da00-07b2-4f26-9504-44f1c50da3e8	660071e0-8f17-4c48-9d80-d4cac306de3a	uploads/product/0b9ffad0-fcd9-4a36-b661-15bcff855593.jpg	uploads/product/fc1d5b8d-bd84-411e-af20-b5a306966579.jpg	2022-09-17 14:58:40.131788+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05
\.


--
-- Data for Name: languages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.languages (id, name_short, flag, created_at, updated_at, deleted_at) FROM stdin;
aea98b93-7bdf-455b-9ad4-a259d69dc76e	ru	uploads/language1c24e3a6-173e-4264-a631-f099d15495dd.jpeg	2022-06-15 19:53:21.29491+05	2022-06-15 19:53:21.29491+05	\N
8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	tm	uploads/language17b99bd1-f52d-41db-b4e6-1ecff03e0fd0.jpeg	2022-06-15 19:53:06.041686+05	2022-10-16 18:53:27.82538+05	\N
55a387df-6d38-42ea-bfba-379327b53cbd	fr	uploads/language/3535a022-0d14-4030-9658-1a720798ce03.jpg	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.771213+05	\N
198695b5-579a-4f80-ac10-8380e17e5d98	tr	uploads/language/54ebb99b-f894-4540-b75d-1e9dde5b8007.jpg	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
b62a1c1c-0a29-4756-8e9d-5c9680758d18	pl	uploads/language/uploads/language/5726efb8-3c34-4f83-b39a-cf68ce04acc3.jpg	2022-10-20 01:44:26.912355+05	2022-10-20 01:46:04.093537+05	\N
\.


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.likes (id, product_id, customer_id, created_at, updated_at, deleted_at) FROM stdin;
d45b89a8-6676-412d-bd6c-a4e655aa41b8	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	f25a66d3-93ac-4da4-b237-d34867d5ca8f	2022-10-17 00:50:48.875499+05	2022-10-17 00:50:48.875499+05	\N
a14b8219-81d6-4f2c-999b-d6ec8ef7fece	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	f25a66d3-93ac-4da4-b237-d34867d5ca8f	2022-10-17 00:50:48.957722+05	2022-10-17 00:50:48.957722+05	\N
ded6953a-f3ae-4463-82e5-998f114abf8d	b2b165a3-2261-4d67-8160-0e239ecd99b5	f25a66d3-93ac-4da4-b237-d34867d5ca8f	2022-10-22 02:39:51.17081+05	2022-10-22 02:39:51.17081+05	\N
\.


--
-- Data for Name: main_image; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.main_image (id, product_id, small, medium, large, created_at, updated_at, deleted_at) FROM stdin;
489304cb-a16a-4f78-841e-797b341f224b	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uploads/product/139044f5-bb94-4f31-9b5a-c7c28056398f.jpg	uploads/product/a45a8c58-bc17-44ad-807d-719607bdd031.jpg	uploads/product/7d873556-bfe5-4991-8cc7-0ab6609eb45e.jpg	2022-09-17 14:59:14.066244+05	2022-09-17 14:59:14.066244+05	\N
24fdc6b1-afb2-4735-b406-70addc0dd8d9	8df705a5-2351-4aca-b03e-3357a23840b4	uploads/product/c0a88fd0-2374-49af-95d0-5c692c626b94.jpg	uploads/product/8a67c3a7-94fc-4831-9f90-80ae409c684f.jpg	uploads/product/c7a5b0ce-d9fc-449c-b039-726d716c62a7.jpg	2022-09-17 15:00:15.178822+05	2022-09-17 15:00:15.178822+05	\N
ea8b26f1-1e34-4587-9b72-3747e29930cf	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uploads/product/5938e42c-700a-48c3-8461-ad3df445dac1.jpg	uploads/product/9ac3b3f6-6d4b-4be0-8a98-4d1b10f4d954.jpg	uploads/product/5516cc1d-8f21-4006-b353-b41518ef1bef.jpg	2022-10-06 11:07:41.522494+05	2022-10-06 11:07:41.522494+05	\N
af383593-cacb-4440-8144-4560c1887921	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uploads/product/0bb0fc0e-6b9a-48cb-a005-067818339a9a.jpg	uploads/product/a07bf6d8-837d-4743-9752-3a5a2b178b4d.jpg	uploads/product/3d48f030-4edb-4510-af2f-4688667f4e0b.jpg	2022-09-17 14:54:57.018537+05	2022-10-22 01:26:07.108522+05	\N
045bb7ae-6d64-4366-acdd-7f342a7600a2	b2b165a3-2261-4d67-8160-0e239ecd99b5	uploads/product/a7c19215-5db7-4ebb-b373-a8b58b678def.jpg	uploads/product/58fae77e-a7d4-4f4b-bdfd-399e1e7994e8.jpg	uploads/product/22c2eb3b-cf3e-4369-b08f-92a6f8665bdd.jpg	2022-09-17 14:55:35.485706+05	2022-10-22 01:26:07.108522+05	\N
85d25308-e2a0-40f7-ab57-4dd081e59ed8	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uploads/product/061bb6bc-d603-4a28-8737-246018133728.jpg	uploads/product/1e562ca8-31bf-4712-af03-feb1654bf1f1.jpg	uploads/product/cba805fa-8c99-4992-9315-60071832bb59.jpg	2022-09-17 14:56:05.43168+05	2022-10-22 01:26:07.108522+05	\N
0b9cc77d-87fd-4603-bf2d-e7203adeb4e8	d731b17a-ae8d-4561-ad67-0f431d5c529b	uploads/product/a2d59c2f-e312-4541-b4a7-79fc679c5e0c.jpg	uploads/product/87aa78b0-6f8a-4b37-a7f6-790e77398e8b.jpg	uploads/product/5d7e8980-f050-4936-9198-eaa3f3236370.jpg	2022-09-17 14:56:36.174972+05	2022-10-22 01:26:07.108522+05	\N
7ab33a3b-0195-4024-a035-e54268762d3b	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uploads/product/bf635780-dc34-4478-b936-8e073eefa79e.jpg	uploads/product/ea2ce473-4110-4129-8d5e-0fa456d10530.jpg	uploads/product/a92d33a6-8180-4355-8f4b-bd23079177a4.jpg	2022-09-17 14:57:07.221569+05	2022-10-22 01:26:07.108522+05	\N
286f3f21-8750-499f-b35c-9de67b236316	81b84c5d-9759-4b86-978a-649c8ef79660	uploads/product/bc78c4bb-35f7-4873-88eb-af1c985e9f34.jpg	uploads/product/356b93e2-be40-4ed0-bf77-b229210de3e7.jpg	uploads/product/d7641dc6-6e2a-4d21-9403-98c872f93b25.jpg	2022-09-17 14:58:10.020986+05	2022-10-22 01:26:07.108522+05	\N
7982b9f8-bfca-42f5-b13a-37b1ed41e1a2	d4156225-082e-4f0f-9b2c-85268114433a	uploads/product/706c3355-01f9-4648-850c-eacc3557f435.jpg	uploads/product/e585a68a-d6c3-4a1f-8660-041cade9899e.jpg	uploads/product/03902970-8446-4180-9200-902a6aa7fa23.jpg	2022-09-17 14:57:34.609435+05	2022-10-22 02:53:49.519217+05	\N
df561c2b-1eda-4ca0-ada8-9c2b930ab0bd	939da5d3-2f7a-40e2-b133-0b4113280647	uploads/product/fa0e1e4a-3f97-4f1c-b0db-d37381c27895.jpg	uploads/product/c9b756d4-7854-4dea-babb-550e3f12d878.jpg	uploads/product/d99cbd21-0cbe-4b3d-8192-63f2878b82fb.jpg	2022-10-22 12:46:26.484513+05	2022-10-22 12:46:26.484513+05	\N
b1a5fb1e-1b93-4a32-899a-8480f0c863be	442ffe07-6c0b-459d-80cd-8e12e2147568	uploads/product/67266c0a-61c6-408f-932e-f326e8159664.jpg	uploads/product/a1734eec-d099-46a1-82c5-755e800be6a0.jpg	uploads/product/76732327-cabf-43ce-933a-35f7e01e7d09.jpg	2022-10-22 12:47:01.973883+05	2022-10-22 12:47:01.973883+05	\N
3141d941-c1fa-41f0-b542-44090d4ba2a1	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	uploads/product/7afdb808-d815-4a44-ae9b-6715ecd98518.jpg	uploads/product/7bebeee4-7151-4afb-b70e-0d410ca51c4a.jpg	uploads/product/19b45018-2178-42d5-9d55-d63df7297f3c.jpg	2022-09-17 14:59:44.866884+05	2022-10-24 00:14:29.289334+05	\N
9ddabe0c-bad0-493e-b084-a5d6d96e894d	660071e0-8f17-4c48-9d80-d4cac306de3a	uploads/product/64e1692d-bdd3-4081-88d5-74c8cc71ae51.jpg	uploads/product/74278164-f5c5-423f-8d84-ba7a122a8171.jpg	uploads/product/018a0b4a-8593-4287-93b3-37aaa1a04f0f.jpg	2022-09-17 14:58:40.111413+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05
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
f4290ceb-9743-448e-abc8-a7a39b784664	8df705a5-2351-4aca-b03e-3357a23840b4	1	ed6f4a78-b4fa-4035-a8e0-d884c65a5889	2022-10-24 02:21:24.275457+05	2022-10-24 02:21:24.275457+05	\N
a4e19672-ec52-4d51-9175-563d4e86f6e1	b2b165a3-2261-4d67-8160-0e239ecd99b5	3	ed6f4a78-b4fa-4035-a8e0-d884c65a5889	2022-10-24 02:21:24.293498+05	2022-10-24 02:21:24.293498+05	\N
02ec1f62-996f-4ff2-8497-b4da23cfdce0	8df705a5-2351-4aca-b03e-3357a23840b4	1	2698cad6-a4f2-493d-b86d-d3d1659c1896	2022-10-24 02:23:26.457082+05	2022-10-24 02:23:26.457082+05	\N
c72e2d77-4b2c-4cdc-8fae-8da97845ab7e	b2b165a3-2261-4d67-8160-0e239ecd99b5	3	2698cad6-a4f2-493d-b86d-d3d1659c1896	2022-10-24 02:23:26.478042+05	2022-10-24 02:23:26.478042+05	\N
c8a0cd47-4056-41d7-9a3c-34bdc7674014	d4156225-082e-4f0f-9b2c-85268114433a	25	f767105b-51b9-47a2-a161-1646934f5acd	2022-10-24 11:16:38.481312+05	2022-10-24 11:16:38.481312+05	\N
671aa93d-dbbb-4770-b895-f537aa605f11	81b84c5d-9759-4b86-978a-649c8ef79660	3	f767105b-51b9-47a2-a161-1646934f5acd	2022-10-24 11:16:38.492814+05	2022-10-24 11:16:38.492814+05	\N
cc02eb93-96ab-4176-b072-380626ef1be5	8df705a5-2351-4aca-b03e-3357a23840b4	3	f767105b-51b9-47a2-a161-1646934f5acd	2022-10-24 11:16:38.503666+05	2022-10-24 11:16:38.503666+05	\N
974bc5f8-a43f-4316-8ff8-ba9ab05f18bc	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	3	f767105b-51b9-47a2-a161-1646934f5acd	2022-10-24 11:16:38.514362+05	2022-10-24 11:16:38.514362+05	\N
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, customer_id, customer_mark, order_time, payment_type, total_price, created_at, updated_at, deleted_at, order_number) FROM stdin;
ed6f4a78-b4fa-4035-a8e0-d884c65a5889	fec81bff-8264-403f-b2ba-58afd53c821b	ajkedbhewj	18:00 - 21:00	nagt_tm	845	2022-10-24 02:21:24.249433+05	2022-10-24 02:21:24.249433+05	\N	22
2698cad6-a4f2-493d-b86d-d3d1659c1896	655c5504-1547-4daf-abe0-80b4116684f0	jenew	18:00 - 21:00	nagt_tm	845	2022-10-24 02:23:26.441517+05	2022-10-24 02:23:26.441517+05	\N	23
f767105b-51b9-47a2-a161-1646934f5acd	1f5bc917-fc85-46cf-a1b2-7c14cfe940be	isleg market bet cykypdyr	12:00 - 16:00	nagt	1788	2022-10-24 11:16:38.47027+05	2022-10-24 11:16:38.47027+05	\N	24
\.


--
-- Data for Name: payment_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_types (id, lang_id, type, created_at, updated_at, deleted_at) FROM stdin;
83e6589c-0cb6-4267-bcc5-e06cc93b36d8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	наличные	2022-09-20 14:33:50.780468+05	2022-09-20 14:33:50.780468+05	\N
7a6a313d-8fcd-4c56-9fa5-aefb12552b82	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	töleg terminaly	2022-09-20 14:34:46.329459+05	2022-09-20 14:34:46.329459+05	\N
cb7e8cc9-9b2e-4cd8-921f-91b3bb5e5564	aea98b93-7bdf-455b-9ad4-a259d69dc76e	платежный терминал	2022-09-20 14:34:46.359276+05	2022-09-20 14:34:46.359276+05	\N
38696743-82e5-4644-9c86-4a99ae45f912	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	nagt_tm	2022-09-20 14:33:50.755689+05	2022-09-20 14:40:04.959827+05	\N
c188243d-a553-4fd3-ae05-cf8db9beb43e	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
29ec9186-1ccf-4c9c-b4fa-e0abc4b45291	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
f55dc383-4e8e-4e78-9cd6-981bf79cf925	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, brend_id, price, old_price, amount, product_code, created_at, updated_at, deleted_at, limit_amount, is_new) FROM stdin;
c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	fdd259c2-794a-42b9-a3ad-9e91502af23e	65	68	1000	151fwe51we	2022-09-17 14:59:14.037001+05	2022-09-17 14:59:14.037001+05	\N	100	f
8df705a5-2351-4aca-b03e-3357a23840b4	46b13f0a-d584-4ad3-b270-437ecdc51449	65	68	1000	151fwe51we	2022-09-17 15:00:15.148583+05	2022-09-17 15:00:15.148583+05	\N	100	f
3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	46b13f0a-d584-4ad3-b270-437ecdc51449	67	68	1000	151fwe51we	2022-10-06 11:07:41.410248+05	2022-10-06 11:07:41.410248+05	\N	100	f
0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	214be879-65c3-4710-86b4-3fc3bce2e974	65	68	1000	151fwe51we	2022-09-17 14:54:56.989242+05	2022-10-22 01:26:07.108522+05	\N	100	f
b2b165a3-2261-4d67-8160-0e239ecd99b5	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	65	68	1000	151fwe51we	2022-09-17 14:55:35.441733+05	2022-10-22 01:26:07.108522+05	\N	100	f
a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	fdd259c2-794a-42b9-a3ad-9e91502af23e	65	68	1000	151fwe51we	2022-09-17 14:56:05.406905+05	2022-10-22 01:26:07.108522+05	\N	100	f
d731b17a-ae8d-4561-ad67-0f431d5c529b	f53a27b4-7810-4d8f-bd45-edad405d92b9	65	68	1000	151fwe51we	2022-09-17 14:56:36.153769+05	2022-10-22 01:26:07.108522+05	\N	100	f
bb6c3bdb-79e2-44b3-98b1-c1cee0976777	46b13f0a-d584-4ad3-b270-437ecdc51449	65	68	1000	151fwe51we	2022-09-17 14:57:07.191142+05	2022-10-22 01:26:07.108522+05	\N	100	f
81b84c5d-9759-4b86-978a-649c8ef79660	214be879-65c3-4710-86b4-3fc3bce2e974	65	68	1000	151fwe51we	2022-09-17 14:58:09.998335+05	2022-10-22 01:26:07.108522+05	\N	100	f
d4156225-082e-4f0f-9b2c-85268114433a	c4bcda34-7332-4ae5-8129-d7538d63fee4	65	68	1000	151fwe51we	2022-09-17 14:57:34.582228+05	2022-10-22 02:53:49.519217+05	\N	100	f
939da5d3-2f7a-40e2-b133-0b4113280647	fdd259c2-794a-42b9-a3ad-9e91502af23e	67	68	1000	151fwe51we	2022-10-22 12:46:26.452543+05	2022-10-22 12:46:26.452543+05	\N	100	f
442ffe07-6c0b-459d-80cd-8e12e2147568	fdd259c2-794a-42b9-a3ad-9e91502af23e	67	68	1000	151fwe51we	2022-10-22 12:47:01.951588+05	2022-10-22 12:47:01.951588+05	\N	100	f
e3c33ead-3c30-40f1-9d28-7bb8b71b767f	46b13f0a-d584-4ad3-b270-437ecdc51449	46	46.5	998	w5f2we6f598	2022-09-17 14:59:44.837302+05	2022-10-24 00:14:29.271173+05	\N	990	t
660071e0-8f17-4c48-9d80-d4cac306de3a	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	65	68	1000	151fwe51we	2022-09-17 14:58:40.084476+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05	100	f
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
8f802660-b581-41c9-8e08-77adf0c8d9d7	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
0d2a5c3c-8e0a-492e-bcb2-ac80d5038364	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
8a6fed6c-c718-4d1b-9eb4-a0fdcf709210	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
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
806a5f9a-7882-46f7-bd0d-a3f4cc24fb6e	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N	uytget
8e1d1766-a39e-480d-aa23-8a1fe477ad69	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N	uytget
6dcc6257-cc87-44bf-8bb4-d46560867f34	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N	uytget
\.


--
-- Data for Name: translation_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_category (id, lang_id, category_id, name, created_at, updated_at, deleted_at) FROM stdin;
21520180-13e2-4c2b-a5f9-866c2e59ba87	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f745d171-68e6-42e2-b339-cb3c210cda55	Kiçi paket kofeler	2022-06-16 13:45:48.889727+05	2022-06-16 13:45:48.889727+05	\N
ee2f97fb-8c6c-4e38-bdb3-bf769bc95d3b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d4cb1359-6c23-4194-8e3c-21ed8cec8373	Batonçikler	2022-06-16 13:48:04.581888+05	2022-06-16 13:48:04.581888+05	\N
ab35a97a-dfd1-4100-8e84-d34e74e9a02e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f745d171-68e6-42e2-b339-cb3c210cda55	Кофе в пакетиках	2022-06-16 13:45:48.906024+05	2022-06-16 13:45:48.906024+05	\N
ea104eaf-c3fd-4f2d-88bf-dffc14d48dc5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d4cb1359-6c23-4194-8e3c-21ed8cec8373	Батончики	2022-06-16 13:48:04.597499+05	2022-06-16 13:48:04.597499+05	\N
4eef5d40-9aad-4101-b36b-9026dd3dfb51	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b982bd86-0a0f-4950-baad-5a131e9b728e	name_tm	2022-06-16 13:44:16.499713+05	2022-06-16 13:44:16.499713+05	\N
10a8b5ec-a3ca-448d-975b-83b3a7a8c0d2	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b982bd86-0a0f-4950-baad-5a131e9b728e	name_ru	2022-06-16 13:44:16.515874+05	2022-06-16 13:44:16.515874+05	\N
4eb6bcbf-91f2-4505-a27e-cc3f96f2b829	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	Plitkalar	2022-06-16 13:47:18.888998+05	2022-06-16 13:47:18.888998+05	\N
53fb44c7-45fb-49f0-a433-aaed23b2dfc0	aea98b93-7bdf-455b-9ad4-a259d69dc76e	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	Плитки	2022-06-16 13:47:18.942159+05	2022-06-16 13:47:18.942159+05	\N
bff34c21-04c1-4cea-bfaf-c8f9ce7e2bfe	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	02bd4413-8586-49ab-802e-16304e756a8b	name_tm	2022-06-16 13:43:22.674562+05	2022-06-16 13:43:22.674562+05	\N
0e400414-a80c-449d-8842-dd6667b45c73	aea98b93-7bdf-455b-9ad4-a259d69dc76e	02bd4413-8586-49ab-802e-16304e756a8b	name_ru	2022-06-16 13:43:22.681932+05	2022-06-16 13:43:22.681932+05	\N
e099e7f6-1b97-4f70-8f29-f586ab6697d0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	5bb9a4e7-9992-418f-b551-537844d371da	Şokolad we Keksler	2022-06-16 13:46:44.657849+05	2022-06-16 13:46:44.657849+05	\N
415a0711-2482-44b3-8f03-923dca28bd5d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	5bb9a4e7-9992-418f-b551-537844d371da	Шоколады и Кексы	2022-06-16 13:46:44.673892+05	2022-06-16 13:46:44.673892+05	\N
1c287f79-c467-4530-aafd-77c294ac4091	55a387df-6d38-42ea-bfba-379327b53cbd	f745d171-68e6-42e2-b339-cb3c210cda55	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
c66cc626-d744-487d-8fc6-87ebc6a19535	55a387df-6d38-42ea-bfba-379327b53cbd	d4cb1359-6c23-4194-8e3c-21ed8cec8373	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
63cd3daa-d15e-4bcb-b08f-f088e8ebded8	55a387df-6d38-42ea-bfba-379327b53cbd	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
53b2031b-a0da-4853-8d98-d4630b1f64b1	55a387df-6d38-42ea-bfba-379327b53cbd	02bd4413-8586-49ab-802e-16304e756a8b	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
7f55b5a8-fbf2-4266-a55a-6c2a52b8b63a	55a387df-6d38-42ea-bfba-379327b53cbd	5bb9a4e7-9992-418f-b551-537844d371da	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
77950afa-46c3-41ba-a1f5-98ef3511ad5c	55a387df-6d38-42ea-bfba-379327b53cbd	b982bd86-0a0f-4950-baad-5a131e9b728e	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
de41f1cc-4429-4c88-a10c-14f42dc568b8	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7605172f-7a12-4781-a892-6e3b5cf11490	Name_tm	2022-10-18 11:09:26.77206+05	2022-10-18 11:09:26.77206+05	\N
d71a47b2-f7c4-4bae-8fca-8c945579e09b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7605172f-7a12-4781-a892-6e3b5cf11490	Name_ru	2022-10-18 11:09:26.786998+05	2022-10-18 11:09:26.786998+05	\N
1423ae8b-850b-47f3-a6fb-93ea15648405	55a387df-6d38-42ea-bfba-379327b53cbd	7605172f-7a12-4781-a892-6e3b5cf11490	NAME_FR	2022-10-18 11:09:26.79722+05	2022-10-18 11:09:26.79722+05	\N
98a6ba52-a644-4069-b03c-0a6bf6388ddd	198695b5-579a-4f80-ac10-8380e17e5d98	f745d171-68e6-42e2-b339-cb3c210cda55	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
f819bc8c-792c-4688-baf5-50fb1531ad2e	198695b5-579a-4f80-ac10-8380e17e5d98	d4cb1359-6c23-4194-8e3c-21ed8cec8373	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
b3f670f8-e22d-472e-bb2f-5c5a548e29e9	198695b5-579a-4f80-ac10-8380e17e5d98	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
3995129d-88f4-4085-9b2a-e21932161f23	198695b5-579a-4f80-ac10-8380e17e5d98	02bd4413-8586-49ab-802e-16304e756a8b	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
d5c65d5b-6769-4c06-93c7-b78ad123a924	198695b5-579a-4f80-ac10-8380e17e5d98	5bb9a4e7-9992-418f-b551-537844d371da	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
243b63ec-6959-4ba4-a323-0963180eaebd	198695b5-579a-4f80-ac10-8380e17e5d98	b982bd86-0a0f-4950-baad-5a131e9b728e	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
7b643341-1396-4f55-9dc2-aaf14f63572d	198695b5-579a-4f80-ac10-8380e17e5d98	7605172f-7a12-4781-a892-6e3b5cf11490	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
10ada4ab-ead3-4123-8825-c7b4be23c2c4	b62a1c1c-0a29-4756-8e9d-5c9680758d18	f745d171-68e6-42e2-b339-cb3c210cda55	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
4a713859-2a34-43ad-9570-fa7f7547b41d	b62a1c1c-0a29-4756-8e9d-5c9680758d18	d4cb1359-6c23-4194-8e3c-21ed8cec8373	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
e3a7ce6e-ad96-4d53-a9c8-1edd7d3f6123	b62a1c1c-0a29-4756-8e9d-5c9680758d18	fdc10d33-043b-4ee0-9d6e-e2a12a3e150a	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
199363e0-6044-44f9-8b11-7293b509e0cb	b62a1c1c-0a29-4756-8e9d-5c9680758d18	02bd4413-8586-49ab-802e-16304e756a8b	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
7e39ab42-0e0d-495a-b6e2-6d80503236bd	b62a1c1c-0a29-4756-8e9d-5c9680758d18	5bb9a4e7-9992-418f-b551-537844d371da	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
04b695dc-145e-4b00-b08a-6644c587bb18	b62a1c1c-0a29-4756-8e9d-5c9680758d18	b982bd86-0a0f-4950-baad-5a131e9b728e	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
ad15d99a-a266-4bf9-bbce-54deaa9e6e4e	b62a1c1c-0a29-4756-8e9d-5c9680758d18	7605172f-7a12-4781-a892-6e3b5cf11490	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
78aef3f3-a0e8-4e55-8b44-2e0005f2cad9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	8f6af238-80ef-40f0-8c34-531b9c06b373	Name_tm	2022-10-21 03:21:08.009168+05	2022-10-21 03:21:08.009168+05	\N
17d71716-2d15-47b2-9775-84bb60989bfa	aea98b93-7bdf-455b-9ad4-a259d69dc76e	8f6af238-80ef-40f0-8c34-531b9c06b373	Name_ru	2022-10-21 03:21:08.037577+05	2022-10-21 03:21:08.037577+05	\N
9aff7fdc-323a-41cc-a853-f79f95734af8	55a387df-6d38-42ea-bfba-379327b53cbd	8f6af238-80ef-40f0-8c34-531b9c06b373	NAME_FR	2022-10-21 03:21:08.062555+05	2022-10-21 03:21:08.062555+05	\N
7593a61a-57d8-4617-aa07-e23a626b4351	198695b5-579a-4f80-ac10-8380e17e5d98	8f6af238-80ef-40f0-8c34-531b9c06b373	name_tr	2022-10-21 03:21:08.084635+05	2022-10-21 03:21:08.084635+05	\N
25cf06bf-4829-482f-b0a6-212df761fadc	b62a1c1c-0a29-4756-8e9d-5c9680758d18	8f6af238-80ef-40f0-8c34-531b9c06b373	name_pl	2022-10-21 03:21:08.105367+05	2022-10-21 03:21:08.105367+05	\N
905b8603-1bb8-4bb1-b79d-9646aa1f3cfe	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	417a385e-6a74-44f3-a536-405eb8251978	Name_tm	2022-10-21 03:24:33.340982+05	2022-10-21 03:24:33.340982+05	\N
84661578-b7d0-4468-aede-6cf523470f56	aea98b93-7bdf-455b-9ad4-a259d69dc76e	417a385e-6a74-44f3-a536-405eb8251978	Name_ru	2022-10-21 03:24:33.359312+05	2022-10-21 03:24:33.359312+05	\N
bb7cf0c4-e8bc-47db-953a-4e35b61b1b02	55a387df-6d38-42ea-bfba-379327b53cbd	417a385e-6a74-44f3-a536-405eb8251978	NAME_FR	2022-10-21 03:24:33.382628+05	2022-10-21 03:24:33.382628+05	\N
a60eb529-6cec-4b10-8f87-1ecd93ff1b55	198695b5-579a-4f80-ac10-8380e17e5d98	417a385e-6a74-44f3-a536-405eb8251978	name_tr	2022-10-21 03:24:33.40283+05	2022-10-21 03:24:33.40283+05	\N
7200b5d7-fa0d-41f5-86bd-649c6bcd67c8	b62a1c1c-0a29-4756-8e9d-5c9680758d18	417a385e-6a74-44f3-a536-405eb8251978	name_pl	2022-10-21 03:24:33.413325+05	2022-10-21 03:24:33.413325+05	\N
598d59a5-9964-476e-a659-afe8367efb20	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b1bae1ce-4295-4268-bf2d-71c8761e5679	Name_tm	2022-10-21 03:26:21.900843+05	2022-10-21 03:26:21.900843+05	\N
db035c85-6e20-439b-ad1a-462f1f3e9a78	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b1bae1ce-4295-4268-bf2d-71c8761e5679	Name_ru	2022-10-21 03:26:21.916767+05	2022-10-21 03:26:21.916767+05	\N
75c6459d-bc2e-4da0-b941-fe1364d07197	55a387df-6d38-42ea-bfba-379327b53cbd	b1bae1ce-4295-4268-bf2d-71c8761e5679	NAME_FR	2022-10-21 03:26:21.930256+05	2022-10-21 03:26:21.930256+05	\N
c26ca683-a991-452d-a734-316e9d98959b	198695b5-579a-4f80-ac10-8380e17e5d98	b1bae1ce-4295-4268-bf2d-71c8761e5679	name_tr	2022-10-21 03:26:21.937746+05	2022-10-21 03:26:21.937746+05	\N
c1a3bf99-9bb3-46f7-b883-5be9dbca1564	b62a1c1c-0a29-4756-8e9d-5c9680758d18	b1bae1ce-4295-4268-bf2d-71c8761e5679	name_pl	2022-10-21 03:26:21.950506+05	2022-10-21 03:26:21.950506+05	\N
ecbe74d3-aea6-4209-85e5-a0ac9beac84d	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d8ded28c-d4fb-4c11-a84c-4d4f81a22e28	Name_tm	2022-10-21 03:28:14.66099+05	2022-10-21 03:28:14.66099+05	\N
3bbec4b4-693f-4079-be36-154c74feab05	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d8ded28c-d4fb-4c11-a84c-4d4f81a22e28	Name_ru	2022-10-21 03:28:14.679857+05	2022-10-21 03:28:14.679857+05	\N
8a91bcb0-fcce-4a4f-80ff-a2896c0cc36a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	Arzanladyşdaky harytlar	2022-06-20 09:43:07.368782+05	2022-10-22 01:26:07.108522+05	\N
ce573dfd-6af8-4e64-8260-8746a090acd7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	Продукция со скидкой	2022-06-20 09:43:07.377729+05	2022-10-22 01:26:07.108522+05	\N
49c2be64-bb43-4696-8af6-988187f99466	55a387df-6d38-42ea-bfba-379327b53cbd	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N
2c625f79-ad57-48ae-a87b-67f69d947d41	198695b5-579a-4f80-ac10-8380e17e5d98	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N
44ccdd47-b8e3-46a9-827d-9e7e4ab9fe8e	b62a1c1c-0a29-4756-8e9d-5c9680758d18	7f453dd0-7b2e-480d-a8be-fcfa23bd863e	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N
0451777f-8681-4d54-b562-5aea3bfddb22	55a387df-6d38-42ea-bfba-379327b53cbd	d8ded28c-d4fb-4c11-a84c-4d4f81a22e28	NAME_FR	2022-10-21 03:28:14.700081+05	2022-10-21 03:28:14.700081+05	\N
78b0a57d-7838-4d9c-9f8c-26fbbf4edfe8	198695b5-579a-4f80-ac10-8380e17e5d98	d8ded28c-d4fb-4c11-a84c-4d4f81a22e28	name_tr	2022-10-21 03:28:14.720587+05	2022-10-21 03:28:14.720587+05	\N
0aa92e51-c95b-443e-ad7e-6ca2c043f890	b62a1c1c-0a29-4756-8e9d-5c9680758d18	d8ded28c-d4fb-4c11-a84c-4d4f81a22e28	name_pl	2022-10-21 03:28:14.73378+05	2022-10-21 03:28:14.73378+05	\N
00578b2c-c394-46c1-9dd3-1a99ce6e8efb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	cdb681a2-98e4-4716-a136-a5e4888e9c32	name_tm	2022-10-21 03:28:30.772796+05	2022-10-21 11:09:53.910704+05	\N
40d0bb2d-93e0-41d0-9ee1-7dac06cf6500	aea98b93-7bdf-455b-9ad4-a259d69dc76e	cdb681a2-98e4-4716-a136-a5e4888e9c32	name_ru	2022-10-21 03:28:30.792425+05	2022-10-21 11:09:53.922362+05	\N
0e7ebc37-c1fe-45cf-a074-6e55b7f52507	55a387df-6d38-42ea-bfba-379327b53cbd	cdb681a2-98e4-4716-a136-a5e4888e9c32	name_fr	2022-10-21 03:28:30.810261+05	2022-10-21 11:09:53.934863+05	\N
21630f55-c6a6-40c7-8d60-24339fe36835	198695b5-579a-4f80-ac10-8380e17e5d98	cdb681a2-98e4-4716-a136-a5e4888e9c32	name_tr	2022-10-21 03:28:30.823715+05	2022-10-21 11:09:53.945098+05	\N
1e16a03a-8ce4-477c-9bfb-209f50358696	b62a1c1c-0a29-4756-8e9d-5c9680758d18	cdb681a2-98e4-4716-a136-a5e4888e9c32	name_pl	2022-10-21 03:28:30.836012+05	2022-10-21 11:09:53.95513+05	\N
d6d74648-e65e-4b18-8bff-7704f48faa27	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	a4277afa-1c92-4f4e-809e-dfbb54ddbc9b	Name_tm	2022-10-21 11:48:26.239328+05	2022-10-21 11:48:26.239328+05	\N
81a3a041-dffe-4431-8cdc-d846d343989c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	a4277afa-1c92-4f4e-809e-dfbb54ddbc9b	Name_ru	2022-10-21 11:48:26.266282+05	2022-10-21 11:48:26.266282+05	\N
21efedd5-2b4b-4700-a8d4-5642c405b042	55a387df-6d38-42ea-bfba-379327b53cbd	a4277afa-1c92-4f4e-809e-dfbb54ddbc9b	NAME_FR	2022-10-21 11:48:26.289638+05	2022-10-21 11:48:26.289638+05	\N
33ad859f-8e12-4761-b9bb-ecd52ad90a77	198695b5-579a-4f80-ac10-8380e17e5d98	a4277afa-1c92-4f4e-809e-dfbb54ddbc9b	name_tr	2022-10-21 11:48:26.298321+05	2022-10-21 11:48:26.298321+05	\N
9f7d33a5-2e70-4c1f-9654-68f7f444074f	b62a1c1c-0a29-4756-8e9d-5c9680758d18	a4277afa-1c92-4f4e-809e-dfbb54ddbc9b	name_pl	2022-10-21 11:48:26.30947+05	2022-10-21 11:48:26.30947+05	\N
ff67a77c-a31e-4761-96e1-804936f88c51	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	849a1c59-45fb-429b-8fe3-a6e34a6dafaa	name_tm	2022-10-21 11:48:48.639641+05	2022-10-21 12:15:12.870392+05	\N
8f16c2d0-53c4-458a-93ad-61c92a11bed1	aea98b93-7bdf-455b-9ad4-a259d69dc76e	849a1c59-45fb-429b-8fe3-a6e34a6dafaa	name_ru	2022-10-21 11:48:48.65354+05	2022-10-21 12:15:12.883594+05	\N
99e80b09-8738-452a-ba4e-d8bf68238bc0	55a387df-6d38-42ea-bfba-379327b53cbd	849a1c59-45fb-429b-8fe3-a6e34a6dafaa	name_fr	2022-10-21 11:48:48.665756+05	2022-10-21 12:15:12.89389+05	\N
35fa0235-38a9-4953-9a13-bc609e3b23d7	198695b5-579a-4f80-ac10-8380e17e5d98	849a1c59-45fb-429b-8fe3-a6e34a6dafaa	name_tr	2022-10-21 11:48:48.67566+05	2022-10-21 12:15:12.905776+05	\N
7214b362-eb97-4e16-b650-af5515e05941	b62a1c1c-0a29-4756-8e9d-5c9680758d18	849a1c59-45fb-429b-8fe3-a6e34a6dafaa	name_pl	2022-10-21 11:48:48.686741+05	2022-10-21 12:15:12.915902+05	\N
032fd254-f202-4fff-8650-e6a5dcb81fd5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	bdabc7aa-a567-48d5-a1d9-b1ff61c6af4b	Name_tm	2022-10-21 22:28:26.651906+05	2022-10-21 22:28:26.651906+05	\N
5f102237-f6f9-4464-872e-61edce3a6a7e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	bdabc7aa-a567-48d5-a1d9-b1ff61c6af4b	Name_ru	2022-10-21 22:28:26.66675+05	2022-10-21 22:28:26.66675+05	\N
aab84f8a-0ba2-4104-bdf6-f142fb5fe3b7	55a387df-6d38-42ea-bfba-379327b53cbd	bdabc7aa-a567-48d5-a1d9-b1ff61c6af4b	NAME_FR	2022-10-21 22:28:26.681302+05	2022-10-21 22:28:26.681302+05	\N
2bccde95-68aa-46f7-a9a6-079319371699	198695b5-579a-4f80-ac10-8380e17e5d98	bdabc7aa-a567-48d5-a1d9-b1ff61c6af4b	name_tr	2022-10-21 22:28:26.702262+05	2022-10-21 22:28:26.702262+05	\N
5e7067dc-1f01-4234-8066-0bb6cb3a433b	b62a1c1c-0a29-4756-8e9d-5c9680758d18	bdabc7aa-a567-48d5-a1d9-b1ff61c6af4b	name_pl	2022-10-21 22:28:26.725956+05	2022-10-21 22:28:26.725956+05	\N
e224ecfc-6daa-4df5-8112-74846fc44867	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	Sowgatlyk toplumlar	2022-06-20 09:46:01.148565+05	2022-10-22 01:26:07.108522+05	\N
53959762-0b63-4100-ae13-4bbf8c015fec	aea98b93-7bdf-455b-9ad4-a259d69dc76e	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	Подарочные наборы	2022-06-20 09:46:01.408239+05	2022-10-22 01:26:07.108522+05	\N
67f168de-ecf8-4885-99d5-df3c0bb9b3d6	55a387df-6d38-42ea-bfba-379327b53cbd	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N
ddc67ec8-ad5e-4d81-a714-734748d23e26	198695b5-579a-4f80-ac10-8380e17e5d98	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N
a121d39a-1a3f-4142-ac4d-9cbd87733027	b62a1c1c-0a29-4756-8e9d-5c9680758d18	338906f1-dbe2-4ba7-84fc-fe7a4d7856ec	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N
85469cf2-f48a-4e73-800d-ebf599aaeaba	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	29ed85bb-11eb-4458-bbf3-5a5644d167d6	Arzanladyş we Aksiýalar	2022-06-20 09:41:17.756928+05	2022-10-22 01:26:07.073419+05	\N
bbdd06a4-2dce-4c99-bf05-cf4e911776c7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	29ed85bb-11eb-4458-bbf3-5a5644d167d6	Распродажи и Акции	2022-06-20 09:41:17.941489+05	2022-10-22 01:26:07.073419+05	\N
72a79790-1880-4338-b929-0edd99c64f93	55a387df-6d38-42ea-bfba-379327b53cbd	29ed85bb-11eb-4458-bbf3-5a5644d167d6	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.073419+05	\N
3e895146-5077-4e2d-9d03-8a59bff095c6	198695b5-579a-4f80-ac10-8380e17e5d98	29ed85bb-11eb-4458-bbf3-5a5644d167d6	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.073419+05	\N
bbf190d2-e0e9-41ec-af82-68715e98b057	b62a1c1c-0a29-4756-8e9d-5c9680758d18	29ed85bb-11eb-4458-bbf3-5a5644d167d6	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.073419+05	\N
34f4cdb5-04b9-48c0-b5b0-0045a02aa094	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	66772380-c161-4c45-9350-a45e765193e2	Aksiýadaky harytlar	2022-06-20 09:45:34.450534+05	2022-10-22 01:26:07.108522+05	\N
713cc05f-6a9d-4dae-88b5-dde2e564480c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	66772380-c161-4c45-9350-a45e765193e2	Продукция в категории Акции	2022-06-20 09:45:34.466904+05	2022-10-22 01:26:07.108522+05	\N
623e54d4-b912-4ad3-8dcc-7800448d2bfb	55a387df-6d38-42ea-bfba-379327b53cbd	66772380-c161-4c45-9350-a45e765193e2	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N
9e55e45f-cb53-49b1-a445-a44ed8e76faa	198695b5-579a-4f80-ac10-8380e17e5d98	66772380-c161-4c45-9350-a45e765193e2	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N
1bfbf71e-e81b-4142-affc-7bcd6c3d31a3	b62a1c1c-0a29-4756-8e9d-5c9680758d18	66772380-c161-4c45-9350-a45e765193e2	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N
3b756a33-bf2c-4d04-af57-962a3226d00b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	45765130-7f97-4f0c-b886-f70b75e02610	Täze harytlar	2022-06-20 10:11:06.719528+05	2022-10-22 01:26:07.108522+05	\N
2d22961c-ef08-4238-ae54-c00593c0073c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	45765130-7f97-4f0c-b886-f70b75e02610	Новые продукты	2022-06-20 10:11:06.735056+05	2022-10-22 01:26:07.108522+05	\N
c5cc497b-2bb8-47b8-b59e-912f18b0fa4a	55a387df-6d38-42ea-bfba-379327b53cbd	45765130-7f97-4f0c-b886-f70b75e02610	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N
de3f2982-1850-4dd5-be06-aa66de1300d1	198695b5-579a-4f80-ac10-8380e17e5d98	45765130-7f97-4f0c-b886-f70b75e02610	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N
3e5fb890-f005-4f07-9793-5bc79166c306	b62a1c1c-0a29-4756-8e9d-5c9680758d18	45765130-7f97-4f0c-b886-f70b75e02610	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N
\.


--
-- Data for Name: translation_contact; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_contact (id, lang_id, full_name, email, phone, letter, company_phone, imo, company_email, instagram, created_at, updated_at, deleted_at, button_text) FROM stdin;
f1693167-0c68-4a54-9831-56f124d629a3	aea98b93-7bdf-455b-9ad4-a259d69dc76e	at_ru	mail_ru	phone_ru	letter ru	cp ru	imo ru	ce ru	instagram ru	2022-06-27 11:29:48.050553+05	2022-06-27 11:29:48.050553+05	\N	Отправить
73253999-7355-42b4-8700-94de76f0058a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	at_tm	mail_tm	phone_tm	letter_tm	cp_tm	imo_tm	ce_tm	ins_tm	2022-06-27 11:29:47.914891+05	2022-06-27 11:29:47.914891+05	\N	ugrat
ea0fe324-a8c6-4426-b132-e36b3b4c08fb	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N	\N
833db713-0ba0-4232-9966-5632573445aa	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N	\N
298d9ec7-ffa1-41ac-88df-378d064a2dc9	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N	\N
\.


--
-- Data for Name: translation_district; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_district (id, lang_id, district_id, name, created_at, updated_at, deleted_at) FROM stdin;
ad9f94d3-05e7-43b3-aa77-7b7f3754d003	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	a58294d3-efe5-4cb7-82d3-8df8c37563c5	Parahat 2	2022-06-25 10:23:25.712337+05	2022-06-25 10:23:25.712337+05	\N
aa1cfa48-3132-4dd4-abfb-070a2986690b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	a58294d3-efe5-4cb7-82d3-8df8c37563c5	Mir 2	2022-06-25 10:23:25.774504+05	2022-06-25 10:23:25.774504+05	\N
987bb3c7-59d3-4f2b-b5ca-6905ec581952	55a387df-6d38-42ea-bfba-379327b53cbd	a58294d3-efe5-4cb7-82d3-8df8c37563c5	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
60a22007-7e91-4cc9-9ec2-bd0dcfd0425a	198695b5-579a-4f80-ac10-8380e17e5d98	a58294d3-efe5-4cb7-82d3-8df8c37563c5	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
a48ce5fd-aeb3-472d-96b3-bc1b56160ff0	b62a1c1c-0a29-4756-8e9d-5c9680758d18	a58294d3-efe5-4cb7-82d3-8df8c37563c5	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
\.


--
-- Data for Name: translation_footer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_footer (id, lang_id, about, payment, contact, secure, word, created_at, updated_at, deleted_at) FROM stdin;
84b5504f-1056-4b44-94dd-a7819148da66	aea98b93-7bdf-455b-9ad4-a259d69dc76e	О нас	Порядок доставки и оплаты	Коммуникация	Обслуживания и Политика Конфиденциальности	Все права защищены	2022-06-22 15:23:32.793161+05	2022-06-22 15:23:32.793161+05	\N
12dc4c16-5712-4bff-a957-8e16d450b4fb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Biz Barada	Eltip bermek we töleg tertibi	Aragatnaşyk	Ulanyş düzgünleri we gizlinlik şertnamasy	Ähli hukuklary goraglydyr	2022-06-22 15:23:32.716064+05	2022-06-22 15:23:32.716064+05	\N
a50a6d02-3604-467b-ae88-4a764483882f	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	uytget	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
fe85940e-0f31-47f9-a9f9-33c0b609d66a	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	uytget	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
b43b6c81-31d6-418b-8d08-bc7c6a03a2b4	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	uytget	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
\.


--
-- Data for Name: translation_header; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_header (id, lang_id, research, phone, password, forgot_password, sign_in, sign_up, name, password_verification, verify_secure, my_information, my_favorites, my_orders, log_out, created_at, updated_at, deleted_at, basket, email, add_to_basket) FROM stdin;
eaf206e6-d515-4bdb-9323-a047cd0edae5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	gözleg	telefon	parol	Acar sozumi unutdym	ulgama girmek	agza bolmak	Ady	Acar sozi tassyklamak	Ulanyş Düzgünlerini we Gizlinlik Şertnamasyny okadym we kabul edýärin	maglumatym	halanlarym	sargytlarym	cykmak	2022-06-16 04:48:26.460534+05	2022-06-16 04:48:26.460534+05	\N	sebet	uytget	uytget
9154e800-2a92-47de-b4ff-1e63b213e5f7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	поиск	tелефон	пароль	забыл пароль	войти	зарегистрироваться	имя	Подтвердить Пароль	Я прочитал и принимаю Условия Обслуживания и Политика Конфиденциальности	моя информация	мои любимые	мои заказы	выйти	2022-06-16 04:48:26.491672+05	2022-06-16 04:48:26.491672+05	\N	корзина	uytget	uytget
cc96bb49-8073-47e0-b733-c8af7cea2df4	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N	uytget	uytget	uytget
3fabc4d9-21cc-41e2-8ae6-7f5c10de0bb6	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N	uytget	uytget	uytget
8105af8e-a161-4620-a000-d6ffa890f092	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N	uytget	uytget	uytget
\.


--
-- Data for Name: translation_my_information_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_my_information_page (id, lang_id, address, created_at, updated_at, deleted_at, birthday, update_password, save) FROM stdin;
d294138e-b808-41ae-9ac5-1826751fda3d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ваш адрес	2022-07-04 19:28:46.603058+05	2022-07-04 19:28:46.603058+05	\N	дата рождения	изменить пароль	запомнить
11074158-69f2-473a-b4fe-94304ff0d8a7	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	salgyňyz	2022-07-04 19:28:46.529935+05	2022-07-04 19:28:46.529935+05	\N	doglan senäň	açar sözi üýtget	ýatda sakla
6f731337-0faf-45f0-8d2d-b378c29907ee	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N	uytget	uytegt	uytegt
4b048c75-163e-4e4f-8af7-336a78234a91	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N	uytget	uytegt	uytegt
a1b328d5-4c55-4a1c-871d-63d65f12f0a3	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N	uytget	uytegt	uytegt
\.


--
-- Data for Name: translation_my_order_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_my_order_page (id, lang_id, orders, date, price, currency, image, name, brend, code, amount, total_price, created_at, updated_at, deleted_at) FROM stdin;
6f30b588-94d8-49f5-a558-a90c2ec9150e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	orders_ru	date_ru	price_ru	currency_ru	image_ru	name_ru	brend_ru	code_ru	amount_ru	total_price_ru	2022-09-02 13:04:39.394714+05	2022-09-02 13:04:39.394714+05	\N
ff43b90d-e22d-4364-b358-6fd56bb3a305	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	orders	date	price	currency	image	name	brend	code	amount	total_price	2022-09-02 13:04:39.36328+05	2022-09-02 13:12:48.119751+05	\N
2f318dd6-890b-46b4-a984-cb3cbbbc5299	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
6e6825cd-e563-40fe-9d6a-ea5b992cd2c0	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
57c2883a-c8a5-46fa-80d6-4030fd798dda	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
\.


--
-- Data for Name: translation_order_dates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_order_dates (id, lang_id, order_date_id, date, created_at, updated_at, deleted_at) FROM stdin;
dcd0c70b-9fa2-4327-8b35-de29bd3febcb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	32646376-c93f-412b-9e75-b3a5fa70df9e	şu gün	2022-09-28 17:35:33.812812+05	2022-09-28 17:35:33.812812+05	\N
3338d831-f091-4574-a0bf-f9cb07dd4893	aea98b93-7bdf-455b-9ad4-a259d69dc76e	32646376-c93f-412b-9e75-b3a5fa70df9e	Cегодня	2022-09-28 17:35:33.82453+05	2022-09-28 17:35:33.82453+05	\N
1aa5185f-9815-4e3f-9c34-718bfb587d91	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c1f2beca-a6b6-4971-a6a7-ed50079c6912	Ertir	2022-09-28 17:36:46.836838+05	2022-09-28 17:36:46.836838+05	\N
9e7a3752-fce2-4b66-bf3e-d915bf463f92	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c1f2beca-a6b6-4971-a6a7-ed50079c6912	Завтра	2022-09-28 17:36:46.847888+05	2022-09-28 17:36:46.847888+05	\N
e7986920-39ff-4d7a-b805-05341516d42d	55a387df-6d38-42ea-bfba-379327b53cbd	32646376-c93f-412b-9e75-b3a5fa70df9e	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
2b504ee9-e7aa-4472-bbee-583eb0abec44	55a387df-6d38-42ea-bfba-379327b53cbd	c1f2beca-a6b6-4971-a6a7-ed50079c6912	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
248034ff-59c0-4957-90b1-f5a11fa152d6	198695b5-579a-4f80-ac10-8380e17e5d98	32646376-c93f-412b-9e75-b3a5fa70df9e	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
e32d37c3-81a9-4420-b3f3-cdae72456285	198695b5-579a-4f80-ac10-8380e17e5d98	c1f2beca-a6b6-4971-a6a7-ed50079c6912	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
8b140e0d-8074-41eb-8b3d-147f8dec413a	b62a1c1c-0a29-4756-8e9d-5c9680758d18	32646376-c93f-412b-9e75-b3a5fa70df9e	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
334dc106-6cbf-4259-a0af-9ffc8d8378cf	b62a1c1c-0a29-4756-8e9d-5c9680758d18	c1f2beca-a6b6-4971-a6a7-ed50079c6912	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
\.


--
-- Data for Name: translation_order_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_order_page (id, lang_id, content, type_of_payment, choose_a_delivery_time, your_address, mark, to_order, tomorrow, cash, payment_terminal, created_at, updated_at, deleted_at) FROM stdin;
474a15e9-1a05-49aa-9a61-c92837d9c9a8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	content_ru	type_of_payment_ru	choose_a_delivery_time_ru	your_address_ru	mark_ru	to_order_ru	tomorrow_ru	cash_ru	payment_terminal_ru	2022-09-01 12:47:16.802639+05	2022-09-01 12:47:16.802639+05	\N
75810722-07fd-400e-94b4-cd230de08cbf	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	content	type_of_payment	choose_a_delivery_time	your_address	mark	to_order	tomorrow	cash	payment_terminal	2022-09-01 12:47:16.720956+05	2022-09-01 12:55:25.638676+05	\N
17338d5e-a818-4465-9697-ad089bc1f11b	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
ce5a982b-d1be-4579-965d-687b2420b573	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
4625c11c-7e60-4782-aca7-e53350fd5478	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
\.


--
-- Data for Name: translation_payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_payment (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
5748ec03-5278-425c-babf-f7f2bf8d2efa	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Eltip bermek we töleg tertibi	Eltip bermek hyzmaty Aşgabat şäheriniň çägi bilen bir hatarda Büzmeýine we Änew şäherine hem elýeterlidir. Hyzmat mugt amala aşyrylýar;\nHer bir sargydyň jemi bahasy azyndan 150 manat bolmalydyr;\nSaýtdan sargyt edeniňizden soňra operator size jaň edip sargydy tassyklar (eger hemişelik müşderi bolsaňyz sargytlaryňyz islegiňize görä awtomatik usulda hem tassyklanýar);\nGirizen salgyňyz we telefon belgiňiz esasynda hyzmat amala aşyrylýar;\nSargyt tassyklanmadyk ýagdaýynda ol hasaba alynmaýar we ýerine ýetirilmeýär. Sargydyň tassyklanmagy üçin girizen telefon belgiňizden jaň kabul edip bilýändigiňize göz ýetiriň. Şeýle hem girizen salgyňyzyň dogrulygyny barlaň;\nSargydy barlap alanyňyzdan soňra töleg amala aşyrylýar. Eltip berijiniň size gowşurýan töleg resminamasynda siziň tölemeli puluňyz bellenendir. Töleg nagt we nagt däl görnüşde milli manatda amala aşyrylýar. Kabul edip tölegini geçiren harydyňyz yzyna alynmaýar;\nSargyt tassyklanandan soňra 24 sagadyň dowamynda eýesi tapylmasa ol güýjüni ýitirýär;	2022-06-25 11:37:47.362666+05	2022-06-25 11:37:47.362666+05	\N
ea7f4c0c-4b1a-41d3-94eb-e058aba9c99f	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Порядок доставки и оплаты	В настоящее время услуга по доставке осуществляется по городу Ашхабад, Бюзмеин и Анау. Услуга предоставляется бесплатно.\nМинимальный заказ должен составлять не менее 150 манат;\nПосле Вашего заказа по сайту, оператор позвонит Вам для подтверждения заказа (постоянным клиентам по их желанию подтверждение осуществляется автоматизированно);\nУслуга доставки выполняется по указанному Вами адресу и номеру телефона;\nЕсли заказ не подтвержден то данный заказ не регистрируется и не выполняется. Для подтверждения заказа, удостоверьтесь, что можете принять звонок по указанному Вами номеру телефона. Также проверьте правильность указанного Вами адреса;\nОплата выполняется после того, как Вы проверите и примите заказ. На платежном документе курьера указана сумма Вашей оплаты. Оплата выполняется наличными и через карту в национальной валюте. Принятый и оплаченный товар возврату не подлежит;\nЕсли не удается найти владельца заказа в течение 24 часов после подтверждения заказа, то данный заказ аннулируется;	2022-06-25 11:37:47.39047+05	2022-06-25 11:37:47.39047+05	\N
a1da8202-2df0-419c-90c0-bb68e4558174	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
aaf32df7-2667-4589-92b5-42493afbf1db	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
07ac02ec-853e-4652-9820-64e0f3076649	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
\.


--
-- Data for Name: translation_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_product (id, lang_id, product_id, name, description, created_at, updated_at, deleted_at, slug) FROM stdin;
2edb91d0-4d17-4128-9bf8-0eb594418ee5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:54:57.051301+05	2022-10-22 01:26:07.108522+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
e12a74a8-f5c4-4dcd-b8c7-038a8d27624d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:54:57.063296+05	2022-10-22 01:26:07.108522+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
65875507-77b2-42a9-9923-45508ae8b156	b62a1c1c-0a29-4756-8e9d-5c9680758d18	81b84c5d-9759-4b86-978a-649c8ef79660	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N	uytget
e156d07d-c7e9-40da-8480-93512f474f80	aea98b93-7bdf-455b-9ad4-a259d69dc76e	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:57:07.264006+05	2022-10-22 01:26:07.108522+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
dc55a19f-d194-4fd9-b3f8-66e8122f0219	55a387df-6d38-42ea-bfba-379327b53cbd	d4156225-082e-4f0f-9b2c-85268114433a	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 02:53:49.519217+05	\N	uytget
1bd8d363-04cd-4ab0-b03e-086ae6179d78	198695b5-579a-4f80-ac10-8380e17e5d98	d4156225-082e-4f0f-9b2c-85268114433a	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 02:53:49.519217+05	\N	uytget
8a7fc25f-4776-498c-818d-95b9fb34fd2d	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:59:14.099998+05	2022-09-17 14:59:14.099998+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
eeccc321-6c74-42c0-8ea7-acda231fc47b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	8df705a5-2351-4aca-b03e-3357a23840b4	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 15:00:15.211143+05	2022-09-17 15:00:15.211143+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
c5e4d0da-dbe9-46c6-87b4-3b70226ca2a9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-10-06 11:07:41.682197+05	2022-10-06 11:07:41.682197+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
925e8b5a-eec6-4f0b-8f46-9718a8f4f653	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:59:14.111821+05	2022-09-17 14:59:14.111821+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
54b061df-a71c-405e-8ddd-0155b034dcd5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	8df705a5-2351-4aca-b03e-3357a23840b4	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 15:00:15.223949+05	2022-09-17 15:00:15.223949+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
535e4de6-ad74-4f19-bda4-4f09baf299f9	b62a1c1c-0a29-4756-8e9d-5c9680758d18	d4156225-082e-4f0f-9b2c-85268114433a	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 02:53:49.519217+05	\N	uytget
ab660f8c-5ba8-45c4-be0b-ad2ef1450d1d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-10-06 11:07:41.692074+05	2022-10-06 11:07:41.692074+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
ff0d6c88-4ec2-49f3-b8b4-a3b1861cccb9	55a387df-6d38-42ea-bfba-379327b53cbd	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N	uytget
0878f53e-cd11-40c8-a16e-7a74e1c5d145	55a387df-6d38-42ea-bfba-379327b53cbd	8df705a5-2351-4aca-b03e-3357a23840b4	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N	uytget
458efe76-79d3-4643-9c35-cffcaa33ff3e	55a387df-6d38-42ea-bfba-379327b53cbd	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N	uytget
f8db1561-1f73-4323-9521-7b9a340b2bd4	198695b5-579a-4f80-ac10-8380e17e5d98	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N	uytget
677b11d4-3fd6-4116-953c-4cbba1d506c7	198695b5-579a-4f80-ac10-8380e17e5d98	8df705a5-2351-4aca-b03e-3357a23840b4	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N	uytget
547e7dbe-1d97-4d70-9250-27836e222977	198695b5-579a-4f80-ac10-8380e17e5d98	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N	uytget
6c9720a9-7ead-4918-a5d8-b5a1e1ad81c5	b62a1c1c-0a29-4756-8e9d-5c9680758d18	c866d5e4-284c-4bea-a94f-cc23f6c7e5d0	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N	uytget
8b6acc76-e3e1-4f43-8b3c-e8f7245fda1c	b62a1c1c-0a29-4756-8e9d-5c9680758d18	8df705a5-2351-4aca-b03e-3357a23840b4	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N	uytget
e43a09fb-e429-42fd-b630-cd5a484ee850	b62a1c1c-0a29-4756-8e9d-5c9680758d18	3e81d4cd-c3c6-4b01-832b-383b8bea5a6a	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N	uytget
b458c6ad-96f9-4a87-8ed4-871cf86611a9	55a387df-6d38-42ea-bfba-379327b53cbd	b2b165a3-2261-4d67-8160-0e239ecd99b5	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N	uytget
c2a16cb7-56ec-45d9-86bc-57834fbae5da	198695b5-579a-4f80-ac10-8380e17e5d98	b2b165a3-2261-4d67-8160-0e239ecd99b5	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N	uytget
5abc86ef-4ec1-45d0-a671-497b029597e3	b62a1c1c-0a29-4756-8e9d-5c9680758d18	b2b165a3-2261-4d67-8160-0e239ecd99b5	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N	uytget
1327fb24-9a60-47bb-aab7-10f1bee360c4	55a387df-6d38-42ea-bfba-379327b53cbd	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N	uytget
4241288e-3c3e-4853-9a43-32eef9867cd8	198695b5-579a-4f80-ac10-8380e17e5d98	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N	uytget
73e9e605-25e6-44fa-b91c-f0745e8285e5	b62a1c1c-0a29-4756-8e9d-5c9680758d18	0d4a6c3c-cc5d-457b-ac9a-ce60eacb94de	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N	uytget
b1de5982-ee89-4dae-afb9-1116fa1259b4	55a387df-6d38-42ea-bfba-379327b53cbd	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N	uytget
2f9185aa-968c-46fa-a942-f20fc28291c9	198695b5-579a-4f80-ac10-8380e17e5d98	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N	uytget
9ff13666-abf2-4179-9f36-de152c95478d	b62a1c1c-0a29-4756-8e9d-5c9680758d18	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N	uytget
12b07b81-9e9c-47fe-a051-5e3f7288ecdd	55a387df-6d38-42ea-bfba-379327b53cbd	d731b17a-ae8d-4561-ad67-0f431d5c529b	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N	uytget
53df335e-5e07-4bc3-84cc-fb303daf047d	55a387df-6d38-42ea-bfba-379327b53cbd	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	name_fr	description_fr	2022-10-17 02:31:43.703806+05	2022-10-24 00:14:29.337348+05	\N	name_fr
3e58887c-4a35-439b-b4bc-9a7c90aa8bb7	198695b5-579a-4f80-ac10-8380e17e5d98	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	name_tr	description_tr	2022-10-19 11:00:40.050132+05	2022-10-24 00:14:29.359168+05	\N	name_tr
f162152a-0262-4e87-813c-0c932399dc47	b62a1c1c-0a29-4756-8e9d-5c9680758d18	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	name_pl	description_pl	2022-10-20 01:44:26.912355+05	2022-10-24 00:14:29.368743+05	\N	name_pl
9bc5a5be-72d2-4494-ac02-dd20954a83ab	aea98b93-7bdf-455b-9ad4-a259d69dc76e	660071e0-8f17-4c48-9d80-d4cac306de3a	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:58:40.15465+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
dcedd91a-b7e9-4e9f-ae0a-f58308e6d751	aea98b93-7bdf-455b-9ad4-a259d69dc76e	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	name_ru	description_ru	2022-09-17 14:59:44.910924+05	2022-10-24 00:14:29.323205+05	\N	name_ru
a4e1fdc0-0330-4238-bebf-566018205273	55a387df-6d38-42ea-bfba-379327b53cbd	660071e0-8f17-4c48-9d80-d4cac306de3a	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05	uytget
fdba8cdd-2eeb-415a-a744-f4219966db66	198695b5-579a-4f80-ac10-8380e17e5d98	d731b17a-ae8d-4561-ad67-0f431d5c529b	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N	uytget
6fff60d6-4ec9-46d9-80a9-df657ec267e7	55a387df-6d38-42ea-bfba-379327b53cbd	81b84c5d-9759-4b86-978a-649c8ef79660	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N	uytget
a56de9a0-8e03-4bd7-b09a-fa3be1d9b001	198695b5-579a-4f80-ac10-8380e17e5d98	81b84c5d-9759-4b86-978a-649c8ef79660	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N	uytget
ce61fdba-2628-4f09-aff2-27ce8ac6b37c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d4156225-082e-4f0f-9b2c-85268114433a	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:57:34.654635+05	2022-10-22 02:53:49.519217+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
e008099d-ff03-4182-86ce-d91ca984ca76	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d4156225-082e-4f0f-9b2c-85268114433a	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:57:34.642037+05	2022-10-22 02:53:49.519217+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
2756e684-ad6a-4e95-89f8-75b509f63290	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b2b165a3-2261-4d67-8160-0e239ecd99b5	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:55:35.51906+05	2022-10-22 01:26:07.108522+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
33011809-8da8-4563-8dd0-22ca01e5caee	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b2b165a3-2261-4d67-8160-0e239ecd99b5	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:55:35.530179+05	2022-10-22 01:26:07.108522+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
0577ae3a-d145-4a22-9d4e-a13fc9136c2c	b62a1c1c-0a29-4756-8e9d-5c9680758d18	d731b17a-ae8d-4561-ad67-0f431d5c529b	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N	uytget
ad74bc57-3cd1-4c50-9287-a6a21b4beca4	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d731b17a-ae8d-4561-ad67-0f431d5c529b	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:56:36.207944+05	2022-10-22 01:26:07.108522+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
ceeb7241-e4b8-4fa0-b99f-80ba9c141589	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d731b17a-ae8d-4561-ad67-0f431d5c529b	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:56:36.219444+05	2022-10-22 01:26:07.108522+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
1f038393-3955-4cef-a6a1-a1cf087173c5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:56:05.464418+05	2022-10-22 01:26:07.108522+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
c499a61c-948d-41d4-9dad-9d12ea7324d4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	a2bb8745-1f3a-4de9-ad66-11b0bb3bb754	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:56:05.475198+05	2022-10-22 01:26:07.108522+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
f6791ed2-e2a4-4d3a-a32f-59ae4b34d832	55a387df-6d38-42ea-bfba-379327b53cbd	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-22 01:26:07.108522+05	\N	uytget
70894bc5-12cf-4910-9631-ece1658d1449	198695b5-579a-4f80-ac10-8380e17e5d98	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-22 01:26:07.108522+05	\N	uytget
66d53b03-1f2b-4a7a-bc86-dd90baaec6ef	aea98b93-7bdf-455b-9ad4-a259d69dc76e	81b84c5d-9759-4b86-978a-649c8ef79660	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-09-17 14:58:10.064659+05	2022-10-22 01:26:07.108522+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
8965478c-0afe-4b65-af05-0d151c8dd462	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	81b84c5d-9759-4b86-978a-649c8ef79660	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:58:10.054293+05	2022-10-22 01:26:07.108522+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
4ae2447a-d920-47b0-a248-e8db464c1796	b62a1c1c-0a29-4756-8e9d-5c9680758d18	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-22 01:26:07.108522+05	\N	uytget
5aa614ae-8c7d-47d7-867e-44c3d4d2015c	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	bb6c3bdb-79e2-44b3-98b1-c1cee0976777	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:57:07.254121+05	2022-10-22 01:26:07.108522+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
69ca7230-f4e4-4a20-a502-e88df834ca44	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	939da5d3-2f7a-40e2-b133-0b4113280647	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-10-22 12:46:26.52067+05	2022-10-22 12:46:26.52067+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
1a58c5bb-a6d8-4ebe-a450-ded3a0fdcb3c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	939da5d3-2f7a-40e2-b133-0b4113280647	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-10-22 12:46:26.532275+05	2022-10-22 12:46:26.532275+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
9409f397-3c64-49c6-ad01-2130be26cbba	55a387df-6d38-42ea-bfba-379327b53cbd	939da5d3-2f7a-40e2-b133-0b4113280647	name_fr	description_fr	2022-10-22 12:46:26.544378+05	2022-10-22 12:46:26.544378+05	\N	name_fr
2b454e26-3631-4694-92a7-e04d6cb6dd10	198695b5-579a-4f80-ac10-8380e17e5d98	939da5d3-2f7a-40e2-b133-0b4113280647	name_tr	description_tr	2022-10-22 12:46:26.554082+05	2022-10-22 12:46:26.554082+05	\N	name_tr
cc92195a-bcce-4970-a7a6-a5ee104ef5c6	b62a1c1c-0a29-4756-8e9d-5c9680758d18	939da5d3-2f7a-40e2-b133-0b4113280647	name_pl	description_pl	2022-10-22 12:46:26.565398+05	2022-10-22 12:46:26.565398+05	\N	name_pl
663c123b-4590-442d-9026-a562e085c681	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	442ffe07-6c0b-459d-80cd-8e12e2147568	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-10-22 12:47:02.011325+05	2022-10-22 12:47:02.011325+05	\N	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
a31747b4-4faf-4d60-a196-10025f56b0b9	aea98b93-7bdf-455b-9ad4-a259d69dc76e	442ffe07-6c0b-459d-80cd-8e12e2147568	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	Жидкое крем-мыло увлажняющее Aura Clean "Черничный йогурт" 1 л	2022-10-22 12:47:02.022757+05	2022-10-22 12:47:02.022757+05	\N	zhidkoe-krem-mylo-uvlazhniaiushchee-aura-clean-chernichnyi-iogurt-1-l
6d43ddba-2959-4292-a401-b314f3187ade	55a387df-6d38-42ea-bfba-379327b53cbd	442ffe07-6c0b-459d-80cd-8e12e2147568	name_fr	description_fr	2022-10-22 12:47:02.033862+05	2022-10-22 12:47:02.033862+05	\N	name_fr
dccec077-5d56-40c9-b8d1-8d416992bfa1	198695b5-579a-4f80-ac10-8380e17e5d98	442ffe07-6c0b-459d-80cd-8e12e2147568	name_tr	description_tr	2022-10-22 12:47:02.043856+05	2022-10-22 12:47:02.043856+05	\N	name_tr
daf541ff-0c4b-4c7b-81bd-6f42c08d17de	b62a1c1c-0a29-4756-8e9d-5c9680758d18	442ffe07-6c0b-459d-80cd-8e12e2147568	name_pl	description_pl	2022-10-22 12:47:02.054955+05	2022-10-22 12:47:02.054955+05	\N	name_pl
8f666b20-35be-41df-b276-472ae2d5dd3d	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	e3c33ead-3c30-40f1-9d28-7bb8b71b767f	name_tm	description_tm	2022-09-17 14:59:44.899731+05	2022-10-24 00:14:29.314539+05	\N	name_tm
3ec435d8-394a-4002-959e-c2d61d242307	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	660071e0-8f17-4c48-9d80-d4cac306de3a	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	Nemlendiriji suwuk sabyn Aura Clean "Черничный йогурт" 1 ltr	2022-09-17 14:58:40.143451+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05	nemlendiriji-suwuk-sabyn-aura-clean-chernichnyi-iogurt-1-ltr
cd06d549-1037-4cec-8d83-5e2e5a8ffeae	198695b5-579a-4f80-ac10-8380e17e5d98	660071e0-8f17-4c48-9d80-d4cac306de3a	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05	uytget
198df259-91ee-40ea-8ace-19519e6e53e7	b62a1c1c-0a29-4756-8e9d-5c9680758d18	660071e0-8f17-4c48-9d80-d4cac306de3a	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-24 00:48:25.038363+05	2022-10-24 00:48:25.038363+05	uytget
\.


--
-- Data for Name: translation_secure; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_secure (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
3579a847-ce74-4fbe-b10d-8aba83867857	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Пользовательское соглашение	Между Ынамдар – Интернет Маркетом (далее – “Ынамдар”) и интернет сайтом www.ynamdar.com (далее – “Сайт”), а также его клиентом (далее - “Клиент”) достигнуто соглашение по нижеследующим условиям.\n	2022-06-25 10:46:54.221498+05	2022-06-25 10:46:54.221498+05	\N
5988b64a-82ad-4ed0-bd1b-bdd0b3b05912	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ÖZARA YLALAŞYGY	Ynamdar - Internet Marketi (Mundan beýläk – “Ynamdar”) we www.ynamdar.com internet saýty (Mundan beýläk – “Saýt”) bilen, onuň agzasynyň (“Agza”) arasynda aşakdaky şertleri ýerine ýetirmek barada ylalaşyga gelindi.	2022-06-25 10:46:54.190131+05	2022-06-25 10:46:54.190131+05	\N
869da7b2-efb5-40d6-ba4c-cb8bb5c12fe1	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N
1fa9f7fa-9430-43cc-8aac-f0afb1aef4b1	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N
ea183b95-1613-48eb-b425-44be17f427e9	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N
\.


--
-- Data for Name: translation_update_password_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_update_password_page (id, lang_id, title, verify_password, explanation, save, created_at, updated_at, deleted_at, password) FROM stdin;
5190ca93-7007-4db4-8105-65cc3b1af868	aea98b93-7bdf-455b-9ad4-a259d69dc76e	изменить пароль	Подтвердить Пароль	ключевое слово должно быть буквой или цифрой длиной от 5 до 20	запомнить	2022-07-05 10:35:08.984141+05	2022-07-05 10:35:08.984141+05	\N	ключевое слово
de12082b-baab-4b83-ac07-119df09d1230	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	açar sözi üýtgetmek	açar sözi tassykla	siziň açar sözüňiz 5-20 uzynlygynda harp ýa-da sandan ybarat bolmalydyr	ýatda sakla	2022-07-05 10:35:08.867617+05	2022-07-05 10:35:08.867617+05	\N	açar sözi
06503847-0b5a-4b39-8124-6f89c7d9ece7	55a387df-6d38-42ea-bfba-379327b53cbd	uytget	uytget	uytget	uytget	2022-10-17 02:31:43.703806+05	2022-10-17 11:32:22.801107+05	\N	uytget
1652ca85-8e9c-4661-a121-3281ccddd010	198695b5-579a-4f80-ac10-8380e17e5d98	uytget	uytget	uytget	uytget	2022-10-19 11:00:40.050132+05	2022-10-19 12:55:44.565405+05	\N	uytget
d9747bdd-a010-470a-8e2a-2d1dc54faf21	b62a1c1c-0a29-4756-8e9d-5c9680758d18	uytget	uytget	uytget	uytget	2022-10-20 01:44:26.912355+05	2022-10-20 01:44:26.912355+05	\N	uytget
\.


--
-- Name: orders_order_number_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_order_number_seq', 24, true);


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

