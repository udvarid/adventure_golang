package main

import (
	"errors"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"

	myStory "github.com/udvarid/adventure/story"
)

var validPath = regexp.MustCompile("^/(story)/([a-zA-Z0-9-]+)$")

var fullStory myStory.Story

var templates = template.Must(template.ParseFiles("story.html"))

func makeHandler(fn func(http.ResponseWriter, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, m[2])
	}
}

func loadChapter(title string) (*myStory.Chapter, error) {
	currentChapter, ok := fullStory[title]
	if !ok {
		return nil, errors.New("Not found this chapter")
	}
	return &currentChapter, nil
}

func storyHandler(w http.ResponseWriter, title string) {
	p, err := loadChapter(title)
	if err != nil {
		return
	}
	renderTemplate(w, "story", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *myStory.Chapter) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	filename := flag.String("file", "gopher.json", "the Json file contains the story")
	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := myStory.JsonStory(f)
	if err != nil {
		panic(err)
	}
	fullStory = story

	http.HandleFunc("/", makeHandler(storyHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
