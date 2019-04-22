package ds

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/ds/put", put)
	http.HandleFunc("/ds/putmulti", putMulti)
	http.HandleFunc("/ds/get", get)
	http.HandleFunc("/ds/getmulti", getMulti)
	http.HandleFunc("/ds/query", query)
	http.HandleFunc("/ds/query2", query2)
	http.HandleFunc("/ds/tx", tx)
}

// start put1 OMIT
type Book struct {
	Title     string
	Author    string
	Price     int
	CreatedAt time.Time
}

// end put1 OMIT

func put(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start put2 OMIT
	book := Book{
		"Perfect Go",
		"Some Gopher",
		1000,
		time.Now(),
	}
	// end put2 OMIT

	// start put3 OMIT
	key := datastore.NewKey(c, "Book", "book1", 0, nil) // HL
	//key := datastore.NewIncompleteKey(c, "Book", nil) // 自動ID付与
	key, err := datastore.Put(c, key, &book) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end put3 OMIT

	w.Write([]byte("success"))
}

func putMulti(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start putMulti1 OMIT
	keys := make([]*datastore.Key, 10)
	for i, _ := range keys {
		keys[i] = datastore.NewKey(c, "Book", "", int64(i+1), nil) // HL
	}

	books := make([]Book, 10)
	for i, _ := range books {
		number := i + 1
		books[i] = Book{
			fmt.Sprintf("book-%d", number),
			fmt.Sprintf("author-%d", number%2),
			number * 100,
			time.Now(),
		}
	}
	_, err := datastore.PutMulti(c, keys, books) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end putMulti1 OMIT

	w.Write([]byte("success"))
}

func get(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start get1 OMIT
	key := datastore.NewKey(c, "Book", "book1", 0, nil) // HL

	var book Book
	err := datastore.Get(c, key, &book) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end get1 OMIT

	je := json.NewEncoder(w)
	if err := je.Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getMulti(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start getMulti1 OMIT
	keys := make([]*datastore.Key, 10)
	for i, _ := range keys {
		keys[i] = datastore.NewKey(c, "Book", "", int64(i+1), nil) // HL
	}

	books := make([]Book, 10)
	err := datastore.GetMulti(c, keys, books) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end getMulti1 OMIT

	je := json.NewEncoder(w)
	if err := je.Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func query(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start query1 OMIT
	q := datastore.NewQuery("Book").Filter("Author=", "author-1").Order("-CreatedAt").Offset(2).Limit(5) // HL

	var books []Book
	keys, err := q.GetAll(c, &books) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end query1 OMIT

	c.Debugf("#v", keys)

	je := json.NewEncoder(w)
	if err := je.Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func query2(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	// start cursor1 OMIT
	q := datastore.NewQuery("Book").Filter("Author=", "author-1").Order("-CreatedAt")

	pCursor := r.FormValue("cursor")
	if pCursor != "" {
		cursor, err := datastore.DecodeCursor(pCursor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		q.Start(cursor)
	}
	// end cursor1 OMIT

	// start cursor2 OMIT
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
	// end cursor2 OMIT

	response := struct {
		Cursor string
		Books  []Book
	}{
		Books: books,
	}

	// start cursor3 OMIT
	response.Books = books
	if cursor, err := t.Cursor(); err == nil {
		response.Cursor = cursor.String()
	}
	// end cursor3 OMIT

	je := json.NewEncoder(w)
	if err := je.Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func tx(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start tx OMIT
	key := datastore.NewKey(c, "Book", "book1", 0, nil) // HL

	if err := datastore.RunInTransaction(c, func(c appengine.Context) error { // HL
		var book Book
		if err := datastore.Get(c, key, &book); err != nil { // HL
			return err
		}

		book.Price += 200

		if _, err := datastore.Put(c, key, &book); err != nil { // HL
			return err
		}

		return nil
	}, nil); err != nil { // HL
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end tx OMIT

	w.Write([]byte("success"))
}
