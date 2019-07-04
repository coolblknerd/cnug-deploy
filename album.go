package main

// Albums holds a collection of albums
type Albums struct {
	Albums []*Album `json:"albums"`
}

// Album is representative of a musical body of work from a artist
type Album struct {
	Artist string  `json:"artist"`
	Genre  string  `json:"genre"`
	Title  string  `json:"title"`
	Year   float64 `json:"year"`
}

// NewAlbum creates a new instance of a album
func NewAlbum(artist string, genre string, title string, year float64) *Album {
	return &Album{
		Artist: artist,
		Genre:  genre,
		Title:  title,
		Year:   year,
	}
}
