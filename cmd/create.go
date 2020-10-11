/*
Copyright © 2020 DSCVIT dscvitvellore@gmail.com

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
	"katamari/internal/utils"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new katamari project",
	Long: `Use this command to create a new project using katamari.
	
	Example: katamari create gdgvit`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			msg := fmt.Sprintf("Ignoring extra arguments after %s", chalk.Green.Color("'" + args[0] + "'"))
			utils.Warn("optional", msg)
		} else if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		utils.Info("sill", "Initializing a new katamari project...")

		hugoPath, err := exec.LookPath("hugo")
		if err != nil {
			utils.Err("notsup", "could not find hugo! make sure you have hugo installed.")
			os.Exit(1)
		}
		
		hugo := exec.Cmd {
			Path: hugoPath,
			Stdout: nil,
			Stderr: os.Stderr,
			Args: []string{"", "new", "site", args[0]},
		}

		utils.Info("run", hugo.String())

		err = hugo.Run()
		if err != nil {
			utils.Err("fatal", err.Error())
		}

		utils.Info("sill", fmt.Sprintf("Created katamari project %s", chalk.Green.Color(args[0])))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
