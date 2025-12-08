package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"healthy","service":"datadog-demo","timestamp":"%s"}`, time.Now().Format(time.RFC3339))
	})

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"service":"datadog-observability-demo",
			"message":"Data endpoint",
			"timestamp":"%s",
			"metrics":{
				"requests":100,
				"errors":2,
				"latency_ms":45
			}
		}`, time.Now().Format(time.RFC3339))
	})

	fmt.Println("Datadog Observability Demo running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
