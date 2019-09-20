// Copyright © 2019
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/mdsauce/sauced/logger"
	"github.com/spf13/cobra"
)

// CurVersion is a global reference to the version number set at build time
var CurVersion = "DEV"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sauced",
	Short: "Read from a config file at and start some tunnels.",
	Long: `Will read from a user specified config file and start tunnels as they appear.
Each tunnel should be on one line and separated by a newline character.  Use the start cmd
to start all the specified tunnels.  Stop cmd will stop all tunnels that were started by the
program.`,
	Version: fmt.Sprintf("%s", CurVersion),
	// Persistent*Run hooks are inherited by all subcommands if they do not have hooks themselves
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logfile, err := cmd.Flags().GetString("logfile")
		if err != nil {
			logger.Disklog.Warn("Problem retrieving logfile flag: ", err)
		}
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			logger.Disklog.Warn("Problem retrieving verbosity flag: ", err)
		}
		logger.SetupLogfile(logfile, verbose)
	},
	// Run: func(cmd *cobra.Command, args []string) {
	// },
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
	rootCmd.Flags().Bool("version", false, "Print the version of sauced")
	rootCmd.PersistentFlags().Bool("verbose", false, "Print out all sauced logging information")
	rootCmd.PersistentFlags().StringP("logfile", "l", "/tmp/sauced.log", "logfile for meta-status output")
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/sauced.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
