CREATE TABLE IF NOT EXISTS processes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    pid INTEGER NOT NULL,
    log_file TEXT,
    status TEXT,
    command TEXT
);

