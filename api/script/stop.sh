#先过滤出2333端口，用grep过滤出main，同时用awx过滤出第二行的数据
id=$(lsof -i:2333|grep main|awk 'NR==1{print $2}')
if [ ! ${id} ]
then
  echo "项目未启动!"
else
  kill -9 $id
  echo "项目已关闭!"
fi