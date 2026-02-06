package repositories

import (
	"sync"
	"time"

	"telemetron/internal/models"
)

type MockAgentRepository struct {
	mu     sync.RWMutex
	agents []models.Agent
	stop   chan struct{}
}

func NewMockAgentRepository() *MockAgentRepository {
	repo := &MockAgentRepository{
		stop: make(chan struct{}),
		agents: []models.Agent{
			{
				Name:                   "agent-1",
				Description:            "Data processing agent",
				MaxParallelInvocations: 5,
				DeploymentName:         "agent-deployment-1",
				Models:                 []string{"gpt-4", "gpt-3.5-turbo"},
				Activity: models.Activity{
					ActiveTaskIDs: []models.TaskStatus{
						{ID: "task-1", Status: "running"},
						{ID: "task-2", Status: "pending"},
					},
					UpdatedAt: time.Now().Format(time.RFC3339),
				},
			},
			{
				Name:                   "agent-2",
				Description:            "Analytics agent",
				MaxParallelInvocations: 3,
				DeploymentName:         "agent-deployment-2",
				Models:                 []string{"gpt-4"},
				Activity: models.Activity{
					ActiveTaskIDs: []models.TaskStatus{
						{ID: "task-3", Status: "running"},
					},
					UpdatedAt: time.Now().Format(time.RFC3339),
				},
			},
		},
	}

	go repo.simulateActivity()
	return repo
}

func (r *MockAgentRepository) GetAll() ([]models.Agent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	agents := make([]models.Agent, len(r.agents))
	copy(agents, r.agents)
	return agents, nil
}

func (r *MockAgentRepository) GetByName(name string) (*models.Agent, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, agent := range r.agents {
		if agent.Name == name {
			agentCopy := agent
			return &agentCopy, true
		}
	}
	return nil, false
}

func (r *MockAgentRepository) simulateActivity() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	statuses := []string{"running", "pending", "completed", "failed"}

	for {
		select {
		case <-ticker.C:
			r.mu.Lock()
			for i := range r.agents {
				// Simulate task status changes
				for j := range r.agents[i].Activity.ActiveTaskIDs {
					r.agents[i].Activity.ActiveTaskIDs[j].Status = statuses[time.Now().Unix()%int64(len(statuses))]
				}
				r.agents[i].Activity.UpdatedAt = time.Now().Format(time.RFC3339)
			}
			r.mu.Unlock()
		case <-r.stop:
			return
		}
	}
}

func (r *MockAgentRepository) Close() {
	close(r.stop)
}
