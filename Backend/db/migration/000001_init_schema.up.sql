CREATE TABLE "recruiter" (
  "id" bigint PRIMARY KEY,
  "name" varchar,
  "company" varchar
);

CREATE TABLE "job" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL,
  "skill" varchar NOT NULL,
  "location" varchar NOT NULL,
  "experince" varchar,
  "compensation" varchar,
  "postby" bigint NOT NULL
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "softdelete" bool NOT NULL DEFAULT false
);

CREATE TABLE "jobseeker" (
  "id" bigint PRIMARY KEY NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "firstname" varchar NOT NULL,
  "lastname" varchar,
  "photo" blob,
  "resume" blob NOT NULL,
  "address" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "pincode" varchar NOT NULL
);

CREATE TABLE "experience" (
  "id" bigint,
  "title" varchar,
  "company" varchar,
  "startddate" date,
  "enddate" date,
  "skills" varchar
);

CREATE TABLE "education" (
  "id" bigint,
  "school" varchar,
  "course" varchar,
  "startdate" date,
  "enddate" date
);

CREATE TABLE "appliedjob" (
  "userid" bigint PRIMARY KEY NOT NULL,
  "jobid" bigint NOT NULL,
  "recruiterid" bigint NOT NULL,
  "current_status" varchar NOT NULL
);

CREATE TABLE "recommendedjob" (
  "userid" bigint PRIMARY KEY NOT NULL,
  "jobid" bigint NOT NULL
);

ALTER TABLE "recruiter" ADD FOREIGN KEY ("id") REFERENCES "user" ("id") on DELETE CASCADE;

ALTER TABLE "jobseeker" ADD FOREIGN KEY ("id") REFERENCES "user" ("id") on DELETE CASCADE;

ALTER TABLE "experience" ADD FOREIGN KEY ("id") REFERENCES "user" ("id") on DELETE CASCADE;

ALTER TABLE "education" ADD FOREIGN KEY ("id") REFERENCES "user" ("id") on DELETE CASCADE;

ALTER TABLE "job" ADD FOREIGN KEY ("postby") REFERENCES "recruiter" ("id") on DELETE CASCADE;

ALTER TABLE "appliedjob" ADD FOREIGN KEY ("userid") REFERENCES "jobseeker" ("id") on DELETE CASCADE;

ALTER TABLE "appliedjob" ADD FOREIGN KEY ("jobid") REFERENCES "job" ("id") on DELETE CASCADE;

ALTER TABLE "recommendedjob" ADD FOREIGN KEY ("userid") REFERENCES "user" ("id") on DELETE CASCADE;

ALTER TABLE "recommendedjob" ADD FOREIGN KEY ("jobid") REFERENCES "job" ("id") on DELETE CASCADE;

ALTER TABLE "appliedjob" ADD FOREIGN KEY ("recruiterid") REFERENCES "recruiter" ("id") on DELETE CASCADE;
