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
		{kind: LocalFile{}, src: "./path/to/image", excpected: "./path/to/image"},
		{kind: LocalFile{}, src: "/path/to:some/image", excpected: "/path/to:some/image"},
		{kind: NetConnHTTP{}, src: "https://www.pinclipart.com/picdir/middle/571-5718168_go-gopher-stickers-clipart.png", excpected: "https://www.pinclipart.com/picdir/middle/571-5718168_go-gopher-stickers-clipart.png"},
		{kind: NetConnSFTP{}, src: "sftp://user@hostname:224/image/path", excpected: struct{ user, host, port, src string }{user: "user", host: "hostname", port: "224", src: "/image/path"}},
		{kind: NetConnSFTP{}, src: "sftp://user@192.168.0.232:224/image/path", excpected: struct{ user, host, port, src string }{user: "user", host: "192.168.0.232", port: "224", src: "/image/path"}},
	}

	for _, tc := range tt {
		resource, err := New(tc.src)
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
