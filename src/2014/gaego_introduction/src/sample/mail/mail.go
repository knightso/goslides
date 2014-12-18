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
	// start mailreceive1 OMIT
	http.HandleFunc("/_ah/mail/", receive)
	// end mailreceive1 OMIT
}

func send(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start mailsend OMIT
	msg := &mail.Message{ // HL
		Sender:  "Hoge <hoge@example.com>",
		To:      []string{"to@example.com"},
		Subject: "This is test mail.",
		Body:    "Hello Gopher!",
	}

	if err := mail.Send(c, msg); err != nil { // HL
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end mailsend OMIT

	fmt.Fprint(w, "success")
}

func receive(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start mailreceive2 OMIT
	defer r.Body.Close()

	var b bytes.Buffer
	if _, err := b.ReadFrom(r.Body); err != nil { // HL
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end mailreceive2 OMIT

	c.Debugf("%s", b.String())
}
