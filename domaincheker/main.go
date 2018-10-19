package main

import (
    "fmt"
    "log"
    "net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("jbfewfwufe")
    //message:=domain_search.DomainSearch("python")
    message:="Stasichuy!!!"
    fmt.Fprintf(w, message)
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)
    http.HandleFunc("/log",postHandler)

    log.Println("Listening...")
    http.ListenAndServe(":3000", nil)
}
