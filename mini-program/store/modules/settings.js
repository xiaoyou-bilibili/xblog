import * as setting from  "../../config/api"
import requestProcess from "../../utils/request-process"
export const state = {
}

// 获取友链界面设置
export function getFriend(){return requestProcess(setting.settingGetFriend)}
// 获取赞助界面设置
export function getDonate(){return requestProcess(setting.settingGetDonate)}
// 获取微信小程序设置
export function getWechat(){return requestProcess(setting.settingGetWechat)}