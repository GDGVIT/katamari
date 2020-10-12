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
	"github.com/GDGVIT/katamari/internal/utils"
	"github.com/spf13/viper"
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
			msg := fmt.Sprintf("Ignoring extra arguments after %s", chalk.Green.Color("'"+args[0]+"'"))
			utils.Warn("optional", msg)
		} else if len(args) == 0 {
			_ = cmd.Help()
			os.Exit(0)
		}

		utils.Info("sill", "Initializing a new katamari project...")
		viper.Set("site", args[0])
		viper.Set("theme", "ananke")

		hugoPath, err := exec.LookPath("hugo")
		if err != nil {
			utils.Err("notsup", "could not find hugo! make sure you have hugo installed.")
			os.Exit(1)
		}

		hugo := exec.Cmd{
			Path:   hugoPath,
			Stdout: nil,
			Stderr: os.Stderr,
			Args:   []string{"", "new", "site", args[0]},
		}

		utils.Info("run", hugo.String())

		err = hugo.Run()
		if err != nil {
			utils.Err("fatal", err.Error())
			os.Exit(1)
		}

		err = os.Chdir(fmt.Sprintf("./%s", args[0]))
		if err != nil {
			utils.Err("fatal", err.Error())
			os.Exit(1)
		}

		err = viper.SafeWriteConfigAs(".katamari.toml")
		if err != nil {
			utils.Err("fatal", err.Error())
			os.Exit(1)
		}

		utils.Info("config", "Generated katamari config file in project directory")
		utils.Info("sill", fmt.Sprintf("Created katamari project %s", chalk.Green.Color(args[0])))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
