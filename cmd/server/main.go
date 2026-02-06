package main

// @title Telemetron API
// @version 1.0
// @description Telemetry and agent orchestration API.
// @host localhost:8080
// @BasePath /

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "telemetron/docs" // Import generated docs
	"telemetron/internal/repositories"
	"telemetron/internal/services"
	"telemetron/pkg/config"
	"telemetron/pkg/logger"

	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

// @Summary Get system state
// @Description Returns the current state of agents, workloads, queues, and LiteLLM models
// @Tags system
// @Produce json
// @Success 200 {object} models.SystemState
// @Failure 500 {string} string "Internal server error"
// @Router /system/state [get]
func systemStateHandler(systemService *services.SystemService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state, err := systemService.GetSystemState()
		if err != nil {
			logger.Log.Error("Failed to get system state", zap.Error(err))
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(state); err != nil {
			logger.Log.Error("Failed to encode response", zap.Error(err))
		}
	}
}

func main() {
	cfg := config.Load()

	if err := logger.Init(cfg.LogLevel); err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Close()

	// Initialize repositories
	agentRepo := repositories.NewMockAgentRepository()
	workloadRepo := repositories.NewMockWorkloadRepository()
	queueRepo := repositories.NewMockQueueRepository()
	llmRepo := repositories.NewMockLiteLLMRepository()

	// Initialize service
	systemService := services.NewSystemService(agentRepo, workloadRepo, queueRepo, llmRepo)
	defer systemService.Close()

	// Setup handlers
	http.HandleFunc("/system/state", systemStateHandler(systemService))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Telemetron API - visit /system/state or /swagger/"))
	})

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	addr := ":" + cfg.ServerPort
	logger.Log.Info("Starting Telemetron server", zap.String("address", addr))

	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Log.Fatal("Server error", zap.Error(err))
	}
}
