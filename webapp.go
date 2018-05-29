package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.gohtml")
	name := "Zack"
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", process)
	http.ListenAndServe(":8000", nil)
}
