CREATE TABLE playlists(
       id serial PRIMARY KEY NOT NULL,
       name varchar (50),
       user_id int,
       tracks int[] default '{}',
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL
);

CREATE TABLE queues(
       id serial PRIMARY KEY NOT NULL,
       user_id int NOT NULL,
       name varchar (40) default 'Untitled Queue' NOT NULL,
       max_size int default 10 NOT NULL,
       track_queue int[] default '{}',
       seed_artists int[] default '{}',
       seed_tracks int[] default '{}',
       listened_tracks int[] default '{}',
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL
);

CREATE TABLE users(
       id serial PRIMARY KEY NOT NULL,
       name varchar (50) NOT NULL,
       email varchar (100) NOT NULL UNIQUE,
       avatar_url varchar (200),
       spotify_token varchar (40),
       spotify_refresh_token varchar (40),
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL,
       track_blacklist int[] default '{}'
);

CREATE TABLE tracks(
       id serial PRIMARY KEY NOT NULL,
       name varchar (50),
       artists text[] default '{}',
       image_url varchar (200),
       spotify_id varchar(20),
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL,
       features real[] default '{}'
);

CREATE TABLE genres(
       id serial PRIMARY KEY NOT NULL,
       user_id int NOT NULL,
       queue_id int NOT NULL,
       name varchar (50),
       description varchar (200),
       seed_artists int[] default '{}',
       seed_tracks int[] default '{}',
       avatar_url varchar (200),
       created timestamptz default current_timestamp NOT NULL,
       updated timestamptz default current_timestamp NOT NULL,
       track_blacklist int[] default '{}',
       track_whitelist int[] default '{}' 
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

CREATE TRIGGER update_timestamp BEFORE UPDATE ON genres FOR EACH ROW EXECUTE PROCEDURE update_updated_preserve_created();

CREATE TRIGGER update_timestamp BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_updated_preserve_created();

CREATE TRIGGER update_timestamp BEFORE UPDATE ON tracks FOR EACH ROW EXECUTE PROCEDURE update_updated_preserve_created();

CREATE TRIGGER update_timestamp BEFORE UPDATE ON queues FOR EACH ROW EXECUTE PROCEDURE update_updated_preserve_created();
