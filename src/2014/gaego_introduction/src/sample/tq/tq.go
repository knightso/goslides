package tq
/*
import (
	"appengine"
	"appengine/memcache"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/memcache/push", push)
	http.HandleFunc("/memcache/add4pull", get)
	http.HandleFunc("/memcache/handler", handler)
}

func push(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	t := taskqueue.NewPOSTTask("/memcache/worker", map[string][]string{
		"Title": {"Perfect Go"},
		"Author": {"Some Gopher"},
		"Price": {"1000"},
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
*/
