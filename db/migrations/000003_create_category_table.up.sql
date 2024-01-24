CREATE TABLE IF NOT EXISTS "categories" (
    id serial PRIMARY KEY,
    name varchar(255) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_updated_at BEFORE UPDATE ON "categories" FOR EACH ROW EXECUTE
PROCEDURE trigger_update_timestamp();

INSERT INTO "categories"(name)
VALUES('horror'),('romance'),('fiction');