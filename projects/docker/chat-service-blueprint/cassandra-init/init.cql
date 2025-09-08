CREATE KEYSPACE IF NOT EXISTS chat
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

CREATE TABLE IF NOT EXISTS chat.messages (
    id UUID PRIMARY KEY,
    room_id TEXT,
    sender TEXT,
    content TEXT,
    created_at TIMESTAMP
);
