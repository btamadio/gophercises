package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type StoryArc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []struct {
		Text    string `json:"text"`
		ArcName string `json:"arc"`
	} `json:"options"`
}

func arcHandler(s StoryArc) http.HandlerFunc {

	// TODO: Use HTML templates to display the title, paragraphs, and options, with links to other pages

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("%+v", s))
	}
}

func main() {
	var book map[string]StoryArc
	jsonData, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		log.Fatal("failed to open json file")
	}
	if err := json.Unmarshal(jsonData, &book); err != nil {
		log.Fatal("failed to unmarshal json")
	}

	for arcName := range book {
		http.HandleFunc("/"+arcName, arcHandler(book[arcName]))
	}
	log.Fatal(http.ListenAndServe(":8080", nil))

}
