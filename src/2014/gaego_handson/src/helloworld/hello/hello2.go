package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

// start OMIT
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello, World!")
	fmt.Fprint(w, "Hello, gopher!")
}
// end OMIT
