package main

import (
	"log"
	"net/http"
	"text/template"
)

func renderHtml(w http.ResponseWriter, dir string) error {
	t, err := template.ParseFiles(dir)
	if err != nil {
		return err
	}
	err = t.Execute(w, nil)
	return err
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderHtml(w, "login.html"); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", MainHandler)
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("listenAndServe：", err.Error())
	}
}
