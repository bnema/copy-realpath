package path

import (
	"fmt"
	"os"
	"path/filepath"
)

// Resolver handles path resolution to absolute paths.
type Resolver struct{}

// NewResolver creates a new path resolver.
func NewResolver() *Resolver {
	return &Resolver{}
}

// Resolve converts the given input to an absolute path.
// If input is empty or ".", it returns the current working directory.
// Otherwise, it resolves the input relative to the current working directory.
func (r *Resolver) Resolve(input string) (string, error) {
	// Handle empty input or "." - return current directory
	if input == "" || input == "." {
		cwd, err := os.Getwd()
		if err != nil {
			return "", fmt.Errorf("failed to get current directory: %w", err)
		}
		return cwd, nil
	}

	// If already absolute, clean and return
	if filepath.IsAbs(input) {
		return filepath.Clean(input), nil
	}

	// Resolve relative path against current directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	absPath := filepath.Join(cwd, input)
	return filepath.Clean(absPath), nil
}
