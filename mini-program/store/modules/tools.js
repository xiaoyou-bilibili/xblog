import * as tool from  "../../config/api"
import requestProcess from "../../utils/request-process"
export const state = {

}
// 定义一个空函数
const none = _=>{}
// 微信小程序获取openid
export function getOpenid(data){return requestProcess(tool.toolGetOpenid,none,none,data)}
// 提交意见反馈
export function submitAdvice(data){return requestProcess(tool.toolSubmitAdvice,none,none,data)}

