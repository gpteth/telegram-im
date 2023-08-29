#!/usr/bin/env bash

echo "run databus ..."
LOG_DIR=../logs/databus nohup ./databus -conf=../configs/databus/databus.toml > ../logs/databus/all.log 2>&1 &
sleep 1

echo "run media ..."
LOG_DIR=../logs/media MINIO_ENDPOINT=127.0.0.1:9000 ETCD_ENDPOINTS=127.0.0.1:2379 nohup ./media -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/media > ../logs/media/all.log 2>&1 &
sleep 1

echo "run auth_session ..."
LOG_DIR=../logs/auth_session ETCD_ENDPOINTS=127.0.0.1:2379 nohup ./auth_session -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/auth_session > ../logs/auth_session/all.log 2>&1 &
sleep 1

echo "run relay ..."
LOG_DIR=../logs/relay ETCD_ENDPOINTS=127.0.0.1:2379 nohup ./relay -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/relay > ../logs/relay/all.log 2>&1 &
sleep 1

echo "run inbox ..."
LOG_DIR=../logs/inbox ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40302" nohup ./inbox -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/inbox -predefined=false > ../logs/inbox/all.log 2>&1 &
sleep 1

echo "run msg ..."
LOG_DIR=../logs/msg ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40402" nohup ./msg -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/msg -predefined=false > ../logs/msg/all.log 2>&1 &
sleep 1

echo "run push ..."
LOG_DIR=../logs/push ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40603" nohup ./push -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/push -predefined=false > ../logs/push/all.log 2>&1 &
sleep 1

echo "run botfather ..."
LOG_DIR=../logs/botfather ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40032" nohup ./botfather -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/botfather -predefined=false > ../logs/botfather/all.log 2>&1 &
sleep 1

echo "run gif ..."
LOG_DIR=../logs/gif ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40903" nohup ./gif -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/gif -predefined=false > ../logs/gif/all.log 2>&1 &
sleep 1

echo "run scheduled ..."
LOG_DIR=../logs/scheduled ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40702" nohup ./scheduled -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/scheduled -predefined=false > ../logs/scheduled/all.log 2>&1 &
sleep 1

echo "run admin_log ..."
LOG_DIR=../logs/admin_log ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40802" nohup ./admin_log -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/admin_log -predefined=false > ../logs/admin_log/all.log 2>&1 &
sleep 1

echo "run webpage ..."
LOG_DIR=../logs/webpage ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:41022" nohup ./webpage -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/webpage -predefined=false > ../logs/webpage/all.log 2>&1 &
sleep 1

echo "run biz_server ..."
LOG_DIR=../logs/biz_server MINIO_ENDPOINT=127.0.0.1:9000 ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40102" nohup ./biz_server -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/biz_server -predefined=false > ../logs/biz_server/all.log 2>&1 &
sleep 5

echo "run gateway ..."
LOG_DIR=../logs/gateway ETCD_ENDPOINTS=127.0.0.1:2379 nohup ./gateway -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/gateway -predefined=false > ../logs/gateway/all.log 2>&1 &
sleep 1

echo "run botway ..."
HOSTNAME=botway001 LOG_DIR=../logs/botway MINIO_ENDPOINT=127.0.0.1:9000 ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40082" nohup ./botway -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/botway -predefined=false > ../logs/botway/all.log 2>&1 &
sleep 1

echo "run api_server ..."
LOG_DIR=../logs/api_server MINIO_ENDPOINT=127.0.0.1:9000 ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:8883" nohup ./api_server -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/api_server -predefined=false > ../logs/api_server/all.log 2>&1 &
sleep 1

echo "run wsserver ..."
LOG_DIR=../logs/wsserver nohup ./wsserver -conf=../configs/wsserver >> ../logs/wsserver/all.log 2>1 &
sleep 3

echo "run session ..."
HOSTNAME=session001 LOG_DIR=../logs/session ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:10013" nohup ./session -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/session -predefined=false > ../logs/session/all.log 2>&1 &
sleep 3

echo "run sync ..."
LOG_DIR=../logs/sync ETCD_ENDPOINTS=127.0.0.1:2379 HTTP_PERF="tcp://0.0.0.0:40503" nohup ./sync -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/sync -predefined=false > ../logs/sync/all.log 2>&1 &
sleep 1
