package search

import (
	"appengine"
	"appengine/search"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/search/save", saveIndex)
	http.HandleFunc("/search/search", searchBooks)
}

type Book struct {
	Title     string
	Author    string
	Price     float64
	CreatedAt time.Time
}

func saveIndex(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	books := make([]*Book, 0, 3)
	books = append(books, &Book{
		Title:     "Perfect Go",
		Author:    "Some Gopher",
		Price:     3000,
		CreatedAt: time.Now(),
	})
	books = append(books, &Book{
		Title:     "Go In Practice",
		Author:    "One Gopher",
		Price:     2500,
		CreatedAt: time.Now(),
	})
	books = append(books, &Book{
		Title:     "Let it go",
		Author:    "hogehoge",
		Price:     3000,
		CreatedAt: time.Now(),
	})

	// start searchput OMIT
	index, err := search.Open("Book") // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, book := range books {
		_, err := index.Put(c, fmt.Sprintf("book%d", i), book) // HL
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// end searchput OMIT

	fmt.Fprint(w, "success")
}

func searchBooks(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start search OMIT
	index, err := search.Open("Book") // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	books := make([]*Book, 0, 10)
	t := index.Search(c, "Gopher Price >= 3000", nil) // HL
	for {
		var book Book
		_, err := t.Next(&book)
		if err == search.Done {
			break
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, &book)
	}
	// end search OMIT

	je := json.NewEncoder(w)
	if err := je.Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
