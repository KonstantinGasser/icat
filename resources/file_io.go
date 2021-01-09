package resources

import (
	"io"
	"os"
)

// NewFile returns a pointer to a file implementing the File interface
func NewFile() Resource {
	return &file{}
}

type file struct {
	file *os.File
}

// Open opens a file to a given resource
// used as wrapper for os.Open
func (f *file) Open(src string) (io.Reader, error) {
	fi, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	f.file = fi

	return fi, nil
}

// Close closes the file resource after use
func (f *file) Close() {
	f.file.Close()
}
