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
	"github.com/mdsauce/sauced/output"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "List all last known tunnels.",
	Long:  `The tunnel state list will be pruned and then all active tunnels will be shown.`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO: This should be it's function since it's called on all files.
		logfile, err := cmd.Flags().GetString("logfile")
		if err != nil {
			logger.Disklog.Warn("Problem retrieving logfile flag", err)
		}
		logger.SetupLogfile(logfile)

		pretty, err := cmd.Flags().GetBool("pretty")
		if err != nil {
			logger.Disklog.Warn("Problem retrieving pretty flag", err)
		}

		logger.Disklog.Debug("show called by the user. Pruning then listing all tunnels.")
		manager.PruneState()

		pool, _ := cmd.Flags().GetString("pool")
		id, _ := cmd.Flags().GetString("id")

		logger.Disklog.Debug("Pool name searched: ", pool)
		logger.Disklog.Debug("ID searched: ", id)

		switch pretty {
		case true:
			output.PrettyPrint(id, pool)
		case false:
			if pool == "" && id != "" {
				output.ShowTunnelJSON(id)
			} else if pool == "" && id == "" {
				output.ShowStateJSON()
			} else {
				output.ShowPool(pool)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.PersistentFlags().Bool("pretty", false, "pretty print the state")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	showCmd.Flags().String("pool", "", "Pool name of tunnels. May return one or more results. Takes precedence over --id")
	showCmd.Flags().String("id", "", "Assigned ID for a given tunnel.")

}
