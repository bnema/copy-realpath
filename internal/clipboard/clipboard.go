package clipboard

// Copier defines the interface for clipboard operations.
type Copier interface {
	// Copy copies the given text to the system clipboard.
	Copy(text string) error

	// Available returns true if this clipboard backend is available.
	Available() bool

	// Name returns the name of this clipboard backend.
	Name() string
}
