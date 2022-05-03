#!/bin/bash
echo -e "\033[36m====[欢迎使用博客博客系统web端镜像自动构建脚本 V1.0]==== \033[0m"
# 更新grep
echo "正在更新grep..."
apk add --no-cache --upgrade grep
#echo "使用淘宝镜像..."
#npm config set registry https://registry.npm.taobao.org
# 开始安装
echo "正在安装依赖...."
npm install
# 构建项目
echo "正在构建项目...."
npm run build
echo -e "\033[36m 镜像构建完毕,感谢你使用本系统，祝你玩的愉快φ(>ω<*)。系统稍后会自动结束构建，请勿取消操作! \033[0m"
