package main

import (
    "fmt"
    "net/http"
)

func server(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "THIS SERVER IS STARTED  FOR ENCRYPTION ! ")
}

func main() {
    http.HandleFunc("/", server)
    http.ListenAndServe(":8888", nil)
}