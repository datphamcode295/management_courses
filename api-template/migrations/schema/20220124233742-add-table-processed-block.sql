-- +migrate Up

CREATE TABLE "processed_block" (
  "id" text NOT NULL,
  "block_number" int8 NOT NULL
)
;

-- +migrate Down
DROP TABLE IF EXISTS "processed_block";