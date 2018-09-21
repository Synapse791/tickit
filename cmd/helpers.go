package cmd

import (
	"os/user"
	"path/filepath"
)

// UserHomeDir Returns the current user's home directory as a string
func UserHomeDir() string {
	current, _ := user.Current()
	return current.HomeDir
}

// UserHomeFile returns the provided path appended to the current user's home directory
func UserHomeFile(path string) string {
	return filepath.Join(UserHomeDir(), path)
}
