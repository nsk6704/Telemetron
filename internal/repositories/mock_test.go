package repositories

import (
	"testing"
)

func TestMockAgentRepository(t *testing.T) {
	repo := NewMockAgentRepository()

	agents, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(agents) == 0 {
		t.Error("Expected agents, got empty slice")
	}

	// Verify agent structure
	agent := agents[0]
	if agent.Name == "" {
		t.Error("Expected agent name, got empty string")
	}

	if agent.DeploymentName == "" {
		t.Error("Expected deployment name, got empty string")
	}

	if agent.Activity.UpdatedAt == "" {
		t.Error("Expected activity updated timestamp")
	}
}

func TestMockWorkloadRepository(t *testing.T) {
	repo := NewMockWorkloadRepository()

	workloads, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(workloads) == 0 {
		t.Error("Expected workloads, got empty slice")
	}

	// Verify workload structure
	workload := workloads[0]
	if workload.DeploymentName == "" {
		t.Error("Expected deployment name, got empty string")
	}

	if workload.MaxPods == 0 {
		t.Error("Expected MaxPods > 0")
	}
}

func TestMockQueueRepository(t *testing.T) {
	repo := NewMockQueueRepository()
	queues, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(queues) == 0 {
		t.Fatal("Expected queues")
	}

	if queues[0].Name == "" {
		t.Error("Expected queue name")
	}
}

func TestMockLiteLLMRepository(t *testing.T) {
	repo := NewMockLiteLLMRepository()
	llms, err := repo.GetAll()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(llms) == 0 {
		t.Fatal("Expected LiteLLM models")
	}

	if llms[0].Model == "" {
		t.Error("Expected model name")
	}
}
