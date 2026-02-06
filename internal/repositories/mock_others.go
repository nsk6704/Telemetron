package repositories

import (
	"sync"
	"time"

	"telemetron/internal/models"
)

// MockWorkloadRepository implementation
type MockWorkloadRepository struct {
	mu        sync.RWMutex
	workloads []models.Workload
}

func NewMockWorkloadRepository() *MockWorkloadRepository {
	return &MockWorkloadRepository{
		workloads: []models.Workload{
			{
				DeploymentName: "agent-deployment-1",
				MaxPods:        10,
				PodMaxRAM:      "2Gi",
				PodMaxCPU:      "1000m",
				Live: models.LiveWorkload{
					ActivePods: 3,
					UpdatedAt:  time.Now().Format(time.RFC3339),
				},
				Pods: []models.Pod{
					{PodID: "pod-1", CPU: 0.5, Memory: 1024, Status: "running"},
					{PodID: "pod-2", CPU: 0.3, Memory: 512, Status: "running"},
					{PodID: "pod-3", CPU: 0.2, Memory: 256, Status: "running"},
				},
			},
		},
	}
}

func (r *MockWorkloadRepository) GetAll() ([]models.Workload, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	workloads := make([]models.Workload, len(r.workloads))
	copy(workloads, r.workloads)
	return workloads, nil
}

func (r *MockWorkloadRepository) Close() {}

// MockQueueRepository implementation
type MockQueueRepository struct {
	mu     sync.RWMutex
	queues []models.Queue
}

func NewMockQueueRepository() *MockQueueRepository {
	return &MockQueueRepository{
		queues: []models.Queue{
			{
				Name:      "default",
				UpdatedAt: time.Now().Format(time.RFC3339),
				Tasks: []models.QueueTask{
					{ID: "task-1", Priority: models.Priority{Level: "high"}, SubmittedAt: time.Now().Add(-5 * time.Minute).Format(time.RFC3339)},
					{ID: "task-2", Priority: models.Priority{Level: "medium"}, SubmittedAt: time.Now().Add(-2 * time.Minute).Format(time.RFC3339)},
				},
			},
			{
				Name:      "priority",
				UpdatedAt: time.Now().Format(time.RFC3339),
				Tasks: []models.QueueTask{
					{ID: "task-3", Priority: models.Priority{Level: "high"}, SubmittedAt: time.Now().Add(-1 * time.Minute).Format(time.RFC3339)},
				},
			},
		},
	}
}

func (r *MockQueueRepository) GetAll() ([]models.Queue, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	queues := make([]models.Queue, len(r.queues))
	copy(queues, r.queues)
	return queues, nil
}

func (r *MockQueueRepository) Close() {}

// MockLiteLLMRepository implementation
type MockLiteLLMRepository struct {
	mu      sync.RWMutex
	litellm []models.LiteLLM
}

func NewMockLiteLLMRepository() *MockLiteLLMRepository {
	return &MockLiteLLMRepository{
		litellm: []models.LiteLLM{
			{Model: "gpt-4", Provider: "openai", TPM: 45000, RPM: 200, TPMMax: 90000, RPMMax: 3500, PaymentType: "pay-per-request"},
			{Model: "gpt-3.5-turbo", Provider: "openai", TPM: 120000, RPM: 3400, TPMMax: 240000, RPMMax: 3500, PaymentType: "pay-per-request"},
			{Model: "claude-3-opus", Provider: "anthropic", TPM: 30000, RPM: 150, TPMMax: 80000, RPMMax: 3000, PaymentType: "pay-per-request"},
		},
	}
}

func (r *MockLiteLLMRepository) GetAll() ([]models.LiteLLM, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	litellm := make([]models.LiteLLM, len(r.litellm))
	copy(litellm, r.litellm)
	return litellm, nil
}

func (r *MockLiteLLMRepository) Close() {}
