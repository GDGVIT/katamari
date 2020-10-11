package styles

import "github.com/ttacon/chalk"

// WarnStyle style for decorating WARN text
var WarnStyle = chalk.Bold.NewStyle().WithBackground(chalk.Yellow).WithForeground(chalk.Black)

// InfoStyle style for decorating INFO text
var InfoStyle = chalk.Bold.NewStyle().WithBackground(chalk.ResetColor).WithForeground(chalk.Green)

// ErrorStyle style for decorating ERROR text
var ErrorStyle = chalk.Bold.NewStyle().WithBackground(chalk.ResetColor).WithForeground(chalk.Red)
