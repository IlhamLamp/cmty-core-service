-- 001_create_members.up.sql
CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    types VARCHAR(10),
    core_id INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    profile_id VARCHAR(50) NOT NULL,
    experience VARCHAR(20),
    role_id VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
);
