package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	myStory "udvari/adventure"
)

func main() {
	filename := flag.String("file", "gopher.json", "the Json file contains the story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)
	var story myStory.Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("+v\n", story)

}
