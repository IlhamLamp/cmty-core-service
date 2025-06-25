-- 002_create_projects.up.sql
CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    logo TEXT,
    owner VARCHAR (50),
    title VARCHAR(100),
    company VARCHAR(100),
    start_date DATE,
    end_date DATE,
    types VARCHAR(10),
    duration VARCHAR(10),
    participation VARCHAR(10),
    address JSONB,
    approval VARCHAR(5),
    description TEXT,
    tags JSONB,
    salary NUMERIC(15, 2),
    priority VARCHAR(10),
    is_private BOOLEAN,
    is_closed BOOLEAN,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
);
