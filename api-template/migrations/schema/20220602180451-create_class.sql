-- +migrate Up
CREATE TABLE "class" (
  "id" TEXT PRIMARY KEY,
  "name" TEXT NOT NULL

);
INSERT INTO "class" ("id","name") VALUES ('1','CC03');

-- +migrate Down
DROP TABLE IF EXISTS "class";