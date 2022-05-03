// 对request请求进行处理 ... 表示可变长的参数
export default function requestProcess(request, okHandle = () => {}, errHandle = () => {},...data){
  return new Promise((resolve,reject) => {
    // ... 这里会把data解析为参数
    request(...data)
    .then((res) => {okHandle(res);resolve(res)})
    .catch(res=>{errHandle(res);reject(res.message)})
  })
}