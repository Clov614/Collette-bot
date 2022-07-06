package traceMoe

type ResTracemoe struct {
	FrameCount int    `json:"frameCount"`
	Error      string `json:"error"`
	Result     []struct {
		Anilist struct {
			ID    int `json:"id"`
			IDMal int `json:"idMal"`
			Title struct {
				Native  string `json:"native"`
				Romaji  string `json:"romaji"`
				English string `json:"english"`
			} `json:"title"`
			Synonyms []string `json:"synonyms"`
			IsAdult  bool     `json:"isAdult"`
		} `json:"anilist"`
		Filename   string  `json:"filename"`
		Episode    int     `json:"episode"`
		From       float64 `json:"from"`
		To         float64 `json:"to"`
		Similarity float64 `json:"similarity"`
		Video      string  `json:"video"`
		Image      string  `json:"image"`
	} `json:"result"`
}
