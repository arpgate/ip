package main

import (
        "io"
        "net/http"
        "strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
        listIP  := strings.Split(r.Header.Get("X-FORWARDED-FOR"),",")
        io.WriteString(w, listIP[0] )
}

func main() {
        http.HandleFunc("/", hello)
        http.ListenAndServe(":3001", nil)
}

