DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
    id          SERIAL PRIMARY KEY,
	title       TEXT NOT NULL UNIQUE,
	description TEXT NOT NULL,
	date        TEXT NOT NULL,
	link        TEXT NOT NULL
);
