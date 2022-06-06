-- +migrate Up
CREATE TABLE "course_at_class" (
  "id" TEXT PRIMARY KEY,
  class_id TEXT,
  course_id TEXT,
  FOREIGN KEY (class_id)
    REFERENCES class(id)
,
  FOREIGN KEY (course_id)
    REFERENCES course(id)
);


-- +migrate Down
DROP TABLE IF EXISTS "course_at_class";