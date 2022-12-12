package main

import (
	"fmt"
	"net/http"
)


func main(){
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter,req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "coockie",
		Value: "some value",
	})
	fmt.Fprintln(w, "COOCKIE WRITTEN")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("coockie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE :", c)
}