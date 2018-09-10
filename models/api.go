package models

const (
	OK = iota + 1
	ERROR_REQ_FORMAT
	ERROR_JOB_WORLING
	ERROR_UNKNOWN_JOBID
)

type ResponseData struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}
