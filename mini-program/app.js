//app.js 这里是小程序的主入口
import  {getOpenid} from './store/modules/tools'
import  {getWechat} from './store/modules/settings'
import  {wechatGetInfo} from './utils/tools'

App({
  globalData:{
    userInfo:null // 用户信息
  },
  onLaunch: function () {
    //自定义导航栏
    wx.getSystemInfo({
      success: e => {
        this.globalData.StatusBar = e.statusBarHeight;
        let custom = wx.getMenuButtonBoundingClientRect();
        this.globalData.Custom = custom;
        this.globalData.CustomBar = custom.bottom + custom.top - e.statusBarHeight;
      }
    })
    // 获取用户的openid
    wx.login({
      success: res => {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
        getOpenid(res.code).then(data=>{
          wx.setStorageSync("openid", data.openid)
          wx.setStorageSync("session_key", data.session_key)
        })
      }
    })
    // 获取微信小程序的设置
    getWechat().then(data=>{
      this.globalData.setting = data
    })
    // 主动获取用户信息
    wechatGetInfo().then(res=>{this.globalData.userInfo = res.userInfo})
  }
})