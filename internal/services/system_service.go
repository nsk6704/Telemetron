package services

import (
	"telemetron/internal/models"
	"telemetron/internal/repositories"
)

type SystemService struct {
	agentRepo    repositories.AgentRepository
	workloadRepo repositories.WorkloadRepository
	queueRepo    repositories.QueueRepository
	llmRepo      repositories.LiteLLMRepository
}

func NewSystemService(
	agent repositories.AgentRepository,
	workload repositories.WorkloadRepository,
	queue repositories.QueueRepository,
	llm repositories.LiteLLMRepository,
) *SystemService {
	return &SystemService{
		agentRepo:    agent,
		workloadRepo: workload,
		queueRepo:    queue,
		llmRepo:      llm,
	}
}

func (s *SystemService) GetSystemState() (*models.SystemState, error) {
	agents, err := s.agentRepo.GetAll()
	if err != nil {
		return nil, err
	}

	workloads, err := s.workloadRepo.GetAll()
	if err != nil {
		return nil, err
	}

	queues, err := s.queueRepo.GetAll()
	if err != nil {
		return nil, err
	}

	litellm, err := s.llmRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return &models.SystemState{
		ID:       "system-1",
		Agents:   agents,
		Workload: workloads,
		Queues:   queues,
		LiteLLM:  litellm,
	}, nil
}

func (s *SystemService) Close() {
	if s.agentRepo != nil {
		s.agentRepo.Close()
	}
	if s.workloadRepo != nil {
		s.workloadRepo.Close()
	}
	if s.queueRepo != nil {
		s.queueRepo.Close()
	}
	if s.llmRepo != nil {
		s.llmRepo.Close()
	}
}
