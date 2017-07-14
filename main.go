package main

import (
    "io"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request){
     io.WriteString(w, "Hey there")
}

func main() {
    http.HandlFunc("/",hello)
    http.ListenAndServe(":80",nil)
}