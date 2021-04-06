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
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/GDGVIT/katamari/internal/styles"
	"github.com/kardianos/osext"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// Warn prints a formatted and colored WARN message
func Warn(t string, b string) {
	fmt.Println(chalk.White.Color("katamari"), styles.WarnStyle.Style("WARN"), chalk.Magenta.Color(t), chalk.Yellow.Color(b))
}

// Info prints a formatted and colored INFO message
func Info(t string, b string) {
	fmt.Println(chalk.White.Color("katamari"), styles.InfoStyle.Style("INFO"), chalk.Magenta.Color(t), chalk.White.Color(b))
}

// Err prints a formatted and colored ERROR message
func Err(t string, b string) {
	fmt.Println(chalk.White.Color("katamari"), styles.ErrorStyle.Style("ERROR"), chalk.Magenta.Color(t), chalk.White.Color(b))
}

//ExecLoc to get the install loc of binary
func ExecLoc() string {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	return folderPath
}

//CheckConfig to check if config file exists
func CheckConfig() {
	folderPath := ExecLoc()
	_, err := ioutil.ReadFile(folderPath + "/.katamari/config.json")
	if err != nil {
		generateConfig()
	}
}

//generateConfig to generate an empty config file
func generateConfig() {
	credString := `
	{
		"GITHUB_ACCESS_TOKEN":""
	}
	`
	mode := int(0666)
	folderPath := ExecLoc()
	err := os.Mkdir(folderPath+"/.katamari", os.FileMode(0777))
	if err != nil {
		Err("fatal", "unable to create folder: "+err.Error())
	}
	ioutil.WriteFile(folderPath+"/.katamari/config.json", []byte(credString), os.FileMode(mode))
}

//UnsetAndSaveConfig remove key from viper and save as filename
func UnsetAndSaveConfig(key string, fileName string) error {
	configMap := viper.AllSettings()
	delete(configMap, key)
	encodedConfig, _ := json.MarshalIndent(configMap, "", " ")
	err := viper.ReadConfig(bytes.NewReader(encodedConfig))
	if err != nil {
		return err
	}
	return viper.SafeWriteConfigAs(".katamari.toml")
}
