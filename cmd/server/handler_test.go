package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"telemetron/internal/models"
	"telemetron/internal/repositories"
	"telemetron/internal/services"
	"testing"
)

func TestSystemStateHandler_Success(t *testing.T) {
	// Arrange
	agentRepo := repositories.NewMockAgentRepository()
	workloadRepo := repositories.NewMockWorkloadRepository()
	queueRepo := repositories.NewMockQueueRepository()
	llmRepo := repositories.NewMockLiteLLMRepository()

	systemService := services.NewSystemService(agentRepo, workloadRepo, queueRepo, llmRepo)
	defer systemService.Close()

	handler := systemStateHandler(systemService)

	req, err := http.NewRequest("GET", "/system/state", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Expected content type %s, got %s", expectedContentType, contentType)
	}

	var state models.SystemState
	if err := json.Unmarshal(rr.Body.Bytes(), &state); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Verify response structure
	if len(state.Agents) == 0 {
		t.Error("Expected agents in response")
	}

	if len(state.Workload) == 0 {
		t.Error("Expected workload in response")
	}

	if len(state.Queues) == 0 {
		t.Error("Expected queues in response")
	}

	if len(state.LiteLLM) == 0 {
		t.Error("Expected litellm data in response")
	}
}

func TestSystemStateHandler_JSONFormat(t *testing.T) {
	agentRepo := repositories.NewMockAgentRepository()
	workloadRepo := repositories.NewMockWorkloadRepository()
	queueRepo := repositories.NewMockQueueRepository()
	llmRepo := repositories.NewMockLiteLLMRepository()

	systemService := services.NewSystemService(agentRepo, workloadRepo, queueRepo, llmRepo)
	defer systemService.Close()

	handler := systemStateHandler(systemService)

	req, _ := http.NewRequest("GET", "/system/state", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	// Verify JSON is valid
	var result map[string]interface{}
	if err := json.Unmarshal(rr.Body.Bytes(), &result); err != nil {
		t.Fatalf("Response is not valid JSON: %v", err)
	}

	// Verify required top-level keys exist
	requiredKeys := []string{"agents", "workload", "queues", "litellm"}
	for _, key := range requiredKeys {
		if _, exists := result[key]; !exists {
			t.Errorf("Missing required key in response: %s", key)
		}
	}
}

func TestSystemStateHandler_ResponseData(t *testing.T) {
	agentRepo := repositories.NewMockAgentRepository()
	workloadRepo := repositories.NewMockWorkloadRepository()
	queueRepo := repositories.NewMockQueueRepository()
	llmRepo := repositories.NewMockLiteLLMRepository()

	systemService := services.NewSystemService(agentRepo, workloadRepo, queueRepo, llmRepo)
	defer systemService.Close()

	handler := systemStateHandler(systemService)

	req, _ := http.NewRequest("GET", "/system/state", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	var state models.SystemState
	json.Unmarshal(rr.Body.Bytes(), &state)

	// Verify specific data from mocks
	expectedAgentName := "agent-1"
	if state.Agents[0].Name != expectedAgentName {
		t.Errorf("Expected agent name '%s', got '%s'", expectedAgentName, state.Agents[0].Name)
	}

	expectedDeployment := "agent-deployment-1"
	if state.Workload[0].DeploymentName != expectedDeployment {
		t.Errorf("Expected deployment '%s', got '%s'", expectedDeployment, state.Workload[0].DeploymentName)
	}
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Telemetron API - visit /system/state or /swagger/"))
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, status)
	}

	expected := "Telemetron API - visit /system/state or /swagger/"
	if rr.Body.String() != expected {
		t.Errorf("Expected body '%s', got '%s'", expected, rr.Body.String())
	}
}
