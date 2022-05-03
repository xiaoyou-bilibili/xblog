// 管理员评论管理接口
import request from '@/utils/requests-v3'

// 基本路径
const base = process.env.SERVER + 'api/v3/admin/comments'

// 获取评论列表
export function getComments (data) { return request(base, data, 'get') }
// 管理员更新评论状态
export function updateComments (id, data) { return request(base + `/${id}`, data, 'put') }
// 管理员删除评论
export function deleteComments (id) { return request(base + `/${id}`, null, 'delete') }
