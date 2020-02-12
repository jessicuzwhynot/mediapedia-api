-- CREATE DATABASE mediapedia;

--
-- Table structure for table game
--

CREATE TABLE game (
  id serial CONSTRAINT game_pk PRIMARY KEY,
  title VARCHAR (255) NOT NULL,
  description VARCHAR (255) NOT NULL,
  cover_art VARCHAR (255) NOT NULL,
  mediapedia_rating INTEGER NOT NULL,
  metacritic_rating INTEGER NOT NULL,
  publisher VARCHAR (255) NOT NULL,
  release_date DATE NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at DATE DEFAULT CURRENT_TIMESTAMP
);


-- Game Seed for testing, TBD remove after AddGame() implemented
INSERT INTO game (title, cover_art, description, mediapedia_rating, metacritic_rating, publisher, release_date) 
VALUES ('Bloodborne', 'https://gamecoverart.com', 'Lorem Ipsum', 96, 95, 'FromSoftware', '2019-12-31');
