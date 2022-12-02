package main

import (
	"io"
	"net/http"
)

// ListenAndServe on port ":8080" using the default serveMux .

// Use HandleFunc to add the following routes to the default serve mux

// "/"
// "/dog"
// "/me"

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

func a(res http.ResponseWriter, req *http.Request ) {
	io.WriteString(res, "This is the home page")
}

func b(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the dog page")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Ali esmail pour")
}

func main() {

	http.HandleFunc("/" , a)
	http.HandleFunc("/dog", b)
	http.HandleFunc("/me", c)


	http.ListenAndServe(":8080", nil)
}