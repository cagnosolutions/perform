package main

import (
	"net/http"
	"time"

	"github.com/cagnosolutions/web"
)

var performAdmin = web.Route{"GET", "/performAdmin", func(w http.ResponseWriter, r *http.Request) {
	var PersonnelDatas []PersonnelData
	db.All("personnelData", &PersonnelDatas)
	tmpl.Render(w, r, "performAdminUser.tmpl", web.Model{
		"personnelDatas": PersonnelDatas,
	})
}}

var performAdminPOST = web.Route{"POST", "/performAdmin", func(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var personnelData PersonnelData
	FormToStruct(&personnelData, r.Form, "")
	company := Company{
		Id:             GenId(),
		Name:           r.FormValue("companyName"),
		RegisteredDate: time.Now().UnixNano(),
	}
	personnelData.Id = GenId()
	personnelData.Password = personnelData.Email
	personnelData.Role = "COMPANY_ADMIN"
	personnelData.Active = true
	personnelData.DateAdded = time.Now().UnixNano()
	personnelData.AdminFlag = true
	personnelData.CompanyID = company.Id

	db.Add("personnelData", personnelData.Id, personnelData)
	db.Add("company", company.Id, company)
	web.SetSuccessRedirect(w, r, "/performAdmin", "Added company and admin")
	return
}}

var performCompany = web.Route{"GET", "/perform/company", func(w http.ResponseWriter, r *http.Request) {
	var companies []Company
	db.All("company", &companies)
	tmpl.Render(w, r, "performCompany.tmpl", web.Model{
		"companies": companies,
	})
}}

var performCompanySave = web.Route{"POST", "/perform/company/save", func(w http.ResponseWriter, r *http.Request) {
	var company Company
	db.Get("company", r.FormValue("Id"), &company)
	FormToStruct(&company, r.Form, "")
	if company.Id == "" {
		company.Id = GenId()
		company.RegisteredDate = time.Now().UnixNano()
	}
	db.Set("company", company.Id, company)
	web.SetSuccessRedirect(w, r, "/perform/company/"+company.Id, "Successfully saved sompany")
}}

var performCompanyView = web.Route{"GET", "/perform/company/:id", func(w http.ResponseWriter, r *http.Request) {
	var company Company
	if !db.Get("company", r.FormValue(":id"), &company) {
		web.SetErrorRedirect(w, r, "/perform/company", "Error finding company")
		return
	}

	tmpl.Render(w, r, "performCompanyView.tmpl", web.Model{
		"company": company,
	})
}}
