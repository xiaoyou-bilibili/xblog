// v3版本接口request方法
import axios from 'axios'
import Cookie from 'js-cookie'
// 对axios函数进行封装，用来发api请求，post使用qs进行处理，避免自己把from数据转换为json字符串
export default async function request (url, data, type) {
  // 默认数据
  const res = {
    status: 450,
    data: {
      message: '发送请求失败'
    }
  }
  // 在发送请求前自己手动给每个请求加上token等信息
  let auth = {
    user_id: 0,
    token: ''
  }
  try {
    // 先判断cookie是否存在
    if (Cookie.get('token') !== undefined) {
      // 把cookie数据解析为对象
      auth = JSON.parse(Cookie.get('token'))
    }
  } catch (e) {
    // 发生错误不管，继续发送请求
  }
  // 判断请求类型
  if (type === 'get') {
    return await axios.get(url, { params: data, headers: auth, timeout: 1000 * 60 * 10 })
  } else if (type === 'post') {
    return await axios.post(url, data, { headers: auth })
  } else if (type === 'put') {
    return await axios.put(url, data, { headers: auth })
  } else if (type === 'delete') {
    return await axios.delete(url, {
      params: data,
      headers: auth
    })
  } else if (type === 'patch') {
    return await axios.patch(url, data, { headers: auth })
  } else {
    return res
  }
}
