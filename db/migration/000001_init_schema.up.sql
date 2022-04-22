CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);