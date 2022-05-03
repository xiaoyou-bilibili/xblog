// 管理员用户管理接口
import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'api/v3/admin/users'

// 获取用户列表
export function getUserList (data) { return request(base, data, 'get') }
// 管理员更新用户状态
export function updateUser (id, data) { return request(base + `/${id}`, data, 'put') }
// 管理员删除用户
export function deleteUser (id) { return request(base + `/${id}`, null, 'delete') }
