// 函数的错误处理v3版本
// 这个函数用来处理api请求( okHandle 是请求发送成功的时候调用的函数，errorHandle 是请求发送失败的时候调用的函数，如果不传参数，那么就默认为空函数）
export function requestProcess (request, okHandle = () => {}, errHandle = () => {}, ...data) {
  // 初始化一个promise用来处理api请求
  return new Promise((resolve, reject) => {
    // 调用api函数,并传入数据
    request(...data).then((res) => {
      // 判断api函数的返回状态码
      if (res.status >= 200 && res.status < 300) {
        // 请求成功执行的操作
        okHandle(res.data)
        // 请求成功的回调
        resolve(res.data)
      } else {
        // 请求失败执行的操作
        errHandle(res.data.message)
        // 请求失败的回调
        reject(res.data.message)
      }
    }).catch((error) => {
      let message = '请求出现异常'
      // 这里说明是服务器异常
      if (error.response) {
        message = error.response.data.message
      } else if (error.request) {
        message = '发送请求失败'
      }
      // 请求失败执行的操作
      errHandle(message)
      // 请求失败的回调
      reject(message)
    })
  })
}
