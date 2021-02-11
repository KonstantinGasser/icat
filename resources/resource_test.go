package resources

import (
	"reflect"
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
	}

	for _, tc := range tt {
		resource, err := Get(tc.src)
		is.NoErr(err)

		if reflect.TypeOf(tc.kind) == reflect.TypeOf(LocalFile{}) {
			lf, ok := resource.(LocalFile)
			if !ok {
				t.Errorf("returned resource dose not match expected type of test-case")
			}
			is.Equal(lf.src, tc.excpected)
		}

	}
}
