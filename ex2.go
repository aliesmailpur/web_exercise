package main

import (
	"io"
	"io/ioutil"
	"net/http"
)


func main(){
	http.HandleFunc("/", ali)
	http.ListenAndServe(":8080", nil)
}


func ali(res http.ResponseWriter, req *http.Request) {
	var s string
    
	if req.Method == http.MethodPost {
	// catching the file :
	f, _, err := req.FormFile("q")
	if err != nil{
		http.Error(res, "file not found", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// reads the file
	bs , err := ioutil.ReadAll(f)
	if err != nil{
		http.Error(res, "could not read", http.StatusInternalServerError)
		return
	}

	s = string(bs)

}
    res.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(res, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)

    
}