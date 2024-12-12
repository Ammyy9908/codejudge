package execution

// Submission represents a task consumed from the submission queue
type Submission struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
}

// ExecutionResult represents the output of the code execution
type ExecutionResult struct {
	ID     string `json:"id"`
	Output string `json:"output"`
	Error  string `json:"error"`
}
