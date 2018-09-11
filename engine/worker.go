package engine

import (
	"crontab_worker/config"
	"fmt"
	"github.com/google/uuid"
	"os"
	"zuji/common/dlog"
	"zuji/common/exec"
)

func IsJobWorking(JobId string) bool {
	_, ok := config.JobIdsMap.LoadOrStore(JobId, 1)
	return ok
}

func FinishJobWorking(JobId string) {
	config.JobIdsMap.Delete(JobId)
}

func DoWork(job config.Job) error {
	mypath := "/tmp"
	shellCmd := "#!/bin/bash\n"
	shellCmd += job.Exec + "\n"
	filepath := mypath + "/" + uuid.New().String() + ".sh"

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	fmt.Println(job.Exec)

	f.WriteString(shellCmd)
	f.Chmod(0755)
	f.Close()

	out := exec.RunCMDSync("/bin/bash -c " + filepath)
	dlog.LogColor(dlog.TextGreen, "DoWork  RunCMDSync JobId"+job.JobId+" out:"+out)

	err = os.Remove(filepath)
	if err != nil {
		return err
	}

	FinishJobWorking(job.JobId)
	return nil
}
