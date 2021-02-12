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
	imageSrc := strings.Join(splitSet[0:], ":")
	switch true {
	// src of image is URL
	case (prefix == "http" || prefix == "https"):
		return NetConnHTTP{src: imageSrc}, nil
	// src of image is SFTP request
	case (prefix == "sftp"):
		sftpConn := NetConnSFTP{}
		// exclude "sftp://"
		if err := sftpConn.resolve(imageSrc[7:]); err != nil {
			return nil, err
		}
		return sftpConn, nil
	// default use local file system to open src
	default:
		return LocalFile{src: imageSrc}, nil
	}

}
