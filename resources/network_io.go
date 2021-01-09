package resources

import (
	"io"
	"net/http"
)

// NewNetConn returns a pointer to netconn implementing the NetConn interface
func NewNetConn() Resource {
	return &netconn{}
}

type netconn struct {
	resp *http.Response
}

// HTTPRequest calles the requested URL for an image. if successful returns the *Response
// else an error
func (conn *netconn) Open(src string) (io.Reader, error) {

	resp, err := http.Get(src)
	if err != nil {
		return nil, err
	}
	conn.resp = resp
	return resp.Body, nil
}

// Close closes the response body after use
func (conn *netconn) Close() {
	conn.resp.Body.Close()
}
