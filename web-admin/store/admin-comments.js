// 和评论相关的vuex管理
import { requestProcess } from '@/utils/request-process-v3'
import { deleteComments, updateComments } from '@/api/admin/comments'

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
  // 更新文章状态
  updateComments ({ commit }, data) { return requestProcess(updateComments, none, none, data.id, data.data) },
  // 删除评论
  deleteComments ({ commit }, data) { return requestProcess(deleteComments, none, none, data) }
}
