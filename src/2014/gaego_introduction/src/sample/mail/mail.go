package mail

import (
	"appengine"
	"appengine/mail"
	"bytes"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/mail/send", send)
	http.HandleFunc("/_ah/mail/", receive)
}

func send(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	msg := &mail.Message{
		Sender:  "Hoge <hoge@example.com>",
		To:      []string{"to@example.com"},
		Subject: "This is test mail.",
		Body:    "Hello Gopher!",
	}

	if err := mail.Send(c, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "success")
}

func receive(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	defer r.Body.Close()

	var b bytes.Buffer
	if _, err := b.ReadFrom(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.Debugf("%s", b.String())
}
