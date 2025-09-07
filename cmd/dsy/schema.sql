CREATE TABLE IF NOT EXISTS processes (
    title TEXT PRIMARY KEY,
    pid INTEGER NOT NULL,
    log_file TEXT,
    status TEXT,
    command TEXT
);

