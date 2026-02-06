package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/system/state", systemStateHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Telemetron API - visit /system/state"))
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func systemStateHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"id": "system-1",
		"agents": []map[string]string{
			{"name": "agent-1", "status": "active"},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
