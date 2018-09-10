// +build !windows

package exec

import (
	"os/exec"
	"strings"
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
	return ""
}

func getKillCmd() (cmd string) {
	cmd = "killall"
	return
}

//SetEnv set env
func SetEnv(mypath string) {

}
