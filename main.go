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
	mux.AddRoutes(index)
	// mux.AddSecureRoutes()
	fmt.Println("Dont forget to register all routes*******************************")
	log.Fatal(http.ListenAndServe(":8088", mux))
}

var index = web.Route{"GET", "/", func(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, r, "index.tmpl", web.Model{})
	return
}}
