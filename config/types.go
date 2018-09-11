package config

//Type:shell, cmd
//static: 每个执行器写配置
//dyanamic: java tls动态传递
type Job struct {
	JobId string `json:"job_id" binding:"required"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Exec  string `json:"exec"`
}

type AddJob struct {
	Name string `json:name`
}

type RunJob struct {
	Name string `json:name`
}

type Data struct {
	AddJob AddJob
	RunJob RunJob
}

type ResponseData struct {
	Data   Data
	Status string `json:status`
}
