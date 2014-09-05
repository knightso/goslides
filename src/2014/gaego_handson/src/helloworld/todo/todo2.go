package todo

// start 1 OMIT
import (
	"appengine"
	"appengine/user"
	"html/template" // HL
	"net/http"
)

// end 1 OMIT

func init() {
	http.HandleFunc("/todo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		loginUrl, _ := user.LoginURL(c, "/todo")
		http.Redirect(w, r, loginUrl, http.StatusFound)
		return
	}

	// start 2 OMIT
	t, err := template.ParseFiles("todo/todo.tmpl") // テンプレート読み込み // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	logoutUrl, _ := user.LogoutURL(c, "/")

	params := struct {
		LogoutUrl string
		User      *user.User
	}{
		logoutUrl,
		u,
	}

	err = t.Execute(w, params) // テンプレート適用! // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end 2 OMIT
}
