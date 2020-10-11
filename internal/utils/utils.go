package utils

import (
	"fmt"
	"katamari/internal/styles"

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