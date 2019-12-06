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
	"strings"
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
		launch(configFile)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file for tunnels to start.")
}

func initConfig() {
	if configFile != "" {
		// Use config file from the --config flag.
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logger.Disklog.Warn(err)
			os.Exit(1)
		}

		viper.SetConfigName("sauced")
		xdgHome, _ := os.LookupEnv("XDG_CONFIG_HOME")
		viper.SetConfigType("toml")
		viper.AddConfigPath(home)
		viper.AddConfigPath(xdgHome)
		viper.AddConfigPath("$HOME/.config")
		viper.AddConfigPath("$HOME/.config/sauced")
		viper.AddConfigPath("$HOME/.sauced/config")
		viper.AddConfigPath("$HOME/.sauced")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Disklog.Info("No 'default' or 'sauced' config file found. ", err)
		}
	} else if err == nil {
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

	//TODO: Add unit/integration test
	configFile := xdgConfigPath

	return configFile
}

func launch(config string) {
	if !configUsable(viper.ConfigFileUsed()) {
		os.Exit(1)
	}
	if strings.Contains(config, ".yaml") {
		logger.Disklog.Error("YAML is not supported at this time.  Cannot use config file: ", configFile)
	} else if strings.Contains(viper.ConfigFileUsed(), ".toml") || strings.Contains(config, ".toml") {
		logger.Disklog.Info("Found TOML config file: ", viper.ConfigFileUsed())
		configMap := make(map[string]manager.TunnelArgs)
		if err := viper.ReadInConfig(); err != nil {
			logger.Disklog.Fatalf("Error reading config file, %s", err)
		}
		err := viper.Unmarshal(&configMap)
		if err != nil {
			logger.Disklog.Fatalf("Unable to unmarshal config file %s: %v", viper.ConfigFileUsed(), err)
		}
		configs := mapToSlice(configMap)
		fmt.Println(configs)
	} else if viper.ConfigFileUsed() == "" && configFile == "" {
		unstructuredConfig(findXdgConfigHome())
	} else {
		unstructuredConfig(configFile)
	}
}

func mapToSlice(config map[string]manager.TunnelArgs) []manager.TunnelArgs {
	c := []manager.TunnelArgs{}
	i := 0
	for k, v := range config {
		c = append(c, v)
		c[i].TunnelIdentifier = k
		i++
	}
	logger.Disklog.Debugf("Converted map %v to slice %v", c, config)
	logger.Disklog.Debug("Pool Size is: ", c[0].PoolSize)
	return c
}

func unstructuredConfig(configFile string) {
	logger.Disklog.Debug("Using config: ", configFile)
	if !configUsable(configFile) {
		logger.Disklog.Warn("You did not specify a config file! Pass in a file like 'sauced start --config /path/to/sauced-config.txt'")
		logger.Disklog.Debug("Checking for config file in any XDG_HOME_CONFIG environment variables.")
		configFile = findXdgConfigHome()

		if len(configFile) < 2 {
			logger.Disklog.Error("Problem retrieving config file flag.")
			os.Exit(1)

		}
	}
	if !strings.Contains(configFile, ".config") {
		logger.Disklog.Error("Config file is not a .config: ", configFile)
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
}
