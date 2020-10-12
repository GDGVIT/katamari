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
	"fmt"
	"github.com/GDGVIT/katamari/internal/styles"

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