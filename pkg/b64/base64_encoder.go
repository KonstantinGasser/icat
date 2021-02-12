package b64

import (
	"encoding/base64"
	"fmt"
	"io"
)

type Encoder interface {
	Stream(dst io.Writer) io.WriteCloser
	Copy(stream func(dst io.Writer) io.WriteCloser, src io.Reader) *io.PipeReader
}

type b64Encoder struct{}

func NewEncoder() Encoder {
	return &b64Encoder{}
}

// Stream takes a io.Writer and calls the NewEncoder function on base64
func (enc *b64Encoder) Stream(dst io.Writer) io.WriteCloser {
	return base64.NewEncoder(base64.StdEncoding, dst)
}

// Copy takes a function stream which can modify data in the stream.
// Copy creates a pipeReader and pipeWriter - the pipeReader gets returned.
// It spins up a goroutine to consume from the stream and pipes the data in the io.Pipe
func (enc *b64Encoder) Copy(stream func(dst io.Writer) io.WriteCloser, src io.Reader) *io.PipeReader {
	pipeR, pipeW := io.Pipe()

	// stream function can be any function returning a io.WriterCloser
	// to modify the data before it gets piped to the io.Pipe
	go func() {
		defer pipeW.Close()

		// pipe the content from the src to the encoder stream
		streamer := stream(pipeW)
		
		// read from encoder stream and pipe to pipewrtier
		if _, err := io.Copy(streamer, src); err != nil {
			pipeW.CloseWithError(fmt.Errorf("could not stream base64 content to pipeWriter: %s", err.Error()))
			return
		}
		if err := streamer.Close(); err != nil {
			pipeW.CloseWithError(fmt.Errorf("could not close io.Writer from base64.NewEncoder: %s", err.Error()))
			return
		}
	}()

	return pipeR
}
