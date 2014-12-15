package memcache

import (
	"appengine"
	"appengine/memcache"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/memcache/set", set)
	http.HandleFunc("/memcache/get", get)
}

type Book struct {
	Title     string
	Author    string
	Price     int
	CreatedAt time.Time
}

func set(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	book := Book{
		"Perfect Go",
		"Some Gopher",
		1000,
		time.Now(),
	}

	item := memcache.Item{
		Key:    "book1",
		Object: &book,
	}
	if err := memcache.Gob.Set(c, &item); err != nil {
		fmt.Fprintf(w, "failure: %s", err)
	} else {
		w.Write([]byte("success"))
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var book Book
	_, err := memcache.Gob.Get(c, "book1", &book)
	if err != nil {
		w.Write([]byte("not found"))
		return
	}

	je := json.NewEncoder(w)
	if err := je.Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
