package todo

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"net/http"
)

func init() {
	http.HandleFunc("/todo/done", delete)
}

func delete(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		http.Error(w, "login required.", http.StatusForbidden)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// start 1 OMIT
	// 以下errorチェック省略

	key, err := datastore.DecodeKey(r.FormValue("key"))
	if err != nil { // OMIT
		http.Error(w, err.Error(), http.StatusBadRequest) //OMIT
		return                                            //OMIT
	} //OMIT

	var todo Todo
	err = datastore.Get(c, key, &todo)
	if err != nil { //OMIT
		http.Error(w, err.Error(), http.StatusNotFound) //OMIT
		return                                          //OMIT
	} //OMIT

	if u.ID != todo.UserId {
		http.Error(w, "forbidden access.", http.StatusForbidden)
		return
	}

	todo.Done = true

	_, err = datastore.Put(c, key, &todo)
	if err != nil { //OMIT
		http.Error(w, err.Error(), http.StatusInternalServerError) //OMIT
		return                                                     //OMIT
	} //OMIT

	http.Redirect(w, r, "/todo", http.StatusFound)
	// end 1 OMIT
}
