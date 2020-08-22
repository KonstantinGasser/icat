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

// rawCmd represents the raw command
var rawCmd = &cobra.Command{
	Use:   "raw",
	Short: "Outputs the raw base64 encoded img to the command-line",
	Long:  `not implemented yet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := cmd.Flags().GetString("select")
		if err != nil || path == "nil" {
			return fmt.Errorf("🤨 ~ You forgort to specify the path to the image")
			os.Exit(1)
		}
		if err := internal.Raw(os.Stdout, path); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rawCmd)
	rawCmd.Flags().StringP("select", "s", "nil", "📍Select location of the image")
}
