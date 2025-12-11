package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthStatus struct {
	Status    string            `json:"status"`
	Service   string            `json:"service"`
	Timestamp string            `json:"timestamp"`
	Checks    map[string]string `json:"checks,omitempty"`
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	status := HealthStatus{
		Status:    "healthy",
		Service:   "datadog-observability-demo",
		Timestamp: time.Now().Format(time.RFC3339),
		Checks: map[string]string{
			"database": "ok",
			"cache":    "ok",
			"api":      "ok",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
