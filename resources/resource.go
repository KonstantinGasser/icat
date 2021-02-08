package resources

import (
	"fmt"
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

// resource can look like:
// 1) /path/to/image -> from local file system
// 		-> image_path
//		exceptions:
//			- path/to:image -> look out for colon in path - re-join path
// 2) http(s)://www.images.com/my_image -> http resource
// 		-> protocol, URL
//		exceptions:
//			- check for https or http protocol
// 3) sftp:user@host:/path/to/image -> sftp resource
//		-> protocol, user, host, path
// 		exceptions:
//			- path/to:image -> look out for colon in path - re-join path
//			- sftp:user@host:post:/path/to/image -> port can be provided check for host-colon-port
func Get(src string) (Resource, error) {
	// var protocol int
	// src: options
	// /path/to/image
	// http(s)://www.images.com/my_image
	// sftp:user@host:/path/to/image

	splitSet := strings.Split(src, ":")
	// 1) check for protocol
	switch true {
	case (splitSet[0] == "http" || splitSet[0] == "https"):
		fmt.Println("Protocol: HTTP(S)")
	case (splitSet[0] == "sftp"):
		fmt.Println("Protocol: STFP")
	default:
		fmt.Println("Protocol: LocalFile")
		// in case the path looks like (path/to:some/image)
		// re-join the path in order to not mess up the original path
		filePath := strings.Join(splitSet[0:], ":")
		return LocalFile{path: filePath}, nil
	}
	return nil, fmt.Errorf("could not find a suitable resource for: %s", src)
}
