// 和设置相关的vuex管理

// 全局参数
import { requestProcess } from '@/utils/request-process-v3'
import {
  deletePlugins,
  getPluginsSetting,
  reloadPlugins,
  uploadPlugins,
  pluginRequest
} from '@/api/admin/plugins'

export const state = () => ({
  pluginOptions: []
})

// 属性获取
export const getters = {
  pluginOptions (state) { return state.pluginOptions }
}

// 数据修改
export const mutations = {
  setPluginsOption (state, data) { state.pluginOptions = data }
}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  uploadPlugins ({ commit }, data) { return requestProcess(uploadPlugins, none, none, data) },
  deletePlugins ({ commit }, id) { return requestProcess(deletePlugins, none, none, id) },
  reloadPlugins ({ commit }, id) { return requestProcess(reloadPlugins, none, none, id) },
  getPluginsSetting ({ commit }, id) { return requestProcess(getPluginsSetting, data => commit('setPluginsOption', data), none, id) },
  // 插件请求
  pluginRequest ({ commit }, data) { return requestProcess(pluginRequest, none, none, data) }
}
