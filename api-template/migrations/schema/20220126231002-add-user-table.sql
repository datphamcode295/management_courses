-- +migrate Up
CREATE TABLE "user" (
  "id" uuid NOT NULL PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text)::uuid,
  "wallet_address" VARCHAR NOT NULL,
  "referral_code" VARCHAR NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS "user";