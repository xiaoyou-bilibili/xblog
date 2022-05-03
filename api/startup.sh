# 先判断资源文件夹是否有映射
if [ -z "$(ls -A assets)" ];
then
#  如果为空，那么就复制备份文件夹
   echo "正在复制资源文件夹..."
   cp -rf assets-back/* assets/
fi
# 判断配置文件夹是否为空
if [ -z "$(ls -A configs)" ];
then
#  如果为空，那么就复制备份文件夹
  echo "正在复制配置文件夹..."
   cp -rf configs-back/* configs/
fi

echo "正在启动api服务..."
./main
