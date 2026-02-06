package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestSystemStateJSONSerialization(t *testing.T) {
	// Create test data
	now := time.Now().Format(time.RFC3339)
	state := SystemState{
		ID: "system-1",
		Agents: []Agent{
			{
				Name:           "test-agent",
				Description:    "Initial test agent",
				DeploymentName: "test-deployment",
				Activity: Activity{
					UpdatedAt: now,
				},
			},
		},
		Workload: []Workload{
			{
				DeploymentName: "test-deployment",
				MaxPods:        5,
				PodMaxRAM:      "1Gi",
				PodMaxCPU:      "500m",
				Live: LiveWorkload{
					ActivePods: 2,
					UpdatedAt:  now,
				},
			},
		},
		Queues: []Queue{
			{
				Name:      "test-queue",
				UpdatedAt: now,
				Tasks: []QueueTask{
					{ID: "task-1", Priority: Priority{Level: "high"}},
				},
			},
		},
		LiteLLM: []LiteLLM{
			{
				Model:    "gpt-4",
				Provider: "openai",
				TPM:      1000,
				RPM:      10,
			},
		},
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("Failed to marshal SystemState to JSON: %v", err)
	}

	// Test JSON unmarshaling
	var unmarshaled SystemState
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal JSON to SystemState: %v", err)
	}

	// Verify data integrity
	if unmarshaled.ID != state.ID {
		t.Errorf("System ID mismatch: expected %s, got %s", state.ID, unmarshaled.ID)
	}

	if unmarshaled.Agents[0].Name != state.Agents[0].Name {
		t.Errorf("Agent Name mismatch: expected %s, got %s",
			state.Agents[0].Name, unmarshaled.Agents[0].Name)
	}

	if unmarshaled.Workload[0].DeploymentName != state.Workload[0].DeploymentName {
		t.Errorf("Deployment name mismatch: expected %s, got %s",
			state.Workload[0].DeploymentName, unmarshaled.Workload[0].DeploymentName)
	}
}

func TestEmptySystemState(t *testing.T) {
	state := SystemState{}

	jsonData, err := json.Marshal(state)
	if err != nil {
		t.Fatalf("Failed to marshal empty SystemState: %v", err)
	}

	var unmarshaled SystemState
	if err := json.Unmarshal(jsonData, &unmarshaled); err != nil {
		t.Fatalf("Failed to unmarshal empty SystemState: %v", err)
	}

	// Should have empty slices (not nil if initialized, but here they might be nil)
	if len(unmarshaled.Agents) != 0 {
		t.Error("Expected 0 agents")
	}
}
