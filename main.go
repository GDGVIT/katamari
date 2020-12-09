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
package main

import (
	"log"
	"os"

	"github.com/GDGVIT/katamari/cmd"
	"github.com/GDGVIT/katamari/internal/utils"
	"github.com/spf13/viper"
)

func main() {
	utils.CheckConfig()
	viper.SetConfigFile(utils.ExecLoc() + "/.katamari/config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}
	os.Setenv("GITHUB_ACCESS_TOKEN", viper.GetString("GITHUB_ACCESS_TOKEN"))

	cmd.Execute()
}
