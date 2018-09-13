package main

import (
	"fmt"
	"net/http"

	"bytes"
	"io/ioutil"
	"strconv"
	"time"
	"zuji/common/post"
)

//curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"12345\",\"exec\": \"cd /tmp && date >> 12345.txt\"}" http://127.0.0.1:8080/ReceiveDiyJob
func httpPostSameJob() {
	http_url := "http://127.0.0.1:8080/ReceiveDiyJob"
	//http_url := "http://127.0.0.1/test.php"

	var jsonStr = []byte(`{"job_id":"12345", "exec":"cd /tmp && date >> 12345.txt && sleep 1"}`)
	req, err := http.NewRequest("POST", http_url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func httpPostDiffJob(id int) {
	http_url := "http://127.0.0.1:8080/ReceiveDiyJob"
	id_str := strconv.Itoa(id)
	json_str := fmt.Sprintf(`{"job_id": "%s", "exec":"cd /tmp && date >> %d.txt && sleep 1"}`, id_str, id)
	jsonStr := []byte(json_str)

	fmt.Println(json_str)

	req, err := http.NewRequest("POST", http_url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//主要是测相同任务,只能接受一个
func SameJobTest() {
	for id := 0; id < 30; id++ {
		httpPostSameJob()
	}
	time.Sleep(time.Second * 3)
}

//主要是发不同的任务,数量大于server的协程数,希望能在任务队列看能任务的挤压
//可以看到任务的加压,然后被处理
//服务需要开启
//"IsDebugGoroutineNum":true,
//"IsDebugQueueStatus":true,
func DiffJobTest() {
	for id := 0; id < 100; id++ {
		httpPostDiffJob(id)
	}
	time.Sleep(time.Second * 3)
}

func checkTest() {
	out, err := post.Get("http://127.0.0.1:8080/Check")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func ReloadTest() {
	out, err := post.Get("http://127.0.0.1:8080/Reload")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func main() {
	//ReloadTest()
	//SameJobTest()
	DiffJobTest()
	//checkTest()
}
