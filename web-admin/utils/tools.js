// 一些通用的工具路由
export default {
  install (Vue, options) {
    Vue.prototype.tools = {
      // 判断邮箱格式是否正确
      checkEmail (email) {
        const re = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/
        return re.test(email)
      }
    }
  }
}
