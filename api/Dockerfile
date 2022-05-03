# 上线阶段
FROM alpine

# 设置工作目录
WORKDIR /app/xblog-api

# 拷贝文件夹，因为我们设置了ignore文件，所以会只拷贝我们需要的内容
COPY  . .

# 给文件执行权限
RUN chmod +x ./build.sh && sh ./build.sh

# 每次启动容器时运行启动脚本
CMD chmod +x ./startup.sh && sh ./startup.sh
