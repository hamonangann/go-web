package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name string
}

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := Person{Name: "abdul"}
		tmpl := template.Must(template.ParseFiles(
			"views/index.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			log.Print(err.Error())
			//http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := Person{Name: "nito"}
		tmpl := template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))
		err := tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
