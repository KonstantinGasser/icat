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

	"github.com/KonstantinGasser/icat/b64"
	"github.com/KonstantinGasser/icat/resources"
	"github.com/spf13/cobra"
)

var version = "1.0.0"

const (
	ENV_iTerm_KEY   = "TERM_PROGRAM"
	ENV_iTerm_VALUE = "iTerm.app"
)

var (
	resource            resources.File
	encoder             b64.Encoder
	iTermCmdStartRender = strings.NewReader("\033]1337;File=inline=1:")
	iTermCmdStopRender  = strings.NewReader("\a\n") // \n to force terminal to do new line
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "icat",
	Short: "icat enables you to display images on your iTerm command line",
	Long: `icat encodes image to base64 which then can be rendered by iTerm.
You can also output or write the plain base64 the command line / a file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	if os.Getenv(ENV_iTerm_KEY) != ENV_iTerm_VALUE {
		fmt.Printf("Mhm looks like you are trying to run icat on a standard terminal..\nicat only works on iTerm\n")
		os.Exit(1)
	}

	// init deps
	// deps are accessable for all cmds!
	resource = resources.NewResource()
	encoder = b64.NewEncoder()
}
