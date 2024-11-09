CREATE TABLE vocabulary (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    word VARCHAR(100) NOT NULL,
    translation VARCHAR(100) NOT NULL,
    difficulty INTEGER NOT NULL
);
