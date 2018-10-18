package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"./domain-search"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("dff")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Print(string(body))
    message:=domain_search.DomainSearch(string(body)[5:])
   // message:="Stasichuy!!!"
    fmt.Fprintf(w, message)
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)
    http.HandleFunc("/log",postHandler)

    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}
