package resources

import (
	"fmt"
	"io"
)

type stdout struct{}

func NewStdOut() StdOut {
	return &stdout{}
}

// Copy wrappes the io.Copy function to copy data from one src to a given destination
func (stdout *stdout) Copy(dst io.Writer, src io.Reader) error {
	if _, err := io.Copy(dst, src); err != nil {
		return fmt.Errorf("cloud not copy data from %v to %v: %s", src, dst, err.Error())
	}
	return nil
}

// MultiCopy is used to copy data from a multiple-reader to one destination writer
func (stdout *stdout) MultiCopy(dst io.Writer, src ...io.Reader) error {

	if _, err := io.Copy(dst, io.MultiReader(src...)); err != nil {
		return fmt.Errorf("could not copy data from %v to given src: %s", dst, err.Error())
	}
	return nil
}
