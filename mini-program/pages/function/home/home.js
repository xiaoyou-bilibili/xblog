// pages/index/home/home.js
import create from '../../../utils/create'
import store from '../../../store/index'
const app = getApp()
create.Page(store, {
  data: {
    functions:[]
  },
  onLoad: function (options) {
    // 这里我们根据微信小程序的设置自动显示相关的界面
    const setting = app.globalData.setting
    // console.log(setting)
    let tool = []
    if(setting.animation){
      tool.push({icon:'/images/icon/bilibili.png',text:'我的追番',page:'/pages/function/bilibili/bilibili'})
    }
    if(setting.friend){
      tool.push({icon:'/images/icon/friend.png',text:'友人帐',page:'/pages/function/friend/friend'})
    }
    if(setting.donate){
      tool.push({icon:'/images/icon/donate.png',text:'赞助博主',page:'/pages/function/donate/donate'})
    }
    if(setting.dou_ban){
      tool.push({icon:'/images/icon/douban.png',text:'我的豆瓣',page:'/pages/function/douban/douban'})
    }
    // tool.push({icon:'/images/icon/bilibili.png',text:'我的追番',page:'/pages/function/bilibili/bilibili'})
    // tool.push({icon:'/images/icon/friend.png',text:'友人帐',page:'/pages/function/friend/friend'})
    // tool.push({icon:'/images/icon/donate.png',text:'赞助博主',page:'/pages/function/donate/donate'})
    // tool.push({icon:'/images/icon/douban.png',text:'我的豆瓣',page:'/pages/function/douban/douban'})
    // 设置数据
    this.setData({functions:tool})
  },
  goto(e){
    wx.navigateTo({url: e.currentTarget.dataset.src})
  }
})