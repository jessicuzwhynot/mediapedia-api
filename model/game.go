package model

type (
	// Store contains info for accessing a store
	Store struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	// Game contains game info
	Game struct {
		ID               string `json:"id"`
		Title            string `json:"title"`
		CoverArt         string `json:"coverArt"`
		Description      string `json:"description"`
		MediapediaRating int    `json:"mediapediaRating"`
		MetacriticRating int    `json:"metacritic"`
		Publisher        string `json:"publisher"`
		ReleaseDate      string `json:"releaseDate"`
		// Stores           []Store       `json:"stores"`
	}
)
