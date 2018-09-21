package types

import (
	"encoding/json"
	"io/ioutil"
)

// TaskList Stores the list of tasks
type TaskList []*Task

// WriteToPath Converts the TaskList to a JSON []byte and writes it to the given path
func (t TaskList) WriteToPath(path string) error {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0600)
}
