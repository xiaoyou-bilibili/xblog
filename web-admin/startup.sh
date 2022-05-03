#!/bin/bash
#export web=192.168.123.131:3000
#export api=192.168.123.131:2333
#export https=1
#export tao=0

# 配置文件
file="package.json"

# 构建代码
Build () {
  echo "检测到你的配置文件发生变化，正在修改配置文件...."
  # 自动宁进行/转义
  api=${api//\//\\\/}
  web=${web//\//\\\/}
  ws=${ws//\//\\\/}
  apiT=${apiT//\//\\\/}
  webT=${webT//\//\\\/}
  wsT=${wsT//\//\\\/}
  # 修改配置文件
  sed -i "s/${apiT}/${api}/g" ${file}
  sed -i "s/${webT}/${web}/g" ${file}
  sed -i "s/${wsT}/${ws}/g" ${file}
  # 是否需要设置淘宝镜像
  if [ "$tao" -eq 1 ]
  then
    echo "已设置为淘宝镜像..."
    npm config set registry https://registry.npm.taobao.org
  else
    echo "已设置为默认镜像..."
    npm config set registry https://registry.npmjs.org/
  fi
  echo "正在安装依赖...."
  npm install
  # 构建项目
  echo "正在构建项目...."
  npm run build
  Start
}

# 运行程序
Start () {
  # 启动项目
  echo "正在启动项目..."
  npm start
}

##########
# 主程序
##########

# 判断是否有https
if [ "$https" -eq 1 ]
then
  ws="wss://$api"
  api="https://$api"
  web="https://$web"
else
  ws="ws://$api"
  api="http://$api"
  web="http://$web"
fi

# 获取配置文件里面的地址
apiT=$( sed -n '7p' ${file} | grep -oP '(?<=SERVER=).*(?=\ L)')
webT=$( sed -n '7p' ${file} | grep -oP '(?<=LOCAL=).*(?=\ W)')
wsT=$( sed -n '7p' ${file} | grep -oP '(?<=WS=).*(?=\ n)')

# 判断是否需要重新编译
if [ "$apiT" == "$api" ]
then
  Start
else
  Build
fi
