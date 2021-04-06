/*
Copyright © 2020 DSCVIT

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/GDGVIT/katamari/internal/utils"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "katamari <command> <subcommand> [flags]",
	Short:         "katamari CLI",
	Version:       "1.0.4",
	Long:          `Use katamari to aggregate and build static pages for all the READMEs in your organization`,
	SilenceErrors: true,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/.katamari.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println(cfgFile)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		cwd, err := os.Getwd()
		if err != nil {
			utils.Err("enoent", err.Error())
			os.Exit(1)
		}

		// Search config in cwd directory with name ".katamari" (without extension).
		viper.AddConfigPath(cwd)
		viper.SetConfigName(".katamari")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		utils.Info("sill", fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
	}
}
