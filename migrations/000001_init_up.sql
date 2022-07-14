-- Active: 1657388736800@@127.0.0.1@5432@product_service@public
BEGIN;

SET client_encoding = 'UTF-8';

-- TABLES --

CREATE TABLE public.currency
(
    id SERIAL PRIMARY KEY,
    name TEXT,
    symbol TEXT
);

CREATE TABLE public.category
(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE public.product
(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    image_id UUID,
    price BIGINT,
    currency_id INT,
    rating INT,
    specification JSONB,
    category_id INT NOT NULL,
    created_at TEXT,
    updated_at TEXT
);


-- DATA --

INSERT INTO public.currency (name, symbol)
VALUES ('рубль', '₽');
INSERT INTO public.currency (name, symbol)
VALUES ('доллар', '$');


COMMIT;