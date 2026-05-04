CREATE TABLE IF NOT EXISTS activities (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    child_profile_id BIGINT NOT NULL,
    type TEXT NOT NULL,              -- reading / art / sport
    activity_id TEXT NOT NULL,       -- book-id / game-id
    confirmed_by_parent BOOLEAN DEFAULT false,
    created_at TIMESTAMPTZ DEFAULT NOW()
);