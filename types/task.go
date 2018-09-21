package types

// Task represents a task in the list
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Tags        Tags   `json:"tags"`
}
