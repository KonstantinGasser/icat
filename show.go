package main

import (
	"fmt"
	"os"

	"github.com/KonstantinGasser/icat/pkg/resources"
	"github.com/KonstantinGasser/icat/pkg/stream"
)

func printImage(resource resources.Resource) error {
	// request to open file from resource
	content, teardown, err := resource.Open()
	// teardown holds resource specific instructions on how to close
	// and clean up the opened resource
	defer teardown()
	if err != nil {
		return fmt.Errorf("could not open resource: %v", err)
	}
	// pipe resource content through encoder-stream and pipeWriter
	pipeR := stream.Pipe(content, stream.Base64)
	if _, err := Copy(os.Stdout, iTermCmdStartRender, pipeR, iTermCmdStopRender); err != nil {
		return fmt.Errorf("could not copy content to os.Stdout: %v", err)
	}
	return nil
}
