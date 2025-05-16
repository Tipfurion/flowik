package workflow

type Job struct {
	ID          string            `json:"id"`
	Type        string            `json:"type"` // Always "job"
	Name        *string           `json:"name,omitempty"`
	Image       string            `json:"image"`
	Command     *string           `json:"command,omitempty"`
	Args        map[string]string `json:"args,omitempty"`
	Timeout     *int              `json:"timeout,omitempty"`
	Environment map[string]string `json:"environment,omitempty"`
}
	