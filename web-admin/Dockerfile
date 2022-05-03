FROM node:14.14.0-alpine3.10

# 指定工作目录
WORKDIR /app/xblog-web

# 复制当前路径文件
COPY . .

# 先解压文件，然后系统自动进行构建
RUN chmod +x ./build.sh && sh ./build.sh

# 每次启动容器时自动启动web项目
CMD  chmod +x ./startup.sh && sh ./startup.sh
