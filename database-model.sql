CREATE DATABASE dev_database;

CREATE TABLE IF NOT EXISTS public.book (
  id uuid primary key,
  title varchar(255) not null,
  isbn varchar(13) not null,
  number_of_pages integer,
  publishing_company varchar(255) not null,
  year_of_edition timestamp with time zone not null
);