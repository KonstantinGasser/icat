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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, errP := cmd.Flags().GetString("select")
		isURL, errU := cmd.Flags().GetBool("url")
		if errP != nil || errU != nil || path == "nil" {
			return fmt.Errorf("🤨 ~ You forgort to specify the path to the image")

		}
		if err := internal.Show(os.Stdout, path, isURL); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("select", "s", "nil", "📍Select location of the image")
	showCmd.Flags().Bool("url", false, "🌐 If set fethes image from given URL")
}
