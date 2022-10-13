CREATE TABLE "public"."animal" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "species" text NOT NULL, PRIMARY KEY ("id") , UNIQUE ("species"));COMMENT ON TABLE "public"."animal" IS E'animal';
CREATE EXTENSION IF NOT EXISTS pgcrypto;
