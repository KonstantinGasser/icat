package internal

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

// initLine tells iTerm to render base64 while finishLine tells when to stop
var (
	initLine   = strings.NewReader("\033]1337;File=inline=1:")
	finishLine = strings.NewReader("\a")
)

func WriteView(dst io.Writer, src string) error {

	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Could not open the file at `%s`: %s", src, err.Error())
	}
	defer f.Close()

	pipeReader, pipeWriter := io.Pipe()

	// spin up gorutine to write base64 image in pipeWriter
	go func() {
		defer pipeWriter.Close()
		// get base64 stream Writer
		b64Stream := base64.NewEncoder(base64.StdEncoding, pipeWriter)
		if _, err := io.Copy(b64Stream, f); err != nil {
			pipeWriter.CloseWithError(fmt.Errorf("Could not encode file to base64: %s", err.Error()))
			return
		}
		if err := b64Stream.Close(); err != nil {
			pipeWriter.CloseWithError(fmt.Errorf("Could not encode file to base64: %s", err.Error()))
			return
		}
	}()

	// final copy to the passed io.Writer
	if _, err := io.Copy(dst, io.MultiReader(initLine, pipeReader, finishLine)); err != nil {
		return fmt.Errorf("Could not view src: %s", err.Error())
	}
	return nil
}
