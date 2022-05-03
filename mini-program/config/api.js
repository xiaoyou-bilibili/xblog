import request from "../utils/request"

// API地址
export const API = "https://xiaoyou66.com"
// WEB端地址
export const WEB = "https://xiaoyou66.com"
// 各个板块的基本地址
const postBase = API + '/api/v3/posts'
const pluginsBase = API + '/api/v3/plugins'
const settingBase = API + '/api/v3/settings'
const toolsBase = API + '/api/v3/tools'
const userBase  = API + '/api/v3/user'

export const wechatCode = id=>API + `/api/v3/posts/${id}/mini_program_code`

/*主页调用的api*/
// 获取文章列表
export function postGetPostList(data) { return request( postBase, 'get', data) }
// 获取文章内容
export function postGetContent(id) { return request( postBase + `/${id}`, 'get') }
// 获取加密文章的内容
export function postGetEncryption(id,data) { return request( postBase + `/${id}/encryption`, 'get',data) }
// 判断用户是否收藏文章
export function postGetPostCollection(id,openid) { return request( postBase + `/${id}/wechat_status/${openid}`, 'get') }
// 用户给文章收藏或点赞
export function postUpdateCollection(id,openid,data) { return request( postBase+`/${id}/wechat_status/${openid}`, 'put', data) }
// 获取文章评论
export function postGetComment(id) { return request( postBase + `/${id}/comments`, 'get') }
// 用户提交评论
export function postCommitComment(id,data) { return request( postBase + `/${id}/wechat_comments`, 'post', data) }



/*插件板块的api*/
// 获取我的追番
export function moreGetAnimation(data) { return request( pluginsBase + '/animations', 'get', data) }
// 获取我的友链
export function moreGetFriend(data) { return request( pluginsBase + '/friends', 'get', data) }
// 用户提交友链
export function moreSubmitFriend(data) { return request( pluginsBase + '/friends', 'post', data) }
// 获取用户赞助信息
export function moreGetDonate(data) { return request( pluginsBase + '/sponsors', 'get', data) }
// 获取豆瓣记录
export function moreGetDouBan(type,data) { return request( pluginsBase + `/dou_ban/${type}`, 'get', data) }
// 获取我的日记
export function moreGetDiary(data) { return request( pluginsBase + '/diary', 'get', data) }


/* 通用工具板块 */
export function toolGetOpenid(code) { return request( toolsBase + `/openid/${code}`, 'get') }
export function toolSubmitAdvice(data) { return request( toolsBase + '/advice', 'post', data) }


/*设置界面调用的api*/
// 获取友链界面的设置
export function settingGetFriend(data) { return request( settingBase + '/plugins/friends', 'get', data) }
// 获取赞助界面的设置
export function settingGetDonate(data) { return request( settingBase + '/plugins/sponsor', 'get', data) }
// 获取小程序界面设置
export function settingGetWechat(data) { return request( settingBase + '/wechat', 'get', data) }


// 获取我的收藏
export function userGetCollection(openid) { return request( userBase + `/wechat_collections/${openid}`, 'get') }
