// 用来发送request请求
import axios from 'axios'
import QS from 'qs'
import Cookie from 'js-cookie'
// 对axios函数进行封装，用来发api请求，post使用qs进行处理，避免自己把from数据转换为json字符串
export default async function requestsV2 (url, data, type) {
  // 默认数据
  const res = {
    code: 104,
    data: null,
    msg: '发送请求失败！'
  }
  try {
    // 在发送请求前自己手动给每个请求加上token等信息
    let auth = {
      user_id: 0,
      token: ''
    }
    try {
      // 先判断cookie是否存在
      if (Cookie.get('blog-token') !== undefined) {
        // 把cookie数据解析为对象
        auth = JSON.parse(Cookie.get('blog-token'))
      }
    } catch (e) {
      // 发生错误不管，继续发送请求
    }
    // 判断请求类型
    if (type === 'get') {
      return (await axios.get(url, { params: data, headers: auth })).data
    } else if (type === 'post') {
      return (await axios.post(url, QS.stringify(data), { headers: { 'content-type': 'application/x-www-form-urlencoded;charset=utf-8', token: auth.token, user_id: auth.user_id } })).data
    } else if (type === 'form') {
      return (await axios.post(url, data, { headers: auth })).data
    } else {
      return res
    }
  } catch (e) {
    return res
  }
}
