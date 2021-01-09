/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/KonstantinGasser/icat/resources"
	"github.com/spf13/cobra"
)

var (
	// user     string
	// host     string
	// port     int
	isBase64 bool
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "view allows you to render images in your iTerm terminal",
	Long:  "view allows you to render images in your iTerm terminal. The file can either be an image or a text file containing base64",
	Run: func(cmd *cobra.Command, args []string) {
		var src string
		var resource resources.Resource

		// check for src path
		if len(os.Args) < 2 {
			fmt.Printf("you forgot to provide a src path to an image or base64 text file")
			return
		}

		src = os.Args[2]

		// TODO: change code to be clear!
		// determine which resource is required
		if string(src[0:4]) == "http" || string(src[0:5]) == "https" {
			resource = resources.NewNetConn()
		} else {
			resource = resources.NewFile()
		}
		defer resource.Close()

		// fetch content from image
		target, err := resource.Open(src)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}

		// content of file is already in base64
		if isBase64 {
			if err := stdout.MultiCopy(os.Stdout, iTermCmdStartRender, target, iTermCmdStopRender); err != nil {
				fmt.Printf("%s\n", err.Error())
				return
			}
		}

		// encode image to bas64 and pipe it to the io.Pipe
		// encoder.Copy triggers a goroutine to pipe the data from the encoder stream to
		// the pipeWriter
		pipeReader := encoder.Copy(encoder.Stream, target)

		// copy starting and stop command and base64 image conntent
		// to os.Stdout (iTerm window)
		if err := stdout.MultiCopy(os.Stdout, iTermCmdStartRender, pipeReader, iTermCmdStopRender); err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
	},
}

func init() {

	rootCmd.AddCommand(viewCmd)

	// flags for sftp access
	// viewCmd.Flags().StringVarP(&user, "user", "u", "", "user name for remote server")
	// viewCmd.Flags().StringVarP(&host, "host", "H", "", "hostname of remote server")
	// viewCmd.Flags().IntVarP(&port, "port", "p", 22, "sft port for remote server (default 22)")
	viewCmd.Flags().BoolVarP(&isBase64, "base64", "b", false, "use if image is in base64")
}
