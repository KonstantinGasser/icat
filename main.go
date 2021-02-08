package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/KonstantinGasser/icat/resources"
)

const (
	// ENV_iTerm_KEY environment variable to tell which terminal app is used
	ENV_iTerm_KEY = "TERM_PROGRAM"
	// ENV_iTerm_VALUE value telling iTerm terminal is used to execute icat command
	ENV_iTerm_VALUE = "iTerm.app"
)

var (
	iTermCmdStartRender = strings.NewReader("\033]1337;File=inline=1:")
	iTermCmdStopRender  = strings.NewReader("\a\n") // \n to force terminal to do new line
)

func main() {

	// imagePath := flag.String("path", "", "path to image file")
	// determine which resource needs to be called (HTTP, SFTP, File(file is default if others not set))
	// isURL := flag.Bool("url", false, "use if path is a URL to an image")
	// isSFTP := flag.Bool("sftp", false, "use if image needs to be accessed on remote server (use user@host:/path/to/image")

	// determine which mode to apply on file
	// displayImage := flag.Bool("show", true, "displays image in iTerm window")
	displayBase64 := flag.Bool("base64", false, "display base64 encoding of image")
	writeBase64 := flag.String("out", "", "write base64 encoding of image to file")

	imagePath := os.Args[1]

	flag.Parse()

	// var resource resources.Resource
	resource, err := resources.Get(imagePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// check which mode to use (display or raw base64)
	// var mode int
	switch true {
	case *displayBase64: // write base64 of image to os.Stdout
		fallthrough
	case (*writeBase64 != ""): // write base64 of image to given output file
	default: // display image in iTerm window
		if err := showImage(resource); err != nil {
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
