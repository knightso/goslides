package urlfetch

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/urlfetch/fetch", fetch)
}

func fetch(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// start urlfetch OMIT
	client := urlfetch.Client(c)                      // HL
	resp, err := client.Get("http://www.google.com/") // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end urlfetch OMIT

	c.Debugf("%s", buf.String())

	fmt.Fprint(w, "success")
}
