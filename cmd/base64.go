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
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	src string
	out string
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(src)
		if err != nil {
			fmt.Printf("Could not open src file: %s", err.Error())
			return
		}
		defer f.Close()

		var w = os.Stdout
		if out != "" {
			w, err = os.Create(out)
			if err != nil {
				fmt.Printf("Could not create out file: %s", err.Error())
				return
			}
			defer w.Close()
		}
		enc := base64.NewEncoder(base64.StdEncoding, w)
		if _, err := io.Copy(enc, f); err != nil {
			fmt.Printf("Cloud not write to writer: %s", err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.Flags().StringVarP(&src, "src", "s", "", "path from source file")
	base64Cmd.Flags().StringVarP(&out, "out", "o", "", "write base64 of image to given output path")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// base64Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// base64Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
