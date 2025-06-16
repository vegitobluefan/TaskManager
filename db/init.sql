CREATE TABLE IF NOT EXISTS tasks (
  id UUID PRIMARY KEY,
  type TEXT NOT NULL,
  payload TEXT,
  status TEXT,
  result TEXT
);