// 引入post的 api
import {
  getSettingAdmin
} from '@/api/setting'
import { requestProcess } from '@/utils/request-process-v3'

// 全局参数
export const state = () => ({
  // 后台管理系统设置
  settingAdmin: {
    icon: '', // 图标
    title: '', // 标题
    version: '' // 版本信息
  }
})

// 属性获取
export const getters = {
  settingAdmin (state) { return state.settingAdmin }
}

// 数据修改
export const mutations = {
  setSettingAdmin (state, data) { state.settingAdmin = data }
}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  // 获取后台管理界面的设置
  getAdminSetting ({ commit }, data) { return requestProcess(getSettingAdmin, res => commit('setSettingAdmin', res), none, data) }
}
