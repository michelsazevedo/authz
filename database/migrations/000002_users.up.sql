-- users
CREATE TABLE IF NOT EXISTS users (
    id         uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    first_name VARCHAR(125) NOT NULL,
    last_name  VARCHAR(125) NOT NULL,
    email      VARCHAR(125) UNIQUE NOT NULL,
    password   VARCHAR(255) NOT NULL,
    locale     VARCHAR(125) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
    );

--
CREATE OR REPLACE TRIGGER set_timestamp
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
