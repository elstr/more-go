package main

/*
	Small webserver in go
	Makes use of gorilla mux multiplexor
	It allows complex routing and param use
*/

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Greeting parameters for the template to use
type Greeting struct {
	Hello string
	Name  string
}

// RootHandler handles requests to /
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello world! </h1>")
}

// NameHandler handles requests to /name
// Takes hello and name from querystring
// E.g. ?hello=greetings&name=lele
func NameHandler(w http.ResponseWriter, r *http.Request) {
	hello := r.FormValue("hello")
	name := r.FormValue("name")
	fmt.Fprintf(w, "<h1>%s, %s</h1>", hello, name)
}

// ParamHandler handles requests to /string/string
// Takes parameters from the url not from querystring
func ParamHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // used to parse path variables
	fmt.Fprintf(w, "<h1>%s, %s</h1>", params["first"], params["second"])
}

// UseTemplate makes use of a template for the response
func UseTemplate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	t := template.New("my-template")

	g := Greeting{
		params["hello"],
		params["name"],
	}

	t.Parse(`<h1>{{.Hello}}, {{.Name}}! this is a template :)</h1>`)
	t.Execute(w, g)

}

func main() {
	gmux := mux.NewRouter()
	gmux.HandleFunc("/", RootHandler)
	gmux.HandleFunc("/name", NameHandler)
	gmux.HandleFunc("/{first}/{second}", ParamHandler)
	gmux.HandleFunc("/template/{hello}/{name}", UseTemplate)
	http.ListenAndServe(":8000", gmux)
}
