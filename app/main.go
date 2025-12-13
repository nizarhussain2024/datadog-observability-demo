package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", healthCheckHandler)

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		w.Header().Set("Content-Type", "application/json")
		
		stats := metrics.GetStats()
		fmt.Fprintf(w, `{
			"service":"datadog-observability-demo",
			"message":"Data endpoint",
			"timestamp":"%s",
			"metrics":%s
		}`, time.Now().Format(time.RFC3339), toJSON(stats))
		
		duration := time.Since(start)
		metrics.RecordRequest(duration, false)
		enhancedMetrics.RecordRequest("/api/data", duration, false)
	})

	http.HandleFunc("/api/metrics/enhanced", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		stats := enhancedMetrics.GetEnhancedStats()
		fmt.Fprintf(w, `%s`, toJSON(stats))
	})

	http.HandleFunc("/api/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		stats := metrics.GetStats()
		fmt.Fprintf(w, `%s`, toJSON(stats))
	})

	fmt.Println("Datadog Observability Demo running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
