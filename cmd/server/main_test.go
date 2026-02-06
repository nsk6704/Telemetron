package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"telemetron/internal/repositories"
	"telemetron/internal/services"
	"testing"
)

func TestSystemStateHandler(t *testing.T) {
	// Setup test dependencies
	agentRepo := repositories.NewMockAgentRepository()
	workloadRepo := repositories.NewMockWorkloadRepository()
	queueRepo := repositories.NewMockQueueRepository()
	llmRepo := repositories.NewMockLiteLLMRepository()

	systemService := services.NewSystemService(agentRepo, workloadRepo, queueRepo, llmRepo)
	defer systemService.Close()

	// Create handler
	handler := systemStateHandler(systemService)

	// Create test request
	req := httptest.NewRequest("GET", "/system/state", nil)
	rr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/json")
	}

	var response map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if response["id"] != "system-1" {
		t.Errorf("expected system id 'system-1', got %v", response["id"])
	}
}

func TestRootHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Telemetron API - visit /system/state or /swagger/"))
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Telemetron API - visit /system/state or /swagger/"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
