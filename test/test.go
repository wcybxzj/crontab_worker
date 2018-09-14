package main

import (
	"net/http"
	"net/url"
	"fmt"
	"time"
)

//curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"12345\",\"exec\": \"cd /tmp && date >> 12345.txt\"}" http://127.0.0.1/1.php
func postTest()  {
	resp, _ := http.PostForm("http://127.0.0.1/1.php",
		url.Values{"key": {"Value"}, "id": {"123"}})
	fmt.Println(resp)
	}

func sleepTest(num int)  {
	time.Sleep(time.Second*time.Duration(num))
}



func main() {
	str :="abc"
	str+="def"

	fmt.Println(str)
}
