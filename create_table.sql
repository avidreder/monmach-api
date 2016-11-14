CREATE TABLE playlists(
       id serial PRIMARY KEY NOT NULL,
       name varchar (50),
       user_id int,
       tracks int[],
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL
);

CREATE TABLE users(
       id serial PRIMARY KEY NOT NULL,
       name varchar (50),
       email varchar (100),
       avatar_url varchar (200),
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL,
       track_blacklist int[],
       track_whitelist int[],
       listened_tracks int[]
);

CREATE TABLE tracks(
       id serial PRIMARY KEY NOT NULL,
       name varchar (50),
       artists varchar(50)[][]
       image_url varchar (200),
       spotify_id varchar(20)[]
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL,
       features real[]
);

CREATE TABLE genres(
       id serial PRIMARY KEY NOT NULL,
       user_id int,
       name varchar (50),
       description varchar (200),
       seed_artists int[],
       seed_tracks int[],
       avatar_url varchar (200),
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL,
       track_blacklist int[],
       track_whitelist int[]
);

CREATE OR REPLACE FUNCTION update_updated_preserve_created()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated = now();
    NEW.created = OLD.created;
    RETURN NEW;	
END;
$$ language 'plpgsql';

CREATE TRIGGER update_timestamp BEFORE UPDATE ON playlists FOR EACH ROW EXECUTE PROCEDURE update_updated_preserve_created();
