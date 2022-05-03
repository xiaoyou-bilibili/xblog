// 全局变量
export default {
  install (Vue, options) {
    Vue.prototype.G = {
      // api服务地址
      SERVER: process.env.SERVER,
      // web服务地址
      LOCAL: process.env.LOCAL,
      // websocket服务地址
      WS: process.env.WS
    }
  }
}
