package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've reached %s!", r.URL.Path)
    })

    fmt.Println("Server is starting...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
