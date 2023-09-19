#!/usr/bin/bash
# 下面这两行的目的，是更改rabbitmq写入的路径为普通用户可以写入的路径。
# 它默认的路径，需要root权限
export RABBITMQ_LOG_BASE=$PWD/srv/rabbitmq-base/log
export RABBITMQ_MNESIA_BASE=$PWD/srv/rabbitmq-base/mnesia

# 这是1024code独有的内容，如果你只需要跑在本地，可以修改他们
# 1024code上，这些变量会变化，但是总归是环境变量
echo "db_username: ${MYSQL_USER}
db_password: ${MYSQL_PASSWORD}
db_name: douyin
db_addr: ${MYSQL_HOST}:${MYSQL_PORT}
redis_addr: ${REDIS_HOST}:${REDIS_PORT}
redis_password: ${REDISCLI_AUTH}
redis_db_no: 0
# 下面这个链接，必须以/结尾，
# 本来传递它很困难，但是好在1024code提供了该环境变量
video_srv_prefix: https://${paas_url}/
# 为什么不配置consul和rabbitmq？原因是，它们目前跑在本地。

" > cmd/config.yaml

# 避免这些文件夹不存在
mkdir $RABBITMQ_LOG_BASE $RABBITMQ_MNESIA_BASE -p
mkdir cmd/storage/static -p

consul agent -dev -client=0.0.0.0 -log-level error &
rabbitmq-server &
echo 等待服务启动
# 因为1024上rabbitmq启动较慢，一般要15秒，所以这里写20
sleep 20

for p in $@; do
	$(cd $p; make)&
done
wait
