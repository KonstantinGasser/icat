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

	"github.com/KonstantinGasser/icat/internal"
	"github.com/spf13/cobra"
)

var (
	path     string
	isBase64 bool
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "use view to render a image on the command line",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if isBase64 {
			if err := internal.RenderFromBase64(os.Stdout, path); err != nil {
				fmt.Errorf("Cloud not view image: %s", err.Error())
				return
			}
			return
		}

		if err := internal.WriteView(os.Stdout, path); err != nil {
			fmt.Printf("Could not view image: %s", err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().StringVarP(&path, "src", "s", "<not given>", "path to image")
	viewCmd.Flags().BoolVarP(&isBase64, "base64", "b", false, "use if file is base64 encoded")
}
