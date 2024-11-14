-- +goose Up
CREATE TYPE role AS ENUM('manager','superadmin');
CREATE TYPE status AS ENUM('pending','accepted','rejected');
CREATE TABLE "user" (
  id bigserial PRIMARY KEY,
  email VARCHAR NOT NULL UNIQUE,
  role role,
  status status,
  created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "user";
DROP TYPE IF EXISTS role;
DROP TYPE IF EXISTS status;
