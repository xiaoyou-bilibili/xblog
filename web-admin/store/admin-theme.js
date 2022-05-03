// 和主题相关的vuex管理

// 全局参数
import { requestProcess } from '@/utils/request-process-v3'
import {
  chooseThemes,
  getThemeSetting
} from '@/api/admin/theme'

export const state = () => ({
  themeOptions: []
})

// 属性获取
export const getters = {
  themeOptions (state) { return state.themeOptions }
}

// 数据修改
export const mutations = {
  setThemeOption (state, data) { state.themeOptions = data }
}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  // 选择主题
  chooseTheme ({ commit }, data) { return requestProcess(chooseThemes, none, none, data) },
  // 获取主题设置
  getThemeSetting ({ commit }, data) { return requestProcess(getThemeSetting, data => commit('setThemeOption', data), none, data) }
}
