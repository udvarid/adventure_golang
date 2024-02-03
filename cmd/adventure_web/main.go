package main

import (
	"flag"
	"fmt"
	"os"

	myStory "github.com/adventure"
)

func main() {
	filename := flag.String("file", "gopher.json", "the Json file contains the story")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *filename)

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := myStory.JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}
