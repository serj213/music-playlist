CREATE TABLE playlist (
    id serial PRIMARY KEY NOT NULL,
    title text NOT NULL,
    artist text NOT NULL,
    duration text NOT NULL,
    position int NOT NULL
);