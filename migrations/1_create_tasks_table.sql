CREATE TYPE task_status AS ENUM ('pending','running','done','failed','canceled');

CREATE TABLE IF NOT EXISTS tasks (
  id              BIGSERIAL PRIMARY KEY,
  name            TEXT NOT NULL,
  job_type        TEXT NOT NULL,
  payload         JSONB NOT NULL,
  run_at          TIMESTAMPTZ NOT NULL,         -- store UTC
  status          task_status NOT NULL DEFAULT 'pending',
  attempts        INT NOT NULL DEFAULT 0,
  max_attempts    INT NOT NULL DEFAULT 10,
  last_error      TEXT,
  dedup_key       TEXT,                          -- optional idempotency
  visibility_deadline TIMESTAMPTZ,               -- for heartbeat/timeouts
  created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- claimable jobs index
CREATE INDEX IF NOT EXISTS idx_tasks_claim ON tasks (status, run_at);

-- dedup if you want "at-most-once" per logical key
CREATE UNIQUE INDEX IF NOT EXISTS uniq_tasks_dedup ON tasks (dedup_key) WHERE dedup_key IS NOT NULL;