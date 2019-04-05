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
	"os/exec"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sauced",
	Short: "Read from a YAML file at $HOME/.config/sauced.yaml and list the changes",
	Long:  `First test to read and watch a YAML config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		scbinary := viper.GetString("sc-path")
		fmt.Println("Launching Sauce Connect Proxy binary at", scbinary)

		scCmd := exec.Command(scbinary)
		stdout, _ := scCmd.StdoutPipe()
		err := scCmd.Start()
		if err != nil {
			fmt.Println("Something went wrong with the sc binary! ", err)
		}

		fmt.Printf("Sauce Connect started as process %d.\n", scCmd.Process.Pid)
		scanner := bufio.NewScanner(stdout)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			m := scanner.Text()
			fmt.Println(m)
			if strings.Contains(m, "Sauce Connect is up") {
				fmt.Println("Sauce Connect started!  Killing it for you now so you don't forget!")
				// can't send interrupts on Windows!! Beware, must use scCmd.Process.Kill
				scCmd.Process.Signal(os.Interrupt)
				break
			}
		}

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
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/sauced.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".sauced" (without extension).
		viper.SetConfigType("yaml")
		viper.AddConfigPath(home)
		viper.AddConfigPath("$HOME/.config")
		viper.AddConfigPath(home + "/.config")
		viper.SetConfigName("sauced")
		viper.WatchConfig() // call this AFTER setting paths
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println("Problem reading config file:", err)
	}

	tunnels := viper.GetStringMap("tunnels")
	for key, tunnel := range tunnels {
		fmt.Println()
		fmt.Println("Tunnel: ", key)
		fmt.Println(tunnel)
	}
	fmt.Println()
}
