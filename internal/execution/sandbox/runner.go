package sandbox

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	execution "github.com/ammyy9908/codejudge/internal/execution/models"
)

// ExecuteCode runs the submitted code and returns the result
func ExecuteCode(submission execution.Submission) execution.ExecutionResult {
	var cmd *exec.Cmd

	// Handle the code execution based on the language
	switch submission.Language {
	case "python":
		cmd = exec.Command("python3", "-c", submission.Code)

	case "go":
		fmt.Println("Running Go code...")
		// Write Go code to a temporary file
		tmpFile, err := os.CreateTemp("", "*.go")
		if err != nil {
			return execution.ExecutionResult{
				ID:    submission.ID,
				Error: fmt.Sprintf("Failed to create temporary file: %v", err),
			}
		}
		defer os.Remove(tmpFile.Name()) // Ensure the file is deleted after execution

		// Write the code to the temporary file
		if _, err := tmpFile.WriteString(submission.Code); err != nil {
			return execution.ExecutionResult{
				ID:    submission.ID,
				Error: fmt.Sprintf("Failed to write code to file: %v", err),
			}
		}

		// Close the file before execution
		tmpFile.Close()

		// Use go run to execute the file
		cmd = exec.Command("go", "run", tmpFile.Name())

	default:
		return execution.ExecutionResult{
			ID:    submission.ID,
			Error: "Unsupported language",
		}
	}

	// Capture output and errors
	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return execution.ExecutionResult{
			ID:    submission.ID,
			Error: stderr.String(),
		}
	}

	return execution.ExecutionResult{
		ID:     submission.ID,
		Output: out.String(),
	}
}
