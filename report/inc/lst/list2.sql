CREATE ROLE guest with
    NOSUPERUSER
    NOCREATEDB
    NOCREATEROLE
    NOINHERIT
    NOBYPASSRLS
    NOREPLICATION
    LOGIN
    PASSWORD 'guest'
    connection limit -1;

GRANT select
	on public."client"
    TO guest;

GRANT INSERT
    on public."client"
    TO guest;