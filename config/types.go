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
