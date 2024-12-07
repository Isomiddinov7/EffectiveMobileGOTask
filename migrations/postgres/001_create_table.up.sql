CREATE TABLE "songs"(
	"id" UUID NOT NULL PRIMARY KEY,
    "group_name" VARCHAR NOT NULL,
    "song_name" VARCHAR NOT NULL,
    "release_date" VARCHAR NOT NULL,
    "lyrics" VARCHAR NOT NULL,
    "link" VARCHAR NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

CREATE  TABLE "song_genres"(
    "song_id" UUID NOT NULL REFERENCES "songs"("id"),
    "genre" VARCHAR NOT NULL
);