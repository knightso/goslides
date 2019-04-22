package tq

import (
	"appengine"
	"appengine/taskqueue"
	"appengine/delay"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func init() {
	http.HandleFunc("/tq/push", push)
	http.HandleFunc("/tq/handler", handler)
	http.HandleFunc("/tq/add4pull", add4pull)
	http.HandleFunc("/tq/lease", lease)
	http.HandleFunc("/tq/delay", addDelay)
}

func push(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	t := taskqueue.NewPOSTTask("/tq/handler", map[string][]string{
		"Title":  {"Perfect Go"},
		"Author": {"Some Gopher"},
		"Price":  {"1000"},
	})

	if _, err := taskqueue.Add(c, t, "pushtest"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("success"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	price, err := strconv.Atoi(r.PostFormValue("Price"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s := struct {
		Title  string
		Author string
		Price  int
	}{
		r.PostFormValue("Title"),
		r.PostFormValue("Author"),
		price,
	}

	c.Infof("executed handler!")
	c.Infof("%v", s)

	w.Write([]byte("success"))
}

func add4pull(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "pull me!! %s", time.Now())

	t := &taskqueue.Task{
		Payload: buf.Bytes(),
		Method:  "PULL",
	}

	if _, err := taskqueue.Add(c, t, "pulltest"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("success"))
}

func lease(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	tasks, err := taskqueue.Lease(c, 100, "pulltest", 3600)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payloads := make([]string, 0, len(tasks))
	for _, task := range tasks {
		payloads = append(payloads, string(task.Payload))
	}

	je := json.NewEncoder(w)
	if err := je.Encode(payloads); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type Book struct {
	Title     string
	Author    string
	Price     int
	CreatedAt time.Time
}

var laterFunc = delay.Func("key", func(c appengine.Context, book *Book) {
	c.Infof("executed delay!")
	c.Infof("%v", book)
})

func addDelay(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	laterFunc.Call(c, &Book{
		"Perfect Go",
		"Some Gopher",
		1000,
		time.Now(),
	})

	laterFunc.Call(c, &Book{
		"Let It Go",
		"Hoge Gopher",
		3000,
		time.Now(),
	})
	
	w.Write([]byte("success"))
}
