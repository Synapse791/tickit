package types

import "testing"

func TestCloseTask(t *testing.T) {
	list := TaskList{
		&Task{
			ID:          1,
			Description: "task 1",
		},
		&Task{
			ID:          2,
			Description: "task 2",
		},
	}

	if err := list.CloseTask(2); err != nil {
		t.Error("Expected nil, got", err)
	}

	if list[0].Complete {
		t.Error("Expected false, got true. Wrong task was closed")
	}

	if !list[1].Complete {
		t.Error("Expected true, got false. Task was not closed")
	}
}

func TestReopenTask(t *testing.T) {
	list := TaskList{
		&Task{
			ID:          1,
			Description: "task 1",
			Complete:    true,
		},
		&Task{
			ID:          2,
			Description: "task 2",
			Complete:    true,
		},
	}

	if err := list.ReopenTask(2); err != nil {
		t.Error("Expected nil, got", err)
	}

	if !list[0].Complete {
		t.Error("Expected true, got false. Wrong task was reopened")
	}

	if list[1].Complete {
		t.Error("Expected false, got true. Task was not reopened")
	}
}
