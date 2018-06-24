package main

import (
	"log"
	"net/http"
	"text/template"
)

func renderHTML(w http.ResponseWriter, dir string) error {
	t, err := template.ParseFiles(dir)
	if err != nil {
		return err
	}
	err = t.Execute(w, nil)
	return err
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderHTML(w, "search.html"); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderHTML(w, "test.html"); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderHTML(w, "login.html"); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/", mainHandler)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("listenAndServe：", err.Error())
	}
}
