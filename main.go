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
	mux.AddRoutes(login_page, employee_information, employee_overview, employee_profile, job_information, job_information_overview, register_page, login, logout, observation_goals, business_results)
	mux.AddRoutes(one_to_one_measurement, business_results_overview, job_information_overview, learning_activities, learning_activities_overview, index, course_registrations, course_sessions)
	mux.AddRoutes(multi_measurement_performance, peer_to_peer, career_planning, job_canidates, organization_structure, accountability_reviews, job_canidates_overview)
	// mux.AddSecureRoutes()
	fmt.Println("Dont forget to register all routes*******************************")
	log.Fatal(http.ListenAndServe(":9999", mux))
}

var job_canidates_overview = web.Route{"GET", "/job_canidates_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "job_canidates_overview.tmpl", web.Model{})
	return
}}

var accountability_reviews = web.Route{"GET", "/accountability_reviews", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "accountability_reviews.tmpl", web.Model{})
	return
}}

var organization_structure = web.Route{"GET", "/organization_structure", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "organization_structure.tmpl", web.Model{})
	return
}}

var one_to_one_measurement = web.Route{"GET", "/one_to_one_measurement", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "one_to_one_measurement.tmpl", web.Model{})
	return
}}
var job_canidates = web.Route{"GET", "/job_canidates", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "job_canidates.tmpl", web.Model{})
	return
}}
var career_planning = web.Route{"GET", "/career_planning", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "career_planning.tmpl", web.Model{})
	return
}}

var peer_to_peer = web.Route{"GET", "/peer_to_peer", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "peer_to_peer.tmpl", web.Model{})
	return
}}

var multi_measurement_performance = web.Route{"GET", "/multi_measurement_performance", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "multi_measurement_performance.tmpl", web.Model{})
	return
}}

var course_registrations = web.Route{"GET", "/course_registrations", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "course_registrations.tmpl", web.Model{})
	return
}}

var course_sessions = web.Route{"GET", "/course_sessions", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "course_sessions.tmpl", web.Model{})
	return
}}

var index = web.Route{"GET", "/", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "index.tmpl", web.Model{})
	return
}}

var login_page = web.Route{"GET", "/login", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "login_page.tmpl", web.Model{})
	return
}}

var learning_activities_overview = web.Route{"GET", "/learning_activities_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "learning_activities_overview.tmpl", web.Model{
		"page":    "employeedevelopement",
		"subpage": "learningactivities",
	})
	return
}}

var learning_activities = web.Route{"GET", "/learning_activities", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "learning_activities.tmpl", web.Model{
		"page":    "employeedevelopement",
		"subpage": "learningactivities",
	})
	return
}}

var employee_information = web.Route{"GET", "/employee_information", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_information.tmpl", web.Model{
		"page":    "employeeinformation",
		"subpage": "employeeinformation",
	})
	return
}}

var employee_overview = web.Route{"GET", "/employee_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_overview.tmpl", web.Model{
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
	tmpl.Render(w, r, "register_page.tmpl", web.Model{})
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
	company := r.FormValue("company")
	username := r.FormValue("username")
	password := r.FormValue("password")
	if company == "admin" && username == "admin" && password == "admin" {
		web.Login(w, r, "ADMIN")
		web.SetSuccessRedirect(w, r, "/", "Welcome Admin!")
		return
	}
	web.SetErrorRedirect(w, r, "/login", "Incorrect Email or Password")
}}

var logout = web.Route{"GET", "/logout", func(w http.ResponseWriter, r *http.Request) {
	web.Logout(w)
	web.SetSuccessRedirect(w, r, "/loginpage", "Logged Out")
	return
}}
