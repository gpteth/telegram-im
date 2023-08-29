databus.exe -conf=../configs/databus/databus.toml  > ../logs/databus/all.log

gateway.exe  -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/gateway -predefined=false > ../logs/gateway/all.log

wsserver.exe -conf=../configs/wsserver >> ../logs/wsserver/all.log
auth_session.exe -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/auth_session > ../logs/auth_session/all.log

session.exe -app_name=e-chat -site_name=e-waychat.cn -conf=../configs/session -predefined=false > ../logs/session/all.log