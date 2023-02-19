package shortenurl

import (
	"gorm.io/gorm"
	"net/http"
)

type URL struct {
	gorm.Model
	Original  string
	Shortened string
}

func Redirect(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	shortened := "http://localhost:8000/" + r.URL.Path[1:]
	var url URL
	db.First(url, "shortened = ?", shortened)
	http.Redirect(w, r, url.Original, http.StatusFound)
}
