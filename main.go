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
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tmpl := template.Must(template.New("index").ParseFiles(
				"views/form.html",
				"views/_header.html",
			))
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "", http.StatusBadRequest)
	})
	http.HandleFunc("/message-delivered", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			tmpl := template.Must(template.New("result").ParseFiles(
				"views/form.html",
				"views/_header.html"))
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			name := r.FormValue("name")
			message := r.FormValue("message")

			data := map[string]string{"name": name, "message": message}

			if err := tmpl.Execute(w, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		http.Error(w, "", http.StatusBadRequest)
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := Person{Name: "abdulbari"}
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

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
