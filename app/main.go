package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Datadog Observability Running")
    })
    http.ListenAndServe(":8081", nil)
}
