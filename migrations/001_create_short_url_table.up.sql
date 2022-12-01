CREATE TABLE IF NOT EXISTS "short_url" (
  "id" serial NOT NULL,
  PRIMARY KEY ("id"),
  "code" character(255) NOT NULL,
  "original" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "lifetime" integer NULL
);
