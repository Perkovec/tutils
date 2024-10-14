package tutils

// Size is the size of the terminal window
type Size struct {
	Rows    int
	Columns int
}

// GetSize returns the size of the terminal window
// TODO: Add support for other platforms like windows
func GetSize() (Size, error) {
	return getSize()
}
