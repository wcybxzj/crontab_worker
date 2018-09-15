#/bin/bash

# */5  * * * * cd /data/go_www/src/crontab_worker/ && ./run.sh
mpath=$(dirname $0)
crontab_worker=${mpath}/crontab_worker

RUN=`screen -list | grep crontab_worker -c`
if [ $RUN -lt 1 ]; then
  echo "run crontab_worker"
  screen -dmS crontab_worker ${crontab_worker}
fi