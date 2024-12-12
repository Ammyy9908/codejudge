package models

// Submission represents a code submission
type Submission struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Language string `json:"language"`
}
