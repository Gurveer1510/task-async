CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name TEXT,
    job_type TEXT,
    payload JSONB,
    run_at TIMESTAMP,
    status TEXT DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW()
);
    