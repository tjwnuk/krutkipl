CREATE TABLE IF NOT EXISTS Url (
    ID INTEGER PRIMARY KEY NOT NULL,
    CreatedAt TEXT,
    UpdatedAt TEXT,
    DeletedAt TEXT,
    OriginalURL TEXT,
    Token TEXT,
    ShortenedURL TEXT
)