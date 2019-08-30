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
	"github.com/mdsauce/sauced/logger"
	"github.com/mdsauce/sauced/manager"
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop all running tunnels and close this program.",
	Long:  `Use the last known tunnel state to stop all tunnels.  This process will close after SIGINT or kill signal has been deliverd to all tunnels.`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: This should be it's function since it's called on all files.
		logfile, err := cmd.Flags().GetString("logfile")
		if err != nil {
			logger.Disklog.Warn("Problem retrieving logfile flag", err)
		}
		logger.SetupLogfile(logfile)

		pool, _ := cmd.Flags().GetString("pool")
		id, _ := cmd.Flags().GetString("id")
		all, _ := cmd.Flags().GetBool("all")

		logger.Disklog.Debug("All flag: ", all)
		logger.Disklog.Debug("Pool name searched: ", pool)
		logger.Disklog.Debug("ID searched: ", id)

		if all { // stop all tunnel regardless of other flags
			logger.Disklog.Info("Stop command sent.  Stopping all tunnels.")
			manager.StopAll()
			logger.Disklog.Info("All tunnels sent the Kill, Interrupt, or SIGINT signal.  Sauced closing.")
		} else if pool == "" && id != "" { // id is the only one set
			// find tunnel by ID
			// get pid
			// stop tunnel
			// remove tunnel from state
		} else if pool != "" { // delete pool. doesn't matter if id is set too
			// find tunnel(s) by pool
			// get pid(s)
			// stop tunnel(s)
			// remove tunnel(s) from state
		} else {
			logger.Disklog.Error("No flag was added. Please use -h for options")
		}

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	stopCmd.Flags().String("pool", "", "Pool name of tunnels. May return one or more results. Takes precedence over --id")
	stopCmd.Flags().String("id", "", "Assigned ID for a given tunnel.")
	stopCmd.Flags().Bool("all", false, "Allows stopping all active tunnels. Takes precedence over all other flags.")
}
