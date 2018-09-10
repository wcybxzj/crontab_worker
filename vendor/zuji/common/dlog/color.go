package dlog

import (
	"fmt"
	"runtime"
	"time"
)

const TextNormal = iota

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func printColorText(color int, val ...interface{}) {
	if color <= TextNormal || isWindows() {
		PrintLine(getVal(val)...)
	} else {
		PrintLine(fmt.Sprintf("\x1b[%dm%s %+v\x1b[0m", color, time.Now().Format(timeFormat), val))
	}
}

func isWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	}

	return false
}
