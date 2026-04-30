CREATE TABLE sagas (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    title       TEXT NOT NULL,
    description TEXT,
    image_url   TEXT;
);

CREATE TABLE songs (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    saga_id     INTEGER NOT NULL,
    title       TEXT NOT NULL,
    description TEXT,
    singers     TEXT,
    image_url   TEXT,
    FOREIGN KEY (saga_id) REFERENCES sagas(id)
);