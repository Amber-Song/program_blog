package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseGlob("demo.html"))
	_ = templates.ExecuteTemplate(w, "demo.html", "")
}

func demo1Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	password := r.Form.Get("password")
	fmt.Println("This is form 1.")
	fmt.Println("Name is ", name, " . Password is ", password)
	var templates = template.Must(template.ParseGlob("demo1.html"))
	_ = templates.ExecuteTemplate(w, "demo1.html", name)
}

func demo2Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	password := r.Form.Get("password")
	fmt.Println("This is form 2.")
	fmt.Println("Name is ", name, " . Password is ", password)
	var templates = template.Must(template.ParseGlob("demo2.html"))
	_ = templates.ExecuteTemplate(w, "demo2.html", map[string]string{"name": name, "password": password})
}

func main() {
	http.HandleFunc("/demo/", homeHandler)
	http.HandleFunc("/demo1/", demo1Handler)
	http.HandleFunc("/demo2/", demo2Handler)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
