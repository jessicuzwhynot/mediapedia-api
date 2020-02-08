package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Store struct {
		ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name string        `json:"name" bson:"name"`
		URL  string        `json:"url" bson:"url"`
	}

	Game struct {
		ID               bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name             string        `json:"name" bson:"name"`
		CoverArt         string        `json:"coverArt" bson:"coverArt"`
		Description      string        `json:"description" bson:"description"`
		MediapediaRating int           `json:"mediapediaRating" bson:"mediapediaRating"`
		Metacritic       int           `json:"metacritic" bson:"metacritic"`
		Publisher        string        `json:"publisher" bson:"publisher"`
		ReleaseDate      string        `json:"releaseDate" bson:"releaseDate"`
		Stores           []Store       `json:"stores" bson:"stores"`
	}
)
