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

	fmt.Println("Executing code...", submission.Language)
	fmt.Println("Code:", submission.Code)

	switch submission.Language {
	case "python":
		cmd = exec.Command("python3", "-c", submission.Code)
		fmt.Println("Running Python code:", cmd.String())

	case "go":
		fmt.Println("Running Go code...")
		tmpFile, err := os.CreateTemp("", "*.go")
		if err != nil {
			return execution.ExecutionResult{
				ID:    submission.ID,
				Error: fmt.Sprintf("Failed to create temporary file: %v", err),
			}
		}
		defer os.Remove(tmpFile.Name())

		if _, err := tmpFile.WriteString(submission.Code); err != nil {
			return execution.ExecutionResult{
				ID:    submission.ID,
				Error: fmt.Sprintf("Failed to write code to file: %v", err),
			}
		}
		tmpFile.Close()
		fmt.Printf("Temporary file created: %s\n", tmpFile.Name())
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

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Execution error: %v\n", err)
		fmt.Printf("Stderr: %s\n", stderr.String())
		return execution.ExecutionResult{
			ID:    submission.ID,
			Error: stderr.String(),
		}
	}

	fmt.Printf("Execution successful. Stdout: %s\n", out.String())
	return execution.ExecutionResult{
		ID:     submission.ID,
		Output: out.String(),
	}
}
