package stream

import (
	"encoding/base64"
	"fmt"
	"io"
)

// Base64 returns a new base64.NewEncoder stream
func Base64(dst io.Writer) io.WriteCloser {
	return base64.NewEncoder(base64.StdEncoding, dst)
}

// Pipe takes in a src to read from and a function which can modify the data (here a base64.NewEncoder stream)
// the streamed data gets written into the pipeWriter from where it can be consumed by the returned pipeReader
func Pipe(src io.Reader, stream func(streamDst io.Writer) io.WriteCloser) *io.PipeReader {
	pipeReader, pipeWriter := io.Pipe()

	// copy data into pipeWriter via go-routine
	// since as pointed out by the docs io.Pipe creates
	// a synchronous in-memory pipe
	go func() {
		defer pipeWriter.Close()

		// call the passed stream function
		// set the output src to be pipeWriter
		encodedStream := stream(pipeWriter)
		if _, err := io.Copy(encodedStream, src); err != nil {
			pipeWriter.CloseWithError(fmt.Errorf("cloud not copy data from passed stream to pipeWriter: %v", err))
			return
		}
		if err := encodedStream.Close(); err != nil {
			pipeWriter.CloseWithError(fmt.Errorf("could not close io.WriteCloser from stream: %v", err))
			return
		}
	}()
	return pipeReader
}
