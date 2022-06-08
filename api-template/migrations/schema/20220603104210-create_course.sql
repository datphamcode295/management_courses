-- +migrate Up
CREATE TABLE "course" (
  "id" TEXT PRIMARY KEY,
  "name" TEXT,
  "lecturer" TEXT
);
insert into "course" (id, name, lecturer) values ('1','math', 'thanh');
-- +migrate Down
DROP TABLE IF EXISTS "course";