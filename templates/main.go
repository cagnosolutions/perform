package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cagnosolutions/repono"
	"github.com/cagnosolutions/web"
)

var mux *web.Mux
var tmpl *web.TmplCache
var db *repono.Client

func init() {
	db = repono.Embed(repono.EmbedAndServeDB(":9090"))
	db.AddStore("user")
	tmpl = web.NewTmplCache()
}

func main() {
	mux = web.NewMux()
	mux.AddSecureRoutes(ADMIN, job_information, index, employee_information, employee_overview)
	mux.AddSecureRoutes(ADMIN, job_information_overview, job_canidates_overview, job_canidates, company_settings)
	mux.AddRoutes(login_page, register_page, login, register)
	// mux.AddSecureRoutes()
	fmt.Println("Dont forget to register all routes*******************************")
	log.Fatal(http.ListenAndServe(":9090", mux))
}

var register = web.Route{"POST", "/register", func(w http.ResponseWriter, r *http.Request) {
	id := strconv.Itoa(int(time.Now().UnixNano()))

	user := User{
		ID:       id,
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	db.Add("user", id, user)
}}

var login = web.Route{"POST", "/login", func(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "admin@admin" && password == "admin" {
		web.Login(w, r, "ADMIN")
		web.SetSuccessRedirect(w, r, "/", "Welcome Admin!")
		return
	}
	web.SetErrorRedirect(w, r, "/login", "Incorrect Email or Password")
}}

var performance_reviews = web.Route{"GET", "/performance_reviews", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "performance_reviews.tmpl", web.Model{})
	return
}}

var company_settings_general = web.Route{"GET", "/company_settings_general", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "company_settings_general.tmpl", web.Model{})
	return
}}

var company_settings_branches = web.Route{"GET", "/company_settings_branches", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "company_settings_branches.tmpl", web.Model{})
	return
}}

var company_settings = web.Route{"GET", "/company_settings", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "company_settings.tmpl", web.Model{})
	return
}}

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
	tmpl.Render(w, r, "learning_activities_overview.tmpl", web.Model{})
	return
}}

var learning_activities = web.Route{"GET", "/learning_activities", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "learning_activities.tmpl", web.Model{})
	return
}}

var employee_information = web.Route{"GET", "/employee_information", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_information.tmpl", web.Model{})
	return
}}

var employee_overview = web.Route{"GET", "/employee_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_overview.tmpl", web.Model{})
	return
}}

var employee_profile = web.Route{"GET", "/employee_profile", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "employee_profile.tmpl", web.Model{})
	return
}}

var job_information = web.Route{"GET", "/job_information", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "job_information.tmpl", web.Model{})
	return
}}

var job_information_overview = web.Route{"GET", "/job_information_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "job_information_overview.tmpl", web.Model{})
	return
}}

var business_results = web.Route{"GET", "/business_results", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "business_results.tmpl", web.Model{})
	return
}}
var business_results_overview = web.Route{"GET", "/business_results_overview", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "business_results_overview.tmpl", web.Model{})
	return
}}

var register_page = web.Route{"GET", "/register", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "register_page.tmpl", web.Model{})
	return
}}

var observation_goals = web.Route{"GET", "/observation_goals", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "observation_goals.tmpl", web.Model{})
	return
}}
