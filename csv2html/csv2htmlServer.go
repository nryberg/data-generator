package main

import (
	"log"
	"net/http"

	"github.com/alecthomas/template"
)

func main() {
	log.Println("Starting webserver at http://localhost:8080/show")
	http.HandleFunc("/show", handler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tableTemplate.html")
	t.Execute(w, r)
}
