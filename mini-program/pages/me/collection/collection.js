// pages/functions/love/love.js
import create from '../../../utils/create'
import store from '../../../store/index'
// 引入自己封装好的函数
// import {Dialog,SetClipboard} from '../../../utils/tools'

create.Page(store, {
  // 声明依赖(computed需要用上)
  use: ['user'],
  // 计算属性
  computed: {
    collections() { return this.user.collections }
  },
  data: {
    hascollect:false//用户是否有收藏
  },
  onLoad: function (options) {
    // 获取我的收藏
    this.store.user.getCollection(wx.getStorageSync('openid')).then(res=>console.log('你好',res))
  },
  // 点击文章卡片显示文章
  postDetail(e) {
    const id=e.currentTarget.dataset.id;
    const img=e.currentTarget.dataset.img;
    wx.navigateTo({url: '/pages/index/post/post?id='+id+'&img='+img})
  }
})