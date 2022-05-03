# 后端源码

配置文件在 `api\configs\app.ini` 里。自己看情况修改

如果想本地启动需要先 `export XBLOG_RUNMODE=1` 才能正常启动，要不然会报错

## 自带的控制台

启动应用时添加 blog 参数即可

![](../images/2022-05-03-18-06-02.png)

## 文章导出功能

输入下面的命令即可导出文章为markdown文件

```bash
go run main.go export
```
