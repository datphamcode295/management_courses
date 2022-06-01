-- +migrate Up

CREATE TABLE "transaction" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text)::uuid,
  "block_number" INT8 NOT NULL,
  "tx_hash" VARCHAR NOT NULL,
  "action_name" VARCHAR NOT NULL,
  "version" INT8,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now()
);

-- +migrate Down

DROP TABLE IF EXISTS "transaction";