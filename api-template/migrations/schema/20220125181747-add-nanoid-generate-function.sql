
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS pgcrypto;
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION nanoid(size int DEFAULT 12)
RETURNS text AS $$
DECLARE
  id text := '';
  i int := 0;
  urlAlphabet char(64) := '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
  bytes bytea := gen_random_bytes(size);
  byte int;
  pos int;
BEGIN
  WHILE i < size LOOP
    byte := get_byte(bytes, i);
    pos := (byte & 63) + 1; -- + 1 because substr starts at 1 for some reason
    id := id || substr(urlAlphabet, pos, 1);
    i = i + 1;
  END LOOP;
  RETURN id;
END
$$ LANGUAGE PLPGSQL STABLE;
-- +migrate StatementEnd
-- +migrate Down
DROP EXTENSION IF EXISTS pgcrypto;
DROP FUNCTION IF EXISTS nanoid;
