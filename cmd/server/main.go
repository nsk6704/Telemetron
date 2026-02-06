package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"telemetron/internal/models"
	"telemetron/pkg/config"
	"telemetron/pkg/logger"

	"go.uber.org/zap"
)

func systemStateHandler(w http.ResponseWriter, r *http.Request) {
	response := models.SystemState{
		ID: "system-1",
		Agents: []models.Agent{
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
					UpdatedAt: "2025-01-15T10:30:00Z",
				},
			},
		},
		Workload: []models.Workload{
			{
				DeploymentName: "agent-deployment-1",
				MaxPods:        10,
				PodMaxRAM:      "2Gi",
				PodMaxCPU:      "1000m",
				Live: models.LiveWorkload{
					ActivePods: 3,
					UpdatedAt:  "2025-01-15T10:30:00Z",
				},
				Pods: []models.Pod{
					{PodID: "pod-1", CPU: 0.5, Memory: 1024, Status: "running"},
					{PodID: "pod-2", CPU: 0.3, Memory: 512, Status: "running"},
					{PodID: "pod-3", CPU: 0.2, Memory: 256, Status: "running"},
				},
			},
		},
		Queues: []models.Queue{
			{
				Name:      "default",
				UpdatedAt: "2025-01-15T10:30:00Z",
				Tasks: []models.QueueTask{
					{ID: "task-1", Priority: models.Priority{Level: "high"}, SubmittedAt: "2025-01-15T10:25:00Z"},
					{ID: "task-2", Priority: models.Priority{Level: "medium"}, SubmittedAt: "2025-01-15T10:28:00Z"},
				},
			},
		},
		LiteLLM: []models.LiteLLM{
			{Model: "gpt-4", Provider: "openai", TPM: 45000, RPM: 200, TPMMax: 90000, RPMMax: 3500, PaymentType: "pay-per-request"},
			{Model: "gpt-3.5-turbo", Provider: "openai", TPM: 120000, RPM: 3400, TPMMax: 240000, RPMMax: 3500, PaymentType: "pay-per-request"},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	cfg := config.Load()

	if err := logger.Init(cfg.LogLevel); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Close()

	http.HandleFunc("/system/state", systemStateHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Telemetron API - visit /system/state"))
	})

	addr := ":" + cfg.ServerPort
	logger.Log.Info("Starting Telemetron server", zap.String("address", addr))

	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Log.Fatal("Server error", zap.Error(err))
	}
}
