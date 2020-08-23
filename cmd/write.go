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
	"time"

	"github.com/KonstantinGasser/icat/internal"
	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Writes the base64 of the image to the givien output file",
	Long:  `not implemented yet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		from, err := cmd.Flags().GetString("select")
		to, err := cmd.Flags().GetString("output")
		isURL, err := cmd.Flags().GetBool("url")
		if from == "nil" || err != nil {
			return fmt.Errorf("🤨 ~ You messed up the command: %v", err.Error())
		}
		if err := internal.Write(from, to, isURL); err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "🥳 ~ File with base64 of %s saved at %s", from, to)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
	writeCmd.Flags().StringP("select", "s", "nil", "📍Select location of the image")

	// default for outtput
	outputDefault := strings.Join([]string{"./icat-", time.Now().String(), ".txt"}, "")
	writeCmd.Flags().StringP("output", "o", outputDefault, "💾 Location of the outputfile including the file name,\nif not given icat-currentTime is the name of the file")
	writeCmd.Flags().Bool("url", false, "🌐 If set fethes image from given URL")
}
