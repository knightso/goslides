package todo

// start 1 OMIT
import (
	"appengine"      // HL
	"appengine/user" // HL
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/todo", handler)
}

// end 1 OMIT

// start 2 OMIT
func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c) // ログインユーザー取得 // HL
	if u == nil {
		// 未ログインの場合はログインURLへリダイレクト // HL
		loginUrl, _ := user.LoginURL(c, "/todo")        // HL
		http.Redirect(w, r, loginUrl, http.StatusFound) // HL
		return                                          // HL
	}
	logoutUrl, _ := user.LogoutURL(c, "/") // ログアウトURL取得 // HL

	w.Header().Set("Content-type", "text/html; charset=utf-8")

	html := `
<html><body>
Hello, %s ! - <a href="%s">sign out</a><br>
<hr>
This is TODO page under constuction!
</body></html>
`
	fmt.Fprintf(w, html, u.Email, logoutUrl)
}

// end 2 OMIT
