/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"path"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const pointsKey = "points"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jd",
	Short: "jd (jump-directory): CD with kicks.",
	Long:  ``,

	// make root cmd act as a catch all
	Args: cobra.ArbitraryArgs,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("invalid command")
		}
		points := viper.GetStringMapString(pointsKey)
		dir, ok := points[args[0]]
		if !ok {
			return fmt.Errorf("jump point \"%s\" not found", args[0])
		}
		inject(fmt.Sprintf("cd \"%s\"\r", dir))
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	configPath := path.Join(home, ".jd.yaml")
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		if _, err := os.Create(configPath); err != nil { // perm 0666
			cobra.CheckErr(err)
		}
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(fmt.Errorf("failed to read config file"))
	}
}
