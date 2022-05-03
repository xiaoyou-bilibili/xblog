// pages/me/home/home.js
const app = getApp()
import create from '../../../utils/create'
import store from '../../../store/index'
// 获取用户信息
import  {wechatGetInfo} from '../../../utils/tools'

create.Page(store, {
  data: {
    // 用户信息
    userInfo: {},
    // 用户是否登录
    hasUserInfo: false,
    // 设置数据
    setting: {}
  },
  onLoad: function () {
    // 获取用户信息
    wechatGetInfo().then(res=> {
      // 设置全局变量
      app.globalData.userInfo = res.userInfo
      // 设置我的界面内容显示
      this.setData({
        userInfo: res.userInfo,
        hasUserInfo: true
      })
    })
    // 设置头部数据等其他东西
    const setting = app.globalData.setting
    this.setData({setting})
  },
  // 用户点击的登录按钮
  getUserInfo: function(e) {
    app.globalData.userInfo = e.detail.userInfo
    // 判断是否获取到了个人信息
    if (e.detail.errMsg === 'getUserInfo:ok'){
      this.setData({
        userInfo: e.detail.userInfo,
        hasUserInfo: true
      })
    }
  },
  // 个人中心
  person(){
    wx.navigateTo({url: '/pages/me/person/person'})
  },
  // 我的收藏
  collect(){
    wx.navigateTo({url: '/pages/me/collection/collection'})
  },
  // 意见反馈
  advice(){
    wx.navigateTo({url: '/pages/me/advice/advice'})
  },
  // 关于作者
  about(){
    wx.showModal({
      title: '关于博主',
      content: this.data.setting.about,
      confirmText: '确定',
      showCancel: false,
      success (res) {
        // wx.switchTab({url: '/pages/me/home/home'})
      }
    })
  },
  // 赞助博主
  donate(){
    wx.previewImage({
      current: [this.data.setting.price],
      urls: [this.data.setting.price]
    })
  }
})