// 数据可视化状态管理

// 全局参数
import { requestProcess } from '@/utils/request-process-v3'
import { getPostDetail, getPostDistributed, getTotal } from '@/api/admin/visual'

export const state = () => ({
  total: {
    post: 0, // 文章总数
    user: 0, // 用户总数
    view: 0, // 浏览量
    comment: 0 // 评论数
  }
})

// 属性获取
export const getters = {
  total (state) { return state.total }
}

// 数据修改
export const mutations = {
  setTotal (state, data) { state.total = data }
}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  // 获取统计信息
  getTotal ({ commit }) { return requestProcess(getTotal, res => commit('setTotal', res)) },
  // 获取文章分布信息
  getDistributed ({ commit }) { return requestProcess(getPostDistributed) },
  // 获取文章详细分布数据
  getPostDetail ({ commit }) { return requestProcess(getPostDetail) }
}
