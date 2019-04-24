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
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/maxdobeck/sauced/logger"
	"github.com/maxdobeck/sauced/manager"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sauced",
	Short: "Read from a YAML file at $HOME/.config/sauced.yaml and list the changes",
	Long:  `First test to read and watch a YAML config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get logfile
		logfile, err := cmd.Flags().GetString("logfile")
		if err != nil {
			logger.Disklog.Warn("Problem retrieving logfile flag", err)
		}
		logger.SetupLogfile(logfile)

		manager.PruneState()

		var wg sync.WaitGroup
		// read in the sc startup commands
		file, _ := os.Open(args[0])
		fscanner := bufio.NewScanner(file)
		for fscanner.Scan() {
			if fscanner.Text() != "" {
				wg.Add(1)
				go manager.Start(fscanner.Text(), &wg)
			}
		}
		wg.Wait()
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/sauced.yaml)")
	rootCmd.PersistentFlags().StringP("logfile", "l", "/tmp/sauced.log", "logfile for meta-status output (default is /tmp/sauced.log)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
