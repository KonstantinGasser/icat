package resources

import (
	"testing"

	"github.com/matryer/is"
)

// TestGet tests the evaluation of the given src string
// in order to get the correct resource required to open the file
func TestGet(t *testing.T) {
	is := is.New(t)

	tt := []struct {
		kind      Resource
		src       string
		excpected interface{}
	}{
		{
			kind:      LocalFile{},
			src:       "./path/to/image",
			excpected: "./path/to/image"},
		{
			kind:      LocalFile{},
			src:       "/path/to:some/image",
			excpected: "/path/to:some/image",
		},
		{
			kind:      NetConnHTTP{},
			src:       "https://www.pinclipart.com/picdir/middle/571-5718168_go-gopher-stickers-clipart.png",
			excpected: "https://www.pinclipart.com/picdir/middle/571-5718168_go-gopher-stickers-clipart.png",
		},
		{
			kind: NetConnSFTP{},
			src:  "sftp://user@hostname:224/image/path",
			excpected: struct{ user, host, port, src string }{
				user: "user", host: "hostname", port: "224", src: "/image/path",
			},
		},
		{
			kind: NetConnSFTP{},
			src:  "sftp://user@192.168.0.232:224/image/path",
			excpected: struct{ user, host, port, src string }{
				user: "user", host: "192.168.0.232", port: "224", src: "/image/path",
			},
		},
		{
			kind: NetConnSFTP{},
			src:  "sftp://user@192.168.0.232/image/path",
			excpected: struct{ user, host, port, src string }{
				user: "user", host: "192.168.0.232", port: "22", src: "/image/path",
			},
		},
		// these are all valid URLs according to [RFC 3986]
		{
			kind:      NetConnHTTP{},
			src:       "http://:pass@example.org:123/some/directory/file.html?query=string#fragment",
			excpected: "http://:pass@example.org:123/some/directory/file.html?query=string#fragment",
		},
		{
			kind:      NetConnHTTP{},
			src:       "http://[2010:836B:4179::836B:4179]",
			excpected: "http://[2010:836B:4179::836B:4179]",
		},
		{
			kind:      NetConnHTTP{},
			src:       "http://i.xss.com\\www.example.org/some/directory/file.html?query=string#fragment",
			excpected: "http://i.xss.com\\www.example.org/some/directory/file.html?query=string#fragment",
		},
	}

	for _, tc := range tt {
		resource, err := New(tc.src)
		if resource == nil {
			t.Fatal("resources.New returned nil")
		}
		is.NoErr(err)

		switch tc.kind.(type) {
		case LocalFile:
			lf := resource.(LocalFile)
			is.Equal(lf.src, tc.excpected.(string))
		case NetConnHTTP:
			conn := resource.(NetConnHTTP)
			is.Equal(conn.src, tc.excpected.(string))
		case NetConnSFTP:
			conn := resource.(NetConnSFTP)
			is.Equal(conn.user, tc.excpected.(struct{ user, host, port, src string }).user)
			is.Equal(conn.host, tc.excpected.(struct{ user, host, port, src string }).host)
			is.Equal(conn.port, tc.excpected.(struct{ user, host, port, src string }).port)
			is.Equal(conn.src, tc.excpected.(struct{ user, host, port, src string }).src)
		default:
			t.Errorf("no valid resource returned by resources.New(src)")
		}
	}
}
