package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Albums is a colletion of albums
var Albums []Album

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "homePage")
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getAlbum")
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getAlbums")
	json.NewEncoder(w).Encode(Albums)
}

func removeAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "removeAlbum")
}

func addAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "addAlbum")
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "updateAlbum")
}

func handleRoutes() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", homePage).Methods("GET")
	s.HandleFunc("/albums", getAlbum).Methods("GET")
	s.HandleFunc("/albums", getAlbums).Methods("GET")
	s.HandleFunc("/albums", removeAlbum).Methods("DELETE")
	s.HandleFunc("/albums", addAlbum).Methods("POST")
	s.HandleFunc("/albums", updateAlbum).Methods("PUT")
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", s))
}

func main() {
	albums := []Album{
		Album{Artist: "Larry June", Genre: "Hip Hop", Title: "Mr. Midnight", Year: 2019},
		Album{Artist: "Jaden", Genre: "Hip Hop", Title: "Erys", Year: 2019},
		Album{Artist: "Freddie Gibbs", Genre: "Hip Hop", Title: "Bandana", Year: 2019},
		Album{Artist: "Benny the Butcher", Genre: "Hip Hop", Title: "The Plug I Met", Year: 2019},
	}

	handleRoutes()
}
