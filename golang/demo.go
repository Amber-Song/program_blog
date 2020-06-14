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

func demo3Handler(w http.ResponseWriter, r *http.Request) {
	var templates = template.Must(template.ParseGlob("parse_demo.html"))
	_ = templates.ExecuteTemplate(w, "parse_demo.html", 100)
}

type User struct {
	Name  string
	Width []int
}

func demo4Handler(w http.ResponseWriter, r *http.Request) {
	width := []int{100, 200, 300}
	name := "Amber"
	user := User{Name: name, Width: width}
	t, _ := template.New("demo").Parse(`<html>
	<body>
	  {{ range .Width }}
		<div style="width:{{.}}px; background-color:green;">Hi {{$.Name}}</div>
	  {{ end }}
  
	  {{ $name := .Name}}
	  {{ range .Width }}
		<div style="width:{{.}}px; background-color:green;">Hi {{$name}}</div>
	  {{ end }}
	</body>
  </html>`)
	t.Execute(w, user)
}

func main() {
	http.HandleFunc("/demo/", homeHandler)
	http.HandleFunc("/demo1/", demo1Handler)
	http.HandleFunc("/demo2/", demo2Handler)
	http.HandleFunc("/demo3/", demo3Handler)
	http.HandleFunc("/demo4/", demo4Handler)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
