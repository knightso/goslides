package todo

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"html/template"
	"net/http"
)

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

	logoutUrl, _ := user.LogoutURL(c, "/")

	t, err := template.ParseFiles("todo/todo.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	q := datastore.NewQuery("Todo").Filter("UserId =", u.ID).Filter("Done =", false).Order("-DueDate")

	// start 1 OMIT
	var todos []Todo
	keys, err := q.GetAll(c, &todos) // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end 1 OMIT

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	// start 2 OMIT
	params := struct {
		LogoutUrl string
		User      *user.User
		Todos     []Todo
		Keys      []*datastore.Key // HL
	}{
		logoutUrl,
		u,
		todos,
		keys, // HL
	}

	err = t.Execute(w, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// end 2 OMIT
}
