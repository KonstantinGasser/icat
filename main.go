package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/KonstantinGasser/icat/pkg/resources"
)

const (
	// icat version
	version = "0.0.1"
	// ENV_iTerm_KEY environment variable to tell which terminal app is used
	ENV_iTerm_KEY = "TERM_PROGRAM"
	// ENV_iTerm_VALUE value telling iTerm terminal is used to execute icat command
	ENV_iTerm_VALUE = "iTerm.app"
)

var (
	// tell iTerm to start render base64 content
	iTermCmdStartRender = strings.NewReader("\033]1337;File=inline=1:")
	// tell iTerm to stop render base64 content
	iTermCmdStopRender = strings.NewReader("\a\n") // \n to force terminal to do new line
)

func main() {

	displayBase64 := flag.Bool("base64", false, "display base64 encoding of image")
	writeBase64 := flag.String("out", "", "write base64 encoding of image to file")
	flag.Parse()

	if len(os.Args) <= 1 {
		fmt.Println("please provide either a path to a local file, a http(s) like to an image or a smpt conntection to load an image from a remote server")
		os.Exit(1)
	}
	if os.Args[1] == "version" {
		fmt.Printf("icat version: %s\n", version)
		return
	}

	imagePath := os.Args[len(os.Args)-1]
	resource, err := resources.New(imagePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// check which mode to use (display or raw base64)
	switch true {
	case *displayBase64: // write base64 of image to os.Stdout
		fallthrough
	case (*writeBase64 != ""): // write base64 of image to given output file
	default: // display image in iTerm window
		if err := printImage(resource); err != nil {
			fmt.Println(err)
			return
		}
	}

}

// isiTerm returns true if used terminal is iTerm.app
// images can only be displayed in iTerm -required for showImage(resource)
func isiTerm() bool {
	return os.Getenv(ENV_iTerm_KEY) != ENV_iTerm_VALUE
}
