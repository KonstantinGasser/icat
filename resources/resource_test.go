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
		excpected string
	}{
		{kind: LocalFile{}, src: "./path/to/image", excpected: "./path/to/image"},
		{kind: LocalFile{}, src: "/path/to:some/image", excpected: "/path/to:some/image"},
		{kind: NetConnHTTP{}, src: "https://www.pinclipart.com/picdir/middle/571-5718168_go-gopher-stickers-clipart.png", excpected: "https://www.pinclipart.com/picdir/middle/571-5718168_go-gopher-stickers-clipart.png"},
	}

	for _, tc := range tt {
		resource, err := New(tc.src)
		is.NoErr(err)

		switch tc.kind.(type) {
		case LocalFile:
			lf := resource.(LocalFile)
			is.Equal(lf.src, tc.excpected)
		case NetConnHTTP:
			conn := resource.(NetConnHTTP)
			is.Equal(conn.src, tc.excpected)
		default:
			t.Errorf("no valid resource returned by resources.New(src)")

		}
	}
}
