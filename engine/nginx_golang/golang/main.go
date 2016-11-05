package main

import (  
    "fmt"
    "log"
    "net/http"
)

func main() {  
    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Powered-By", "go version go1.7.3 linux/amd64")
        fmt.Fprintln(w, "")
        log.Println("Served Response: Favicon")
    })
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Powered-By", "go version go1.7.3 linux/amd64")
        fmt.Fprintln(w, "Hello World!")
        log.Println("Served Response: Hello World")
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}