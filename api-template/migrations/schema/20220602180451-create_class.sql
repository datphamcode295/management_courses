-- +migrate Up
CREATE TABLE "class" (
  "id" TEXT PRIMARY KEY

);
INSERT INTO "class" ("id") VALUES ('1');

-- +migrate Down
DROP TABLE IF EXISTS "class";