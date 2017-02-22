package main

import "github.com/cagnosolutions/web"

var PERFORM_ADMIN = web.Auth{
	Roles:    []string{"PERFORM_ADMIN"},
	Redirect: "/login",
	Msg:      "You're not an admin",
}

var COMPANY_ADMIN = web.Auth{
	Roles:    []string{"PERFORM_ADMIN, COMPANY_ADMIN"},
	Redirect: "/login",
	Msg:      "You're not an admin",
}

var COMPANY_EMPLOYEE = web.Auth{
	Roles:    []string{"PERFORM_ADMIN, COMPANY_ADMIN, COMPANY_EMPLOYEE"},
	Redirect: "/login",
	Msg:      "You're not an admin",
}
