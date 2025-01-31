CREATE TABLE playlist (
    id serial NOT NULL PRIMARY KEY,
    title text NOT NULL,
    created_at 		timestamp with time zone 	DEFAULT now() NOT NULL,
    updated_at 		timestamp with time zone
);

CREATE TABLE song (
    id serial NOT NULL PRIMARY KEY,
    title text NOT NULL,
    artist text NOT NULL,
    duration int NOT NULL
);


CREATE TABLE playlist_songs (
    id serial NOT NULL PRIMARY KEY,
    playlist_id int,
    song_id int,
    FOREIGN KEY(playlist_id) REFERENCES playlist(id),
    FOREIGN KEY(song_id) REFERENCES song(id)
);