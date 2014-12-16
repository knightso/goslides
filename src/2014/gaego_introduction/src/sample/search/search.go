package search

import (
	"appengine"
	"appengine/delay"
	"appengine/taskqueue"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func init() {
	http.HandleFunc("/search/saveIndx", saveIndex)
	http.HandleFunc("/search/search", search)
}

type Book struct {
	Title     string
	Author    string
	Price     int
	CreatedAt time.Time
}

func saveIndx(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	index, err := Search.Open("Book")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	book := &Book{
		Title:     "Perfect Go",
		Author:    "Some Gopher",
		Price:     3000,
		CreatedAt: time.Now(),
	}

	_, err := index.Put(c, "book1", book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "success")
}

func search(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	index, err := Search.Open("Book")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	books := make([]*Book, 0, 10)
	t := index.Search(c, "Gopher", nil)
	for {
		var book Book
		id, err := t.Next(&doc)
		if err == search.Done {
			break
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	je := json.NewEncoder(w)
	if err := je.Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

