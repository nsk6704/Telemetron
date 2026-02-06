package services

import (
	"telemetron/internal/repositories"
	"testing"
)

func TestNewSystemService(t *testing.T) {
	agentRepo := repositories.NewMockAgentRepository()
	workloadRepo := repositories.NewMockWorkloadRepository()
	queueRepo := repositories.NewMockQueueRepository()
	llmRepo := repositories.NewMockLiteLLMRepository()

	service := NewSystemService(agentRepo, workloadRepo, queueRepo, llmRepo)
	defer service.Close()

	if service == nil {
		t.Fatal("Expected non-nil service")
	}
}

func TestGetSystemState_Success(t *testing.T) {
	// Arrange
	agentRepo := repositories.NewMockAgentRepository()
	workloadRepo := repositories.NewMockWorkloadRepository()
	queueRepo := repositories.NewMockQueueRepository()
	llmRepo := repositories.NewMockLiteLLMRepository()

	service := NewSystemService(agentRepo, workloadRepo, queueRepo, llmRepo)
	defer service.Close()

	// Act
	state, err := service.GetSystemState()

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if state == nil {
		t.Fatal("Expected non-nil state")
	}

	if state.ID != "system-1" {
		t.Errorf("Expected ID 'system-1', got %s", state.ID)
	}

	if len(state.Agents) == 0 {
		t.Error("Expected agents in system state")
	}

	if len(state.Workload) == 0 {
		t.Error("Expected workload in system state")
	}
}
