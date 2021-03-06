package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"zuji/common/debug"
)

const JOB_NAME = "crontab_worker"

var JobIdsMap sync.Map

//Config config
var Config struct {
	IsDebug       bool
	SaveLog       bool
	Jobs          []Job
	LogURI        string
	LogHost       string
	LogService    string
	MaxGoroutines string
}

func init() {
	Config.IsDebug = true
	Config.SaveLog = true
}

//LoadConfig load config
func LoadConfig() error {
	jsonBytes, err := ioutil.ReadFile("data/jobs.json")
	if err != nil {
		return err
	}

	Config.Jobs = nil

	err = json.Unmarshal(jsonBytes, &Config)
	if err != nil {
		return err
	}

	debug.IsDebug = Config.IsDebug
	debug.SaveLog = Config.SaveLog

	//apibeatlog.Init(Config.LogURI, Config.LogService, Config.LogHost)

	return nil
}
