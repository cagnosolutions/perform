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
	mux.AddRoutes(login_page, employee_information, employee_information_overview, employee_profile, employee_profile_overview, job_information, job_information_overview, register_page, login, logout, observation_goals, business_results, business_results_overview, job_information_overview, learning_activities, learning_activities_overview)
	// mux.AddSecureRoutes()
	fmt.Println("Dont forget to register all routes*******************************")
	log.Fatal(http.ListenAndServe(":8088", mux))
}

var login_page = web.Route{"GET", "/login_page", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "login_page.tmpl", web.Model{})
	return
}}

var learning_activities_overview = web.Route{"GET", "/learning_activities_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "learning_activities_overview.tmpl", web.Model{
		"page":    "employeedevelopement",
		"subpage": "learningactivities",
	})
}}

var learning_activities = web.Route{"GET", "/learning_activities", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "learning_activities.tmpl", web.Model{
		"page":    "employeedevelopement",
		"subpage": "learningactivities",
	})
}}

var employee_information = web.Route{"GET", "/employee_information", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_information.tmpl", web.Model{
		"page":    "employeeinformation",
		"subpage": "employeeinformation",
	})
	return
}}

var employee_information_overview = web.Route{"GET", "/employee_information_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_information_overview.tmpl", web.Model{
		"page":    "employeeinformation",
		"subpage": "employeeinformation",
	})
	return
}}

var employee_profile = web.Route{"GET", "/employee_profile", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_profile.tmpl", web.Model{
		"page":    "employeeinformation",
		"subpage": "employeeprofile",
	})
	return
}}

var employee_profile_overview = web.Route{"GET", "/employee_profile_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_profile_overview.tmpl", web.Model{
		"page":    "employeeinformation",
		"subpage": "employeeprofile",
	})
	return
}}

var job_information = web.Route{"GET", "/job_information", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "job_information.tmpl", web.Model{
		"page":    "companyinformation",
		"subpage": "jobinformation",
	})
	return
}}

var job_information_overview = web.Route{"GET", "/job_information_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "job_information_overview.tmpl", web.Model{
		"page":    "companyinformation",
		"subpage": "jobinformation",
	})
	return
}}

var business_results = web.Route{"GET", "/business_results", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "business_results.tmpl", web.Model{
		"page":    "companyinformation",
		"subpage": "businessresults",
	})
}}
var business_results_overview = web.Route{"GET", "/business_results_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "business_results_overview.tmpl", web.Model{
		"page":    "companyinformation",
		"subpage": "businessresults",
	})
}}

var register_page = web.Route{"GET", "/register", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "registerpage.tmpl", web.Model{})
	return
}}

var observation_goals = web.Route{"GET", "/observation_goals", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "observation_goals.tmpl", web.Model{
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
