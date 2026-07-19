package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Hub struct {
}

func main() {

	// Optional: Serve a minimal HTML page to test the client locally
	http.HandleFunc("/", serveHome)

	fmt.Println("Chat server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("static/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
