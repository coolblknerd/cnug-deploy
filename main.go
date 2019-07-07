package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Albums is a colletion of albums
var Albums []*Album

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
	vars := mux.Vars(r)
	id := vars["id"]

	for i, album := range Albums {
		if album.ID == id {
			Albums = append(Albums[:i], Albums[i+1:]...)
		}
	}
}

func createAlbum(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var album *Album

	json.Unmarshal(data, &album)

	Albums = append(Albums, album)

	json.NewEncoder(w).Encode(album)
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "updateAlbum")

	data, _ := ioutil.ReadAll(r.Body)
	var album *Album
	json.Unmarshal(data, &album)

	vars := mux.Vars(r)
	id := vars["id"]

	for _, project := range Albums {
		if project.ID == id {
			project.Artist = album.Artist
			project.Title = album.Title
			project.Genre = album.Genre
		}
	}

	json.NewEncoder(w).Encode(album)
}

func handleRoutes() {
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", homePage)
	s.HandleFunc("/albums", readAlbums)
	s.HandleFunc("/album", createAlbum).Methods("POST")
	s.HandleFunc("/album/{id}", deleteAlbum).Methods("DELETE")
	s.HandleFunc("/albums/{id}", updateAlbum).Methods("PUT")
	s.HandleFunc("/album/{id}", readAlbum)
	log.Fatal(http.ListenAndServe(":8080", s))
}

func main() {
	Albums = []*Album{
		&Album{ID: "1", Artist: "Larry June", Genre: "Hip Hop", Title: "Mr. Midnight", Year: 2019},
		&Album{ID: "2", Artist: "Jaden", Genre: "Hip Hop", Title: "Erys", Year: 2019},
		&Album{ID: "3", Artist: "Freddie Gibbs", Genre: "Hip Hop", Title: "Bandana", Year: 2019},
		&Album{ID: "4", Artist: "Benny the Butcher", Genre: "Hip Hop", Title: "The Plug I Met", Year: 2019},
	}

	handleRoutes()
}
