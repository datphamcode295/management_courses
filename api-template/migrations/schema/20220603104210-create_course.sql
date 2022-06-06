-- +migrate Up
CREATE TABLE "course" (
  "id" TEXT PRIMARY KEY,
  "class_id" TEXT,
  FOREIGN KEY (class_id)
    REFERENCES class(id)
);

-- +migrate Down
DROP TABLE IF EXISTS "course";