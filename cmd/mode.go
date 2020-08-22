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

	"github.com/spf13/cobra"
)

// modeCmd represents the mode command
var modeCmd = &cobra.Command{
	Use:   "mode",
	Short: "🤔Tell me what you want to do with the image",
	Long: `you can print the image in RGB to the iTerm2 session you are in
or display the raw base64 encdoing of the image. Of cause you can also write the base64 encoded image to a file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mode called")
	},
}

func init() {
	rootCmd.AddCommand(modeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
