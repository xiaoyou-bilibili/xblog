// 各个页面的设置api
import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'api/v3/settings/'
const site = base + 'site'
// 获取登录界面的设置
export function getSettingAccess () { return request(site + '/login', null, 'get') }
// 获取后台管理界面的设置
export function getSettingAdmin (data) { return request(site + '/admin', data, 'get') }
