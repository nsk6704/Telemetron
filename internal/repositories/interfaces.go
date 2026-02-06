// internal/repositories/interfaces.go
package repositories

import "telemetron/internal/models"

type AgentRepository interface {
	GetAll() ([]models.Agent, error)
	Close()
}

type WorkloadRepository interface {
	GetAll() ([]models.Workload, error)
	Close()
}

type QueueRepository interface {
	GetAll() ([]models.Queue, error)
	Close()
}

type LiteLLMRepository interface {
	GetAll() ([]models.LiteLLM, error)
	Close()
}
