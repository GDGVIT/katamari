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
	"os/exec"

	"github.com/GDGVIT/katamari/internal/utils"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [organization's/user's github username]",
	Short: "Create a new katamari project",
	Long: `Use this command to create a new project using katamari.
	
	Examples:
	katamari create GDGVIT
	katamari create -u BRO3886
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		utils.Info("sill", "Initializing a new katamari project...")
		viper.Set("site", args[0])
		viper.Set("theme", "smol")

		if isUser {
			utils.Info("sill", fmt.Sprintf("creating katamari project for user %s", chalk.Green.Color(args[0])))
			viper.Set("isUser", true)
		}

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

		gitPath, err := exec.LookPath("git")
		if err != nil {
			utils.Err("notsup", "could not find git! make sure you have git installed, to install a theme.")
			os.Exit(1)
		}

		init := exec.Cmd{
			Path:   gitPath,
			Stdout: nil,
			Stderr: os.Stderr,
			Args:   []string{"", "init"},
		}

		err = init.Run()
		if err != nil {
			utils.Err("fatal", err.Error())
			os.Exit(1)
		}

		addTheme := exec.Cmd{
			Path:   gitPath,
			Stdout: nil,
			Stderr: os.Stderr,
			Args:   []string{"", "submodule", "add", "https://github.com/colorchestra/smol", "themes/smol"},
		}

		err = addTheme.Run()
		if err != nil {
			utils.Err("fatal", err.Error())
			os.Exit(1)
		}

		f, err := os.OpenFile("config.toml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			utils.Err("fatal", "unable to open config.toml")
			os.Exit(1)
		}

		defer f.Close()

		if _, err := f.WriteString(`theme="smol"`); err != nil {
			utils.Err("fatal", "unable to write to config.toml")
			os.Exit(1)
		}

		err = viper.SafeWriteConfigAs(".katamari.toml")
		if err != nil {
			utils.Err("fatal", err.Error())
			os.Exit(1)
		}

		utils.Info("config", "Generated katamari config file in project directory")
		utils.Info("sill", fmt.Sprintf("Created katamari project %s", chalk.Green.Color(args[0])))
		utils.Info("config", fmt.Sprintf("Configure hugo by editing %s", chalk.Green.Color(args[0]+"/config.toml")))
	},
}

var isUser bool

func init() {
	createCmd.Flags().BoolVarP(&isUser, "user", "u", false, "use this flag to aggregate READMEs for a User")
	rootCmd.AddCommand(createCmd)
}
