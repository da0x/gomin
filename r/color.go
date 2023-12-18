package r

import (
	"github.com/fatih/color"
)

var Red = color.New(color.FgRed).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
var Blue = color.New(color.FgBlue).SprintFunc()
var Purple = color.New(color.FgMagenta).SprintFunc()
var Cyan = color.New(color.FgCyan).SprintFunc()
var Gray = color.New(color.FgWhite).SprintFunc() // FgWhite is often used as gray in terminal color schemes
var White = color.New(color.FgHiWhite).SprintFunc()

var BrightRed = color.New(color.FgHiRed).SprintFunc()
var BrightGreen = color.New(color.FgHiGreen).SprintFunc()
var BrightYellow = color.New(color.FgHiYellow).SprintFunc()
var BrightBlue = color.New(color.FgHiBlue).SprintFunc()
var BrightPurple = color.New(color.FgHiMagenta).SprintFunc()
var BrightCyan = color.New(color.FgHiCyan).SprintFunc()
var BrightGray = color.New(color.FgHiBlack).SprintFunc() // FgHiBlack is often used as bright gray in terminal color schemes
