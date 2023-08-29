 REM cmd 以管理员形式打开命令行
 cd F:\feij\gagachat-chat
 f:
 REM 下面命令逐条执行
 go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64
 
go build -o .\bin\databus  .\app\infra\databus\cmd
start /b go build -o .\bin\media  .\app\service\media\cmd && go build -o .\bin\auth_session  .\app\service\auth_session\cmd && go build -o .\bin\inbox  .\app\messenger\msg\inbox\cmd
start /b go build -o .\bin\sync  .\app\messenger\sync\cmd && go build -o .\bin\push  .\app\messenger\push\cmd && go build -o .\bin\webpage  .\app\messenger\webpage\cmd
start /b go build -o .\bin\gif  .\app\bots\gif\cmd && go build -o .\bin\admin_log  .\app\job\admin_log\cmd && go build -o .\bin\scheduled  .\app\job\scheduled\cmd 
start /b go build -o .\bin\session  .\app\interface\session\cmd && go build -o .\bin\botway  .\app\interface\botway\cmd && go build -o .\bin\api_server  .\app\admin\api_server\cmd 
start /b go build -o .\bin\gateway  .\app\interface\gateway\cmd && go build -o .\bin\botfather  .\app\bots\botfather\cmd && go build -o .\bin\msg  .\app\messenger\msg\msg\cmd
go build -o .\bin\relay  .\app\interface\relay\cmd
go build -o .\bin\wsserver  .\app\interface\wsserver\cmd
go build -o .\bin\biz_server  .\app\messenger\biz_server\cmd



