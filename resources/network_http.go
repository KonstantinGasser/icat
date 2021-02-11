package resources

import (
	"fmt"
	"io"
	"net/http"
)

// NetConnHTTP represents a HTTP connection to locate some
// website's image
type NetConnHTTP struct {
	src string
}

// Open calls a given url returns the response.Body and a cleanup function
// to close the resource
func (conn NetConnHTTP) Open() (io.Reader, func(), error) {
	resp, err := http.Get(conn.src)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()

		}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, cleanup, fmt.Errorf("URL: %s - returned status: %d", conn.src, resp.StatusCode)
	}
	return resp.Body, cleanup, nil
}
