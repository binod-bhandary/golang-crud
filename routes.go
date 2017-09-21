package main

import (
	"net/http"

	"./books"
	"./login"
	"github.com/julienschmidt/httprouter"
)

func defaultRoutes() {

	// http.HandleFunc("/", index)
	// http.HandleFunc("/books", books.Index)
	// http.HandleFunc("/books/show", books.Show)
	// http.HandleFunc("/books/create", books.Create)
	// http.HandleFunc("/books/create/process", books.CreateProcess)
	// http.HandleFunc("/books/update", books.Update)
	// http.HandleFunc("/books/update/process", books.UpdateProcess)
	// http.HandleFunc("/books/delete/process", books.DeleteProcess)
	// http.ListenAndServe(":8080", nil)

}

func loadRoutes() {

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/books", books.Index)
	mux.GET("/books/show", books.Show)
	mux.GET("/books/create", books.Create)
	mux.POST("/books/create/process", books.CreateProcess)
	mux.GET("/books/update", books.Update)
	mux.POST("/books/update/process", books.UpdateProcess)
	mux.GET("/books/delete/process", books.DeleteProcess)

	mux.GET("/register", login.Register)
	mux.POST("/register/process", login.RegisterProcess)

	mux.GET("/login", login.Index)
	mux.POST("/login/process", login.LoginProcess)
	mux.GET("/logout", login.Logout)

	http.ListenAndServe(":8080", mux)

}
