package resources

import (
	"fmt"
	"io"
)

type Resource interface {
	Open(src string) (io.Reader, error)
	Close()
}

func NewResource(kind string) (Resource, error) {
	switch kind {
	case "remote":
		return &sftpclient{}, nil
	case "net":
		return &netconn{}, nil
	case "local":
		return &file{}, nil
	default:
		return nil, fmt.Errorf("unknow resource type: %s", kind)
	}
}

type StdOut interface {
	Copy(dst io.Writer, src io.Reader) error
	MultiCopy(dst io.Writer, src ...io.Reader) error
}
