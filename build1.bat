@echo off
echo 正在设置运行环境为linux...
go env -w CGO_ENABLED=0
go env -w GOOS=linux
go env -w GOARCH=amd64
echo 设置完成
echo 编译开始........
echo 正在编译databus,请勿关闭此窗口!!!
go build -o .\bin\databus  .\app\infra\databus\cmd
echo 正在编译剩下的文件,请勿强行关闭窗口!
echo 正在编译media,请勿关闭此窗口!!!
go build -o .\bin\media  .\app\service\media\cmd
echo 正在编译auth_session,请勿关闭此窗口!!!
go build -o .\bin\auth_session  .\app\service\auth_session\cmd
echo 正在编译inbox,请勿关闭此窗口!!!
go build -o .\bin\inbox  .\app\messenger\msg\inbox\cmd
echo 正在编译sync,请勿关闭此窗口!!!
go build -o .\bin\sync  .\app\messenger\sync\cmd
echo 正在编译push,请勿关闭此窗口!!!
go build -o .\bin\push  .\app\messenger\push\cmd
echo 正在编译webpage,请勿关闭此窗口!!!
go build -o .\bin\webpage  .\app\messenger\webpage\cmd
echo 正在编译gif,请勿关闭此窗口!!!
go build -o .\bin\gif  .\app\bots\gif\cmd
echo 正在编译admin_log,请勿关闭此窗口!!!
go build -o .\bin\admin_log  .\app\job\admin_log\cmd
echo 正在编译scheduled,请勿关闭此窗口!!!
go build -o .\bin\scheduled  .\app\job\scheduled\cmd 
echo 正在编译session,请勿关闭此窗口!!!
go build -o .\bin\session  .\app\interface\session\cmd
echo 正在编译botway,请勿关闭此窗口!!!
go build -o .\bin\botway  .\app\interface\botway\cmd
echo 正在编译api_server,请勿关闭此窗口!!!
go build -o .\bin\api_server  .\app\admin\api_server\cmd 
echo 正在编译gateway,请勿关闭此窗口!!!
go build -o .\bin\gateway  .\app\interface\gateway\cmd
echo 正在编译botfather,请勿关闭此窗口!!!
go build -o .\bin\botfather  .\app\bots\botfather\cmd
echo 正在编译msg,请勿关闭此窗口!!!
go build -o .\bin\msg  .\app\messenger\msg\msg\cmd
echo 正在编译relay,请勿关闭此窗口!!!
go build -o .\bin\relay  .\app\interface\relay\cmd
echo 正在编译wsserver,请勿关闭此窗口!!!
go build -o .\bin\wsserver  .\app\interface\wsserver\cmd
echo 正在编译biz_server,请勿关闭此窗口!!!
go build -o .\bin\biz_server  .\app\messenger\biz_server\cmd
echo 编译成功
echo 所有编译好的文件存放在bin目录
echo 正在还原运行环境
go env -w CGO_ENABLED=1
go env -w GOOS=windows
go env -w GOARCH=amd64
echo 设置完成
echo 所有步骤已完毕!
pause