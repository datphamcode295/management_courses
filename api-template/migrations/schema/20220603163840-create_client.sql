-- +migrate Up
CREATE TABLE "client" (
  "id" TEXT PRIMARY KEY,
  "role" VARCHAR NOT NULL,
  "class_id" TEXT,
  FOREIGN KEY (class_id)
      REFERENCES class(id)
);

-- +migrate Down
DROP TABLE IF EXISTS "client";