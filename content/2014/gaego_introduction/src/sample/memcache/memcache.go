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
	http.HandleFunc("/memcache/setGob", setGob)
	http.HandleFunc("/memcache/getGob", getGob)
}

type Book struct {
	Title     string
	Author    string
	Price     int
	CreatedAt time.Time
}

func set(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start memset OMIT
	item := memcache.Item{ // HL
		Key:   "test1",
		Value: []byte("hello memcache"),
	}
	if err := memcache.Set(c, &item); err != nil { // HL
		fmt.Fprintf(w, "failure: %s", err)
	} else {
		fmt.Fprint(w, "success")
	}
	// end memset OMIT
}

func get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start memget OMIT
	item, err := memcache.Get(c, "test1") // HL
	if err != nil {
		w.Write([]byte("not found"))
		return
	}
	// end memget OMIT

	je := json.NewEncoder(w)
	if err := je.Encode(string(item.Value)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func setGob(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start memgobset OMIT
	book := Book{
		"Perfect Go",
		"Some Gopher",
		1000,
		time.Now(),
	}

	item := memcache.Item{ // HL
		Key:    "book1",
		Object: &book,
	}
	if err := memcache.Gob.Set(c, &item); err != nil { // HL
		fmt.Fprintf(w, "failure: %s", err)
	} else {
		w.Write([]byte("success"))
	}
	// end memgobset OMIT
}

func getGob(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start memgobget OMIT
	var book Book
	_, err := memcache.Gob.Get(c, "book1", &book) // HL
	if err != nil {
		w.Write([]byte("not found"))
		return
	}
	// end memgobget OMIT

	je := json.NewEncoder(w)
	if err := je.Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
