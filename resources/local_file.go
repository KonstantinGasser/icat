package resources

import (
	"io"
	"os"
)

// LocalFile implements the Resource interface
// is able to open and close resources (files) on the local system
type LocalFile struct {
	src string
}

// Open returns the requested file
// returns the file, a cleanup function to close the resource
func (file LocalFile) Open() (io.Reader, func(), error) {
	f, err := os.Open(file.src)
	cleanup := func() {
		if f != nil {
			f.Close()
		}
	}
	return f, cleanup, err
}
