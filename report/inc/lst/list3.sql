CREATE ROLE authorized with
    NOSUPERUSER
    NOCREATEDB
    NOCREATEROLE
    NOINHERIT
    NOBYPASSRLS
    NOREPLICATION
    LOGIN
    PASSWORD 'guest'
    CONNECTION LIMIT -1;

GRANT SELECT
    ON ALL TABLES IN SCHEMA public
    TO authorized;

GRANT INSERT (homeid, coords, name)
    ON public.home
    TO authorized;