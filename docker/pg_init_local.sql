-- init-locale.sql
-- This script ensures the database is created with EN-GB locale settings

-- Set the default locale for the database
-- DO $$
-- BEGIN
--     IF NOT EXISTS (SELECT FROM pg_catalog.pg_roles WHERE rolname = 'postgres') THEN
--         CREATE ROLE postgres WITH LOGIN PASSWORD 'mypassword';
--     END IF;
-- END $$;

-- Create the database with specific locale settings
CREATE DATABASE takeaway
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_GB.UTF-8'
    LC_CTYPE = 'en_GB.UTF-8'
    TEMPLATE = template0;

-- Connect to the new database and set up locale-specific settings
\c takeaway

-- Optional: Set session-level locale if needed
SET lc_time TO 'en_GB.UTF-8';