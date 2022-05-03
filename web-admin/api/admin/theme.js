// 管理员主题管理
import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'php/api/v1/'
const base2 = process.env.SERVER + 'api/v3/admin/settings/'

const themes = base + 'themes'
const themesServer = base2 + 'theme'

// 获取所有主题
export function getThemes () { return request(themes, null, 'get') }
// 启用某个主题
export function chooseThemes (name) { return request(themes + `?name=${name}`, null, 'put') }
// 获取主题的设置
export function getThemeSetting (id) { return request(themesServer + `/${id}`, null, 'get') }
