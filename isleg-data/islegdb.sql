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
declare category_uuid uuid;
begin
FOR category_uuid IN SELECT id FROM categories WHERE parent_category_id = cat_id
LOOP UPDATE translation_category SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE category_product SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE products SET deleted_at = now() FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = category_uuid;
UPDATE translation_product SET deleted_at = now() FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE images SET deleted_at = now() FROM products,category_product WHERE images.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
end loop; end; $$;


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
declare category_uuid uuid;
begin
FOR category_uuid IN SELECT id FROM categories WHERE parent_category_id = cat_id
LOOP UPDATE translation_category SET deleted_at = NULL WHERE category_id = category_uuid;
UPDATE category_product SET deleted_at = NULL WHERE category_id = category_uuid;
UPDATE products SET deleted_at = NULL FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = category_uuid;
UPDATE translation_product SET deleted_at = NULL FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE images SET deleted_at = NULL FROM products,category_product WHERE images.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
end loop; end; $$;


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
begin
UPDATE brends SET deleted_at = now() WHERE id = b_id;
UPDATE products SET deleted_at = now() WHERE brend_id = b_id;
UPDATE translation_product SET deleted_at = now() FROM products WHERE translation_product.product_id=products.id AND products.brend_id = b_id;
UPDATE images SET deleted_at = now() FROM products WHERE images.product_id=products.id AND products.brend_id = b_id;
end; $$;


ALTER PROCEDURE public.delete_brend(b_id uuid) OWNER TO postgres;

--
-- Name: delete_category(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_category(category_uuid uuid)
    LANGUAGE plpgsql
    AS $$
begin
UPDATE categories SET deleted_at = now() WHERE id = category_uuid;
UPDATE translation_category SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE categories SET deleted_at = now() WHERE parent_category_id = category_uuid;
UPDATE category_product SET deleted_at = now() WHERE category_id = category_uuid;
UPDATE products SET deleted_at = now() FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = category_uuid;
UPDATE translation_product SET deleted_at = now() FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
UPDATE images SET deleted_at = now() FROM products,category_product WHERE images.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = category_uuid;
end; $$;


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
UPDATE translation_notification SET deleted_at = now() WHERE lang_id = language_id;
END; $$;


ALTER PROCEDURE public.delete_language(language_id uuid) OWNER TO postgres;

--
-- Name: delete_notification(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_notification(n_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE notifications SET deleted_at = now() WHERE id = n_id;
UPDATE translation_notification SET deleted_at = now() WHERE notification_id = n_id;
END; $$;


ALTER PROCEDURE public.delete_notification(n_id uuid) OWNER TO postgres;

--
-- Name: delete_order_date(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_order_date(od_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE order_dates SET deleted_at = now() WHERE id = od_id;
UPDATE order_times SET deleted_at = now() WHERE order_date_id = od_id;
UPDATE translation_order_dates SET deleted_at = now() WHERE order_date_id = od_id;
END; $$;


ALTER PROCEDURE public.delete_order_date(od_id uuid) OWNER TO postgres;

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
UPDATE images SET deleted_at = now() WHERE product_id = p_id;
END; $$;


ALTER PROCEDURE public.delete_product(p_id uuid) OWNER TO postgres;

--
-- Name: delete_shop(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.delete_shop(s_id uuid)
    LANGUAGE plpgsql
    AS $$
begin
UPDATE shops SET deleted_at = now() WHERE id = s_id;
UPDATE products SET deleted_at = now() WHERE shop_id = s_id;
UPDATE translation_product SET deleted_at = now() FROM products WHERE translation_product.product_id = products.id AND products.shop_id = s_id;
UPDATE images SET deleted_at = now() FROM products WHERE images.product_id = products.id AND products.shop_id = s_id;
end; $$;


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
begin
UPDATE brends SET deleted_at = NULL WHERE id = b_id;
UPDATE products SET deleted_at = NULL WHERE brend_id = b_id;
UPDATE translation_product SET deleted_at = NULL FROM products WHERE translation_product.product_id=products.id AND products.brend_id = b_id;
UPDATE images SET deleted_at = NULL FROM products WHERE images.product_id=products.id AND products.brend_id = b_id;
end; $$;


ALTER PROCEDURE public.restore_brend(b_id uuid) OWNER TO postgres;

--
-- Name: restore_category(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_category(cat_id uuid)
    LANGUAGE plpgsql
    AS $$
begin
UPDATE categories SET deleted_at = NULL WHERE id = cat_id;
UPDATE translation_category SET deleted_at = NULL WHERE category_id = cat_id;
UPDATE categories SET deleted_at = NULL WHERE parent_category_id = cat_id;
UPDATE category_product SET deleted_at = NULL WHERE category_id = cat_id;
UPDATE products SET deleted_at = NULL FROM category_product WHERE category_product.product_id = products.id AND category_product.category_id = cat_id;
UPDATE translation_product SET deleted_at = NULL FROM products,category_product WHERE translation_product.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = cat_id;
UPDATE images SET deleted_at = NULL FROM products,category_product WHERE images.product_id = products.id AND category_product.product_id = products.id  AND category_product.category_id = cat_id;
end; $$;


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
UPDATE translation_notification SET deleted_at = NULL WHERE lang_id = language_id;
END; $$;


ALTER PROCEDURE public.restore_language(language_id uuid) OWNER TO postgres;

--
-- Name: restore_notification(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_notification(n_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE notifications SET deleted_at = NULL WHERE id = n_id;
UPDATE translation_notification SET deleted_at = NULL WHERE notification_id = n_id;
END; $$;


ALTER PROCEDURE public.restore_notification(n_id uuid) OWNER TO postgres;

--
-- Name: restore_order_date(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_order_date(od_id uuid)
    LANGUAGE plpgsql
    AS $$
BEGIN
UPDATE order_dates SET deleted_at = NULL WHERE id = od_id;
UPDATE order_times SET deleted_at = NULL WHERE order_date_id = od_id;
UPDATE translation_order_dates SET deleted_at = NULL WHERE order_date_id = od_id; 
END; $$;


ALTER PROCEDURE public.restore_order_date(od_id uuid) OWNER TO postgres;

--
-- Name: restore_product(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_product(p_id uuid)
    LANGUAGE plpgsql
    AS $$
begin
UPDATE products SET deleted_at = NULL WHERE id = p_id;
UPDATE category_product SET deleted_at = NULL WHERE product_id = p_id;
UPDATE translation_product SET deleted_at = NULL WHERE product_id = p_id;
UPDATE images SET deleted_at = NULL WHERE product_id = p_id;
end; $$;


ALTER PROCEDURE public.restore_product(p_id uuid) OWNER TO postgres;

--
-- Name: restore_shop(uuid); Type: PROCEDURE; Schema: public; Owner: postgres
--

CREATE PROCEDURE public.restore_shop(s_id uuid)
    LANGUAGE plpgsql
    AS $$
begin
UPDATE shops SET deleted_at = NULL WHERE id = s_id;
UPDATE products SET deleted_at = NULL WHERE shop_id = s_id;
UPDATE translation_product SET deleted_at = NULL FROM products WHERE translation_product.product_id = products.id AND products.shop_id = s_id;
UPDATE images SET deleted_at = NULL FROM products WHERE images.product_id = products.id AND products.shop_id = s_id;
end; $$;


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
-- Name: admins; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.admins (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    full_name character varying(50) NOT NULL,
    phone_number character varying(20) NOT NULL,
    password character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    type character varying(15) NOT NULL
);


ALTER TABLE public.admins OWNER TO postgres;

--
-- Name: afisa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.afisa (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    image character varying(100),
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
    image character varying(100) NOT NULL,
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
    name character varying(1000) NOT NULL,
    image character varying(100) NOT NULL,
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
    product_id uuid NOT NULL,
    customer_id uuid NOT NULL,
    quantity_of_product bigint NOT NULL,
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
    image character varying(100),
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
    category_id uuid NOT NULL,
    product_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.category_product OWNER TO postgres;

--
-- Name: company_address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.company_address (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid NOT NULL,
    address character varying DEFAULT 'uytget'::character varying NOT NULL,
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
    phone character varying(20) NOT NULL,
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
    logo character varying(100) NOT NULL,
    favicon character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    instagram character varying(100) NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp without time zone,
    imo character varying(20) NOT NULL
);


ALTER TABLE public.company_setting OWNER TO postgres;

--
-- Name: customer_address; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer_address (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    customer_id uuid NOT NULL,
    address character varying NOT NULL,
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
    full_name character varying(50) NOT NULL,
    phone_number character varying(20) NOT NULL,
    password character varying,
    birthday date,
    gender character varying(10),
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    email character varying(100),
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
    product_id uuid NOT NULL,
    image character varying(100) NOT NULL,
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
    name_short character varying(10) NOT NULL,
    flag character varying(100) NOT NULL,
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
    product_id uuid NOT NULL,
    customer_id uuid NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.likes OWNER TO postgres;

--
-- Name: notifications; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.notifications (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.notifications OWNER TO postgres;

--
-- Name: order_dates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.order_dates (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    date character varying(50),
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
    order_date_id uuid NOT NULL,
    "time" character varying(50) NOT NULL,
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
    product_id uuid NOT NULL,
    quantity_of_product integer NOT NULL,
    order_id uuid NOT NULL,
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
    customer_id uuid NOT NULL,
    customer_mark character varying,
    order_time character varying(50) NOT NULL,
    payment_type character varying(50) NOT NULL,
    total_price numeric NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    order_number integer NOT NULL,
    shipping_price numeric NOT NULL,
    excel character varying(100),
    address character varying DEFAULT 'uytget'::character varying NOT NULL
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
    lang_id uuid NOT NULL,
    type character varying(100) DEFAULT 'uytget'::character varying NOT NULL,
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
    price numeric NOT NULL,
    old_price numeric,
    amount bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    limit_amount bigint NOT NULL,
    is_new boolean DEFAULT false,
    shop_id uuid,
    main_image character varying(100) NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: shops; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shops (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    owner_name character varying(50) NOT NULL,
    address character varying NOT NULL,
    phone_number character varying(20) NOT NULL,
    running_time character varying(20) NOT NULL,
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
    lang_id uuid NOT NULL,
    title character varying DEFAULT 'uytget'::character varying NOT NULL,
    content text DEFAULT 'uytget'::text NOT NULL,
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
    afisa_id uuid NOT NULL,
    lang_id uuid NOT NULL,
    title character varying DEFAULT 'uytget'::character varying NOT NULL,
    description text DEFAULT 'uytget'::text NOT NULL,
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
    lang_id uuid NOT NULL,
    quantity_of_goods character varying DEFAULT 'uytget'::character varying NOT NULL,
    total_price character varying DEFAULT 'uytget'::character varying NOT NULL,
    discount character varying DEFAULT 'uytget'::character varying NOT NULL,
    delivery character varying DEFAULT 'uytget'::character varying NOT NULL,
    total character varying DEFAULT 'uytget'::character varying NOT NULL,
    to_order character varying DEFAULT 'uytget'::character varying NOT NULL,
    your_basket character varying DEFAULT 'uytget'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    empty_the_basket character varying DEFAULT 'uytget'::character varying NOT NULL,
    empty_the_like_page character varying DEFAULT 'uytget'::character varying NOT NULL
);


ALTER TABLE public.translation_basket_page OWNER TO postgres;

--
-- Name: translation_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_category (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid NOT NULL,
    category_id uuid NOT NULL,
    name character varying DEFAULT 'uytget'::character varying NOT NULL,
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
    lang_id uuid NOT NULL,
    full_name character varying DEFAULT 'uytget'::character varying NOT NULL,
    email character varying DEFAULT 'uytget'::character varying NOT NULL,
    phone character varying DEFAULT 'uytget'::character varying NOT NULL,
    letter character varying DEFAULT 'uytget'::character varying NOT NULL,
    company_phone character varying DEFAULT 'uytget'::character varying NOT NULL,
    imo character varying DEFAULT 'uytget'::character varying NOT NULL,
    company_email character varying DEFAULT 'uytget'::character varying NOT NULL,
    instagram character varying DEFAULT 'uytget'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    button_text character varying NOT NULL
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
    lang_id uuid NOT NULL,
    about character varying DEFAULT 'uytget'::character varying NOT NULL,
    payment character varying DEFAULT 'uytget'::character varying NOT NULL,
    contact character varying DEFAULT 'uytget'::character varying NOT NULL,
    secure character varying DEFAULT 'uytget'::character varying NOT NULL,
    word character varying DEFAULT 'uytget'::character varying NOT NULL,
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
    lang_id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    research character varying DEFAULT 'uytget'::character varying NOT NULL,
    phone character varying DEFAULT 'uytget'::character varying NOT NULL,
    password character varying DEFAULT 'uytget'::character varying NOT NULL,
    forgot_password character varying DEFAULT 'uytget'::character varying NOT NULL,
    sign_in character varying DEFAULT 'uytget'::character varying NOT NULL,
    sign_up character varying DEFAULT 'uytget'::character varying NOT NULL,
    name character varying DEFAULT 'uytget'::character varying NOT NULL,
    password_verification character varying DEFAULT 'uytget'::character varying NOT NULL,
    verify_secure character varying DEFAULT 'uytget'::character varying NOT NULL,
    my_information character varying DEFAULT 'uytget'::character varying NOT NULL,
    my_favorites character varying DEFAULT 'uytget'::character varying NOT NULL,
    my_orders character varying DEFAULT 'uytget'::character varying NOT NULL,
    log_out character varying DEFAULT 'uytget'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    basket character varying DEFAULT 'uytget'::character varying NOT NULL,
    email character varying DEFAULT 'uytget'::character varying NOT NULL,
    add_to_basket character varying DEFAULT 'uytget'::character varying NOT NULL,
    add_button character varying DEFAULT 'uytget'::character varying NOT NULL
);


ALTER TABLE public.translation_header OWNER TO postgres;

--
-- Name: translation_my_information_page; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_my_information_page (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid NOT NULL,
    address character varying DEFAULT 'uytget'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    birthday character varying DEFAULT 'uytget'::character varying NOT NULL,
    update_password character varying DEFAULT 'uytegt'::character varying NOT NULL,
    save character varying DEFAULT 'uytegt'::character varying NOT NULL,
    gender character varying DEFAULT 'uytget'::character varying NOT NULL,
    male character varying DEFAULT 'uytget'::character varying NOT NULL,
    female character varying DEFAULT 'uytget'::character varying NOT NULL
);


ALTER TABLE public.translation_my_information_page OWNER TO postgres;

--
-- Name: translation_my_order_page; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_my_order_page (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid NOT NULL,
    orders character varying DEFAULT 'uytget'::character varying NOT NULL,
    date character varying DEFAULT 'uytget'::character varying NOT NULL,
    price character varying DEFAULT 'uytget'::character varying NOT NULL,
    image character varying DEFAULT 'uytget'::character varying NOT NULL,
    name character varying DEFAULT 'uytget'::character varying NOT NULL,
    brend character varying DEFAULT 'uytget'::character varying NOT NULL,
    product_price character varying DEFAULT 'uytget'::character varying NOT NULL,
    amount character varying DEFAULT 'uytget'::character varying NOT NULL,
    total_price character varying DEFAULT 'uytget'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_my_order_page OWNER TO postgres;

--
-- Name: translation_notification; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_notification (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    notification_id uuid NOT NULL,
    lang_id uuid NOT NULL,
    translation character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.translation_notification OWNER TO postgres;

--
-- Name: translation_order_dates; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_order_dates (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid NOT NULL,
    order_date_id uuid NOT NULL,
    date character varying(50) DEFAULT 'uytget'::character varying NOT NULL,
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
    lang_id uuid NOT NULL,
    content character varying DEFAULT 'uytget'::character varying NOT NULL,
    type_of_payment character varying DEFAULT 'uytget'::character varying NOT NULL,
    choose_a_delivery_time character varying DEFAULT 'uytget'::character varying NOT NULL,
    your_address character varying DEFAULT 'uytget'::character varying NOT NULL,
    mark character varying DEFAULT 'uytget'::character varying NOT NULL,
    to_order character varying DEFAULT 'uytget'::character varying NOT NULL,
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
    lang_id uuid NOT NULL,
    title character varying DEFAULT 'uytget'::character varying NOT NULL,
    content text DEFAULT 'uytget'::text NOT NULL,
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
    lang_id uuid NOT NULL,
    product_id uuid NOT NULL,
    name character varying DEFAULT 'uytget'::character varying NOT NULL,
    description text DEFAULT 'uytget'::text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    slug character varying DEFAULT 'uytget'::character varying NOT NULL
);


ALTER TABLE public.translation_product OWNER TO postgres;

--
-- Name: translation_secure; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_secure (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    lang_id uuid NOT NULL,
    title character varying DEFAULT 'uytget'::character varying NOT NULL,
    content text DEFAULT 'uytget'::text NOT NULL,
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
    lang_id uuid NOT NULL,
    title character varying DEFAULT 'uytget'::character varying NOT NULL,
    verify_password character varying DEFAULT 'uytget'::character varying NOT NULL,
    explanation character varying DEFAULT 'uytget'::character varying NOT NULL,
    save character varying DEFAULT 'uytget'::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    password character varying DEFAULT 'uytget'::character varying NOT NULL
);


ALTER TABLE public.translation_update_password_page OWNER TO postgres;

--
-- Name: orders order_number; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders ALTER COLUMN order_number SET DEFAULT nextval('public.orders_order_number_seq'::regclass);


--
-- Data for Name: admins; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.admins (id, full_name, phone_number, password, created_at, updated_at, deleted_at, type) FROM stdin;
5e0e9a0a-e07a-4911-bbc0-8dad29a6abbd	Allanur Bayramgeldiyev	+99362420377	$2a$14$w5DuoaqTJEdPpT3fzrLdh.Q3BGm/wX76NFGPGhy9G4nU/fDUXvtPy	2022-11-02 16:46:19.827563+00	2022-11-02 16:46:19.827563+00	\N	super_admin
c97bfc6a-fd85-4aa8-82db-e788f6b0d70a	Muhammet Bayramov	+99363747155	$2a$14$IXSdYxI0f.qQ8kDuLq.DU.F4ZRnMuq58VErTjaFNdquqzcZaenImu	2022-11-02 16:48:30.593337+00	2022-11-02 16:48:30.593337+00	\N	super_admin
6989254b-79c7-412c-acb2-19f67a3277d5	Seyit Batyrov	+99361111111	$2a$14$qD9HYvqPHUVITgfAJLfj3uODG.hcGiI7.ayv3jc1NlSD34QA5drv2	2022-11-02 17:30:44.622854+00	2022-11-02 17:30:44.622854+00	\N	admin
42ba1c9b-f56d-44e4-a72b-5031d0f3ce64	Kakajan Batyrov	+99362222222	$2a$14$QxhaMFkztA7x1jlR9PT3zOwCystNe/7QDEp7C04xLNFp2MDMCgpHK	2022-11-02 17:31:53.959553+00	2022-11-02 17:31:53.959553+00	\N	admin
54f172bc-e13b-4dd8-95af-c3e364f09e3b	Maya Kerimova	+99363333333	$2a$14$tR7soepOOiWQ0jtHpmTLy.kYOzMfbu2RNk2iICMeYHFqwiI1AvKBm	2022-11-02 17:43:54.030609+00	2022-11-02 17:43:54.030609+00	\N	super_admin
\.


--
-- Data for Name: afisa; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.afisa (id, image, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: banner; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.banner (id, image, url, created_at, updated_at, deleted_at) FROM stdin;
785d48b7-8600-4ff4-8608-23e5114ac3f5	uploads/banner/f07abf59-09ab-4900-92df-2f6d54f9864b.jpeg	/category/44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	2022-12-16 04:07:23.802027+00	2022-12-16 04:07:23.802027+00	\N
\.


--
-- Data for Name: brends; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.brends (id, name, image, created_at, updated_at, deleted_at) FROM stdin;
9b838628-fd75-4232-862d-998635f24f52	Yalong	uploads/brend/4fe27aa5-d228-4e75-97fd-02291bea4db6.png	2023-01-17 04:54:38.997513+00	2023-01-17 04:54:38.997513+00	\N
30723733-79c0-4f1b-9365-a0943a11557c	КанцПарк	uploads/brend/45aae190-a1d3-43a2-b25c-3ddd12cf9508.png	2023-01-17 04:57:31.559121+00	2023-01-17 04:57:31.559121+00	\N
1b5981d8-5136-4c3d-9a0c-35370bb586f5	ПЕРО	uploads/brend/89ce9bd6-f7a4-4f84-a7a7-2cd82da54f50.png	2023-01-17 05:00:49.7397+00	2023-01-17 05:00:49.7397+00	\N
e9b7cc3b-bc80-4682-95e1-7db3bfd8f9b7	MAPED	uploads/brend/9834457a-af44-4764-b601-dd795a42bfaa.png	2023-01-17 05:01:38.076937+00	2023-01-17 05:01:38.076937+00	\N
639ae6a7-a3d0-4c68-b99c-1620086761b4	Комус	uploads/brend/c75f6fac-1485-4c41-8038-da02b35279ba.png	2023-01-17 05:02:57.039469+00	2023-01-17 05:02:57.039469+00	\N
014a8138-7caa-40a6-aed2-d08a8e7e72c6	Old Spice	uploads/brend/8039b305-9eb4-4ec9-be8f-d184b96557c1.png	2023-01-17 05:03:20.250687+00	2023-01-17 05:03:20.250687+00	\N
\.


--
-- Data for Name: cart; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cart (id, product_id, customer_id, quantity_of_product, created_at, updated_at, deleted_at) FROM stdin;
96197f01-8f8f-4b33-bf26-a3df612718bd	14d95413-2c8a-472f-8f89-9458dc1bde33	12c4d76a-e3a6-4f35-97ba-efed264f849a	9	2023-01-25 18:27:15.654486+00	2023-01-25 18:33:34.603621+00	\N
900a226f-9565-4b88-ad69-f4b55de939b2	14d95413-2c8a-472f-8f89-9458dc1bde33	19cdcf1a-f110-4510-a52b-063329d98607	1	2023-01-26 18:30:34.127449+00	2023-01-26 18:30:34.127449+00	\N
\.


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, parent_category_id, image, is_home_category, created_at, updated_at, deleted_at) FROM stdin;
d7862d17-0742-4bd5-8fc8-478fd7e868c4	d154a3f1-7086-439f-b343-3998d6521efa		t	2022-10-27 07:38:27.306873+00	2022-10-27 07:38:27.306873+00	\N
71994790-1b7b-41ab-90a8-b3df0d68e3e6	d154a3f1-7086-439f-b343-3998d6521efa		t	2022-10-27 07:38:49.13083+00	2022-10-27 07:38:49.13083+00	\N
75dd289a-f72b-42fa-975e-ee10cd796135	d154a3f1-7086-439f-b343-3998d6521efa		t	2022-10-27 07:39:23.932688+00	2022-10-27 07:39:23.932688+00	\N
f47ad001-2fbf-49bd-948d-e5c7fa373712	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd		f	2022-10-27 18:50:35.895337+00	2022-10-27 18:50:35.895337+00	\N
723bd96d-ea4e-44d3-8052-a1579f32b216	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd		f	2022-10-27 18:50:54.908079+00	2022-10-27 18:50:54.908079+00	\N
533f8773-0034-42b0-9269-33bc73ae9cd2	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd		f	2022-10-27 18:51:19.473581+00	2022-10-27 18:51:19.473581+00	\N
28a5bd8a-318a-4acf-b3c9-8ba04be5a979	789cbced-9141-4748-94d3-93476d276057		f	2022-10-27 20:19:43.093638+00	2022-10-27 20:19:43.093638+00	\N
7e3eeef8-4748-483c-bbf8-3767943135ee	789cbced-9141-4748-94d3-93476d276057		f	2022-10-27 20:20:05.222471+00	2022-10-27 20:20:05.222471+00	\N
5e16c816-a24a-42a4-92a8-8f765e72a149	28a5bd8a-318a-4acf-b3c9-8ba04be5a979		f	2022-10-27 20:20:44.316185+00	2022-10-27 20:20:44.316185+00	\N
57d072c8-4952-44c5-845e-d2d706677e16	7e3eeef8-4748-483c-bbf8-3767943135ee		f	2022-10-27 20:21:12.997082+00	2022-10-27 20:21:12.997082+00	\N
d154a3f1-7086-439f-b343-3998d6521efa	\N	uploads/category/0a624e64-811a-4dde-a968-e69724427dd7.png	t	2022-10-27 07:35:14.821001+00	2023-01-08 12:17:19.426111+00	\N
44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	\N	uploads/category/1e7621c4-14c2-46f3-a884-2af5f7549adc.png	f	2022-10-27 18:49:57.62672+00	2023-01-18 09:08:11.189134+00	\N
789cbced-9141-4748-94d3-93476d276057	\N	uploads/category/e15130f5-4bc4-4b62-8f69-a0c0e6ec1688.png	f	2022-10-27 20:19:00.352561+00	2023-01-18 09:08:11.189134+00	\N
a6b395b6-01de-4d0a-a8bf-542ece2eef3a	\N	uploads/category/95bf23db-b2ec-4133-b910-6f77f54a33a0.png	f	2023-01-08 12:30:18.222126+00	2023-01-18 09:08:11.189134+00	\N
573ae82d-f7bc-4da2-926b-9a822d75a4a0	\N	uploads/category/9a810e7d-7c60-4d69-87af-ea5e3c0c2c24.png	f	2023-01-08 12:25:18.082623+00	2023-01-18 09:08:11.189134+00	\N
f04e644c-6a37-4c18-ac8a-a90f08599d71	\N	uploads/category/5e894220-e385-496d-8f59-084196803085.png	f	2023-01-08 12:42:27.577603+00	2023-01-18 09:08:11.189134+00	\N
f67e67e6-db32-45ce-86f2-b35f70dc5792	\N	uploads/category/6ddc4cba-3f2e-4d23-a016-53c1f51e9215.png	f	2023-01-08 12:44:05.721057+00	2023-01-18 09:08:11.189134+00	\N
05d315e2-2859-4b33-af12-43b7882e175e	\N	uploads/category/226293b0-cb11-4bd8-a8d6-b83bedb5830c.png	f	2023-01-08 12:45:24.871968+00	2023-01-18 09:08:11.189134+00	\N
e17f79c4-a118-4c9e-895a-769da3b9f243	\N	uploads/category/be392a60-2d4c-472e-9d5b-1ebfe963f16f.png	f	2023-01-08 12:47:43.637707+00	2023-01-18 09:08:11.189134+00	\N
d5e7a59e-b272-4a77-9a95-5efebee00eb0	\N	uploads/category/1eac1a11-22d6-45cb-8e9b-d784a5a4bdd5.png	f	2023-01-08 12:49:32.346804+00	2023-01-18 09:08:11.189134+00	\N
5797f94d-ffd4-484d-b4a6-359a0b0196df	\N	uploads/category/a0f7b1b4-41a5-496b-96a5-4212b5404d79.png	f	2023-01-08 10:59:20.666498+00	2023-01-18 09:08:11.189134+00	\N
7d605280-1fa9-4bae-91b5-160c309200aa	\N	uploads/category/14056f02-f36e-4cf7-9d3b-38310545dfac.png	f	2023-01-08 13:12:55.482788+00	2023-01-18 09:08:11.189134+00	\N
b969cd61-af6a-4bae-88c5-cbd3cdb36a53	\N	uploads/category/716a71c4-aa9e-4cd2-9124-0b1e70f2e324.png	f	2023-01-08 13:14:12.955254+00	2023-01-18 09:08:11.189134+00	\N
7581d22a-dbba-4e2a-bf0d-98eb3e9076d9	\N	uploads/category/48ff4298-b8b3-4b93-addd-b1e07a8a290e.png	f	2023-01-08 13:15:08.543619+00	2023-01-18 09:08:11.189134+00	\N
9a4cc9ad-948c-4949-8fd8-2babda70a2b9	\N	uploads/category/693abed2-2560-4709-93c5-03ce7e666531.png	f	2023-01-08 13:16:56.324285+00	2023-01-18 09:08:11.189134+00	\N
\.


--
-- Data for Name: category_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category_product (id, category_id, product_id, created_at, updated_at, deleted_at) FROM stdin;
48f453a1-525f-441c-8582-a41e328baa4d	75dd289a-f72b-42fa-975e-ee10cd796135	14d95413-2c8a-472f-8f89-9458dc1bde33	2023-01-25 05:18:32.597208+00	2023-01-25 05:18:32.597208+00	\N
210556dd-4742-44d7-b046-ce46f6d0bc05	d5e7a59e-b272-4a77-9a95-5efebee00eb0	14d95413-2c8a-472f-8f89-9458dc1bde33	2023-01-25 05:18:32.597208+00	2023-01-25 05:18:32.597208+00	\N
\.


--
-- Data for Name: company_address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_address (id, lang_id, address, created_at, updated_at, deleted_at) FROM stdin;
75706251-06ea-41c1-905f-95ed8b4132f8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Улица Азади 23, Ашхабад	2022-06-22 13:44:50.239558+00	2022-06-22 13:44:50.239558+00	\N
d2c66808-e5fe-435f-ba01-cb717f80d9e0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	adres_tm	2022-06-22 13:44:50.21776+00	2022-08-22 04:33:42.14835+00	\N
\.


--
-- Data for Name: company_phone; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_phone (id, phone, created_at, updated_at, deleted_at) FROM stdin;
96c30e15-274c-49a0-bcc5-e2f8deac248f	+993 12 227475	2022-09-29 07:49:06.569246+00	2022-09-29 07:49:06.569246+00	\N
\.


--
-- Data for Name: company_setting; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.company_setting (id, logo, favicon, email, instagram, created_at, updated_at, deleted_at, imo) FROM stdin;
7d193677-e0b1-4df0-be88-dc6e16a47ca7	uploads/logode9c4f45-acba-42ce-b435-e744631a98ba.jpeg	uploads/favicon8a413c02-108d-4d2f-8e92-d24a18cea1d3.jpeg	isleg-bazar@gmail.com	@islegbazarinstagram	2022-06-15 14:57:04.54457+00	2022-12-14 10:11:38.550502+00	\N	+99362946805
\.


--
-- Data for Name: customer_address; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customer_address (id, customer_id, address, created_at, updated_at, deleted_at, is_active) FROM stdin;
41e86110-cc92-442e-b788-81959e56f668	19cdcf1a-f110-4510-a52b-063329d98607	Mir 2/2 jay 7 oy 36	2022-11-01 06:36:50.869926+00	2022-11-01 06:36:50.869926+00	\N	f
226a81d9-f317-43cd-808e-76fb3b54e783	95595b3c-c1ca-4363-b9c0-9916ce88b82a	sdfgsdfhsdfghdf	2023-01-08 13:31:51.302567+00	2023-01-08 13:31:51.302567+00	\N	f
b556ebd9-cae4-4265-bf80-cba92cb39da7	e9db1ea7-5ce4-4284-bdc2-0fc39da2b7f2	qwkjdwidwnfjkfnwefnjwef	2023-01-13 17:18:08.808941+00	2023-01-13 17:18:08.808941+00	\N	f
71dd6ebd-bdb2-4cc2-8010-d5a2771f2192	392e6586-b3aa-4086-92d9-35e1dc29253e	Mir 2/2	2023-01-15 14:41:59.97412+00	2023-01-15 14:41:59.97412+00	\N	f
27fe094a-0093-4837-9f36-5e251f8b77c9	eb132b01-c9e7-4836-af89-1c2184439544	,ksnedkjewnf	2023-01-26 06:50:59.62467+00	2023-01-26 06:50:59.62467+00	\N	f
3a52b3d9-4bd3-4d36-8dcc-31ea123540fe	ec5da332-10e5-4c82-9195-a2479a200c25	jabdjhegfh	2023-01-26 06:52:53.847892+00	2023-01-26 06:52:53.847892+00	\N	f
d3496ac8-f36f-40f8-8b40-0ba3a7e226e5	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	Mir 6/3 jay 56 	2022-11-22 05:02:14.214333+00	2023-01-27 04:06:06.424171+00	\N	t
f9e7f338-aad1-4f25-aef9-7c597979f346	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	Hitrowka	2023-01-26 19:09:11.693+00	2023-01-27 04:06:06.436138+00	\N	f
567546cf-85e4-482d-9ad1-0c551982aaf5	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	Howdan	2023-01-26 19:17:12.00161+00	2023-01-27 04:06:06.436138+00	\N	f
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (id, full_name, phone_number, password, birthday, gender, created_at, updated_at, deleted_at, email, is_register) FROM stdin;
19cdcf1a-f110-4510-a52b-063329d98607	Allanur Bayramgeldiyew	+99362420377	$2a$14$QLQ.Mkd6Oi3Qz4djp38KS.Y1BBwKNJL1Hy6qKS0piHnoNP4rvIMd2	\N	\N	2022-11-01 06:33:32.61818+00	2022-11-01 06:33:32.61818+00	\N	abb@gmail.com	t
1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	Muhammetmyrat Bayramov	+99363747155	$2a$14$Ag0N0Otwyu7qmHaDCVVmWOz2UxHsYhqoEMkZcnCgzMzB1rAGqMZO2	1997-03-11	\N	2022-11-02 02:34:24.403632+00	2022-11-28 03:39:44.990107+00	\N	muhammetmyrat@gmail.com	t
1f949ae7-f59b-4846-9b3f-f9ed1938174c	Kemal	+99362766780	$2a$14$6qycGfXrbnySmfTLBroXE.lXMJuF2VT.9Arn9F34SzQM9qtFB6GiW	\N	\N	2023-01-04 10:31:22.767892+00	2023-01-04 10:31:22.767892+00	\N	Kemal@gmail.com	t
37e44e45-ed7e-4ff8-8167-886595349855	dsgjhdgfjhsdfg	+99363463456	$2a$14$vmmZV8x31Otg25QbOTAyae1enF1Leo0fzjK4JdaNWUUx6SMgLxZFi	\N	\N	2023-01-04 17:44:12.445059+00	2023-01-04 17:44:12.445059+00	\N	awgserhgdfh@gmail.com	t
938fb63a-e7c8-41a7-b2d6-624156094d8f	sdagarfgadf	+99363254634	$2a$14$26kRxXVAglPoYNrzynb2tuM9k4LF1NfVYFZkfpJTChqVA/5.VLAx2	\N	\N	2023-01-08 11:19:22.357132+00	2023-01-08 11:19:22.357132+00	\N	345gdfbndfg@gmail.com	t
95595b3c-c1ca-4363-b9c0-9916ce88b82a	Kemal	+99361899737	$2a$14$q2UJe9rq1Q9Mb/yLYBP3POzZrRuSWMQ43Z02y6ARZ5nE8fQot8Qb6	\N	\N	2023-01-08 13:31:34.981358+00	2023-01-08 13:31:34.981358+00	\N	kemalhanow97@gmail.com	t
e9db1ea7-5ce4-4284-bdc2-0fc39da2b7f2	Allanur	+99365684712	\N	\N	\N	2023-01-13 17:18:08.805908+00	2023-01-13 17:18:08.805908+00	\N	\N	f
a28db6ca-17a6-4879-ba0b-4741633d1395	Oraz	+99362986819	$2a$14$k/3rRQzdZWN026OZg5HCCeUZzb2QbHO8FZB7v8CdGNNlMmEOGZBIm	\N	\N	2023-01-14 11:34:16.599034+00	2023-01-14 11:34:16.599034+00	\N	orazdurdyyew3762@gmail.com	t
392e6586-b3aa-4086-92d9-35e1dc29253e	Serdar Berdiyew	+99365432187	\N	\N	\N	2023-01-15 14:41:59.970261+00	2023-01-15 14:41:59.970261+00	\N	\N	f
dad96943-886e-4ee3-9ca5-3723423b4191	Eziz	+99363558110	$2a$14$IZz8MRoQNJkMhC4hJ9ey7uLnrYNAQ0cakkf52lJA6mDgx/s6bODmC	\N	\N	2023-01-15 15:04:02.885068+00	2023-01-15 15:04:02.885068+00	\N	eziz@gmail.com	t
75cc7b9d-96b6-4189-a62e-1acd6f16851d	dnkjndkjndkwed	+99363747656	$2a$14$WPsa3EMQWzA5gZavgRmInuWMsPzpBgmWk39lh1Mys45gkN6/D6F4u	\N	\N	2023-01-23 15:30:56.851327+00	2023-01-23 15:30:56.851327+00	\N	ewjdnjwkednw@gmail.com	t
f12aec9a-7e10-4c6b-8121-4b5ce334bfa9	jkbdjcbewjhbdwebh	+99363744664	$2a$14$.wtw3vCqNxVizxEqy1w1ruld8hk9V5sFvmuisv.hThfMpFHtF90fm	\N	\N	2023-01-23 15:48:42.721254+00	2023-01-23 15:48:42.721254+00	\N	wejdwkedwehb@gmail.com	t
878502bc-3aaf-4ffc-9b8c-1423efb829dc	w3jeduw3h32	+99363747154	$2a$14$B8RCoD0zD1VOoZ4sjFgDPeLoOVl.mzw/2MMx9MGMiyQfpblor.Cle	\N	\N	2023-01-23 15:53:02.024195+00	2023-01-23 15:53:02.024195+00	\N	23oei23jen@gmail.com	t
12c4d76a-e3a6-4f35-97ba-efed264f849a	54jkwnefjkewbhd	+99363745454	$2a$14$mriNhEaRJYqLbrpygPJeUOG1MKgKLHPdmDE/CX9arGhcLSpR0Ovsy	\N	\N	2023-01-23 15:54:38.267322+00	2023-01-23 15:54:38.267322+00	\N	wejdweu@gmail.com	t
eb132b01-c9e7-4836-af89-1c2184439544	jshfjkew	+99365454121	\N	\N	\N	2023-01-26 06:50:59.620772+00	2023-01-26 06:50:59.620772+00	\N	\N	f
ec5da332-10e5-4c82-9195-a2479a200c25	jewdkjweb	+99363265326	\N	\N	\N	2023-01-26 06:52:53.844755+00	2023-01-26 06:52:53.844755+00	\N	\N	f
\.


--
-- Data for Name: district; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.district (id, price, created_at, updated_at, deleted_at) FROM stdin;
a58294d3-efe5-4cb7-82d3-8df8c37563c5	15	2022-06-25 05:23:25.640364+00	2022-06-25 05:23:25.640364+00	\N
\.


--
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.images (id, product_id, image, created_at, updated_at, deleted_at) FROM stdin;
abe6b7ca-5982-4ff4-815b-37d30e22d515	14d95413-2c8a-472f-8f89-9458dc1bde33	uploads/product/fcc38775-f791-4f8f-9b24-2bce41200729.jpg	2023-01-25 05:18:32.564523+00	2023-01-25 05:18:32.564523+00	\N
7d51bb70-d816-43ea-ad40-27a7348895cc	14d95413-2c8a-472f-8f89-9458dc1bde33	uploads/product/271ebb94-45a8-4578-a91a-67d2886f9efd.jpg	2023-01-25 05:18:32.564523+00	2023-01-25 05:18:32.564523+00	\N
\.


--
-- Data for Name: languages; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.languages (id, name_short, flag, created_at, updated_at, deleted_at) FROM stdin;
aea98b93-7bdf-455b-9ad4-a259d69dc76e	ru	uploads/language/22a2ad57-4686-44d2-aded-01261006d2be.png	2022-06-15 14:53:21.29491+00	2023-01-08 12:08:20.10097+00	\N
8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	tm	uploads/language/f1ce871a-07b1-4199-97e2-3b213229085c.png	2022-06-15 14:53:06.041686+00	2023-01-08 12:08:57.877665+00	\N
\.


--
-- Data for Name: likes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.likes (id, product_id, customer_id, created_at, updated_at, deleted_at) FROM stdin;
3251b463-7870-403c-bb16-e442e4b2e5c7	14d95413-2c8a-472f-8f89-9458dc1bde33	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	2023-01-26 05:07:03.878968+00	2023-01-26 05:07:03.878968+00	\N
\.


--
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.notifications (id, name, created_at, updated_at, deleted_at) FROM stdin;
f832e5da-d969-43d7-9cd0-7eae6c6c59e9	sargyt_ucin	2022-11-08 18:08:35.4564+00	2022-11-08 18:08:35.4564+00	\N
\.


--
-- Data for Name: order_dates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_dates (id, date, created_at, updated_at, deleted_at) FROM stdin;
c1f2beca-a6b6-4971-a6a7-ed50079c6912	tomorrow	2022-09-28 12:36:46.804343+00	2022-09-28 12:36:46.804343+00	\N
32646376-c93f-412b-9e75-b3a5fa70df9e	today	2022-09-28 12:35:33.772335+00	2022-11-01 18:58:35.636948+00	\N
\.


--
-- Data for Name: order_times; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_times (id, order_date_id, "time", created_at, updated_at, deleted_at) FROM stdin;
7d47a77a-b8f3-4e96-aa56-5ec7fb328e86	c1f2beca-a6b6-4971-a6a7-ed50079c6912	09:00 - 12:00	2022-09-28 12:36:46.825964+00	2022-09-28 12:36:46.825964+00	\N
de31361b-9fba-48f2-9341-9e3dd08cf9fd	c1f2beca-a6b6-4971-a6a7-ed50079c6912	18:00 - 21:00	2022-09-28 12:36:46.825964+00	2022-09-28 12:36:46.825964+00	\N
67c488ef-6021-4cc5-96cc-25408e71dbe3	32646376-c93f-412b-9e75-b3a5fa70df9e	09:00 - 12:00	2022-11-01 17:51:27.754592+00	2022-11-01 18:58:35.636948+00	\N
861ae017-67b5-45ae-88d7-a6990d7c49fd	32646376-c93f-412b-9e75-b3a5fa70df9e	18:00 - 21:00	2022-11-01 17:51:27.767647+00	2022-11-01 18:58:35.636948+00	\N
\.


--
-- Data for Name: ordered_products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ordered_products (id, product_id, quantity_of_product, order_id, created_at, updated_at, deleted_at) FROM stdin;
31380923-020c-476b-8393-d65b5f878256	14d95413-2c8a-472f-8f89-9458dc1bde33	9	4171967e-d9e6-47bc-bc93-250603e1ec95	2023-01-25 20:17:31.165542+00	2023-01-25 20:17:31.165542+00	\N
45444772-0c7f-4556-b44d-b1c9baffdd9d	14d95413-2c8a-472f-8f89-9458dc1bde33	7	478588e5-28ce-4a95-be74-a8e3cbbf9021	2023-01-25 20:19:55.21329+00	2023-01-25 20:19:55.21329+00	\N
1c6fbaf9-3697-45a2-8e4c-753de5c58399	14d95413-2c8a-472f-8f89-9458dc1bde33	3	b95c90ab-5238-4539-99c0-67d7b1535302	2023-01-25 20:22:00.673763+00	2023-01-25 20:22:00.673763+00	\N
7d578802-d6f7-4f13-ac73-d8e612a267da	14d95413-2c8a-472f-8f89-9458dc1bde33	1	792c651e-9ca5-430d-b7e0-0bde7756b4c1	2023-01-25 20:26:22.085655+00	2023-01-25 20:26:22.085655+00	\N
11042686-0159-4546-a1d3-db12a9b569e5	14d95413-2c8a-472f-8f89-9458dc1bde33	48	4d9c0319-cf86-4b9a-9b89-e2db2ebab627	2023-01-26 06:50:59.634617+00	2023-01-26 06:50:59.634617+00	\N
0c469c5a-901d-430d-b878-9c6de765cbe0	14d95413-2c8a-472f-8f89-9458dc1bde33	32	8565bf0e-bb0d-4e91-9a2c-8e619635fa40	2023-01-26 06:52:53.855888+00	2023-01-26 06:52:53.855888+00	\N
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, customer_id, customer_mark, order_time, payment_type, total_price, created_at, updated_at, deleted_at, order_number, shipping_price, excel, address) FROM stdin;
4171967e-d9e6-47bc-bc93-250603e1ec95	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a		18:00 - 21:00	nagt_tm	900	2023-01-25 20:17:31.160054+00	2023-01-25 20:17:31.216885+00	\N	64	0	uploads/orders/64.xlsx	Mir 6/3 jay 56 
6e94ad1b-863e-4041-89d9-0b700c59107d	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a		09:00 - 12:00	töleg terminaly	900	2023-01-25 20:18:06.695167+00	2023-01-25 20:18:06.711328+00	\N	65	0	uploads/orders/65.xlsx	Mir 6/3 jay 56 
478588e5-28ce-4a95-be74-a8e3cbbf9021	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a		09:00 - 12:00	töleg terminaly	700	2023-01-25 20:19:55.209153+00	2023-01-25 20:19:55.259347+00	\N	66	0	uploads/orders/66.xlsx	Mir 6/3 jay 56 
b95c90ab-5238-4539-99c0-67d7b1535302	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a		09:00 - 12:00	töleg terminaly	300	2023-01-25 20:22:00.668768+00	2023-01-25 20:22:00.703872+00	\N	67	0	uploads/orders/67.xlsx	Mir 6/3 jay 56 
792c651e-9ca5-430d-b7e0-0bde7756b4c1	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a		09:00 - 12:00	nagt_tm	100	2023-01-25 20:26:22.080437+00	2023-01-25 20:26:22.125813+00	\N	68	0	uploads/orders/68.xlsx	Mir 6/3 jay 56 
4d9c0319-cf86-4b9a-9b89-e2db2ebab627	eb132b01-c9e7-4836-af89-1c2184439544		09:00 - 12:00	töleg terminaly	4800	2023-01-26 06:50:59.630586+00	2023-01-26 06:50:59.674048+00	\N	69	0	uploads/orders/69.xlsx	,ksnedkjewnf
8565bf0e-bb0d-4e91-9a2c-8e619635fa40	ec5da332-10e5-4c82-9195-a2479a200c25		09:00 - 12:00	töleg terminaly	3200	2023-01-26 06:52:53.852555+00	2023-01-26 06:52:53.896105+00	\N	70	0	uploads/orders/70.xlsx	jabdjhegfh
\.


--
-- Data for Name: payment_types; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.payment_types (id, lang_id, type, created_at, updated_at, deleted_at) FROM stdin;
83e6589c-0cb6-4267-bcc5-e06cc93b36d8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	наличные	2022-09-20 09:33:50.780468+00	2022-09-20 09:33:50.780468+00	\N
7a6a313d-8fcd-4c56-9fa5-aefb12552b82	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	töleg terminaly	2022-09-20 09:34:46.329459+00	2022-09-20 09:34:46.329459+00	\N
cb7e8cc9-9b2e-4cd8-921f-91b3bb5e5564	aea98b93-7bdf-455b-9ad4-a259d69dc76e	платежный терминал	2022-09-20 09:34:46.359276+00	2022-09-20 09:34:46.359276+00	\N
38696743-82e5-4644-9c86-4a99ae45f912	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	nagt_tm	2022-09-20 09:33:50.755689+00	2022-09-20 09:40:04.959827+00	\N
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, brend_id, price, old_price, amount, created_at, updated_at, deleted_at, limit_amount, is_new, shop_id, main_image) FROM stdin;
14d95413-2c8a-472f-8f89-9458dc1bde33	9b838628-fd75-4232-862d-998635f24f52	100	101	1000000	2023-01-25 05:18:32.560308+00	2023-01-26 07:25:06.298016+00	\N	50	t	74cce5dc-6fc2-487c-8553-1f00850df257	uploads/product/8ec51cf3-6770-42b3-912b-a1226f234996.JPG
\.


--
-- Data for Name: shops; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shops (id, owner_name, address, phone_number, running_time, created_at, updated_at, deleted_at) FROM stdin;
74cce5dc-6fc2-487c-8553-1f00850df257	Owez Myradow	Asgabat saher Mir 4/1 jay 2 magazyn 56	+99361254689	7:00-21:00	2022-11-09 09:02:42.724172+00	2022-11-09 10:09:48.522999+00	\N
a283d9a4-f38e-43ee-a228-6584b7406cc4	Arslan Kerimow	Asgabat saher Mir 2/2 jay 2 magazyn 23	+99362420387	8:00-22:00	2022-11-09 08:53:47.243923+00	2022-11-09 10:11:14.496257+00	\N
\.


--
-- Data for Name: translation_about; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_about (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
7abeb5cf-2fbb-43b9-94ca-251dd5f40d5a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Sizi Isleg onlaýn marketimizde hoş gördük!	Onlaýn marketimiz 2019-njy ýylyň iýul aýyndan bäri hyzmat berýär. Häzirki wagtda Size ýüzlerçe brendlere degişli bolan müňlerçe haryt görnüşlerini hödürleýäris! Haryt görnüşlerimizi sizden gelýän isleg we teklipleriň esasynda köpeltmäge dowam edýäris. Biziň maksadymyz müşderilerimize ýokary hilli hyzmat bermek bolup durýar. Indi Siz öýüňizden çykmazdan özüňizi gerekli zatlar bilen üpjün edip bilersiňiz! Munuň bilen bir hatarda Siz wagtyňyzy we transport çykdajylaryny hem tygşytlaýarsyňyz. Tölegi harytlar size gowuşandan soňra nagt ýa-da bank kartlarynyň üsti bilen amala aşyryp bilersiňiz!\n\nBiziň gapymyz hyzmatdaşlyklara we tekliplere hemişe açyk!	2022-06-25 07:07:15.62033+00	2022-06-25 07:07:15.62033+00	\N
e50bb3d1-14a1-400e-83d9-8bc15969b914	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Рады приветствовать Вас в интернет-маркете Isleg!	Мы начали работу в июле 2019 года и на сегодняшний день мы предлагаем Вам тысячи видов товаров, которые принадлежат сотням брендам. Каждый день мы работаем над увеличением ассортимента, привлечением новых компаний к сотрудничеству. Целью нашей работы является создание выгодных условий для наших клиентов-экономия времени на походы в магазины, оплата наличными или картой, доставка в удобное время, и конечно же качественная продукция по лучшим ценам!\n\nМы открыты для сотрудничества и пожеланий!	2022-06-25 07:07:15.653744+00	2022-06-25 07:07:15.653744+00	\N
\.


--
-- Data for Name: translation_afisa; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_afisa (id, afisa_id, lang_id, title, description, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: translation_basket_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_basket_page (id, lang_id, quantity_of_goods, total_price, discount, delivery, total, to_order, your_basket, created_at, updated_at, deleted_at, empty_the_basket, empty_the_like_page) FROM stdin;
51b3699e-1c7b-442a-be7b-6b2ad1f111b4	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Haryt mukdary	Doly bahasy	Arzanladyş	Eltip berme	Jemi	Sargyt et	Sebet	2022-08-30 07:36:24.978404+00	2023-01-27 06:38:00.127486+00	\N	Siziň sebediňiz boş	Halan harytlaryňyzyň sanawy boş
456dcb5a-fabb-47f8-b216-0cddd3077124	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Количество товара	Полная стоимость	Скидка	Доставкa	Общее	Заказ	Корзина	2022-08-30 07:36:24.978404+00	2023-01-27 06:39:27.405953+00	\N	Ваша корзина пуста	У вас нет любимого продукта
\.


--
-- Data for Name: translation_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_category (id, lang_id, category_id, name, created_at, updated_at, deleted_at) FROM stdin;
c6086874-26c3-4ea3-bb70-3c88bf67643b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d7862d17-0742-4bd5-8fc8-478fd7e868c4	Sowgatlyk toplumlar	2022-10-27 07:38:27.36353+00	2022-10-27 07:38:27.36353+00	\N
c9f9ac63-172d-450f-80c0-eecfba4284d1	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d7862d17-0742-4bd5-8fc8-478fd7e868c4	Подарочные наборы	2022-10-27 07:38:27.377387+00	2022-10-27 07:38:27.377387+00	\N
0cf94c44-aab0-4180-8e07-b076babf5865	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	71994790-1b7b-41ab-90a8-b3df0d68e3e6	Aksiýadaky harytlar	2022-10-27 07:38:49.186311+00	2022-10-27 07:38:49.186311+00	\N
fc576a2d-3ab3-420c-856b-704daf0cc3ed	aea98b93-7bdf-455b-9ad4-a259d69dc76e	71994790-1b7b-41ab-90a8-b3df0d68e3e6	Продукция в категории Акции	2022-10-27 07:38:49.200519+00	2022-10-27 07:38:49.200519+00	\N
3a7f44fb-5e22-416b-9bdc-f20be8485b1b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	75dd289a-f72b-42fa-975e-ee10cd796135	Täze harytlar	2022-10-27 07:39:23.997701+00	2022-10-27 07:39:23.997701+00	\N
25b0fa42-8466-4cd0-a322-b1d02332b918	aea98b93-7bdf-455b-9ad4-a259d69dc76e	75dd289a-f72b-42fa-975e-ee10cd796135	Новые продукты	2022-10-27 07:39:24.012575+00	2022-10-27 07:39:24.012575+00	\N
ed1bfdf4-e479-4c1c-8bd4-5d2e8459d05b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f47ad001-2fbf-49bd-948d-e5c7fa373712	Miweler	2022-10-27 18:50:35.915849+00	2022-10-27 18:50:35.915849+00	\N
ef104ee3-9004-40dd-b7d7-4220c65e5783	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f47ad001-2fbf-49bd-948d-e5c7fa373712	Фрукты	2022-10-27 18:50:35.987027+00	2022-10-27 18:50:35.987027+00	\N
42dc017a-ec97-46d8-aa3e-1b8caeb3c920	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	723bd96d-ea4e-44d3-8052-a1579f32b216	Gök otlar	2022-10-27 18:50:54.9831+00	2022-10-27 18:50:54.9831+00	\N
2f7f0906-7d14-4bbf-9d54-fe6bb53ff155	aea98b93-7bdf-455b-9ad4-a259d69dc76e	723bd96d-ea4e-44d3-8052-a1579f32b216	Зелень\n	2022-10-27 18:50:54.997372+00	2022-10-27 18:50:54.997372+00	\N
a2a49708-3863-46fd-a364-f273893c48b9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	533f8773-0034-42b0-9269-33bc73ae9cd2	Gök önümler	2022-10-27 18:51:19.516869+00	2022-10-27 18:51:19.516869+00	\N
68cb6490-4a07-4242-b42f-aee6c72c0266	aea98b93-7bdf-455b-9ad4-a259d69dc76e	533f8773-0034-42b0-9269-33bc73ae9cd2	Овощи	2022-10-27 18:51:19.53169+00	2022-10-27 18:51:19.53169+00	\N
9e665bec-e985-49cc-92f5-8b9312f49670	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	Kosmetika	2022-10-27 20:19:43.143664+00	2022-10-27 20:19:43.143664+00	\N
cb61081e-e86e-41d8-bece-7be188b25a85	aea98b93-7bdf-455b-9ad4-a259d69dc76e	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	Косметика	2022-10-27 20:19:43.15858+00	2022-10-27 20:19:43.15858+00	\N
358b416e-ddfc-422e-b382-e379809e4854	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7e3eeef8-4748-483c-bbf8-3767943135ee	Diş saglygy we arassaçylygy	2022-10-27 20:20:05.288219+00	2022-10-27 20:20:05.288219+00	\N
142572af-4ed1-4036-b9d6-85b7e3eeeceb	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7e3eeef8-4748-483c-bbf8-3767943135ee	Здоровье и чистота зубов	2022-10-27 20:20:05.302795+00	2022-10-27 20:20:05.302795+00	\N
538296a0-e2cf-4b7e-a4fb-164b7faf4eda	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	5e16c816-a24a-42a4-92a8-8f765e72a149	Makiaž	2022-10-27 20:20:44.388622+00	2022-10-27 20:20:44.388622+00	\N
07517ceb-d7a1-481e-88a9-39cb10eda852	aea98b93-7bdf-455b-9ad4-a259d69dc76e	5e16c816-a24a-42a4-92a8-8f765e72a149	Макияж	2022-10-27 20:20:44.404066+00	2022-10-27 20:20:44.404066+00	\N
b4c2a314-4687-47e0-8a5a-cad0b922b191	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	57d072c8-4952-44c5-845e-d2d706677e16	Diş pastasy	2022-10-27 20:21:13.055793+00	2022-10-27 20:21:13.055793+00	\N
426ae99c-25c0-4de1-81aa-a8122d081e4b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	57d072c8-4952-44c5-845e-d2d706677e16	Зубная паста	2022-10-27 20:21:13.075148+00	2022-10-27 20:21:13.075148+00	\N
feaf5735-1815-4567-8239-8b6e9346e554	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7d605280-1fa9-4bae-91b5-160c309200aa	Öý haýwanlary	2023-01-08 13:12:55.485361+00	2023-01-08 13:12:55.485361+00	\N
0a7bb101-babb-4f5a-b5d9-30972595fb5e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7d605280-1fa9-4bae-91b5-160c309200aa	Домашние животные	2023-01-08 13:12:55.490105+00	2023-01-08 13:12:55.490105+00	\N
870e94b3-7f07-4e2b-9a2e-e5d132ad7deb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	b969cd61-af6a-4bae-88c5-cbd3cdb36a53	Mebel	2023-01-08 13:14:12.95748+00	2023-01-08 13:14:12.95748+00	\N
0970c7b0-c935-4504-a717-b4f3857ca6d6	aea98b93-7bdf-455b-9ad4-a259d69dc76e	b969cd61-af6a-4bae-88c5-cbd3cdb36a53	Мебель	2023-01-08 13:14:12.961279+00	2023-01-08 13:14:12.961279+00	\N
fc397dde-64f2-471b-b5d3-5de7ef697121	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7581d22a-dbba-4e2a-bf0d-98eb3e9076d9	Saglyk	2023-01-08 13:15:08.548158+00	2023-01-08 13:15:08.548158+00	\N
f4f96f28-6156-4987-9994-69f11c10f25d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7581d22a-dbba-4e2a-bf0d-98eb3e9076d9	Здоровье	2023-01-08 13:15:08.553463+00	2023-01-08 13:15:08.553463+00	\N
50594841-dc20-40cf-9575-a690c95be1f9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	9a4cc9ad-948c-4949-8fd8-2babda70a2b9	Hyzmatlar	2023-01-08 13:16:56.328572+00	2023-01-08 13:16:56.328572+00	\N
23b79fa3-cb54-425d-8883-596d1c0a6705	aea98b93-7bdf-455b-9ad4-a259d69dc76e	9a4cc9ad-948c-4949-8fd8-2babda70a2b9	Услуги	2023-01-08 13:16:56.336167+00	2023-01-08 13:16:56.336167+00	\N
5ef15568-e39e-4a3a-bf80-3695ae6e5367	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d154a3f1-7086-439f-b343-3998d6521efa	Arzanladyş we Aksiýalar	2022-10-27 07:35:14.838875+00	2023-01-08 12:17:19.438081+00	\N
00638721-b67f-41fd-b332-cdd96f66bf0c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d154a3f1-7086-439f-b343-3998d6521efa	Распродажи и Акции	2022-10-27 07:35:14.853749+00	2023-01-08 12:17:19.446354+00	\N
9dc67a0a-f78a-4094-9b40-faafd55e87b1	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	Iýmit	2022-10-27 18:49:57.648716+00	2023-01-08 12:19:26.649996+00	\N
906373de-0e0b-4f74-b6d8-5d20d7c9a53d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	Питания	2022-10-27 18:49:57.66415+00	2023-01-08 12:19:26.65804+00	\N
09388897-a1b0-4a89-abf0-68a9528f7e61	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	789cbced-9141-4748-94d3-93476d276057	Şahsy ideg	2022-10-27 20:19:00.465181+00	2023-01-08 12:22:33.068596+00	\N
00cb656d-3639-4895-a71e-f44f93f0bd89	aea98b93-7bdf-455b-9ad4-a259d69dc76e	789cbced-9141-4748-94d3-93476d276057	Косметика	2022-10-27 20:19:00.493038+00	2023-01-08 12:22:33.078608+00	\N
52e62a02-7965-4021-8dea-f1e23d50ffb6	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	a6b395b6-01de-4d0a-a8bf-542ece2eef3a	Arassaçylyk	2023-01-08 12:30:18.225211+00	2023-01-08 12:30:18.225211+00	\N
7fbccd06-1bb9-44d5-ba86-52f8d5897208	aea98b93-7bdf-455b-9ad4-a259d69dc76e	a6b395b6-01de-4d0a-a8bf-542ece2eef3a	Чистота	2023-01-08 12:30:18.230027+00	2023-01-08 12:30:18.230027+00	\N
20ca970c-4763-4f9f-aa64-d1d0a232594c	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	573ae82d-f7bc-4da2-926b-9a822d75a4a0	Gurluşyk harytlary	2023-01-08 12:25:18.085496+00	2023-01-08 12:34:34.961153+00	\N
9927e053-2b77-4d01-ac18-05059ac5be3c	aea98b93-7bdf-455b-9ad4-a259d69dc76e	573ae82d-f7bc-4da2-926b-9a822d75a4a0	Стройматериалы	2023-01-08 12:25:18.091383+00	2023-01-08 12:34:34.970239+00	\N
a2ce2d66-b984-4fb6-b847-efe46b71765b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f04e644c-6a37-4c18-ac8a-a90f08599d71	Awtoulag	2023-01-08 12:42:27.584041+00	2023-01-08 12:42:27.584041+00	\N
abe4f8e2-60ef-43ec-9703-cfed3355825d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f04e644c-6a37-4c18-ac8a-a90f08599d71	Aвтомобиль	2023-01-08 12:42:27.590423+00	2023-01-08 12:42:27.590423+00	\N
91bfe695-ba1d-40a4-8e0d-a0a8e74e4eef	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f67e67e6-db32-45ce-86f2-b35f70dc5792	Bilim	2023-01-08 12:44:05.724576+00	2023-01-08 12:44:05.724576+00	\N
6cdca87e-5bed-47b7-8025-52e626839ece	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f67e67e6-db32-45ce-86f2-b35f70dc5792	Образование	2023-01-08 12:44:05.730148+00	2023-01-08 12:44:05.730148+00	\N
e7b63dab-f824-48cc-9d6b-15419205c27e	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	05d315e2-2859-4b33-af12-43b7882e175e	Egin-eşik	2023-01-08 12:45:24.876644+00	2023-01-08 12:45:24.876644+00	\N
fded73ff-c5b8-4846-88f8-7cbe7758a578	aea98b93-7bdf-455b-9ad4-a259d69dc76e	05d315e2-2859-4b33-af12-43b7882e175e	Одежда	2023-01-08 12:45:24.886171+00	2023-01-08 12:45:24.886171+00	\N
6e71de96-68e7-4e13-97ec-973a2d6e6105	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	e17f79c4-a118-4c9e-895a-769da3b9f243	Hojalyk	2023-01-08 12:47:43.643268+00	2023-01-08 12:47:43.643268+00	\N
78c0db6a-56ae-4e34-98b1-da47bf364ea1	aea98b93-7bdf-455b-9ad4-a259d69dc76e	e17f79c4-a118-4c9e-895a-769da3b9f243	Хозяйственные	2023-01-08 12:47:43.650051+00	2023-01-08 12:47:43.650051+00	\N
9f8e2036-7e2b-4746-83b7-49db7d215497	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d5e7a59e-b272-4a77-9a95-5efebee00eb0	Konselýariýa	2023-01-08 12:49:32.350447+00	2023-01-08 12:49:32.350447+00	\N
86e4fca1-0e46-4283-9e3a-81ccb02805dc	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d5e7a59e-b272-4a77-9a95-5efebee00eb0	Kанцелярия	2023-01-08 12:49:32.356138+00	2023-01-08 12:49:32.356138+00	\N
0b2ea6f6-a7f9-4217-a9c8-d00eda3dfb60	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	5797f94d-ffd4-484d-b4a6-359a0b0196df	Elektronika	2023-01-08 10:59:20.66978+00	2023-01-08 13:10:02.959523+00	\N
2a53ee62-d41b-43c5-8d95-a667b6f83265	aea98b93-7bdf-455b-9ad4-a259d69dc76e	5797f94d-ffd4-484d-b4a6-359a0b0196df	Электроника	2023-01-08 10:59:20.676556+00	2023-01-08 13:10:02.971264+00	\N
\.


--
-- Data for Name: translation_contact; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_contact (id, lang_id, full_name, email, phone, letter, company_phone, imo, company_email, instagram, created_at, updated_at, deleted_at, button_text) FROM stdin;
73253999-7355-42b4-8700-94de76f0058a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Doly Adyňyz	Email	Telefon	Hatyňyz	Telefon belgimiz	Imo	E-mail	Instagram	2022-06-27 06:29:47.914891+00	2023-01-26 18:26:20.379585+00	\N	Ugrat
f1693167-0c68-4a54-9831-56f124d629a3	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Ваше полное имя	Эл. адрес	Телефон	Твое письмо	Наш номер телефона	Имо	Эл. адрес	Инстаграм	2022-06-27 06:29:48.050553+00	2023-01-26 18:31:03.105358+00	\N	Отправить
\.


--
-- Data for Name: translation_district; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_district (id, lang_id, district_id, name, created_at, updated_at, deleted_at) FROM stdin;
ad9f94d3-05e7-43b3-aa77-7b7f3754d003	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	a58294d3-efe5-4cb7-82d3-8df8c37563c5	Parahat 2	2022-06-25 05:23:25.712337+00	2022-06-25 05:23:25.712337+00	\N
aa1cfa48-3132-4dd4-abfb-070a2986690b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	a58294d3-efe5-4cb7-82d3-8df8c37563c5	Mir 2	2022-06-25 05:23:25.774504+00	2022-06-25 05:23:25.774504+00	\N
\.


--
-- Data for Name: translation_footer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_footer (id, lang_id, about, payment, contact, secure, word, created_at, updated_at, deleted_at) FROM stdin;
84b5504f-1056-4b44-94dd-a7819148da66	aea98b93-7bdf-455b-9ad4-a259d69dc76e	О нас	Порядок доставки и оплаты	Коммуникация	Обслуживания и Политика Конфиденциальности	Все права защищены	2022-06-22 10:23:32.793161+00	2022-06-22 10:23:32.793161+00	\N
12dc4c16-5712-4bff-a957-8e16d450b4fb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Biz Barada	Eltip bermek we töleg tertibi	Aragatnaşyk	Ulanyş düzgünleri we gizlinlik şertnamasy	Ähli hukuklary goraglydyr	2022-06-22 10:23:32.716064+00	2022-06-22 10:23:32.716064+00	\N
\.


--
-- Data for Name: translation_header; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_header (id, lang_id, research, phone, password, forgot_password, sign_in, sign_up, name, password_verification, verify_secure, my_information, my_favorites, my_orders, log_out, created_at, updated_at, deleted_at, basket, email, add_to_basket, add_button) FROM stdin;
eaf206e6-d515-4bdb-9323-a047cd0edae5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Gözleg	Telefon	Parol	Açar sözümi unutdym	Ulgama girmek	Agza bolmak	Ady	Açar sözi tassyklamak	Ulanyş Düzgünlerini we Gizlinlik Şertnamasyny okadym we kabul edýärin	Maglumatym	Halanlarym	Sargytlarym	Çykmak	2022-06-15 23:48:26.460534+00	2023-01-26 18:50:20.020483+00	\N	Sebet	Mail adres	Sebede goş	Goşmak
9154e800-2a92-47de-b4ff-1e63b213e5f7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Поиск	Телефон	Пароль	Забыл пароль	Войти	Зарегистрироваться	Имя	Подтвердить Пароль	Я прочитал и принимаю Условия Обслуживания и Политика Конфиденциальности	Моя информация	Мои любимые	Мои заказы	Выйти	2022-06-15 23:48:26.491672+00	2023-01-26 18:50:56.407227+00	\N	Корзина	Почта	Добавить в корзину	Добавить
\.


--
-- Data for Name: translation_my_information_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_my_information_page (id, lang_id, address, created_at, updated_at, deleted_at, birthday, update_password, save, gender, male, female) FROM stdin;
11074158-69f2-473a-b4fe-94304ff0d8a7	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Salgyňyz	2022-07-04 14:28:46.529935+00	2023-01-26 12:17:39.361424+00	\N	Doglan senäň	Açar sözi üýtget	Ýatda sakla	Jynsy	Oglan	Gyz
d294138e-b808-41ae-9ac5-1826751fda3d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Ваш адрес	2022-07-04 14:28:46.603058+00	2023-01-26 12:21:46.5462+00	\N	Дата рождения	Изменить пароль	Запомнить	Пол	Мужчина	Женщина
\.


--
-- Data for Name: translation_my_order_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_my_order_page (id, lang_id, orders, date, price, image, name, brend, product_price, amount, total_price, created_at, updated_at, deleted_at) FROM stdin;
ff43b90d-e22d-4364-b358-6fd56bb3a305	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Edilen sargytlar	Senesi	Bahasy	Surat	Ady	Brendy	Bahasy	Sany	Jemi	2022-09-02 08:04:39.36328+00	2023-01-26 13:50:31.058384+00	\N
6f30b588-94d8-49f5-a558-a90c2ec9150e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Заказы	Дата	Цена	Рисунок	Имя	Бренди	Цена	Количество	Общая цена	2022-09-02 08:04:39.394714+00	2023-01-26 18:16:58.168531+00	\N
\.


--
-- Data for Name: translation_notification; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_notification (id, notification_id, lang_id, translation, created_at, updated_at, deleted_at) FROM stdin;
bb82aa0f-dd88-49cd-9f6a-bd27b2930505	f832e5da-d969-43d7-9cd0-7eae6c6c59e9	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Ваш заказ успешно получен	2022-11-08 18:08:35.531894+00	2022-11-08 18:08:35.531894+00	\N
b0e087dd-60c8-4f47-8755-95e9678d4405	f832e5da-d969-43d7-9cd0-7eae6c6c59e9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	siziň sargydyňyz üstünlikli kabul edildi	2022-11-08 18:08:35.616593+00	2022-11-08 18:08:35.616593+00	\N
\.


--
-- Data for Name: translation_order_dates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_order_dates (id, lang_id, order_date_id, date, created_at, updated_at, deleted_at) FROM stdin;
1aa5185f-9815-4e3f-9c34-718bfb587d91	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c1f2beca-a6b6-4971-a6a7-ed50079c6912	Ertir	2022-09-28 12:36:46.836838+00	2022-09-28 12:36:46.836838+00	\N
9e7a3752-fce2-4b66-bf3e-d915bf463f92	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c1f2beca-a6b6-4971-a6a7-ed50079c6912	Завтра	2022-09-28 12:36:46.847888+00	2022-09-28 12:36:46.847888+00	\N
3338d831-f091-4574-a0bf-f9cb07dd4893	aea98b93-7bdf-455b-9ad4-a259d69dc76e	32646376-c93f-412b-9e75-b3a5fa70df9e	Segodnya	2022-09-28 12:35:33.82453+00	2022-11-01 18:58:35.636948+00	\N
dcd0c70b-9fa2-4327-8b35-de29bd3febcb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	32646376-c93f-412b-9e75-b3a5fa70df9e	Su gun	2022-09-28 12:35:33.812812+00	2022-11-01 18:58:35.636948+00	\N
\.


--
-- Data for Name: translation_order_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_order_page (id, lang_id, content, type_of_payment, choose_a_delivery_time, your_address, mark, to_order, created_at, updated_at, deleted_at) FROM stdin;
75810722-07fd-400e-94b4-cd230de08cbf	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Eltip bermek hyzmaty Aşgabat şäheriniň çägi bilen bir hatarda Büzmeýine we Änew şäherine hem elýeterlidir. Hyzmat mugt amala aşyrylýar. Saýtdan sargyt edeniňizden soňra operator size jaň edip sargydy tassyklar (eger hemişelik müşderi bolsaňyz sargytlaryňyz islegiňize görä awtomatik usulda hem tassyklanýar). Sargydy barlap alanyňyzdan soňra töleg amala aşyrylýar. Eltip berijiniň size gowşurýan töleg resminamasynda siziň tölemeli puluňyz bellenendir. Töleg nagt we nagt däl görnüşde milli manatda amala aşyrylýar. Kabul edip tölegini geçiren harydyňyz yzyna alynmaýar	Töleg şekili	Eltip berme wagtyny saýlaň	Salgyňyz	Bellik	Sargyt et	2022-09-01 07:47:16.720956+00	2023-01-26 06:33:07.714052+00	\N
474a15e9-1a05-49aa-9a61-c92837d9c9a8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Помимо Ашхабада, услуга доставки доступна в Бузмеин и Анью. Услуга предоставляется бесплатно. После того, как вы сделаете заказ с сайта, вам позвонит оператор для подтверждения заказа (если вы постоянный клиент, ваш заказ при желании будет подтвержден автоматически). Оплата производится после подтверждения заказа. Платежный документ, который дает вам поставщик, является суммой, которую вы должны заплатить. Оплата производится в манатах в наличной и безналичной форме. Ваша покупка после принятия и оплаты не подлежит возврату	Форма оплаты	Выберите время доставки	Ваш адресс	Примечание	Разместить заказ	2022-09-01 07:47:16.802639+00	2023-01-26 06:35:14.87903+00	\N
\.


--
-- Data for Name: translation_payment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_payment (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
5748ec03-5278-425c-babf-f7f2bf8d2efa	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	Eltip bermek we töleg tertibi	Eltip bermek hyzmaty Aşgabat şäheriniň çägi bilen bir hatarda Büzmeýine we Änew şäherine hem elýeterlidir. Hyzmat mugt amala aşyrylýar;\nHer bir sargydyň jemi bahasy azyndan 150 manat bolmalydyr;\nSaýtdan sargyt edeniňizden soňra operator size jaň edip sargydy tassyklar (eger hemişelik müşderi bolsaňyz sargytlaryňyz islegiňize görä awtomatik usulda hem tassyklanýar);\nGirizen salgyňyz we telefon belgiňiz esasynda hyzmat amala aşyrylýar;\nSargyt tassyklanmadyk ýagdaýynda ol hasaba alynmaýar we ýerine ýetirilmeýär. Sargydyň tassyklanmagy üçin girizen telefon belgiňizden jaň kabul edip bilýändigiňize göz ýetiriň. Şeýle hem girizen salgyňyzyň dogrulygyny barlaň;\nSargydy barlap alanyňyzdan soňra töleg amala aşyrylýar. Eltip berijiniň size gowşurýan töleg resminamasynda siziň tölemeli puluňyz bellenendir. Töleg nagt we nagt däl görnüşde milli manatda amala aşyrylýar. Kabul edip tölegini geçiren harydyňyz yzyna alynmaýar;\nSargyt tassyklanandan soňra 24 sagadyň dowamynda eýesi tapylmasa ol güýjüni ýitirýär;	2022-06-25 06:37:47.362666+00	2022-06-25 06:37:47.362666+00	\N
ea7f4c0c-4b1a-41d3-94eb-e058aba9c99f	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Порядок доставки и оплаты	В настоящее время услуга по доставке осуществляется по городу Ашхабад, Бюзмеин и Анау. Услуга предоставляется бесплатно.\nМинимальный заказ должен составлять не менее 150 манат;\nПосле Вашего заказа по сайту, оператор позвонит Вам для подтверждения заказа (постоянным клиентам по их желанию подтверждение осуществляется автоматизированно);\nУслуга доставки выполняется по указанному Вами адресу и номеру телефона;\nЕсли заказ не подтвержден то данный заказ не регистрируется и не выполняется. Для подтверждения заказа, удостоверьтесь, что можете принять звонок по указанному Вами номеру телефона. Также проверьте правильность указанного Вами адреса;\nОплата выполняется после того, как Вы проверите и примите заказ. На платежном документе курьера указана сумма Вашей оплаты. Оплата выполняется наличными и через карту в национальной валюте. Принятый и оплаченный товар возврату не подлежит;\nЕсли не удается найти владельца заказа в течение 24 часов после подтверждения заказа, то данный заказ аннулируется;	2022-06-25 06:37:47.39047+00	2022-06-25 06:37:47.39047+00	\N
\.


--
-- Data for Name: translation_product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_product (id, lang_id, product_id, name, description, created_at, updated_at, deleted_at, slug) FROM stdin;
98621106-ba5c-410f-a9c0-a37d97460047	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	14d95413-2c8a-472f-8f89-9458dc1bde33	"Bianyo" firmaň çyzgy üçin niýetlenen gara galamy	12 sany naborly karton gutyda - ýokary hilli ajaýyp kombinasiýa. Urga çydamly we aňsat arassalanýan berk agaç	2023-01-25 05:18:32.58376+00	2023-01-25 05:18:32.58376+00	\N	bianyo-firman-cyzgy-ucin-niyetlenen-gara-galamy
93450fc6-1de4-4bfc-b407-d7f5bfef9578	aea98b93-7bdf-455b-9ad4-a259d69dc76e	14d95413-2c8a-472f-8f89-9458dc1bde33	Набор карандашей чернографитных фирма "Bianyo"	Набор карандашей чернографитных, 12 шт., "Bianyo" гранённые, заточенные, в картонной коробке — прекрасное сочетание высокого качества. Прочный стержень не крошится, а деревянный корпус легко затачивается	2023-01-25 05:18:32.59094+00	2023-01-25 05:18:32.59094+00	\N	nabor-karandashei-chernografitnykh-firma-bianyo
\.


--
-- Data for Name: translation_secure; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_secure (id, lang_id, title, content, created_at, updated_at, deleted_at) FROM stdin;
3579a847-ce74-4fbe-b10d-8aba83867857	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Пользовательское соглашение	Между Ынамдар – Интернет Маркетом (далее – “Ынамдар”) и интернет сайтом www.ynamdar.com (далее – “Сайт”), а также его клиентом (далее - “Клиент”) достигнуто соглашение по нижеследующим условиям.\n	2022-06-25 05:46:54.221498+00	2022-06-25 05:46:54.221498+00	\N
5988b64a-82ad-4ed0-bd1b-bdd0b3b05912	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ÖZARA YLALAŞYGY	Ynamdar - Internet Marketi (Mundan beýläk – “Ynamdar”) we www.ynamdar.com internet saýty (Mundan beýläk – “Saýt”) bilen, onuň agzasynyň (“Agza”) arasynda aşakdaky şertleri ýerine ýetirmek barada ylalaşyga gelindi.	2022-06-25 05:46:54.190131+00	2022-06-25 05:46:54.190131+00	\N
\.


--
-- Data for Name: translation_update_password_page; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_update_password_page (id, lang_id, title, verify_password, explanation, save, created_at, updated_at, deleted_at, password) FROM stdin;
5190ca93-7007-4db4-8105-65cc3b1af868	aea98b93-7bdf-455b-9ad4-a259d69dc76e	изменить пароль	Подтвердить Пароль	ключевое слово должно быть буквой или цифрой длиной от 5 до 20	запомнить	2022-07-05 05:35:08.984141+00	2022-07-05 05:35:08.984141+00	\N	ключевое слово
de12082b-baab-4b83-ac07-119df09d1230	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	açar sözi üýtgetmek	açar sözi tassykla	siziň açar sözüňiz 5-20 uzynlygynda harp ýa-da sandan ybarat bolmalydyr	ýatda sakla	2022-07-05 05:35:08.867617+00	2022-07-05 05:35:08.867617+00	\N	açar sözi
\.


--
-- Name: orders_order_number_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_order_number_seq', 70, true);


--
-- Name: admins admins_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.admins
    ADD CONSTRAINT admins_pkey PRIMARY KEY (id);


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
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


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
-- Name: translation_notification translation_notification_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_notification
    ADD CONSTRAINT translation_notification_pkey PRIMARY KEY (id);


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
-- Name: admins updated_admins_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_admins_updated_at BEFORE UPDATE ON public.admins FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


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
-- Name: notifications updated_notifications_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_notifications_updated_at BEFORE UPDATE ON public.notifications FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


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
-- Name: translation_notification updated_translation_notification_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_translation_notification_updated_at BEFORE UPDATE ON public.translation_notification FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


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
-- Name: translation_notification languages_translation_notification; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_notification
    ADD CONSTRAINT languages_translation_notification FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_order_dates languages_translation_order_dates; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_order_dates
    ADD CONSTRAINT languages_translation_order_dates FOREIGN KEY (lang_id) REFERENCES public.languages(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: translation_notification notifications_translation_notification; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.translation_notification
    ADD CONSTRAINT notifications_translation_notification FOREIGN KEY (notification_id) REFERENCES public.notifications(id) ON UPDATE CASCADE ON DELETE CASCADE;


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
-- Name: ordered_products products_ordered_products; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ordered_products
    ADD CONSTRAINT products_ordered_products FOREIGN KEY (product_id) REFERENCES public.products(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: products shops_products; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT shops_products FOREIGN KEY (shop_id) REFERENCES public.shops(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

