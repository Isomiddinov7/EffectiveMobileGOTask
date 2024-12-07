package models

type Song struct {
	Id          string   `json:"id"`
	GroupName   string   `json:"group_name"`
	SongName    string   `json:"song_name"`
	ReleaseDate string   `json:"release_date"`
	Lyrics      string   `json:"lyrics"`
	Link        string   `json:"link"`
	Genres      []string `json:"genres"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type CreateSong struct {
	GroupName   string   `json:"group_name"`
	SongName    string   `json:"song_name"`
	ReleaseDate string   `json:"release_date"`
	Lyrics      string   `json:"lyrics"`
	Link        string   `json:"link"`
	Genres      []string `json:"genres"`
}

type GetSongRequest struct {
	Limit  int64  `json:"limit"`
	Offset int64  `json:"offset"`
	Search string `json:"search"`
}

type GetSongResponse struct {
	Songs []*Song `json:"songs"`
}

type SongPrimaryKey struct {
	Id string `json:"id"`
}

type UpdateSong struct {
	Id          string   `json:"id"`
	GroupName   string   `json:"group_name"`
	SongName    string   `json:"song_name"`
	ReleaseDate string   `json:"release_date"`
	Lyrics      string   `json:"lyrics"`
	Link        string   `json:"link"`
	Genres      []string `json:"genres"`
}
