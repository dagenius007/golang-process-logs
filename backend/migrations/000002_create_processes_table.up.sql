
CREATE TABLE IF NOT EXISTS processes (
    id integer not null primary key autoincrement,
    user text,
    pid integer not null UNIQUE,
    cpu_usage decimal(10,2),
    memory_usage decimal(10,2),
    resident_memory_size integer,
    virtual_memory_size integer,
    state text,
    total_time text,
    cpu_time text,
    command text,
    priority text,
    created_at timestamp,
    updated_at timestamp
);

CREATE INDEX idx_user ON processes (user);
CREATE INDEX idx_state ON processes (state);
