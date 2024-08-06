CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TEXT NOT NULL DEFAULT now(),
    updated_at TEXT NOT NULL DEFAULT now(),
    deleted_at INT DEFAULT 0
);
