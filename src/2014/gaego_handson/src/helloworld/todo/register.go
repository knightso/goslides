package todo

// start 1 OMIT
import (
	"appengine"
	"appengine/datastore" // HL
	"appengine/user"
	"net/http"
)

// end 1 OMIT

// start 2 OMIT
func init() {
	http.HandleFunc("/todo/register", register)
}

// Todo保存用構造体
type Todo struct {
	UserId  string
	Todo    string
	Notes   string
	DueDate string
	Done    bool
}

// end 2 OMIT

func register(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		http.Error(w, "login required.", http.StatusForbidden)
		return
	}

	// start 3 OMIT
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Entity保存用構造体を用意 // HL
	todo := Todo{
		u.ID,
		r.FormValue("Todo"),
		r.FormValue("Notes"),
		r.FormValue("DueDate"),
		false,
	}
	// end 3 OMIT

	// start 4 OMIT
	key := datastore.NewIncompleteKey(c, "Todo", nil) // HL
	key, err := datastore.Put(c, key, &todo)          // HL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/todo", http.StatusFound) // "/todo"にリダイレクト // HL
	// end 4 OMIT
}
