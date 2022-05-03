// 封装微信的request请求
export default function request(url, type, data = null){
  // 发送请求
  return new Promise((resove,reject) => {
    wx.request({
      url: url,
      method: type,
      data: data,
      success(response) {
        // 获取响应码
        let code = response.statusCode
        // 如果响应码是2xx那么就说明请求成功
        if(code>=200 && code <=299){
          resove(response.data)
        }else{
          reject(response.data)
        }
      },
      fail(err){
        // 返回请求失败的内容
        reject({message:'请求发送失败'})
      }
    })
  })
}