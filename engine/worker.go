package engine

import (
	"crontab_worker/config"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
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
	fmt.Printf("jobid:%s exec:%s", job.JobId, job.Exec)

	mypath, err := os.Getwd()
	if err != nil {
		return err
	}

	shellCmd := "#!/bin/bash\n"
	shellCmd += job.Exec + "\n"
	filepath := mypath + "/" + uuid.New().String() + ".sh"

	f, err := os.Create(filepath)
	if err != nil {
		dlog.LogColor(dlog.TextRed, "DoWork create shell file", err)
		return err
	}

	f.WriteString(shellCmd)
	f.Chmod(0755)
	f.Close()

	out := exec.RunCMDSync("/bin/bash -c " + filepath)
	dlog.LogColor(dlog.TextGreen, "DoWork", out)

	err = os.Remove(filepath)
	if err != nil {
		dlog.LogColor(dlog.TextRed, "DoWork delete file", err)
	}
	time.Sleep(time.Second * 10)

	FinishJobWorking(job.JobId)

	return nil

}
