CREATE TABLE IF NOT EXISTS "profiles" (
    id serial PRIMARY KEY,
    user_id INTEGER,
    profile_pictures TEXT,
    bio TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    CONSTRAINT fk_user_id_profile
        FOREIGN KEY (user_id)
        REFERENCES "users"(id)
);
CREATE TRIGGER set_updated_at BEFORE UPDATE ON "profiles" FOR EACH ROW EXECUTE
PROCEDURE trigger_update_timestamp();