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
	"fmt"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"

	"github.com/mdsauce/sauced/logger"
	"github.com/mdsauce/sauced/manager"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var defaultConfigPath = "/sauced/sauced.config"
var configFile string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start all tunnels listed in your config file you reference.",
	Long:  `Start all tunnels in the config file you reference like $ sauced start ~/my-config.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, err := cmd.Flags().GetString("config")
		if err != nil {
			logger.Disklog.Warn("Problem retrieving config file flag", err)
		}

		if !configUsable(configFile) {
			logger.Disklog.Warn("You did not specify a config file! Please pass in a file like 'sauced start --config /path/to/sauced-config.txt'")
			logger.Disklog.Debug("Checking for config file in any XDG_HOME_CONFIG environment variables.")
			configFile = findXdgConfigHome()

			if len(configFile) < 2 {
				logger.Disklog.Error("Problem retrieving config file flag.")
				os.Exit(1)

			}
		}

		manager.PruneState()
		meta := manager.CollectMetadata(configFile)

		var wg sync.WaitGroup
		// read in the sc startup commands
		file, _ := os.Open(configFile)
		fscanner := bufio.NewScanner(file)
		stop := make(chan os.Signal, 1)
		for fscanner.Scan() {
			if fscanner.Text() != "" || len(fscanner.Text()) != 0 {
				wg.Add(1)
				pool := manager.PoolName(fscanner.Text())
				logger.Disklog.Debugf("%s pool is %s.  Metadata is %v", fscanner.Text(), pool, meta[pool])
				go manager.Start(fscanner.Text(), &wg, meta[pool])
			}
		}
		signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-stop
			logger.Disklog.Warn("User pressed CTRL-C (SIGINT). Killing tunnels now.  Active jobs using these tunnels may die.")
		}()
		wg.Wait()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringP("config", "c", "", "config file for tunnels to start.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".sauced" (without extension).
		viper.SetConfigType("toml")
		viper.AddConfigPath(home)
		viper.AddConfigPath("$HOME/.config")
		viper.AddConfigPath("$HOME/.config/sauced")
		viper.AddConfigPath("$HOME/.sauced/config")
		viper.AddConfigPath("$HOME/.sauced")
		//check for XDG_CONFIG_HOME somewhere in here
		viper.SetConfigName("sauced")
		viper.SetConfigName("default")
		// call this AFTER setting paths
		viper.WatchConfig()
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Disklog.Debug("Using config file:", viper.ConfigFileUsed())
	} else {
		logger.Disklog.Error("Problem reading config file: ", err)
	}
}

func configUsable(file string) bool {
	if file == "" || len(file) == 0 {
		logger.Disklog.Warn("No config file passed in with -c or --config", file)
		return false
	}

	fd, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Disklog.Warnf("Cannot find file '%s' it does not exist.", file)
		}
		logger.Disklog.Warnf("Problem getting information about file '%s' it may not exist.", file)
		return false
	}

	if fd.IsDir() {
		logger.Disklog.Warnf("%s is a directory and cannot be read as a config file.", file)
		return false
	}
	return true
}

func findXdgConfigHome() string {
	xdgHome, isXdgSet := os.LookupEnv("XDG_CONFIG_HOME")
	xdgConfigPath := path.Join(xdgHome, defaultConfigPath)
	logger.Disklog.Debugf("Looking for %s on XDG_CONFIG_HOME: %s", defaultConfigPath, xdgHome)

	if !isXdgSet {
		logger.Disklog.Debug("XDG_CONFIG_HOME not set. Exiting.")
		os.Exit(1)
	}

	if !configUsable(xdgConfigPath) {
		logger.Disklog.Debugf("No config file at %s", xdgConfigPath)
		os.Exit(1)
	}
	// If we got here it means that there is a config file located on XDG_CONFIG_HOME
	// assigning configFile to it
	logger.Disklog.Debugf("Found config file %s. Setting to %s", defaultConfigPath, xdgConfigPath)

	//TODO: Add unit/integration test to check that we are actually
	configFile := xdgConfigPath

	return configFile
}
