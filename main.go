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
	err := config.TPL.ExecuteTemplate(w, "home.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
