export XBLOG_RUNMODE=1
go run main.go
# 查看服务日志
#journalctl -u service-name.service
#清空日志 rm -rf /var/log/journal/