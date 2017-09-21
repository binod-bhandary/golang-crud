package login

import (
	"fmt"
	"log"
	"net/http"

	"../config"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var dbUsers = map[string]User{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// http.Redirect(w, r, "/books", http.StatusSeeOther)
	err := config.TPL.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func LoginProcess(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		user, err := LogUser(req)
		//encrypt db password
		bs, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		//store db to variable
		dbUsers[un] = User{user.ID, user.Fullname, user.Email, user.Username, string(bs)}
		fmt.Println(dbUsers)
		//store value to variable
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))

		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	config.TPL.ExecuteTemplate(w, "login.gohtml", nil)

}

func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// http.Redirect(w, r, "/books", http.StatusSeeOther)
	err := config.TPL.ExecuteTemplate(w, "register.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}

func RegisterProcess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	bk, err := PutUser(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "login.gohtml", bk)
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := config.TPL.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
