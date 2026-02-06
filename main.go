package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SystemState struct {
	ID       string     `json:"id"`
	Agents   []Agent    `json:"agents"`
	Workload []Workload `json:"workload"`
	Queues   []Queue    `json:"queues"`
	LiteLLM  []LiteLLM  `json:"litellm"`
}

type Agent struct {
	Name                   string   `json:"name"`
	Description            string   `json:"description"`
	MaxParallelInvocations int      `json:"max_parallel_invocations"`
	DeploymentName         string   `json:"deployment_name"`
	Models                 []string `json:"models"`
	Activity               Activity `json:"activity"`
}
type Activity struct {
	ActiveTaskIDs []TaskStatus `json:"active_task_ids"`
	UpdatedAt     string       `json:"updated_at"`
}

type TaskStatus struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type Workload struct {
	DeploymentName string       `json:"deployment_name"`
	MaxPods        int          `json:"max_pods"`
	PodMaxRAM      string       `json:"pod_max_ram"`
	PodMaxCPU      string       `json:"pod_max_cpu"`
	Live           LiveWorkload `json:"live"`
	Pods           []Pod        `json:"pods"`
}

type LiveWorkload struct {
	ActivePods int    `json:"active_pods"`
	UpdatedAt  string `json:"updated_at"`
}

type Pod struct {
	PodID  string  `json:"pod_id"`
	CPU    float64 `json:"cpu"`
	Memory int     `json:"memory"`
	Status string  `json:"status"`
}

type Queue struct {
	Name      string      `json:"name"`
	UpdatedAt string      `json:"updated_at"`
	Tasks     []QueueTask `json:"tasks"`
}

type QueueTask struct {
	ID          string   `json:"id"`
	Priority    Priority `json:"priority"`
	SubmittedAt string   `json:"submitted_at"`
}

type Priority struct {
	Level string `json:"level"`
}

type LiteLLM struct {
	Model       string `json:"model"`
	Provider    string `json:"provider"`
	TPM         int    `json:"tpm"`
	RPM         int    `json:"rpm"`
	TPMMax      int    `json:"tpm_max"`
	RPMMax      int    `json:"rpm_max"`
	PaymentType string `json:"payment_type"`
}

func main() {
	http.HandleFunc("/system/state", systemStateHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Telemetron API - visit /system/state"))
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func systemStateHandler(w http.ResponseWriter, r *http.Request) {
	response := SystemState{
		ID: "system-1",
		Agents: []Agent{
			{
				Name:                   "agent-1",
				Description:            "Data processing agent",
				MaxParallelInvocations: 5,
				DeploymentName:         "agent-deployment-1",
				Models:                 []string{"gpt-4", "gpt-3.5-turbo"},
				Activity: Activity{
					ActiveTaskIDs: []TaskStatus{
						{ID: "task-1", Status: "running"},
						{ID: "task-2", Status: "pending"},
					},
					UpdatedAt: "2025-01-15T10:30:00Z",
				},
			},
		},
		Workload: []Workload{
			{
				DeploymentName: "agent-deployment-1",
				MaxPods:        10,
				PodMaxRAM:      "2Gi",
				PodMaxCPU:      "1000m",
				Live: LiveWorkload{
					ActivePods: 3,
					UpdatedAt:  "2025-01-15T10:30:00Z",
				},
				Pods: []Pod{
					{PodID: "pod-1", CPU: 0.5, Memory: 1024, Status: "running"},
					{PodID: "pod-2", CPU: 0.3, Memory: 512, Status: "running"},
					{PodID: "pod-3", CPU: 0.2, Memory: 256, Status: "running"},
				},
			},
		},
		Queues: []Queue{
			{
				Name:      "default",
				UpdatedAt: "2025-01-15T10:30:00Z",
				Tasks: []QueueTask{
					{ID: "task-1", Priority: Priority{Level: "high"}, SubmittedAt: "2025-01-15T10:25:00Z"},
					{ID: "task-2", Priority: Priority{Level: "medium"}, SubmittedAt: "2025-01-15T10:28:00Z"},
				},
			},
		},
		LiteLLM: []LiteLLM{
			{Model: "gpt-4", Provider: "openai", TPM: 45000, RPM: 200, TPMMax: 90000, RPMMax: 3500, PaymentType: "pay-per-request"},
			{Model: "gpt-3.5-turbo", Provider: "openai", TPM: 120000, RPM: 3400, TPMMax: 240000, RPMMax: 3500, PaymentType: "pay-per-request"},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
