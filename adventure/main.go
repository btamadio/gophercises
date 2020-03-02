package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type StoryArc struct{
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []struct{
		Text string `json:"text"`
		ArcName string `json:"arc"`
	} `json:"options"`
}

func main(){
	var book map[string]StoryArc
	jsonData, err := ioutil.ReadFile("gopher.json")
	if err != nil{
		log.Fatal("failed to open json file")
	}
	if err := json.Unmarshal(jsonData, &book); err != nil{
		log.Fatal("failed to unmarshal json")
	}
	fmt.Printf("%+v\n", book)
}