-- +migrate Up
CREATE TABLE "user" (
  "id" TEXT PRIMARY KEY,
  "wallet_address" VARCHAR NOT NULL,
  "referral_code" VARCHAR NOT NULL,
  "created_at" timestamptz DEFAULT now(),
  "updated_at" timestamptz DEFAULT now()
);

-- +migrate Down
DROP TABLE IF EXISTS "user";