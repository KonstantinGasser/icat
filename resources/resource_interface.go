package resources

import (
	"io"
)

type Resource interface {
	Open(src string) (io.Reader, error)
	Close()
}

type StdOut interface {
	Copy(dst io.Writer, src io.Reader) error
	MultiCopy(dst io.Writer, src ...io.Reader) error
}
