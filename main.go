package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)
type Film struct {
	Title string
	Director string
}

func main() {
	fmt.Println("server run at port:8000")

	h1 := func (w http.ResponseWriter, r *http.Request)  {
		tmpl := template.Must(template.ParseFiles("index.html"))
	
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Fracncis Ford Coppola"},
				{Title: "Blade Runer", Director: "Rideley Scott"},					
			},
		}

		tmpl.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}