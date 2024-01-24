CREATE TABLE IF NOT EXISTS "books" (
    id serial PRIMARY KEY,
    user_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    title varchar(255) NOT NULL,
    synopsis text NOT NULL,
    cover text NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user_id_book
        FOREIGN KEY(user_id)
        REFERENCES "users"(id),
    CONSTRAINT fk_category_id_book
        FOREIGN KEY(category_id)
        REFERENCES "categories"(id)
);
CREATE TRIGGER set_updated_at BEFORE UPDATE ON "books" FOR EACH ROW EXECUTE
PROCEDURE trigger_update_timestamp();