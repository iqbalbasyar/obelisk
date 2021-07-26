CREATE TABLE "students" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "joined_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "class" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "class_enrollment" (
  "id" bigserial PRIMARY KEY,
  "student_id" bigSerial NOT NULL,
  "class_id" bigSerial NOT NULL
);

CREATE TABLE "materials" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "class_materials" (
  "id" bigserial PRIMARY KEY,
  "material_id" bigserial NOT NULL,
  "class_id" bigserial NOT NULL
);

CREATE TABLE "mentors" (
  "id" bigserial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "joined_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "assignment_tickets" (
  "id" bigserial PRIMARY KEY,
  "material_id" bigserial NOT NULL,
  "student_id" bigserial NOT NULL,
  "mentor_id" bigserial,
  "score" int
);

ALTER TABLE "class_enrollment" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "class_enrollment" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "class_materials" ADD FOREIGN KEY ("material_id") REFERENCES "materials" ("id");

ALTER TABLE "class_materials" ADD FOREIGN KEY ("class_id") REFERENCES "class" ("id");

ALTER TABLE "assignment_tickets" ADD FOREIGN KEY ("material_id") REFERENCES "class_materials" ("id");

ALTER TABLE "assignment_tickets" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");

ALTER TABLE "assignment_tickets" ADD FOREIGN KEY ("mentor_id") REFERENCES "mentors" ("id");

CREATE INDEX ON "students" ("name");

CREATE INDEX ON "assignment_tickets" ("mentor_id");

CREATE INDEX ON "assignment_tickets" ("student_id");

CREATE INDEX ON "assignment_tickets" ("material_id");
