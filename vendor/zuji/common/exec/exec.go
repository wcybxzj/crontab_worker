package exec

import (
	"io"
	"strings"
	"sync"
	"zuji/common/dlog"
)

// RunCMD Run Shell Command
func RunCMD(inCmd string) (out string, err error) {
	execmd := getCommand(inCmd)
	output, err := execmd.CombinedOutput()

	out = ""
	if err != nil {
		dlog.LogColor(dlog.TextRed, inCmd, err.Error())
		out = string(output)
		return
	}

	out = string(output)

	return
}

// KillCMDSync kill command
func KillCMDSync(cmd string) (out string) {
	out = RunCMDSync(getKillCmd() + " " + cmd + getKillOption())
	return
}

// RunCMDSync Run Shell Command
func RunCMDSync(inCmd string) (out string) {
	out, _ = RunCMD(inCmd)

	out = strings.Trim(out, "\n")
	out = strings.Trim(out, "\r")
	out = strings.TrimSpace(out)

	return
}

// RunCmdASync Run async Shell Command
func RunCmdASync(inCmd string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	waitGroup.Add(1)

	execmd := getCommand(inCmd)

	stdout, err := execmd.StdoutPipe()
	if err != nil {
	}
	stderr, err := execmd.StderrPipe()
	if err != nil {
	}

	execmd.Start()

	defer execmd.Wait()

	go handleOutput(stdout)
	go handleOutput(stderr)
}

// RunCmdSyncWithHandler Run async Shell Command
func RunCmdSyncWithHandler(inCmd string, waitGroup *sync.WaitGroup, fHandler func(reader io.ReadCloser)) {
	defer waitGroup.Done()
	waitGroup.Add(1)

	execmd := getCommand(inCmd)

	stdout, err := execmd.StdoutPipe()
	if err != nil {
		dlog.LogColor(dlog.TextRed, "RunCmdSyncWithHandler stdout:", err)
	}
	stderr, err := execmd.StderrPipe()
	if err != nil {
		dlog.LogColor(dlog.TextRed, "RunCmdSyncWithHandler stderr:", err)
	}

	execmd.Start()

	defer execmd.Wait()

	go fHandler(stdout)
	go fHandler(stderr)
}

func handleOutput(reader io.ReadCloser) {
	for {
		buf := make([]byte, 4096)
		reqLen, err := reader.Read(buf)
		if err == nil {
			dlog.DebugLog("Cmd Output : " + string(buf[:reqLen-1]))
		} else {
			return
		}
	}
}
