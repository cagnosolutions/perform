package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cagnosolutions/adb"
	"github.com/cagnosolutions/web"
)

var mux *web.Mux
var tmpl *web.TmplCache
var db *adb.DB

func init() {
	db = adb.NewDB()
	db.AddStore("user")
	db.AddStore("personnelData")
	db.AddStore("company")
	mux = web.NewMux()
	tmpl = web.NewTmplCache()
}

func main() {
	mux.AddRoutes(login, register, logout, login_page, register_page)

	mux.AddSecureRoutes(PERFORM_ADMIN, performAdmin, performAdminPOST, performCompany, performCompanyView, performCompanySave)
	mux.AddSecureRoutes(PERFORM_ADMIN, index, employee_information, employee_overview, business_results_overview, business_results, job_information, job_information_overview)
	mux.AddSecureRoutes(PERFORM_ADMIN, course_registrations, job_canidates_overview, job_canidates, employee_profile, course_overview, course, course_sessions, career_planning, performance_reviews, accountability_reviews)

	fmt.Println("Listening on :9090")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=- Register Your Routes -=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	log.Fatal(http.ListenAndServe(":9090", mux))
}
