#!/bin/bash
dir=`pwd`
exec_name=$dir/bin/route7
log_dir=$dir/log
# cfg for kill
cfg=R7_BIZ_20140708

export REALTIME_SERVER_URL="http://localhost:8082"
export CACHE_SERVER_PORT=":8182"

case "$1" in
  start)
    echo "Starting $exec_name......"
    nohup $exec_name -log_dir=$log_dir $cfg 2>&1 >> $log_dir/stdout.log &
    echo "Done"
    ;;
  stop)
    echo "Stop $exec_name......"
    pid=`ps aux | grep -v grep | grep "$exec_name" | grep "${cfg}" | awk '{print $2}'`
    kill -9 $pid
    echo "Done"
    ;;
  stat)
    ps aux | grep -v grep | grep "$exec_name" | grep "${cfg}"
    ;;
  restart)
    echo "Restart $exec_name......"
    pid=`ps aux | grep -v grep | grep "$exec_name" | grep "${cfg}" | awk '{print $2}'`
    kill -9 $pid
    sleep 1
    echo "Old server stoped ...... "
    ps aux | grep -v grep | grep "$exec_name" | grep "${cfg}"
    echo "Starting $exec_name ......"
    nohup $exec_name -log_dir=$log_dir $cfg 2>&1 >> $log_dir/stdout.log &
    echo "Done"
    ;;
  *)
    echo "Usage: ./server_ctl.sh {start|stop|restart|stat}" >&2
    ;;
esac




