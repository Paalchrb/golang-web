package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct{
	UserName string
	Password []byte
	First string
	Last string
}

var tpl *template.Template
var dbUsers = map[string]user{} 			//userID, user-struct
var dbSessions = map[string]string{} 	//sessionID, userID

func init() {
	//initialize templates
	tpl = template.Must(template.ParseGlob("templates/*"))

	//create initial user:
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	initUser := user{
		UserName: "test@test.com", 
		Password: bs, 
		First: "Hames", 
		Last: "Bondage",
	}

	//add user to DB:
	dbUsers[initUser.UserName] = initUser
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	//redirect if user not logged in
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	//redirect if already logged in:
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	//process from submission
	if req.Method == http.MethodPost {
		//get form values
		un := req.FormValue("email")
		p := req.FormValue("passeord")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		//check if username is taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		sID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un

		//store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	//check if already logged in
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	//process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")

		//check if user exists
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Invalid credentials", http.StatusForbidden)
			return
		}

		//check if provided password is valid
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusForbidden)
			return
		}

		//create session
		sID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	cookie, _ := req.Cookie("session")
	//delete the session
	delete(dbSessions, cookie.Value)
	//remove the cookie
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	//redirect to login
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}