package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cagnosolutions/web"
)

var mux *web.Mux
var tmpl *web.TmplCache

func init() {
	mux = web.NewMux()
	tmpl = web.NewTmplCache()
}

func main() {
	mux.AddRoutes(loginpage, empinfo, empprof, jobinfo, registerpage, login, logout, observationgoals, businessresults, businessresultsoverview)
	// mux.AddSecureRoutes()
	fmt.Println("Dont forget to register all routes*******************************")
	log.Fatal(http.ListenAndServe(":8088", mux))
}

var loginpage = web.Route{"GET", "/loginpage", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "loginpage.tmpl", web.Model{})
	return
}}

var empinfo = web.Route{"GET", "/employeeinformation", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "empinfo.tmpl", web.Model{
		"page":    "employeeinformation",
		"subpage": "employeeinformation",
	})
	return
}}

var empprof = web.Route{"GET", "/employeeprofile", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "empprofile.tmpl", web.Model{
		"page":    "employeeinformation",
		"subpage": "employeeprofile",
	})
	return
}}

var jobinfo = web.Route{"GET", "/jobinformation", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "jobinfo.tmpl", web.Model{
		"page":    "companyinformation",
		"subpage": "jobinformation",
	})
	return
}}

var businessresults = web.Route{"GET", "/businessresults", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "businessresults.tmpl", web.Model{
		"page":    "companyinformation",
		"subpage": "businessresults",
	})
}}
var businessresultsoverview = web.Route{"GET", "/businessresultsoverview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "businessresultsoverview.tmpl", web.Model{
		"page":    "companyinformation",
		"subpage": "businessresults",
	})
}}

var registerpage = web.Route{"GET", "/register", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "registerpage.tmpl", web.Model{})
	return
}}

var observationgoals = web.Route{"GET", "/observationgoals", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "observationgoals.tmpl", web.Model{
		"page":    "employeeperformance",
		"subpage": "observationgoals",
	})
	return
}}

var login = web.Route{"POST", "/login", func(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("username")
	password := r.FormValue("password")
	if email == "admin" && password == "admin" {
		web.Login(w, r, "ADMIN")
		web.SetSuccessRedirect(w, r, "/employeeinformation", "Welcome Admin!")
		return
	}
	web.SetErrorRedirect(w, r, "/login", "Incorrect Email or Password")
}}

var logout = web.Route{"GET", "/logout", func(w http.ResponseWriter, r *http.Request) {
	web.Logout(w)
	web.SetSuccessRedirect(w, r, "/loginpage", "Logged Out")
	return
}}
