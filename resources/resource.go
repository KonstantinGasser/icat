package resources

import (
	"io"
	"strings"
)

const (
	// protocolLOCAL file located on local file system
	protocolLOCAL = iota
	// protocolHTTP file located on www site
	protocolHTTP
	// protocolSFTP file located on remote server
	protocolSFTP
)

// Resource represents a file, http or sftp connection
type Resource interface {
	Open() (io.Reader, func(), error)
}

// New returns based on the src the correct resource to handle the request
// support for http/https, sftp, local-file-system (default)
func New(src string) (Resource, error) {

	splitSet := strings.Split(src, ":")
	// http/https, sftp, -
	prefix := splitSet[0]
	mergedSrc := strings.Join(splitSet[0:], ":")
	switch true {
	// src of image is URL
	case (prefix == "http" || prefix == "https"):
		return NetConnHTTP{src: mergedSrc}, nil
	// src of image is SFTP request
	case (prefix == "sftp"):
		panic("get resource SFTP - not implemented")
	// default use local file system to open image
	default:
		return LocalFile{src: mergedSrc}, nil
	}

}
