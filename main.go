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

func readAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getAlbum")
	vars := mux.Vars(r)
	key := vars["id"]

	for _, album := range Albums {
		if album.ID == key {
			json.NewEncoder(w).Encode(album)
		}
	}
}

func readAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "getAlbums\n")
	json.NewEncoder(w).Encode(Albums)
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "removeAlbum")
}

func createAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "addAlbum\n")
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "updateAlbum")
}

func handleRoutes() {
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", homePage)
	s.HandleFunc("/albums/{id}", readAlbum)
	s.HandleFunc("/albums", readAlbums)
	s.HandleFunc("/albums/{id}", deleteAlbum).Methods("DELETE")
	s.HandleFunc("/album", createAlbum).Methods("POST")
	s.HandleFunc("/albums/{id}", updateAlbum).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", s))
}

func main() {
	Albums = []Album{
		Album{ID: "1", Artist: "Larry June", Genre: "Hip Hop", Title: "Mr. Midnight", Year: 2019},
		Album{ID: "2", Artist: "Jaden", Genre: "Hip Hop", Title: "Erys", Year: 2019},
		Album{ID: "3", Artist: "Freddie Gibbs", Genre: "Hip Hop", Title: "Bandana", Year: 2019},
		Album{ID: "4", Artist: "Benny the Butcher", Genre: "Hip Hop", Title: "The Plug I Met", Year: 2019},
	}

	handleRoutes()
}
