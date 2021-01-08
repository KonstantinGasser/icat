package resources

import (
	"fmt"
	"io"
	"os"
)

// File is the interface to work with the file resource
type File interface {
	Open(src string) (*os.File, error)
	Copy(dst io.Writer, src io.Reader) error
	MultiCopy(dst io.Writer, src ...io.Reader) error
}

type file struct{}

// NewResource returns a pointer to a file implementing the File interface
func NewResource() File {
	return &file{}
}

// Open opens a file to a given resource
// used as wrapper for os.Open
func (f *file) Open(src string) (*os.File, error) {
	return os.Open(src)
}

// Copy wrappes the io.Copy function to copy data from one src to a given destination
func (f *file) Copy(dst io.Writer, src io.Reader) error {
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("cloud not copy data from %v to %v: %s", src, dst, err.Error())
	}
	return nil
}

// MultiCopy is used to copy data from a multiple-reader to one destination writer
func (f *file) MultiCopy(dst io.Writer, src ...io.Reader) error {

	if _, err := io.Copy(dst, io.MultiReader(src...)); err != nil {
		return fmt.Errorf("could not copy data from %v to given src: %s", dst, err.Error())
	}
	return nil
}
