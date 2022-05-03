// pages/functions/love/love.js
import create from '../../../utils/create'
import store from '../../../store/index'
// 引入自己封装好的函数
create.Page(store, {
  data: {
    hascollect: false//用户是否有收藏
  },
  onLoad: function (options) {

  },
  // 订阅 
  subscription(e) {
    console.log('订阅')
    wx.requestSubscribeMessage({
      tmplIds: ['ncUhOFtIqKQalVqq_zICwPxIfCAZ4McnSrPMwwxfwJM'],
      success (res) {
        console.log('订阅成功')
       }
    })
  }
})