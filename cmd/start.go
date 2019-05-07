// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
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
	"os"
	"sync"

	"github.com/maxdobeck/sauced/logger"
	"github.com/maxdobeck/sauced/manager"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start all tunnels listed in your config file you reference.",
	Long:  `Start all tunnels in the config file you reference like $ sauced start ~/my-config.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logger.Disklog.Warn("You did not specify a config file!  Please write a .txt file and pass it in like $ sauced start /path/to/config.txt")
			os.Exit(1)
		}
		configFile := args[0]
		// Get logfile
		logfile, err := cmd.Flags().GetString("logfile")
		if err != nil {
			logger.Disklog.Warn("Problem retrieving logfile flag", err)
		}
		logger.SetupLogfile(logfile)

		manager.PruneState()
		meta := manager.CollectMetadata(configFile)
		manager.UpdateState(meta)

		var wg sync.WaitGroup
		// read in the sc startup commands
		file, _ := os.Open(configFile)
		fscanner := bufio.NewScanner(file)
		for fscanner.Scan() {
			if fscanner.Text() != "" || len(fscanner.Text()) != 0 {
				wg.Add(1)
				c := make(chan string)
				go manager.PoolName(fscanner.Text(), c)
				pool := <-c
				logger.Disklog.Debugf("%s pool is %s.  Metadata is %v", fscanner.Text(), pool, meta[pool])
				go manager.Start(fscanner.Text(), &wg, meta[pool])
			}
		}
		wg.Wait()
		logger.Disklog.Info("All tunnels must be closed.  Goodbye :)")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
