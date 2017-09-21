package main

import (
	"log"
	"net/http"

	"./config"
	"github.com/julienschmidt/httprouter"
)

func main() {
	loadRoutes()
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// http.Redirect(w, r, "/books", http.StatusSeeOther)
	u := getUser(w, r)
	err := config.TPL.ExecuteTemplate(w, "home.gohtml", u)
	if err != nil {
		log.Println(err)
	}
}
