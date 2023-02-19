package main

import (
	"fmt"
	"github.com/najmiddinyodgor/urlshortener/shortenurl"
	"net/http"
)

func main() {
	db, err := shortenurl.Connect()

	if err != nil {
		panic("Could not establish a connection")
	}

	db.AutoMigrate(&shortenurl.URL{})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		original := r.FormValue("url")
		shortened := shortenurl.ShortenURL(original)
		fmt.Printf(shortened)

		db.Create(&shortenurl.URL{Original: original, Shortened: shortened})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		shortenurl.Redirect(db, w, r)
	})

	err = http.ListenAndServe(":8000", nil)

	if err == nil {
		fmt.Println("Connected successfully")
	}
}
