package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	PageTitle string
}

func main() {

	page := Page{
		PageTitle: "Login",
	}
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal("Error in loading template")
	}

	base := func(w http.ResponseWriter, _ *http.Request) {
		tmpl.Execute(w, page)
	}

	server := &http.ServeMux{}
	server.HandleFunc("/", base)

	fmt.Println("Auth Server running on port 8080")
	http.ListenAndServe(":8080", server)

}
