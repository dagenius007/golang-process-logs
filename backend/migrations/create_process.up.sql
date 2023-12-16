CREATE TABLE IF NOT EXISTS processes (
    id INTEGER NOT NULL PRIMARY KEY,
    user text,
    pid integer NOT NULL UNIQUE,
    cpuUsage integer,
    memoryPercentageUsage integer,
    virtualMemorySize integer,
    residentMemorySize integer,
    tty text,
    state text,
    started text
    totalTime text,
    command text,
    createdAt timestamp,
    updatedAt timestamp
);