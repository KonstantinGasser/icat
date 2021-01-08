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
	"strings"

	"github.com/spf13/cobra"
)

var (
	out string
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "convert image file to base64 text files",
	Long:  `If the --out is not given the base64 of the file gets printed to the command line`,
	Run: func(cmd *cobra.Command, args []string) {
		// check for src path
		if len(os.Args) < 2 {
			fmt.Printf("you forgot to provide a src path to an image or base64 text file")
			return
		}

		src := os.Args[2]

		// fetch content from image
		target, err := resource.Open(src)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
		defer target.Close()

		// encode image to bas64 and pipe it to the io.Pipe
		// encoder.Copy triggers a goroutine to pipe the data from the encoder stream to
		// the pipeWriter
		pipeReader := encoder.Copy(encoder.Stream, target)

		// check if output path is given
		if out != "" {
			outF, err := os.Create(out)
			if err != nil {
				fmt.Printf("cloud not create out put file at:%s :%s\n", out, err.Error())
				return
			}
			// copy base64 of an image to a given output file
			if err := resource.Copy(outF, pipeReader); err != nil {
				fmt.Printf("could not copy content to file: %s\n", err.Error())
				return
			}
		}

		// no output path given: print to terminal
		if out == "" {
			if err := resource.MultiCopy(os.Stdout, pipeReader, strings.NewReader("\n")); err != nil {
				fmt.Printf("could not copy content to os.Stdout: %s\n", err.Error())
				return
			}
			// force new line after content is printed
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().StringVarP(&out, "out", "o", "", "write base64 of image to given output path")
}
