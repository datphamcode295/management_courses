-- +migrate Up
CREATE TYPE role AS ENUM ('admin', 'student');
CREATE TABLE "client" (
  "id" TEXT PRIMARY KEY,
  "username" TEXT NOT NULL,
  "password" TEXT NOT NULL,
  "name" TEXT NOT NULL,
  "role" role NOT NULL,
  "class_id" TEXT,
  FOREIGN KEY (class_id)
      REFERENCES class(id)
);

insert into "client" (id, username, password, name, "role", class_id) values ('1','datpham','123456','me','admin','1');

-- +migrate Down
DROP TABLE IF EXISTS "client";