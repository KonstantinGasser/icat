package resources

import (
	"io"
	"os"
)

// LocalFile implements the Resource interface
// is able to open and close resources (files) on the local system
type LocalFile struct {
	path string
}

// Open returns the requested file
// returns the file, a cleanup function to close the resource
func (file LocalFile) Open() (io.Reader, func(), error) {
	f, err := os.Open(file.path)
	return f, func() { f.Close() }, err
}
