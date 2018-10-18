package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"./domain-search"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}
    message:=domain_search.DomainSearch(string(body)[5:]) //my frontend developer made this crutch
    fmt.Fprintf(w, message)
}

func main() {
	// upload html,css,js file for response
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)
    http.HandleFunc("/log",postHandler)

    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}
