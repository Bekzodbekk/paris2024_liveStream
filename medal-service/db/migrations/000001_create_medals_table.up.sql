CREATE TABLE IF NOT EXISTS medals (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    country_id VARCHAR(255) NOT NULL,
    type INT NOT NULL,
    event_id VARCHAR(255) NOT NULL,
    athlete_id VARCHAR(255) NOT NULL,
    created_at TEXT NOT NULL DEFAULT now(),
    updated_at TEXT NOT NULL DEFAULT now(),
    deleted_at INT DEFAULT 0
);
