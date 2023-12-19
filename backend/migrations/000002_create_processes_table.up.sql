BEGIN;

CREATE TABLE IF NOT EXISTS processes (
    id INTEGER NOT NULL PRIMARY KEY,
    user text,
    pid integer NOT NULL UNIQUE,
    cpuUsage integer,
    memoryUsage integer,
    residentMemorySize integer,
    virtualMemorySize integer,
    state text,
    totalTime text,
    cpuTime text,
    command text,
    priority text,
    createdAt timestamp,
    updatedAt timestamp
);

CREATE INDEX idx_user ON user (processes);
CREATE INDEX idx_state ON state (processes);

COMMIT;