package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/cagnosolutions/web"
)

var register = web.Route{"POST", "/register", func(w http.ResponseWriter, r *http.Request) {
	id := strconv.Itoa(int(time.Now().UnixNano()))
	user := PersonnelData{
		Id:        id,
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Role:      "ADMIN",
		Active:    true,
		Gender:    r.FormValue("gender"),
	}
	db.Add("user", id, user)
	web.SetSuccessRedirect(w, r, "/login", "Registered Please Log In")
}}

var logout = web.Route{"GET", "/logout", func(w http.ResponseWriter, r *http.Request) {
	web.Logout(w)
	web.SetMsgRedirect(w, r, "/login", "Later Bro")
	return
}}

var login = web.Route{"POST", "/login", func(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "admin@admin" && password == "admin" {
		web.Login(w, r, "PERFORM_ADMIN")
		web.SetSuccessRedirect(w, r, "/performAdmin", "Welcome In Memory Admin")
		return
	}
	var users []PersonnelData
	db.All("user", &users)
	for _, user := range users {
		if user.Email == email && user.Password == password && user.Active {
			web.Login(w, r, user.Role)
			web.SetMsgRedirect(w, r, "/", "WELCOME "+user.FirstName)
			return
		}
	}
	web.SetErrorRedirect(w, r, "/login", "Incorrect Email or Password")
}}
