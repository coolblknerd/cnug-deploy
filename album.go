package main

// Album is representative of a musical body of work from a artist
type Album struct {
	ID     string  `json:"id"`
	Artist string  `json:"artist"`
	Genre  string  `json:"genre"`
	Title  string  `json:"title"`
	Year   float64 `json:"year"`
}
