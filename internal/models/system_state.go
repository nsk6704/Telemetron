// internal/models/system_state.go
package models

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
