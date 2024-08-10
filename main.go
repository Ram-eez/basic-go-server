package main

import (
	"fmt"
	"log"
	"net/http"
)

func formpageHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./static/form.html")
}

func formHandler(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(response, "ParseForm() err: %v", err)
	}

	name := request.FormValue("name")
	if len(name) == 0 {
		http.Redirect(response, request, "/form-page", http.StatusSeeOther)
		return
	}

	address := request.FormValue("address")

	fmt.Fprintf(response, "POST successfull")
	fmt.Fprintf(response, "Name = %s\n", name)
	fmt.Fprintf(response, "Address = %s\n", address)

}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "Method not supported", http.StatusNotFound)
	}

	fmt.Fprintf(response, "HELLOOO!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form-page", formpageHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
