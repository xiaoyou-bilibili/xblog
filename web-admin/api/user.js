// 和用户信息相关的路由

import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'api/v3/user'

// 获取用户信息
export function getUserInfo () { return request(base, null, 'get') }
// 更新用户信息
export function updateInfo (data) { return request(base, data, 'put') }
// 用户登录
export function login (data) { return request(base + '/token', data, 'post') }
// 用户是否已经登录
export function isLogin (data) { return request(base + '/token', data, 'get') }
// 用户忘记密码邮件
export function userForget (data) { return request(base + '/password/email', data, 'post') }
// 用户或邮箱是否存在
export function userIsExist (data) { return request(base + '/username', data, 'get') }
// 用户注册
export function userRegister (data) { return request(base, data, 'post') }
