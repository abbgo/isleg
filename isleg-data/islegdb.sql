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
UPDATE products SET deleted_at = now() WHERE shop_id = s_id;
UPDATE translation_product SET deleted_at = now() FROM products WHERE translation_product.product_id = products.id AND products.shop_id = s_id;
UPDATE main_image SET deleted_at = now() FROM products WHERE main_image.product_id = products.id AND products.shop_id = s_id;
UPDATE images SET deleted_at = now() FROM products WHERE images.product_id = products.id AND products.shop_id = s_id;
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
UPDATE products SET deleted_at = NULL WHERE shop_id = s_id;
UPDATE translation_product SET deleted_at = NULL FROM products WHERE translation_product.product_id = products.id AND products.shop_id = s_id;
UPDATE main_image SET deleted_at = NULL FROM products WHERE main_image.product_id = products.id AND products.shop_id = s_id;
UPDATE images SET deleted_at = NULL FROM products WHERE images.product_id = products.id AND products.shop_id = s_id;
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
-- Name: admins; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.admins (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    full_name character varying,
    phone_number character varying,
    password character varying,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    type character varying
);


ALTER TABLE public.admins OWNER TO postgres;

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
    total_price numeric,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    order_number integer NOT NULL,
    shipping_price numeric,
    excel character varying,
    address character varying DEFAULT 'uytget'::character varying
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
    is_new boolean DEFAULT false,
    shop_id uuid
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
-- Name: translation_notification; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.translation_notification (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    notification_id uuid,
    lang_id uuid,
    translation character varying,
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
-- Data for Name: admins; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.admins (id, full_name, phone_number, password, created_at, updated_at, deleted_at, type) FROM stdin;
5e0e9a0a-e07a-4911-bbc0-8dad29a6abbd	Allanur Bayramgeldiyev	+99362420377	$2a$14$w5DuoaqTJEdPpT3fzrLdh.Q3BGm/wX76NFGPGhy9G4nU/fDUXvtPy	2022-11-02 21:46:19.827563+05	2022-11-02 21:46:19.827563+05	\N	super_admin
c97bfc6a-fd85-4aa8-82db-e788f6b0d70a	Muhammet Bayramov	+99363747155	$2a$14$IXSdYxI0f.qQ8kDuLq.DU.F4ZRnMuq58VErTjaFNdquqzcZaenImu	2022-11-02 21:48:30.593337+05	2022-11-02 21:48:30.593337+05	\N	super_admin
6989254b-79c7-412c-acb2-19f67a3277d5	Seyit Batyrov	+99361111111	$2a$14$qD9HYvqPHUVITgfAJLfj3uODG.hcGiI7.ayv3jc1NlSD34QA5drv2	2022-11-02 22:30:44.622854+05	2022-11-02 22:30:44.622854+05	\N	admin
42ba1c9b-f56d-44e4-a72b-5031d0f3ce64	Kakajan Batyrov	+99362222222	$2a$14$QxhaMFkztA7x1jlR9PT3zOwCystNe/7QDEp7C04xLNFp2MDMCgpHK	2022-11-02 22:31:53.959553+05	2022-11-02 22:31:53.959553+05	\N	admin
54f172bc-e13b-4dd8-95af-c3e364f09e3b	Maya Kerimova	+99363333333	$2a$14$tR7soepOOiWQ0jtHpmTLy.kYOzMfbu2RNk2iICMeYHFqwiI1AvKBm	2022-11-02 22:43:54.030609+05	2022-11-02 22:43:54.030609+05	\N	super_admin
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
b96b4869-c152-48e1-9c32-54d906fbe689	45a9f186-2521-4eef-a4e0-b5c253c70878	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	1	2022-11-21 09:20:21.645886+05	2022-11-21 09:46:32.685328+05	\N
68cf8e16-0cfd-45df-bbfa-4f25ac481cc8	32055a0a-2d59-45a9-89b0-761d1f6ad047	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	2	2022-11-21 09:46:10.30059+05	2022-11-21 09:46:34.957809+05	\N
a26e3bd2-06b9-4c69-bbc1-750ea14141ea	83da5c7b-bffe-4450-97c9-0f376441b1d4	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	2	2022-11-22 09:45:46.362287+05	2022-11-22 10:15:21.439857+05	\N
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
44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	\N	uploads/category/bcf63001-13d5-4630-8a94-11a7b6154b28.png	f	2022-10-27 23:49:57.62672+05	2022-10-27 23:49:57.62672+05	\N
f47ad001-2fbf-49bd-948d-e5c7fa373712	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd		f	2022-10-27 23:50:35.895337+05	2022-10-27 23:50:35.895337+05	\N
723bd96d-ea4e-44d3-8052-a1579f32b216	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd		f	2022-10-27 23:50:54.908079+05	2022-10-27 23:50:54.908079+05	\N
533f8773-0034-42b0-9269-33bc73ae9cd2	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd		f	2022-10-27 23:51:19.473581+05	2022-10-27 23:51:19.473581+05	\N
789cbced-9141-4748-94d3-93476d276057	\N	uploads/category/a13801b4-aa3a-48a6-9f6f-1809f469a59e.png	f	2022-10-28 01:19:00.352561+05	2022-10-28 01:19:00.352561+05	\N
28a5bd8a-318a-4acf-b3c9-8ba04be5a979	789cbced-9141-4748-94d3-93476d276057		f	2022-10-28 01:19:43.093638+05	2022-10-28 01:19:43.093638+05	\N
7e3eeef8-4748-483c-bbf8-3767943135ee	789cbced-9141-4748-94d3-93476d276057		f	2022-10-28 01:20:05.222471+05	2022-10-28 01:20:05.222471+05	\N
5e16c816-a24a-42a4-92a8-8f765e72a149	28a5bd8a-318a-4acf-b3c9-8ba04be5a979		f	2022-10-28 01:20:44.316185+05	2022-10-28 01:20:44.316185+05	\N
57d072c8-4952-44c5-845e-d2d706677e16	7e3eeef8-4748-483c-bbf8-3767943135ee		f	2022-10-28 01:21:12.997082+05	2022-10-28 01:21:12.997082+05	\N
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
395621d1-8269-4fd3-8626-8a6712916d74	d154a3f1-7086-439f-b343-3998d6521efa	83da5c7b-bffe-4450-97c9-0f376441b1d4	2022-10-27 13:30:49.724642+05	2022-10-27 13:30:49.724642+05	\N
94a6f4a3-6d81-4c9e-a99e-b8dbfdc1b71a	d7862d17-0742-4bd5-8fc8-478fd7e868c4	83da5c7b-bffe-4450-97c9-0f376441b1d4	2022-10-27 13:30:49.724642+05	2022-10-27 13:30:49.724642+05	\N
ecc4b4b6-9fb7-4371-8df8-d4c780c53098	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	83da5c7b-bffe-4450-97c9-0f376441b1d4	2022-10-27 13:30:49.724642+05	2022-10-27 13:30:49.724642+05	\N
b83c9c2c-e6f5-4a04-9f4b-67e531a87e64	d154a3f1-7086-439f-b343-3998d6521efa	0946a0f5-d23f-4660-9151-80ef91ae9747	2022-10-27 13:32:30.048139+05	2022-10-27 13:32:30.048139+05	\N
0069bea4-e091-4894-82f2-0057952b2ce7	d7862d17-0742-4bd5-8fc8-478fd7e868c4	0946a0f5-d23f-4660-9151-80ef91ae9747	2022-10-27 13:32:30.048139+05	2022-10-27 13:32:30.048139+05	\N
d9d36d7e-bc22-4b4f-aef9-9375aca69223	d154a3f1-7086-439f-b343-3998d6521efa	03050bc6-6223-49f3-b729-397fd3b6b285	2022-10-27 13:36:08.830146+05	2022-10-27 13:36:08.830146+05	\N
14e3d2d1-eac0-45bd-9c2e-73ad51ad9c07	71994790-1b7b-41ab-90a8-b3df0d68e3e6	03050bc6-6223-49f3-b729-397fd3b6b285	2022-10-27 13:36:08.830146+05	2022-10-27 13:36:08.830146+05	\N
93325946-9d80-40a4-9761-f2d3a8d7393a	d154a3f1-7086-439f-b343-3998d6521efa	8b481e58-cd39-4761-a052-75e30124689a	2022-10-27 13:38:14.209569+05	2022-10-27 13:38:14.209569+05	\N
aeed9d4e-aa12-4874-859f-13cb5f178cfc	71994790-1b7b-41ab-90a8-b3df0d68e3e6	8b481e58-cd39-4761-a052-75e30124689a	2022-10-27 13:38:14.209569+05	2022-10-27 13:38:14.209569+05	\N
9703c0b6-8106-424d-b8cc-245f5da75b7d	d154a3f1-7086-439f-b343-3998d6521efa	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	2022-10-27 23:15:36.102749+05	2022-10-27 23:15:36.102749+05	\N
0a1a751b-ffa4-46f9-8ccd-2e3c60cfa24a	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	2022-10-27 23:15:36.102749+05	2022-10-27 23:15:36.102749+05	\N
c59db997-c4f9-4ed1-b596-7560787e775a	71994790-1b7b-41ab-90a8-b3df0d68e3e6	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	2022-10-27 23:15:36.102749+05	2022-10-27 23:15:36.102749+05	\N
21ac7f67-ccec-4527-aeb1-d5312b30a4ff	d154a3f1-7086-439f-b343-3998d6521efa	81462bfa-36df-4e09-aa46-c6fa1ab86de6	2022-10-27 23:17:42.840691+05	2022-10-27 23:17:42.840691+05	\N
1015b7f8-5f03-4479-bb01-a10abc974247	71994790-1b7b-41ab-90a8-b3df0d68e3e6	81462bfa-36df-4e09-aa46-c6fa1ab86de6	2022-10-27 23:17:42.840691+05	2022-10-27 23:17:42.840691+05	\N
3dc2735c-d5e6-48be-ab75-cd85a52592c7	d154a3f1-7086-439f-b343-3998d6521efa	360ebeac-853e-45a5-ab7f-838430b0c442	2022-10-27 23:19:20.813712+05	2022-10-27 23:19:20.813712+05	\N
77e24a17-c899-4556-89b5-c90cd9599ff3	75dd289a-f72b-42fa-975e-ee10cd796135	360ebeac-853e-45a5-ab7f-838430b0c442	2022-10-27 23:19:20.813712+05	2022-10-27 23:19:20.813712+05	\N
230c98aa-e238-4048-81df-19ff30320195	71994790-1b7b-41ab-90a8-b3df0d68e3e6	360ebeac-853e-45a5-ab7f-838430b0c442	2022-10-27 23:19:20.813712+05	2022-10-27 23:19:20.813712+05	\N
59a4c9bc-c2dc-45a5-bd4f-fe46bf15bdef	d154a3f1-7086-439f-b343-3998d6521efa	fe309360-c5dd-406a-9957-3d898ea85dfc	2022-10-27 23:20:14.847982+05	2022-10-27 23:20:14.847982+05	\N
df4d2c4c-d9d0-46c3-a0b0-d888114a1741	75dd289a-f72b-42fa-975e-ee10cd796135	fe309360-c5dd-406a-9957-3d898ea85dfc	2022-10-27 23:20:14.847982+05	2022-10-27 23:20:14.847982+05	\N
db3a3850-f5f5-4f99-9865-66d2afcf464b	71994790-1b7b-41ab-90a8-b3df0d68e3e6	fe309360-c5dd-406a-9957-3d898ea85dfc	2022-10-27 23:20:14.847982+05	2022-10-27 23:20:14.847982+05	\N
59026659-a8bf-42ca-8fd2-7c104023872b	d154a3f1-7086-439f-b343-3998d6521efa	febf699d-ca37-458a-b121-b5b70bbc7db0	2022-10-27 23:21:07.496807+05	2022-10-27 23:21:07.496807+05	\N
7bc4f65b-dada-4f0a-8390-144655d07b6d	75dd289a-f72b-42fa-975e-ee10cd796135	febf699d-ca37-458a-b121-b5b70bbc7db0	2022-10-27 23:21:07.496807+05	2022-10-27 23:21:07.496807+05	\N
604072c0-008e-47c7-817a-0a3d9fc7c8b1	71994790-1b7b-41ab-90a8-b3df0d68e3e6	febf699d-ca37-458a-b121-b5b70bbc7db0	2022-10-27 23:21:07.496807+05	2022-10-27 23:21:07.496807+05	\N
47ad0ddd-79a6-4b06-86b9-79cb7338c936	d154a3f1-7086-439f-b343-3998d6521efa	802b422b-710a-420b-860e-59b7f49d10bd	2022-10-27 23:22:15.741114+05	2022-10-27 23:22:15.741114+05	\N
8e14d85e-86ba-416d-9557-05de48cb641d	75dd289a-f72b-42fa-975e-ee10cd796135	802b422b-710a-420b-860e-59b7f49d10bd	2022-10-27 23:22:15.741114+05	2022-10-27 23:22:15.741114+05	\N
37996541-c786-4c00-8502-c57c6b7fa8b3	71994790-1b7b-41ab-90a8-b3df0d68e3e6	802b422b-710a-420b-860e-59b7f49d10bd	2022-10-27 23:22:15.741114+05	2022-10-27 23:22:15.741114+05	\N
0e676d4f-e8db-4adc-a230-51dd9b3b6f11	d154a3f1-7086-439f-b343-3998d6521efa	35f5f2d8-9271-469f-bde1-2314c18ea574	2022-10-27 23:23:19.488924+05	2022-10-27 23:23:19.488924+05	\N
aeb7cfe2-7d53-4669-8c82-c39ab8778604	75dd289a-f72b-42fa-975e-ee10cd796135	35f5f2d8-9271-469f-bde1-2314c18ea574	2022-10-27 23:23:19.488924+05	2022-10-27 23:23:19.488924+05	\N
6f37b390-dec0-48f5-9dca-e5ff3145120e	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	d987b7ad-257e-4ae2-befb-b7d369252a54	2022-10-27 23:53:01.858222+05	2022-10-27 23:53:01.858222+05	\N
587cc263-ab41-4a1e-87b3-9d800f227645	f47ad001-2fbf-49bd-948d-e5c7fa373712	d987b7ad-257e-4ae2-befb-b7d369252a54	2022-10-27 23:53:01.858222+05	2022-10-27 23:53:01.858222+05	\N
3183e7c7-c2d5-4104-a213-8c2cdb0e2a5d	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	77ecf422-b48b-45fd-8e58-380e23d74c4c	2022-10-28 00:08:25.834073+05	2022-10-28 00:08:25.834073+05	\N
cc5bd31b-b76b-45f8-b8eb-ff4548ba92e5	f47ad001-2fbf-49bd-948d-e5c7fa373712	77ecf422-b48b-45fd-8e58-380e23d74c4c	2022-10-28 00:08:25.834073+05	2022-10-28 00:08:25.834073+05	\N
ad658941-5122-4846-9936-72adf29acea9	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	9b9ef1ce-2f3d-4051-8e88-5c301bd68554	2022-10-28 00:09:12.523423+05	2022-10-28 00:09:12.523423+05	\N
6ca62718-079c-4640-971c-c40874f51294	f47ad001-2fbf-49bd-948d-e5c7fa373712	9b9ef1ce-2f3d-4051-8e88-5c301bd68554	2022-10-28 00:09:12.523423+05	2022-10-28 00:09:12.523423+05	\N
2222586d-6870-4f4c-ae5a-eae99c1ae792	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	20a8c487-9a56-4fb6-8fa0-13facaf96109	2022-10-28 00:09:53.625505+05	2022-10-28 00:09:53.625505+05	\N
b7460a96-873f-42b2-9910-097e46488592	f47ad001-2fbf-49bd-948d-e5c7fa373712	20a8c487-9a56-4fb6-8fa0-13facaf96109	2022-10-28 00:09:53.625505+05	2022-10-28 00:09:53.625505+05	\N
c322eae3-659e-4241-9a33-2bbc88f56883	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	7078e107-dd52-4da1-8007-29ed7cf731fb	2022-10-28 00:10:28.382834+05	2022-10-28 00:10:28.382834+05	\N
b85f5f1d-e5c3-48b9-b9ab-3a8d6d2810d6	f47ad001-2fbf-49bd-948d-e5c7fa373712	7078e107-dd52-4da1-8007-29ed7cf731fb	2022-10-28 00:10:28.382834+05	2022-10-28 00:10:28.382834+05	\N
b21d7db0-0a7b-45c2-9317-f73b62cd7352	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	5f6aba1c-66df-4791-b85e-b0a90ccffc20	2022-10-28 00:11:26.405889+05	2022-10-28 00:11:26.405889+05	\N
2847d349-79b2-48a5-87b5-be0ff614145b	723bd96d-ea4e-44d3-8052-a1579f32b216	5f6aba1c-66df-4791-b85e-b0a90ccffc20	2022-10-28 00:11:26.405889+05	2022-10-28 00:11:26.405889+05	\N
45fac4d2-2a82-49ed-a499-865873bc48d2	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	01dc8537-7ec1-4c48-bcce-3734f1ac598a	2022-10-28 00:11:58.385324+05	2022-10-28 00:11:58.385324+05	\N
b9c35c21-1d4c-408b-9747-b4786282b7ab	723bd96d-ea4e-44d3-8052-a1579f32b216	01dc8537-7ec1-4c48-bcce-3734f1ac598a	2022-10-28 00:11:58.385324+05	2022-10-28 00:11:58.385324+05	\N
e8841d5f-5f18-467a-b9ab-96681d41b3c5	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	214befd8-68bd-484a-a8d5-8e2d0b73931c	2022-10-28 00:12:34.843712+05	2022-10-28 00:12:34.843712+05	\N
3ec1a9b3-a6f5-47a9-97c4-b7cb41223911	723bd96d-ea4e-44d3-8052-a1579f32b216	214befd8-68bd-484a-a8d5-8e2d0b73931c	2022-10-28 00:12:34.843712+05	2022-10-28 00:12:34.843712+05	\N
28a3a520-8952-4b4f-ac02-02fe1d9d91fe	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	a316070c-409c-42c2-85df-bb7509a24c54	2022-10-28 00:13:05.587469+05	2022-10-28 00:13:05.587469+05	\N
5663182f-1804-4cad-89df-2df48ecf9015	723bd96d-ea4e-44d3-8052-a1579f32b216	a316070c-409c-42c2-85df-bb7509a24c54	2022-10-28 00:13:05.587469+05	2022-10-28 00:13:05.587469+05	\N
d4d15239-3e60-4208-8b7f-1e626977ddad	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	89172d2f-b5b3-4b26-a299-dc7e8a71d16e	2022-10-28 00:13:39.000719+05	2022-10-28 00:13:39.000719+05	\N
20c66dbf-69b7-4272-9c18-f9d641c532f4	723bd96d-ea4e-44d3-8052-a1579f32b216	89172d2f-b5b3-4b26-a299-dc7e8a71d16e	2022-10-28 00:13:39.000719+05	2022-10-28 00:13:39.000719+05	\N
7e1ab8c3-ec1e-4d9b-9f72-0d2a2bbc414d	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	c82fef0a-ad15-4b07-8855-910fc4708af1	2022-10-28 00:14:34.715942+05	2022-10-28 00:14:34.715942+05	\N
47e266ea-e072-48f5-950b-4166b0054443	723bd96d-ea4e-44d3-8052-a1579f32b216	c82fef0a-ad15-4b07-8855-910fc4708af1	2022-10-28 00:14:34.715942+05	2022-10-28 00:14:34.715942+05	\N
a53d2d1b-dd0e-45b7-9c8b-314e39593aa9	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	311aa4c1-6002-4acf-b1b5-e2aa7896def7	2022-10-28 00:15:19.629305+05	2022-10-28 00:15:19.629305+05	\N
d850652d-e357-4e80-bda3-bae6858eb523	533f8773-0034-42b0-9269-33bc73ae9cd2	311aa4c1-6002-4acf-b1b5-e2aa7896def7	2022-10-28 00:15:19.629305+05	2022-10-28 00:15:19.629305+05	\N
dedee2d4-659e-4c26-be7b-76a359711ad6	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	32055a0a-2d59-45a9-89b0-761d1f6ad047	2022-10-28 01:13:10.788073+05	2022-10-28 01:13:10.788073+05	\N
d43d4a79-6eb0-4d60-8e66-c8bd8d9140ab	533f8773-0034-42b0-9269-33bc73ae9cd2	32055a0a-2d59-45a9-89b0-761d1f6ad047	2022-10-28 01:13:10.788073+05	2022-10-28 01:13:10.788073+05	\N
f480d337-415c-4c91-92e9-01ca8ec86171	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	19e94bc7-398a-45d9-b3cf-5c8d50550e48	2022-10-28 01:13:56.406224+05	2022-10-28 01:13:56.406224+05	\N
6d552f8c-7ee7-48ee-9cb1-36b8d233d042	533f8773-0034-42b0-9269-33bc73ae9cd2	19e94bc7-398a-45d9-b3cf-5c8d50550e48	2022-10-28 01:13:56.406224+05	2022-10-28 01:13:56.406224+05	\N
b3b1d154-80ca-4606-b7aa-ac1598c4ae29	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	fd2c148f-70b5-48ba-8cf6-65f26438b46d	2022-10-28 01:15:13.871321+05	2022-10-28 01:15:13.871321+05	\N
4efb39df-2e52-448a-94d6-ec868f4b449a	533f8773-0034-42b0-9269-33bc73ae9cd2	fd2c148f-70b5-48ba-8cf6-65f26438b46d	2022-10-28 01:15:13.871321+05	2022-10-28 01:15:13.871321+05	\N
fa7ffddb-271d-408c-b68a-706daeeaa1c2	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	bcb52dfc-c957-4d5e-9bbc-1fcb607d3fd6	2022-10-28 01:15:51.325621+05	2022-10-28 01:15:51.325621+05	\N
a14c9771-3282-422b-bbfb-7fd379630621	533f8773-0034-42b0-9269-33bc73ae9cd2	bcb52dfc-c957-4d5e-9bbc-1fcb607d3fd6	2022-10-28 01:15:51.325621+05	2022-10-28 01:15:51.325621+05	\N
135eba09-4e30-4e65-8bbb-bc653f9920e9	789cbced-9141-4748-94d3-93476d276057	badd0869-99df-4df3-8a27-5e27c10a861d	2022-10-28 01:22:44.80739+05	2022-10-28 01:22:44.80739+05	\N
60763f36-5e09-4528-aef7-a1460a3a03b8	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	badd0869-99df-4df3-8a27-5e27c10a861d	2022-10-28 01:22:44.80739+05	2022-10-28 01:22:44.80739+05	\N
ebd8d48b-83ea-4faf-b4bf-ff01dc052e45	5e16c816-a24a-42a4-92a8-8f765e72a149	badd0869-99df-4df3-8a27-5e27c10a861d	2022-10-28 01:22:44.80739+05	2022-10-28 01:22:44.80739+05	\N
83ce2ee6-af54-4338-914e-a775dbcbf2c4	789cbced-9141-4748-94d3-93476d276057	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	2022-10-28 01:24:16.831459+05	2022-10-28 01:24:16.831459+05	\N
7bf5cb40-f59f-47b9-806e-db503df0f7fc	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	2022-10-28 01:24:16.831459+05	2022-10-28 01:24:16.831459+05	\N
32a85a29-2d39-48fd-ac05-2a498e5940b7	5e16c816-a24a-42a4-92a8-8f765e72a149	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	2022-10-28 01:24:16.831459+05	2022-10-28 01:24:16.831459+05	\N
52dc8810-0407-464f-9bdf-a6d85c40bb8c	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	2022-10-28 01:24:16.831459+05	2022-10-28 01:24:16.831459+05	\N
d52dec4a-bd78-4df6-82b7-19acb1014878	789cbced-9141-4748-94d3-93476d276057	45a9f186-2521-4eef-a4e0-b5c253c70878	2022-10-28 01:24:52.910549+05	2022-10-28 01:24:52.910549+05	\N
9ede1eaf-b4c5-4ade-915e-694e72228863	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	45a9f186-2521-4eef-a4e0-b5c253c70878	2022-10-28 01:24:52.910549+05	2022-10-28 01:24:52.910549+05	\N
7b9ffe2b-321e-4094-8d8b-26ac32bd2379	5e16c816-a24a-42a4-92a8-8f765e72a149	45a9f186-2521-4eef-a4e0-b5c253c70878	2022-10-28 01:24:52.910549+05	2022-10-28 01:24:52.910549+05	\N
7b42704a-dbd3-4f8d-9eb0-88a9bff47ef7	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	45a9f186-2521-4eef-a4e0-b5c253c70878	2022-10-28 01:24:52.910549+05	2022-10-28 01:24:52.910549+05	\N
7df6c589-1ee6-49aa-8204-b9420fc0d98f	789cbced-9141-4748-94d3-93476d276057	fa148eb2-520f-430e-bd8d-9d5a166d0600	2022-10-28 01:25:37.444609+05	2022-10-28 01:25:37.444609+05	\N
3f8ee642-3979-4a0d-9c42-3c3bf583192a	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	fa148eb2-520f-430e-bd8d-9d5a166d0600	2022-10-28 01:25:37.444609+05	2022-10-28 01:25:37.444609+05	\N
7b447e81-2987-47f0-82ce-b20424c24b06	5e16c816-a24a-42a4-92a8-8f765e72a149	fa148eb2-520f-430e-bd8d-9d5a166d0600	2022-10-28 01:25:37.444609+05	2022-10-28 01:25:37.444609+05	\N
b7493b49-73d7-4c80-92f1-36a5b869ca2e	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	fa148eb2-520f-430e-bd8d-9d5a166d0600	2022-10-28 01:25:37.444609+05	2022-10-28 01:25:37.444609+05	\N
dd448fd5-1e9e-4caf-834d-2f0a335a1c16	789cbced-9141-4748-94d3-93476d276057	9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	2022-10-28 01:26:46.993977+05	2022-10-28 01:26:46.993977+05	\N
106200d4-a2b1-4927-aab0-ff7716a9dda8	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	2022-10-28 01:26:46.993977+05	2022-10-28 01:26:46.993977+05	\N
bf272483-8b92-4fef-8184-e5ac6335de62	5e16c816-a24a-42a4-92a8-8f765e72a149	9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	2022-10-28 01:26:46.993977+05	2022-10-28 01:26:46.993977+05	\N
da1c15dd-1fe4-46c8-b6c0-f21ba7a2ad33	789cbced-9141-4748-94d3-93476d276057	74ec9e27-de0b-44d9-8036-3fae8be486a9	2022-10-28 01:27:54.417587+05	2022-10-28 01:27:54.417587+05	\N
5d107627-e0ec-45cc-8cfb-3b764836d442	7e3eeef8-4748-483c-bbf8-3767943135ee	74ec9e27-de0b-44d9-8036-3fae8be486a9	2022-10-28 01:27:54.417587+05	2022-10-28 01:27:54.417587+05	\N
72b94fdd-f438-419f-8931-dab1fa3de736	57d072c8-4952-44c5-845e-d2d706677e16	74ec9e27-de0b-44d9-8036-3fae8be486a9	2022-10-28 01:27:54.417587+05	2022-10-28 01:27:54.417587+05	\N
4f11ac49-e128-4cad-a8da-4e45b2ef4c97	789cbced-9141-4748-94d3-93476d276057	332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	2022-10-28 01:28:28.471997+05	2022-10-28 01:28:28.471997+05	\N
8dca84ea-658c-4f9a-8f2c-8080513a52b4	7e3eeef8-4748-483c-bbf8-3767943135ee	332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	2022-10-28 01:28:28.471997+05	2022-10-28 01:28:28.471997+05	\N
5f20ea2c-4991-4c4d-a7c1-430bd293475f	57d072c8-4952-44c5-845e-d2d706677e16	332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	2022-10-28 01:28:28.471997+05	2022-10-28 01:28:28.471997+05	\N
0d8daa5a-34c3-4d81-9791-e9941ae53578	789cbced-9141-4748-94d3-93476d276057	2eb8a13f-3edc-4422-b772-e57bfd8f8797	2022-10-28 01:29:09.063799+05	2022-10-28 01:29:09.063799+05	\N
0cc3e871-5cc4-45b4-9780-7f47b4481ea7	7e3eeef8-4748-483c-bbf8-3767943135ee	2eb8a13f-3edc-4422-b772-e57bfd8f8797	2022-10-28 01:29:09.063799+05	2022-10-28 01:29:09.063799+05	\N
74b12c89-7790-46f4-90a0-d4686c2a9358	57d072c8-4952-44c5-845e-d2d706677e16	2eb8a13f-3edc-4422-b772-e57bfd8f8797	2022-10-28 01:29:09.063799+05	2022-10-28 01:29:09.063799+05	\N
e28acbed-0577-4266-b9de-ae883ed50847	789cbced-9141-4748-94d3-93476d276057	18f957f2-216d-4810-b4d7-bd4dd49efd0d	2022-10-28 01:29:49.70851+05	2022-10-28 01:29:49.70851+05	\N
b1c8dbb0-1caf-4331-9e45-b777906b94d9	7e3eeef8-4748-483c-bbf8-3767943135ee	18f957f2-216d-4810-b4d7-bd4dd49efd0d	2022-10-28 01:29:49.70851+05	2022-10-28 01:29:49.70851+05	\N
4751429c-7ab4-4ff2-8a79-2fac616bba35	57d072c8-4952-44c5-845e-d2d706677e16	18f957f2-216d-4810-b4d7-bd4dd49efd0d	2022-10-28 01:29:49.70851+05	2022-10-28 01:29:49.70851+05	\N
f0f22da7-7bd4-444f-b006-06282cf0ab76	789cbced-9141-4748-94d3-93476d276057	ad24153a-997a-46d1-87bb-27aa1e3e8aea	2022-10-28 01:30:33.08544+05	2022-10-28 01:30:33.08544+05	\N
713e42c5-d389-42dc-b738-23b915adad02	7e3eeef8-4748-483c-bbf8-3767943135ee	ad24153a-997a-46d1-87bb-27aa1e3e8aea	2022-10-28 01:30:33.08544+05	2022-10-28 01:30:33.08544+05	\N
ce1dfb43-c4ae-4fde-be63-115bef5883c2	57d072c8-4952-44c5-845e-d2d706677e16	ad24153a-997a-46d1-87bb-27aa1e3e8aea	2022-10-28 01:30:33.08544+05	2022-10-28 01:30:33.08544+05	\N
ceffda14-f32c-4878-baa0-bb3563f4327d	789cbced-9141-4748-94d3-93476d276057	9c655c36-1832-48ca-9f88-c04197f191af	2022-10-28 01:31:07.066813+05	2022-10-28 01:31:07.066813+05	\N
14c04bed-1cc8-41ff-868d-2646986d4dd2	7e3eeef8-4748-483c-bbf8-3767943135ee	9c655c36-1832-48ca-9f88-c04197f191af	2022-10-28 01:31:07.066813+05	2022-10-28 01:31:07.066813+05	\N
0f01d1a6-3a83-4d83-8caf-a938f8c0cac2	57d072c8-4952-44c5-845e-d2d706677e16	9c655c36-1832-48ca-9f88-c04197f191af	2022-10-28 01:31:07.066813+05	2022-10-28 01:31:07.066813+05	\N
eced4828-b523-446b-af5b-b5b78e5053b5	d154a3f1-7086-439f-b343-3998d6521efa	ccb43083-1c9e-4e84-bffd-ecb28474165e	2022-10-27 13:29:08.279964+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05
2fd6bbe7-0b16-4605-a6e2-c8d82c62eff4	d7862d17-0742-4bd5-8fc8-478fd7e868c4	ccb43083-1c9e-4e84-bffd-ecb28474165e	2022-10-27 13:29:08.279964+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05
f961661a-de6d-4cfe-8c52-4686406cce6b	ab28ad8f-72af-4e9e-841b-38a6e6881a6e	ccb43083-1c9e-4e84-bffd-ecb28474165e	2022-10-27 13:29:08.279964+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05
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
41e86110-cc92-442e-b788-81959e56f668	19cdcf1a-f110-4510-a52b-063329d98607	Mir 2/2 jay 7 oy 36	2022-11-01 11:36:50.869926+05	2022-11-01 11:36:50.869926+05	\N	f
d3496ac8-f36f-40f8-8b40-0ba3a7e226e5	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	Mir 6/3 jay 56 	2022-11-22 10:02:14.214333+05	2022-11-22 10:02:14.214333+05	\N	t
d16fe7ec-9024-4745-8dab-0e79b13cc343	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	Mir 2/2 jay 7 oy 36	2022-11-03 22:00:31.880539+05	2022-11-22 10:02:14.271009+05	\N	f
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (id, full_name, phone_number, password, birthday, gender, created_at, updated_at, deleted_at, email, is_register) FROM stdin;
19cdcf1a-f110-4510-a52b-063329d98607	Allanur Bayramgeldiyew	+99362420377	$2a$14$QLQ.Mkd6Oi3Qz4djp38KS.Y1BBwKNJL1Hy6qKS0piHnoNP4rvIMd2	\N	\N	2022-11-01 11:33:32.61818+05	2022-11-01 11:33:32.61818+05	\N	abb@gmail.com	t
1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	Muhammetmyrat	+99363747155	$2a$14$Ag0N0Otwyu7qmHaDCVVmWOz2UxHsYhqoEMkZcnCgzMzB1rAGqMZO2	\N	\N	2022-11-02 07:34:24.403632+05	2022-11-02 07:34:24.403632+05	\N	bayramovmuhammetmyrat97@gmail.com	t
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
c1150204-74a8-4d1d-96bf-350c495ed4a0	81462bfa-36df-4e09-aa46-c6fa1ab86de6	uploads/product/cc72451b-28a7-4c37-b6df-fb2f980cf05b.jpg	uploads/product/e565562b-38c3-406b-abc4-9c6c7fe18b6a.jpg	2022-10-27 23:17:42.808526+05	2022-11-09 15:11:14.496257+05	\N
d2ed160f-8df5-472b-bd68-79839dbe1b02	81462bfa-36df-4e09-aa46-c6fa1ab86de6	uploads/product/bb42c64f-0e97-4683-b121-e319cf4726cb.jpg	uploads/product/ba398d68-ee79-4866-b93a-bae00978f3e1.jpg	2022-10-27 23:17:42.808526+05	2022-11-09 15:11:14.496257+05	\N
ea9ec679-e5fe-4e12-b74a-cd1f3e3dabb3	35f5f2d8-9271-469f-bde1-2314c18ea574	uploads/product/2c60a680-05b2-461f-9c45-9eaa159bcb0f.jpg	uploads/product/d0de12b7-26a7-4f18-a2af-13648d055f50.jpg	2022-10-27 23:23:19.432859+05	2022-11-09 15:11:14.496257+05	\N
7a7f59d9-881d-4ae2-9aad-52385c861ccc	35f5f2d8-9271-469f-bde1-2314c18ea574	uploads/product/1c51123c-feef-4aac-af55-d900d0887c5b.jpg	uploads/product/7c4aaa6e-5bef-4219-b107-18e7bfcc50a2.jpg	2022-10-27 23:23:19.432859+05	2022-11-09 15:11:14.496257+05	\N
abadfa0d-6999-40b1-b27b-b0084da7fcec	83da5c7b-bffe-4450-97c9-0f376441b1d4	uploads/product/162dec1d-64d9-4dda-84d2-0d065b7e0fe2.jpg	uploads/product/fda9bc4e-3727-48e8-aee2-b46737f0bf36.jpg	2022-10-27 13:30:49.679942+05	2022-11-09 15:09:41.955371+05	\N
f1eedc22-c922-4aea-93e0-c5b44449a23e	83da5c7b-bffe-4450-97c9-0f376441b1d4	uploads/product/bc01f470-6db7-4e73-bfe9-f6ea46ba34c6.jpg	uploads/product/b152ac0e-e7f2-4668-a3d8-abdcd118a2b2.jpg	2022-10-27 13:30:49.679942+05	2022-11-09 15:09:41.955371+05	\N
93d0b190-39e5-4419-89a3-c4b8d12b14c7	d987b7ad-257e-4ae2-befb-b7d369252a54	uploads/product/0b1d8d54-d7ec-445a-9536-254e74ccbf9d.jpg	uploads/product/d709e7ef-c65d-4cca-9275-7555f8ff0b58.jpg	2022-10-27 23:53:01.811227+05	2022-11-09 15:11:14.496257+05	\N
2382f52e-38c3-4e0b-81a0-b992ade243cb	d987b7ad-257e-4ae2-befb-b7d369252a54	uploads/product/79da974e-ca2c-447e-b345-db8305b363c8.jpg	uploads/product/669e5975-3cef-46d1-b3fa-ee9980c35575.jpg	2022-10-27 23:53:01.811227+05	2022-11-09 15:11:14.496257+05	\N
27e41045-ac6a-4fa8-953e-6c2576158f07	77ecf422-b48b-45fd-8e58-380e23d74c4c	uploads/product/dcb6fbeb-c8a0-4e80-b63e-bae92fbb0976.jpg	uploads/product/d329ff32-3386-4855-81c3-02a3fa745dba.jpg	2022-10-28 00:08:25.802076+05	2022-11-09 15:11:14.496257+05	\N
c13de883-9df2-4606-8e5c-fbd9749dcaf4	77ecf422-b48b-45fd-8e58-380e23d74c4c	uploads/product/01be5183-ab9b-4219-be51-6d30a0bdeb97.jpg	uploads/product/1a89dbd8-a0b0-4394-a49d-6c2a00a95766.jpg	2022-10-28 00:08:25.802076+05	2022-11-09 15:11:14.496257+05	\N
bd3b7206-488c-4151-89ab-ce75d8968147	9b9ef1ce-2f3d-4051-8e88-5c301bd68554	uploads/product/5b615aea-ebe8-4c8c-ad40-c599ef434b89.jpg	uploads/product/22ff7898-dbf6-471b-8608-7ce76204dc24.jpg	2022-10-28 00:09:12.485443+05	2022-11-09 15:11:14.496257+05	\N
5461340e-92fc-4845-b7cf-a6ada7d274ca	9b9ef1ce-2f3d-4051-8e88-5c301bd68554	uploads/product/a8b8e530-0418-4161-932a-f34cea1c86be.jpg	uploads/product/d5f7e4d3-9dea-46ca-aa82-5e0dda8e629b.jpg	2022-10-28 00:09:12.485443+05	2022-11-09 15:11:14.496257+05	\N
3bc24ef3-5f6e-40d9-b50d-2ec46bcbe668	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	uploads/product/1211f76d-734e-4c49-8688-b6b11406b2e8.jpg	uploads/product/e543ea1f-3001-438d-9b98-6a5ef1e7fa4f.jpg	2022-10-27 23:15:36.04716+05	2022-11-09 15:09:41.955371+05	\N
06b374a0-8af5-4cdf-b228-ee5668feb802	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	uploads/product/c175237d-bb04-42c2-a673-285c7297f671.jpg	uploads/product/fbb79b30-d4d8-4476-b584-59980e7484c4.jpg	2022-10-27 23:15:36.04716+05	2022-11-09 15:09:41.955371+05	\N
8202b1ec-dffd-49cb-8ec6-7857487ea9ab	20a8c487-9a56-4fb6-8fa0-13facaf96109	uploads/product/5921b54b-490e-4594-9d49-37b43acb7ea1.jpg	uploads/product/816742ca-efc9-485a-a363-5873086f6aaf.jpg	2022-10-28 00:09:53.581469+05	2022-11-09 15:11:14.496257+05	\N
a75ef72a-634c-47f5-8ca1-f5bc92a2fe22	20a8c487-9a56-4fb6-8fa0-13facaf96109	uploads/product/7c1acb87-c0dc-4127-81d0-dd6f02580485.jpg	uploads/product/bbe835cb-3738-4400-baf7-0181f8c6744d.jpg	2022-10-28 00:09:53.581469+05	2022-11-09 15:11:14.496257+05	\N
56400764-56ba-450c-8f61-f892a74720fd	360ebeac-853e-45a5-ab7f-838430b0c442	uploads/product/d93f2aed-5e17-4e7c-aa80-f6f47edb28ad.jpg	uploads/product/75592e9f-1f90-41b6-9d62-600eba932325.jpg	2022-10-27 23:19:20.758383+05	2022-11-09 15:09:41.955371+05	\N
91b6b528-a17a-40fb-ac13-bd5069a752c0	360ebeac-853e-45a5-ab7f-838430b0c442	uploads/product/497900ad-38aa-4838-840d-f3e8dca35f04.jpg	uploads/product/fc9f1610-cbc2-434a-8c48-655d5b96f40a.jpg	2022-10-27 23:19:20.758383+05	2022-11-09 15:09:41.955371+05	\N
62fad7e9-b766-42e5-a111-722c0da6caab	fe309360-c5dd-406a-9957-3d898ea85dfc	uploads/product/9df3e247-c08b-4fce-9a93-1484395d3a75.jpg	uploads/product/10db54bb-033a-4219-8be4-e8af50a48fa3.jpg	2022-10-27 23:20:14.814499+05	2022-11-09 15:09:41.955371+05	\N
79d5fba6-b793-4297-a8d0-94b6a4fc5b7b	fe309360-c5dd-406a-9957-3d898ea85dfc	uploads/product/c0b65dac-0f34-44ff-98bd-795e78e549b2.jpg	uploads/product/8a729026-030e-4d97-aca1-b8393224dc04.jpg	2022-10-27 23:20:14.814499+05	2022-11-09 15:09:41.955371+05	\N
4c819434-14df-4f48-b675-b2becf2609ad	febf699d-ca37-458a-b121-b5b70bbc7db0	uploads/product/2820441c-6ab8-429d-8c39-041c8844df7f.jpg	uploads/product/9a0f4f52-a334-4ed9-9306-b3b4d43ff2df.jpg	2022-10-27 23:21:07.460715+05	2022-11-09 15:09:41.955371+05	\N
d38c521d-5a2a-430f-96c8-51663e9a63d8	febf699d-ca37-458a-b121-b5b70bbc7db0	uploads/product/b81acb7c-cb9c-4464-a2c9-ae1fc06f4a56.jpg	uploads/product/b9b20fea-2ef0-4041-8c39-822ead59af67.jpg	2022-10-27 23:21:07.460715+05	2022-11-09 15:09:41.955371+05	\N
80c94ca2-f069-4183-b0e0-4a473e67b7a9	802b422b-710a-420b-860e-59b7f49d10bd	uploads/product/47827170-3c02-43ab-a6ad-2a29f8fd613e.jpg	uploads/product/4dce1c2f-f552-47b6-838d-cbbf86d054c0.jpg	2022-10-27 23:22:15.697191+05	2022-11-09 15:09:41.955371+05	\N
a1c0f7db-4bb4-45cd-9c71-f50d016d7f51	802b422b-710a-420b-860e-59b7f49d10bd	uploads/product/9741a4cc-8ff4-4915-8460-7c53163f0516.jpg	uploads/product/bca4c0eb-a73c-4afd-808d-90b43547c821.jpg	2022-10-27 23:22:15.697191+05	2022-11-09 15:09:41.955371+05	\N
bccfca4d-428c-4c6e-9dde-449334d685ef	7078e107-dd52-4da1-8007-29ed7cf731fb	uploads/product/844b4e80-9437-4e4b-9ec6-b45a9318ccfe.jpg	uploads/product/f80dd79d-68b7-4de4-8ea8-c3c835105072.jpg	2022-10-28 00:10:28.350298+05	2022-11-09 15:11:14.496257+05	\N
8cd61e93-a3db-4400-b32c-31803ddc9085	7078e107-dd52-4da1-8007-29ed7cf731fb	uploads/product/7120b944-278f-4098-a4c1-e6a0733a0588.jpg	uploads/product/0bbfd269-7cb6-43b5-a960-d0f3b9867871.jpg	2022-10-28 00:10:28.350298+05	2022-11-09 15:11:14.496257+05	\N
f0e654b1-f13e-467e-bcda-c9b29984b2c6	5f6aba1c-66df-4791-b85e-b0a90ccffc20	uploads/product/a06f5df4-c9cc-476f-b84b-29f83c94fa1b.jpg	uploads/product/6abd71ff-ffdf-4495-8128-1b68fe8d1db0.jpg	2022-10-28 00:11:26.340855+05	2022-11-09 15:11:14.496257+05	\N
a808e6fa-600f-485d-add0-8bf422ce7082	5f6aba1c-66df-4791-b85e-b0a90ccffc20	uploads/product/76dad0fb-d666-4c6e-976a-a503d2ae45d7.jpg	uploads/product/70fc8e33-ddd8-4874-8a15-0db290bbd721.jpg	2022-10-28 00:11:26.340855+05	2022-11-09 15:11:14.496257+05	\N
9f1428aa-7867-4922-93cf-ee0aaab53843	01dc8537-7ec1-4c48-bcce-3734f1ac598a	uploads/product/037a94b6-27fe-412b-b3b3-ab105fa46108.jpg	uploads/product/2a8a96ea-f606-408e-b063-9816bccd4569.jpg	2022-10-28 00:11:58.351962+05	2022-11-09 15:11:14.496257+05	\N
1d08a72f-647c-4d22-9cad-ac51ace2ca9f	01dc8537-7ec1-4c48-bcce-3734f1ac598a	uploads/product/5fe57de2-d775-4a9f-af4d-6560543deff6.jpg	uploads/product/0c2aeb3f-fcaf-4542-9e45-48fa649ebf02.jpg	2022-10-28 00:11:58.351962+05	2022-11-09 15:11:14.496257+05	\N
87e6fe7f-4832-4510-ae19-46ef4d9ba20b	214befd8-68bd-484a-a8d5-8e2d0b73931c	uploads/product/facbfc3f-0727-4594-94c2-5d1e2d5dc9bb.jpg	uploads/product/03efc6ab-7b8a-4d77-8a2b-36549e47ad11.jpg	2022-10-28 00:12:34.796369+05	2022-11-09 15:11:14.496257+05	\N
f7d5ef49-fb07-4949-a644-a6caba59c550	214befd8-68bd-484a-a8d5-8e2d0b73931c	uploads/product/74af4f1a-c1d1-4eab-b90d-d889b6376868.jpg	uploads/product/599c3436-2bd9-42f1-ad3d-e69b1b4b422e.jpg	2022-10-28 00:12:34.796369+05	2022-11-09 15:11:14.496257+05	\N
7caf7c30-a9bc-4a73-9a07-e39e37384403	a316070c-409c-42c2-85df-bb7509a24c54	uploads/product/e3d76095-10d7-4d57-a0c8-aacf2e2cc7de.jpg	uploads/product/e8d7ba92-239a-4fa4-8715-2e39b698ad5f.jpg	2022-10-28 00:13:05.554456+05	2022-11-09 15:11:14.496257+05	\N
272b0167-2a50-4877-99d1-5576a5f0f973	a316070c-409c-42c2-85df-bb7509a24c54	uploads/product/e643fd62-e07a-459b-8bab-6d2b103b5fee.jpg	uploads/product/1a5d8d0b-a8e8-46c9-9f9b-c96b3586b465.jpg	2022-10-28 00:13:05.554456+05	2022-11-09 15:11:14.496257+05	\N
229779ac-eb8a-4bae-b5e8-187e38dc2343	89172d2f-b5b3-4b26-a299-dc7e8a71d16e	uploads/product/eb2e2d0b-0f3d-47e5-8f3b-a57b498974dd.jpg	uploads/product/acea27f2-f517-42c5-8d37-2dd427ababb6.jpg	2022-10-28 00:13:38.968059+05	2022-11-09 15:11:14.496257+05	\N
d23332be-54a7-4487-8448-6d19f95a58d4	89172d2f-b5b3-4b26-a299-dc7e8a71d16e	uploads/product/df3b989b-6f45-4a03-ab3a-63535665d7b3.jpg	uploads/product/c502a2c7-eeb0-4ff1-b997-90f63a74a0cc.jpg	2022-10-28 00:13:38.968059+05	2022-11-09 15:11:14.496257+05	\N
6600ce61-3124-4f14-b519-95f5b23ed844	c82fef0a-ad15-4b07-8855-910fc4708af1	uploads/product/401d6961-3148-4f19-9330-b618c3664859.jpg	uploads/product/61f6200f-ba08-4b70-86af-c62dab6ef0b9.jpg	2022-10-28 00:14:34.679499+05	2022-11-09 15:11:14.496257+05	\N
c92b5c54-ba83-4ed0-ac59-f5805d7a968a	c82fef0a-ad15-4b07-8855-910fc4708af1	uploads/product/3059b7e8-515f-4180-98c1-af0c91d98b0f.jpg	uploads/product/853ac031-11c2-4e28-b957-4ff1c6200b42.jpg	2022-10-28 00:14:34.679499+05	2022-11-09 15:11:14.496257+05	\N
4dea0a7a-e0a4-4663-b606-5e2ba02f9860	311aa4c1-6002-4acf-b1b5-e2aa7896def7	uploads/product/890360f1-417d-4782-ba42-2ab58fa314a0.jpg	uploads/product/64131571-2787-4cac-9313-2ceeee12d5f6.jpg	2022-10-28 00:15:19.584828+05	2022-11-09 15:11:14.496257+05	\N
580ab999-7569-4e88-b345-27543cc96e8b	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	uploads/product/0150db4e-e9bf-4cc8-b8f6-7f7f62ef898f.jpg	uploads/product/cf067d6f-ffcf-44ec-a125-758f35d85b37.jpg	2022-10-27 13:24:26.617604+05	2022-11-09 15:11:14.496257+05	\N
865750f3-5d69-47a0-b312-9c2ab0cc2289	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	uploads/product/a724631a-74e7-48ab-b976-7b89f808388d.jpg	uploads/product/4402bb7b-6d14-444d-a18e-d2c34b72a727.jpg	2022-10-27 13:24:26.617604+05	2022-11-09 15:11:14.496257+05	\N
aecb0fb0-5a38-40a1-8af2-d09614a9b6bd	3f397126-6d8d-4a0d-982c-01fd00526957	uploads/product/2ea5faef-2ce3-44b6-80d1-7a7c15701f8c.jpg	uploads/product/9a0b4714-1ee7-4077-be1c-edd2018ade1c.jpg	2022-10-27 13:26:05.297633+05	2022-11-09 15:11:14.496257+05	\N
4eb50f79-8829-4079-b0cd-8331435823d5	3f397126-6d8d-4a0d-982c-01fd00526957	uploads/product/6109cab7-8d18-4e29-8223-710a9b721ca8.jpg	uploads/product/281fe419-0068-4fb2-bdfc-072a82c0c082.jpg	2022-10-27 13:26:05.297633+05	2022-11-09 15:11:14.496257+05	\N
b6c7be30-57c3-4406-963e-a0c74d0e731b	0946a0f5-d23f-4660-9151-80ef91ae9747	uploads/product/583d21a2-5b92-4d30-9062-a7b646e981ca.jpg	uploads/product/4783671a-16a8-4858-b7f6-17ba9f0f8305.jpg	2022-10-27 13:32:30.014698+05	2022-11-09 15:11:14.496257+05	\N
80a99f3b-1586-4305-ab97-0af6b57785bd	0946a0f5-d23f-4660-9151-80ef91ae9747	uploads/product/3600b0d8-27ac-4bc8-adb1-e5826df56cff.jpg	uploads/product/802e22e1-b0a9-4f6b-9456-47e9dc309442.jpg	2022-10-27 13:32:30.014698+05	2022-11-09 15:11:14.496257+05	\N
b93797c8-b747-4e9b-bc2e-6c83bc6b9b8b	03050bc6-6223-49f3-b729-397fd3b6b285	uploads/product/ee8ef999-a666-4c66-a775-61e3acabda1d.jpg	uploads/product/db4aec4d-e284-4eda-9441-77926a80cc35.jpg	2022-10-27 13:36:08.795988+05	2022-11-09 15:11:14.496257+05	\N
615494e5-4be4-47c9-96f8-06dafaa1f898	03050bc6-6223-49f3-b729-397fd3b6b285	uploads/product/085cff96-14ce-4b4b-be05-71ca8590867f.jpg	uploads/product/c547330c-c275-4575-aa7a-543c7811d02c.jpg	2022-10-27 13:36:08.795988+05	2022-11-09 15:11:14.496257+05	\N
95e50227-58e1-438e-ac2e-f1288215cc73	8b481e58-cd39-4761-a052-75e30124689a	uploads/product/51b146be-a1a9-410e-8314-9ed26680d363.jpg	uploads/product/869b8b7c-5643-4596-b8db-468af8a4f312.jpg	2022-10-27 13:38:14.176348+05	2022-11-09 15:11:14.496257+05	\N
c92f6bfa-9139-433c-8177-52370121605e	8b481e58-cd39-4761-a052-75e30124689a	uploads/product/70598020-bd76-4bb8-b021-af48e70bf5d6.jpg	uploads/product/92eb371e-2832-41b8-bdbb-df223fdcd9a9.jpg	2022-10-27 13:38:14.176348+05	2022-11-09 15:11:14.496257+05	\N
3cd664f7-b765-4171-b354-0e944a2fdd13	311aa4c1-6002-4acf-b1b5-e2aa7896def7	uploads/product/1bbf53cd-ca19-43fa-8d1c-d481475efd2d.jpg	uploads/product/e6b19777-9274-4f0d-aed6-8b1f4d3f2b19.jpg	2022-10-28 00:15:19.584828+05	2022-11-09 15:11:14.496257+05	\N
e038a02e-ddda-4e31-a81a-6d92d8dbd75f	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	uploads/product/604b9373-5efa-4530-a7dc-e38e5633ad20.jpg	uploads/product/6d8a9371-b9e3-4e10-9a12-98f5e5614f8e.jpg	2022-10-28 01:24:16.798646+05	2022-11-09 15:09:41.955371+05	\N
45c9eb98-6d8d-4f29-b211-856b5ed5e90b	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	uploads/product/ae8f6c4e-c034-4c75-974c-121631dce46e.jpg	uploads/product/4e29447b-65ab-487c-8e1b-4dea309423b0.jpg	2022-10-28 01:24:16.798646+05	2022-11-09 15:09:41.955371+05	\N
62271c1b-e49e-4cd4-a098-cf23809b0dd1	45a9f186-2521-4eef-a4e0-b5c253c70878	uploads/product/398d721e-d844-4694-bcae-6e0c53563a98.jpg	uploads/product/afe3c3f5-c21f-4086-a917-085116c07705.jpg	2022-10-28 01:24:52.876387+05	2022-11-09 15:09:41.955371+05	\N
e11d840f-dcde-4e3e-91f1-6c29d0d00b0d	45a9f186-2521-4eef-a4e0-b5c253c70878	uploads/product/6bc98fb2-e88e-42b2-bd28-b977bdcd335e.jpg	uploads/product/c824f9b4-7100-47ca-adfb-6db4c3cc721c.jpg	2022-10-28 01:24:52.876387+05	2022-11-09 15:09:41.955371+05	\N
53b6b788-2dd3-4f99-a207-b30889ea21a6	fa148eb2-520f-430e-bd8d-9d5a166d0600	uploads/product/da73ddaa-a306-4c33-89cf-a807174187b3.jpg	uploads/product/3266c7cc-840e-4ea0-985c-d8012533c4c8.jpg	2022-10-28 01:25:37.4003+05	2022-11-09 15:09:41.955371+05	\N
69caa4ae-d0da-4e87-9076-443648e46b98	fa148eb2-520f-430e-bd8d-9d5a166d0600	uploads/product/982e61b2-9fcd-489b-be86-cdc06b483c4b.jpg	uploads/product/0b326dc6-99a7-4ae2-b747-a4da0461573b.jpg	2022-10-28 01:25:37.4003+05	2022-11-09 15:09:41.955371+05	\N
b541c4e4-525a-4fa5-878e-c4789c057284	9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	uploads/product/ab04e2bd-2402-4dbf-882a-9f5f1a7bb5ae.jpg	uploads/product/ce2630d7-f8d1-429c-8340-2485308d4404.jpg	2022-10-28 01:26:46.936658+05	2022-11-09 15:11:14.496257+05	\N
f2bb4825-a7f1-4568-9e24-4ea3a5aa33ee	9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	uploads/product/f1fe6411-785c-457f-959e-b552428c7342.jpg	uploads/product/cf4fd86b-4acc-45df-a3b6-ad624bf917af.jpg	2022-10-28 01:26:46.936658+05	2022-11-09 15:11:14.496257+05	\N
b8089d9f-4685-4eb7-a802-16190993aec1	74ec9e27-de0b-44d9-8036-3fae8be486a9	uploads/product/723a49df-3e86-4382-90f4-cf373f2c8b0d.jpg	uploads/product/bc1ca9ca-b263-463b-a1cd-b767b50affe7.jpg	2022-10-28 01:27:54.381766+05	2022-11-09 15:11:14.496257+05	\N
1bb8cc15-212e-4d8c-8d2d-90bf172366ac	74ec9e27-de0b-44d9-8036-3fae8be486a9	uploads/product/79a9c1fd-6eae-4822-85ad-4e9d22086625.jpg	uploads/product/25d444ab-c899-4cfa-b58e-bf7959b8aa9b.jpg	2022-10-28 01:27:54.381766+05	2022-11-09 15:11:14.496257+05	\N
78e75b7a-49d5-4cf8-92a7-ad8bb0019cd2	332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	uploads/product/e5035dab-c458-496f-a017-648225ddcdaa.jpg	uploads/product/08899946-150e-495c-babe-42de130321fb.jpg	2022-10-28 01:28:28.306962+05	2022-11-09 15:11:14.496257+05	\N
6f745c3a-560c-4959-b889-4335ddaab878	332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	uploads/product/a314f4c5-78c7-4530-b175-4453288c98d6.jpg	uploads/product/697de944-76da-4cec-93c5-c752a9d12748.jpg	2022-10-28 01:28:28.306962+05	2022-11-09 15:11:14.496257+05	\N
077a05d0-52f0-4ad1-91a7-23f8e3e0a9a0	2eb8a13f-3edc-4422-b772-e57bfd8f8797	uploads/product/50321fbf-0ee5-4df5-bc54-c5b086cc24db.jpg	uploads/product/a685b0a3-234b-450a-8046-8a1d37c36c1a.jpg	2022-10-28 01:29:09.030326+05	2022-11-09 15:11:14.496257+05	\N
3aaa588f-e843-4312-8675-445f175cfc1d	2eb8a13f-3edc-4422-b772-e57bfd8f8797	uploads/product/9095b8c6-5732-4080-a776-7a069cf70be7.jpg	uploads/product/28c40eb2-ca18-4f62-a2ee-a2e3b2de16e9.jpg	2022-10-28 01:29:09.030326+05	2022-11-09 15:11:14.496257+05	\N
6453d954-31ac-4f5c-9557-99579cc1a9b7	18f957f2-216d-4810-b4d7-bd4dd49efd0d	uploads/product/3408e90d-a220-4f54-b30d-3b05cc14a617.jpg	uploads/product/6602be43-a50a-4eac-ab9d-903862d9d744.jpg	2022-10-28 01:29:49.5879+05	2022-11-09 15:11:14.496257+05	\N
f438f98c-5f5a-491c-8aab-f7f5ca8f6e52	18f957f2-216d-4810-b4d7-bd4dd49efd0d	uploads/product/2fd806a9-35a8-4570-b2f2-0d8ca6a5b924.jpg	uploads/product/7322e3ce-b7e6-45bb-9e7f-82103f389ff7.jpg	2022-10-28 01:29:49.5879+05	2022-11-09 15:11:14.496257+05	\N
40bebf8a-9010-4a3b-88a2-5c37b4ed225e	ad24153a-997a-46d1-87bb-27aa1e3e8aea	uploads/product/7e044516-87a5-4fd9-aacf-bde265270123.jpg	uploads/product/c1f5b9bc-4598-4584-9256-ae31611f3f5b.jpg	2022-10-28 01:30:32.867315+05	2022-11-09 15:11:14.496257+05	\N
f72f0509-474e-458f-8d29-9dd0e34fe2a1	ad24153a-997a-46d1-87bb-27aa1e3e8aea	uploads/product/d0edaf49-83b5-43d7-b6bc-58074f9741c6.jpg	uploads/product/6a5e0a24-7f7a-4d25-af59-7b7c9be5bc50.jpg	2022-10-28 01:30:32.867315+05	2022-11-09 15:11:14.496257+05	\N
76ac5895-6c11-4383-a187-65e4a195449b	d085e5a4-8229-4177-b5e1-623e80846017	uploads/product/efb6fc81-c9b0-42f7-9ce0-1f414f2c715e.jpg	uploads/product/5fd3bef7-09d0-4b51-a4fb-59a6c8537d6a.jpg	2022-10-27 12:45:23.487097+05	2022-11-09 15:09:41.955371+05	\N
87f3cebb-c979-42db-a4bb-77b74398edb4	d085e5a4-8229-4177-b5e1-623e80846017	uploads/product/fd40bda0-de78-47b5-af32-191063f68243.jpg	uploads/product/8872c8a5-1be6-4cac-9b96-94a579dc3458.jpg	2022-10-27 12:45:23.487097+05	2022-11-09 15:09:41.955371+05	\N
c77af923-69c8-4cfa-8f66-1ef682b5c53c	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	uploads/product/fe37a494-a709-4d84-9eaf-44fba5ed9b0a.jpg	uploads/product/f5a137fa-a069-460b-8b6c-653b663f6b6a.jpg	2022-10-27 12:47:51.157834+05	2022-11-09 15:09:41.955371+05	\N
7dd56932-41ee-45f8-a69b-cf86ded111e1	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	uploads/product/aeec2982-e369-4663-a55a-049811b002e6.jpg	uploads/product/ba67d67d-bf0c-4a0e-adf5-6d54d6a05eb8.jpg	2022-10-27 12:47:51.157834+05	2022-11-09 15:09:41.955371+05	\N
418b0c99-69e3-41f8-bb88-2d3c187e827d	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	uploads/product/d22424be-72d3-44ea-bbfa-0de2e12c19c8.jpg	uploads/product/13c91aca-a93b-4b9a-9af1-4ef2357c6be4.jpg	2022-10-27 12:49:37.905319+05	2022-11-09 15:09:41.955371+05	\N
be8c51d7-ed1a-41ed-b237-30405b801211	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	uploads/product/4721e6b4-43d6-48ae-aec4-e96a2f45e7b7.jpg	uploads/product/21cc543f-79b8-4bfb-8197-0fe6141346ae.jpg	2022-10-27 12:49:37.905319+05	2022-11-09 15:09:41.955371+05	\N
c5e41ca8-f8e3-4383-bf82-3031ee3dfeb1	70a75d8b-d570-41d4-95cb-2199f4417542	uploads/product/05745cf6-6299-4d35-81e0-517930cb843e.jpg	uploads/product/2ed5df2e-1fe9-4e89-94e9-4a57199a2e75.jpg	2022-10-27 13:05:10.321846+05	2022-11-09 15:09:41.955371+05	\N
7f74e7cd-21ee-4400-896e-d93cdc88fb59	70a75d8b-d570-41d4-95cb-2199f4417542	uploads/product/6348f985-5db5-4d9e-b927-ea5b76b975ca.jpg	uploads/product/4d560c92-1fb7-4cb1-8090-447e94966c31.jpg	2022-10-27 13:05:10.321846+05	2022-11-09 15:09:41.955371+05	\N
5889a8ca-d478-40bd-a148-fdeb6b24237a	ee1d67ed-5862-4dfc-8424-52531a240a6c	uploads/product/fb1af5c5-3c6f-480a-83a4-d7a324b27ef1.jpg	uploads/product/4444ce85-b4af-4126-a078-d91970788872.jpg	2022-10-27 13:07:14.091292+05	2022-11-09 15:09:41.955371+05	\N
136f7f16-a271-46da-8577-e756f3aa6847	ee1d67ed-5862-4dfc-8424-52531a240a6c	uploads/product/fae54003-461b-43d7-84aa-c6ae5128c4af.jpg	uploads/product/31d00bf3-7de5-463b-b591-b5775638cc75.jpg	2022-10-27 13:07:14.091292+05	2022-11-09 15:09:41.955371+05	\N
5850c6bd-8c14-4fc0-bc76-ff909cb4f2ff	c14c7f18-77db-4e3c-8939-e6001cb95db0	uploads/product/304771ec-b010-40f3-92df-82b6a5b934de.jpg	uploads/product/2df895c6-4267-4232-82df-53212d785429.jpg	2022-10-27 13:09:04.147542+05	2022-11-09 15:09:41.955371+05	\N
b4cd94cc-0c7f-4f92-8360-89ff7d73fe8f	c14c7f18-77db-4e3c-8939-e6001cb95db0	uploads/product/cbc2c8cd-052a-4d6f-b2ae-ccf785bf31fb.jpg	uploads/product/3b4a1295-860c-4066-8e87-d5515cb81d5c.jpg	2022-10-27 13:09:04.147542+05	2022-11-09 15:09:41.955371+05	\N
8c1c80a9-64cb-4d9a-9986-bb056697eaa3	793be71f-b0fa-43a2-b527-5fb09236f530	uploads/product/f8f1e4f7-be4a-4b6d-a571-c632f36add50.jpg	uploads/product/7b417a83-b629-497d-b757-b03beb029f85.jpg	2022-10-27 13:22:11.393609+05	2022-11-09 15:11:14.496257+05	\N
e0eaf897-245e-421b-af9e-e899c6617e0e	793be71f-b0fa-43a2-b527-5fb09236f530	uploads/product/fa161b81-e4e1-4c22-bfa6-f5f7c7ae1756.jpg	uploads/product/36314b1f-a86d-4bf0-ae1d-f9a80745cdbf.jpg	2022-10-27 13:22:11.393609+05	2022-11-09 15:11:14.496257+05	\N
179de5f5-cab0-4e18-95cf-9ae110b981b4	32055a0a-2d59-45a9-89b0-761d1f6ad047	uploads/product/3758efe5-ea12-469d-a241-5ee078119218.jpg	uploads/product/dfaf2986-a56e-43f7-90bd-f84a194fa7eb.jpg	2022-10-28 01:13:10.606082+05	2022-11-09 15:11:14.496257+05	\N
b16e2d4b-b6bf-429b-b4c5-c02cfff2efa1	32055a0a-2d59-45a9-89b0-761d1f6ad047	uploads/product/fd0cc1d1-5763-4bc0-8af4-5eb73d0ff392.jpg	uploads/product/02a4afef-99ed-46fb-8993-7947bbeb9f3c.jpg	2022-10-28 01:13:10.606082+05	2022-11-09 15:11:14.496257+05	\N
b45eab7d-0771-4c9f-a000-35c9fa458fd8	19e94bc7-398a-45d9-b3cf-5c8d50550e48	uploads/product/701cda76-fb48-4fcc-82bf-ef26401b3dc1.jpg	uploads/product/2f24f533-456d-4492-82d2-626972a4879c.jpg	2022-10-28 01:13:56.372855+05	2022-11-09 15:11:14.496257+05	\N
45f9832b-b035-4dbb-b4d0-583680c9dc20	19e94bc7-398a-45d9-b3cf-5c8d50550e48	uploads/product/d3857efd-18f7-4890-ad0e-cc2413616964.jpg	uploads/product/97d35042-aff7-45d1-885f-d416376c12c4.jpg	2022-10-28 01:13:56.372855+05	2022-11-09 15:11:14.496257+05	\N
8039f565-2e12-4160-bb10-dd6b89879fe9	fd2c148f-70b5-48ba-8cf6-65f26438b46d	uploads/product/1e25a6e5-5f52-4d62-be09-d2dee0202de2.jpg	uploads/product/2453010d-5c5f-46a5-9d1a-d67c09d99f7b.jpg	2022-10-28 01:15:13.838814+05	2022-11-09 15:11:14.496257+05	\N
65b90803-2ac7-453b-af19-b658e603d2b0	fd2c148f-70b5-48ba-8cf6-65f26438b46d	uploads/product/6bea5ddd-1775-44d3-9897-d3bb9140fae0.jpg	uploads/product/5ddf0e79-5173-4393-ae4c-5850b01dc958.jpg	2022-10-28 01:15:13.838814+05	2022-11-09 15:11:14.496257+05	\N
e82719d1-fc5f-4eb6-88c6-30e3d3bc42a6	bcb52dfc-c957-4d5e-9bbc-1fcb607d3fd6	uploads/product/3d9c25b9-5542-43e4-b8b5-afc453938176.jpg	uploads/product/9b53a0ee-31e4-436f-b160-965345129cb2.jpg	2022-10-28 01:15:51.291841+05	2022-11-09 15:11:14.496257+05	\N
27305a59-36e9-465f-8b1a-b22222372128	bcb52dfc-c957-4d5e-9bbc-1fcb607d3fd6	uploads/product/e43ed7de-49a3-40fa-9bf3-18c3f50c9100.jpg	uploads/product/0d48146a-ef45-4269-af24-b84a05097519.jpg	2022-10-28 01:15:51.291841+05	2022-11-09 15:11:14.496257+05	\N
954150ce-308e-4936-b621-c57e70b1b055	badd0869-99df-4df3-8a27-5e27c10a861d	uploads/product/e9af0490-6399-4068-a2d1-2b6e02220598.jpg	uploads/product/1cb45e4c-b843-4159-8585-5750d6e474ea.jpg	2022-10-28 01:22:44.772789+05	2022-11-09 15:11:14.496257+05	\N
47644433-677f-4c40-89c8-cb19ba965fc1	badd0869-99df-4df3-8a27-5e27c10a861d	uploads/product/67eedac3-24bb-4966-aa00-facb4317e2c1.jpg	uploads/product/b649a985-03f8-4412-a899-c2a9758a783b.jpg	2022-10-28 01:22:44.772789+05	2022-11-09 15:11:14.496257+05	\N
4be76e8d-89d5-45ab-b2ff-5f0f04737675	9c655c36-1832-48ca-9f88-c04197f191af	uploads/product/8addee00-1d2e-4d34-af3b-240e16345757.jpg	uploads/product/6a41e8eb-9bcc-421c-96e8-2c02996048f1.jpg	2022-10-28 01:31:06.889648+05	2022-11-09 15:11:14.496257+05	\N
f61c2ad2-7b9a-414d-8ecf-0fa233499af0	9c655c36-1832-48ca-9f88-c04197f191af	uploads/product/9d639611-4a70-4a4b-a2ae-d36a76caa194.jpg	uploads/product/cf6c2ff9-6209-446b-967b-fbc02309e21e.jpg	2022-10-28 01:31:06.889648+05	2022-11-09 15:11:14.496257+05	\N
6c1d1042-f8b2-4404-aa6e-8c8a64532545	ccb43083-1c9e-4e84-bffd-ecb28474165e	uploads/product/42f9cfe5-3c5d-4b69-a8e4-a3d9867a0f5b.jpg	uploads/product/c97ca423-724f-4713-b799-a011287c6aa5.jpg	2022-10-27 13:29:08.23347+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05
987bfb34-af2c-40f1-b540-f3ab8430326d	ccb43083-1c9e-4e84-bffd-ecb28474165e	uploads/product/3127d757-72a9-44d9-9595-b7ed97963f3b.jpg	uploads/product/2275f3ec-9942-41de-9caa-9ef406ecfd00.jpg	2022-10-27 13:29:08.23347+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05
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
d11a3e9a-1d64-4fa9-931e-e649d4600665	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	2022-11-10 08:54:14.379992+05	2022-11-10 08:54:14.379992+05	\N
7de0dbcc-5b39-4ab7-8637-2fb5c092d6f3	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	2022-11-10 08:54:14.401298+05	2022-11-10 08:54:14.401298+05	\N
\.


--
-- Data for Name: main_image; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.main_image (id, product_id, small, medium, large, created_at, updated_at, deleted_at) FROM stdin;
df75d53d-8b81-4e97-838c-67ec42faef1a	214befd8-68bd-484a-a8d5-8e2d0b73931c	uploads/product/ba7c054e-9646-445a-9a5b-94c71c02aed8.jpg	uploads/product/23308225-162c-4dfc-ae81-44fd303a013b.jpg	uploads/product/3a5760b7-2746-4f80-acdd-ca5bd0242ad7.jpg	2022-10-28 00:12:34.781513+05	2022-11-09 15:11:14.496257+05	\N
4c888f3e-923a-40da-8f98-967bf65ff288	a316070c-409c-42c2-85df-bb7509a24c54	uploads/product/e22bcbf1-f5af-4082-8b07-0b9e6fd14ed1.jpg	uploads/product/266c1aa3-aecb-4fb4-9bb7-c55ddc8626a0.jpg	uploads/product/9a8ebf6b-d70e-440c-93ac-0cf51ad45d38.jpg	2022-10-28 00:13:05.537968+05	2022-11-09 15:11:14.496257+05	\N
deea0ded-b341-4c61-8d9d-3d7429997e92	89172d2f-b5b3-4b26-a299-dc7e8a71d16e	uploads/product/56e8ad30-0584-4428-bb96-970bcb09f9af.jpg	uploads/product/1e38780c-6718-417b-adb3-e4d712d92f77.jpg	uploads/product/5a613837-e650-4789-af89-da52b4ca30c1.jpg	2022-10-28 00:13:38.949885+05	2022-11-09 15:11:14.496257+05	\N
54830b38-6bcd-4874-9100-8941800d16f1	c82fef0a-ad15-4b07-8855-910fc4708af1	uploads/product/a839f817-9102-4ead-9946-68f48abf2fb0.jpg	uploads/product/56b9ceaa-05c1-459a-9cb3-a1f677b72328.jpg	uploads/product/8dfc0fcd-74cc-42af-b198-d95a736f271f.jpg	2022-10-28 00:14:34.662577+05	2022-11-09 15:11:14.496257+05	\N
9077315a-0326-4129-8a5a-1c4e29838adf	311aa4c1-6002-4acf-b1b5-e2aa7896def7	uploads/product/bc5e18fc-bb14-4fab-bc87-7657857667db.jpg	uploads/product/aee04cc5-8198-45a5-b99a-4f13eb296d7d.jpg	uploads/product/7242e6f2-1764-47f4-ae96-79f0ab84dd68.jpg	2022-10-28 00:15:19.556846+05	2022-11-09 15:11:14.496257+05	\N
4447e3c9-bfb8-4dc4-8fb3-b7e2c0b81905	32055a0a-2d59-45a9-89b0-761d1f6ad047	uploads/product/1131d5c8-ab7c-4045-a0e5-fa2ea40f1d1c.jpg	uploads/product/082347ae-6923-4110-8c90-23134e92f8f0.jpg	uploads/product/d664c51c-6428-4bd3-9e6e-ddd3db054c0b.jpg	2022-10-28 01:13:10.534589+05	2022-11-09 15:11:14.496257+05	\N
5cc978bf-c6c7-4fc5-b40b-efc037d8cf56	19e94bc7-398a-45d9-b3cf-5c8d50550e48	uploads/product/a10a976a-0f64-4b5a-8636-b1277861e714.jpg	uploads/product/2a0d66c1-9664-491b-89f6-930c5e153fa9.jpg	uploads/product/1fccaa81-d9ef-4bcb-a788-bd26b83761df.jpg	2022-10-28 01:13:56.355921+05	2022-11-09 15:11:14.496257+05	\N
4c1cc717-9a3e-4150-b30e-e3424c97cf24	fd2c148f-70b5-48ba-8cf6-65f26438b46d	uploads/product/5622869a-fa4a-4c71-a0fb-38d0f3067c52.jpg	uploads/product/49a800c2-b41e-4e2f-89ab-3bf74e299930.jpg	uploads/product/fb4ad49b-eb81-4d64-b183-bea7550c4202.jpg	2022-10-28 01:15:13.821376+05	2022-11-09 15:11:14.496257+05	\N
dbacca4a-7c05-4c01-881d-404df4b3ca7d	bcb52dfc-c957-4d5e-9bbc-1fcb607d3fd6	uploads/product/0528fd2d-de7e-48f7-96a6-f6f17852ce3b.jpg	uploads/product/f952de2d-2846-4a4c-9730-c0e5fd0b1a41.jpg	uploads/product/ad90475c-265e-4752-be8b-89eb187ea3cd.jpg	2022-10-28 01:15:51.276734+05	2022-11-09 15:11:14.496257+05	\N
e8727b67-96aa-4884-911d-7e3b555578e6	badd0869-99df-4df3-8a27-5e27c10a861d	uploads/product/a60a6bc9-e1b5-4ab3-8a56-8d60e9fe273a.jpg	uploads/product/1c7947fb-4076-4b74-a84b-8383e4764978.jpg	uploads/product/b46b0942-0c3f-4954-99fb-5aa2088dc85e.jpg	2022-10-28 01:22:44.750879+05	2022-11-09 15:11:14.496257+05	\N
93dac158-087b-4322-a871-f03252381eaf	9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	uploads/product/0d265a9e-5194-4b58-8af3-2804aaec1214.jpg	uploads/product/ef7d9ec3-d2f5-4169-b77e-6843596112a1.jpg	uploads/product/f2f81037-f3da-40b8-9d98-2a18022eb616.jpg	2022-10-28 01:26:46.920307+05	2022-11-09 15:11:14.496257+05	\N
39a0ab39-38ba-431a-a0db-3155172c67d9	74ec9e27-de0b-44d9-8036-3fae8be486a9	uploads/product/fc80734b-373e-4db6-b8d0-bb0617480c21.jpg	uploads/product/554d4bad-e665-48b8-a62d-0cfbe2accada.jpg	uploads/product/b568dafd-d71f-44c1-829f-7f34f2676b21.jpg	2022-10-28 01:27:54.366644+05	2022-11-09 15:11:14.496257+05	\N
da3408ea-1cd6-41b0-aa7d-361ed2325c55	332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	uploads/product/c5366042-fe78-4142-aec1-ed1593d0e4ed.jpg	uploads/product/620544e9-9592-40b1-98f2-2705914e735e.jpg	uploads/product/3628cf7c-b12b-4558-9c03-7f9ff0489ea3.jpg	2022-10-28 01:28:28.256912+05	2022-11-09 15:11:14.496257+05	\N
23dc8d1a-8e83-40ce-80a3-5b86b630fdcc	2eb8a13f-3edc-4422-b772-e57bfd8f8797	uploads/product/71bf32cd-a2ee-4389-ba23-405be37ac8c7.jpg	uploads/product/37c14428-7bbd-4b2c-8bca-d41fc2263afa.jpg	uploads/product/b6824542-c343-4c3a-9be1-ee4f14e01c27.jpg	2022-10-28 01:29:09.014209+05	2022-11-09 15:11:14.496257+05	\N
89b144d9-20eb-42bb-a809-3806bcb0fe3c	18f957f2-216d-4810-b4d7-bd4dd49efd0d	uploads/product/d888ef4f-fec3-4a6a-a18e-2b04aff11532.jpg	uploads/product/90a27ff5-7b1d-4be6-9d94-086c83342bcf.jpg	uploads/product/5d4871fd-84e0-45aa-9124-befd6e078bac.jpg	2022-10-28 01:29:49.572854+05	2022-11-09 15:11:14.496257+05	\N
738556d7-6c7a-4a37-9258-ccac8809e5f7	ad24153a-997a-46d1-87bb-27aa1e3e8aea	uploads/product/9b3307e2-9237-45d7-bea8-664acc26d6f1.jpg	uploads/product/b3ca3781-9d7b-4409-9a4c-6ed1371aac3a.jpg	uploads/product/8dbc5909-6341-4639-83bc-1020da63b4ea.jpg	2022-10-28 01:30:32.817453+05	2022-11-09 15:11:14.496257+05	\N
8784f06e-d7cc-4516-951a-e9539e80ecf7	9c655c36-1832-48ca-9f88-c04197f191af	uploads/product/af808e6d-eab1-44f3-844c-51d59a1d01df.jpg	uploads/product/df743a98-6d36-48c0-ba39-5316b28c55cd.jpg	uploads/product/36639cb1-50bd-468b-8915-5f521884c198.jpg	2022-10-28 01:31:06.807562+05	2022-11-09 15:11:14.496257+05	\N
69bbae4e-e69b-4e46-a0f5-8f6a6574843f	793be71f-b0fa-43a2-b527-5fb09236f530	uploads/product/c61c3627-e507-48ee-8287-e023d20a1339.jpg	uploads/product/5bf97605-7983-4767-84fb-6e11c3ac00cc.jpg	uploads/product/b302a6cb-eb01-4afb-991d-a95f2be25d9b.jpg	2022-10-27 13:22:11.378944+05	2022-11-09 15:11:14.496257+05	\N
c325207e-dfbf-4986-96e6-26fd7e17ccd2	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	uploads/product/d6338cc4-f07e-4068-a9f2-656e64886d62.jpg	uploads/product/ce1ced74-7e0d-4a97-bea7-63dfe2026a75.jpg	uploads/product/06ccab77-b6c9-4356-a182-46ff5de5e8d1.jpg	2022-10-27 13:24:26.603629+05	2022-11-09 15:11:14.496257+05	\N
57233981-9813-4f95-b436-7962f56f2889	3f397126-6d8d-4a0d-982c-01fd00526957	uploads/product/4d7d4295-07e6-4756-af1c-789cf6b8512f.jpg	uploads/product/c766e9c4-787f-41c6-b66c-13adab39c4ac.jpg	uploads/product/ce597557-b988-4a31-a2aa-c22b9f6e960a.jpg	2022-10-27 13:26:05.283184+05	2022-11-09 15:11:14.496257+05	\N
fa340958-a6c9-437d-9928-1f388892fb56	0946a0f5-d23f-4660-9151-80ef91ae9747	uploads/product/b627f961-60f8-40c9-91fa-21818c09cddb.jpg	uploads/product/09db62e8-d9ee-458a-be46-af4988bfe8cb.jpg	uploads/product/37db2574-b67a-4bca-b606-46273f27d63c.jpg	2022-10-27 13:32:30.000631+05	2022-11-09 15:11:14.496257+05	\N
e948a556-f723-4e50-b9d2-2d94b7c7e619	03050bc6-6223-49f3-b729-397fd3b6b285	uploads/product/bb8433b9-3f0b-4f47-8e67-3acad07820a4.jpg	uploads/product/a8edb71f-4a7c-4940-851f-ea281b3d6039.jpg	uploads/product/2e34230d-7ce7-4403-9646-61581a914d65.jpg	2022-10-27 13:36:08.781822+05	2022-11-09 15:11:14.496257+05	\N
8d74ecec-e8c9-4c36-8b90-361c5665a1b7	8b481e58-cd39-4761-a052-75e30124689a	uploads/product/17f209e5-1d6c-437e-b34d-adf5c62c8433.jpg	uploads/product/a998b549-d68a-4e47-8eb0-7cd345a2183a.jpg	uploads/product/96483dec-8e60-4c7c-9b45-f7f28713bac0.jpg	2022-10-27 13:38:14.095266+05	2022-11-09 15:11:14.496257+05	\N
1ecfc54d-0c07-4577-97e1-29503ca60c86	d085e5a4-8229-4177-b5e1-623e80846017	uploads/product/6a2b2732-6585-428e-b6d0-52e94fb9a4f2.jpg	uploads/product/810f948e-6a52-4319-84f8-1cd25b3f8bb7.jpg	uploads/product/6ca77765-43e0-48cb-86ca-573568ab4201.jpg	2022-10-27 12:45:23.472753+05	2022-11-09 15:09:36.259608+05	\N
4de19867-d401-4538-b7ce-015967d0c6f0	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	uploads/product/fd224ec4-c237-4079-811c-f5fa39e3e885.jpg	uploads/product/6c4ae0ab-58bb-4405-9cd4-657605995d1d.jpg	uploads/product/a301fd4e-2e75-4503-8530-8fdcf80b15b1.jpg	2022-10-27 12:47:51.142922+05	2022-11-09 15:09:36.259608+05	\N
b3460c07-2e1f-4110-94c2-14750748eec6	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	uploads/product/f498f90e-89a0-43e9-9116-0216348d10ab.jpg	uploads/product/4d44c079-787a-45a2-8b90-6a7d83d9ef2c.jpg	uploads/product/483dbc69-d63c-4202-9d55-c7f283837ae9.jpg	2022-10-27 12:49:37.845528+05	2022-11-09 15:09:36.259608+05	\N
258471e1-e002-4e3d-9be8-4a442a6d6d50	70a75d8b-d570-41d4-95cb-2199f4417542	uploads/product/1ec30d87-b8a3-4de8-a71f-1749328bf20e.jpg	uploads/product/88cf205f-c937-4f89-a29e-be0cc8b0d99e.jpg	uploads/product/82ad8ac8-251a-4432-9b07-19e3596198fb.jpg	2022-10-27 13:05:10.30663+05	2022-11-09 15:09:36.259608+05	\N
a95f9bba-0dc1-4096-90c3-db2cc308bda7	ee1d67ed-5862-4dfc-8424-52531a240a6c	uploads/product/38457790-da1e-4ffb-bdae-c28c40fc6534.jpg	uploads/product/60eeed69-6577-42d9-a7f8-08000ea7dbc2.jpg	uploads/product/8c6655d0-eb65-4431-81fd-916d35f42850.jpg	2022-10-27 13:07:14.075976+05	2022-11-09 15:09:36.259608+05	\N
bfffc5b6-956d-4460-9683-973289a4a76b	c14c7f18-77db-4e3c-8939-e6001cb95db0	uploads/product/0c009049-b607-44e3-8d25-140040506225.jpg	uploads/product/ed283bde-c753-49c7-8e3a-e40911da1ef6.jpg	uploads/product/48fc7ed3-5e18-4d57-88ee-bcd389377593.jpg	2022-10-27 13:09:04.132921+05	2022-11-09 15:09:36.259608+05	\N
ef3d46dc-e961-472e-ad8c-e4e61a10e7a2	83da5c7b-bffe-4450-97c9-0f376441b1d4	uploads/product/e4dacbb0-3fe1-456d-bd26-23e382a0d536.jpg	uploads/product/05d9a4d6-fddf-4a72-8c0b-650c6940217f.jpg	uploads/product/0e3b5a4d-daef-4cfa-8847-dd1538ac6ce1.jpg	2022-10-27 13:30:49.665842+05	2022-11-09 15:09:36.259608+05	\N
fa50db2f-5ceb-46b6-9941-62800d9d1aab	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	uploads/product/6006c020-ea15-4a58-a80a-3026d0624550.jpg	uploads/product/1d3f3d7a-fbbd-440c-8115-fb107e6c6356.jpg	uploads/product/5303fca2-9314-4bc6-88c8-6262e7b9c045.jpg	2022-10-27 23:15:36.031827+05	2022-11-09 15:09:36.259608+05	\N
1c0b5103-5844-4678-b939-1d853989398b	81462bfa-36df-4e09-aa46-c6fa1ab86de6	uploads/product/f9f8496c-e46e-442c-8812-c665e3b776a4.jpg	uploads/product/c891cceb-618e-45ee-b350-a5e0c73ed169.jpg	uploads/product/d726dc68-f99d-4dbc-ad7d-3d5a876444a3.jpg	2022-10-27 23:17:42.792181+05	2022-11-09 15:11:14.496257+05	\N
0e9f5a75-6930-495d-bf93-c9b1a0fd0b0e	35f5f2d8-9271-469f-bde1-2314c18ea574	uploads/product/36682327-39cd-44f2-834a-809257a2bd7d.jpg	uploads/product/3603e0b8-7d14-40ab-8d80-ab64517a16ea.jpg	uploads/product/e3a9704a-db38-4bbc-9a86-6c79fb884d26.jpg	2022-10-27 23:23:19.416485+05	2022-11-09 15:11:14.496257+05	\N
d40e2562-52e0-490c-b047-2928955dd4a7	360ebeac-853e-45a5-ab7f-838430b0c442	uploads/product/de4b3324-0485-4abd-a568-37c4f6d470b8.jpg	uploads/product/c3fada3a-4421-4413-9ca2-bac4d3911016.jpg	uploads/product/919d1dbf-5348-42a8-a669-9c2bdb025114.jpg	2022-10-27 23:19:20.741606+05	2022-11-09 15:09:36.259608+05	\N
de10aba0-5cdf-4b28-8c9c-c3f20639c5fc	fe309360-c5dd-406a-9957-3d898ea85dfc	uploads/product/c793782c-c880-4490-ae08-34547ad180cd.jpg	uploads/product/44fa28dc-23f5-432c-a1c4-d3d0d2656cb5.jpg	uploads/product/3b5a7411-60c6-4704-a153-4c4b1cbe12b8.jpg	2022-10-27 23:20:14.774758+05	2022-11-09 15:09:36.259608+05	\N
5e9fffdf-06cf-4944-b949-ee5715fc3c05	febf699d-ca37-458a-b121-b5b70bbc7db0	uploads/product/127368ed-e77d-4639-8ebd-aedef27c200c.jpg	uploads/product/25e133af-0198-4a14-a66e-36022db617b1.jpg	uploads/product/73eb6faf-fb5c-4395-9579-a5bba1bc8322.jpg	2022-10-27 23:21:07.444791+05	2022-11-09 15:09:36.259608+05	\N
6dee238b-10ff-455a-ae38-f1f63d55f567	802b422b-710a-420b-860e-59b7f49d10bd	uploads/product/f24f0f5f-6954-462d-9795-1f895a89f707.jpg	uploads/product/623e606d-53cd-45cd-b67c-63bb74507b2a.jpg	uploads/product/9c325d22-1224-4586-a79a-508eef75124a.jpg	2022-10-27 23:22:15.680856+05	2022-11-09 15:09:36.259608+05	\N
9aff278a-615b-440c-81df-5e33634b8ae3	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	uploads/product/2d36f97c-397f-41d8-b503-9c404e0713da.jpg	uploads/product/3a05513c-310f-45ed-82c1-fb8f57b8f475.jpg	uploads/product/f4f607bf-c802-4cb2-8bdb-50ebb3619d77.jpg	2022-10-28 01:24:16.782147+05	2022-11-09 15:09:36.259608+05	\N
e7a54110-ad06-4cf4-bea8-9935e99df168	45a9f186-2521-4eef-a4e0-b5c253c70878	uploads/product/44742d95-3475-4e0f-8846-a45a8781231b.jpg	uploads/product/af09bf0f-f441-4995-a508-400bf1e3d8f4.jpg	uploads/product/84f9c74f-8180-4273-8b77-95c69e95c808.jpg	2022-10-28 01:24:52.861397+05	2022-11-09 15:09:36.259608+05	\N
2f3beb18-0836-465b-bc4f-c851c34aec1e	fa148eb2-520f-430e-bd8d-9d5a166d0600	uploads/product/38efbe3b-80d4-422a-8e02-d5e77c6c5191.jpg	uploads/product/2d09ae6a-ade2-481a-a073-39771a06ebf4.jpg	uploads/product/39774041-20ce-440f-8591-2e8e0a8a7615.jpg	2022-10-28 01:25:37.384738+05	2022-11-09 15:09:36.259608+05	\N
322c9ce5-12a2-43ec-a533-2bbfd9b9dbb8	d987b7ad-257e-4ae2-befb-b7d369252a54	uploads/product/09dce57f-03b6-48c3-932b-dbb520eafac7.jpg	uploads/product/5ad27340-5ad3-4839-a5ef-3d09484afef7.jpg	uploads/product/f30bbffb-3cbc-4c2f-b847-da4a5c5752ae.jpg	2022-10-27 23:53:01.79677+05	2022-11-09 15:11:14.496257+05	\N
1809a348-4705-4511-9485-5e7e72058159	77ecf422-b48b-45fd-8e58-380e23d74c4c	uploads/product/81e7eb9d-f900-4206-9e53-9b9947f0aa8f.jpg	uploads/product/aae92bed-0e94-4536-b79b-9266648ae1ac.jpg	uploads/product/d3ce4211-5aa2-4ef5-b18e-37fe2793f345.jpg	2022-10-28 00:08:25.784901+05	2022-11-09 15:11:14.496257+05	\N
fa5d4ac7-e082-4589-9af2-18fc434b3d63	9b9ef1ce-2f3d-4051-8e88-5c301bd68554	uploads/product/69eb5492-4037-475b-9469-335ebc33b03d.jpg	uploads/product/59a9015f-446f-459f-93e6-d4f5b4cc4da3.jpg	uploads/product/693f8675-5166-4d12-8e5c-73e63e8956ed.jpg	2022-10-28 00:09:12.464121+05	2022-11-09 15:11:14.496257+05	\N
3fea5404-4a4f-4e0c-8428-53a4c72c7612	20a8c487-9a56-4fb6-8fa0-13facaf96109	uploads/product/2ec8d062-34ac-466b-89f4-b5e63e63de31.jpg	uploads/product/e2dc63e5-55da-429f-ad78-cc911941776e.jpg	uploads/product/6e4f5f79-7664-4483-87a3-7a7f699b56a8.jpg	2022-10-28 00:09:53.565536+05	2022-11-09 15:11:14.496257+05	\N
15dab5fb-955b-4b97-b27c-bf1c2c582358	7078e107-dd52-4da1-8007-29ed7cf731fb	uploads/product/8ada7e6a-e8f4-4164-a1e6-2c32ef6ba3eb.jpg	uploads/product/83e05a0b-5c84-4a22-ac61-1be1d65cb293.jpg	uploads/product/2a17b554-6652-4dec-bfdd-09ffb2550627.jpg	2022-10-28 00:10:28.333242+05	2022-11-09 15:11:14.496257+05	\N
8c9d3a09-89f9-4830-b896-4f39b9a006b9	5f6aba1c-66df-4791-b85e-b0a90ccffc20	uploads/product/d5b2e5ac-6e29-49cf-bdbe-c8ac48927fc0.jpg	uploads/product/7d96bc51-3d16-48b0-b205-03a567da4142.jpg	uploads/product/2978b50f-ffe7-492a-9683-735e430e9182.jpg	2022-10-28 00:11:26.323714+05	2022-11-09 15:11:14.496257+05	\N
ecd99db9-df34-4cad-b8a3-0b030e4e188d	01dc8537-7ec1-4c48-bcce-3734f1ac598a	uploads/product/f535c1ed-f6b8-41d7-9d44-9b3d602bcd05.jpg	uploads/product/f0930527-7aa3-4aa9-82c7-630847b894cf.jpg	uploads/product/74ce7c46-f94b-4498-9ff4-c4be74cd401a.jpg	2022-10-28 00:11:58.336604+05	2022-11-09 15:11:14.496257+05	\N
45456031-0230-493e-b897-bf218a376fbf	ccb43083-1c9e-4e84-bffd-ecb28474165e	uploads/product/e122a649-5132-45f5-99a7-ef299f683bcf.jpg	uploads/product/25d480da-32bc-42a2-a1ab-0780d1207171.jpg	uploads/product/63ad2aaf-0c61-4380-9b88-113b6b1eac5d.jpg	2022-10-27 13:29:08.219368+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05
\.


--
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.notifications (id, name, created_at, updated_at, deleted_at) FROM stdin;
f832e5da-d969-43d7-9cd0-7eae6c6c59e9	sargyt_ucin	2022-11-08 23:08:35.4564+05	2022-11-08 23:08:35.4564+05	\N
\.


--
-- Data for Name: order_dates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_dates (id, date, created_at, updated_at, deleted_at) FROM stdin;
c1f2beca-a6b6-4971-a6a7-ed50079c6912	tomorrow	2022-09-28 17:36:46.804343+05	2022-09-28 17:36:46.804343+05	\N
32646376-c93f-412b-9e75-b3a5fa70df9e	today	2022-09-28 17:35:33.772335+05	2022-11-01 23:58:35.636948+05	\N
\.


--
-- Data for Name: order_times; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.order_times (id, order_date_id, "time", created_at, updated_at, deleted_at) FROM stdin;
7d47a77a-b8f3-4e96-aa56-5ec7fb328e86	c1f2beca-a6b6-4971-a6a7-ed50079c6912	09:00 - 12:00	2022-09-28 17:36:46.825964+05	2022-09-28 17:36:46.825964+05	\N
de31361b-9fba-48f2-9341-9e3dd08cf9fd	c1f2beca-a6b6-4971-a6a7-ed50079c6912	18:00 - 21:00	2022-09-28 17:36:46.825964+05	2022-09-28 17:36:46.825964+05	\N
67c488ef-6021-4cc5-96cc-25408e71dbe3	32646376-c93f-412b-9e75-b3a5fa70df9e	09:00 - 12:00	2022-11-01 22:51:27.754592+05	2022-11-01 23:58:35.636948+05	\N
861ae017-67b5-45ae-88d7-a6990d7c49fd	32646376-c93f-412b-9e75-b3a5fa70df9e	18:00 - 21:00	2022-11-01 22:51:27.767647+05	2022-11-01 23:58:35.636948+05	\N
\.


--
-- Data for Name: ordered_products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ordered_products (id, product_id, quantity_of_product, order_id, created_at, updated_at, deleted_at) FROM stdin;
9074c15c-0f85-4952-a267-2d791326e7e4	03050bc6-6223-49f3-b729-397fd3b6b285	10	668a5cbf-f1ce-4804-a811-b9fd25fd7c10	2022-11-19 15:41:21.066704+05	2022-11-19 15:41:21.066704+05	\N
81c5b30a-b8bc-4823-b220-2d069cdb9a12	214befd8-68bd-484a-a8d5-8e2d0b73931c	4	668a5cbf-f1ce-4804-a811-b9fd25fd7c10	2022-11-19 15:41:21.247433+05	2022-11-19 15:41:21.247433+05	\N
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, customer_id, customer_mark, order_time, payment_type, total_price, created_at, updated_at, deleted_at, order_number, shipping_price, excel, address) FROM stdin;
668a5cbf-f1ce-4804-a811-b9fd25fd7c10	1ae12390-03ae-49ac-a9ad-d7ba5c95b51a	isleg market bet cykypdyr	12:00 - 16:00	nagt	934.8	2022-11-19 15:41:20.784455+05	2022-11-19 15:41:21.402377+05	\N	49	10	uploads/orders/49.xlsx	Mir 2/2 jay 7 oy 36
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

COPY public.products (id, brend_id, price, old_price, amount, created_at, updated_at, deleted_at, limit_amount, is_new, shop_id) FROM stdin;
793be71f-b0fa-43a2-b527-5fb09236f530	fdd259c2-794a-42b9-a3ad-9e91502af23e	72.5	0	2	2022-10-27 13:22:11.35263+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
0946a0f5-d23f-4660-9151-80ef91ae9747	214be879-65c3-4710-86b4-3fc3bce2e974	141.2	0	2	2022-10-27 13:32:29.977781+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
8b481e58-cd39-4761-a052-75e30124689a	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	161	0	2	2022-10-27 13:38:14.068965+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
81462bfa-36df-4e09-aa46-c6fa1ab86de6	214be879-65c3-4710-86b4-3fc3bce2e974	67	0	2	2022-10-27 23:17:42.767975+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
35f5f2d8-9271-469f-bde1-2314c18ea574	c4bcda34-7332-4ae5-8129-d7538d63fee4	1	0	2	2022-10-27 23:23:19.34493+05	2022-11-09 15:11:14.496257+05	\N	5	t	a283d9a4-f38e-43ee-a228-6584b7406cc4
d987b7ad-257e-4ae2-befb-b7d369252a54	214be879-65c3-4710-86b4-3fc3bce2e974	10.8	0	2	2022-10-27 23:53:01.778407+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
9b9ef1ce-2f3d-4051-8e88-5c301bd68554	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	31.2	0	2	2022-10-28 00:09:12.441851+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
20a8c487-9a56-4fb6-8fa0-13facaf96109	fdd259c2-794a-42b9-a3ad-9e91502af23e	22.8	0	2	2022-10-28 00:09:53.537927+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
7078e107-dd52-4da1-8007-29ed7cf731fb	fdd259c2-794a-42b9-a3ad-9e91502af23e	3.7	0	2	2022-10-28 00:10:28.312591+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
5f6aba1c-66df-4791-b85e-b0a90ccffc20	46b13f0a-d584-4ad3-b270-437ecdc51449	15	0	2	2022-10-28 00:11:26.306096+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
01dc8537-7ec1-4c48-bcce-3734f1ac598a	46b13f0a-d584-4ad3-b270-437ecdc51449	23	0	2	2022-10-28 00:11:58.30745+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
fd2c148f-70b5-48ba-8cf6-65f26438b46d	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	8.4	0	2	2022-10-28 01:15:13.800473+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
bcb52dfc-c957-4d5e-9bbc-1fcb607d3fd6	fdd259c2-794a-42b9-a3ad-9e91502af23e	3.6	0	2	2022-10-28 01:15:51.252579+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
badd0869-99df-4df3-8a27-5e27c10a861d	fdd259c2-794a-42b9-a3ad-9e91502af23e	165.2	0	2	2022-10-28 01:22:44.491613+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	46b13f0a-d584-4ad3-b270-437ecdc51449	146.9	0	2	2022-10-28 01:26:46.876844+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
70a75d8b-d570-41d4-95cb-2199f4417542	46b13f0a-d584-4ad3-b270-437ecdc51449	74.8	94.4	2	2022-10-27 13:05:10.286344+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
ee1d67ed-5862-4dfc-8424-52531a240a6c	f53a27b4-7810-4d8f-bd45-edad405d92b9	74.8	99.3	2	2022-10-27 13:07:14.050311+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
c14c7f18-77db-4e3c-8939-e6001cb95db0	f53a27b4-7810-4d8f-bd45-edad405d92b9	68.9	74.8	2	2022-10-27 13:09:04.112256+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	25.9	37.5	2	2022-10-27 23:15:35.997022+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
360ebeac-853e-45a5-ab7f-838430b0c442	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	52.4	66	2	2022-10-27 23:19:20.710619+05	2022-11-09 15:09:28.523722+05	\N	5	t	74cce5dc-6fc2-487c-8553-1f00850df257
fe309360-c5dd-406a-9957-3d898ea85dfc	fdd259c2-794a-42b9-a3ad-9e91502af23e	37.9	47	2	2022-10-27 23:20:14.720704+05	2022-11-09 15:09:28.523722+05	\N	5	t	74cce5dc-6fc2-487c-8553-1f00850df257
febf699d-ca37-458a-b121-b5b70bbc7db0	fdd259c2-794a-42b9-a3ad-9e91502af23e	73.6	92.7	2	2022-10-27 23:21:07.422604+05	2022-11-09 15:09:28.523722+05	\N	5	t	74cce5dc-6fc2-487c-8553-1f00850df257
802b422b-710a-420b-860e-59b7f49d10bd	46b13f0a-d584-4ad3-b270-437ecdc51449	55.2	69.6	2	2022-10-27 23:22:15.654712+05	2022-11-09 15:09:28.523722+05	\N	5	t	74cce5dc-6fc2-487c-8553-1f00850df257
d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	fdd259c2-794a-42b9-a3ad-9e91502af23e	79.5	88.3	2	2022-10-28 01:24:16.759592+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
45a9f186-2521-4eef-a4e0-b5c253c70878	fdd259c2-794a-42b9-a3ad-9e91502af23e	79.5	88.3	2	2022-10-28 01:24:52.841332+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
fa148eb2-520f-430e-bd8d-9d5a166d0600	fdd259c2-794a-42b9-a3ad-9e91502af23e	49.3	54.8	2	2022-10-28 01:25:37.342021+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	c4bcda34-7332-4ae5-8129-d7538d63fee4	86.3	158	2	2022-10-27 12:49:37.648623+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
d085e5a4-8229-4177-b5e1-623e80846017	c4bcda34-7332-4ae5-8129-d7538d63fee4	52.4	61.6	2	2022-10-27 12:45:23.437479+05	2022-11-09 15:09:28.523722+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
a316070c-409c-42c2-85df-bb7509a24c54	f53a27b4-7810-4d8f-bd45-edad405d92b9	2.8	0	2	2022-10-28 00:13:05.512248+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
89172d2f-b5b3-4b26-a299-dc7e8a71d16e	f53a27b4-7810-4d8f-bd45-edad405d92b9	12.7	0	2	2022-10-28 00:13:38.928331+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
c82fef0a-ad15-4b07-8855-910fc4708af1	c4bcda34-7332-4ae5-8129-d7538d63fee4	4	0	2	2022-10-28 00:14:34.637038+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
311aa4c1-6002-4acf-b1b5-e2aa7896def7	c4bcda34-7332-4ae5-8129-d7538d63fee4	3.6	0	2	2022-10-28 00:15:19.537659+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
32055a0a-2d59-45a9-89b0-761d1f6ad047	c4bcda34-7332-4ae5-8129-d7538d63fee4	5.7	0	2	2022-10-28 01:13:10.464164+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
19e94bc7-398a-45d9-b3cf-5c8d50550e48	214be879-65c3-4710-86b4-3fc3bce2e974	4	0	2	2022-10-28 01:13:56.310076+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
74ec9e27-de0b-44d9-8036-3fae8be486a9	46b13f0a-d584-4ad3-b270-437ecdc51449	17.3	0	2	2022-10-28 01:27:54.313788+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	46b13f0a-d584-4ad3-b270-437ecdc51449	17.3	0	2	2022-10-28 01:28:28.098708+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
2eb8a13f-3edc-4422-b772-e57bfd8f8797	c4bcda34-7332-4ae5-8129-d7538d63fee4	18.7	0	2	2022-10-28 01:29:08.991169+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
18f957f2-216d-4810-b4d7-bd4dd49efd0d	c4bcda34-7332-4ae5-8129-d7538d63fee4	42.4	0	2	2022-10-28 01:29:49.510321+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
ad24153a-997a-46d1-87bb-27aa1e3e8aea	c4bcda34-7332-4ae5-8129-d7538d63fee4	32.7	0	2	2022-10-28 01:30:32.769134+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
9c655c36-1832-48ca-9f88-c04197f191af	c4bcda34-7332-4ae5-8129-d7538d63fee4	30.5	0	2	2022-10-28 01:31:06.759114+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
e34a20fa-3aef-4ba6-92ba-79d3649c61a6	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	109	0	2	2022-10-27 13:24:26.583694+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
77ecf422-b48b-45fd-8e58-380e23d74c4c	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	31.8	0	2	2022-10-28 00:08:25.761854+05	2022-11-09 15:11:14.496257+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	c4bcda34-7332-4ae5-8129-d7538d63fee4	82.7	91.8	10	2022-10-27 12:47:51.121005+05	2022-11-19 15:52:19.01745+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
03050bc6-6223-49f3-b729-397fd3b6b285	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	115	0	10	2022-10-27 13:36:08.764157+05	2022-11-19 15:52:19.01745+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
214befd8-68bd-484a-a8d5-8e2d0b73931c	f53a27b4-7810-4d8f-bd45-edad405d92b9	4.2	0	10	2022-10-28 00:12:34.7569+05	2022-11-19 15:52:19.01745+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
ccb43083-1c9e-4e84-bffd-ecb28474165e	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	92	96.9	2	2022-10-27 13:29:08.198279+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
3f397126-6d8d-4a0d-982c-01fd00526957	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	95	0	2	2022-10-27 13:26:05.260347+05	2022-11-22 09:57:08.025573+05	\N	5	f	a283d9a4-f38e-43ee-a228-6584b7406cc4
83da5c7b-bffe-4450-97c9-0f376441b1d4	ddccb2dc-9697-4f4e-acf5-26b8bc2c8b72	50.1	85.8	1	2022-10-27 13:30:49.646484+05	2022-11-22 09:57:30.851479+05	\N	5	f	74cce5dc-6fc2-487c-8553-1f00850df257
\.


--
-- Data for Name: shops; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shops (id, owner_name, address, phone_number, running_time, created_at, updated_at, deleted_at) FROM stdin;
74cce5dc-6fc2-487c-8553-1f00850df257	Owez Myradow	Asgabat saher Mir 4/1 jay 2 magazyn 56	+99361254689	7:00-21:00	2022-11-09 14:02:42.724172+05	2022-11-09 15:09:48.522999+05	\N
a283d9a4-f38e-43ee-a228-6584b7406cc4	Arslan Kerimow	Asgabat saher Mir 2/2 jay 2 magazyn 23	+99362420387	8:00-22:00	2022-11-09 13:53:47.243923+05	2022-11-09 15:11:14.496257+05	\N
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
9dc67a0a-f78a-4094-9b40-faafd55e87b1	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	Gök we bakja önümleri	2022-10-27 23:49:57.648716+05	2022-10-27 23:49:57.648716+05	\N
906373de-0e0b-4f74-b6d8-5d20d7c9a53d	aea98b93-7bdf-455b-9ad4-a259d69dc76e	44d2783f-133e-4bb7-b4c2-9e03dc04e2dd	Фрукты и овощи	2022-10-27 23:49:57.66415+05	2022-10-27 23:49:57.66415+05	\N
ed1bfdf4-e479-4c1c-8bd4-5d2e8459d05b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	f47ad001-2fbf-49bd-948d-e5c7fa373712	Miweler	2022-10-27 23:50:35.915849+05	2022-10-27 23:50:35.915849+05	\N
ef104ee3-9004-40dd-b7d7-4220c65e5783	aea98b93-7bdf-455b-9ad4-a259d69dc76e	f47ad001-2fbf-49bd-948d-e5c7fa373712	Фрукты	2022-10-27 23:50:35.987027+05	2022-10-27 23:50:35.987027+05	\N
42dc017a-ec97-46d8-aa3e-1b8caeb3c920	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	723bd96d-ea4e-44d3-8052-a1579f32b216	Gök otlar	2022-10-27 23:50:54.9831+05	2022-10-27 23:50:54.9831+05	\N
2f7f0906-7d14-4bbf-9d54-fe6bb53ff155	aea98b93-7bdf-455b-9ad4-a259d69dc76e	723bd96d-ea4e-44d3-8052-a1579f32b216	Зелень\n	2022-10-27 23:50:54.997372+05	2022-10-27 23:50:54.997372+05	\N
a2a49708-3863-46fd-a364-f273893c48b9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	533f8773-0034-42b0-9269-33bc73ae9cd2	Gök önümler	2022-10-27 23:51:19.516869+05	2022-10-27 23:51:19.516869+05	\N
68cb6490-4a07-4242-b42f-aee6c72c0266	aea98b93-7bdf-455b-9ad4-a259d69dc76e	533f8773-0034-42b0-9269-33bc73ae9cd2	Овощи	2022-10-27 23:51:19.53169+05	2022-10-27 23:51:19.53169+05	\N
09388897-a1b0-4a89-abf0-68a9528f7e61	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	789cbced-9141-4748-94d3-93476d276057	Şahsy ideg, kosmetika	2022-10-28 01:19:00.465181+05	2022-10-28 01:19:00.465181+05	\N
00cb656d-3639-4895-a71e-f44f93f0bd89	aea98b93-7bdf-455b-9ad4-a259d69dc76e	789cbced-9141-4748-94d3-93476d276057	Косметика и уход	2022-10-28 01:19:00.493038+05	2022-10-28 01:19:00.493038+05	\N
9e665bec-e985-49cc-92f5-8b9312f49670	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	Kosmetika	2022-10-28 01:19:43.143664+05	2022-10-28 01:19:43.143664+05	\N
cb61081e-e86e-41d8-bece-7be188b25a85	aea98b93-7bdf-455b-9ad4-a259d69dc76e	28a5bd8a-318a-4acf-b3c9-8ba04be5a979	Косметика	2022-10-28 01:19:43.15858+05	2022-10-28 01:19:43.15858+05	\N
358b416e-ddfc-422e-b382-e379809e4854	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7e3eeef8-4748-483c-bbf8-3767943135ee	Diş saglygy we arassaçylygy	2022-10-28 01:20:05.288219+05	2022-10-28 01:20:05.288219+05	\N
142572af-4ed1-4036-b9d6-85b7e3eeeceb	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7e3eeef8-4748-483c-bbf8-3767943135ee	Здоровье и чистота зубов	2022-10-28 01:20:05.302795+05	2022-10-28 01:20:05.302795+05	\N
538296a0-e2cf-4b7e-a4fb-164b7faf4eda	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	5e16c816-a24a-42a4-92a8-8f765e72a149	Makiaž	2022-10-28 01:20:44.388622+05	2022-10-28 01:20:44.388622+05	\N
07517ceb-d7a1-481e-88a9-39cb10eda852	aea98b93-7bdf-455b-9ad4-a259d69dc76e	5e16c816-a24a-42a4-92a8-8f765e72a149	Макияж	2022-10-28 01:20:44.404066+05	2022-10-28 01:20:44.404066+05	\N
b4c2a314-4687-47e0-8a5a-cad0b922b191	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	57d072c8-4952-44c5-845e-d2d706677e16	Diş pastasy	2022-10-28 01:21:13.055793+05	2022-10-28 01:21:13.055793+05	\N
426ae99c-25c0-4de1-81aa-a8122d081e4b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	57d072c8-4952-44c5-845e-d2d706677e16	Зубная паста	2022-10-28 01:21:13.075148+05	2022-10-28 01:21:13.075148+05	\N
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
-- Data for Name: translation_notification; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_notification (id, notification_id, lang_id, translation, created_at, updated_at, deleted_at) FROM stdin;
bb82aa0f-dd88-49cd-9f6a-bd27b2930505	f832e5da-d969-43d7-9cd0-7eae6c6c59e9	aea98b93-7bdf-455b-9ad4-a259d69dc76e	Ваш заказ успешно получен	2022-11-08 23:08:35.531894+05	2022-11-08 23:08:35.531894+05	\N
b0e087dd-60c8-4f47-8755-95e9678d4405	f832e5da-d969-43d7-9cd0-7eae6c6c59e9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	siziň sargydyňyz üstünlikli kabul edildi	2022-11-08 23:08:35.616593+05	2022-11-08 23:08:35.616593+05	\N
\.


--
-- Data for Name: translation_order_dates; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.translation_order_dates (id, lang_id, order_date_id, date, created_at, updated_at, deleted_at) FROM stdin;
1aa5185f-9815-4e3f-9c34-718bfb587d91	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c1f2beca-a6b6-4971-a6a7-ed50079c6912	Ertir	2022-09-28 17:36:46.836838+05	2022-09-28 17:36:46.836838+05	\N
9e7a3752-fce2-4b66-bf3e-d915bf463f92	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c1f2beca-a6b6-4971-a6a7-ed50079c6912	Завтра	2022-09-28 17:36:46.847888+05	2022-09-28 17:36:46.847888+05	\N
3338d831-f091-4574-a0bf-f9cb07dd4893	aea98b93-7bdf-455b-9ad4-a259d69dc76e	32646376-c93f-412b-9e75-b3a5fa70df9e	Segodnya	2022-09-28 17:35:33.82453+05	2022-11-01 23:58:35.636948+05	\N
dcd0c70b-9fa2-4327-8b35-de29bd3febcb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	32646376-c93f-412b-9e75-b3a5fa70df9e	Su gun	2022-09-28 17:35:33.812812+05	2022-11-01 23:58:35.636948+05	\N
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
8f518241-a80c-467d-bc40-c1c0f3ab8585	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ee1d67ed-5862-4dfc-8424-52531a240a6c	Kofe Idee Kaffee "100% Arabica" 250 gr	Kofe Idee Kaffee "100% Arabica" 250 gr	2022-10-27 13:07:14.101052+05	2022-11-09 15:09:20.835667+05	\N	kofe-idee-kaffee-100-arabica-250-gr
7f4a55fd-9237-4af0-ae60-2fc5191cca08	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ee1d67ed-5862-4dfc-8424-52531a240a6c	Кофе Idee Kaffee "100% Arabica" молотый 250 г	Кофе Idee Kaffee "100% Arabica" молотый 250 г	2022-10-27 13:07:14.112977+05	2022-11-09 15:09:20.835667+05	\N	kofe-idee-kaffee-100-arabica-molotyi-250-g
8bf6e83c-0394-4e3d-8e65-5cd87a82b8a9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c14c7f18-77db-4e3c-8939-e6001cb95db0	Kofe Nescafe Gold "Sumatra" çüýşe gapda 85 gr	Kofe Nescafe Gold "Sumatra" çüýşe gapda 85 gr	2022-10-27 13:09:04.159628+05	2022-11-09 15:09:20.835667+05	\N	kofe-nescafe-gold-sumatra-cuyse-gapda-85-gr
e504255e-51cb-4583-b37d-d91624aa94b1	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c14c7f18-77db-4e3c-8939-e6001cb95db0	Кофе Nescafe Gold "Sumatra" банка 85 гр	Кофе Nescafe Gold "Sumatra" банка 85 гр	2022-10-27 13:09:04.169875+05	2022-11-09 15:09:20.835667+05	\N	kofe-nescafe-gold-sumatra-banka-85-gr
68216fcf-1e95-4adf-9069-b46251e8d061	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	77ecf422-b48b-45fd-8e58-380e23d74c4c	Banan (~900-1.1 kg)	Banan (~900-1.1 kg)	2022-10-28 00:08:25.813149+05	2022-11-09 15:11:14.496257+05	\N	banan-900-1-1-kg
78a56e75-aa39-4cbd-9108-489192e01286	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	0946a0f5-d23f-4660-9151-80ef91ae9747	Sowgatlyk toplumy Head & Shoulders "Saç üçin balzam 275 ml + Goňaga garşy şampun 400 ml	Sowgatlyk toplumy Head & Shoulders "Saç üçin balzam 275 ml + Goňaga garşy şampun 400 ml	2022-10-27 13:32:30.026538+05	2022-11-09 15:11:14.496257+05	\N	sowgatlyk-toplumy-head-and-shoulders-sac-ucin-balzam-275-ml-gonaga-garsy-sampun-400-ml
9f1877ae-af11-4021-a62c-8e7ff715a6f4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	0946a0f5-d23f-4660-9151-80ef91ae9747	Подарочный Набор Head & Shoulders "Бальзам-ополаскиватель для волос 275 мл + Шампунь против перхоти 400 мл	Подарочный Набор Head & Shoulders "Бальзам-ополаскиватель для волос 275 мл + Шампунь против перхоти 400 мл	2022-10-27 13:32:30.038394+05	2022-11-09 15:11:14.496257+05	\N	podarochnyi-nabor-head-and-shoulders-bal-zam-opolaskivatel-dlia-volos-275-ml-shampun-protiv-perkhoti-400-ml
e8b956d7-8e65-44bc-a7a1-02fd3df23338	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	03050bc6-6223-49f3-b729-397fd3b6b285	(Aksiýa!) Kofe Nescafe Gold, paket gapda 220 gr + Kofe Nescafe Classic, 3x1 kiçi paket 14.5 gr	(Aksiýa!) Kofe Nescafe Gold, paket gapda 220 gr + Kofe Nescafe Classic, 3x1 kiçi paket 14.5 gr	2022-10-27 13:36:08.808395+05	2022-11-09 15:11:14.496257+05	\N	aksiya-kofe-nescafe-gold-paket-gapda-220-gr-kofe-nescafe-classic-3x1-kici-paket-14-5-gr
ee30076c-a648-4de4-9f62-238333d88481	aea98b93-7bdf-455b-9ad4-a259d69dc76e	03050bc6-6223-49f3-b729-397fd3b6b285	(Акция!) Кофе Nescafe Gold, пакет 220 г + Кофе Nescafe Classic 3в1, стик 14.5 гр	(Акция!) Кофе Nescafe Gold, пакет 220 г + Кофе Nescafe Classic 3в1, стик 14.5 гр	2022-10-27 13:36:08.818396+05	2022-11-09 15:11:14.496257+05	\N	aktsiia-kofe-nescafe-gold-paket-220-g-kofe-nescafe-classic-3v1-stik-14-5-gr
015b37fc-77f0-49ff-818d-a4443e60a454	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	8b481e58-cd39-4761-a052-75e30124689a	(4+1) Kofe Jacobs Monarch 47.5 gr (5 sany)	(4+1) Kofe Jacobs Monarch 47.5 gr (5 sany)	2022-10-27 13:38:14.188215+05	2022-11-09 15:11:14.496257+05	\N	4-1-kofe-jacobs-monarch-47-5-gr-5-sany
7bed1227-d2e6-4ba7-844f-3ffc70027a90	aea98b93-7bdf-455b-9ad4-a259d69dc76e	8b481e58-cd39-4761-a052-75e30124689a	(4+1) Кофе Jacobs Monarch 47.5 г (5 шт)	(4+1) Кофе Jacobs Monarch 47.5 г (5 шт)	2022-10-27 13:38:14.199712+05	2022-11-09 15:11:14.496257+05	\N	4-1-kofe-jacobs-monarch-47-5-g-5-sht
db6e8eac-1e4b-48ab-9082-040f0635ebe1	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	83da5c7b-bffe-4450-97c9-0f376441b1d4	Sowgatlyk toplum Le Petit Marseillais erkekler üçin gel-şampun "Narpyz we Laým" 250 ml	Sowgatlyk toplum Le Petit Marseillais erkekler üçin gel-şampun "Narpyz we Laým" 250 ml	2022-10-27 13:30:49.691056+05	2022-11-09 15:09:20.835667+05	\N	sowgatlyk-toplum-le-petit-marseillais-erkekler-ucin-gel-sampun-narpyz-we-laym-250-ml
e6457a78-48e1-4f73-b380-75f76892d146	aea98b93-7bdf-455b-9ad4-a259d69dc76e	83da5c7b-bffe-4450-97c9-0f376441b1d4	Подарочный набор Le Petit Marseillais гель-шампунь для мужчин "Мята и Лайм" 250 мл	Подарочный набор Le Petit Marseillais гель-шампунь для мужчин "Мята и Лайм" 250 мл	2022-10-27 13:30:49.702691+05	2022-11-09 15:09:20.835667+05	\N	podarochnyi-nabor-le-petit-marseillais-gel-shampun-dlia-muzhchin-miata-i-laim-250-ml
eee6b1f8-7db8-492a-889d-70e232d1e2d7	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	81462bfa-36df-4e09-aa46-c6fa1ab86de6	(2+1) Dezodorant XO "Aqua joy" 150 ml, Dezodorant XO "Nice girl" 150 ml + Dezodorant XO MEN "Absolute blue" 150 ml	(2+1) Dezodorant XO "Aqua joy" 150 ml, Dezodorant XO "Nice girl" 150 ml + Dezodorant XO MEN "Absolute blue" 150 ml	2022-10-27 23:17:42.819236+05	2022-11-09 15:11:14.496257+05	\N	2-1-dezodorant-xo-aqua-joy-150-ml-dezodorant-xo-nice-girl-150-ml-dezodorant-xo-men-absolute-blue-150-ml
33020f9e-708f-4ab9-8c41-2296b4daa6f2	aea98b93-7bdf-455b-9ad4-a259d69dc76e	81462bfa-36df-4e09-aa46-c6fa1ab86de6	(2+1) Дезодорант XO "Aqua joy" 150 мл, Дезодорант XO "Nice girl" 150 мл + Дезодорант XO MEN "Absolute blue" 150 мл	(2+1) Дезодорант XO "Aqua joy" 150 мл, Дезодорант XO "Nice girl" 150 мл + Дезодорант XO MEN "Absolute blue" 150 мл	2022-10-27 23:17:42.829392+05	2022-11-09 15:11:14.496257+05	\N	2-1-dezodorant-xo-aqua-joy-150-ml-dezodorant-xo-nice-girl-150-ml-dezodorant-xo-men-absolute-blue-150-ml
e4bb9f51-5ca4-4899-9ca0-05a70da2ec19	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	35f5f2d8-9271-469f-bde1-2314c18ea574	Kofe Жокей "Триумф" 2 gr	Kofe Жокей "Триумф" 2 gr	2022-10-27 23:23:19.455084+05	2022-11-09 15:11:14.496257+05	\N	kofe-zhokei-triumf-2-gr
00964051-1dfc-41bf-a4f5-b0f519f45aa4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	35f5f2d8-9271-469f-bde1-2314c18ea574	Кофе Жокей "Триумф" sublimirlenen 2 г	Кофе Жокей "Триумф" sublimirlenen 2 г	2022-10-27 23:23:19.467008+05	2022-11-09 15:11:14.496257+05	\N	kofe-zhokei-triumf-sublimirlenen-2-g
66e6eca6-3d7e-4143-9180-263c074d4a7b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d987b7ad-257e-4ae2-befb-b7d369252a54	Alma gülgüne Gök Önüm 1 kg (± 50 gr)	Alma gülgüne Gök Önüm 1 kg (± 50 gr)	2022-10-27 23:53:01.823656+05	2022-11-09 15:11:14.496257+05	\N	alma-gulgune-gok-onum-1-kg-50-gr
ebafc0f4-eb8c-4ff8-94b3-2c34043bc2f7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d987b7ad-257e-4ae2-befb-b7d369252a54	Яблоки Gök Önüm 1 кг (± 50 г)	Яблоки Gök Önüm 1 кг (± 50 г)	2022-10-27 23:53:01.837748+05	2022-11-09 15:11:14.496257+05	\N	iabloki-gok-onum-1-kg-50-g
42d862b6-68c9-4ecc-906c-9ee070414105	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	(2+1) Ýüz üçin gara maska plýonka "Zenix" paket 15 gr (3 sany)	(2+1) Ýüz üçin gara maska plýonka "Zenix" paket 15 gr (3 sany)	2022-10-27 23:15:36.076305+05	2022-11-09 15:09:20.835667+05	\N	2-1-yuz-ucin-gara-maska-plyonka-zenix-paket-15-gr-3-sany
f81b6304-e768-45bc-9354-df9ba609cebe	aea98b93-7bdf-455b-9ad4-a259d69dc76e	4bb06dbd-e4b2-4148-bb61-b1429d8cfc40	(2+1) Черная маска пленка "Zenix" пакетик 15 г (3 шт)	(2+1) Черная маска пленка "Zenix" пакетик 15 г (3 шт)	2022-10-27 23:15:36.091709+05	2022-11-09 15:09:20.835667+05	\N	2-1-chernaia-maska-plenka-zenix-paketik-15-g-3-sht
e1a326cb-e956-4695-a313-54410dc914ab	aea98b93-7bdf-455b-9ad4-a259d69dc76e	77ecf422-b48b-45fd-8e58-380e23d74c4c	Банан (~900-1.1 кг)	Банан (~900-1.1 кг)	2022-10-28 00:08:25.821742+05	2022-11-09 15:11:14.496257+05	\N	banan-900-1-1-kg
b71c3749-9a42-44d0-bc07-0253c0d85a93	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	9b9ef1ce-2f3d-4051-8e88-5c301bd68554	Alma gyzyl Türkiýe Gök Önüm (1 kg ±50 gr)	Alma gyzyl Türkiýe Gök Önüm (1 kg ±50 gr)	2022-10-28 00:09:12.503503+05	2022-11-09 15:11:14.496257+05	\N	alma-gyzyl-turkiye-gok-onum-1-kg-50-gr
c49cacd8-f994-4e52-9b65-3286c8c4e159	aea98b93-7bdf-455b-9ad4-a259d69dc76e	9b9ef1ce-2f3d-4051-8e88-5c301bd68554	Яблоки красные Турция Gök Önüm (1 кг ±50 г)	Яблоки красные Турция Gök Önüm (1 кг ±50 г)	2022-10-28 00:09:12.515179+05	2022-11-09 15:11:14.496257+05	\N	iabloki-krasnye-turtsiia-gok-onum-1-kg-50-g
f2011438-66f4-41f1-873b-f12728b5b2a1	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	214befd8-68bd-484a-a8d5-8e2d0b73931c	Kinza	Kinza	2022-10-28 00:12:34.811612+05	2022-11-09 15:11:14.496257+05	\N	kinza
46dd8b80-ce85-4522-9ac9-ec67900551d3	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	20a8c487-9a56-4fb6-8fa0-13facaf96109	Armyt ýerli Gök Önüm (1 kg)	Armyt ýerli Gök Önüm (1 kg)	2022-10-28 00:09:53.593406+05	2022-11-09 15:11:14.496257+05	\N	armyt-yerli-gok-onum-1-kg
73615719-aba4-4ea5-a724-ebf130063ffd	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	360ebeac-853e-45a5-ab7f-838430b0c442	Kofe Jardin 5 "Colombia Medellin" sublimirlenen 95 gr	Kofe Jardin 5 "Colombia Medellin" sublimirlenen 95 gr	2022-10-27 23:19:20.778041+05	2022-11-09 15:09:20.835667+05	\N	kofe-jardin-5-colombia-medellin-sublimirlenen-95-gr
6699fccd-fb74-4b18-bbe6-9bdf8773cf10	aea98b93-7bdf-455b-9ad4-a259d69dc76e	360ebeac-853e-45a5-ab7f-838430b0c442	Кофе Jardin 5 "Colombia Medellin" сублимированный 95 г	Кофе Jardin 5 "Colombia Medellin" сублимированный 95 г	2022-10-27 23:19:20.795882+05	2022-11-09 15:09:20.835667+05	\N	kofe-jardin-5-colombia-medellin-sublimirovannyi-95-g
13937dc7-d066-4c16-9b66-221fd9b53710	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	fe309360-c5dd-406a-9957-3d898ea85dfc	Kofe Jardin 5 "Colombia Medellin" sublimirlenen 75 gr	Kofe Jardin 5 "Colombia Medellin" sublimirlenen 75 gr	2022-10-27 23:20:14.825585+05	2022-11-09 15:09:20.835667+05	\N	kofe-jardin-5-colombia-medellin-sublimirlenen-75-gr
6c4163b1-d636-4325-ac2c-abe574a244b8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	fe309360-c5dd-406a-9957-3d898ea85dfc	Кофе Jardin 5 "Colombia Medellin" сублимированный 75 г	Кофе Jardin 5 "Colombia Medellin" сублимированный 75 г	2022-10-27 23:20:14.835738+05	2022-11-09 15:09:20.835667+05	\N	kofe-jardin-5-colombia-medellin-sublimirovannyi-75-g
c0d4842e-33b6-4290-b395-57e350b437a2	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	febf699d-ca37-458a-b121-b5b70bbc7db0	Kofe Jardin 5 "Colombia Medellin" sublimirlenen 150 gr	Kofe Jardin 5 "Colombia Medellin" sublimirlenen 150 gr	2022-10-27 23:21:07.471715+05	2022-11-09 15:09:20.835667+05	\N	kofe-jardin-5-colombia-medellin-sublimirlenen-150-gr
0c72b07c-633a-4fdc-aa9b-fc811e8c0d61	aea98b93-7bdf-455b-9ad4-a259d69dc76e	febf699d-ca37-458a-b121-b5b70bbc7db0	Кофе Jardin 5 "Colombia Medellin" сублимированный 150 г	Кофе Jardin 5 "Colombia Medellin" сублимированный 150 г	2022-10-27 23:21:07.482116+05	2022-11-09 15:09:20.835667+05	\N	kofe-jardin-5-colombia-medellin-sublimirovannyi-150-g
4a1514bb-befc-4e7e-8e54-49b5d1232141	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	802b422b-710a-420b-860e-59b7f49d10bd	Kofe Жокей "Импер" sublimirlenen 150 gr	Kofe Жокей "Импер" sublimirlenen 150 gr	2022-10-27 23:22:15.71738+05	2022-11-09 15:09:20.835667+05	\N	kofe-zhokei-imper-sublimirlenen-150-gr
e17bf906-e560-429b-9e8f-b0656b428ec4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	802b422b-710a-420b-860e-59b7f49d10bd	Кофе Жокей "Импер" sublimirlenen 150 г	Кофе Жокей "Импер" sublimirlenen 150 г	2022-10-27 23:22:15.728698+05	2022-11-09 15:09:20.835667+05	\N	kofe-zhokei-imper-sublimirlenen-150-g
2e49076c-2fc3-47ef-9b34-128197ab4b0e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	20a8c487-9a56-4fb6-8fa0-13facaf96109	Груша местная Gök Önüm (1 кг)	Груша местная Gök Önüm (1 кг)	2022-10-28 00:09:53.605862+05	2022-11-09 15:11:14.496257+05	\N	grusha-mestnaia-gok-onum-1-kg
cc8cb40c-1ffa-4d8b-9fa9-6d9f70f80085	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	7078e107-dd52-4da1-8007-29ed7cf731fb	Limon ýerli (1 sany)	Limon ýerli (1 sany)	2022-10-28 00:10:28.3597+05	2022-11-09 15:11:14.496257+05	\N	limon-yerli-1-sany
d1c3f296-ea82-4422-8dab-dfe6bfa8ffc5	aea98b93-7bdf-455b-9ad4-a259d69dc76e	7078e107-dd52-4da1-8007-29ed7cf731fb	Лимон (1 шт)	Лимон (1 шт)	2022-10-28 00:10:28.370068+05	2022-11-09 15:11:14.496257+05	\N	limon-1-sht
1ca5933c-e476-46c1-a240-ec996227aafb	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	5f6aba1c-66df-4791-b85e-b0a90ccffc20	Micro Green "Kelem" 50 gr	Micro Green "Kelem" 50 gr	2022-10-28 00:11:26.36221+05	2022-11-09 15:11:14.496257+05	\N	micro-green-kelem-50-gr
e67a775a-8e6c-4cba-ac45-489adc167098	aea98b93-7bdf-455b-9ad4-a259d69dc76e	5f6aba1c-66df-4791-b85e-b0a90ccffc20	Micro Green "Капуста" 50 gr	Micro Green "Капуста" 50 gr	2022-10-28 00:11:26.38452+05	2022-11-09 15:11:14.496257+05	\N	micro-green-kapusta-50-gr
af5869cf-4d49-41c4-b864-439520dd11dc	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	01dc8537-7ec1-4c48-bcce-3734f1ac598a	Micro Green "Rukkola" 100 gr	Micro Green "Rukkola" 100 gr	2022-10-28 00:11:58.362561+05	2022-11-09 15:11:14.496257+05	\N	micro-green-rukkola-100-gr
530e577c-a292-4f68-8b3f-b935df1a9885	aea98b93-7bdf-455b-9ad4-a259d69dc76e	01dc8537-7ec1-4c48-bcce-3734f1ac598a	Micro Green "Руккола" 100 г	Micro Green "Руккола" 100 г	2022-10-28 00:11:58.373677+05	2022-11-09 15:11:14.496257+05	\N	micro-green-rukkola-100-g
42560a2d-79a5-4578-a3c4-14de767855ed	aea98b93-7bdf-455b-9ad4-a259d69dc76e	214befd8-68bd-484a-a8d5-8e2d0b73931c	Кинза пучок	Кинза пучок	2022-10-28 00:12:34.831585+05	2022-11-09 15:11:14.496257+05	\N	kinza-puchok
bbf1c432-2594-43ee-9eee-95a228655707	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	a316070c-409c-42c2-85df-bb7509a24c54	Gök ot assorti	Gök ot assorti	2022-10-28 00:13:05.564306+05	2022-11-09 15:11:14.496257+05	\N	gok-ot-assorti
6da583d1-6b56-49dd-be6f-6fd01513f2d6	aea98b93-7bdf-455b-9ad4-a259d69dc76e	a316070c-409c-42c2-85df-bb7509a24c54	Зелень ассорти	Зелень ассорти	2022-10-28 00:13:05.575646+05	2022-11-09 15:11:14.496257+05	\N	zelen-assorti
618a3351-d17f-4bb9-b1b8-012609af2e13	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	89172d2f-b5b3-4b26-a299-dc7e8a71d16e	Duzlanan gök önümler 200 gr	Duzlanan gök önümler 200 gr	2022-10-28 00:13:38.976924+05	2022-11-09 15:11:14.496257+05	\N	duzlanan-gok-onumler-200-gr
e3f5fb9c-19e0-46f3-bf2b-04672ad01a65	aea98b93-7bdf-455b-9ad4-a259d69dc76e	89172d2f-b5b3-4b26-a299-dc7e8a71d16e	Квашеные овощи 200 г	Квашеные овощи 200 г	2022-10-28 00:13:38.987966+05	2022-11-09 15:11:14.496257+05	\N	kvashenye-ovoshchi-200-g
449ec269-1171-414e-8ea4-8a6259a15df2	aea98b93-7bdf-455b-9ad4-a259d69dc76e	c82fef0a-ad15-4b07-8855-910fc4708af1	Базилик Рейхан пучок	Базилик Рейхан пучок	2022-10-28 00:14:34.6995+05	2022-11-09 15:11:14.496257+05	\N	bazilik-reikhan-puchok
e40e3a9a-d64b-473d-9450-39f23cc41f75	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	311aa4c1-6002-4acf-b1b5-e2aa7896def7	Şugundyr (1 kg ±50 gr)	Şugundyr (1 kg ±50 gr)	2022-10-28 00:15:19.604609+05	2022-11-09 15:11:14.496257+05	\N	sugundyr-1-kg-50-gr
2f272516-1bf8-4cd3-8520-a86449b13688	aea98b93-7bdf-455b-9ad4-a259d69dc76e	311aa4c1-6002-4acf-b1b5-e2aa7896def7	Свекла (1 кг ±50 г)	Свекла (1 кг ±50 г)	2022-10-28 00:15:19.617495+05	2022-11-09 15:11:14.496257+05	\N	svekla-1-kg-50-g
5ef3ad2a-5d58-4fc3-b8fd-0229ea6e492a	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	32055a0a-2d59-45a9-89b0-761d1f6ad047	Kelem (1-1.5 kg)	Kelem (1-1.5 kg)	2022-10-28 01:13:10.64922+05	2022-11-09 15:11:14.496257+05	\N	kelem-1-1-5-kg
b80cafd5-6396-4d9a-bbd1-47b59adc92c4	aea98b93-7bdf-455b-9ad4-a259d69dc76e	32055a0a-2d59-45a9-89b0-761d1f6ad047	Капуста ( 1-1.5 кг)	Капуста ( 1-1.5 кг)	2022-10-28 01:13:10.738455+05	2022-11-09 15:11:14.496257+05	\N	kapusta-1-1-5-kg
188357e4-7b6b-4af8-ab40-7807e8e841e4	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	19e94bc7-398a-45d9-b3cf-5c8d50550e48	Sogan 1 kg (±20 gr)	Sogan 1 kg (±20 gr)	2022-10-28 01:13:56.381852+05	2022-11-09 15:11:14.496257+05	\N	sogan-1-kg-20-gr
cd65b8ac-a598-4aba-a908-a082685db458	aea98b93-7bdf-455b-9ad4-a259d69dc76e	19e94bc7-398a-45d9-b3cf-5c8d50550e48	Лук репчетый 1 кг (±20 г )	Лук репчетый 1 кг (±20 г )	2022-10-28 01:13:56.394192+05	2022-11-09 15:11:14.496257+05	\N	luk-repchetyi-1-kg-20-g
d4199db1-5df5-4584-ac20-821c2606d7f8	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	fd2c148f-70b5-48ba-8cf6-65f26438b46d	Pomidor "Nowça" 1 kg	Pomidor "Nowça" 1 kg	2022-10-28 01:15:13.848781+05	2022-11-09 15:11:14.496257+05	\N	pomidor-nowca-1-kg
63b6c3ed-1d61-4f00-9f0e-c3f5accd9147	aea98b93-7bdf-455b-9ad4-a259d69dc76e	fd2c148f-70b5-48ba-8cf6-65f26438b46d	Помидор "Новча" (1 кг)	Помидор "Новча" (1 кг)	2022-10-28 01:15:13.859782+05	2022-11-09 15:11:14.496257+05	\N	pomidor-novcha-1-kg
6899fc36-f2b9-4e2e-81a9-7f517cd95079	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	793be71f-b0fa-43a2-b527-5fb09236f530	Sowgatlyk toplumy MEN DEEP CLEANING duş üçin krem-gel 300 ml+ýuwmak üçin gel HYDRO ENERgETIC 150 ml	Sowgatlyk toplumy MEN DEEP CLEANING duş üçin krem-gel 300 ml+ýuwmak üçin gel HYDRO ENERgETIC 150 ml	2022-10-27 13:22:11.416525+05	2022-11-09 15:11:14.496257+05	\N	sowgatlyk-toplumy-men-deep-cleaning-dus-ucin-krem-gel-300-ml-yuwmak-ucin-gel-hydro-energetic-150-ml
6d8cc27d-aa60-490a-b282-03f413bfa908	aea98b93-7bdf-455b-9ad4-a259d69dc76e	793be71f-b0fa-43a2-b527-5fb09236f530	Подарочный набор MEN DEEP CLEANINg крем-грель для душа 300 мл+грель для умыв HYDRO ENERgETIC 150 мл	Подарочный набор MEN DEEP CLEANINg крем-грель для душа 300 мл+грель для умыв HYDRO ENERgETIC 150 мл	2022-10-27 13:22:11.428036+05	2022-11-09 15:11:14.496257+05	\N	podarochnyi-nabor-men-deep-cleaning-krem-grel-dlia-dusha-300-ml-grel-dlia-umyv-hydro-energetic-150-ml
950aa440-28cf-4c13-8685-d8baf19c2bf8	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	Sowgatlyk toplumy UFC x EXXE sakgal syrmak üçin köpürjik + sakgal syrmak üçin gel + duş geli Ultimate freshness	Sowgatlyk toplumy UFC x EXXE sakgal syrmak üçin köpürjik + sakgal syrmak üçin gel + duş geli Ultimate freshness	2022-10-27 13:24:26.628997+05	2022-11-09 15:11:14.496257+05	\N	sowgatlyk-toplumy-ufc-x-exxe-sakgal-syrmak-ucin-kopurjik-sakgal-syrmak-ucin-gel-dus-geli-ultimate-freshness
42c64158-3179-4b2e-9d3b-8766e3586931	aea98b93-7bdf-455b-9ad4-a259d69dc76e	e34a20fa-3aef-4ba6-92ba-79d3649c61a6	Подарочный набор UFC x EXXE пена для бритья + крем-бальзам после бритья + гель для душа Ultimate Freshness	Подарочный набор UFC x EXXE пена для бритья + крем-бальзам после бритья + гель для душа Ultimate Freshness	2022-10-27 13:24:26.641508+05	2022-11-09 15:11:14.496257+05	\N	podarochnyi-nabor-ufc-x-exxe-pena-dlia-brit-ia-krem-bal-zam-posle-brit-ia-gel-dlia-dusha-ultimate-freshness
5c87b6d0-7b5d-40df-a32d-9a12f8205e69	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	3f397126-6d8d-4a0d-982c-01fd00526957	Sowgatlyk toplumy UFC x EXXE duş geli + şampun Carbon Hit	Sowgatlyk toplumy UFC x EXXE duş geli + şampun Carbon Hit	2022-10-27 13:26:05.309306+05	2022-11-09 15:11:14.496257+05	\N	sowgatlyk-toplumy-ufc-x-exxe-dus-geli-sampun-carbon-hit
6caa3ccc-e904-413b-8730-80dbddf15790	aea98b93-7bdf-455b-9ad4-a259d69dc76e	3f397126-6d8d-4a0d-982c-01fd00526957	Подарочный набор UFC x EXXE гель для душа + шампунь Carbon Hit	Подарочный набор UFC x EXXE гель для душа + шампунь Carbon Hit	2022-10-27 13:26:05.319675+05	2022-11-09 15:11:14.496257+05	\N	podarochnyi-nabor-ufc-x-exxe-gel-dlia-dusha-shampun-carbon-hit
0283d1be-d69a-467f-af88-d78dfbb76c59	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	Dodak üçin suwuk pomada Farmasi "Nude Essence" Matte 4 ml (03)	Dodak üçin suwuk pomada Farmasi "Nude Essence" Matte 4 ml (03)	2022-10-28 01:24:16.809364+05	2022-11-09 15:09:20.835667+05	\N	dodak-ucin-suwuk-pomada-farmasi-nude-essence-matte-4-ml-03
b39c1d8a-997c-4141-b54e-9b2d22bd9658	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d6dfc8f9-27d9-45f9-811d-2a93fa0f7d35	Матовая жидкая губная помада Farmasi "Nude Essence" Matte 4 мл (03)	Матовая жидкая губная помада Farmasi "Nude Essence" Matte 4 мл (03)	2022-10-28 01:24:16.821316+05	2022-11-09 15:09:20.835667+05	\N	matovaia-zhidkaia-gubnaia-pomada-farmasi-nude-essence-matte-4-ml-03
832430ab-524a-46d5-b34b-29904fc4ea57	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	45a9f186-2521-4eef-a4e0-b5c253c70878	Matte pomada Farmasi "Heat Wave" 4 ml (211)	Matte pomada Farmasi "Heat Wave" 4 ml (211)	2022-10-28 01:24:52.887674+05	2022-11-09 15:09:20.835667+05	\N	matte-pomada-farmasi-heat-wave-4-ml-211
f9ae1956-4812-4412-a93c-e3d7460c8977	aea98b93-7bdf-455b-9ad4-a259d69dc76e	45a9f186-2521-4eef-a4e0-b5c253c70878	Матовая помада для губ Farmasi "Heat Wave" 4 мл (211)	Матовая помада для губ Farmasi "Heat Wave" 4 мл (211)	2022-10-28 01:24:52.900396+05	2022-11-09 15:09:20.835667+05	\N	matovaia-pomada-dlia-gub-farmasi-heat-wave-4-ml-211
857cebf6-8605-4f4a-9962-d30f1d5f98fd	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	fa148eb2-520f-430e-bd8d-9d5a166d0600	Göz galamy Farmasi "Express" 1.14 gr (09)	Göz galamy Farmasi "Express" 1.14 gr (09)	2022-10-28 01:25:37.413225+05	2022-11-09 15:09:20.835667+05	\N	goz-galamy-farmasi-express-1-14-gr-09
59ce55bb-932c-4a82-8cc7-208e43b9abfe	aea98b93-7bdf-455b-9ad4-a259d69dc76e	fa148eb2-520f-430e-bd8d-9d5a166d0600	Карандаш для глаз Farmasi "Express" 1.14 гр (09)	Карандаш для глаз Farmasi "Express" 1.14 гр (09)	2022-10-28 01:25:37.4344+05	2022-11-09 15:09:20.835667+05	\N	karandash-dlia-glaz-farmasi-express-1-14-gr-09
fdecb62e-e026-400c-8146-269552268363	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	d085e5a4-8229-4177-b5e1-623e80846017	Tebigy ereýän kofe Maxwell House 150 gr	Tebigy ereýän kofe Maxwell House 150 gr	2022-10-27 12:45:23.512856+05	2022-11-09 15:09:20.835667+05	\N	tebigy-ereyan-kofe-maxwell-house-150-gr
b5441981-2b60-4d64-8ded-f6cd68e0bec7	aea98b93-7bdf-455b-9ad4-a259d69dc76e	d085e5a4-8229-4177-b5e1-623e80846017	Растворимый кофе Maxwell House 150 г	Растворимый кофе Maxwell House 150 г	2022-10-27 12:45:23.533062+05	2022-11-09 15:09:20.835667+05	\N	rastvorimyi-kofe-maxwell-house-150-g
cc5a5be3-8749-4d81-8a96-fc25163d045b	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	Tebigy ereýän kofe Maxwell House 200+50 gr	Tebigy ereýän kofe Maxwell House 200+50 gr	2022-10-27 12:47:51.168874+05	2022-11-09 15:09:20.835667+05	\N	tebigy-ereyan-kofe-maxwell-house-200-50-gr
0d87cb5f-4d1f-49e8-a1b4-c26353bd8c8f	aea98b93-7bdf-455b-9ad4-a259d69dc76e	4a5bdcbf-712d-45ca-baa8-1318c6e2fb3c	Растворимый кофе Maxwell House 200+50 г	Растворимый кофе Maxwell House 200+50 г	2022-10-27 12:47:51.1792+05	2022-11-09 15:09:20.835667+05	\N	rastvorimyi-kofe-maxwell-house-200-50-g
4c993d4a-f4d6-4d33-afca-8cf9fbf652c0	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	Kofe Mövenpick "Gold Original 100% Arabica" 100 gr	Kofe Mövenpick "Gold Original 100% Arabica" 100 gr	2022-10-27 12:49:37.972094+05	2022-11-09 15:09:20.835667+05	\N	kofe-movenpick-gold-original-100-arabica-100-gr
6ba071ce-7548-48e4-ad83-444f175281ba	aea98b93-7bdf-455b-9ad4-a259d69dc76e	2e05c0d9-f7a0-4dc8-ab1a-171f8d725d33	Кофе Mövenpick "Gold Original 100% Arabica" 100 г	Кофе Mövenpick "Gold Original 100% Arabica" 100 г	2022-10-27 12:49:37.981834+05	2022-11-09 15:09:20.835667+05	\N	kofe-movenpick-gold-original-100-arabica-100-g
c7344a8d-9a89-4acc-8cf3-be8643124045	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	70a75d8b-d570-41d4-95cb-2199f4417542	Üwelen kofe Mövenpick Der Himmlische 250 gr	Üwelen kofe Mövenpick Der Himmlische 250 gr	2022-10-27 13:05:10.332333+05	2022-11-09 15:09:20.835667+05	\N	uwelen-kofe-movenpick-der-himmlische-250-gr
e0565164-63dd-4bf6-8bc2-d1399961c610	aea98b93-7bdf-455b-9ad4-a259d69dc76e	70a75d8b-d570-41d4-95cb-2199f4417542	Молотый кофе Mövenpick Der Himmlische 250 гр	Молотый кофе Mövenpick Der Himmlische 250 гр	2022-10-27 13:05:10.705616+05	2022-11-09 15:09:20.835667+05	\N	molotyi-kofe-movenpick-der-himmlische-250-gr
86cac080-d98c-4a2c-ba86-6703e153d6d0	aea98b93-7bdf-455b-9ad4-a259d69dc76e	9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	Увлажняющая сыворотка Farmasi Dr. C.Tuna "Aqua" Hydra Drops 30 мл	Увлажняющая сыворотка Farmasi Dr. C.Tuna "Aqua" Hydra Drops 30 мл	2022-10-28 01:26:46.982412+05	2022-11-09 15:11:14.496257+05	\N	uvlazhniaiushchaia-syvorotka-farmasi-dr-c-tuna-aqua-hydra-drops-30-ml
0ce73b75-e1ac-40e7-8598-829bdcc69b87	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	74ec9e27-de0b-44d9-8036-3fae8be486a9	Diş pastasy Colgate "Herbal Smoothie" serginlediji 75 ml	Diş pastasy Colgate "Herbal Smoothie" serginlediji 75 ml	2022-10-28 01:27:54.394089+05	2022-11-09 15:11:14.496257+05	\N	dis-pastasy-colgate-herbal-smoothie-serginlediji-75-ml
557fe4fa-2520-4689-90d5-05d89a73aa77	aea98b93-7bdf-455b-9ad4-a259d69dc76e	74ec9e27-de0b-44d9-8036-3fae8be486a9	Зубная паста Colgate "Herbal Smoothie" освежающий 75 мл	Зубная паста Colgate "Herbal Smoothie" освежающий 75 мл	2022-10-28 01:27:54.405186+05	2022-11-09 15:11:14.496257+05	\N	zubnaia-pasta-colgate-herbal-smoothie-osvezhaiushchii-75-ml
6cece3d5-d3ea-4aad-84e1-72f7e91bb089	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	Diş pastasy Colgate "Vitamin Coctail" serginlediji 75 ml	Diş pastasy Colgate "Vitamin Coctail" serginlediji 75 ml	2022-10-28 01:28:28.417763+05	2022-11-09 15:11:14.496257+05	\N	dis-pastasy-colgate-vitamin-coctail-serginlediji-75-ml
10c8d57d-0e30-424c-a32a-864458e3654e	aea98b93-7bdf-455b-9ad4-a259d69dc76e	332d15a5-8f2a-4ea5-8eac-a0e571fcdce5	Зубная паста Colgate "Vitamin Coctail" освежающий 75 мл	Зубная паста Colgate "Vitamin Coctail" освежающий 75 мл	2022-10-28 01:28:28.46043+05	2022-11-09 15:11:14.496257+05	\N	zubnaia-pasta-colgate-vitamin-coctail-osvezhaiushchii-75-ml
8254faf7-9f84-475d-a829-1c7b2c6e0d39	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	2eb8a13f-3edc-4422-b772-e57bfd8f8797	Diş pastasy Colgate "Bejergi otlary" 100 ml	Diş pastasy Colgate "Bejergi otlary" 100 ml	2022-10-28 01:29:09.04184+05	2022-11-09 15:11:14.496257+05	\N	dis-pastasy-colgate-bejergi-otlary-100-ml
da57822b-a5eb-4c96-bbfe-578bd40c389f	aea98b93-7bdf-455b-9ad4-a259d69dc76e	2eb8a13f-3edc-4422-b772-e57bfd8f8797	Зубная паста Colgate "Лечебные травы" 100 мл	Зубная паста Colgate "Лечебные травы" 100 мл	2022-10-28 01:29:09.051954+05	2022-11-09 15:11:14.496257+05	\N	zubnaia-pasta-colgate-lechebnye-travy-100-ml
c7548ef9-4811-423c-b342-e8d5a55295b7	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	18f957f2-216d-4810-b4d7-bd4dd49efd0d	Diş pastasy Colgate "Total 12" Professional arassalaýjy pasta 75 ml	Diş pastasy Colgate "Total 12" Professional arassalaýjy pasta 75 ml	2022-10-28 01:29:49.665417+05	2022-11-09 15:11:14.496257+05	\N	dis-pastasy-colgate-total-12-professional-arassalayjy-pasta-75-ml
6b47b833-242d-4d47-93cb-8d37ad93481b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	18f957f2-216d-4810-b4d7-bd4dd49efd0d	Зубная паста Colgate "Total 12" Профессиональная чистка паста 75 мл	Зубная паста Colgate "Total 12" Профессиональная чистка паста 75 мл	2022-10-28 01:29:49.688606+05	2022-11-09 15:11:14.496257+05	\N	zubnaia-pasta-colgate-total-12-professional-naia-chistka-pasta-75-ml
657c1f2d-54c6-4604-b1c4-d32b43fc2cbe	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ad24153a-997a-46d1-87bb-27aa1e3e8aea	Diş pastasy 32 Жемчужины tutuş maşgala üçin hemmetaraplaýyn ideg 100 gr	Diş pastasy 32 Жемчужины tutuş maşgala üçin hemmetaraplaýyn ideg 100 gr	2022-10-28 01:30:32.911801+05	2022-11-09 15:11:14.496257+05	\N	dis-pastasy-32-zhemchuzhiny-tutus-masgala-ucin-hemmetaraplayyn-ideg-100-gr
51b82883-c7aa-4539-a5a3-e5013f2406b8	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ad24153a-997a-46d1-87bb-27aa1e3e8aea	Зубная паста 32 Жемчужины Комплексный уход для всей семьи 100 гр	Зубная паста 32 Жемчужины Комплексный уход для всей семьи 100 гр	2022-10-28 01:30:32.955604+05	2022-11-09 15:11:14.496257+05	\N	zubnaia-pasta-32-zhemchuzhiny-kompleksnyi-ukhod-dlia-vsei-sem-i-100-gr
d59df776-a4b9-48f6-a222-1c426813e3e5	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	9c655c36-1832-48ca-9f88-c04197f191af	Diş pastasy "Blend-a-Med Complete 7" Herbal, 100 ml	Diş pastasy "Blend-a-Med Complete 7" Herbal, 100 ml	2022-10-28 01:31:06.952867+05	2022-11-09 15:11:14.496257+05	\N	dis-pastasy-blend-a-med-complete-7-herbal-100-ml
8fcde624-a402-46c5-9266-fd2c126d3c06	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	c82fef0a-ad15-4b07-8855-910fc4708af1	Reýhan	Reýhan	2022-10-28 00:14:34.6898+05	2022-11-09 15:11:14.496257+05	\N	reyhan
ad644522-21ea-44dc-b0c7-1ddb0d86dc22	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	bcb52dfc-c957-4d5e-9bbc-1fcb607d3fd6	Sary käşir "Daşoguz" 1 kg	Sary käşir "Daşoguz" 1 kg	2022-10-28 01:15:51.303298+05	2022-11-09 15:11:14.496257+05	\N	sary-kasir-dasoguz-1-kg
e5f5a0cf-3ad4-418a-ac1d-50cb919b432b	aea98b93-7bdf-455b-9ad4-a259d69dc76e	bcb52dfc-c957-4d5e-9bbc-1fcb607d3fd6	Желтая морковь "Daşoguz" 1 кг	Желтая морковь "Daşoguz" 1 кг	2022-10-28 01:15:51.313414+05	2022-11-09 15:11:14.496257+05	\N	zheltaia-morkov-dasoguz-1-kg
78132a22-b4af-4cd8-8f39-6171aacb3d74	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	badd0869-99df-4df3-8a27-5e27c10a861d	Tonalnyý krem Farmasi VFX PRO Camera Ready 30 ml (07)	Tonalnyý krem Farmasi VFX PRO Camera Ready 30 ml (07)	2022-10-28 01:22:44.78443+05	2022-11-09 15:11:14.496257+05	\N	tonalnyy-krem-farmasi-vfx-pro-camera-ready-30-ml-07
5861b129-de86-46ea-8ece-0523d36934db	aea98b93-7bdf-455b-9ad4-a259d69dc76e	badd0869-99df-4df3-8a27-5e27c10a861d	Тональний крем Farmasi VFX PRO Camera Ready 30 мл (07)	Тональний крем Farmasi VFX PRO Camera Ready 30 мл (07)	2022-10-28 01:22:44.795164+05	2022-11-09 15:11:14.496257+05	\N	tonal-nii-krem-farmasi-vfx-pro-camera-ready-30-ml-07
8a2b1f41-c968-49e9-812d-ae3e04639cd9	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	9cd1e4e4-b15c-4ceb-a03c-01e9cfbb224b	Nemlendiriji syworotka Farmasi Dr. C.Tuna "Aqua" Hydra Drops 30 ml	Nemlendiriji syworotka Farmasi Dr. C.Tuna "Aqua" Hydra Drops 30 ml	2022-10-28 01:26:46.960127+05	2022-11-09 15:11:14.496257+05	\N	nemlendiriji-syworotka-farmasi-dr-c-tuna-aqua-hydra-drops-30-ml
84f32104-7086-4282-a1f9-d82afac8d1ff	aea98b93-7bdf-455b-9ad4-a259d69dc76e	9c655c36-1832-48ca-9f88-c04197f191af	Зубная паста "Blend-a-Med Complete 7" с ополаскивателем 100 мл	Зубная паста "Blend-a-Med Complete 7" с ополаскивателем 100 мл	2022-10-28 01:31:07.002492+05	2022-11-09 15:11:14.496257+05	\N	zubnaia-pasta-blend-a-med-complete-7-s-opolaskivatelem-100-ml
a440c88c-0030-445b-b675-94029afeacbc	8723c1c7-aa6d-429f-b8af-ee9ace61f0d7	ccb43083-1c9e-4e84-bffd-ecb28474165e	Sowgatlyk toplumy UFC x EXXE duş geli + dezodorant Ultimate Freshness	Sowgatlyk toplumy UFC x EXXE duş geli + dezodorant Ultimate Freshness	2022-10-27 13:29:08.24936+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05	sowgatlyk-toplumy-ufc-x-exxe-dus-geli-dezodorant-ultimate-freshness
f2d39fe8-264d-4394-964a-88b04f133187	aea98b93-7bdf-455b-9ad4-a259d69dc76e	ccb43083-1c9e-4e84-bffd-ecb28474165e	Подарочный набор UFC x EXXE гель для душа + дезодорант Ultimate Freshness	Подарочный набор UFC x EXXE гель для душа + дезодорант Ultimate Freshness	2022-10-27 13:29:08.266915+05	2022-11-22 09:49:06.741919+05	2022-11-22 09:49:06.741919+05	podarochnyi-nabor-ufc-x-exxe-gel-dlia-dusha-dezodorant-ultimate-freshness
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

SELECT pg_catalog.setval('public.orders_order_number_seq', 49, true);


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
-- Name: main_image main_image_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.main_image
    ADD CONSTRAINT main_image_pkey PRIMARY KEY (id);


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
-- Name: main_image updated_main_image_updated_at; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER updated_main_image_updated_at BEFORE UPDATE ON public.main_image FOR EACH ROW EXECUTE FUNCTION public.update_updated_at();


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
-- Name: products shops_products; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT shops_products FOREIGN KEY (shop_id) REFERENCES public.shops(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

