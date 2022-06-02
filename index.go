package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Albums struct {
	StudioAlbums []Album `json:"studioAlbums"`
	LiveAlbums   []Album `json:"liveAlbums"`
	WeirdStuff   []Album `json:"weirdStuff"`
}

type Album struct {
	ReleaseDate string `json:"releaseDate"`
	Name        string `json:"name"`
	Art         string `json:"art"`
	Player      string `json:"player"`
	Label       string `json:"label"`
}

var AllAlbums Albums

func returnAllAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllAlbums")
	json.NewEncoder(w).Encode((AllAlbums))
}

func returnStudioAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnStudioAlbums")
	json.NewEncoder(w).Encode((AllAlbums.StudioAlbums))
}

func returnLiveAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnLiveAlbums")
	json.NewEncoder(w).Encode((AllAlbums.LiveAlbums))
}

func returnWeirdStuff(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnWeirdStuff")
	json.NewEncoder(w).Encode((AllAlbums.WeirdStuff))
}

func handleRequests() {
	port := os.Getenv("PORT")

	fmt.Println(port)

	http.HandleFunc("/albums", returnAllAlbums)
	http.HandleFunc("/albums/studio", returnStudioAlbums)
	http.HandleFunc("/albums/live", returnLiveAlbums)
	http.HandleFunc("/albums/weird", returnWeirdStuff)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func main() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	jsonFile, err := os.Open("album-data.json")

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll((jsonFile))

	json.Unmarshal(byteValue, &AllAlbums)

	println("API Running")

	handleRequests()
}
