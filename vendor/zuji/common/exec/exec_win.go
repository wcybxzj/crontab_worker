// +build windows

package exec

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"zuji/common/dlog"
	"zuji/common/str"
)

func getCommand(inCmd string) (execmd *exec.Cmd) {

	inCmd = strings.Trim(inCmd, " ")

	parts := strings.Fields(inCmd)

	if len(parts) > 1 {
		head := parts[0]
		parts = parts[1:len(parts)]

		execmd = exec.Command(head, parts...)
	} else {
		execmd = exec.Command(inCmd)
	}

	return
}

func getKillOption() string {
	return ".exe /F"
}

func getKillCmd() (cmd string) {
	cmd = "taskkill /IM"
	return
}

//SetEnv set env
func SetEnv(mypath string) {

	dir, err := filepath.Abs(filepath.Dir(mypath))
	if err != nil {
		dlog.LogColor(dlog.TextRed, err)
		return
	}

	path := os.Getenv("PATH")

	if !str.HasSuffix(path, ";") {
		path += ";"
	}

	path += dir + "\\tools;"

	dlog.DebugLog(path)

	os.Setenv("PATH", path)

}
