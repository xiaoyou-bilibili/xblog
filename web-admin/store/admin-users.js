// 和用相关的vuex管理
import { requestProcess } from '@/utils/request-process-v3'
import { deleteUser, updateUser } from '@/api/admin/users'
import { deleteComments } from '@/api/admin/comments'

// 全局参数
export const state = () => ({
})

// 属性获取
export const getters = {
}

// 数据修改
export const mutations = {
}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  // 更新用户状态
  updateUser ({ commit }, data) { return requestProcess(updateUser, none, none, data.id, data.data) },
  // 删除用户
  deleteUser ({ commit }, data) { return requestProcess(deleteUser, none, none, data) }
}
