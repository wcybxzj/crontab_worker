#!/usr/bin/env bash
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"1\",\"exec\": \"cd /tmp && date >> 1.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"2\",\"exec\": \"cd /tmp && date >> 2.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"3\",\"exec\": \"cd /tmp && date >> 3.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"4\",\"exec\": \"cd /tmp && date >> 4.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"5\",\"exec\": \"cd /tmp && date >> 5.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"6\",\"exec\": \"cd /tmp && date >> 6.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
