package types

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// TaskList Stores the list of tasks
type TaskList []*Task

// PrintTo Prints a formatted list of the tasks in the list
func (tl TaskList) PrintTo(out io.Writer) {
	for _, task := range tl {
		status := " "
		if task.Complete {
			status = "X"
		}
		fmt.Fprintf(out, "[%s] - %d: %s [%s]\n", status, task.ID, task.Description, strings.Join(task.Tags, ","))
	}
}

// WriteToPath Converts the TaskList to a JSON []byte and writes it to the given path
func (tl TaskList) WriteToPath(path string) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0600)
}

// CloseTask Closes a task from the task list
func (tl TaskList) CloseTask(id int) error {
	for _, t := range tl {
		if t.ID == id {
			t.Complete = true
			return nil
		}
	}

	return fmt.Errorf("task with ID '%d' not found", id)
}

// ReopenTask Marks a task in the task list as incomplete
func (tl TaskList) ReopenTask(id int) error {
	for _, t := range tl {
		if t.ID == id {
			if !t.Complete {
				return fmt.Errorf("task with ID '%d' is already open", id)
			}
			t.Complete = false
			return nil
		}
	}

	return fmt.Errorf("task with ID '%d' not found", id)
}
