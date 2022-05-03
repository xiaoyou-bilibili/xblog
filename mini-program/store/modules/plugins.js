import * as api from  "../../config/api"
import requestProcess from "../../utils/request-process"
export const state = {
  friends: []
}

// 定义一个空函数
const none = _=>{}
// 获取日记
export function getDiary(data){return requestProcess(api.moreGetDiary,none,none,data)}
// 获取我的追番
export function getAnimation(data){return requestProcess(api.moreGetAnimation,none,none,data)}
// 获取友链数据
export function getFriend(data){return requestProcess(api.moreGetFriend,data => state.friends=data,none,data)}

// 用户提交友链
export function submitFriend(data){return requestProcess(api.moreSubmitFriend,none,none,data)}
// 获取赞助数据
export function getDonate(data){return requestProcess(api.moreGetDonate,none,none,data)}
// 获取豆瓣记录
export function getDouBan(type,data){return requestProcess(api.moreGetDouBan,none,none,type,data)}
