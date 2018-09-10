package dlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"zuji/common/debug"
)

const timeFormat = "01-02 15:04:05"

func getVal(a ...interface{}) []interface{} {
	ss := make([]interface{}, 0, len(a)+1)
	ss = append(ss, time.Now().Format(timeFormat))
	ss = append(ss, a...)

	return ss
}

var logFile *os.File
var debugLog *log.Logger

//OpenLogfile open log file
func OpenLogfile(path string) {
	if debug.SaveLog {
		var err error
		logFile, err = os.OpenFile(path, os.O_WRONLY, 0666)
		if err != nil {
			LogColor(TextRed, "OpenLogfile OpenFile", err)
			logFile, err = os.Create(path)
			if err != nil {
				LogColor(TextRed, "OpenLogfile Create", err)
				return
			}
		}

		logFile.Seek(0, io.SeekEnd)

		// 创建一个日志对象
		debugLog = log.New(logFile, "[Debug]", log.LstdFlags)
	}
}

func writeLog(a ...interface{}) {
	if debug.SaveLog {
		if debugLog != nil {
			debugLog.Println(a...)
		}
	}
}

//CloseLogfile close log file
func CloseLogfile() {
	if logFile != nil {
		logFile.Close()
	}
}

//PrintLine print line
func PrintLine(a ...interface{}) {
	if debug.IsDebug {
		fmt.Println(a...)
	}
}

//Printf Printf log
func Printf(format string, a ...interface{}) {
	if debug.IsDebug {
		fmt.Printf(format, getVal(a)...)
	}

	writeLog(fmt.Sprintf(format, a...))
}

//DebugLog debug log
func DebugLog(a ...interface{}) {
	PrintLine(getVal(a)...)
	writeLog(a...)
}

//DebugLogColor debug log
func DebugLogColor(str string, color int) {
	printColorText(color, str)
	writeLog(str)
}

//LogColor debug log
func LogColor(color int, val ...interface{}) {
	printColorText(color, val...)
	writeLog(val...)
}

//LogError debug log
func LogError(val ...interface{}) {
	printColorText(TextRed, val...)
	writeLog(val...)
}
