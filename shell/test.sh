#执行自定义任务
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"1\",\"exec\": \"cd /tmp && curl www.baidu.com >> 1.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"2\",\"exec\": \"cd /tmp && curl www.163.com >> 2.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"3\",\"exec\": \"cd /tmp && curl www.qq.com >> 3.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"4\",\"exec\": \"cd /tmp && date >> 4.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"5\",\"exec\": \"cd /tmp && date >> 5.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"6\",\"exec\": \"cd /tmp && date >> 6.txt && sleep 15 \"}" http://127.0.0.1:8080/ReceiveDiyJob

#执行在执行器配置好的任务
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"100\"}" http://127.0.0.1:8080/ReceiveConfigedJob
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"200\"}" http://127.0.0.1:8080/ReceiveConfigedJob

#开发机:宿主sh->screen->docker->php->screen->php子任务
curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"7\",\"exec\":  \" sh /tmp/1.sh  \"}" http://127.0.0.1:8080/ReceiveDiyJob

#测试开发环境
#lts java后台配置
{
    "url":"http://172.18.124.159:8080/ReceiveDiyJob",
    "data":"{\"job_id\": \"1\",\"exec\": \"cd \/tmp && date >> 1.txt && sleep 15 \"}"
}


curl -X POST -H "Content-Type:application/json" -d "{\"job_id\": \"1\",\"exec\": \"cd /tmp && curl www.baidu.com >> 1.txt && sleep 15 \"}" http://172.18.124.159:8080/ReceiveDiyJob

