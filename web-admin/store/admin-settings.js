// 和设置相关的vuex管理

// 全局参数
import { requestProcess } from '@/utils/request-process-v3'
import {
  getSiteOption,
  updateOption,
  getBackgroundOption,
  getWechatOption,
  getAPPOption,
  getSideOption, updateSideOption, getNavOption, updateNavOption, rowRequest
} from '@/api/admin/settings'

export const state = () => ({
  siteOptions: [],
  backgroundOptions: [],
  wechatOption: [],
  sideOptions: [],
  navOptions: [],
  appOptions: []
})

// 属性获取
export const getters = {
  siteOption (state) { return state.siteOptions },
  backgroundOption (state) { return state.backgroundOptions },
  wechatOption (state) { return state.wechatOption },
  sideOptions (state) { return state.sideOptions },
  tableTreeData (state) { return state.navOptions },
  appOptions (state) { return state.appOptions }
}

// 数据修改
export const mutations = {
  setSiteOption (state, data) { state.siteOptions = data },
  setBackgroundOption (state, data) { state.backgroundOptions = data },
  setWechatOption (state, data) { state.wechatOption = data },
  setSideOptions (state, data) { state.sideOptions = data },
  setNavOptions (state, data) { state.navOptions = data },
  setAppOptions (state, data) { state.appOptions = data }
}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  // 获取站点设置
  getSiteOption ({ commit }) { return requestProcess(getSiteOption, data => commit('setSiteOption', data)) },
  // 更新站点设置
  updateOption ({ commit }, data) { return requestProcess(updateOption, none, none, data) },
  // 获取壁纸设置
  getBackgroundOption ({ commit }) { return requestProcess(getBackgroundOption, data => commit('setBackgroundOption', data)) },
  // 获取侧边栏设置
  getSideOption ({ commit }) { return requestProcess(getSideOption, data => commit('setSideOptions', data)) },
  // 更新站点信息
  updateSideOption ({ commit }, data) { return requestProcess(updateSideOption, none, none, data) },
  // 获取导航栏设置
  getNavOption ({ commit }) { return requestProcess(getNavOption, data => commit('setNavOptions', data)) },
  // 更新导航栏设置
  updateNavOption ({ commit }, data) { return requestProcess(updateNavOption, none, none, data) },
  // 获取微信小程序设置
  getWechatOption ({ commit }) { return requestProcess(getWechatOption, data => commit('setWechatOption', data)) },
  // 获取手机APP设置
  getAPPOption ({ commit }) { return requestProcess(getAPPOption, data => commit('setAppOptions', data)) },
  // 发送原始请求
  sendRowRequest ({ commit }, data) { return requestProcess(rowRequest, none, none, data) }
}
