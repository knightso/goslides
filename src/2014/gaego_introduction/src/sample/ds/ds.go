package ds

// start 1 OMIT
import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
// end 1 OMIT

func init() {
	http.HandleFunc("/ds/put", put)
	http.HandleFunc("/ds/putmulti", putMulti)
	http.HandleFunc("/ds/get", get)
	http.HandleFunc("/ds/getmulti", getMulti)
	http.HandleFunc("/ds/query", query)
	http.HandleFunc("/ds/query2", query2)
	http.HandleFunc("/ds/tx", tx)
}

type Book struct {
	Title     string
	Author    string
	Price     int
	CreatedAt time.Time
}

func put(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	book := Book{
		"Perfect Go",
		"Some Gopher",
		1000,
		time.Now(),
	}

	key := datastore.NewKey(c, "Book", "book1", 0, nil) // HL
	//key := datastore.NewIncompleteKey(c, "Book", nil) // HL
	key, err := datastore.Put(c, key, &book) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("success"))
}

func putMulti(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	keys := make([]*datastore.Key, 10)
	for i, _ := range keys {
		keys[i] = datastore.NewKey(c, "Book", "", int64(i + 1), nil) // HL
	}

	books := make([]Book, 10)
	for i, _ := range books {
		number := i + 1
		books[i] = Book{
			fmt.Sprintf("book-%d", number),
			fmt.Sprintf("author-%d", number % 2),
			number * 100,
			time.Now(),
		}
	}
	_, err := datastore.PutMulti(c, keys, books) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("success"))
}

func get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	key := datastore.NewKey(c, "Book", "book1", 0, nil) // HL

	var book Book
	err := datastore.Get(c, key, &book) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	je := json.NewEncoder(w)
	if err := je.Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getMulti(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	keys := make([]*datastore.Key, 10)
	for i, _ := range keys {
		keys[i] = datastore.NewKey(c, "Book", "", int64(i + 1), nil) // HL
	}

	books := make([]Book, 10)
	err := datastore.GetMulti(c, keys, books) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	je := json.NewEncoder(w)
	if err := je.Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func query(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Book").Filter("Author=", "author-1").Order("-CreatedAt").Offset(2).Limit(5)

	var books []Book
	keys, err := q.GetAll(c, &books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Debugf("#v", keys)

	je := json.NewEncoder(w)
	if err := je.Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func query2(w http.ResponseWriter, r *http.Request) {
	pCursor := r.FormValue("cursor")

	c := appengine.NewContext(r)

	q := datastore.NewQuery("Book").Filter("Author=", "author-1").Order("-CreatedAt")

	if pCursor != "" {
		cursor, err := datastore.DecodeCursor(pCursor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		q.Start(cursor)
	}

	var books []Book

	t := q.Run(c)
	for i := 0; i < 10; i++ {
		var book Book
		key, err := t.Next(&book)
		if err == datastore.Done {
			break
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		c.Debugf("#v", key)

		books = append(books, book)
	}

	response := struct {
		Cursor string
		Books  []Book
	}{
		Books: books,
	}

	response.Books = books

	if cursor, err := t.Cursor(); err == nil {
		response.Cursor = cursor.String()
	}

	je := json.NewEncoder(w)
	if err := je.Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func tx(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	key := datastore.NewKey(c, "Book", "book1", 0, nil) // HL

	if err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		var book Book
		if err := datastore.Get(c, key, &book); err != nil { // HL
			return err
		}

		book.Price += 200

		if _, err := datastore.Put(c, key, &book); err != nil {
			return err
		}

		return nil
	}, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("success"))
}
