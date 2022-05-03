// 管理员插件管理
import request from '@/utils/requests-v3'
import requestV2 from '@/utils/requests-v2'

// 基本路径
const base = process.env.SERVER + 'api/v3/admin/plugins'
const shop = process.env.SHOP + '/api/v2/blog/'

const manage = base + '/manage'
// plugin请求
export function pluginRequest (data) { return request(base + data.url, data.data, data.type) }

// 插件请求
export function shopRequest (data) { return requestV2(shop + data.url, data.data, data.type) }

// 获取所有插件
export function getAllPlugins () { return request(manage, null, 'get') }
// 上传新的插件
export function uploadPlugins (data) { return request(manage, data, 'post') }
// 删除插件
export function deletePlugins (id) { return request(`${manage}/${id}`, null, 'delete') }
// 重新加载插件
export function reloadPlugins (id) { return request(manage, null, 'put') }
// 获取插件的设置
export function getPluginsSetting (id) { return request(`${manage}/setting/${id}`, null, 'get') }
