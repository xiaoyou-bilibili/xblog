import { requestProcess } from '@/utils/request-process-v3'
import {
  getUserInfo,
  updateInfo,
  login,
  isLogin,
  userForget,
  userIsExist,
  userRegister
} from '@/api/user'
import Cookie from 'js-cookie'
import { use } from 'element-ui/src/locale'
// 全局参数
export const state = () => ({
  userInfo: { // 用户信息
    avatar: '', // 用户头像
    sign: '', // 个性签名
    level: 0, // 用户等级
    hang: '', // 头像挂件
    username: '', // 用户名
    nickname: '', // 昵称
    email: '', // 邮件
    user_id: 0, // 用户id
    identity: 0, // 用户身份
    subscription: false // 是否订阅邮件
  }
})

// 属性获取
export const getters = {
  userInfo (state) { return state.userInfo }
}

// 数据修改
export const mutations = {
  // 设置用户个人信息
  setUserInfo (state, data) { state.userInfo = data },
  // 用户登录操作
  setAuth (state, data) {
    // 设置cookie
    if (data.remember) {
      Cookie.set('token', data, { expires: 30 })
    } else {
      Cookie.set('token', data)
    }
  },
  // 清除认证信息
  clearAuth (state, data) {
    Cookie.remove('token')
  }
}
// 定义空函数
const none = () => {}
// 函数调用
export const actions = {
  // 获取用户个人信息
  getUserInfo ({ commit }) { return requestProcess(getUserInfo, data => commit('setUserInfo', data)) },
  // 更新用户个人信息
  updateInfo ({ commit }, data) { return requestProcess(updateInfo, none, none, data) },
  // 用户登录
  userLogin ({ commit }, data) { return requestProcess(login, (res) => { res.remember = data.remember; commit('setAuth', res) }, none, data) },
  // 用户是否登录
  userIsLogin ({ commit }) { return requestProcess(isLogin) },
  // 用户忘记密码
  forgetPassword ({ commit }, data) { return requestProcess(userForget, none, none, data) },
  // 用户的邮箱和密码是否存在
  userIsExist ({ commit }, data) { return requestProcess(userIsExist, none, none, data) },
  // 用户注册
  userRegister ({ commit }, data) { return requestProcess(userRegister, none, none, data) }
}
