CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR NOT NULL UNIQUE,
  "password" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE "rooms" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR UNIQUE NOT NULL
);

CREATE TABLE "messages" (
  "id" BIGSERIAL PRIMARY KEY,
  "room_id" BIGINT,
  "user_id" BIGINT,
  "content" VARCHAR NOT NULL,
  "created_at" TIMESTAMPTZ DEFAULT NOW()
);

-- Add foreign key constraints with CASCADE or SET NULL options
ALTER TABLE "messages" 
  ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE CASCADE;

ALTER TABLE "messages" 
  ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
