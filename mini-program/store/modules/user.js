import * as user from  "../../config/api"
import requestProcess from "../../utils/request-process"
export const state = {
  collections: []
}

// 定义一个空函数
const none = _=>{}
// 获取用户收藏的文章
export function getCollection(data){
  return requestProcess(user.userGetCollection,data=>state.collections=data,none,data)
}
